package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(router *gin.RouterGroup) {
	router.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, nil)
	})
}
