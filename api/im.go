package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/kiririx/krim/service"
	"github.com/kiririx/krim/util/callback"
	"github.com/kiririx/krim/wsbus"
	"github.com/kiririx/krutils/httpx"
	"github.com/kiririx/krutils/strx"
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

// Conn uses the token to create a websocket connection
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
	userId := strx.ToStr(userMeta.Id)

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusOK, callback.Error(0, err.Error()))
		return
	}
	defer conn.Close()

	// register client to client map, if the client is exists, then override it
	if wsbus.ClientMap[userId] != nil {
		_conn := wsbus.ClientMap[userId]
		_conn.Conn.Close()
	}
	wsbus.ClientMap[userId] = &wsbus.WsClient{Conn: conn, MsgChan: make(chan *wsbus.WsMessage)}
	// go func() {
	// 	for {
	// 		wsMsg := <-wsbus.ClientMap[username].MsgChan
	// 		if wsMsg != nil {
	// 			err := wsbus.SendUserMessage(wsMsg.TargetId, wsMsg.Message)
	// 			if err != nil {
	// 				// return
	// 				continue
	// 			}
	// 		}
	//
	// 	}
	// }()
	ctx, _ := context.WithCancel(c)
	wsbus.ReceiveMessage(ctx, conn)
}
