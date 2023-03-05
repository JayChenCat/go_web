package domain

//摄像头设置
type Camerasetting struct {
	Id           int64      `json:"id"`           //编号(主键)
	Camera_Name  string     `json:"camera_Name"`  //摄像头名称
	Camera_OnOff int        `json:"camera_OnOff"` //摄像头开关，1开启，0关闭
	Camera_IP    string     `json:"camera_IP"`    //摄像头IP地址
	ChannelNo    int        `json:"channelNo"`    //通道号
	ChannelName  string     `json:"channelName"`  //通道名称
	CommonFileld BaseFileld `json:"commonfileld"` //公共字段(创建时间，创建人等属性)
}
