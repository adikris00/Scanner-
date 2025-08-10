package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func SendToTelegram(message string) {
	token := "BOT TOKEN DISINI"
	chatID := "CHATID DISINI"
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)

	payload := map[string]string{"chat_id": chatID, "text": message}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	resp, err := http.Post(url, "application/json", strings.NewReader(string(jsonPayload)))
	if err != nil {
		fmt.Println("Error sending message to Telegram:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Failed to send message to Telegram. Status code:", resp.StatusCode)
	} else {
		fmt.Println("Message sent to Telegram successfully")
	}
}
