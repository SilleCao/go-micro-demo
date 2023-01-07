package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/SilleCao/golang/go-micro-demo/internal/modules/auth/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const BEARER_SCHEMA = "Bearer "
const HEADER_ATTR_AUTHORIZATION = "Authorization"

func ValidateJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(HEADER_ATTR_AUTHORIZATION)
		if !strings.HasPrefix(authHeader, BEARER_SCHEMA) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})
			return
		}
		tokenStr := authHeader[len(BEARER_SCHEMA):]
		//TODO valid token
		token, err := service.ValidateToken(tokenStr)
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})
			return
		}

		if token.Valid {
			mc := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Name]", mc["name"])
			log.Println("Claims[roles]", mc["roles"])
			log.Println("Claims[Issuer]", mc["iss"])
			log.Println("Claims[IssueAt]", mc["iat"])
			log.Println("Claims[ExpireAt]", mc["exp"])
		} else {
			log.Println("token is invalid")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})
		}
	}
}
