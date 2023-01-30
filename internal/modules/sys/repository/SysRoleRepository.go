package repository

import (
	"errors"

	"github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/model"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/common"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/dao"
	"github.com/SilleCao/golang/go-micro-demo/pkg/logger"
	"github.com/gin-gonic/gin"
)

func CreateRole(ctx *gin.Context, role *model.SysRole) (err error) {
	return dao.DbQuery().SysRole.WithContext(ctx).Create(role)
}

func UpdateRole(ctx *gin.Context, role *model.SysRole) (err error) {
	sr := dao.DbQuery().SysRole
	ri, err := sr.WithContext(ctx).Where(sr.ID.Eq(role.ID)).Updates(role)
	if ri.RowsAffected == 0 {
		err = errors.New("no rows were updated")
		logger.Err("Update failed", ctx, err)
		return err
	}
	return err
}

func GetRoles(ctx *gin.Context, role *model.SysRole, page *common.Pagination) (*common.Pagination, error) {

	sr := dao.DbQuery().SysRole
	sdo := sr.WithContext(ctx)
	if len(ctx.Query("name")) > 0 {
		sdo = sdo.Where(sr.Name.Like("%" + role.Name + "%"))
	}
	if len(ctx.Query("remark")) > 0 {
		sdo = sdo.Where(sr.Remark.Like("%" + role.Remark + "%"))
	}
	if len(ctx.Query("deptId")) > 0 {
		sdo = sdo.Where(sr.DeptID.Eq(role.DeptID))
	}
	result, count, err := sdo.Order(sr.CreateDate.Desc()).FindByPage(page.GetOffset(), page.GetSize())
	page.Content = result
	page.TotalContent = count
	return page, err
}

func DeleteRole(ctx *gin.Context, id int64) (err error) {
	sr := dao.DbQuery().SysRole
	ri, err := sr.WithContext(ctx).Where(sr.ID.Eq(id)).Delete()
	if ri.RowsAffected == 0 {
		err = errors.New("no rows were updated")
		logger.Err("Update failed", ctx, err)
		return err
	}
	return err
}
