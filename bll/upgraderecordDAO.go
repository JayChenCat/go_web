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
	insertUpgraderecordSQL = `INSERT INTO upgraderecord(Id, UpgradeName, UpgradeFileName,UpgradeVersion,CREATED_BY, CREATED_TIME, UPDATED_BY, UPDATED_TIME, IsDeleted, Remarks) 
VALUES (?,?,?,?,?,?,?,?,?,?)`
	//修改信息语句
	updateUpgraderecordSQL = `update upgraderecord set UpgradeName=?,UpgradeFileName=?,UpgradeVersion=?,CREATED_BY=?, CREATED_TIME=?, 
UPDATED_BY=?, UPDATED_TIME=?, IsDeleted=?, Remarks=? where Id=?`
	//删除信息语句
	deleteUpgraderecordSQL = `delete from upgraderecord where Id=?`
	//信息列表语句(分页)
	listUpgraderecordSQL = `select id,UpgradeName,UpgradeFileName,UpgradeVersion,CREATED_BY, CREATED_TIME, UPDATED_BY,
UPDATED_TIME,isDeleted,Remarks from upgraderecord   LIMIT ? OFFSET ?`
	//根据编号查询单个信息
	SingleUpgraderecordSQL = `select id,UpgradeName,UpgradeFileName,UpgradeVersion,CREATED_BY, CREATED_TIME, UPDATED_BY,
       UPDATED_TIME,isDeleted,Remarks from upgraderecord where Id=?`
	//统计用户列表信息数量
	TotalUpgraderecordSQL = `select count(1) total from upgraderecord`
)

func scanUpgraderecordWithSignOnAndBannerData(r *sql.Rows) (*domain.Upgraderecord, error) {
	var id int64
	var upgradeName, upgradeFileName, upgradeVersion, CREATED_BY, CREATED_TIME, UPDATED_BY, UPDATED_TIME, Remarks string
	var isDeleted int

	err := r.Scan(&id, &upgradeName, &upgradeFileName, &upgradeVersion, &CREATED_BY, &CREATED_TIME,
		&UPDATED_BY, &UPDATED_TIME, &isDeleted, &Remarks)

	baseFileld := domain.BaseFileld{
		CREATED_BY:   CREATED_BY,   //创建人
		CREATED_TIME: CREATED_TIME, //创建时间
		UPDATED_BY:   UPDATED_BY,   //更新人
		UPDATED_TIME: UPDATED_TIME, //更新时间
		IsDeleted:    isDeleted,    //是否删除(逻辑删除)，1是，0否
		Remarks:      Remarks,      //备注
	}

	a := &domain.Upgraderecord{
		Id:              id,              //编号(主键)
		UpgradeName:     upgradeName,     //升级的设备名称
		UpgradeFileName: upgradeFileName, //升级文件包名称
		UpgradeVersion:  upgradeVersion,  //升级的版本
		CommonFileld:    baseFileld,      //公共字段(创建时间，创建人等属性)
	}
	return a, err
}

//根据编号获取单个信息
func GetUpgraderecordByUserId(userId int64) (*viewmodel.UpgraderecordViewModel, error) {
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	if err != nil {
		return nil, err
	}
	data, err := conn.Query(SingleUpgraderecordSQL, userId)
	if err != nil {
		return nil, err
	}
	upgraderecordViewModel := &viewmodel.UpgraderecordViewModel{}
	if data.Next() {
		upgraderecord, err := scanUpgraderecordWithSignOnAndBannerData(data)
		if err != nil {
			return nil, err
		}
		upgraderecordViewModel = &viewmodel.UpgraderecordViewModel{
			Id:              strconv.FormatInt(upgraderecord.Id, 10), //编号(主键)
			UpgradeName:     upgraderecord.UpgradeName,               //升级的设备名称
			UpgradeFileName: upgraderecord.UpgradeFileName,           //升级文件包名称
			UpgradeVersion:  upgraderecord.UpgradeVersion,            //升级的版本
		}
		return upgraderecordViewModel, nil
	}
	defer data.Close()
	err = data.Err()
	if err != nil {
		return nil, err
	}
	return upgraderecordViewModel, err
}

// 获取信息列表(分页)-Upgraderecord
func GetUpgraderecordList(pageNumber int, pageSize int) ([]*viewmodel.UpgraderecordViewModel, error) {
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	var result []*viewmodel.UpgraderecordViewModel
	if err != nil {
		return result, err
	}
	offset := (pageNumber - 1) * pageSize
	r, err := conn.Query(listUpgraderecordSQL, pageSize, offset)
	if err != nil {
		return result, err
	}
	for r.Next() {
		upgraderecord, err := scanUpgraderecordWithSignOnAndBannerData(r)
		if err != nil {
			log.Printf("error: %v", err.Error())
			continue
		}
		//createtime := equipmentinfo.CommonFileld.CREATED_TIME
		//_time, _ := time.Parse(time.RFC3339, createtime)
		upgraderecordViewModel := &viewmodel.UpgraderecordViewModel{
			Id:              strconv.FormatInt(upgraderecord.Id, 10), //用户编号(主键)
			UpgradeName:     upgraderecord.UpgradeName,               //升级的设备名称
			UpgradeFileName: upgraderecord.UpgradeFileName,           //升级文件包名称
			UpgradeVersion:  upgraderecord.UpgradeVersion,            //升级的版本
		}
		result = append(result, upgraderecordViewModel)
	}
	defer r.Close()
	err = r.Err()
	if err != nil {
		return result, err
	}
	return result, nil
}

// insert Upgraderecord
func InsertUpgraderecord(upgraderecord *domain.Upgraderecord) error {
	err := InsertOrUpgraderecord(insertUpgraderecordSQL, 0, upgraderecord)
	return err
}

// update Upgraderecord
func UpdateUpgraderecord(upgraderecord *domain.Upgraderecord) error {
	err := InsertOrUpgraderecord(updateUpgraderecordSQL, 1, upgraderecord)
	return err
}

// Delete Upgraderecord
func DeleteUpgraderecord(userId int64) error {
	err := util.Delete(deleteUpgraderecordSQL, userId)
	return err
}

//统计列表信息数量
func GetUpgraderecordTotal() (int, error) {
	var total int
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	if err != nil {
		return 0, err
	}
	r, err := conn.Query(TotalUpgraderecordSQL)
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
func InsertOrUpgraderecord(SQL string, edit int, versioninfo *domain.Upgraderecord) error {
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	if err != nil {
		return err
	}
	var r sql.Result

	//插入
	if edit == 0 {
		r, err = conn.Exec(SQL,
			versioninfo.Id,
			versioninfo.UpgradeName,
			versioninfo.UpgradeFileName,
			versioninfo.UpgradeVersion,
			versioninfo.CommonFileld.CREATED_BY,
			versioninfo.CommonFileld.CREATED_TIME,
			versioninfo.CommonFileld.UPDATED_BY,
			versioninfo.CommonFileld.UPDATED_TIME,
			versioninfo.CommonFileld.IsDeleted,
			versioninfo.CommonFileld.Remarks,
		)
	} else {
		r, err = conn.Exec(SQL,
			versioninfo.UpgradeName,
			versioninfo.UpgradeFileName,
			versioninfo.UpgradeVersion,
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
