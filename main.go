package main

import (
	"github.com/gorilla/websocket"
	"github.com/kiririx/krutils/httpx"
	"github.com/kiririx/krutils/jsonx"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WsClient struct {
	conn *websocket.Conn
	// uid  string
}

var ClientMap = make(map[string]*WsClient)

func main() {
	_ = httpx.Client()
	http.HandleFunc("/im", func(w http.ResponseWriter, r *http.Request) {
		username := httpx.GetQueryParam(r.RequestURI, "username")
		if username == "" {
			return
		}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatalln(err)
		}
		defer conn.Close()
		if ClientMap[username] != nil {
			_conn := ClientMap[username]
			_conn.conn.Close()
		}
		ClientMap[username] = &WsClient{conn: conn}
		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			msgM, err := jsonx.JSON2Map(string(message))
			if err != nil {
				log.Println(err)
				continue
			}
			targetUsername := msgM["targetId"].(string)
			log.Printf("从客户端接收到了消息：%s", msgM)
			log.Printf("从服务端写出了消息：%s", message)
			tc := ClientMap[targetUsername]
			if tc != nil {
				err = ClientMap[targetUsername].conn.WriteMessage(messageType, []byte(msgM["msg"].(string)))
				if err != nil {
					log.Println("error：", err)
					break
				}
			}
		}
	})
	log.Fatalln(http.ListenAndServe(":19993", nil))
}
