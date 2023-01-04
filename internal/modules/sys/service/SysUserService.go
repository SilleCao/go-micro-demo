package service

import (
	"context"
	"time"

	"github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/model"
	repo "github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/repository"
)

func CreateUser(user *model.SysUser, ctx context.Context) (err error) {
	user.CreateDate = time.Now()
	user.UpdateDate = time.Now()
	return repo.CreateUser(user, ctx)
}
