package sender

import (
	"encoding/json"
	"github.com/NoahAmethyst/im/entity"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
)

func SendMsg(conn *websocket.Conn, msg *entity.Message) error {
	if jsonMsg, err := json.Marshal(msg); err != nil {
		return err
	} else {
		conn.Write(neffos.Message{
			IsNative: true,
			Body:     jsonMsg,
		})
	}
	return nil
}
