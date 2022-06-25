package etcd

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wwqdrh/httputil"
	"github.com/wwqdrh/logger"
)

type Confstatus struct {
	httputil.Base
	Request  struct{}
	Response struct{}
}

func (c Confstatus) Run(ctx *gin.Context) {
	if err := c.DoReq(ctx, 0, &c.Request); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}

	c.DoRes(ctx, http.StatusOK, gin.H{
		"data": currConf,
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
		c.DoRes(ctx, httputil.ServerOK, gin.H{"description": "fail"})
	} else {
		c.DoRes(ctx, httputil.ServerOK, gin.H{"description": "ok"})
	}
}

type Confupdate struct {
	httputil.Base
	Request struct {
		Endpoints []string `json:"endpoints"`
		UserName  string   `json:"username"`
		Password  string   `json:"password"`
	}
	Response struct{}
}

func (c Confupdate) Run(ctx *gin.Context) {
	if err := c.DoReq(ctx, httputil.JSON, &c.Request); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}

	if err := currConf.Update(&etcdConf{
		Endpoints: func() []string {
			endpoints := c.Request.Endpoints[:0]
			for _, item := range c.Request.Endpoints {
				if strings.TrimSpace(item) != "" {
					endpoints = append(endpoints, item)
				}
			}
			return endpoints
		}(),
		UserName: strings.TrimSpace(c.Request.UserName),
		Password: strings.TrimSpace(c.Request.Password),
	}); err == nil {
		ctx.String(http.StatusOK, "更新成功")
	} else {
		ctx.String(http.StatusBadRequest, "参数错误")
	}
}

type Keylist struct {
	httputil.Base
	Request struct {
		Prefix string `form:"prefix"`
	}
	Response struct{}
}

func (c Keylist) Run(ctx *gin.Context) {
	if err := c.DoReq(ctx, 0, &c.Request); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}

	if _, err := currConf.VerifyDriver(); err != nil {
		ctx.String(http.StatusOK, "etcd-conf未配置成功")
		return
	}

	res, err := currDriver.List(c.Request.Prefix)
	if err != nil {
		logger.DefaultLogger.Error(err.Error())
		ctx.String(http.StatusInternalServerError, err.Error())
	} else {
		ctx.JSON(http.StatusOK, res)
	}
}

type Keyput struct {
	httputil.Base
	Request struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	Response struct{}
}

func (c Keyput) Run(ctx *gin.Context) {
	if err := c.DoReq(ctx, 0, &c.Request); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}
	if _, err := currConf.VerifyDriver(); err != nil {
		ctx.String(http.StatusOK, "etcd-conf未配置成功")
		return
	}

	err := currDriver.Put(c.Request.Key, c.Request.Value)
	if err != nil {
		logger.DefaultLogger.Error(err.Error())
		ctx.String(http.StatusInternalServerError, err.Error())
	} else {
		ctx.String(http.StatusOK, "ok")
	}
}

type Keydelete struct {
	httputil.Base
	Request  struct{}
	Response struct{}
}

func (c Keydelete) Run(ctx *gin.Context) {

}
