package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	err := ConnectDB()
	if err != nil {
		log.Fatalf("Koneksi MongoDB gagal: %v", err)
	}

	http.HandleFunc("/webhook", webhookHandler)

	fmt.Println("Server berjalan di port 8080 🚀")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
