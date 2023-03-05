package domain

import (
	"encoding/gob"
)

// 用户信息对象
type Account struct {
	Id            int64      `gorm:"column:Id"json:"id"`                       //用户编号(主键)
	UserName      string     `gorm:"column:UserName"json:"username"`           //用户名称
	PassWord      string     `gorm:"column:PassWord"json:"password"`           //密码
	NickName      string     `gorm:"column:NickName"json:"nickName"`           //昵称
	Avatar        string     `gorm:"column:Avatar"json:"avatar"`               //用户头像地址
	FullName      string     `gorm:"column:FullName"json:"fullName"`           //真实姓名(备用字段)
	Mobile        string     `gorm:"column:Mobile"json:"mobile"`               //手机号
	Sex           int        `gorm:"column:Sex"json:"sex"`                     //性别(0 表示女 1表示男)
	Nation        string     `gorm:"column:Nation"json:"nation"`               //民族
	Address       string     `gorm:"column:Address"json:"address"`             //用户地址
	Email         string     `gorm:"column:Email"json:"email"`                 //电子邮件
	IsLock        int        `gorm:"column:IsLock"json:"isLock"`               //是否禁用 1禁用，0为不禁用
	QRCode        string     `gorm:"column:QRCode"json:"qrcode"`               //二维码号(备用字段)
	RoleId        int64      `gorm:"column:RoleId"json:"roleid"`               //所属角色(外键)
	DepartId      int64      `gorm:"column:DepartId"json:"departid"`           //所属部门(外键)
	PositionId    int64      `gorm:"column:PositionId"json:"positionid"`       //所属职位(备用字段)
	CommonFileld  BaseFileld `json:"commonfileld"`                             //公共字段(创建时间，创建人等属性)
	LoginLastTime string     `gorm:"column:LoginLastTime"json:"loginLastTime"` //用户最后登录时间
	RoelName      *string    `gorm:"column:RoelName"json:"roelName"`           //角色名称
}

func init() {
	gob.Register(&Account{})
}
