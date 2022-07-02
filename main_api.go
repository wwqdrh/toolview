//go:build api

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
	"github.com/wwqdrh/httputil"
	"github.com/wwqdrh/logger"

	"github.com/wwqdrh/toolview/etcd"
)

var port = flag.Int("port", 8080, "端口")

type Handler interface {
	httputil.Base
	Run(ctx *gin.Context)
}

var (
	Engine *gin.Engine

	API = []struct {
		method  string
		url     string
		handler Handler
	}{
		{"GET", "/api/etcd/conf/verify", etcd.Confverify{Base: httputil.DefaultHandler}},
		{"GET", "/api/etcd/conf/status", etcd.Confstatus{Base: httputil.DefaultHandler}},
		{"POST", "/api/etcd/conf/update", etcd.Confupdate{Base: httputil.DefaultHandler}},
		{"GET", "/api/etcd/key/list", etcd.Keylist{Base: httputil.DefaultHandler}},
		{"POST", "/api/etcd/key/put", etcd.Keyput{Base: httputil.DefaultHandler}},
		{"POST", "/api/etcd/key/delete", etcd.Keydelete{Base: httputil.DefaultHandler}},
	}
)

func init() {
	Engine = gin.Default()
	for _, item := range API {
		Engine.Handle(item.method, item.url, item.handler.Run)
	}
}

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
