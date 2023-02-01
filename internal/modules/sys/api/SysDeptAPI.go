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

// CreateDept
// @Summary 	Create Dept
// @Description Create the new dept
// @Tags 		sys/dept
// @Produce 	json
// @Param 		dept body dto.SysDeptRequest true "SysDeptRequest JSON"
// @Success		200
// @Router		/sys/depts [post]
// @Security BearerAuth
func CreateDept(router *gin.RouterGroup) {
	router.POST("/sys/depts", func(ctx *gin.Context) {
		var dept dto.SysDeptRequest
		err := ctx.BindJSON(&dept)
		if err != nil {
			logger.Err("invalid request body", ctx, err)
			common.SetErrResp(ctx, http.StatusBadRequest, "invalid request body", err)
			return
		}
		err = service.CreateDept(ctx, &dept)
		if err != nil {
			logger.Err("Create dept fail", ctx, err)
			common.SetErrResp(ctx, http.StatusBadRequest, "Create dept fail", err)
			return
		}
		common.SetScesResp(ctx, &dept)
	})
}

// UpdateDept
// @Summary 	Update Dept
// @Description Update the dept
// @Tags 		sys/dept
// @Produce 	json
// @Param 		id	path	int false "dept id"
// @Param 		dept body dto.SysDeptRequest true "SysDeptRequest JSON"
// @Success		200
// @Router		/sys/depts [put]
// @Security BearerAuth
func UpdateDept(router *gin.RouterGroup) {
	router.PUT("/sys/depts/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		aid, err := strconv.Atoi(id)
		if err != nil {
			logger.Err("invalid request param[id]", ctx, err)
			common.SetErrResp(ctx, http.StatusBadRequest, err.Error(), err)
			return
		}
		m, err := common.GetReqBodyAsMap(ctx)
		if err != nil {
			logger.Err("invalid request body", ctx, err)
			common.SetErrResp(ctx, http.StatusBadRequest, err.Error(), err)
			return
		}
		m["id"] = int64(aid)
		err = service.UpdateDept(ctx, m)
		if err != nil {
			logger.Err(err.Error(), ctx, err)
			common.SetErrResp(ctx, http.StatusBadRequest, err.Error(), err)
			return
		}
		ctx.JSON(http.StatusOK, common.NewResp())
	})
}

// GetDepts
// @Summary 	Get Depts
// @Description Get the list of Dept
// @Tags 		sys/dept
// @Produce 	json
// @Param 		name	query	string false "dept name"
// @Param 		pid		query	string false "parent dept id"
// @Param 		page	query	int false "page number"
// @Param 		size	query	int false "page size"
// @Success		200
// @Router		/sys/depts [get]
// @Security BearerAuth
func GetDepts(router *gin.RouterGroup) {
	router.GET("/sys/depts", func(ctx *gin.Context) {
		var pagination common.Pagination
		err := ctx.BindQuery(&pagination)
		if err != nil {
			common.SetErrResp(ctx, http.StatusBadRequest, "invaild parameters", err)
			return
		}
		var dept dto.SysDeptRequest
		err = ctx.BindQuery(&dept)
		if err != nil {
			common.SetErrResp(ctx, http.StatusBadRequest, "invaild parameters", err)
			return
		}
		p, err := service.GetDepts(ctx, &dept, &pagination)
		if err != nil {
			common.SetErrResp(ctx, http.StatusBadRequest, "invaild parameters", err)
			return
		}
		common.SetScesResp(ctx, p)
	})
}

// DeleteDept
// @Summary 	Delete Dept
// @Description Delete Dept
// @Tags 		sys/dept
// @Produce 	json
// @Param 		id	path	int false "dept id"
// @Success		200
// @Router		/sys/depts [delete]
// @Security BearerAuth
func DeleteDept(router *gin.RouterGroup) {
	router.DELETE("/sys/depts/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		aid, err := strconv.Atoi(id)
		if err != nil {
			logger.Err("invalid request param[id]", ctx, err)
			common.SetErrResp(ctx, http.StatusBadRequest, err.Error(), err)
			return
		}
		err = service.DeleteDept(ctx, int64(aid))
		if err != nil {
			logger.Err("Delete dept fail", ctx, err)
			common.SetErrResp(ctx, http.StatusBadRequest, err.Error(), err)
			return
		}
		ctx.JSON(http.StatusOK, common.NewResp())
	})
}
