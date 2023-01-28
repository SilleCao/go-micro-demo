package repository

import (
	"errors"

	"github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/model"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/common"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/dao"
	"github.com/SilleCao/golang/go-micro-demo/pkg/logger"
	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context, user *model.SysUser) (err error) {
	q := dao.Use(dao.Db())
	err = q.SysUser.WithContext(ctx).Create(user)
	return err
}

func UpdateUser(ctx *gin.Context, user *model.SysUser) (err error) {
	q := dao.Use(dao.Db())
	su := q.SysUser
	rstInfo, err := su.WithContext(ctx).Where(su.ID.Eq(user.ID)).Updates(&user)
	if rstInfo.RowsAffected == 0 {
		err2 := errors.New("no rows were updated")
		logger.Err("Update failed", ctx, err2)
		return err2
	}
	return err
}

func GetUsers(ctx *gin.Context, page *common.Pagination) (*common.Pagination, error) {
	q := dao.Use(dao.Db())
	su := q.SysUser
	result, count, err := q.SysUser.WithContext(ctx).Order(su.CreateDate.Desc()).FindByPage(page.GetOffset(), page.GetSize())
	page.Content = result
	page.TotalContent = count
	return page, err
}

func GetUserById(ctx *gin.Context, id int64) (*model.SysUser, error) {
	q := dao.Use(dao.Db())
	su := q.SysUser
	return su.WithContext(ctx).Where(su.ID.Eq(id)).First()
}

func GetUserByUsername(ctx *gin.Context, username string) (*model.SysUser, error) {
	q := dao.Use(dao.Db())
	su := q.SysUser
	return su.WithContext(ctx).Where(su.Username.Eq(username)).First()
}
