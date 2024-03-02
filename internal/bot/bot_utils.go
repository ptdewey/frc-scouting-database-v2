package bot

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"

	"github.com/bwmarrin/discordgo"
	"github.com/ptdewey/frc-scouting-database-v2/internal/utils"
)


// discord message creation handler
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
    // prevent recursive messages
    if m.Author.ID == s.State.User.ID {
        return
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
        eventKey := m.Content[i + 1:]

        // Define source directory for data (targets year directory)
        sourceDir := filepath.Join("output", m.Content[i + 1:i + 5])

        // check if event key matches
        if eventKey == "all" {
            s.ChannelMessageSend(m.ChannelID, "Getting data for all processed events")
            zipFile := filepath.Join(sourceDir, eventKey + ".zip")
            zipToDiscord(s, m.ChannelID, sourceDir, zipFile)
        } else {
            found := false

            // read directories from data storage dir
            dirs, err := os.ReadDir(sourceDir)
            if err != nil {
                fmt.Println("Invalid data directory:", err)
                s.ChannelMessageSend(m.ChannelID, "Failed to read data directory.")
                return
            }
            
            // check if one of the directories matches the eventKey
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
            zipPath := filepath.Join(sourceDir, eventKey + ".zip")
            fmt.Println(zipToDiscord(s, m.ChannelID, filepath.Join(sourceDir, eventKey), zipPath))
        }
    }
}


// zip source dir, send to discord in specified channel
// TODO: docs
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


// helper function that uploads a file to discord to a specific channel
// TODO: better docs
func uploadToDiscord(s *discordgo.Session, channelID, filePath string) error {
    file, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer file.Close()


    // get files stats
    stats, err := file.Stat()
    if err != nil {
        return err
    }

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

