package service

import (
	"fmt"
	"time"

	"github.com/SilleCao/golang/go-micro-demo/internal/modules/auth/service"
	"github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/dto"
	"github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/model"
	repo "github.com/SilleCao/golang/go-micro-demo/internal/modules/sys/repository"
	"github.com/SilleCao/golang/go-micro-demo/internal/pkg/common"
	"github.com/SilleCao/golang/go-micro-demo/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(ctx *gin.Context, userReq *dto.CreateSysUserRequest) error {

	user := model.SysUser{}
	copier.Copy(&user, userReq)
	su, _ := GetLoginUser(ctx)
	user.CreateDate = time.Now()
	user.UpdateDate = time.Now()
	user.Creator = su.ID
	user.Updater = su.ID
	user.SuperAdmin = common.UserIsSuperAdmin_FALSE
	user.Status = common.UserStatus_Locked

	//crypt user's password
	cryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("generate crypt password fail: %v", err)
	}
	user.Password = string(cryptedPassword)
	return repo.CreateUser(ctx, &user)
}

func UpdateUser(ctx *gin.Context, m map[string]interface{}) error {
	consolidateData(ctx, m, model.SysUser{})
	return repo.UpdateUser(ctx, m)
}

func GetUsers(ctx *gin.Context, userReq *dto.GetSysUsersRequest, pagination *common.Pagination) (*common.Pagination, error) {
	user := model.SysUser{}
	copier.Copy(&user, userReq)

	pagination, err := repo.GetUsers(ctx, &user, pagination)
	if err != nil {
		return pagination, fmt.Errorf("get users fail: %v", err)
	}
	sysUserResps := []dto.SysUserResponse{}
	copier.Copy(&sysUserResps, pagination.Content)
	pagination.Content = sysUserResps
	return pagination, err
}

func GetUserById(ctx *gin.Context, id int64) (*dto.SysUserResponse, error) {
	su, err := repo.GetUserById(ctx, id)
	if err != nil {
		logger.Err("get user by id fail", ctx, err)
		return nil, err
	}
	sur := dto.SysUserResponse{}
	copier.Copy(&sur, su)
	return &sur, nil
}

func UpdateUserStatus(ctx *gin.Context, userReq map[string]interface{}) error {
	// if !CheckLogUserIsAdmin(ctx) {
	// 	return errors.NewUnauthorizedErr("only super admin can lock user", 401)
	// }
	return UpdateUser(ctx, userReq)
}

func CheckLogUserIsAdmin(ctx *gin.Context) bool {
	su, _ := GetLoginUser(ctx)
	return su.SuperAdmin == 1
}

func GetLoginUser(ctx *gin.Context) (*model.SysUser, error) {
	tokenStr, _ := service.GetToken(ctx)
	token, _ := service.ValidateToken(tokenStr)
	return repo.GetUserByUsername(ctx, fmt.Sprintf("%v", token.Claims.(jwt.MapClaims)["name"]))
}
