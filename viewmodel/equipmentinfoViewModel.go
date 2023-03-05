package viewmodel

type EquipmentinfoViewModel struct {
	Id       string `json:"id"`       //编号(主键)
	Sn       string `json:"sn"`       //序列号
	Module   string `json:"module"`   //所属功能模块(1表示DMS)
	DeviceIP string `json:"deviceIP"` //设备IP
	Language string `json:"language"` //系统语言(1表示zh_cn，2表示en)
}
