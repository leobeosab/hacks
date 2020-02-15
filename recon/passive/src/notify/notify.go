package notify

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"passiverecon/models"
)

type DiscordMessage struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

func SendDiscordMessage(message *DiscordMessage, envVar string) {

	webhookurl := os.Getenv(envVar)
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

func NotifyDirBustResults(domain string, results *[]models.DirBustResult) {
	content := "Dirbust Results: \n"
	for _, r := range *results {
		content += r.Path + "\n"
	}

	msg := &DiscordMessage{
		Username: domain,
		Content:  content,
	}

	SendDiscordMessage(msg, Settings().LoggingWHName)
}

func NotifyUniqueDomains(target string, domains *[]models.Domain) {
	content := "Unique Domains Found: \n"
	for _, d := range *domains {
		content += d.Name + "\n"
	}

	msg := &DiscordMessage{
		Username: target,
		Content:  content,
	}

	SendDiscordMessage(msg, Settings().ScanWHName)
}
