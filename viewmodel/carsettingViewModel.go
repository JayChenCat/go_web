package viewmodel

//车辆设置信息表
type CarsettingViewModel struct {
	Id           string `json:"id"`           //编号(主键)
	LocomotiveID string `json:"locomotiveID"` //机车号
	HostSide     int    `json:"hostSide"`     //主机所在端,1:一端,2:二端
}
