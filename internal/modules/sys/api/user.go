package api

import (
	"context"
	"net/http"
	"strconv"

	"github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/model"
	"github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/service"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/common"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func CreateUser(router *gin.RouterGroup) {
	router.POST("/users", func(ctx *gin.Context) {
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
	router.PUT("/users/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, nil)
	})
}

func GetUsers(router *gin.RouterGroup) {
	router.GET("/users", func(ctx *gin.Context) {
		page := ctx.Query("page")
		size := ctx.Query("size")
		if !govalidator.IsInt(page) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "invaild parameters",
			})
			return
		}
		if !govalidator.IsInt(size) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "invaild parameters",
			})
			return
		}

		pageNum, _ := strconv.Atoi(page)
		sizeNum, _ := strconv.Atoi(size)
		p, err := service.GetUsers(pageNum, sizeNum, context.Background())
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
	router.GET("/users/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, nil)
	})
}

func DeleteUser(router *gin.RouterGroup) {
	router.DELETE("/users/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, nil)
	})
}

func LockUser(router *gin.RouterGroup) {
	router.POST("/users/lock", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, nil)
	})
}

func UnlockUser(router *gin.RouterGroup) {
	router.POST("/users/unlock", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, nil)
	})
}
