package route

/**
将路由改为 map 映射的配置
*/

import (
	"go_web/controller"
	"net/http"
)

// 路由映射注册表
var route = map[string]http.HandlerFunc{
	// account
	"/login":   controller.ViewLoginOrPostLogin,
	"/signOut": controller.SignOut,
	//用户信息的路由-account
	"/manager":       controller.ViewManager,
	"/addmanager":    controller.ViewAddManager,
	"/token":         controller.SubmitToken,
	"/SingleAccount": controller.SingleAccount,
	"/DeleteUser":    controller.DeleteUser,
	// view
	"/index":   controller.ViewMain,
	"/system":  controller.ViewSystem,
	"/setting": controller.SubmitSetting,
	"/video":   controller.ViewVideo,
	"/upgrade": controller.ViewUpgrade,
}

// 注册路由
func RegisterRoute() {
	for k, v := range route {
		http.HandleFunc(k, v)
	}
}
