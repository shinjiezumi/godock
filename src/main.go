package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"shinjiezumi.com/godock/chat"
)

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	http.Handle("/chat", &chat.TemplateHandler{Filename: "chat.html"})
	r := chat.NewRoom()
	// SetTracerで出力先を指定するとログが出力される。
	chat.SetTracer(r, os.Stdout)
	http.Handle("/chat/room", r)
	// チャットルームを開始
	go r.Run()
	log.Println("Webサーバーを開始します。ポート", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
