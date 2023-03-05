package viewmodel

//升级记录信息表(备用表)
type UpgraderecordViewModel struct {
	Id              string `json:"id"`              //编号(主键)
	UpgradeName     string `json:"upgradeName"`     //升级的设备名称
	UpgradeFileName string `json:"upgradeFileName"` //升级文件包名称
	UpgradeVersion  string `json:"upgradeVersion"`  //升级的版本
}
