package domain

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//类似构造函数，初始化
func (msg *Result) Init() {
	msg.Code = 2
	msg.Msg = "mysql数据库连接错误，请检查数据库连接字符及数据库服务是否启动或其他未知错误，请联系系统管理员!"
}
