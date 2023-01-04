package repository

import (
	"context"

	"github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/model"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/dao"
)

func CreateUser(user *model.SysUser, ctx context.Context) (err error) {
	q := dao.Use(dao.Db())
	err = q.SysUser.WithContext(ctx).Create(user)
	return err
}
