package viewmodel

//角色信息表
type RoleViewModel struct {
	Id       int64  `json:"id"`       //编号(主键)
	RoleName string `json:"roleName"` //角色名称
}
