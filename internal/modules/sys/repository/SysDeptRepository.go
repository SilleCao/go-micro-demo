package repository

import (
	"errors"

	"github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/model"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/common"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/dao"
	"github.com/SilleCao/golang/go-micro-demo/pkg/logger"
	"github.com/gin-gonic/gin"
)

func CreateDept(ctx *gin.Context, dept *model.SysDept) (err error) {
	return dao.DbQuery().SysDept.WithContext(ctx).Create(dept)
}

func UpdateDept(ctx *gin.Context, m map[string]interface{}) (err error) {
	sm := dao.DbQuery().SysDept
	ri, err := sm.WithContext(ctx).Where(sm.ID.Eq(m["id"].(int64))).Updates(m)
	if ri.RowsAffected == 0 {
		err = errors.New("no rows were updated")
		logger.Err("Update failed", ctx, err)
		return err
	}
	return err
}

func GetDepts(ctx *gin.Context, dept *model.SysDept, page *common.Pagination) (*common.Pagination, error) {
	sm := dao.DbQuery().SysDept
	sdo := sm.WithContext(ctx)
	if len(ctx.Query("name")) > 0 {
		sdo = sdo.Where(sm.Name.Like("%" + dept.Name + "%"))
	}
	if len(ctx.Query("pid")) > 0 {
		sdo = sdo.Where(sm.Pid.Eq(dept.Pid))
	}
	result, count, err := sdo.Order(sm.CreateDate.Desc()).FindByPage(page.GetOffset(), page.GetSize())
	page.Content = result
	page.TotalContent = count
	return page, err
}

func DeleteDept(ctx *gin.Context, id int64) (err error) {
	sm := dao.DbQuery().SysDept
	ri, err := sm.WithContext(ctx).Where(sm.ID.Eq(id)).Delete()
	if ri.RowsAffected == 0 {
		err = errors.New("no rows were updated")
		logger.Err("Update failed", ctx, err)
		return err
	}
	return err
}
