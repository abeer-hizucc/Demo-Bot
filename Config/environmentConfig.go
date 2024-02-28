package Config

import "os"

func Env() {
	err := os.Setenv("SLACK_BOT_TOKEN",
		"xoxb-5536501927799-6714633977250-3kXc0pDypIKVYMVdGgHYHlCc")
	if err != nil {
		return
	}
	os.Setenv("CHANNEL_ID", "C06LU1A87EJ")
}
