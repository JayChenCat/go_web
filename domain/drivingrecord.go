package domain

//行驶记录设置信息表
type Drivingrecord struct {
	Id           int64      `json:"id"`           //编号(主键)
	Name         string     `json:"name"`         //名称
	IsOpen       int        `json:"isOpen"`       //开关，1开启，0关闭
	CommonFileld BaseFileld `json:"commonfileld"` //公共字段(创建时间，创建人等属性)
}
