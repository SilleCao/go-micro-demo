package middlewares

import (
	"fmt"
	"time"

	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s  [%s] %s %s%d %s %s\n",
			params.TimeStamp.Format(time.RFC3339),
			params.Request.Header.Get(common.GetReqIdKey()),
			params.Method,
			params.Path,
			params.StatusCode,
			params.Latency,
			params.ErrorMessage,
		)
	})
}

func TraceRequest() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		reqId := ctx.GetHeader(common.GetReqIdKey())
		if len(reqId) == 0 {
			reqId = uuid.NewString()
		}
		ctx.Set(common.GetReqIdKey(), reqId)
		ctx.Next()
	})
}
