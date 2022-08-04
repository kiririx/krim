package wsbus

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/kiririx/krim/conf"
	"github.com/kiririx/krim/ctx"
	"github.com/kiririx/krim/model"
	"github.com/kiririx/krutils/jsonx"
	"github.com/kiririx/krutils/strx"
	"log"
)

func ReceiveMessage(ctx context.Context, conn *websocket.Conn) {
	for {
		if _, ok := <-ctx.Done(); !ok {
			return
		}
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
			err = ClientMap[targetUsername].Conn.WriteMessage(messageType, []byte(msgM["msg"].(string)))
			if err != nil {
				log.Println("error：", err)
				break
			}
		}
	}
}

// SendUserMessage 发送一对一的消息，先把消息保存到数据库，然后直接通过websocket发送。
//
// 如果消息发送失败，也不返回error，保证对方在离线状态下依然可用
func SendUserMessage(ctx *ctx.Ctx, sourceId, targetId uint64, message string) error {
	// // save the message
	// msgModel := model.Message{
	// 	SourceId: sourceId,
	// 	TargetId: targetId,
	// 	Msg:      message,
	// }
	// err := ctx.SqlCtl().Save(&msgModel).Error
	// if err != nil {
	// 	return err
	// }
	// send the message to peer
	tc := ClientMap[strx.ToStr(targetId)]
	if tc != nil {
		fmtMsg, err := jsonx.Map2JSON(map[string]any{
			"sender_name": ctx.UserName,
			"message":     message,
			"msg_type":    "message",
		})
		if err != nil {
			return err
		}
		err = tc.Conn.WriteMessage(websocket.TextMessage, []byte(fmtMsg))
		if err != nil {
			log.Println("error：", err)
			return err
		}
	}
	return nil
}

func SendGroupMessage(ctx context.Context, sourceId, targetId uint64, message string) error {
	msgModel := model.Message{
		SourceId: sourceId,
		GroupId:  targetId,
		Msg:      message,
	}
	err := conf.Sqlx.Save(&msgModel).Error
	if err != nil {
		return err
	}
	return nil
}
