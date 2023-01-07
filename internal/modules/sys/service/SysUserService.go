package service

import (
	"context"
	"log"
	"time"

	"github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/model"
	repo "github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/repository"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/common"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user *model.SysUser, ctx context.Context) (err error) {
	user.CreateDate = time.Now()
	user.UpdateDate = time.Now()
	//crypt user's password
	cryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	user.Password = string(cryptedPassword)
	return repo.CreateUser(user, ctx)
}

func GetUsers(page int, size int, ctx context.Context) (*common.Pagination, error) {
	pagination := &common.Pagination{
		Page: page,
		Size: size,
	}
	return repo.GetUsers(pagination, ctx)
}

func GetUserByUsername(username string, ctx context.Context) (*model.SysUser, error) {
	return repo.GetUserByUsername(username, ctx)
}
