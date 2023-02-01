package dto

type SysDeptRequest struct {
	ID   int64  `json:"-" form:"id"`      // id
	Pid  int64  `json:"pid" form:"pid"`   // 上级ID
	Pids string `json:"pids" form:"pids"` // 所有上级ID，用逗号分开
	Name string `json:"name" form:"name"` // 部门名称
	Sort int32  `json:"sort" form:"sort"` // 排序
}
