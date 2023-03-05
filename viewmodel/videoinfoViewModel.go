package viewmodel

//点播视频信息表(备用表)
type VideoinfoViewModel struct {
	Id            string  `json:"id"`            //编号(主键)
	VideoFileName string  `json:"videoFileName"` //视频文件名称
	VideoType     int     `json:"videoType"`     //视频类型,（1表示告警视频，2表示长视频）
	VideoAddress  string  `json:"videoAddress"`  //视频存放地址
	Remarks       *string `json:"remarks"`       //备注
}
