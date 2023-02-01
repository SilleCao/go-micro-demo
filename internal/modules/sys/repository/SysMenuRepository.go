package repository

import (
	"errors"

	"github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/model"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/common"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/dao"
	"github.com/SilleCao/golang/go-micro-demo/pkg/logger"
	"github.com/gin-gonic/gin"
)

func CreateMenu(ctx *gin.Context, menu *model.SysMenu) (err error) {
	return dao.DbQuery().SysMenu.WithContext(ctx).Create(menu)
}

func UpdateMenu(ctx *gin.Context, m map[string]interface{}) (err error) {
	sm := dao.DbQuery().SysMenu
	ri, err := sm.WithContext(ctx).Where(sm.ID.Eq(m["id"].(int64))).Updates(m)
	if ri.RowsAffected == 0 {
		err = errors.New("no rows were updated")
		logger.Err("Update failed", ctx, err)
		return err
	}
	return err
}

func GetMenus(ctx *gin.Context, menu *model.SysMenu, page *common.Pagination) (*common.Pagination, error) {
	sm := dao.DbQuery().SysMenu
	sdo := sm.WithContext(ctx)
	if len(ctx.Query("name")) > 0 {
		sdo = sdo.Where(sm.Name.Like("%" + menu.Name + "%"))
	}
	if len(ctx.Query("url")) > 0 {
		sdo = sdo.Where(sm.URL.Like("%" + menu.URL + "%"))
	}
	if len(ctx.Query("pid")) > 0 {
		sdo = sdo.Where(sm.Pid.Eq(menu.Pid))
	}
	result, count, err := sdo.Order(sm.CreateDate.Desc()).FindByPage(page.GetOffset(), page.GetSize())
	page.Content = result
	page.TotalContent = count
	return page, err
}

func DeleteMenu(ctx *gin.Context, id int64) (err error) {
	sm := dao.DbQuery().SysMenu
	ri, err := sm.WithContext(ctx).Where(sm.ID.Eq(id)).Delete()
	if ri.RowsAffected == 0 {
		err = errors.New("no rows were updated")
		logger.Err("Update failed", ctx, err)
		return err
	}
	return err
}
