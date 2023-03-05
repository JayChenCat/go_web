package domain

type Accountmanager struct {
	Id           int64      `json:"id"`           //编号(主键)
	AccountType  int        `json:"accountType"`  //账户类型(1表示主机FTP账号，2表示OCS FTP账号，4表示RTSP账号)
	User         string     `json:"user"`         //用户名
	Pwd          string     `json:"Pwd"`          //密码
	CommonFileld BaseFileld `json:"commonfileld"` //公共字段(创建时间，创建人等属性)
}
