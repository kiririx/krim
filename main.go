package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatalln(err)
		}
		defer conn.Close()
		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			log.Printf("从客户端接收到了消息：%s", message)
			log.Printf("从服务端写出了消息：%s", message)
			err = conn.WriteMessage(messageType, message)
			if err != nil {
				log.Println("error：", err)
				break
			}
		}
	})
	log.Fatalln(http.ListenAndServe(":19993", nil))
}
