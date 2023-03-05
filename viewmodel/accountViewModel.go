package viewmodel

// 用户信息临时对象
type AccountViewModel struct {
	Id           string `json:"id"`       //用户编号(主键)
	UserName     string `json:"username"` //用户名称
	Password     string `json:"password"` //密码
	NickName     string `json:"nickName"` //昵称
	Avatar       string `json:"code"`     //用户头像地址
	FullName     string `json:"avatar"`   //真实姓名(备用字段)
	Mobile       string `json:"mobile"`   //手机号
	Sex          int    `json:"sex"`      //性别(0 表示女 1表示男)
	Nation       string `json:"nation"`   //民族
	Address      string `json:"address"`  //用户地址
	Email        string `json:"email"`    //电子邮件
	CREATED_TIME string `json:"addtime"`  //创建时间
}
