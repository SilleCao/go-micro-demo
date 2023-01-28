package api

import (
	"net/http"
	"strconv"

	"github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/dto"
	"github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/model"
	"github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/service"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/common"
	"github.com/SilleCao/golang/go-micro-demo/pkg/logger"
	"github.com/gin-gonic/gin"
)

// CreateUser
// @Summary 	Create User
// @Description Create the new user
// @Tags 		sys/user
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
		err = service.CreateUser(ctx, &user)
		if err != nil {
			logger.Err("Create user fail", ctx, err)
			ctx.JSON(http.StatusBadRequest, common.NewErrResp(http.StatusBadRequest, "Create user fail", err))
			return
		}
		ctx.JSON(http.StatusOK, common.NewScesResp(&user))
	})
}

// @Summary 	UpdateUser
// @Description Update user
// @Tags 		sys/user
// @Produce 	json
// @Param 		id	path	int false "user id"
// @Param 		user body model.SysUser true "SysUser JSON"
// @Success		200
// @Router		/sys/users [put]
// @Security 	BearerAuth
func UpdateUser(router *gin.RouterGroup) {
	router.PUT("/sys/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		aid, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewErrResp(http.StatusBadRequest, err.Error(), err))
			return
		}
		var user model.SysUser
		err = ctx.BindJSON(&user)
		if err != nil {
			logger.Err(err.Error(), ctx, err)
			ctx.JSON(http.StatusBadRequest, common.NewErrResp(http.StatusBadRequest, err.Error(), err))
			return
		}
		user.ID = int64(aid)
		err = service.UpdateUser(ctx, &user)
		if err != nil {
			logger.Err(err.Error(), ctx, err)
			ctx.JSON(http.StatusBadRequest, common.NewErrResp(http.StatusBadRequest, err.Error(), err))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResp())
	})
}

// @Summary 	GetUsers
// @Description Get users
// @Tags 		sys/user
// @Produce 	json
// @Param 		page	query	int false "page number"
// @Param 		size	query	int false "page size"
// @Success		200
// @Router		/sys/users [get]
// @Security 	BearerAuth
func GetUsers(router *gin.RouterGroup) {
	router.GET("/sys/users", func(ctx *gin.Context) {
		var pagination common.Pagination
		err := ctx.BindQuery(&pagination)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewErrResp(http.StatusBadRequest, "invaild parameters", err))
			return
		}
		p, err := service.GetUsers(ctx, &pagination)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewErrResp(http.StatusBadRequest, "invaild parameters", err))
			return
		}
		ctx.JSON(http.StatusOK, common.NewScesResp(p))
	})
}

// @Summary 	GetUserById
// @Description Get user by id
// @Tags 		sys/user
// @Produce 	json
// @Param 		id	path	int false "user id"
// @Success		200
// @Router		/sys/users/{id} [get]
// @Security 	BearerAuth
func GetUserById(router *gin.RouterGroup) {
	router.GET("/sys/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		aid, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewErrResp(http.StatusBadRequest, err.Error(), err))
			return
		}
		su, err := service.GetUserById(ctx, int64(aid))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewErrResp(http.StatusBadRequest, err.Error(), err))
			return
		}
		ctx.JSON(http.StatusOK, common.NewScesResp(su))
	})
}

// @Summary 	UpdateUserStatus
// @Description Update user's status
// @Tags 		sys/user
// @Produce 	json
// @Param 		user body dto.UpdateUserStatusDTO true "UpdateUserStatusDTO JSON"
// @Success		200
// @Router		/sys/users/status [put]
// @Security 	BearerAuth
func UpdateUserStatus(router *gin.RouterGroup) {
	router.PUT("/sys/users/status", func(ctx *gin.Context) {
		var user dto.UpdateUserStatusDTO
		err := ctx.BindJSON(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewErrResp(http.StatusBadRequest, "invaild parameters", err))
			return
		}
		err = service.UpdateUserStatus(ctx, user)
		if err != nil {
			logger.Err("Update user's status fail", ctx, err)
			ctx.JSON(http.StatusBadRequest, common.NewErrResp(http.StatusBadRequest, err.Error(), err))
			return
		}
		ctx.JSON(http.StatusOK, common.NewResp())
	})
}
