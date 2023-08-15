package main

import (
	"context"
	"github.com/NoahAmethyst/im/manager"
	"github.com/NoahAmethyst/im/utils/log"
	"github.com/kataras/iris/v12"
	"os"
	"os/signal"
	"syscall"
)

// gracefulShutdown waits for termination syscalls and doing clean up operations after received it
func gracefulShutdown(ctx context.Context, app *iris.Application) {
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGSTOP, syscall.SIGKILL, syscall.SIGHUP)
	go func() {
		sig := <-signalChannel
		defer close(signalChannel)
		log.Info().Msgf("receive signal:%+v,graceful shutdown", sig)

		connPool := manager.ConnPool.AllConn()
		log.Info().Msgf("clear connection pool,size:%d", len(connPool))
		for _id, _conn := range connPool {
			_conn.Close()
			manager.ConnPool.DelConn(_id)
		}
		log.Info().Msgf("shutdown wsServer")

		_ = app.Shutdown(ctx)

		log.Info().Msgf("graceful shutdown done")
	}()
}
