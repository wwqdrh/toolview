package etcd

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var etcdAPI = []struct {
	method  string
	url     string
	handler []gin.HandlerFunc
}{
	{"GET", "/conf/status", []gin.HandlerFunc{confStatus}},
	{"POST", "/conf/update", []gin.HandlerFunc{confUpdate}},
	{"GET", "/key/list", []gin.HandlerFunc{keyList}},
	{"POST", "/key/put", []gin.HandlerFunc{keyPut}},
	{"POST", "/key/delete", []gin.HandlerFunc{keyDelete}},
}

func RegisterAPI(engine *gin.Engine) {
	etcd := engine.Group("etcd")
	for _, item := range etcdAPI {
		etcd.Handle(item.method, item.url, item.handler...)
	}
}

func confStatus(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, currConf)
}

type confUpdateReq struct {
	Endpoints []string `json:"endpoints"`
	UserName  string   `json:"username"`
	Password  string   `json:"password"`
}

func confUpdate(ctx *gin.Context) {
	var req confUpdateReq
	ctx.BindJSON(&req)

	if err := currConf.Update(&etcdConf{
		Endpoints: req.Endpoints,
		UserName:  strings.TrimSpace(req.UserName),
		Password:  strings.TrimSpace(req.Password),
	}); err == nil {
		ctx.String(http.StatusOK, "更新成功")
	} else {
		ctx.String(http.StatusBadRequest, "参数错误")
	}
}

func keyList(ctx *gin.Context) {

}

func keyPut(ctx *gin.Context) {

}

func keyDelete(ctx *gin.Context) {

}
