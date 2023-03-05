package service

import (
	"errors"
	"go_web/bll"
	"go_web/domain"
	"go_web/viewmodel"
	"strconv"
	"time"
)

//添加单个信息
func InsertCommunicationsettings(communicationsettings *domain.Communicationsettings) (bool, error) {
	err := bll.InsertCommunicationsettings(communicationsettings)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

//修改单个信息
func UpdateCommunicationsettings(communicationsettings *domain.Communicationsettings) (bool, error) {
	err := bll.UpdateCommunicationsettings(communicationsettings)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

//删除单个信息
func DeleteCommunicationsettings(userId int64) (bool, error) {
	err := bll.DeleteCommunicationsettings(userId)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

//根据编号获取单个信息
func GetSingleCommunicationsettings(userId int64) (*viewmodel.CommunicationsettingsViewModel, error) {
	return bll.GetCommunicationsettingsByUserId(userId)
}

// 获取信息列表(分页)-Accountmanager
func GetCommunicationsettingsPageList(pageNumber int, pageSize int) ([]*viewmodel.CommunicationsettingsViewModel, error) {
	return bll.GetCommunicationsettingsList(pageNumber, pageSize)
}

//统计列表信息数量
func GetCommunicationsettingsTotal() (int, error) {
	return bll.GetCommunicationsettingsTotal()
}

//批量插入sql或批量修改sql
func BatchInsertCommunicationsettings(edit int, communicationsettings []*viewmodel.CommunicationsettingsViewModel) (bool, error) {
	err := errors.New("") //批量插入sql或批量修改错误
	var result []*domain.Communicationsettings
	for _, item := range communicationsettings {
		num, _ := strconv.ParseInt(item.Id, 0, 0)
		port, _ := strconv.ParseInt(item.Port, 0, 0)
		_time := time.Now().Format("2006-01-02 15:04:05")
		commonFileld := domain.BaseFileld{
			CREATED_BY:   "admin", //创建人
			CREATED_TIME: _time,   //创建时间
			UPDATED_BY:   "admin", //更新人
			UPDATED_TIME: _time,   //更新时间
			IsDeleted:    0,       //是否删除(逻辑删除)，1是，0否
			Remarks:      "",      //备注
		}
		accountmanager := &domain.Communicationsettings{
			Id:                num,                    //编号(主键)
			CommunicationType: item.CommunicationType, //通讯类型(1表示本机，2表示OSC)
			IP:                item.IP,                //IP地址
			Port:              port,                   //端口号
			CommonFileld:      commonFileld,
		}
		result = append(result, accountmanager)
	}
	if edit == 0 {
		err = bll.BatchInsertCommunicationsettings(result)
	} /*else {
		err = bll.BatchUpdateCommunicationsettings(result)
	}*/
	if err != nil {
		return false, err
	} else {
		err = nil
	}
	return true, err
}
