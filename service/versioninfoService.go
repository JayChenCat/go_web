package service

import (
	"go_web/bll"
	"go_web/domain"
	"go_web/viewmodel"
)

//添加单个设备信息
func InsertVersioninfo(equipmentinfo *domain.Versioninfo) (bool, error) {
	err := bll.InsertVersioninfo(equipmentinfo)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

//修改单个设备信息
func UpdateVersioninfo(equipmentinfo *domain.Versioninfo) (bool, error) {
	err := bll.UpdateVersioninfo(equipmentinfo)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

//删除单个设备信息
func DeleteVersioninfo(userId int64) (bool, error) {
	err := bll.DeleteVersioninfo(userId)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

//根据编号获取单个设备信息
func GetSingleVersioninfo(userId int64) (*viewmodel.VersioninfoViewModel, error) {
	return bll.GetVersioninfoByUserId(userId)
}

// 获取设备信息列表(分页)-Account
func GetVersioninfoPageList(pageNumber int, pageSize int) ([]*viewmodel.VersioninfoViewModel, error) {
	return bll.GetVersioninfoList(pageNumber, pageSize)
}

//统计设备列表信息数量
func GetVersioninfoTotal() (int, error) {
	return bll.GetVersioninfoTotal()
}
