package api

import (
	"net/http"
	"strconv"

	"github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/dto"
	"github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/service"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/common"
	"github.com/SilleCao/golang/go-micro-demo/pkg/logger"
	"github.com/gin-gonic/gin"
)

// CreateRole
// @Summary 	Create Role
// @Description Create the new role
// @Tags 		sys/role
// @Produce 	json
// @Param 		role body dto.RoleRequest true "RoleRequest JSON"
// @Success		200
// @Router		/sys/roles [post]
// @Security BearerAuth
func CreateRole(router *gin.RouterGroup) {
	router.POST("/sys/roles", func(ctx *gin.Context) {
		var role dto.RoleRequest
		err := ctx.BindJSON(&role)
		if err != nil {
			logger.Err("invalid request body", ctx, err)
			common.SetErrResp(ctx, http.StatusBadRequest, "invalid request body", err)
			return
		}
		err = service.CreateRole(ctx, &role)
		if err != nil {
			logger.Err("Create role fail", ctx, err)
			common.SetErrResp(ctx, http.StatusBadRequest, "Create role fail", err)
			return
		}
		common.SetScesResp(ctx, &role)
	})
}

// UpdateRole
// @Summary 	Update Role
// @Description Update the role
// @Tags 		sys/role
// @Produce 	json
// @Param 		id	path	int false "role id"
// @Param 		role body dto.RoleRequest true "RoleRequest JSON"
// @Success		200
// @Router		/sys/roles [put]
// @Security BearerAuth
func UpdateRole(router *gin.RouterGroup) {
	router.PUT("/sys/roles/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		aid, err := strconv.Atoi(id)
		if err != nil {
			logger.Err("invalid request param[id]", ctx, err)
			common.SetErrResp(ctx, http.StatusBadRequest, err.Error(), err)
			return
		}
		var role dto.RoleRequest
		err = ctx.BindJSON(&role)
		if err != nil {
			logger.Err("invalid request body", ctx, err)
			common.SetErrResp(ctx, http.StatusBadRequest, err.Error(), err)
			return
		}
		role.ID = int64(aid)
		err = service.UpdateRole(ctx, &role)
		if err != nil {
			logger.Err(err.Error(), ctx, err)
			common.SetErrResp(ctx, http.StatusBadRequest, err.Error(), err)
			return
		}
		ctx.JSON(http.StatusOK, common.NewResp())
	})
}

// GetRoles
// @Summary 	Get Roles
// @Description Get the list of Role
// @Tags 		sys/role
// @Produce 	json
// @Param 		name	query	string false "role name"
// @Param 		remark	query	string false "role remark"
// @Param 		page	query	int false "page number"
// @Param 		size	query	int false "page size"
// @Success		200
// @Router		/sys/roles [get]
// @Security BearerAuth
func GetRoles(router *gin.RouterGroup) {
	router.GET("/sys/roles", func(ctx *gin.Context) {
		var pagination common.Pagination
		err := ctx.BindQuery(&pagination)
		if err != nil {
			common.SetErrResp(ctx, http.StatusBadRequest, "invaild parameters", err)
			return
		}
		var role dto.RoleRequest
		err = ctx.BindQuery(&role)
		if err != nil {
			common.SetErrResp(ctx, http.StatusBadRequest, "invaild parameters", err)
			return
		}
		p, err := service.GetRoles(ctx, &role, &pagination)
		if err != nil {
			common.SetErrResp(ctx, http.StatusBadRequest, "invaild parameters", err)
			return
		}
		common.SetScesResp(ctx, p)
	})
}

// DeleteRole
// @Summary 	Delete Role
// @Description Delete Role
// @Tags 		sys/role
// @Produce 	json
// @Param 		id	path	int false "role id"
// @Success		200
// @Router		/sys/roles [delete]
// @Security BearerAuth
func DeleteRole(router *gin.RouterGroup) {
	router.DELETE("/sys/roles/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		aid, err := strconv.Atoi(id)
		if err != nil {
			logger.Err("invalid request param[id]", ctx, err)
			common.SetErrResp(ctx, http.StatusBadRequest, err.Error(), err)
			return
		}
		err = service.DeleteRole(ctx, int64(aid))
		if err != nil {
			logger.Err("Delete role fail", ctx, err)
			common.SetErrResp(ctx, http.StatusBadRequest, err.Error(), err)
			return
		}
		ctx.JSON(http.StatusOK, common.NewResp())
	})
}
