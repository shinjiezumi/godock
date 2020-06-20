package main

import (
	"flag"
	"fmt"
	"godock/chat"
	"log"
	"net/http"
	"os"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
)

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse()
	gomniauth.SetSecurityKey(os.Getenv("OAUTH_CRYPT_KEY"))
	googleClientId := os.Getenv("OAUTH_CLIENT_ID_GOOGLE")
	googleSecret := os.Getenv("OAUTH_CLIENT_SECRET_GOOGLE")
	gomniauth.WithProviders(
		google.New(googleClientId, googleSecret, "http://localhost:8080/auth/callback/google"),
	)

	http.Handle("/", &chat.TemplateHandler{Filename: "index.html"})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})

	http.Handle("/login", &chat.TemplateHandler{Filename: "login.html"})
	http.HandleFunc("/auth/", chat.LoginHandler)

	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "auth",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		w.Header()["Location"] = []string{"/chat"}
		w.WriteHeader(http.StatusTemporaryRedirect)
	})

	// MustAuthヘルパーでラップすると認証必須なページとすることが出来る
	http.Handle("/chat", chat.MustAuth(&chat.TemplateHandler{Filename: "chat.html"}))

	//r := chat.NewRoom(chat.UseAuthAvatar) // OAuthのアバター
	//r := chat.NewRoom(chat.UseGravatar) // Gravatar
	r := chat.NewRoom(chat.UserFileSystemAvatar) // アップロードしたアバター

	// SetTracerで出力先を指定するとログが出力される。
	chat.SetTracer(r, os.Stdout)
	http.Handle("/chat/room", r)

	// アップロードフォーム
	http.Handle("/upload", &chat.TemplateHandler{Filename: "upload.html"})
	http.HandleFunc("/uploader", chat.UploaderHandler)

	http.Handle("/avatars/", http.StripPrefix("/avatars/", http.FileServer(http.Dir("./chat/avatars"))))

	// チャットルームを開始
	go r.Run()
	log.Println("Webサーバーを開始します。ポート", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
