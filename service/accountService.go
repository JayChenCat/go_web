package service

import (
	"go_web/bll"
	"go_web/domain"
	"go_web/viewmodel"
)

// 用户登录
func Login(userName string, password string) (*viewmodel.AccountViewModel, error) {
	return bll.GetAccountByUserName(userName, password)
}

//添加单个用户信息
func InsertUser(account *domain.Account) (bool, error) {
	err := bll.InsertAccount(account)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

//修改单个用户信息
func UpdateUser(account *domain.Account) (bool, error) {
	err := bll.UpdateAccount(account)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

//删除单个用户信息
func DeleteUser(userId int64) (bool, error) {
	err := bll.DeleteAccount(userId)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

//根据用户编号获取单个信息
func GetSingleAccount(userId int64) (*viewmodel.AccountViewModel, error) {
	return bll.GetAccountByUserId(userId)
}

// 获取用户信息列表(分页)-Account
func GetPageList(pageNumber int, pageSize int) ([]*viewmodel.AccountViewModel, error) {
	return bll.GetAccountList(pageNumber, pageSize)
}

//统计用户列表信息数量
func GetAccountTotal() (int, error) {
	return bll.GetAccountTotal()
}

//获取单个用户信息
func GetByByPrimarykeyUserInfo(userId int64) (*viewmodel.AccountViewModel, error) {
	//需要显示的列
	fileds := []string{
		"Id",
		"UserName",
		"PassWord",
		"Email",
	}
	return bll.QueryByByPrimarykey(fileds, userId)
}

//单个或批量删除
func DeleteByPrimarykeyUserInfo(userId int64) (bool, error) {
	ids := []int64{userId}
	return bll.DeleteById(ids)
}
