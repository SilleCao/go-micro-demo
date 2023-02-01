package service

import (
	"time"

	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/common"
	"github.com/gin-gonic/gin"
)

func consolidateData(ctx *gin.Context, m map[string]interface{}, obj any) {
	su, _ := GetLoginUser(ctx)
	m["updater"] = su.ID
	m["updateDate"] = time.Now()
	common.ConvertToGormKey(m, obj)
}

// func addCreatorAndUpdaterInfo(ctx *gin.Context, m map[string]interface{}) {
// 	su, _ := GetLoginUser(ctx)
// 	m["creator"] = su.ID
// 	m["createDate"] = time.Now()
// 	m["updater"] = su.ID
// 	m["updateDate"] = time.Now()
// }
