package viewmodel

type AccountmanagerViewModel struct {
	Id          string `json:"id"`          //编号(主键)
	AccountType int    `json:"accountType"` //账户类型(1表示主机FTP账号，2表示OCS FTP账号，4表示RTSP账号)
	User        string `json:"user"`        //用户名
	Pwd         string `json:"Pwd"`         //密码
}
