package OfficeNotification

import (
	_ "fmt"
	"github.com/robfig/cron/v3"
	"log"
	_ "os"
)

func scheduleNotifications() {
	c := cron.New(cron.WithSeconds())

	// Schedule the start of the office time at 9:00 PM on weekdays
	_, err := c.AddFunc("00 00 09 * * 0-6", func() {
		sendNotification("Office time has started!")
	})
	if err != nil {
		log.Fatal("Failed to schedule start of office time:", err)
	}

	// Schedule the end of the office time at 4:30 PM on weekdays
	_, err = c.AddFunc("00 30 16 * * 0-6", func() {
		sendNotification("Office time has ended!")
	})
	if err != nil {
		log.Fatal("Failed to schedule end of office time:", err)
	}

	c.Start()
}
