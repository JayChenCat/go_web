package bll

import (
	"database/sql"
	"go_web/domain"
	"go_web/util"
	"go_web/viewmodel"
	"log"
	"strconv"
	"time"
)

const (
	//登录sql语句
	getAccountByUsernameSQL = `select u.Id id,roleId,roleName,departId,positionId,userName,PassWord,email,nickName,avatar,fullName
,mobile,nation,address,qRCode,loginLastTime,u.CREATED_BY CREATED_BY,u.CREATED_TIME CREATED_TIME,u.UPDATED_BY UPDATED_BY,
u.UPDATED_TIME UPDATED_TIME,u.Remarks Remarks,sex,isLock,u.isDeleted  isDeleted
from user u left join role r on u.RoleId=r.Id where UserName=? and PassWord=?`
	//插入用户信息语句
	insertUserSQL = `INSERT INTO user(Id, UserName, PassWord, NickName, Avatar, FullName, Mobile, Sex, Nation, Address, Email, IsLock, QRCode, RoleId, DepartId, 
PositionId, CREATED_BY, CREATED_TIME, UPDATED_BY, UPDATED_TIME, IsDeleted, Remarks, LoginLastTime) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?,?, ?, ?, ?,?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	//修改用户信息语句
	updateUserSQL = `update user set UserName=?, PassWord=?, NickName=?, Avatar=?, FullName=?, Mobile=?, Sex=?, Nation=?, Address=?, Email=?, 
IsLock=?, QRCode=?, RoleId=?, DepartId=?, PositionId=?, CREATED_BY=?, CREATED_TIME=?, 
UPDATED_BY=?, UPDATED_TIME=?, IsDeleted=?, Remarks=?, LoginLastTime=? where Id=?`
	//删除用户信息语句
	deleteUserSQL      = `delete from user where Id=?`
	deleteFalseUserSQL = `update user set IsDeleted=1 where Id=?`
	//用户信息列表语句(分页)
	listUserSQL = `select u.Id id,roleId,roleName,departId,positionId,userName,PassWord,email,nickName,avatar,fullName
,mobile,nation,address,qRCode,loginLastTime,u.CREATED_BY CREATED_BY,u.CREATED_TIME CREATED_TIME,u.UPDATED_BY UPDATED_BY,
u.UPDATED_TIME UPDATED_TIME,u.Remarks Remarks,sex,isLock,u.isDeleted  isDeleted
from user u left join role r on u.RoleId=r.Id  where u.isDeleted=0 LIMIT ? OFFSET ?`
	//根据编号查询单个用户信息
	SingleUserSQL = `select u.Id id,roleId,roleName,departId,positionId,userName,PassWord,email,nickName,avatar,fullName
,mobile,nation,address,qRCode,loginLastTime,u.CREATED_BY CREATED_BY,u.CREATED_TIME CREATED_TIME,u.UPDATED_BY UPDATED_BY,
u.UPDATED_TIME UPDATED_TIME,u.Remarks Remarks,sex,isLock,u.isDeleted  isDeleted
from user u left join role r on u.RoleId=r.Id where u.Id=?`
	//统计用户列表信息数量
	TotalUserSQL = `select count(1) total from user`
)

func scanAccountWithSignOnAndBannerData(r *sql.Rows) (*domain.Account, error) {
	var id, roleId, departId, positionId int64
	var userName, passWord, email, nickName, avatar, fullName, mobile, nation, address, qRCode, loginLastTime, CREATED_BY, CREATED_TIME, Remarks string
	var UPDATED_BY, UPDATED_TIME string
	var sex, isLock, isDeleted int
	var roleName *string

	err := r.Scan(&id, &roleId, &roleName, &departId, &positionId, &userName, &passWord, &email, &nickName, &avatar, &fullName, &mobile,
		&nation, &address, &qRCode, &loginLastTime, &CREATED_BY, &CREATED_TIME,
		&UPDATED_BY, &UPDATED_TIME, &Remarks, &sex, &isLock, &isDeleted)

	baseFileld := domain.BaseFileld{
		CREATED_BY:   CREATED_BY,   //创建人
		CREATED_TIME: CREATED_TIME, //创建时间
		UPDATED_BY:   UPDATED_BY,   //更新人
		UPDATED_TIME: UPDATED_TIME, //更新时间
		IsDeleted:    isDeleted,    //是否删除(逻辑删除)，1是，0否
		Remarks:      Remarks,      //备注
	}

	a := &domain.Account{
		Id:            id,            //用户编号(主键)
		UserName:      userName,      //用户名称
		PassWord:      passWord,      //密码
		NickName:      nickName,      //昵称
		Avatar:        avatar,        //用户头像地址
		FullName:      fullName,      //真实姓名(备用字段)
		Mobile:        mobile,        //手机号
		Sex:           sex,           //性别(0 表示女 1表示男)
		Nation:        nation,        //民族
		Address:       address,       //用户地址
		Email:         email,         //电子邮件
		IsLock:        isLock,        //是否禁用 1禁用，0为不禁用
		QRCode:        qRCode,        //二维码号(备用字段)
		RoleId:        roleId,        //所属角色(外键)
		DepartId:      departId,      //所属部门(外键)
		PositionId:    positionId,    //所属职位(备用字段)
		CommonFileld:  baseFileld,    //公共字段(创建时间，创建人等属性)
		LoginLastTime: loginLastTime, //用户最后登录时间
		RoelName:      roleName,      //角色名称
	}
	return a, err
}

//根据用户编号获取单个信息
func GetAccountByUserId(userId int64) (*viewmodel.AccountViewModel, error) {
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	if err != nil {
		return nil, err
	} else {
		data, errs := conn.Query(SingleUserSQL, userId)
		defer data.Close()
		err = data.Err()
		if err != nil {
			return nil, err
		}
		if errs != nil {
			return nil, errs
		} else {
			if data.Next() {
				account, erro := scanAccountWithSignOnAndBannerData(data)
				if erro != nil {
					return nil, erro
				} else {
					if account != nil {
						createtime := account.CommonFileld.CREATED_TIME
						_time, _ := time.Parse(time.RFC3339, createtime)
						result := &viewmodel.AccountViewModel{
							Id:           strconv.FormatInt(account.Id, 10),   //用户编号(主键)
							UserName:     account.UserName,                    //用户名称
							Password:     account.PassWord,                    //密码
							NickName:     account.NickName,                    //昵称
							Avatar:       account.Avatar,                      //用户头像地址
							FullName:     account.FullName,                    //真实姓名(备用字段)
							Mobile:       account.Mobile,                      //手机号
							Sex:          account.Sex,                         //性别(0 表示女 1表示男)
							Nation:       account.Nation,                      //民族
							Address:      account.Address,                     //用户地址
							Email:        account.Email,                       //电子邮件
							CREATED_TIME: _time.Format("2006-01-02 15:04:05"), //创建时间
						}
						return result, nil
					}
				}
			}
		}
	}
	return nil, err
}

// 获取用户信息列表(分页)-Account
func GetAccountList(pageNumber int, pageSize int) ([]*viewmodel.AccountViewModel, error) {
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	var result []*viewmodel.AccountViewModel
	if err != nil {
		return result, err
	}
	offset := (pageNumber - 1) * pageSize
	r, err := conn.Query(listUserSQL, pageSize, offset)
	if err != nil {
		return result, err
	}
	for r.Next() {
		account, err := scanAccountWithSignOnAndBannerData(r)
		if err != nil {
			log.Printf("error: %v", err.Error())
			continue
		} else {
			createtime := account.CommonFileld.CREATED_TIME
			_time, _ := time.Parse(time.RFC3339, createtime)
			accountViewModel := &viewmodel.AccountViewModel{
				Id:           strconv.FormatInt(account.Id, 10),   //用户编号(主键)
				UserName:     account.UserName,                    //用户名称
				Password:     account.PassWord,                    //密码
				NickName:     account.NickName,                    //昵称
				Avatar:       account.Avatar,                      //用户头像地址
				FullName:     account.FullName,                    //真实姓名(备用字段)
				Mobile:       account.Mobile,                      //手机号
				Sex:          account.Sex,                         //性别(0 表示女 1表示男)
				Nation:       account.Nation,                      //民族
				Address:      account.Address,                     //用户地址
				Email:        account.Email,                       //电子邮件
				CREATED_TIME: _time.Format("2006-01-02 15:04:05"), //创建时间
			}
			result = append(result, accountViewModel)
		}
	}
	defer r.Close()
	err = r.Err()
	if err != nil {
		return result, err
	}
	return result, nil
}

//登录
func GetAccountByUserName(userName string, passWord string) (*viewmodel.AccountViewModel, error) {
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	if err != nil {
		return nil, err
	}
	data, err := conn.Query(getAccountByUsernameSQL, userName, passWord)
	if err != nil {
		return nil, err
	}
	accountViewModel := &viewmodel.AccountViewModel{}
	if data.Next() {
		account, err := scanAccountWithSignOnAndBannerData(data)
		if err != nil {
			return nil, err
		}
		createtime := account.CommonFileld.CREATED_TIME
		_time, _ := time.Parse(time.RFC3339, createtime)
		accountViewModel = &viewmodel.AccountViewModel{
			Id:           strconv.FormatInt(account.Id, 10),   //用户编号(主键)
			UserName:     account.UserName,                    //用户名称
			Password:     account.PassWord,                    //密码
			NickName:     account.NickName,                    //昵称
			Avatar:       account.Avatar,                      //用户头像地址
			FullName:     account.FullName,                    //真实姓名(备用字段)
			Mobile:       account.Mobile,                      //手机号
			Sex:          account.Sex,                         //性别(0 表示女 1表示男)
			Nation:       account.Nation,                      //民族
			Address:      account.Address,                     //用户地址
			Email:        account.Email,                       //电子邮件
			CREATED_TIME: _time.Format("2006-01-02 15:04:05"), //创建时间
		}
		return accountViewModel, nil
	}
	defer data.Close()
	err = data.Err()
	if err != nil {
		return nil, err
	}
	return accountViewModel, err
}

// insert account
func InsertAccount(account *domain.Account) error {
	err := InsertOrUpdateAccount(insertUserSQL, 0, account)
	return err
}

// update account
func UpdateAccount(account *domain.Account) error {
	err := InsertOrUpdateAccount(updateUserSQL, 1, account)
	return err
}

// Delete account
func DeleteAccount(userId int64) error {
	err := util.Delete(deleteFalseUserSQL, userId)
	return err
}

//统计用户列表信息数量
func GetAccountTotal() (int, error) {
	var total int
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	if err != nil {
		return 0, err
	}
	r, err := conn.Query(TotalUserSQL)
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
func InsertOrUpdateAccount(SQL string, edit int, account *domain.Account) error {
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
			account.Id,         //用户编号
			account.UserName,   //用户名
			account.PassWord,   //密码
			account.NickName,   //昵称
			account.Avatar,     //用户头像地址
			account.FullName,   //真实姓名(备用字段)
			account.Mobile,     //手机号
			account.Sex,        //性别(0 表示女 1表示男)
			account.Nation,     //民族
			account.Address,    //用户地址
			account.Email,      //电子邮件
			account.IsLock,     //是否禁用 1禁用，0为不禁用
			account.QRCode,     //二维码号(备用字段)
			account.RoleId,     //所属角色(外键)
			account.DepartId,   //所属部门(外键)
			account.PositionId, //所属职位(备用字段)
			account.CommonFileld.CREATED_BY,
			account.CommonFileld.CREATED_TIME,
			account.CommonFileld.UPDATED_BY,
			account.CommonFileld.UPDATED_TIME,
			account.CommonFileld.IsDeleted,
			account.CommonFileld.Remarks,
			account.LoginLastTime,
		)
	} else {
		r, err = conn.Exec(SQL,
			account.UserName,   //用户名
			account.PassWord,   //密码
			account.NickName,   //昵称
			account.Avatar,     //用户头像地址
			account.FullName,   //真实姓名(备用字段)
			account.Mobile,     //手机号
			account.Sex,        //性别(0 表示女 1表示男)
			account.Nation,     //民族
			account.Address,    //用户地址
			account.Email,      //电子邮件
			account.IsLock,     //是否禁用 1禁用，0为不禁用
			account.QRCode,     //二维码号(备用字段)
			account.RoleId,     //所属角色(外键)
			account.DepartId,   //所属部门(外键)
			account.PositionId, //所属职位(备用字段)
			account.CommonFileld.CREATED_BY,
			account.CommonFileld.CREATED_TIME,
			account.CommonFileld.UPDATED_BY,
			account.CommonFileld.UPDATED_TIME,
			account.CommonFileld.IsDeleted,
			account.CommonFileld.Remarks,
			account.LoginLastTime,
			account.Id)
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

/*gorm框架-增删改查*/

//创建表记录-gorm框架
func Add(user *domain.Account) (bool, error) {
	db := util.Init()
	result := db.Table("user").Create(user)
	util.DataBaseClose()
	errs := result.Error //返回 error
	if errs != nil {
		log.Printf("创建表记录错误 error: %v", errs)
		return false, errs
	}
	rows := result.RowsAffected // 返回影响记录的条数
	if rows > 0 {
		return true, nil
	}
	return true, nil
}

//根据主键查询单个记录-gorm框架
func QueryByByPrimarykey(fields []string, id int64) (*viewmodel.AccountViewModel, error) {
	db := util.Init()
	/*fileds := []string{
		"Id",
		"UserName",
		"PassWord",
		"Email",
		"Sex",
		"NickName",
		"CREATED_TIME",
	}*/
	//util.QueryByByPrimarykey("user", fileds, "Id=?", 298957598143545343)
	var user = &domain.Account{}
	//db.Where(map[string]interface{}{"Name": "jinzhu", "Age": 0}).Find(&users)
	query := db.Table("user").Select(fields).Where("Id=?", id).First(user) //
	errs := query.Error                                                    //返回 error
	if errs != nil {
		log.Printf("查询单个记录错误 error: %v", errs)
		//return nil
	}
	util.DataBaseClose()
	if user != nil {
		accountViewModel := &viewmodel.AccountViewModel{
			Id:       strconv.FormatInt(user.Id, 10), //用户编号(主键)
			UserName: user.UserName,                  //用户名称
			Password: user.PassWord,                  //密码
			Email:    user.Email,                     //电子邮件
		}
		return accountViewModel, nil
	}
	return nil, nil
}

//删除记录(参数必须为结构体指针)-写法一
func DeleteData(key string, args int64, user *domain.Account) (bool, error) {
	db := util.Init()
	result := db.Table("user").Where(key, args).Delete(user) //"id =?"
	util.DataBaseClose()
	errs := result.Error //返回 error
	if errs != nil {
		log.Printf("删除表记录错误 error: %v", errs)
		return false, errs
	}
	rows := result.RowsAffected // 返回影响记录的条数
	if rows > 0 {
		return true, nil
	}
	return true, nil
}

//批量删除-写法二(根据多个主键删除)
//批量更新
//IN关键字
//批量更新的话，只能批量对行的某些字段改为相同的值，不能改为不同的值…感觉没啥用
//db.Table("users").Where("id IN (?)", []int{10, 11}).Updates(map[string]interface{}{"name": "hello", "age": 18})

func BatchDeleteDataByPrimarykey(key string, args []int64, user *domain.Account) (bool, error) {
	db := util.Init()
	result := db.Table("user").Where(key, args).Delete(user)
	util.DataBaseClose()
	errs := result.Error //返回 error
	if errs != nil {
		log.Printf("删除表记录错误 error: %v", errs)
		return false, errs
	}
	rows := result.RowsAffected // 返回影响记录的条数
	if rows > 0 {
		return true, nil
	}
	return true, nil
}

func DeleteById(ids []int64) (bool, error) {
	var user = &domain.Account{}
	count := len(ids)
	if count == 1 {
		//单个删除
		return DeleteData("Id=?", ids[0], user)
	}
	if count > 1 {
		//批量删除
		return BatchDeleteDataByPrimarykey("Id in(?)", ids, user)
	}
	return false, nil
}
