package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"

	"passiverecon/models"
)

type DiscordMessage struct {
	Username string         `json:"username"`
	Content  string         `json:"content"`
	Embeds   []DiscordEmbed `json:"embeds"`
}

type DiscordEmbed struct {
}

func SendDiscordMessage(message *DiscordMessage, envVar string) {

	webhookurl := os.Getenv(envVar)
	if webhookurl == "" {
		log.Println("Error sending discord notification, invalid webhookurl")
		return
	}

	if len(message.Content) > 1500 {
		for _, m := range LongMessageToMessageArr(message) {
			SendDiscordMessage(&m, envVar)
		}
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

func LongMessageToMessageArr(msg *DiscordMessage) []DiscordMessage {
	msgA := SplitN(msg.Content, 1500)
	msgs := make([]DiscordMessage, len(msgA))

	for i := 0; i < len(msgA); i++ {
		usr := fmt.Sprintf("%s Part %d of %d", msg.Username, i+1, len(msgA))
		tmp := DiscordMessage{
			Content:  msgA[i],
			Username: usr,
			Embeds:   msg.Embeds,
		}

		msgs[i] = tmp
	}

	return msgs
}

func NotifyDirBustResults(domain string, results *[]models.DirBustResult) {
	if len(*results) == 0 {
		return
	}

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

func NotifyDomains(target string, titleMessage string, domains *[]models.Domain, wh string) {
	if len(*domains) == 0 {
		return
	}

	content := titleMessage + "\n"
	for _, d := range *domains {
		content += d.Name + "\n"
	}

	msg := &DiscordMessage{
		Username: target,
		Content:  content,
	}

	SendDiscordMessage(msg, wh)
}

func SplitN(s string, n int) []string {
	f := float64(len(s)) / float64(n)
	length := int(math.Ceil(f))
	arr := make([]string, length)

	for i := 0; i < length; i++ {
		if i == length-1 {
			arr[i] = s[i*n:]
		} else {
			arr[i] = s[i*n : i*n+n]
		}
	}

	return arr
}
