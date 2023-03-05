package viewmodel

//疲劳驾驶报警信息表
type FatiguedrivingalarmViewModel struct {
	Id               string  `json:"id"`               //编号(主键)
	AlarmTypeName    string  `json:"alarmTypeName"`    //报警类型名称
	Sort             int     `json:"sort"`             //排序
	Speed            float64 `json:"speed"`            //报警触发车速
	Time             float64 `json:"time"`             //报警触发时间(单位:秒)
	TTS              string  `json:"tts"`              //TTS告警语音
	CD               float64 `json:"cd"`               //警报/上传冷却时间(单位:秒)
	AlarmSw          int     `json:"alarmSw"`          //报警总开关，1表示开启，0关闭
	AlarmSoundSw     int     `json:"alarmSoundSw"`     //报警声音开关,1表示开启，0关闭
	UploadingPicSw   int     `json:"uploadingPicSw"`   //上传报警图片开关，1表示开启，0关闭
	UploadingVideoSw int     `json:"uploadingVideoSw"` //上传报警视频开关，1表示开启，0关闭
}
