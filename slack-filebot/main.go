package main

import (
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channel := []string{os.Getenv("CHANNEL_ID")}
	file := []string{""}

	for _, filePath := range file {

		params := slack.FileUploadParameters{
			Channels: channel,
			File:     filePath,
		}

		_, err := api.UploadFile(params)

		if err != nil {
			log.Printf("Error uploading file: %s", err)
			return
		}

		log.Printf("Successfully uploaded file: %s", filePath)
	}
}
