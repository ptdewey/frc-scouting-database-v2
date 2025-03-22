package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
)

func runJob(s *discordgo.Session, channelID string) error {
	year := time.Now().Year()

	fmt.Println("Running scheduled job. Curent time:",
		time.Now().Format("2006-01-02 15:04:05"))

	// Remove old messages
	if _, err := deleteBotMessages(s, channelID); err != nil {
		log.Println("Failed to remove old messages.")
	}

	s.ChannelMessageSend(channelID,
		"Running scheduled job.\nCurrent time: "+
			time.Now().Format("Jan-02 03:04PM"))

	// Define data source directory and zip file target
	sourceDir := filepath.Join("output", fmt.Sprint(year))
	targetZip := filepath.Join("output", fmt.Sprint(year)+"_data.zip")

	// TODO: expand to accept multiple channel IDs?
	return zipToDiscord(s, channelID, sourceDir, targetZip)
}

// Function Start provides a forward-facing wrapper function for starting
// the discord bot application, as well as creating a cron job for
// automatic upload of output data.
// It takes in a discord bot token, a discord channel id which is used
// to specify the channel automated updates should be posted to, and
// a cron schedule string to define the job.
func Start(botToken string, channelID string, cronSchedule string) error {
	// create new discord session
	s, err := discordgo.New("Bot " + botToken)
	if err != nil {
		return err
	}

	// Create new cron job
	c := cron.New()

	// Add cron job handler function
	_, err = c.AddFunc(cronSchedule, func() {
		runJob(s, channelID)
	})
	if err != nil {
		return err
	}

	// Start cron scheduler
	c.Start()

	// Open connection to discord
	s.AddHandler(messageCreate)
	err = s.Open()
	if err != nil {
		return err
	}
	fmt.Println("Bot is now running. Press ctrl+c to exit.")

	// wait until term signal is received
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// stop cron scheduler
	c.Stop()

	// close discord session
	s.Close()

	return nil
}
