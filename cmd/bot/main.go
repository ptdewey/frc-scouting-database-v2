package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/ptdewey/frc-scouting-database-v2/internal/bot"
)

// Function main is the driver function for the exporter bot.
// It creates a discord bot and cron scheduler for automated
// sending of messages.
func main() {
	// Load .env file
	err := godotenv.Load("config/bot.env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// Get token from env
	botToken := os.Getenv("DISCORD_BOT_TOKEN")
	if botToken == "" {
		fmt.Println("No bot token found in environment.")
		return
	}

	// Get Discord channel ID from environment
	channelID := os.Getenv("DISCORD_CHANNEL_ID")
	if channelID == "" {
		fmt.Println("No channel ID found in environment.")
		return
	}

	// Define cron job
	cronSchedule := "2 10-18 * * 6,0"
	// cronSchedule := "1 9-19 * * *"

	// Start bot
	bot.Start(botToken, channelID, cronSchedule)
}
