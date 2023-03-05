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
func InsertAccountmanager(accountmanager *domain.Accountmanager) (bool, error) {
	err := bll.InsertAccountmanager(accountmanager)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

//修改单个信息
func UpdateAccountmanager(accountmanager *domain.Accountmanager) (bool, error) {
	err := bll.UpdateAccountmanager(accountmanager)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

//删除单个信息
func DeleteAccountmanager(userId int64) (bool, error) {
	err := bll.DeleteAccountmanager(userId)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

//根据编号获取单个信息
func GetSingleAccountmanager(userId int64) (*viewmodel.AccountmanagerViewModel, error) {
	return bll.GetAccountmanagerByUserId(userId)
}

// 获取信息列表(分页)-Accountmanager
func GetAccountmanagerPageList(pageNumber int, pageSize int) ([]*viewmodel.AccountmanagerViewModel, error) {
	return bll.GetAccountmanagerList(pageNumber, pageSize)
}

//统计列表信息数量
func GetAccountmanagerTotal() (int, error) {
	return bll.GetAccountmanagerTotal()
}

//批量插入sql或批量修改sql
func BatchInsertAccountmanager(edit int, accountmanages []*viewmodel.AccountmanagerViewModel) (bool, error) {
	err := errors.New("") //批量插入sql或批量修改错误
	var result []*domain.Accountmanager
	for _, item := range accountmanages {
		num, _ := strconv.ParseInt(item.Id, 0, 0)
		_time := time.Now().Format("2006-01-02 15:04:05")
		commonFileld := domain.BaseFileld{
			CREATED_BY:   "admin", //创建人
			CREATED_TIME: _time,   //创建时间
			UPDATED_BY:   "admin", //更新人
			UPDATED_TIME: _time,   //更新时间
			IsDeleted:    0,       //是否删除(逻辑删除)，1是，0否
			Remarks:      "",      //备注
		}
		accountmanager := &domain.Accountmanager{
			Id:           num,              //编号(主键)
			AccountType:  item.AccountType, //账户类型(1表示主机FTP账号，2表示OCS FTP账号，4表示RTSP账号)
			User:         item.User,        //用户名
			Pwd:          item.Pwd,         //密码
			CommonFileld: commonFileld,
		}
		result = append(result, accountmanager)
	}
	if edit == 0 {
		err = bll.BatchInsertAccountmanager(result)
	} /*else {
		err = bll.BatchUpdateAccountmanager(result)
	}*/
	if err != nil {
		return false, err
	} else {
		err = nil
	}
	return true, err
}
