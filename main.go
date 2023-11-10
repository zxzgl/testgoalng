package main

import (
	"aaaaa/chat"
	"log"
	"net/http"
)

func main() {

	// websocket server
	server := chat.NewServer("/entry")
	go server.Listen()

	// static files
	http.Handle("/", http.FileServer(http.Dir("webroot")))

	log.Fatal(http.ListenAndServe(":80", nil))
}
