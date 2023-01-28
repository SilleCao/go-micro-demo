package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/SilleCao/golang/go-micro-demo/internal/modules/auth/model"
	"github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/repository"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/errors"
	"github.com/SilleCao/golang/go-micro-demo/pkg/logger"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type JWTCustomClaims struct {
	Name  string   `json:"name"`
	Roles []string `json:"roles"`
	jwt.StandardClaims
}

const DEFAULT_ISSUER = "sille.cn"
const SECRET_KEY = "0000"
const BEARER_SCHEMA = "Bearer "
const HEADER_ATTR_AUTHORIZATION = "Authorization"

func Authenticate(ctx *gin.Context, cdtl *model.Credentials) (string, error) {
	su, err := repository.GetUserByUsername(ctx, cdtl.Username)
	if err != nil {
		logger.Err("get user fail", ctx, err)
		return "", errors.NewNotFoundErr(err.Error(), 404)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(su.Password), []byte(cdtl.Password)); err != nil {
		logger.Err("password incorrect", ctx, err)
		return "", errors.NewBadRequestErr(err.Error(), 400)
	}

	//login user was verified passed
	return GenerateToken(cdtl.Username, nil)
}

func GenerateToken(username string, roles []string) (string, error) {
	claims := &JWTCustomClaims{
		username,
		roles,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
			Issuer:    DEFAULT_ISSUER,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SECRET_KEY))
}

func GetToken(ctx *gin.Context) (string, error) {
	authHeader := ctx.GetHeader(HEADER_ATTR_AUTHORIZATION)
	if !strings.HasPrefix(authHeader, BEARER_SCHEMA) {
		return "", errors.NewBadRequestErr("invaild token", 400)
	}
	return authHeader[len(BEARER_SCHEMA):], nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(SECRET_KEY), nil
	})
}
