package wsbus

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/kiririx/krim/conf"
	"github.com/kiririx/krim/ctx"
	"github.com/kiririx/krim/repo/model"
	"github.com/kiririx/krutils/jsonx"
	"github.com/kiririx/krutils/strx"
	"github.com/kiririx/krutils/sugar"
	"log"
	"time"
)

func ReceiveMessage(ctx context.Context, conn *websocket.Conn) {
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
//
// msg_time: 如果消息时间是今天，那么不发送日期，如果消息时间不是今天，那么包含日期
func SendUserMessage(ctx *ctx.Ctx, sourceId, targetId uint64, message string, datetime time.Time) error {
	tc := ClientMap[strx.ToStr(targetId)]
	if tc != nil {
		fmtMsg, err := jsonx.Map2JSON(map[string]any{
			"sender_name": ctx.UserName,
			"sender_nick": ctx.NickName,
			"msg_time": func() string {
				return sugar.ThenFunc(datetime.Day() < time.Now().Day(), func() string {
					return strx.TimeToStr(datetime, "yyyy/MM/dd HH:mm")
				}, func() string {
					return strx.TimeToStr(datetime, "HH:mm")
				})
			}(),
			"message":  message,
			"msg_type": "message",
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
