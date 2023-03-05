package bll

import (
	"database/sql"
	"go_web/domain"
	"go_web/util"
	"go_web/viewmodel"
	"log"
	"strconv"
)

const (
	//插入信息语句
	insertVersioninfoSQL = `INSERT INTO versioninfo(Id, ArmVersion, UpgradeVersion, WebVersion, WebServerName, CREATED_BY, CREATED_TIME, UPDATED_BY, UPDATED_TIME, IsDeleted, Remarks) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?,?, ?, ?)`
	//修改信息语句
	updateVersioninfoSQL = `update versioninfo set ArmVersion=?, UpgradeVersion=?, WebVersion=?, WebServerName=?, CREATED_BY=?, CREATED_TIME=?, 
UPDATED_BY=?, UPDATED_TIME=?, IsDeleted=?, Remarks=? where Id=?`
	//删除信息语句
	deleteVersioninfoSQL = `delete from versioninfo where Id=?`
	//信息列表语句(分页)
	listVersioninfoSQL = `select id,ArmVersion,UpgradeVersion,WebVersion,WebServerName,CREATED_BY, CREATED_TIME, UPDATED_BY,
UPDATED_TIME,isDeleted,Remarks from versioninfo   LIMIT ? OFFSET ?`
	//根据编号查询单个信息
	SingleVersioninfoSQL = `select id,ArmVersion,UpgradeVersion,WebVersion,WebServerName,CREATED_BY, CREATED_TIME, UPDATED_BY,
       UPDATED_TIME,isDeleted,Remarks from versioninfo where Id=?`
	//统计用户列表信息数量
	TotalVersioninfoSQL = `select count(1) total from versioninfo`
)

func scanVersioninfoWithSignOnAndBannerData(r *sql.Rows) (*domain.Versioninfo, error) {
	var id int64
	var armVersion, upgradeVersion, webVersion, webServerName, CREATED_BY, CREATED_TIME, UPDATED_BY, UPDATED_TIME, Remarks string
	var isDeleted int

	err := r.Scan(&id, &armVersion, &upgradeVersion, &webVersion, &webServerName, &CREATED_BY, &CREATED_TIME,
		&UPDATED_BY, &UPDATED_TIME, &isDeleted, &Remarks)

	baseFileld := domain.BaseFileld{
		CREATED_BY:   CREATED_BY,   //创建人
		CREATED_TIME: CREATED_TIME, //创建时间
		UPDATED_BY:   UPDATED_BY,   //更新人
		UPDATED_TIME: UPDATED_TIME, //更新时间
		IsDeleted:    isDeleted,    //是否删除(逻辑删除)，1是，0否
		Remarks:      Remarks,      //备注
	}

	a := &domain.Versioninfo{
		Id:             id,             //编号(主键)
		ArmVersion:     armVersion,     //设备应用程序版本
		UpgradeVersion: upgradeVersion, //升级程序版本
		WebVersion:     webVersion,     //Web应用程序版本
		WebServerName:  webServerName,  //Web 服务器名称
		CommonFileld:   baseFileld,     //公共字段(创建时间，创建人等属性)
	}
	return a, err
}

//根据编号获取单个信息
func GetVersioninfoByUserId(userId int64) (*viewmodel.VersioninfoViewModel, error) {
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	if err != nil {
		return nil, err
	}
	data, err := conn.Query(SingleVersioninfoSQL, userId)
	if err != nil {
		return nil, err
	}
	versioninfoViewModel := &viewmodel.VersioninfoViewModel{}
	if data.Next() {
		versioninfo, err := scanVersioninfoWithSignOnAndBannerData(data)
		if err != nil {
			return nil, err
		}
		versioninfoViewModel = &viewmodel.VersioninfoViewModel{
			Id:             strconv.FormatInt(versioninfo.Id, 10), //编号(主键)
			ArmVersion:     versioninfo.ArmVersion,                //设备应用程序版本
			UpgradeVersion: versioninfo.UpgradeVersion,            //升级程序版本
			WebVersion:     versioninfo.WebVersion,                //Web应用程序版本
			WebServerName:  versioninfo.WebServerName,             //Web 服务器名称
		}
		return versioninfoViewModel, nil
	}
	defer data.Close()
	err = data.Err()
	if err != nil {
		return nil, err
	}
	return versioninfoViewModel, err
}

// 获取信息列表(分页)-Account
func GetVersioninfoList(pageNumber int, pageSize int) ([]*viewmodel.VersioninfoViewModel, error) {
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	var result []*viewmodel.VersioninfoViewModel
	if err != nil {
		return result, err
	}
	offset := (pageNumber - 1) * pageSize
	r, err := conn.Query(listVersioninfoSQL, pageSize, offset)
	if err != nil {
		return result, err
	}
	for r.Next() {
		versioninfo, err := scanVersioninfoWithSignOnAndBannerData(r)
		if err != nil {
			log.Printf("error: %v", err.Error())
			continue
		}
		//createtime := equipmentinfo.CommonFileld.CREATED_TIME
		//_time, _ := time.Parse(time.RFC3339, createtime)
		versioninfoViewModel := &viewmodel.VersioninfoViewModel{
			Id:             strconv.FormatInt(versioninfo.Id, 10), //用户编号(主键)
			ArmVersion:     versioninfo.ArmVersion,                //设备应用程序版本
			UpgradeVersion: versioninfo.UpgradeVersion,            //升级程序版本
			WebVersion:     versioninfo.WebVersion,                //Web应用程序版本
			WebServerName:  versioninfo.WebServerName,             //Web 服务器名称
		}
		result = append(result, versioninfoViewModel)
	}
	defer r.Close()
	err = r.Err()
	if err != nil {
		return result, err
	}
	return result, nil
}

// insert Versioninfo
func InsertVersioninfo(versioninfo *domain.Versioninfo) error {
	err := InsertOrUpdateVersioninfo(insertVersioninfoSQL, 0, versioninfo)
	return err
}

// update Versioninfo
func UpdateVersioninfo(versioninfo *domain.Versioninfo) error {
	err := InsertOrUpdateVersioninfo(updateVersioninfoSQL, 1, versioninfo)
	return err
}

// Delete Versioninfo
func DeleteVersioninfo(userId int64) error {
	err := util.Delete(deleteVersioninfoSQL, userId)
	return err
}

//统计列表信息数量
func GetVersioninfoTotal() (int, error) {
	var total int
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	if err != nil {
		return 0, err
	}
	r, err := conn.Query(TotalVersioninfoSQL)
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
func InsertOrUpdateVersioninfo(SQL string, edit int, versioninfo *domain.Versioninfo) error {
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
			versioninfo.Id,
			versioninfo.ArmVersion,
			versioninfo.UpgradeVersion,
			versioninfo.WebVersion,
			versioninfo.WebServerName,
			versioninfo.CommonFileld.CREATED_BY,
			versioninfo.CommonFileld.CREATED_TIME,
			versioninfo.CommonFileld.UPDATED_BY,
			versioninfo.CommonFileld.UPDATED_TIME,
			versioninfo.CommonFileld.IsDeleted,
			versioninfo.CommonFileld.Remarks,
		)
	} else {
		r, err = conn.Exec(SQL,
			versioninfo.ArmVersion,
			versioninfo.UpgradeVersion,
			versioninfo.WebVersion,
			versioninfo.WebServerName,
			versioninfo.CommonFileld.CREATED_BY,
			versioninfo.CommonFileld.CREATED_TIME,
			versioninfo.CommonFileld.UPDATED_BY,
			versioninfo.CommonFileld.UPDATED_TIME,
			versioninfo.CommonFileld.IsDeleted,
			versioninfo.CommonFileld.Remarks,
			versioninfo.Id)
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
