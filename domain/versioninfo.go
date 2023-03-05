package domain

//版本信息
type Versioninfo struct {
	Id             int64      `json:"id"`             //编号(主键)
	ArmVersion     string     `json:"armVersion"`     //设备应用程序版本
	UpgradeVersion string     `json:"upgradeVersion"` //升级程序版本
	WebVersion     string     `json:"webVersion"`     //Web应用程序版本
	WebServerName  string     `json:"webServerName"`  //Web 服务器名称
	CommonFileld   BaseFileld `json:"commonfileld"`   //公共字段(创建时间，创建人等属性)
}
