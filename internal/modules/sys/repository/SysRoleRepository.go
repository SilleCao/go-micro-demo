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
	result, count, err := sr.WithContext(ctx).Where(
		sr.Name.Like("%"+role.Name+"%"),
		sr.Remark.Like("%"+role.Remark+"%"),
		sr.DeptID.Eq(role.DeptID),
	).Order(sr.CreateDate.Desc()).FindByPage(page.GetOffset(), page.GetSize())
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
