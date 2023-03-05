package bll

import (
	"database/sql"
	"fmt"
	"go_web/domain"
	"go_web/util"
	"go_web/viewmodel"
	"strconv"
	"strings"
)

const (
	//插入信息语句
	insertFatiguedrivingalarmSQL = `INSERT INTO carsetting(Id, AlarmTypeName,Sort,Speed,Time,TTS,CREATED_BY,CD,AlarmSw,AlarmSoundSw,UploadingPicSw,
UploadingVideoSw,CREATED_TIME, UPDATED_BY, UPDATED_TIME, IsDeleted, Remarks) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?,?,?,?,?,?,?,?,?,?)`
	insertDrivingrecordSQL = `INSERT INTO soundsetting(Id, Name, IsOpen,CREATED_BY, CREATED_TIME, UPDATED_BY, UPDATED_TIME, IsDeleted, Remarks) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?,?)`
	//信息列表语句(分页)
	listFatiguedrivingalarmSQL = `select id,AlarmTypeName,Sort,Speed,Time,TTS,CD,AlarmSw,AlarmSoundSw,UploadingPicSw,UploadingVideoSw,CREATED_BY, CREATED_TIME, UPDATED_BY,
	   UPDATED_TIME,isDeleted,Remarks from fatiguedrivingalarm order by Sort asc  LIMIT ? OFFSET ?`
	listDrivingrecordSQL = `select id,Name,IsOpen,CREATED_BY,CREATED_TIME, UPDATED_BY,
	   UPDATED_TIME,isDeleted,Remarks from drivingrecord   LIMIT ? OFFSET ?`
)

// 获取信息列表(分页)-报警配置信息
func GetFatiguedrivingalarmPageList(pageNumber int, pageSize int) ([]*viewmodel.FatiguedrivingalarmViewModel, error) {
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	var result []*viewmodel.FatiguedrivingalarmViewModel
	if err != nil {
		return result, err
	}
	offset := (pageNumber - 1) * pageSize
	r, err := conn.Query(listFatiguedrivingalarmSQL, pageSize, offset)
	if err != nil {
		return result, err
	}
	for r.Next() {
		var id int64
		var alarmTypeName, tts, CREATED_BY, CREATED_TIME, UPDATED_BY, UPDATED_TIME, Remarks string
		var isDeleted, sort, alarmSw, alarmSoundSw, uploadingPicSw, uploadingVideoSw int
		var speed, time, cd float64
		err := r.Scan(&id,
			&alarmTypeName,
			&sort,
			&speed,
			&time,
			&tts,
			&cd,
			&alarmSw,
			&alarmSoundSw,
			&uploadingPicSw, &uploadingVideoSw, &CREATED_BY, &CREATED_TIME,
			&UPDATED_BY, &UPDATED_TIME, &isDeleted, &Remarks)
		if err != nil {
			return nil, err
		}
		fatiguedrivingalarmViewModel := &viewmodel.FatiguedrivingalarmViewModel{
			Id:               strconv.FormatInt(id, 10), //编号(主键)
			AlarmTypeName:    alarmTypeName,
			Sort:             sort,
			Speed:            speed,
			Time:             time,
			TTS:              tts,
			CD:               cd,
			AlarmSw:          alarmSw,
			UploadingPicSw:   uploadingPicSw,
			UploadingVideoSw: uploadingVideoSw,
		}
		result = append(result, fatiguedrivingalarmViewModel)
	}
	defer r.Close()
	err = r.Err()
	if err != nil {
		return result, err
	}
	return result, nil
}

// 获取信息列表(分页)-行驶记录设置信息表
func GetDrivingrecordPageList(pageNumber int, pageSize int) ([]*viewmodel.DrivingrecordViewModel, error) {
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	var result []*viewmodel.DrivingrecordViewModel
	if err != nil {
		return result, err
	}
	offset := (pageNumber - 1) * pageSize
	r, err := conn.Query(listDrivingrecordSQL, pageSize, offset)
	if err != nil {
		return result, err
	}
	for r.Next() {
		var id int64
		var name, CREATED_BY, CREATED_TIME, UPDATED_BY, UPDATED_TIME, Remarks string
		var isDeleted, isOpen int
		err := r.Scan(&id, &name, &isOpen, &CREATED_BY, &CREATED_TIME,
			&UPDATED_BY, &UPDATED_TIME, &isDeleted, &Remarks)
		if err != nil {
			return nil, err
		}
		drivingrecordViewModel := &viewmodel.DrivingrecordViewModel{
			Id:     strconv.FormatInt(id, 10), //编号(主键)
			Name:   name,
			IsOpen: isOpen,
		}
		result = append(result, drivingrecordViewModel)
	}
	defer r.Close()
	err = r.Err()
	if err != nil {
		return result, err
	}
	return result, nil
}

// 报警信息，行驶记录设置添加 事务添加两个信息表
func InsertAlarmSettings(fatiguedrivingalarms []*domain.Fatiguedrivingalarm, cdrivingrecords []*domain.Drivingrecord) error {
	// 使用事务，两个表中有一个表的插入有错，则将回滚报错
	return util.ExecTransaction(func(tx *sql.Tx) error {
		err := BatchFatiguedrivingalarm(tx, fatiguedrivingalarms)
		if err != nil {
			tx.Rollback()
			return err
		}
		err = BatchDrivingrecord(tx, cdrivingrecords)
		if err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
}

//批量添加报警信息配置
func BatchFatiguedrivingalarm(tx *sql.Tx, fatiguedrivingalarms []*domain.Fatiguedrivingalarm) error {
	count := len(fatiguedrivingalarms)
	// 存放 (?, ?) 的slice
	vStrings := make([]string, 0, count)
	// 存放values的slice
	vArgs := make([]interface{}, 0, count*2)
	//遍历
	for _, a := range fatiguedrivingalarms {
		// 此处占位符要与插入值的个数对应
		vStrings = append(vStrings, "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
		vArgs = append(vArgs, a.Id)
		vArgs = append(vArgs, a.AlarmTypeName)
		vArgs = append(vArgs, a.Sort)
		vArgs = append(vArgs, a.Speed)
		vArgs = append(vArgs, a.Time)
		vArgs = append(vArgs, a.TTS)
		vArgs = append(vArgs, a.CD)
		vArgs = append(vArgs, a.AlarmSw)
		vArgs = append(vArgs, a.AlarmSoundSw)
		vArgs = append(vArgs, a.UploadingPicSw)
		vArgs = append(vArgs, a.UploadingVideoSw)
		vArgs = append(vArgs, a.CommonFileld.CREATED_BY)
		vArgs = append(vArgs, a.CommonFileld.CREATED_TIME)
		vArgs = append(vArgs, a.CommonFileld.UPDATED_BY)
		vArgs = append(vArgs, a.CommonFileld.UPDATED_TIME)
		vArgs = append(vArgs, a.CommonFileld.IsDeleted)
		vArgs = append(vArgs, a.CommonFileld.Remarks)
	}
	sqlHead := `INSERT INTO fatiguedrivingalarm(Id,AlarmTypeName,Sort,Speed,Time,TTS,CD,AlarmSw,AlarmSoundSw,UploadingPicSw,
UploadingVideoSw,CREATED_BY,CREATED_TIME, UPDATED_BY, UPDATED_TIME, IsDeleted, Remarks)  VALUES %s`
	stmt := fmt.Sprintf(sqlHead, strings.Join(vStrings, ","))
	_, err := tx.Exec(stmt, vArgs...)
	if err != nil {
		return err
	}
	return nil
}

//批量行驶记录信息配置
func BatchDrivingrecord(tx *sql.Tx, drivingrecords []*domain.Drivingrecord) error {
	count := len(drivingrecords)
	// 存放 (?, ?) 的slice
	vStrings := make([]string, 0, count)
	// 存放values的slice
	vArgs := make([]interface{}, 0, count*2)
	//遍历
	for _, a := range drivingrecords {
		// 此处占位符要与插入值的个数对应
		vStrings = append(vStrings, "(?, ?, ?, ?, ?, ?, ?, ?,?)")
		vArgs = append(vArgs, a.Id)
		vArgs = append(vArgs, a.Name)
		vArgs = append(vArgs, a.IsOpen)
		vArgs = append(vArgs, a.CommonFileld.CREATED_BY)
		vArgs = append(vArgs, a.CommonFileld.CREATED_TIME)
		vArgs = append(vArgs, a.CommonFileld.UPDATED_BY)
		vArgs = append(vArgs, a.CommonFileld.UPDATED_TIME)
		vArgs = append(vArgs, a.CommonFileld.IsDeleted)
		vArgs = append(vArgs, a.CommonFileld.Remarks)
	}
	sqlHead := `INSERT INTO drivingrecord(Id,Name,IsOpen,CREATED_BY, CREATED_TIME, UPDATED_BY, UPDATED_TIME, IsDeleted, Remarks) VALUES %s`
	stmt := fmt.Sprintf(sqlHead, strings.Join(vStrings, ","))
	_, err := tx.Exec(stmt, vArgs...)
	if err != nil {
		return err
	}
	return nil
}
