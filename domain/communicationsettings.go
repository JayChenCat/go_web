package domain

//通讯地址设置信息表
type Communicationsettings struct {
	Id                int64      `json:"id"`                //编号(主键)
	CommunicationType int        `json:"communicationType"` //通讯类型(1表示本机，2表示OSC)
	IP                string     `json:"ip"`                //IP地址
	Port              int64      `json:"port"`              //端口号
	CommonFileld      BaseFileld `json:"commonfileld"`      //公共字段(创建时间，创建人等属性)
}
