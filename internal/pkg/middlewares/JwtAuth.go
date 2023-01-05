package middlewares

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const BEARER_SCHEMA = "Bearer "
const HEADER_ATTR_AUTHORIZATION = "Authorization"
const SECRET_KEY = "0000"

func ValidateJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(HEADER_ATTR_AUTHORIZATION)
		tokenStr := authHeader[len(BEARER_SCHEMA):]
		//TODO valid token
		token, err := validateToken(tokenStr)
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}

		if token.Valid {
			mc := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Name]", mc["name"])
			log.Println("Claims[Admin]", mc["admin"])
			log.Println("Claims[Issuer]", mc["iss"])
			log.Println("Claims[IssueAt]", mc["iat"])
			log.Println("Claims[ExpireAt]", mc["exp"])
		} else {
			log.Println("token is invalid")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		}
	}
}

func validateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(SECRET_KEY), nil
	})
}
