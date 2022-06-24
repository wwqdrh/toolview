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
	"github.com/wwqdrh/logger"

	"github.com/wwqdrh/toolview/etcd"
)

var port = flag.Int("port", 8080, "端口")

//go:embed web/dist
var dist embed.FS

var (
	Engine *gin.Engine

	FronentStaticPrefix    []string
	FronentStaticDirPrefix []string

	API = []struct {
		method  string
		url     string
		handler []gin.HandlerFunc
	}{
		{"GET", "/api/etcd/conf/verify", []gin.HandlerFunc{etcd.ConfVerify}},
		{"GET", "/api/etcd/conf/status", []gin.HandlerFunc{etcd.ConfStatus}},
		{"POST", "/api/etcd/conf/update", []gin.HandlerFunc{etcd.ConfUpdate}},
		{"GET", "/api/etcd/key/list", []gin.HandlerFunc{etcd.KeyList}},
		{"POST", "/api/etcd/key/put", []gin.HandlerFunc{etcd.KeyPut}},
		{"POST", "/api/etcd/key/delete", []gin.HandlerFunc{etcd.KeyDelete}},
	}
)

func init() {
	Engine = gin.Default()
	for _, item := range API {
		Engine.Handle(item.method, item.url, item.handler...)
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
