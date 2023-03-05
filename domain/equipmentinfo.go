package domain

//设备基本信息
type Equipmentinfo struct {
	Id           int64      `json:"id"`           //编号(主键)
	Sn           string     `json:"sn"`           //序列号
	Module       int        `json:"module"`       //所属功能模块(1表示DMS)
	DeviceIP     string     `json:"deviceIP"`     //设备IP
	Language     int        `json:"language"`     //系统语言(1表示zh_cn，2表示en)
	CommonFileld BaseFileld `json:"commonfileld"` //公共字段(创建时间，创建人等属性)
}
