package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/SilleCao/golang/go-micro-demo/internal/modules/auth/model"
	"github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/repository"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type JWTCustomClaims struct {
	Name  string   `json:"name"`
	Roles []string `json:"roles"`
	jwt.StandardClaims
}

const DEFAULT_ISSUER = "sille.cn"
const SECRET_KEY = "0000"

func Authenticate(cdtl *model.Credentials, ctx context.Context) (string, error) {
	su, err := repository.GetUserByUsername(cdtl.Username, ctx)
	if err != nil {
		log.Println(err)
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(su.Password), []byte(cdtl.Password)); err != nil {
		log.Println(err)
		return "", err
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

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(SECRET_KEY), nil
	})
}
