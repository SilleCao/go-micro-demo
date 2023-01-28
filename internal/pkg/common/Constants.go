package common

type ctxKey string

const REQ_ID = ctxKey("X-Req-ID")

const (
	UserStatus_Deleted = -1
	UserStatus_Locked  = 0
	UserStatus_Normal  = 1
)
