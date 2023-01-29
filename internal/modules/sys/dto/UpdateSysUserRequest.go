package dto

type UpdateSysUserRequest struct {
	ID         int64  `json:"id"`                           // id
	Username   string `json:"username" form:"username" `    // 用户名
	RealName   string `json:"realName" form:"realName"`     // 姓名
	HeadURL    string `json:"headUrl" form:"headUrl"`       // 头像
	Gender     int32  `json:"gender" form:"gender"`         // 性别   0：男   1：女    2：保密
	Email      string `json:"email" form:"email"`           // 邮箱
	Mobile     string `json:"mobile" form:"mobile"`         // 手机号
	DeptID     int64  `json:"deptId" form:"deptId"`         // 部门ID
	Status     int32  `json:"status" form:"status"`         // 状态 -1 Deteled 0：停用   1：正常
	SuperAdmin int32  `json:"superAdmin" form:"superAdmin"` // 状态 -1 Deteled 0：停用   1：正常
}
