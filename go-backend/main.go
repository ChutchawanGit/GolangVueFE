package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Message struct สำหรับจัดการข้อมูล
type Message struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

// ข้อมูลตัวอย่าง
var messages = []Message{
	{ID: 1, Text: "Hello from Golang Backend!"},
	{ID: 2, Text: "This is an API message"},
}

// Handler สำหรับดึงข้อมูล
func getMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func main() {
	// สร้าง Router
	r := mux.NewRouter()
	r.HandleFunc("/api/messages", getMessages).Methods("GET")

	// ใช้ CORS Middleware
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8081"}, // อนุญาตเฉพาะโดเมนนี้
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	// รัน Server
	http.ListenAndServe(":8080", corsHandler.Handler(r))
}
