package notify

import "fmt"

func SendError(name string, errmsg string, err error) {
	content := fmt.Sprintf("%s\n%s", errmsg, err)
	msg := &DiscordMessage{
		Username: "ERROR: " + name,
		Content:  content,
	}

	SendDiscordMessage(msg, Settings().ErrorWHName)
}
