package OfficeNotification

import (
	"github.com/slack-go/slack"
	"os"
)

func sendNotification(message string) {

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	//fileArr := []string{"invoicesample.pdf"}
	//
	//for _, file := range fileArr {
	//	params := slack.FileUploadParameters{
	//		Channels: channelArr,
	//		File:     file,
	//		Title:    message,
	//	}
	//	file, err := api.UploadFile(params)
	//	if err != nil {
	//		log.Fatal("%s\n", err)
	//		return
	//	}
	//	fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	//}
	_, _, err := api.PostMessage(channelArr[0], slack.MsgOptionText(message, false))
	if err != nil {
		return
	}
}
