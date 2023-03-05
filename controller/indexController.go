package controller

import (
	"go_web/config"
	"go_web/service"
	"log"
	"net/http"
)

// 跳转 首页主页
func ViewMain(w http.ResponseWriter, r *http.Request) {
	//{{range .equipmentinfo}} {{end}}
	equipmentinfo, err := service.GetSingleEquipmentinfo(188888899886666)
	versioninfo, verr := service.GetSingleVersioninfo(898988988888)
	//equipmentinfoList, err := service.GetEquipmentinfoPageList(1, 1000)
	equipmentOrVersioninfo := map[string]interface{}{
		"equipmentinfo": nil,
		"versioninfo":   nil,
	}
	if err != nil {
		log.Printf("绑定首页设备基本信息 error: %v", err.Error())
	}
	if verr != nil {
		log.Printf("绑定首页版本信息 error: %v", err.Error())
	} else {
		equipmentOrVersioninfo = map[string]interface{}{
			"equipmentinfo": equipmentinfo,
			"versioninfo":   versioninfo,
		}
	}
	// 跳转到首页
	skiplinks(w, r, equipmentOrVersioninfo, config.IndexPath)
}
