//go:build api

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wwqdrh/httputil"
	"github.com/wwqdrh/httputil/middleware"
	"github.com/wwqdrh/toolview/etcd"
	"github.com/wwqdrh/toolview/redis"
)

var corsMiddleware = middleware.GinAllowAll()

type Handler interface {
	httputil.Base
	Run(ctx *gin.Context)
}

var API = []struct {
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
	// redis
	{"GET", "/api/redis/conf/verify", redis.Confverify{Base: httputil.DefaultHandler}},
	{"GET", "/api/redis/conf/status", redis.ConfStatus{Base: httputil.DefaultHandler}},
	{"POST", "/api/redis/conf/update", redis.Confupdate{Base: httputil.DefaultHandler}},
	{"GET", "/api/redis/key/get", redis.KeyGet{Base: httputil.DefaultHandler}},
}

func init() {
	Engine = gin.Default()
	Engine.Use(corsMiddleware)
	for _, item := range API {
		Engine.Handle(item.method, item.url, item.handler.Run)
	}
}
