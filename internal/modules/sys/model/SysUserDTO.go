package model

import "time"

type SysUserDTO struct {
	ID         int64     `json:"id" form:"id"`                 // id
	Username   string    `json:"username" form:"username" `    // 用户名
	RealName   string    `json:"realName" form:"realName"`     // 姓名
	HeadURL    string    `json:"headUrl" form:"headUrl"`       // 头像
	Gender     int32     `json:"gender" form:"gender"`         // 性别   0：男   1：女    2：保密
	Email      string    `json:"email" form:"email"`           // 邮箱
	Mobile     string    `json:"mobile" form:"mobile"`         // 手机号
	DeptID     int64     `json:"deptId" form:"deptId"`         // 部门ID
	SuperAdmin int32     `json:"superAdmin" form:"superAdmin"` // 超级管理员   0：否   1：是
	Status     int32     `json:"status" form:"status"`         // 状态  0：停用   1：正常
	Creator    int64     `json:"creator" form:"creator"`       // 创建者
	CreateDate time.Time `json:"createDate" form:"createDate"` // 创建时间
	Updater    int64     `json:"updater" form:"updater"`       // 更新者
	UpdateDate time.Time `json:"updateDate" form:"updateDate"` // 更新时间
}
