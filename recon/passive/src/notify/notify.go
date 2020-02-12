package notify

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type DiscordMessage struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

func SendDiscordMessage(message *DiscordMessage) {

	webhookurl := os.Getenv("D_NOTIFICATION_WH")
	if webhookurl == "" {
		log.Println("Error sending discord notification, invalid webhookurl")
		return
	}

	j, err := json.Marshal(message)
	if err != nil {
		log.Println("Error marshalling message")
		return
	}

	req, err := http.NewRequest("POST", webhookurl, bytes.NewBuffer(j))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Println("Error creating request for Discord notification")
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending Discord notification")
		return
	}
	defer resp.Body.Close()
}
