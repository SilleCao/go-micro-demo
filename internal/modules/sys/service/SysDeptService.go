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

func CreateDept(ctx *gin.Context, dept *dto.SysDeptRequest) error {
	sm := model.SysDept{}
	copier.Copy(&sm, dept)
	su, _ := GetLoginUser(ctx)
	sm.CreateDate = time.Now()
	sm.UpdateDate = time.Now()
	sm.Creator = su.ID
	sm.Updater = su.ID
	return repo.CreateDept(ctx, &sm)
}

func UpdateDept(ctx *gin.Context, m map[string]interface{}) error {
	consolidateData(ctx, m, model.SysDept{})
	return repo.UpdateDept(ctx, m)
}

func GetDepts(ctx *gin.Context, dept *dto.SysDeptRequest, page *common.Pagination) (*common.Pagination, error) {
	sm := model.SysDept{}
	copier.Copy(&sm, dept)
	return repo.GetDepts(ctx, &sm, page)
}

func DeleteDept(ctx *gin.Context, id int64) error {
	return repo.DeleteDept(ctx, id)
}
