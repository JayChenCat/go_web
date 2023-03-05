package viewmodel

//行驶记录设置信息表
type DrivingrecordViewModel struct {
	Id     string `json:"id"`     //编号(主键)
	Name   string `json:"name"`   //名称
	IsOpen int    `json:"isOpen"` //开关，1开启，0关闭
}
