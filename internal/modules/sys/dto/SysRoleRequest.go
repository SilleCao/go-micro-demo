package dto

type SysRoleRequest struct {
	ID     int64  `json:"-"`                    // id
	Name   string `json:"name" form:"name"`     // 角色名称
	Remark string `json:"remark" form:"remark"` // 备注
	DeptID int64  `json:"deptId" form:"deptId"` // 部门ID
}
