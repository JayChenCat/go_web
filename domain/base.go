package domain

//公共字段(创建时间，创建人等属性)

type BaseFileld struct {
	CREATED_BY   string `gorm:"column:CREATED_BY"json:"create_by"`     //创建人
	CREATED_TIME string `gorm:"column:CREATED_TIME"json:"addtime"`     //创建时间
	UPDATED_BY   string `gorm:"column:UPDATED_BY"json:"update_by"`     //更新人
	UPDATED_TIME string `gorm:"column:UPDATED_TIME"json:"update_time"` //更新时间
	IsDeleted    int    `gorm:"column:IsDeleted"json:"isdeleted"`      //是否删除(逻辑删除)，1是，0否
	Remarks      string `gorm:"column:Remarks"json:"remarks"`          //备注
}
