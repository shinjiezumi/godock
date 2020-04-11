package chat

import (
	"github.com/gorilla/websocket"
	"time"
)

type client struct {
	// このクライアントのためのWebsocket
	socket *websocket.Conn
	// メッセージが送られるチャネル
	send chan *message
	// クライアントが参加しているチャットルーム
	room *room
	// ユーザーに関する情報
	userData map[string]interface{}
}

func (c *client) read() {
	for {
		var msg *message
		if err := c.socket.ReadJSON(&msg); err == nil {
			now := time.Now()
			msg.When = now.Format("2006/01/02 15:04:05")
			msg.Name = c.userData["name"].(string)
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
