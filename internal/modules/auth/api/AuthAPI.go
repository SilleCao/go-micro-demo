package api

import (
	"net/http"

	"github.com/SilleCao/golang/go-micro-demo/internal/modules/auth/model"
	"github.com/SilleCao/golang/go-micro-demo/internal/modules/auth/service"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/common"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/errors"
	"github.com/SilleCao/golang/go-micro-demo/pkg/logger"
	"github.com/gin-gonic/gin"
)

// @Summary 	Generate token
// @Description Authenicate login user and return token
// @Tags 		auth
// @Produce 	json
// @Param 		credential body model.Credentials true "Credentials JSON"
// @Success		200
// @Router		/auth/token [post]
func GenerateToken(router *gin.RouterGroup) {
	router.POST("/auth/token", func(ctx *gin.Context) {
		var cdtl model.Credentials
		err := ctx.BindJSON(&cdtl)
		if err != nil {
			baseErr := errors.HandleErr(err)
			ctx.JSON(baseErr.StatusCode, common.NewErrResp(baseErr.ErrCode, "Authenticate failed", err))
			return
		}
		token, err := service.Authenticate(ctx, &cdtl)
		if err != nil {
			logger.Err("Authenticate failed", ctx, err)
			baseErr := errors.HandleErr(err)
			ctx.JSON(baseErr.StatusCode, common.NewErrResp(baseErr.ErrCode, "Authenticate failed", err))
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"issued_token": token,
		})
	})
}
