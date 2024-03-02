package bot

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
)

// TODO: expand to accept multiple channel IDs?
func Start(botToken string, channelID string, cronSchedule string) error {
    // create new discord session
    dg, err := discordgo.New("Bot " + botToken)
    if err != nil {
        return err
    }

    // Create new cron job
    c := cron.New()

    // Add cron job handler function
    _, err = c.AddFunc(cronSchedule, func() {
        year := time.Now().Year()

        fmt.Println("Running scheduled job. Curent time:",
            time.Now().Format("2006-01-02 15:04:05"))

        dg.ChannelMessageSend(channelID, 
            "Running scheduled job.\nCurrent time: " +
            time.Now().Format("Jan-02 03:04PM"))

        // Define data source directory and zip file target
        sourceDir := filepath.Join("output", fmt.Sprint(year))
        targetZip := filepath.Join("output", fmt.Sprint(year) + "_data.zip")

        zipToDiscord(dg, channelID, sourceDir, targetZip)
    })
    if err != nil {
        return err
    }

    // Start cron scheduler
    c.Start()

    // Open connection to discord
    dg.AddHandler(messageCreate)
    err = dg.Open()
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
    dg.Close()

    return nil
}
