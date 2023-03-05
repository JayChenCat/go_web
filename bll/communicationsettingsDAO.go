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
	insertCommunicationsettingsSQL = `INSERT INTO communicationsettings(Id, CommunicationType, IP, Port,CREATED_BY, CREATED_TIME, UPDATED_BY, UPDATED_TIME, IsDeleted, Remarks) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?,?,?)`
	//修改信息语句
	updateCommunicationsettingsSQL = `update communicationsettings set CommunicationType=?, IP=?, Port=?,CREATED_BY=?, CREATED_TIME=?, 
UPDATED_BY=?, UPDATED_TIME=?, IsDeleted=?, Remarks=? where Id=?`
	//删除信息语句
	deleteCommunicationsettingsSQL = `delete from communicationsettings where Id=?`
	//信息列表语句(分页)
	listCommunicationsettingsSQL = `select id,CommunicationType,IP,Port,CREATED_BY, CREATED_TIME, UPDATED_BY,
UPDATED_TIME,isDeleted,Remarks from communicationsettings   LIMIT ? OFFSET ?`
	//根据编号查询单个信息
	SingleCommunicationsettingsSQL = `select id,CommunicationType,IP,Port,CREATED_BY, CREATED_TIME, UPDATED_BY,
       UPDATED_TIME,isDeleted,Remarks from communicationsettings where Id=?`
	//统计列表信息数量
	TotalCommunicationsettingsSQL = `select count(1) total from versioninfo`
)

func scanCommunicationsettingsWithSignOnAndBannerData(r *sql.Rows) (*domain.Communicationsettings, error) {
	var id, port int64
	var ip, CREATED_BY, CREATED_TIME, UPDATED_BY, UPDATED_TIME, Remarks string
	var communicationType, isDeleted int

	err := r.Scan(&id, &communicationType, &ip, &port, &CREATED_BY, &CREATED_TIME,
		&UPDATED_BY, &UPDATED_TIME, &isDeleted, &Remarks)

	baseFileld := domain.BaseFileld{
		CREATED_BY:   CREATED_BY,   //创建人
		CREATED_TIME: CREATED_TIME, //创建时间
		UPDATED_BY:   UPDATED_BY,   //更新人
		UPDATED_TIME: UPDATED_TIME, //更新时间
		IsDeleted:    isDeleted,    //是否删除(逻辑删除)，1是，0否
		Remarks:      Remarks,      //备注
	}

	a := &domain.Communicationsettings{
		Id:                id,                //编号(主键)
		CommunicationType: communicationType, //通讯类型(1表示本机，2表示OSC)
		IP:                ip,                //IP地址
		Port:              port,              //端口号
		CommonFileld:      baseFileld,        //公共字段(创建时间，创建人等属性)
	}
	return a, err
}

//根据编号获取单个信息
func GetCommunicationsettingsByUserId(userId int64) (*viewmodel.CommunicationsettingsViewModel, error) {
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
	communicationsettingsViewModel := &viewmodel.CommunicationsettingsViewModel{}
	if data.Next() {
		communicationsettings, err := scanCommunicationsettingsWithSignOnAndBannerData(data)
		if err != nil {
			return nil, err
		}
		communicationsettingsViewModel = &viewmodel.CommunicationsettingsViewModel{
			Id:                strconv.FormatInt(communicationsettings.Id, 10),   //编号(主键)
			CommunicationType: communicationsettings.CommunicationType,           //通讯类型(1表示本机，2表示OSC)
			IP:                communicationsettings.IP,                          //IP地址
			Port:              strconv.FormatInt(communicationsettings.Port, 10), //端口号
		}
		return communicationsettingsViewModel, nil
	}
	defer data.Close()
	err = data.Err()
	if err != nil {
		return nil, err
	}
	return communicationsettingsViewModel, err
}

// 获取信息列表(分页)-Account
func GetCommunicationsettingsList(pageNumber int, pageSize int) ([]*viewmodel.CommunicationsettingsViewModel, error) {
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	var result []*viewmodel.CommunicationsettingsViewModel
	if err != nil {
		return result, err
	}
	offset := (pageNumber - 1) * pageSize
	r, err := conn.Query(listCommunicationsettingsSQL, pageSize, offset)
	if err != nil {
		return result, err
	}
	for r.Next() {
		communicationsettings, err := scanCommunicationsettingsWithSignOnAndBannerData(r)
		if err != nil {
			log.Printf("error: %v", err.Error())
			continue
		}
		//createtime := equipmentinfo.CommonFileld.CREATED_TIME
		//_time, _ := time.Parse(time.RFC3339, createtime)
		CommunicationsettingsViewModel := &viewmodel.CommunicationsettingsViewModel{
			Id:                strconv.FormatInt(communicationsettings.Id, 10),   //编号(主键)
			CommunicationType: communicationsettings.CommunicationType,           //通讯类型(1表示本机，2表示OSC)
			IP:                communicationsettings.IP,                          //IP地址
			Port:              strconv.FormatInt(communicationsettings.Port, 10), //端口号
		}
		result = append(result, CommunicationsettingsViewModel)
	}
	defer r.Close()
	err = r.Err()
	if err != nil {
		return result, err
	}
	return result, nil
}

// insert Communicationsettings
func InsertCommunicationsettings(communicationsettings *domain.Communicationsettings) error {
	err := InsertOrUpdateCommunicationsettings(insertCommunicationsettingsSQL, 0, communicationsettings)
	return err
}

// update Communicationsettings
func UpdateCommunicationsettings(communicationsettings *domain.Communicationsettings) error {
	err := InsertOrUpdateCommunicationsettings(updateCommunicationsettingsSQL, 1, communicationsettings)
	return err
}

// Delete Communicationsettings
func DeleteCommunicationsettings(userId int64) error {
	err := util.Delete(deleteAccountmanagerSQL, userId)
	return err
}

//统计列表信息数量
func GetCommunicationsettingsTotal() (int, error) {
	var total int
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	if err != nil {
		return 0, err
	}
	r, err := conn.Query(TotalCommunicationsettingsSQL)
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
func InsertOrUpdateCommunicationsettings(SQL string, edit int, communicationsettings *domain.Communicationsettings) error {
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
			communicationsettings.Id,
			communicationsettings.CommunicationType,
			communicationsettings.IP,
			communicationsettings.Port,
			communicationsettings.CommonFileld.CREATED_BY,
			communicationsettings.CommonFileld.CREATED_TIME,
			communicationsettings.CommonFileld.UPDATED_BY,
			communicationsettings.CommonFileld.UPDATED_TIME,
			communicationsettings.CommonFileld.IsDeleted,
			communicationsettings.CommonFileld.Remarks,
		)
	} else {
		r, err = conn.Exec(SQL,
			communicationsettings.CommunicationType,
			communicationsettings.IP,
			communicationsettings.Port,
			communicationsettings.CommonFileld.CREATED_BY,
			communicationsettings.CommonFileld.CREATED_TIME,
			communicationsettings.CommonFileld.UPDATED_BY,
			communicationsettings.CommonFileld.UPDATED_TIME,
			communicationsettings.CommonFileld.IsDeleted,
			communicationsettings.CommonFileld.Remarks,
			communicationsettings.Id)
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
func BatchInsertCommunicationsettings(communicationsettings []*domain.Communicationsettings) error {
	count := len(communicationsettings)
	// 存放 (?, ?) 的slice
	vStrings := make([]string, 0, count)
	// 存放values的slice
	vArgs := make([]interface{}, 0, count*2)
	//遍历
	for _, a := range communicationsettings {
		// 此处占位符要与插入值的个数对应
		vStrings = append(vStrings, "(?, ?, ?, ?, ?, ?, ?, ?,?,?)")
		vArgs = append(vArgs, a.Id)
		vArgs = append(vArgs, a.CommunicationType)
		vArgs = append(vArgs, a.IP)
		vArgs = append(vArgs, a.Port)
		vArgs = append(vArgs, a.CommonFileld.CREATED_BY)
		vArgs = append(vArgs, a.CommonFileld.CREATED_TIME)
		vArgs = append(vArgs, a.CommonFileld.UPDATED_BY)
		vArgs = append(vArgs, a.CommonFileld.UPDATED_TIME)
		vArgs = append(vArgs, a.CommonFileld.IsDeleted)
		vArgs = append(vArgs, a.CommonFileld.Remarks)
	}
	sqlHead := `INSERT INTO communicationsettings(Id,CommunicationType,IP,
		Port,CREATED_BY,CREATED_TIME,UPDATED_BY,UPDATED_TIME,IsDeleted,Remarks) VALUES %s`
	stmt := fmt.Sprintf(sqlHead, strings.Join(vStrings, ","))
	return util.BatchInsertOrUpdate(stmt, vArgs)
}
