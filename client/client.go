package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"strconv"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go mainExec("客户端" + strconv.Itoa(i))
	}
	time.Sleep(time.Hour)
}

func mainExec(client string) {
	u := url.URL{Scheme: "ws", Host: ":19993", Path: "/"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	done := make(chan struct{})
	// 读的协程
	go func() {
		defer close(done)
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			log.Println(client+"::从服务端接收到了：", string(msg))
		}
	}()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			err := conn.WriteMessage(websocket.TextMessage, []byte("（"+client+"你好"+"）"))
			if err != nil {
				log.Println(err)
			}
		}

	}
}
