package domain

//升级记录信息表(备用表)
type Upgraderecord struct {
	Id              int64      `json:"id"`              //编号(主键)
	UpgradeName     string     `json:"upgradeName"`     //升级的设备名称
	UpgradeFileName string     `json:"upgradeFileName"` //升级文件包名称
	UpgradeVersion  string     `json:"upgradeVersion"`  //升级的版本
	CommonFileld    BaseFileld `json:"commonfileld"`    //公共字段(创建时间，创建人等属性)
}
