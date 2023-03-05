package bll

import (
	"database/sql"
	"fmt"
	"go_web/domain"
	"go_web/util"
	"go_web/viewmodel"
	"log"
	"strconv"
	"strings"
)

const (
	//插入信息语句
	insertAccountmanagerSQL = `INSERT INTO accountmanager(Id, AccountType, User, Pwd,CREATED_BY, CREATED_TIME, UPDATED_BY, UPDATED_TIME, IsDeleted, Remarks) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?,?,?)`
	//修改信息语句
	updateAccountmanagerSQL = `update accountmanager set AccountType=?, User=?, Pwd=?,CREATED_BY=?, CREATED_TIME=?, 
UPDATED_BY=?, UPDATED_TIME=?, IsDeleted=?, Remarks=? where Id=?`
	//删除信息语句
	deleteAccountmanagerSQL = `delete from accountmanager where Id=?`
	//信息列表语句(分页)
	listAccountmanagerSQL = `select id,AccountType,User,Pwd,CREATED_BY,CREATED_TIME,UPDATED_BY,
UPDATED_TIME,isDeleted,Remarks from accountmanager   LIMIT ? OFFSET ?`
	//根据编号查询单个信息
	SingleAccountmanagerSQL = `select id,AccountType,User,Pwd,CREATED_BY, CREATED_TIME, UPDATED_BY,
       UPDATED_TIME,isDeleted,Remarks from accountmanager where Id=?`
	//统计列表信息数量
	TotalAccountmanagerSQL = `select count(1) total from accountmanager`
)

func scanAccountmanagerWithSignOnAndBannerData(r *sql.Rows) (*domain.Accountmanager, error) {
	var id int64
	var user, pwd, CREATED_BY, CREATED_TIME, UPDATED_BY, UPDATED_TIME, Remarks string
	var accountType, isDeleted int

	err := r.Scan(&id, &accountType, &user, &pwd, &CREATED_BY, &CREATED_TIME,
		&UPDATED_BY, &UPDATED_TIME, &isDeleted, &Remarks)

	baseFileld := domain.BaseFileld{
		CREATED_BY:   CREATED_BY,   //创建人
		CREATED_TIME: CREATED_TIME, //创建时间
		UPDATED_BY:   UPDATED_BY,   //更新人
		UPDATED_TIME: UPDATED_TIME, //更新时间
		IsDeleted:    isDeleted,    //是否删除(逻辑删除)，1是，0否
		Remarks:      Remarks,      //备注
	}

	a := &domain.Accountmanager{
		Id:           id,          //编号(主键)
		AccountType:  accountType, //账户类型(1表示主机FTP账号，2表示OCS FTP账号，4表示RTSP账号)
		User:         user,        //用户名
		Pwd:          pwd,         //密码
		CommonFileld: baseFileld,  //公共字段(创建时间，创建人等属性)
	}
	return a, err
}

//根据编号获取单个信息
func GetAccountmanagerByUserId(userId int64) (*viewmodel.AccountmanagerViewModel, error) {
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	if err != nil {
		return nil, err
	}
	data, err := conn.Query(SingleAccountmanagerSQL, userId)
	if err != nil {
		return nil, err
	}
	accountmanagerViewModel := &viewmodel.AccountmanagerViewModel{}
	if data.Next() {
		accountmanager, err := scanAccountmanagerWithSignOnAndBannerData(data)
		if err != nil {
			return nil, err
		}
		accountmanagerViewModel = &viewmodel.AccountmanagerViewModel{
			Id:          strconv.FormatInt(accountmanager.Id, 10), //编号(主键)
			AccountType: accountmanager.AccountType,               //账户类型(1表示主机FTP账号，2表示OCS FTP账号，4表示RTSP账号)
			User:        accountmanager.User,                      //用户名
			Pwd:         accountmanager.Pwd,                       //密码
		}
		return accountmanagerViewModel, nil
	}
	defer data.Close()
	err = data.Err()
	if err != nil {
		return nil, err
	}
	return accountmanagerViewModel, err
}

// 获取信息列表(分页)-Account
func GetAccountmanagerList(pageNumber int, pageSize int) ([]*viewmodel.AccountmanagerViewModel, error) {
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	var result []*viewmodel.AccountmanagerViewModel
	if err != nil {
		return result, err
	}
	offset := (pageNumber - 1) * pageSize
	r, err := conn.Query(listAccountmanagerSQL, pageSize, offset)
	if err != nil {
		return result, err
	}
	for r.Next() {
		accountmanager, err := scanAccountmanagerWithSignOnAndBannerData(r)
		if err != nil {
			log.Printf("error: %v", err.Error())
			continue
		}
		//createtime := equipmentinfo.CommonFileld.CREATED_TIME
		//_time, _ := time.Parse(time.RFC3339, createtime)
		accountmanagerViewModel := &viewmodel.AccountmanagerViewModel{
			Id:          strconv.FormatInt(accountmanager.Id, 10), //编号(主键)
			AccountType: accountmanager.AccountType,               //账户类型(1表示主机FTP账号，2表示OCS FTP账号，4表示RTSP账号)
			User:        accountmanager.User,                      //用户名
			Pwd:         accountmanager.Pwd,                       //密码
		}
		result = append(result, accountmanagerViewModel)
	}
	defer r.Close()
	err = r.Err()
	if err != nil {
		return result, err
	}
	return result, nil
}

// insert Accountmanager
func InsertAccountmanager(accountmanager *domain.Accountmanager) error {
	err := InsertOrUpdateAccountmanager(insertAccountmanagerSQL, 0, accountmanager)
	return err
}

// update Accountmanager
func UpdateAccountmanager(accountmanager *domain.Accountmanager) error {
	err := InsertOrUpdateAccountmanager(updateAccountmanagerSQL, 1, accountmanager)
	return err
}

// Delete Accountmanager
func DeleteAccountmanager(userId int64) error {
	err := util.Delete(deleteAccountmanagerSQL, userId)
	return err
}

//统计列表信息数量
func GetAccountmanagerTotal() (int, error) {
	var total int
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	if err != nil {
		return 0, err
	}
	r, err := conn.Query(TotalAccountmanagerSQL)
	if err != nil {
		return 0, err
	}
	if r.Next() {
		err := r.Scan(&total)
		if err != nil {
			return 0, err
		}
	}
	defer r.Close()
	err = r.Err()
	if err != nil {
		return 0, err
	}
	return total, nil
}

//更新和插入语句的逻辑和返回值一致，可封装在一起
func InsertOrUpdateAccountmanager(SQL string, edit int, accountmanager *domain.Accountmanager) error {
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	if err != nil {
		return err
	}
	var r sql.Result

	//INSERT INTO user(Id, UserName, PassWord, NickName, Avatar, FullName, Mobile, Sex, Nation, Address, Email, IsLock, QRCode, RoleId, DepartId,
	//PositionId, CREATED_BY, CREATED_TIME, UPDATED_BY, UPDATED_TIME, IsDeleted, Remarks, LoginLastTime)

	//插入
	if edit == 0 {
		r, err = conn.Exec(SQL,
			accountmanager.Id,
			accountmanager.AccountType,
			accountmanager.User,
			accountmanager.Pwd,
			accountmanager.CommonFileld.CREATED_BY,
			accountmanager.CommonFileld.CREATED_TIME,
			accountmanager.CommonFileld.UPDATED_BY,
			accountmanager.CommonFileld.UPDATED_TIME,
			accountmanager.CommonFileld.IsDeleted,
			accountmanager.CommonFileld.Remarks,
		)
	} else {
		r, err = conn.Exec(SQL,
			accountmanager.AccountType,
			accountmanager.User,
			accountmanager.Pwd,
			accountmanager.CommonFileld.CREATED_BY,
			accountmanager.CommonFileld.CREATED_TIME,
			accountmanager.CommonFileld.UPDATED_BY,
			accountmanager.CommonFileld.UPDATED_TIME,
			accountmanager.CommonFileld.IsDeleted,
			accountmanager.CommonFileld.Remarks,
			accountmanager.Id)
	}
	if err != nil {
		return err
	}
	rowNum, err := r.RowsAffected()
	if err != nil {
		return err
	}
	if rowNum > 0 {
		return nil
	}
	return err
}

//批量插入sql
func BatchInsertAccountmanager(accountmanages []*domain.Accountmanager) error {
	count := len(accountmanages)
	// 存放 (?, ?) 的slice
	vStrings := make([]string, 0, count)
	// 存放values的slice
	vArgs := make([]interface{}, 0, count*2)
	//遍历
	for _, a := range accountmanages {
		// 此处占位符要与插入值的个数对应
		vStrings = append(vStrings, "(?, ?, ?, ?, ?, ?, ?, ?,?,?)")
		vArgs = append(vArgs, a.Id)
		vArgs = append(vArgs, a.AccountType)
		vArgs = append(vArgs, a.User)
		vArgs = append(vArgs, a.Pwd)
		vArgs = append(vArgs, a.CommonFileld.CREATED_BY)
		vArgs = append(vArgs, a.CommonFileld.CREATED_TIME)
		vArgs = append(vArgs, a.CommonFileld.UPDATED_BY)
		vArgs = append(vArgs, a.CommonFileld.UPDATED_TIME)
		vArgs = append(vArgs, a.CommonFileld.IsDeleted)
		vArgs = append(vArgs, a.CommonFileld.Remarks)
	}
	sqlHead := `INSERT INTO accountmanager(Id,AccountType,User,
		Pwd,CREATED_BY,CREATED_TIME,UPDATED_BY,UPDATED_TIME,IsDeleted,Remarks) VALUES %s`
	stmt := fmt.Sprintf(sqlHead, strings.Join(vStrings, ","))
	return util.BatchInsertOrUpdate(stmt, vArgs)
}

//批量修改sql
func BatchUpdateAccountmanager(accountmanages []*domain.Accountmanager) error {
	count := len(accountmanages)
	// 存放 (?, ?) 的slice
	vStrings := make([]string, 0, count)
	// 存放values的slice
	vArgs := make([]interface{}, 0, count*2)
	//遍历
	for _, a := range accountmanages {
		// 此处占位符要与插入值的个数对应
		vStrings = append(vStrings, "update accountmanager set AccountType=?,User=?,Pwd=?,CREATED_BY=?,CREATED_TIME=?,"+
			"UPDATED_BY=?,UPDATED_TIME=?,IsDeleted=?,Remarks=? where Id=?;")
		vArgs = append(vArgs, a.AccountType)
		vArgs = append(vArgs, a.User)
		vArgs = append(vArgs, a.Pwd)
		vArgs = append(vArgs, a.CommonFileld.CREATED_BY)
		vArgs = append(vArgs, a.CommonFileld.CREATED_TIME)
		vArgs = append(vArgs, a.CommonFileld.UPDATED_BY)
		vArgs = append(vArgs, a.CommonFileld.UPDATED_TIME)
		vArgs = append(vArgs, a.CommonFileld.IsDeleted)
		vArgs = append(vArgs, a.CommonFileld.Remarks)
		vArgs = append(vArgs, a.Id)
	}
	stmt := fmt.Sprintf(strings.Join(vStrings, ""))
	return util.BatchInsertOrUpdate(stmt, vArgs)
}
