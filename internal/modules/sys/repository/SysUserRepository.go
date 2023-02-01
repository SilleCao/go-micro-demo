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

func UpdateUser(ctx *gin.Context, m map[string]interface{}) (err error) {
	su := dao.DbQuery().SysUser
	rstInfo, err := su.WithContext(ctx).Where(su.ID.Eq(m["id"].(int64))).Updates(&m)
	if rstInfo.RowsAffected == 0 {
		err = errors.New("no rows were updated")
		logger.Err("Update failed", ctx, err)
		return err
	}
	return err
}

func GetUsers(ctx *gin.Context, user *model.SysUser, page *common.Pagination) (*common.Pagination, error) {
	su := dao.DbQuery().SysUser
	sdo := su.WithContext(ctx)
	if len(ctx.Query("username")) > 0 {
		sdo = sdo.Where(su.Username.Like("%" + user.Username + "%"))
	}
	if len(ctx.Query("realName")) > 0 {
		sdo = sdo.Where(su.RealName.Like("%" + user.RealName + "%"))
	}
	if len(ctx.Query("email")) > 0 {
		sdo = sdo.Where(su.Email.Like("%" + user.Email + "%"))
	}
	if len(ctx.Query("mobile")) > 0 {
		sdo = sdo.Where(su.Username.Like("%" + user.Mobile + "%"))
	}
	if len(ctx.Query("gender")) > 0 {
		sdo = sdo.Where(su.Gender.Eq(user.Gender))
	}
	if len(ctx.Query("status")) > 0 {
		sdo = sdo.Where(su.Status.Eq(user.Status))
	}
	if len(ctx.Query("superAdmin")) > 0 {
		sdo = sdo.Where(su.SuperAdmin.Eq(user.SuperAdmin))
	}
	result, count, err := sdo.Where(
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
