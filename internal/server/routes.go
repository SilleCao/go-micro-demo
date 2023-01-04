package server

import (
	"github.com/SilleCao/golang/go-micro-demo/internal/config"
	sys "github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/api"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, conf *config.Config) {

	v1 := router.Group(config.ApiUri)
	{
		sysRg := v1.Group("/sys")
		{
			sys.CreateUser(sysRg)
			sys.UpdateUser(sysRg)
			sys.GetUsers(sysRg)
			sys.GetUserById(sysRg)
			sys.DeleteUser(sysRg)
			sys.LockUser(sysRg)
			sys.UnlockUser(sysRg)
		}
	}
}
