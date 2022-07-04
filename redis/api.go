package redis

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wwqdrh/httputil"
	"github.com/wwqdrh/logger"
)

type ConfStatus struct {
	httputil.Base
	Request  struct{}
	Response struct {
		Endpoints string `json:"endpoint"`
		Password  string `json:"password"`
	}
}

func (c ConfStatus) Run(ctx *gin.Context) {
	if err := c.DoReq(ctx, 0, &c.Request); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}

	c.Response.Endpoints = currConf.Endpoint
	c.Response.Password = currConf.Password
	c.DoRes(ctx, http.StatusOK, gin.H{
		"data": &c.Response,
	})
}

type Confverify struct {
	httputil.Base
	Request  struct{}
	Response struct{}
}

func (c Confverify) Run(ctx *gin.Context) {
	if err := c.DoReq(ctx, 0, &c.Request); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}

	if _, err := currConf.VerifyDriver(); err != nil {
		logger.DefaultLogger.Error(err.Error())
		c.DoRes(ctx, httputil.ParamInvalid, gin.H{"description": "fail"})
	} else {
		c.DoRes(ctx, httputil.ServerOK, gin.H{"description": "ok"})
	}
}

type Confupdate struct {
	httputil.Base
	Request struct {
		Endpoint string `json:"endpoint"` // a;b;c
		Password string `json:"password"`
	}
	Response struct{}
}

func (c Confupdate) Run(ctx *gin.Context) {
	if err := c.DoReq(ctx, httputil.JSON, &c.Request); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}

	if err := currConf.Update(&redisConf{
		Endpoint: strings.TrimSpace(c.Request.Endpoint),
		Password: strings.TrimSpace(c.Request.Password),
	}); err == nil {
		ctx.String(http.StatusOK, "更新成功")
	} else {
		ctx.String(http.StatusBadRequest, "参数错误")
	}
}

type KeyGet struct {
	httputil.Base
	Request struct {
		Key  string `form:"key"`
		Type uint8  `form:"type"` // 0 表示字符串
	}
	Response struct{}
}

func (c KeyGet) Run(ctx *gin.Context) {
	if err := c.DoReq(ctx, 0, &c.Request); err != nil {
		logger.DefaultLogger.Error(err.Error())
		c.DoRes(ctx, httputil.ParamInvalid, gin.H{"description": "参数校验失败"})
	}

	if _, err := currConf.VerifyDriver(); err != nil {
		logger.DefaultLogger.Error(err.Error())
		c.DoRes(ctx, httputil.ParamInvalid, gin.H{"description": "etcd-conf未配置成功"})
		return
	}

	res, err := currDriver.Get(c.Request.Key, dataTyp(c.Request.Type))
	if err != nil {
		logger.DefaultLogger.Error(err.Error())
		c.DoRes(ctx, httputil.ParamInvalid, gin.H{"description": err.Error(), "data": err.Error()})
	} else {
		c.DoRes(ctx, httputil.ServerOK, gin.H{"description": "ok", "data": res})
	}
}
