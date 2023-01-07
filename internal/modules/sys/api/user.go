package api

import (
	"context"
	"net/http"

	"github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/model"
	"github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/service"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/common"
	"github.com/gin-gonic/gin"
)

// CreateUser
// @Description Create the new user
// @Tags 		sys
// @Produce 	json
// @Param 		user body model.SysUser true "SysUser JSON"
// @Success		200
// @Router		/sys/users [post]
// @Security BearerAuth
func CreateUser(router *gin.RouterGroup) {
	router.POST("/sys/users", func(ctx *gin.Context) {
		var user model.SysUser
		err := ctx.BindJSON(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		err = service.CreateUser(&user, context.Background())
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewErrResponse(http.StatusBadRequest, "", err))
			return
		}
		ctx.JSON(http.StatusOK, nil)
	})
}

func UpdateUser(router *gin.RouterGroup) {
	router.PUT("/sys/users/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, nil)
	})
}

func GetUsers(router *gin.RouterGroup) {
	router.GET("/sys/users", func(ctx *gin.Context) {
		var pagination common.Pagination
		err := ctx.BindQuery(&pagination)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "invaild parameters",
				"err":     err,
			})
			return
		}
		p, err := service.GetUsers(&pagination, context.Background())
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "invaild parameters",
				"err":     err,
			})
			return
		}
		ctx.JSON(http.StatusOK, p)
	})
}

func GetUserById(router *gin.RouterGroup) {
	router.GET("/sys/users/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, nil)
	})
}

func DeleteUser(router *gin.RouterGroup) {
	router.DELETE("/users/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, nil)
	})
}

func LockUser(router *gin.RouterGroup) {
	router.POST("/sys/users/lock", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, nil)
	})
}

func UnlockUser(router *gin.RouterGroup) {
	router.POST("/sys/users/unlock", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, nil)
	})
}
