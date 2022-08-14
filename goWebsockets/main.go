package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ = upgrader.Upgrade(w, r, nil)

		for {
			// read message
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			fmt.Println("%s sent: %s \n", conn.RemoteAddr(), string(msg))

			if err = conn.WriterMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.ServeAndListen(":8000", nil)
}
