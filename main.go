package main

import "Demo-Bot/OfficeNotification"

func main() {
	OfficeNotification.ScheduleNotifications()

	select {}
}
