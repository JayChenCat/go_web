package domain

//点播视频信息表(备用表)
type Videoinfo struct {
	Id            int64      `json:"id"`            //编号(主键)
	VideoFileName string     `json:"videoFileName"` //视频文件名称
	VideoType     int        `json:"videoType"`     //视频类型,（1表示告警视频，2表示长视频）
	VideoAddress  string     `json:"videoAddress"`  //视频存放地址
	CommonFileld  BaseFileld `json:"commonfileld"`  //公共字段(创建时间，创建人等属性)
}
