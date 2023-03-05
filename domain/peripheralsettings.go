package domain

//外设设置信息表
type Peripheralsettings struct {
	Id               int64      `json:"id"`               //编号(主键)
	PeripheralType   int        `json:"peripheralType"`   //外设类型(1表示GPS,2表示WiFi，4表示4G)
	Peripheral_OnOff int        `json:"peripheral_OnOff"` //开关，1开启，0关闭
	CommonFileld     BaseFileld `json:"commonfileld"`     //公共字段(创建时间，创建人等属性)
}
