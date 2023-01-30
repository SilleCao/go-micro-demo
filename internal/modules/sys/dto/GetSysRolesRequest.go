package dto

type GetSysRolesRequest struct {
	ID     int64         `json:"-"`                    // id
	Name   string        `json:"name" form:"name"`     // 角色名称
	Remark string        `json:"remark" form:"remark"` // 备注
	DeptID []interface{} `json:"deptId" form:"deptId"` // 部门ID
}
