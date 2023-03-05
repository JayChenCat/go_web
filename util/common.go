package util

import (
	"encoding/json"
	_ "encoding/json"
	"errors"
	"go_web/domain"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

/*创建文件目录，创建文件等工具公用方法*/

//根据名称创建文件目录
func CreateByNameDir(name string) string {
	folderPath := filepath.Join(name)
	createdir(folderPath)
	return folderPath
}

//根据当前日期来创建文件夹
func CreateDateDir(path string) string {
	folderName := time.Now().Format("20060102")
	folderPath := filepath.Join(path, folderName)
	createdir(folderPath)
	return folderPath
}

//创建目录公共方法(文件夹默认项目根目录)
func createdir(folderPath string) {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 创建文件夹并修改权限
		os.Mkdir(folderPath, os.ModePerm) //0777也可以os.ModePerm
		os.Chmod(folderPath, os.ModePerm)
	}
}

//判断数据库连接错误
func checkDataBaseErro(err error) bool {
	if err != nil {
		//数据库连接错误 10061 dial tcp 119.91.142.127:3306: connectex: No connection could be made because the target machine actively refused it.
		errs := errors.New("connectex").Error()
		errsDesc := err.Error()
		if strings.Contains(errsDesc, errs) {
			return false
		}
	}
	return true
}

//处理公共http请求并返回json格式数据(ajax调用)
func ResponseHtml(w http.ResponseWriter, msg *domain.Result, err error) {
	databseConnErr := checkDataBaseErro(err)
	result_json, _ := json.Marshal(msg)
	if !databseConnErr {
		msg.Code = 2
		msg.Msg = "mysql数据库连接错误，请检查数据库连接字符及数据库服务是否启动!"
		result_json, _ = json.Marshal(msg)
	}
	jsonText := string(result_json)
	io.WriteString(w, jsonText)
}

//返回随机字符串
func GetRandToken(n int) (ret string) {
	textString := "awertyuiopasdfghjklzxcvbnm0123456789"
	for i := 0; i < n; i++ {
		r := rand.Intn(len(textString))
		ret = ret + textString[r:r+1]
	}
	return
}
