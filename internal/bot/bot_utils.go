package bot

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/ptdewey/frc-scouting-database-v2/internal/utils"
)

// Function messageCreate is the message creation handler for a discord
// client session. It reads messages and parses them for commands,
// responding when necessary.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// prevent recursive messages
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Check if message is help command
	if m.Content == ":Help" {
		s.ChannelMessageSend(m.ChannelID,
			"To get data for an event send ':EventsGet event_key' in a channel and replace event_key with a valid event key."+
				"(i.e. ':EventsGet 2024vagle' for the 2024 Glen Allen event)\n"+
				"To get data for all retrieved events from the current year run ':EventsGet all'.",
		)
	}

	// check if correct prefix was used
	if strings.HasPrefix(m.Content, ":EventsGet") {
		// slice message from first space to end (extract event key)
		i := strings.Index(m.Content, " ")
		if i == -1 {
			s.ChannelMessageSend(m.ChannelID,
				"Provide a valid event key (i.e. '2023vagle') or 'all' to get event statistics.")
			return
		}
		eventKey := m.Content[i+1:]

		// check if event key matches
		if eventKey == "all" {
			sourceDir := filepath.Join("output", fmt.Sprint(time.Now().Year()))
			s.ChannelMessageSend(m.ChannelID, "Getting data for all processed events")
			zipPath := filepath.Join(sourceDir, eventKey+".zip")
			zipToDiscord(s, m.ChannelID, sourceDir, zipPath)
		} else {
			// Define source directory for data (targets year directory)
			sourceDir := filepath.Join("output", m.Content[i+1:i+5])

			// read directories from data storage dir
			dirs, err := os.ReadDir(sourceDir)
			if err != nil {
				fmt.Println("Invalid data directory:", err)
				s.ChannelMessageSend(m.ChannelID, "Failed to read data directory.")
				return
			}

			// check if one of the directories matches the eventKey
			found := false
			for _, dir := range dirs {
				if dir.IsDir() && dir.Name() == eventKey {
					found = true
					break
				}
			}

			// no matching directory found
			if !found {
				s.ChannelMessageSend(m.ChannelID, "Event key does not match any existing directory.")
				return
			}

			// create zipfile and add to discord
			zipPath := filepath.Join(sourceDir, eventKey+".zip")
			fmt.Println(zipToDiscord(s, m.ChannelID, filepath.Join(sourceDir, eventKey), zipPath))
		}
	}
}

// Function zipToDiscord takes a source directory, wraps it in a zip file called targetZipPath, and
// uploads it to a specified discord channel.
func zipToDiscord(s *discordgo.Session, channelID string, sourceDir string, targetZipPath string) error {
	fmt.Println("Running Job...")
	err := utils.ZipDir(sourceDir, targetZipPath)
	if err != nil {
		return err
	}

	fmt.Println(channelID)
	err = uploadToDiscord(s, channelID, targetZipPath)
	if err != nil {
		return err
	}
	fmt.Println("File uploaded successfully.")

	// delete zip after sending
	err = os.Remove(targetZipPath)
	if err != nil {
		fmt.Println("Error deleting zip file:", err)
		return err
	} else {
		fmt.Println("Zip file deleted successfully")
	}

	return nil
}

// Function uploadToDiscord is a helper function that uploads a specified file
// to a specified discord channel.
func uploadToDiscord(s *discordgo.Session, channelID, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get file stats
	stats, err := file.Stat()
	if err != nil {
		return err
	}

	// Create message
	message := &discordgo.MessageSend{
		Files: []*discordgo.File{
			{
				Name:   stats.Name(),
				Reader: file,
			},
		},
	}

	// send message with attached file
	_, err = s.ChannelMessageSendComplex(channelID, message)
	return err
}

func deleteBotMessages(s *discordgo.Session, channelID string) (int, error) {
	messages, err := s.ChannelMessages(channelID, 25, "", "", "")
	if err != nil {
		return -1, err
	}

	var botMessageIDs []string
	for _, msg := range messages {
		if msg.Author != nil && msg.Author.Bot {
			botMessageIDs = append(botMessageIDs, msg.ID)
		}
	}

	for _, msgID := range botMessageIDs {
		err := s.ChannelMessageDelete(channelID, msgID)
		if err != nil {
			log.Println("Failed to delete message:", err)
		}
		time.Sleep(50 * time.Millisecond)
	}

	return len(messages), nil
}

// Note: this function is entirely blocking if there are any non-bot messages in the channel
// (Only use it to clean up messages manually)
func deleteAllBotMessages(s *discordgo.Session, channelID string) error {
	for {
		n, err := deleteBotMessages(s, channelID)
		if err != nil {
			log.Println("Failed to remove old messages.")
			return err
		}
		if n <= 0 {
			break
		}
	}

	return nil
}
