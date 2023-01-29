package dto

type UpdateSysUserStatusRequest struct {
	ID     int64 `json:"id" form:"id"`         // id
	Status int32 `json:"status" form:"status"` // 状态 -1 Deteled 0：停用   1：正常
}
