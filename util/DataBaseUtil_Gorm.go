package util

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

//初始化数据库的使用
var db *gorm.DB

func Init() *gorm.DB {
	dataSourceName := userName + ":" + password + tcpPort + dbName + "?" + charset + "&" + local + "&" + parseTime
	conn, err := gorm.Open(databaseType, dataSourceName) //默认使用mysql数据库
	if err != nil {
		log.Printf("初始化数据库错误 error: %v", err.Error())
	}
	db = conn
	//gorm会在创建表的时候去掉”s“的后缀
	db.SingularTable(true)
	//自动迁移
	//db.AutoMigrate(&UserInfo{})
	return conn
}

func DataBaseClose() {
	defer db.Close()
}

//分页封装 var list []user
//db.Scopes(Paginate(1,10)).Find(&list)
//参数即可 Scopes 使你可以复用通用的逻辑，共享的逻辑需要定义为 func(*gorm.DB) *gorm.DB 类型
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
