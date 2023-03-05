package util

//数据库使用的常量
const (
	userName     = "root"
	password     = "123456"
	dbName       = "crrc_dms_web"
	driverName   = "mysql"
	charset      = "charset=utf8"
	local        = "loc=Local"             //Local
	tcpPort      = "@tcp(127.0.0.1:3306)/" // @tcp(119.91.142.127:3306)/
	parseTime    = "parseTime=true"        // 用以解析 数据库 中的 date 类型，否则会解析成 []uint8 不能隐式转为 string
	databaseType = "mysql"                 //数据库类型
)
