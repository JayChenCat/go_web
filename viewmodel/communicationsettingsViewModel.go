package viewmodel

type CommunicationsettingsViewModel struct {
	Id                string `json:"id"`                //编号(主键)
	CommunicationType int    `json:"communicationType"` //通讯类型(1表示本机，2表示OSC)
	IP                string `json:"ip"`                //IP地址
	Port              string `json:"port"`              //端口号
}
