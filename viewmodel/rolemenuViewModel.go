package viewmodel

//角色权限菜单信息表
type RolemenuViewModel struct {
	Id           string `json:"id"`           //编号(主键)
	RoleId       string `json:"roleId"`       //角色ID，外键，关联角色信息
	MenuId       string `json:"menuId"`       //菜单ID,外键,关联菜单信息表
	Jurisdiction int    `json:"jurisdiction"` //权限标识(1表示添加，2表示删除,4表示修改，8表示查询，16表示导入，32导出，依次2的幂次方)
}
