package common

import (
	"github.com/gin-gonic/gin"
)

func GetReqIdKey() string {
	return string(REQ_ID)
}

func GetReqIdValue(ctx *gin.Context) string {
	reqId := ctx.Value(REQ_ID)
	if reqId == nil {
		reqId = ""
	}
	return reqId.(string)
}
