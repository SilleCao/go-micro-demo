package dto

import (
	"time"
)

type SysUserResponse struct {
	ID         int64     `json:"id"`         // id
	Username   string    `json:"username"`   // 用户名
	RealName   string    `json:"realName"`   // 姓名
	HeadURL    string    `json:"headUrl"`    // 头像
	Gender     int32     `json:"gender"`     // 性别   0：男   1：女    2：保密
	Email      string    `json:"email"`      // 邮箱
	Mobile     string    `json:"mobile"`     // 手机号
	DeptID     int64     `json:"deptId"`     // 部门ID
	SuperAdmin int32     `json:"superAdmin"` // 超级管理员   0：否   1：是
	Status     int32     `json:"status"`     // 状态  0：停用   1：正常
	Creator    int64     `json:"creator"`    // 创建者
	CreateDate time.Time `json:"createDate"` // 创建时间
	Updater    int64     `json:"updater"`    // 更新者
	UpdateDate time.Time `json:"updateDate"` // 更新时间
}
