package domain

//车辆设置信息表
type Carsetting struct {
	Id           int64      `json:"id"`           //编号(主键)
	LocomotiveID string     `json:"locomotiveID"` //机车号
	HostSide     int        `json:"hostSide"`     //主机所在端,1:一端,2:二端
	CommonFileld BaseFileld `json:"commonfileld"` //公共字段(创建时间，创建人等属性)
}
