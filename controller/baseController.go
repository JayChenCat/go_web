package controller

import (
	"go_web/config"
	"go_web/domain"
	"go_web/util"
	"log"
	"net/http"
)

func GenerateSnowFlakeId(n int64) int64 {
	worker, erro := util.NewWorker(n)
	if erro != nil {
		log.Printf("生成ID错误 error: %v", erro.Error())
	}
	return worker.GetId()
}

// 跳转界面(判断Session是否失效)
func skiplinks(w http.ResponseWriter, r *http.Request, m map[string]interface{}, filePath string) {
	s, err := util.GetSession(r)
	if err != nil {
		log.Printf("ListOrders GetSession error: %v", err.Error())
	} else {
		re, ok := s.Get(config.AccountKey)
		if ok {
			account, ok := re.(*domain.Account)
			if ok {
				m["UserName"] = account.UserName
				err := util.RenderWithAccountAndCommonTem(w, r, m, filePath)
				if err != nil {
					log.Printf("error: %v", err.Error())
				}
			}
		} else {
			//重定向到登录界面
			skiplinksData(w, nil, config.SignInFormPath)
		}
	}
}

//带有数据的html界面渲染
func skiplinksData(w http.ResponseWriter, data interface{}, filePath string) {
	err := util.RenderWithCommon(w, data, filePath)
	if err != nil {
		log.Printf("view signInForm error: %v", err.Error())
	}
}

//带有数据的渲染
func skipLinkByData(w http.ResponseWriter, r *http.Request, m map[string]interface{}, filePath string) {
	err := util.RenderWithAccountAndCommonTem(w, r, m, filePath)
	if err != nil {
		log.Printf("ListOrders RenderWithAccountAndCommonTem error: %v", err.Error())
	}
}

// session 中保存 任何对象信息
func SaveInfo(w http.ResponseWriter, r *http.Request, tag string, data interface{}) error {
	s, err := util.GetSession(r)
	if err != nil {
		log.Printf("get session error: %v", err.Error())
		return err
	} else {
		if s != nil {
			err = s.Save(tag, data, w, r)
			if err != nil {
				log.Printf("get session error: %v", err.Error())
				return err
			}
		}
	}
	return nil
}

func GetToken(key string, r *http.Request) string {
	s, err := util.GetSession(r)
	if err != nil {
		log.Printf("ListOrders GetSession error: %v", err.Error())
	}
	re, ok := s.Get(key)
	if ok {
		token, ok := re.(string)
		if ok {
			return token
		} else {
			return ""
		}
	} else {
		return ""
	}
}

//移除Session
func RemoveSession(w http.ResponseWriter, r *http.Request, key string) error {
	s, err := util.GetSession(r)
	if err != nil {
		log.Printf("get session error: %v", err.Error())
	}
	if s != nil {
		err = s.Del(key, w, r)
		if err != nil {
			log.Printf("session delete error: %v", err.Error())
		}
	}
	return err
}
