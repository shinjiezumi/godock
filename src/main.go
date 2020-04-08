package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"shinjiezumi.com/godock/chat"
)

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse()
	log.Println(*addr)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	http.Handle("/chat", &chat.TemplateHandler{Filename: "chat.html"})
	r := chat.NewRoom()
	http.Handle("/chat/room", r)
	// チャットルームを開始
	go r.Run()
	log.Println("Webサーバーを開始します。ポート", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
