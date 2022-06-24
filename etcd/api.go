package etcd

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wwqdrh/logger"
)

func ConfStatus(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, currConf)
}

func ConfVerify(ctx *gin.Context) {
	if _, err := currConf.VerifyDriver(); err != nil {
		logger.DefaultLogger.Error(err.Error())
		ctx.String(http.StatusOK, "fail")
	} else {
		ctx.String(http.StatusOK, "ok")
	}
}

type confUpdateReq struct {
	Endpoints []string `json:"endpoints"`
	UserName  string   `json:"username"`
	Password  string   `json:"password"`
}

func ConfUpdate(ctx *gin.Context) {
	var req confUpdateReq
	if err := ctx.BindJSON(&req); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := currConf.Update(&etcdConf{
		Endpoints: func() []string {
			endpoints := req.Endpoints[:0]
			for _, item := range req.Endpoints {
				if strings.TrimSpace(item) != "" {
					endpoints = append(endpoints, item)
				}
			}
			return endpoints
		}(),
		UserName: strings.TrimSpace(req.UserName),
		Password: strings.TrimSpace(req.Password),
	}); err == nil {
		ctx.String(http.StatusOK, "更新成功")
	} else {
		ctx.String(http.StatusBadRequest, "参数错误")
	}
}

type keyListReq struct {
	Prefix string `form:"prefix"`
}

func KeyList(ctx *gin.Context) {
	var req keyListReq
	if err := ctx.BindQuery(&req); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	if _, err := currConf.VerifyDriver(); err != nil {
		ctx.String(http.StatusOK, "etcd-conf未配置成功")
		return
	}

	res, err := currDriver.List(req.Prefix)
	if err != nil {
		logger.DefaultLogger.Error(err.Error())
		ctx.String(http.StatusInternalServerError, err.Error())
	} else {
		ctx.JSON(http.StatusOK, res)
	}
}

type keyPutReq struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func KeyPut(ctx *gin.Context) {
	var req keyPutReq
	if err := ctx.BindJSON(&req); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	if _, err := currConf.VerifyDriver(); err != nil {
		ctx.String(http.StatusOK, "etcd-conf未配置成功")
		return
	}

	err := currDriver.Put(req.Key, req.Value)
	if err != nil {
		logger.DefaultLogger.Error(err.Error())
		ctx.String(http.StatusInternalServerError, err.Error())
	} else {
		ctx.String(http.StatusOK, "ok")
	}

}

func KeyDelete(ctx *gin.Context) {

}
