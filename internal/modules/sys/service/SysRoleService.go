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

func CreateRole(ctx *gin.Context, role *dto.SysRoleRequest) error {
	sr := model.SysRole{}
	copier.Copy(&sr, role)
	su, _ := GetLoginUser(ctx)
	sr.CreateDate = time.Now()
	sr.UpdateDate = time.Now()
	sr.Creator = su.ID
	sr.Updater = su.ID
	return repo.CreateRole(ctx, &sr)
}

func UpdateRole(ctx *gin.Context, m map[string]interface{}) error {
	consolidateData(ctx, m, model.SysRole{})
	return repo.UpdateRole(ctx, m)
}

func GetRoles(ctx *gin.Context, role *dto.SysRoleRequest, page *common.Pagination) (*common.Pagination, error) {
	sr := model.SysRole{}
	copier.Copy(&sr, role)
	return repo.GetRoles(ctx, &sr, page)
}

func DeleteRole(ctx *gin.Context, id int64) error {
	return repo.DeleteRole(ctx, id)
}
