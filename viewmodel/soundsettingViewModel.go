package viewmodel

//TTS语音配置信息表
type SoundsettingViewModel struct {
	Id          string `json:"id"`          //编号(主键)
	SoundType   int    `json:"soundType"`   //声音类型(1表示声音类型1，2声音类型2，4表示声音类型3，0表示自定义类型)
	Volume      int    `json:"volume"`      //音量
	CustomSound string `json:"customSound"` //自定义的声音类型名称(备用字段)
}
