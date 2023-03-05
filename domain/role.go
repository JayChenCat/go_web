package domain

//角色信息表
type Role struct {
	Id           int64      `json:"id"`           //编号(主键)
	RoleName     string     `json:"roleName"`     //角色名称
	CommonFileld BaseFileld `json:"commonfileld"` //公共字段(创建时间，创建人等属性)
}
