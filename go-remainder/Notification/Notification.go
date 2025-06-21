package Notification

import (
	"fmt"
	"github.com/gen2brain/beeep"
)

func SendNotification(message string) {
	fmt.Println("Sending Notification at its scheduled time", message)
	err := beeep.Notify("Reminder", message, "assets/info.png")
	if err != nil {
		fmt.Println("Error sending notification:", err)
	} else {
		// Use the exported channel
		fmt.Println("Notification sent successfully")
	}
}
