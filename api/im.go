package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/kiririx/krim/service"
	"github.com/kiririx/krim/util/callback"
	"github.com/kiririx/krutils/httpx"
	"github.com/kiririx/krutils/jsonx"
	"log"
	"net/http"
)

var Im = &im{}

type im struct {
}

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

func (i *im) Conn(c *gin.Context) {
	token := httpx.GetQueryParam(c.Request.RequestURI, "token")
	if token == "" {
		c.JSON(http.StatusOK, callback.BackFail("非法登录"))
		return
	}
	userMeta, err := service.ValidToken(token)
	if err != nil {
		c.JSON(http.StatusOK, callback.Error(0, err.Error()))
		return
	}
	username := userMeta.Username

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusOK, callback.Error(0, err.Error()))
		return
	}
	defer conn.Close()

	// 注册到客户端map
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
}
