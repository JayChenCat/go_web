package viewmodel

//摄像头设置
type CamerasettingViewModel struct {
	Id           string `json:"id"`           //编号(主键)
	Camera_Name  string `json:"camera_Name"`  //摄像头名称
	Camera_OnOff int    `json:"camera_OnOff"` //摄像头开关，1开启，0关闭
	Camera_IP    string `json:"camera_IP"`    //摄像头IP地址
	ChannelNo    int    `json:"channelNo"`    //通道号
	ChannelName  string `json:"channelName"`  //通道名称
}
