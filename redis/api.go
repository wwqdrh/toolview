package redis

import (
	"github.com/gin-gonic/gin"
	"github.com/wwqdrh/httputil"
)

type ConfStatus struct {
	httputil.Base
	Request  struct{}
	Response struct{}
}

func (c *ConfStatus) Run(ctx *gin.Context) {

}
