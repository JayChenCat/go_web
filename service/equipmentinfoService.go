package service

import (
	"go_web/bll"
	"go_web/domain"
	"go_web/viewmodel"
)

//添加单个设备信息
func InsertEquipmentinfo(equipmentinfo *domain.Equipmentinfo) (bool, error) {
	err := bll.InsertEquipmentinfo(equipmentinfo)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

//修改单个设备信息
func UpdateEquipmentinfo(equipmentinfo *domain.Equipmentinfo) (bool, error) {
	err := bll.UpdateEquipmentinfo(equipmentinfo)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

//删除单个设备信息
func DeleteEquipmentinfo(userId int64) (bool, error) {
	err := bll.DeleteEquipmentinfo(userId)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

//根据编号获取单个设备信息
func GetSingleEquipmentinfo(userId int64) (*viewmodel.EquipmentinfoViewModel, error) {
	return bll.GetEquipmentinfoByUserId(userId)
}

// 获取设备信息列表(分页)-Account
func GetEquipmentinfoPageList(pageNumber int, pageSize int) ([]*viewmodel.EquipmentinfoViewModel, error) {
	return bll.GetEquipmentinfoList(pageNumber, pageSize)
}

//统计设备列表信息数量
func GetEquipmentinfoTotal() (int, error) {
	return bll.GetEquipmentinfoTotal()
}
