package repository

import (
	"errors"

	"github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/model"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/cache"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/common"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/dao"
	"github.com/SilleCao/golang/go-micro-demo/pkg/logger"
	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context, user *model.SysUser) (err error) {
	return dao.DbQuery().SysUser.WithContext(ctx).Create(user)
}

func UpdateUser(ctx *gin.Context, user *model.SysUser) (err error) {
	su := dao.DbQuery().SysUser
	rstInfo, err := su.WithContext(ctx).Where(su.ID.Eq(user.ID)).Updates(&user)
	if rstInfo.RowsAffected == 0 {
		err = errors.New("no rows were updated")
		logger.Err("Update failed", ctx, err)
		return err
	}
	return err
}

func GetUsers(ctx *gin.Context, user *model.SysUser, page *common.Pagination) (*common.Pagination, error) {
	su := dao.DbQuery().SysUser
	result, count, err := su.WithContext(ctx).Where(
		su.Username.Like("%"+user.Username+"%"),
		su.RealName.Like("%"+user.Username+"%"),
		su.Email.Like("%"+user.Username+"%"),
		su.Mobile.Like("%"+user.Mobile+"%"),
		su.Gender.Eq(user.Gender),
		su.Status.Eq(user.Status),
		su.SuperAdmin.Eq(user.SuperAdmin),
	).Order(su.CreateDate.Desc()).FindByPage(page.GetOffset(), page.GetSize())
	page.Content = result
	page.TotalContent = count
	return page, err
}

func GetUserById(ctx *gin.Context, id int64) (*model.SysUser, error) {
	su := dao.DbQuery().SysUser
	return su.WithContext(ctx).Where(su.ID.Eq(id)).First()
}

func GetUserByUsername(ctx *gin.Context, username string) (*model.SysUser, error) {
	var user model.SysUser
	err := cache.Get(ctx, "user:"+username, &user)
	if err == nil {
		return &user, nil
	} else {
		logger.Err("get user from cache fail", ctx, err)
	}
	su := dao.DbQuery().SysUser
	sysUser, err := su.WithContext(ctx).Where(su.Username.Eq(username)).First()
	if err == nil {
		cache.Set(ctx, "user:"+username, sysUser, common.UserDataCacheExpiration)
	}
	return sysUser, err
}
