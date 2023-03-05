package service

import (
	"go_web/bll"
	"go_web/domain"
	"go_web/viewmodel"
)

//添加单个设备信息
func InsertUpgraderecord(upgraderecord *domain.Upgraderecord) (bool, error) {
	err := bll.InsertUpgraderecord(upgraderecord)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

//修改单个信息
func UpdateUpgraderecord(upgraderecord *domain.Upgraderecord) (bool, error) {
	err := bll.UpdateUpgraderecord(upgraderecord)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

//删除单个信息
func DeleteUpgraderecord(userId int64) (bool, error) {
	err := bll.DeleteUpgraderecord(userId)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

//根据编号获取单个信息
func GetSingleUpgraderecord(userId int64) (*viewmodel.UpgraderecordViewModel, error) {
	return bll.GetUpgraderecordByUserId(userId)
}

// 获取信息列表(分页)-Account
func GetUpgraderecordPageList(pageNumber int, pageSize int) ([]*viewmodel.UpgraderecordViewModel, error) {
	return bll.GetUpgraderecordList(pageNumber, pageSize)
}

//统计列表信息数量
func GetUpgraderecordTotal() (int, error) {
	return bll.GetUpgraderecordTotal()
}
