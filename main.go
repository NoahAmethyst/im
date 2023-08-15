package main

import (
	"context"
	"fmt"
	"github.com/NoahAmethyst/im/router"
	"github.com/NoahAmethyst/im/utils/log"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
	"os"
	"time"
)

func main() {
	time.Local = time.FixedZone("Beijing Time", int((8 * time.Hour).Seconds()))
	ctx := context.Background()

	app := iris.New()

	app.Use(Cors)

	wsServer(app)

	gracefulShutdown(ctx, app)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	log.Info().Msgf("server listening in :%s", addr)
	_ = app.Listen(addr, iris.WithConfiguration(iris.Configuration{
		DisableInterruptHandler: true,
	}))
}

func wsServer(app *iris.Application) {
	ws := websocket.New(websocket.DefaultGobwasUpgrader, websocket.Events{
		websocket.OnNativeMessage: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			log.Info().Msgf("Server got: %s from [%s]", msg.Body, nsConn.Conn.ID())
			err := router.RouteMessage(nsConn, msg.Body)
			return err
		},
	})

	ws.OnConnect = func(c *websocket.Conn) error {
		router.RouteConnect(c)
		return nil
	}

	ws.OnDisconnect = func(c *websocket.Conn) {
		router.RouteDisConnect(c)
	}
	app.Get("/ws", websocket.Handler(ws))
}

func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Request().Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "*")
		ctx.Header("Access-Control-Allow-Headers", "Content-TileId, Accept, Authorization")
		ctx.StatusCode(204)
		return
	}
	ctx.Next()
}
