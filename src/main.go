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
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Page not found"))
			return
		}
		fmt.Fprintf(w, "Hello!")
	})

	http.Handle("/login", &chat.TemplateHandler{Filename: "login.html"})
	http.HandleFunc("/auth/", chat.LoginHandler)

	// MustAuthヘルパーでラップすると認証必須なページとすることが出来る
	http.Handle("/chat", chat.MustAuth(&chat.TemplateHandler{Filename: "chat.html"}))
	r := chat.NewRoom()
	// SetTracerで出力先を指定するとログが出力される。
	chat.SetTracer(r, os.Stdout)
	http.Handle("/chat/room", r)

	// チャットルームを開始
	go r.Run()
	log.Println("Webサーバーを開始します。ポート", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
