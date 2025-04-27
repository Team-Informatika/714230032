package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Incoming dari wa.my.id
type IncomingMessage struct {
	Message string `json:"message"`
	Number  string `json:"number"`
}

// Kirim pesan ke wa.my.id
type SendMessageRequest struct {
	ApiKey   string `json:"api_key"`
	DeviceID string `json:"device_id"`
	Number   string `json:"number"`
	Message  string `json:"message"`
}

// Ganti dengan API KEY kamu
const (
	apiKey   = ""
	deviceID = ""
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var msg IncomingMessage
		err := json.NewDecoder(r.Body).Decode(&msg)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		log.Printf("Pesan masuk dari %s: %s", msg.Number, msg.Message)

		// Simpan pesan masuk ke MongoDB
		saveMessage(msg.Number, msg.Message, "incoming")

		// Logika auto-reply
		var reply string
		switch msg.Message {
		case "halo":
			reply = "Halo! Ada yang bisa dibantu? 😊"
		case "siapa kamu?":
			reply = "Saya adalah bot WhatsApp Golang!"
		default:
			reply = "Maaf, saya tidak paham pesanmu 😅"
		}

		// Simpan pesan keluar ke MongoDB
		saveMessage(msg.Number, reply, "outgoing")

		// Kirim balasan ke wa.my.id
		go sendReply(msg.Number, reply)
		fmt.Fprintln(w, "Pesan diterima")
	}
}

func sendReply(number string, message string) {
	payload := SendMessageRequest{
		ApiKey:   apiKey,
		DeviceID: deviceID,
		Number:   number,
		Message:  message,
	}

	payloadBytes, _ := json.Marshal(payload)
	resp, err := http.Post("https://wa.my.id/api/send-message", "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Printf("Gagal kirim balasan: %v", err)
		return
	}
	defer resp.Body.Close()
	log.Printf("Balasan terkirim ke %s", number)
}

func saveMessage(number, message, direction string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newMessage := Message{
		Number:    number,
		Message:   message,
		Direction: direction,
		Timestamp: time.Now().Unix(),
	}

	_, err := MessageCollection.InsertOne(ctx, newMessage)
	if err != nil {
		log.Printf("Gagal menyimpan pesan: %v", err)
	}
}
