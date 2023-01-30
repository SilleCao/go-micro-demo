package service

import (
	"time"

	"github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/dto"
	"github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/model"
	repo "github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/repository"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/common"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func CreateMenu(ctx *gin.Context, menu *dto.SysMenuRequest) error {
	sm := model.SysMenu{}
	copier.Copy(&sm, menu)
	su, _ := GetLoginUser(ctx)
	sm.CreateDate = time.Now()
	sm.UpdateDate = time.Now()
	sm.Creator = su.ID
	sm.Updater = su.ID
	return repo.CreateMenu(ctx, &sm)
}

func UpdateMenu(ctx *gin.Context, menu *dto.SysMenuRequest) error {
	sm := model.SysMenu{}
	copier.Copy(&sm, menu)
	su, _ := GetLoginUser(ctx)
	sm.UpdateDate = time.Now()
	sm.Updater = su.ID
	return repo.UpdateMenu(ctx, &sm)
}

func GetMenus(ctx *gin.Context, menu *dto.SysMenuRequest, page *common.Pagination) (*common.Pagination, error) {
	sm := model.SysMenu{}
	copier.Copy(&sm, menu)
	return repo.GetMenus(ctx, &sm, page)
}

func DeleteMenu(ctx *gin.Context, id int64) error {
	return repo.DeleteMenu(ctx, id)
}
