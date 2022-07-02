package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wwqdrh/logger"
)

var port = flag.Int("port", 8080, "端口")

var (
	Engine *gin.Engine
)

func main() {
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: Engine,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.DefaultLogger.Error(err.Error())
		}
		logger.DefaultLogger.Info("服务退出...")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.DefaultLogger.Error(err.Error())
	}
}
