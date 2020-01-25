package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Message struct {
	Content string `json:"content"`
}

func main() {
	// Discord Notification WebHook
	whurl := os.Getenv("D_NOTIFICATION_WH")

	c := strings.Join(os.Args[1:], "\n")
	m := &Message{
		Content: c,
	}

	j, _ := json.Marshal(m)

	req, err := http.NewRequest("POST", whurl, bytes.NewBuffer(j))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
}
