package api

import (
	"net/http"

	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/common"
	"github.com/gin-gonic/gin"
)

// CreateRole
// @Summary 	Create Role
// @Description Create the new role
// @Tags 		sys/role
// @Produce 	json
// @Param 		role body model.SysRole true "SysRole JSON"
// @Success		200
// @Router		/sys/roles [post]
// @Security BearerAuth
func CreateRole(router *gin.RouterGroup) {
	router.POST("/sys/roles", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, common.NewResp())
	})
}

// UpdateRole
// @Summary 	Update Role
// @Description Update the role
// @Tags 		sys/role
// @Produce 	json
// @Param 		id	path	int false "role id"
// @Param 		role body model.SysRole true "SysRole JSON"
// @Success		200
// @Router		/sys/roles [put]
// @Security BearerAuth
func UpdateRole(router *gin.RouterGroup) {
	router.PUT("/sys/roles/:id", func(ctx *gin.Context) {
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
// @Success		200
// @Router		/sys/roles [get]
// @Security BearerAuth
func GetRoles(router *gin.RouterGroup) {
	router.GET("/sys/roles", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, common.NewResp())
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
		ctx.JSON(http.StatusOK, common.NewResp())
	})
}
