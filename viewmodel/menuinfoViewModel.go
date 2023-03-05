package viewmodel

//导航菜单信息表
type Menuinfo struct {
	Id        string `json:"id"`        //编号(主键)
	MName     string `json:"mname"`     //菜单名称
	PageURL   string `json:"pageURL"`   //菜单页面路径
	TypeID    int    `json:"typeID"`    //菜单级别(0表示一级菜单,1表示二级菜单)
	ParentID  int64  `json:"parentID"`  //上级/父级菜单的ID
	Sort      int    `json:"sort"`      //排序
	Icon      string `json:"icon"`      //菜单的图标
	IsDisplay int    `json:"isDisplay"` //是否显示(禁用) 1启用,0禁用
}
