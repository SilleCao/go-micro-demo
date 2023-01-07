package repository

import (
	"context"

	"github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/model"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/common"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/dao"
)

func CreateUser(user *model.SysUser, ctx context.Context) (err error) {
	q := dao.Use(dao.Db())
	err = q.SysUser.WithContext(ctx).Create(user)
	return err
}

func GetUsers(page *common.Pagination, ctx context.Context) (*common.Pagination, error) {
	q := dao.Use(dao.Db())
	su := q.SysUser
	result, count, err := q.SysUser.WithContext(ctx).Order(su.CreateDate.Desc()).FindByPage(page.GetOffset(), page.GetSize())
	page.Content = result
	page.TotalContent = count
	return page, err
}

func GetUserByUsername(username string, ctx context.Context) (*model.SysUser, error) {
	q := dao.Use(dao.Db())
	su := q.SysUser
	return su.WithContext(ctx).Where(su.Username.Eq(username)).First()
}
