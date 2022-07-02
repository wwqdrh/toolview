//go:build web

package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wwqdrh/httputil"
	"github.com/wwqdrh/logger"

	"github.com/wwqdrh/toolview/etcd"
)

var port = flag.Int("port", 8080, "端口")

//go:embed web/dist
var dist embed.FS

type Handler interface {
	httputil.Base
	Run(ctx *gin.Context)
}

var (
	Engine *gin.Engine

	FronentStaticPrefix    []string
	FronentStaticDirPrefix []string

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

	// 前端文件
	distFiles, err := fs.Sub(dist, "web/dist")
	if err != nil {
		logger.DefaultLogger.Fatal(err.Error())
	}

	files, err := fs.ReadDir(distFiles, ".")
	if err != nil {
		logger.DefaultLogger.Fatal(err.Error())
	}
	for _, item := range files {
		if !item.IsDir() {
			FronentStaticPrefix = append(FronentStaticPrefix, item.Name())
		} else {
			FronentStaticDirPrefix = append(FronentStaticDirPrefix, item.Name())
		}
	}

	f := http.FileServer(http.FS(distFiles))
	Engine.GET("/", func(ctx *gin.Context) {
		data, err := fs.ReadFile(distFiles, "index.html")
		if err != nil {
			logger.DefaultLogger.Error(err.Error())
			ctx.String(http.StatusInternalServerError, err.Error())
		}

		ctx.Data(http.StatusOK, "text/html", data)
	})
	for _, staticFile := range FronentStaticPrefix {
		Engine.HEAD("/"+staticFile, func(ctx *gin.Context) {
			f.ServeHTTP(ctx.Writer, ctx.Request)
		})
		Engine.GET("/"+staticFile, func(ctx *gin.Context) {
			f.ServeHTTP(ctx.Writer, ctx.Request)
		})
	}
	for _, staticDir := range FronentStaticDirPrefix {
		Engine.HEAD(path.Join(staticDir, "/*filepath"), func(ctx *gin.Context) {
			ctx.Request.URL.Path = strings.TrimPrefix(ctx.Request.URL.Path, staticDir)
			f.ServeHTTP(ctx.Writer, ctx.Request)
		})
		Engine.GET(path.Join(staticDir, "/*filepath"), func(ctx *gin.Context) {
			ctx.Request.URL.Path = strings.TrimPrefix(ctx.Request.URL.Path, staticDir)
			f.ServeHTTP(ctx.Writer, ctx.Request)
		})
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
