package server

import (
	"github.com/SilleCao/golang/go-micro-demo/internal/config"
	auth "github.com/SilleCao/golang/go-micro-demo/internal/modules/auth/api"
	sys "github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/api"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/middlewares"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(router *gin.Engine, conf *config.Config) {
	rg := router.Group(conf.Server.ContextPath)
	{
		v1 := rg.Group(config.ApiUri, middlewares.ValidateJWT())
		// v1 := router.Group(config.ApiUri)
		{
			sysRg := v1.Group("")
			{
				sys.CreateUser(sysRg)
				sys.UpdateUser(sysRg)
				sys.GetUsers(sysRg)
				sys.GetUserById(sysRg)
				sys.UpdateUserStatus(sysRg)

				sys.CreateRole(sysRg)
				sys.UpdateRole(sysRg)
				sys.GetRoles(sysRg)
				sys.DeleteRole(sysRg)
			}
		}

		authRg := rg.Group(config.ApiUri)
		{
			auth.GenerateToken(authRg)
		}

	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
