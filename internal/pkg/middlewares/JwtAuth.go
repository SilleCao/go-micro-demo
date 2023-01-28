package middlewares

import (
	"net/http"

	"github.com/SilleCao/golang/go-micro-demo/internal/modules/auth/service"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/common"
	"github.com/SilleCao/golang/go-micro-demo/pkg/logger"
	"github.com/gin-gonic/gin"
)

func ValidateJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr, err := service.GetToken(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common.NewErrResp(http.StatusUnauthorized, "invaild token", err))
			return
		}
		//TODO valid token
		token, err := service.ValidateToken(tokenStr)
		if err != nil {
			logger.Err("invalid token", ctx, err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common.NewErrResp(http.StatusUnauthorized, "invaild token", err))
			return
		}
		if !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common.NewErrResp(http.StatusUnauthorized, "invaild token", err))
			return
		}

		// if token.Valid {
		// 	mc := token.Claims.(jwt.MapClaims)
		// 	log.Println("Claims[Name]", mc["name"])
		// 	log.Println("Claims[roles]", mc["roles"])
		// 	log.Println("Claims[Issuer]", mc["iss"])
		// 	log.Println("Claims[IssueAt]", mc["iat"])
		// 	log.Println("Claims[ExpireAt]", mc["exp"])
		// } else {
		// 	log.Println(ctx.Value(common.GetReqIdKey()), "Token is invalid")
		// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		// 		"message": "invalid token",
		// 	})
		// }
	}
}
