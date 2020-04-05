package main

import (
	"fmt"
	"log"
	"net/http"
	"shinjiezumi.com/godock/chat"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	http.HandleFunc("/chat", chat.Main)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
