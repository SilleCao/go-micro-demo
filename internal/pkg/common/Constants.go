package common

type ctxKey string

const REQ_ID = ctxKey("X-Req-ID")

const (
	UserStatus_Deleted = -1
	UserStatus_Locked  = 0
	UserStatus_Normal  = 1
)

const (
	UserIsSuperAdmin_FALSE = 0
	UserIsSuperAdmin_TURE  = 1
)

const (
	//15d
	UserDataCacheExpiration = "360h"
)
