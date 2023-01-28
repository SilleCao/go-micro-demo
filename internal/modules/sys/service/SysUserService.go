package service

import (
	"fmt"
	"time"

	"github.com/SilleCao/golang/go-micro-demo/internal/modules/auth/service"
	"github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/dto"
	"github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/model"
	repo "github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/repository"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/common"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(ctx *gin.Context, user *model.SysUser) error {
	su, _ := GetLoginUser(ctx)
	user.CreateDate = time.Now()
	user.UpdateDate = time.Now()
	user.Creator = su.ID
	user.Updater = su.ID

	//crypt user's password
	cryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("generate crypt password fail: %v", err)
	}
	user.Password = string(cryptedPassword)
	return repo.CreateUser(ctx, user)
}

func UpdateUser(ctx *gin.Context, user *model.SysUser) error {
	su, _ := GetLoginUser(ctx)
	user.UpdateDate = time.Now()
	user.Updater = su.ID
	return repo.UpdateUser(ctx, user)
}

func GetUsers(ctx *gin.Context, pagination *common.Pagination) (*common.Pagination, error) {
	pagination, err := repo.GetUsers(ctx, pagination)
	if err != nil {
		return pagination, fmt.Errorf("get users fail: %v", err)
	}
	// sud := []model.SysUserDTO{}
	// copier.Copy(&sud, pagination.Content)
	// pagination.Content = sud
	return pagination, err
}

func GetUserById(ctx *gin.Context, id int64) (*model.SysUser, error) {
	return repo.GetUserById(ctx, id)
}

func GetUserByUsername(ctx *gin.Context, username string) (*model.SysUser, error) {
	return repo.GetUserByUsername(ctx, username)
}

func UpdateUserStatus(ctx *gin.Context, uerDto dto.UpdateUserStatusDTO) error {
	if !CheckLogUserIsAdmin(ctx) {
		return errors.NewUnauthorizedErr("only super admin can lock user", 401)
	}
	user := model.SysUser{}
	copier.Copy(&user, uerDto)
	return UpdateUser(ctx, &user)
}

func CheckLogUserIsAdmin(ctx *gin.Context) bool {
	su, _ := GetLoginUser(ctx)
	return su.SuperAdmin == 1
}

func GetLoginUser(ctx *gin.Context) (*model.SysUser, error) {
	tokenStr, _ := service.GetToken(ctx)
	token, _ := service.ValidateToken(tokenStr)
	return GetUserByUsername(ctx, fmt.Sprintf("%v", token.Claims.(jwt.MapClaims)["name"]))
}
