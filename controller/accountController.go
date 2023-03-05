package controller

import (
	"encoding/json"
	"errors"
	"go_web/config"
	"go_web/domain"
	"go_web/service"
	"go_web/util"
	"go_web/viewmodel"
	"log"
	"net/http"
	"strconv"
	"time"
)

// get 跳转到登录页面 或者 post 登录
func ViewLoginOrPostLogin(w http.ResponseWriter, r *http.Request) {
	msg := new(domain.Result)
	msg.Init()
	err := errors.New("")
	if r.Method == "GET" {
		skiplinksData(w, nil, config.SignInFormPath)
	} else {
		err1 := r.ParseForm()
		if err1 != nil {
			err = err1
		}
		userName := r.FormValue("username")
		password := r.FormValue("password")
		//userName == "admin" && password == "aceb162a7bc7d7995529121c9cd647f5"
		login, erro := service.Login(userName, password)
		if erro != nil {
			err = erro
		} else {
			if login != nil {
				if len(login.UserName) > 0 {
					/*m := map[string]interface{}{
						"UserName": userName,
					}*/
					account := domain.Account{
						UserName: login.UserName}
					//将登录的用户信息存入Session
					cookieErr := SaveInfo(w, r, config.AccountKey, account)
					if cookieErr != nil {
						msg.Code = 2
						msg.Msg = "session 错误(清除浏览器缓存):" + cookieErr.Error()
					} else {
						msg.Code = 0
						msg.Msg = "登录成功!"
						//msg.Data = userName
					}
				} else {
					msg.Code = 1
					msg.Msg = "登录失败,用户名或密码错误!"
				}
			} else {
				msg.Code = 3
				msg.Msg = "获取的用户信息为空!"
			}
		}
		util.ResponseHtml(w, msg, err)
	}
}

// 退出登录
func SignOut(w http.ResponseWriter, r *http.Request) {
	err := RemoveSession(w, r, config.AccountKey)
	if err != nil {
		log.Printf("移除Session error: %v", err.Error())
	}
	// 重定向到登录界面
	skiplinks(w, r, make(map[string]interface{}), config.SignInFormPath)
}

//跳转用户列表信息主页
func ViewManager(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 跳转到 Manager主页
		//orderIdStr := util.GetParam(r, "orderId")[0]
		skiplinks(w, r, make(map[string]interface{}), config.ManagerPath)
	} else {
		GetPageListAccount(w, r)
	}
}

//跳转添加用户信息主页
func ViewAddManager(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		skiplinks(w, r, make(map[string]interface{}), config.AddmanagerPath)
	} else {
		AddOrUpdateUser(w, r)
	}
}

//添加，修改单个用户信息-post方式
func AddOrUpdateUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("parse login form error: %v", err.Error())
	}
	userId := r.FormValue("id")
	fromtoken := r.FormValue("token")
	userName := r.FormValue("username")
	password := r.FormValue("password")
	//confirmPassword := r.FormValue("password_confirm")
	email := r.FormValue("email")
	//date := r.FormValue("date")
	//_time, _ := time.Parse("2006-01-02 15:04:05", date)
	_time := time.Now().Format("2006-01-02 15:04:05")
	commonFileld := domain.BaseFileld{
		CREATED_BY:   "admin", //创建人
		CREATED_TIME: _time,   //创建时间
		UPDATED_BY:   "admin", //更新人
		UPDATED_TIME: _time,   //更新时间
		IsDeleted:    0,       //是否删除(逻辑删除)，1是，0否
		Remarks:      "",      //备注
	}
	account := &domain.Account{
		Id:            GenerateSnowFlakeId(1),
		UserName:      userName,
		PassWord:      password,
		Email:         email,
		NickName:      "",           //昵称
		Avatar:        "",           //用户头像地址
		FullName:      "",           //真实姓名(备用字段)
		Mobile:        "",           //手机号
		Sex:           0,            //性别(0 表示女 1表示男)
		Nation:        "汉",          //民族
		Address:       "",           //用户地址
		IsLock:        0,            //是否禁用 1禁用，0为不禁用
		QRCode:        "",           //二维码号(备用字段)
		RoleId:        0,            //所属角色(外键)
		DepartId:      0,            //所属部门(外键)
		PositionId:    0,            //所属职位(备用字段)
		CommonFileld:  commonFileld, //公共字段(创建时间，创建人等属性)
		LoginLastTime: _time,        //用户最后登录时间
	}
	msg := new(domain.Result)
	token := GetToken(config.AddTokenKey, r)

	if fromtoken == token {
		isSucc := false
		err := errors.New("")
		//添加
		if userId == "0" {
			isSucc, err = service.InsertUser(account)
		} else {
			isSucc, err = service.UpdateUser(account)
		}
		if err != nil {
			log.Printf("添加，修改用户信息 error: %v", err.Error())
		} else {
			if isSucc {
				msg.Code = 0
				msg.Msg = "添加成功！"
				//添加成功后移除token
				errr := RemoveSession(w, r, config.AddTokenKey)
				if errr != nil {
					log.Printf("移除Session error: %v", errr.Error())
				}
			} else {
				msg.Code = 1
				msg.Msg = "添加失败！"
				log.Printf("添加用户信息 error: %v", err.Error())
			}
		}
	} else {
		msg.Code = 3
		msg.Msg = "该信息已提交，请勿重复提交!"
	}
	util.ResponseHtml(w, msg, err)
}

//提交防重复的token-post方式
func SubmitToken(w http.ResponseWriter, r *http.Request) {
	msg := new(domain.Result)
	err := r.ParseForm()
	if err != nil {
		log.Printf("parse login form error: %v", err.Error())
	} else {
		fromtoken := util.GetRandToken(20) //r.FormValue("token")
		//将token信息存入Session
		cookieErr := SaveInfo(w, r, config.AddTokenKey, fromtoken)
		if cookieErr != nil {
			msg.Code = 2
			msg.Msg = "获取token 失败 session 错误:" + cookieErr.Error()
		} else {
			msg.Code = 0
			msg.Msg = "获取token成功！"
			msg.Data = fromtoken
		}
	}
	util.ResponseHtml(w, msg, err)
}

//获取用户列表信息-分页
func GetPageListAccount(w http.ResponseWriter, r *http.Request) {
	msg := new(domain.Result)
	err := r.ParseForm()
	if err != nil {
		log.Printf("GetPageListUser error: %v", err.Error())
	}
	pageNumber := r.FormValue("pagenumber")
	pageSize := r.FormValue("pagesize")
	num, err := strconv.Atoi(pageNumber)
	psize, err := strconv.Atoi(pageSize)
	total, err := service.GetAccountTotal()
	if err != nil {
		log.Printf("统计用户列表信息数量 error: %v", err.Error())
	} else {
		list, err := service.GetPageList(num, psize)
		if err != nil {
			msg.Code = 1
			msg.Msg = "获取用户列表信息失败！"
			log.Printf("获取用户列表信息 error: %v", err.Error())
		} else {
			msg.Code = 0
			msg.Msg = "获取用户列表信息成功！"
			accounts := viewmodel.ListAccount{
				Lists: list,
				Count: total,
			}
			dataJson, _ := deal(accounts)
			//解决int64转换精度丢失问题
			msg.Data = dataJson
		}
	}
	util.ResponseHtml(w, msg, err)
}

func deal(data viewmodel.ListAccount) (string, error) {
	account, err := json.Marshal(data)
	/*var list interface{}
	//d := json.NewDecoder(strings.NewReader(string(account)))
	d := json.NewDecoder(bytes.NewBuffer(account))
	d.UseNumber()
	err = d.Decode(&list)
	if err != nil {
		log.Printf("解决int64转换精度丢失问题 error: %v", err.Error())
	}
	jsonText, err := json.Marshal(list)
	if err != nil {
		log.Printf("解决int64转换精度丢失问题-json error: %v", err.Error())
	}*/
	return string(account), err
}

//获取单个用户信息
func SingleAccount(w http.ResponseWriter, r *http.Request) {
	msg := new(domain.Result)
	msg.Init()
	err := r.ParseForm()
	if err != nil {
		log.Printf("GetSingleUser error: %v", err.Error())
	}
	userId := r.FormValue("id")
	num, errs := strconv.ParseInt(userId, 0, 0)
	if errs != nil {
		log.Printf("userId转换整型的错误 error: %v", errs.Error())
	}
	account, err := service.GetByByPrimarykeyUserInfo(num)
	if err != nil {
		msg.Code = 1
		msg.Msg = "获取单个用户信息失败！"
		log.Printf("获取单个用户信息 error: %v", err.Error())
	} else {
		if account != nil {
			msg.Code = 0
			msg.Msg = "获取单个用户信息成功！"
			dataJson, _ := json.Marshal(account)
			msg.Data = string(dataJson)
		} else {
			msg.Code = 2
			msg.Msg = "获取单个用户信息为空！"
			msg.Data = account
		}
	}
	util.ResponseHtml(w, msg, err)
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("parse login form error: %v", err.Error())
	}
	userId := r.FormValue("id")
	num, errs := strconv.ParseInt(userId, 0, 0)
	if errs != nil {
		log.Printf("userId转换整型的错误 error: %v", errs.Error())
	}
	fromtoken := r.FormValue("token")
	userName := r.FormValue("username")
	password := r.FormValue("password")
	//confirmPassword := r.FormValue("password_confirm")
	email := r.FormValue("email")
	//date := r.FormValue("date")
	//_time, _ := time.Parse("2006-01-02 15:04:05", date)
	_time := time.Now().Format("2006-01-02 15:04:05")
	commonFileld := domain.BaseFileld{
		CREATED_BY:   "admin", //创建人
		CREATED_TIME: _time,   //创建时间
		UPDATED_BY:   "admin", //更新人
		UPDATED_TIME: _time,   //更新时间
		IsDeleted:    0,       //是否删除(逻辑删除)，1是，0否
		Remarks:      "",      //备注
	}
	account := &domain.Account{
		Id:            num,
		UserName:      userName,
		PassWord:      password,
		Email:         email,
		NickName:      "",           //昵称
		Avatar:        "",           //用户头像地址
		FullName:      "",           //真实姓名(备用字段)
		Mobile:        "",           //手机号
		Sex:           0,            //性别(0 表示女 1表示男)
		Nation:        "汉",          //民族
		Address:       "",           //用户地址
		IsLock:        0,            //是否禁用 1禁用，0为不禁用
		QRCode:        "",           //二维码号(备用字段)
		RoleId:        0,            //所属角色(外键)
		DepartId:      0,            //所属部门(外键)
		PositionId:    0,            //所属职位(备用字段)
		CommonFileld:  commonFileld, //公共字段(创建时间，创建人等属性)
		LoginLastTime: _time,        //用户最后登录时间
	}
	msg := new(domain.Result)
	token := GetToken(config.AddTokenKey, r)
	if fromtoken == token {
		isSucc, err := service.UpdateUser(account)
		if err != nil {
			log.Printf("修改用户信息 error: %v", err.Error())
		} else {
			if isSucc {
				msg.Code = 0
				msg.Msg = "修改成功！"
				//添加成功后移除token
				errr := RemoveSession(w, r, config.AddTokenKey)
				if errr != nil {
					log.Printf("移除Session error: %v", errr.Error())
				}
			} else {
				msg.Code = 1
				msg.Msg = "修改失败！"
				log.Printf("添加用户信息 error: %v", err.Error())
			}
		}
	} else {
		msg.Code = 3
		msg.Msg = "该信息已提交，请勿重复提交!"
	}
	util.ResponseHtml(w, msg, err)
}

//删除用户信息
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	msg := new(domain.Result)
	msg.Init()
	err := r.ParseForm()
	if err != nil {
		log.Printf("DeleteUser error: %v", err.Error())
	}
	userId := r.FormValue("id")
	num, errs := strconv.ParseInt(userId, 0, 0)
	if errs != nil {
		log.Printf("userId转换整型的错误 error: %v", errs.Error())
	}
	succ, err := service.DeleteByPrimarykeyUserInfo(num)
	if err != nil {
		msg.Code = 1
		msg.Msg = "删除用户信息失败！"
		log.Printf("删除用户信息 error: %v", err.Error())
	} else {
		if succ {
			msg.Code = 0
			msg.Msg = "删除用户信息成功！"
		} else {
			msg.Code = 1
			msg.Msg = "删除用户信息失败！"
		}
	}
	util.ResponseHtml(w, msg, err)
}
