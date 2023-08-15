package router

import (
	"encoding/json"
	"github.com/NoahAmethyst/im/entity"
	"github.com/NoahAmethyst/im/manager"
	"github.com/NoahAmethyst/im/sender"
	"github.com/NoahAmethyst/im/utils/log"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
	"github.com/sirupsen/logrus"
	"time"
)

func RouteMessage(conn *websocket.NSConn, message []byte) error {
	defer func() {
		recover()
	}()

	var msg entity.Message

	if err := json.Unmarshal(message, &msg); err != nil {
		logrus.Error(err)
		return err
	}
	log.Info().Msgf("got message:%+v", msg)

	go func() {
		time.Sleep(500 * time.Millisecond)
		err := sender.SendMsg(conn.Conn, &entity.Message{Content: "hello"})
		if err != nil {
			logrus.Error(err)
		}
	}()

	return nil
}

func RouteConnected(conn *neffos.NSConn, message neffos.Message) {
	log.Info().Msgf("%s connected,message:%+v", conn.Conn.ID(), message)

}

func RouteConnect(conn *websocket.Conn) {
	log.Info().Msgf("%s connected", conn.ID())
	manager.ConnPool.AddConn(conn.ID(), conn)
}

func RouteDisConnect(conn *websocket.Conn) {
	log.Info().Msgf("%s disconnectede", conn.ID())
	manager.ConnPool.DelConn(conn.ID())
}
