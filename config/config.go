package config

import "path/filepath"

// dir name
const (
	Front  = "html"
	Web    = "web"
	Common = "base"

	IndexKey    = "index"
	AccountKey  = "account"
	AddTokenKey = "token"
)

var (
	//模板界面
	CommonPath = filepath.Join(Front, Web, Common, "common.html")
	//登录界面
	SignInFormPath = filepath.Join(Front, "login.html")
	//用户信息列表界面
	ManagerPath = filepath.Join(Front, Web, "manager", "manager.html")
	//添加用户信息界面
	AddmanagerPath = filepath.Join(Front, Web, "manager", "addmanager.html")
	//首页界面
	IndexPath = filepath.Join(Front, Web, "index", "index.html")
	//系统设置界面
	SystemPath = filepath.Join(Front, Web, "system", "system.html")
	//升级界面
	UpgradePath = filepath.Join(Front, Web, "update", "upgrade.html")
	//查看视频界面
	VideoPath = filepath.Join(Front, Web, "video", "video.html")
)
