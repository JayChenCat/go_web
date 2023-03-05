package controller

import (
	"go_web/config"
	"go_web/domain"
	"go_web/service"
	"go_web/util"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

//跳转升级设置主页
func ViewUpgrade(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 跳转到 升级设置主页
		skiplinks(w, r, make(map[string]interface{}), config.UpgradePath)
	} else {
		Upgrade(w, r)
	}
}

func Upgrade(w http.ResponseWriter, request *http.Request) {
	flag := true
	layout := "2006-01-02 15:04:05"
	time := time.Now()
	timeNowText := time.Format("20060102150405")
	erro := request.ParseMultipartForm(1048576)
	if erro != nil {
		log.Printf("升级参数解析 error: %v", erro.Error())
		flag = false
	}
	fromData := request.MultipartForm.Value
	remarks := fromData["remarks"][0]
	// 根据请求body创建一个json解析器实例
	/*decoder := json.NewDecoder(request.Body)
	// 用于存放参数key=value数据
	var params map[string]string
	// 解析参数 存入map
	decoder.Decode(&params)*/

	//remarks := request.PostFormValue("remarks")
	//接收客户端传来的文件 uploadfile 与客户端保持一致
	file, handler, err := request.FormFile("uploadfile")
	if err != nil {
		log.Printf("file upload error: %v", err.Error())
		flag = false
	}
	msg := new(domain.Result)
	if file != nil {
		defer file.Close()
		//上传的文件保存在download路径下
		ext := path.Ext(handler.Filename)                        //获取文件后缀，如:.txt
		filenameall := path.Base(handler.Filename)               //获取文件全名称包括扩展名
		fileprefix := filenameall[0 : len(filenameall)-len(ext)] //获取文件名称不包括扩展名
		fileNewName := string(timeNowText) + strconv.Itoa(time.Nanosecond()) + ext
		filePath := util.CreateByNameDir("download") + "/" + fileNewName
		f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
		defer f.Close()
		if err != nil {
			log.Printf("上传的文件保存在download路径下 error: %v", err.Error())
			flag = false
		} else {
			io.Copy(f, file)
			_time := time.Format(layout)
			commonFileld := domain.BaseFileld{
				CREATED_BY:   "admin", //创建人
				CREATED_TIME: _time,   //创建时间
				UPDATED_BY:   "admin", //更新人
				UPDATED_TIME: _time,   //更新时间
				IsDeleted:    0,       //是否删除(逻辑删除)，1是，0否
				Remarks:      remarks, //备注
			}
			upgraderecord := &domain.Upgraderecord{
				Id:              GenerateSnowFlakeId(1),
				UpgradeName:     "DMS",
				UpgradeFileName: fileprefix,
				UpgradeVersion:  "v1.0.0",
				CommonFileld:    commonFileld}
			success, erro := service.InsertUpgraderecord(upgraderecord)
			if erro != nil {
				log.Printf("提交升级记录到数据库 error: %v", erro.Error())
			} else {
				if flag && success {
					msg.Code = 0
					msg.Msg = "升级成功！"
				} else {
					msg.Code = 1
					msg.Msg = "升级失败！"
				}
			}
		}
	} else {
		msg.Code = 2
		msg.Msg = "上传的文件为空！"
	}
	util.ResponseHtml(w, msg, err)
}
