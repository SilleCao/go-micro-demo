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

// CreateMenu
// @Summary 	Create Menu
// @Description Create the new menu
// @Tags 		sys/menu
// @Produce 	json
// @Param 		menu body dto.SysMenuRequest true "SysMenuRequest JSON"
// @Success		200
// @Router		/sys/menus [post]
// @Security BearerAuth
func CreateMenu(router *gin.RouterGroup) {
	router.POST("/sys/menus", func(ctx *gin.Context) {
		var menu dto.SysMenuRequest
		err := ctx.BindJSON(&menu)
		if err != nil {
			logger.Err("invalid request body", ctx, err)
			common.SetErrResp(ctx, http.StatusBadRequest, "invalid request body", err)
			return
		}
		err = service.CreateMenu(ctx, &menu)
		if err != nil {
			logger.Err("Create menu fail", ctx, err)
			common.SetErrResp(ctx, http.StatusBadRequest, "Create menu fail", err)
			return
		}
		common.SetScesResp(ctx, &menu)
	})
}

// UpdateMenu
// @Summary 	Update Menu
// @Description Update the menu
// @Tags 		sys/menu
// @Produce 	json
// @Param 		id	path	int false "menu id"
// @Param 		menu body dto.SysMenuRequest true "SysMenuRequest JSON"
// @Success		200
// @Router		/sys/menus [put]
// @Security BearerAuth
func UpdateMenu(router *gin.RouterGroup) {
	router.PUT("/sys/menus/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		aid, err := strconv.Atoi(id)
		if err != nil {
			logger.Err("invalid request param[id]", ctx, err)
			common.SetErrResp(ctx, http.StatusBadRequest, err.Error(), err)
			return
		}
		var menu dto.SysMenuRequest
		err = ctx.BindJSON(&menu)
		if err != nil {
			logger.Err("invalid request body", ctx, err)
			common.SetErrResp(ctx, http.StatusBadRequest, err.Error(), err)
			return
		}
		menu.ID = int64(aid)
		err = service.UpdateMenu(ctx, &menu)
		if err != nil {
			logger.Err(err.Error(), ctx, err)
			common.SetErrResp(ctx, http.StatusBadRequest, err.Error(), err)
			return
		}
		ctx.JSON(http.StatusOK, common.NewResp())
	})
}

// GetMenus
// @Summary 	Get Menus
// @Description Get the list of Menu
// @Tags 		sys/menu
// @Produce 	json
// @Param 		name	query	string false "menu name"
// @Param 		url		query	string false "menu url"
// @Param 		pid		query	string false "parent menu id"
// @Param 		page	query	int false "page number"
// @Param 		size	query	int false "page size"
// @Success		200
// @Router		/sys/menus [get]
// @Security BearerAuth
func GetMenus(router *gin.RouterGroup) {
	router.GET("/sys/menus", func(ctx *gin.Context) {
		var pagination common.Pagination
		err := ctx.BindQuery(&pagination)
		if err != nil {
			common.SetErrResp(ctx, http.StatusBadRequest, "invaild parameters", err)
			return
		}
		var menu dto.SysMenuRequest
		err = ctx.BindQuery(&menu)
		if err != nil {
			common.SetErrResp(ctx, http.StatusBadRequest, "invaild parameters", err)
			return
		}
		p, err := service.GetMenus(ctx, &menu, &pagination)
		if err != nil {
			common.SetErrResp(ctx, http.StatusBadRequest, "invaild parameters", err)
			return
		}
		common.SetScesResp(ctx, p)
	})
}

// DeleteMenu
// @Summary 	Delete Menu
// @Description Delete Menu
// @Tags 		sys/menu
// @Produce 	json
// @Param 		id	path	int false "menu id"
// @Success		200
// @Router		/sys/menus [delete]
// @Security BearerAuth
func DeleteMenu(router *gin.RouterGroup) {
	router.DELETE("/sys/menus/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		aid, err := strconv.Atoi(id)
		if err != nil {
			logger.Err("invalid request param[id]", ctx, err)
			common.SetErrResp(ctx, http.StatusBadRequest, err.Error(), err)
			return
		}
		err = service.DeleteMenu(ctx, int64(aid))
		if err != nil {
			logger.Err("Delete menu fail", ctx, err)
			common.SetErrResp(ctx, http.StatusBadRequest, err.Error(), err)
			return
		}
		ctx.JSON(http.StatusOK, common.NewResp())
	})
}
