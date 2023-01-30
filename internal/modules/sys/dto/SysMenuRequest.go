package dto

type SysMenuRequest struct {
	ID          int64  `json:"-"`                              // id
	Pid         int64  `json:"pid" form:"pid"`                 // 上级ID，一级菜单为0
	Name        string `json:"name" form:"name"`               // 名称
	URL         string `json:"url" form:"url"`                 // 菜单URL
	Permissions string `json:"permissions" form:"permissions"` // 授权(多个用逗号分隔，如：sys:user:list,sys:user:save)
	Type        int32  `json:"type" form:"type"`               // 类型   0：菜单   1：按钮
	Icon        string `json:"icon" form:"icon"`               // 菜单图标
	Sort        int32  `json:"sort" form:"sort"`               // 排序
}
