package api

import (
	"context"
	"log"
	"net/http"

	"github.com/SilleCao/golang/go-micro-demo/internal/modules/auth/model"
	"github.com/SilleCao/golang/go-micro-demo/internal/modules/auth/service"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/common"
	"github.com/gin-gonic/gin"
)

func GenerateToken(router *gin.RouterGroup) {
	router.POST("/token", func(ctx *gin.Context) {
		var cdtl model.Credentials
		err := ctx.BindJSON(&cdtl)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		token, err := service.Authenticate(&cdtl, context.Background())
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, common.NewErrResponse(http.StatusBadRequest, "", err))
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"issued_token": token,
		})
	})
}
