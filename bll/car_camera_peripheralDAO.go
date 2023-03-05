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
	insertCarsettingSQL = `INSERT INTO carsetting(Id, LocomotiveID, HostSide,CREATED_BY, CREATED_TIME, UPDATED_BY, UPDATED_TIME, IsDeleted, Remarks) 
VALUES (?,?,?,?,?,?,?,?,?)`
	insertSoundsettingSQL = `INSERT INTO soundsetting(Id,SoundType,Volume,CustomSound,CREATED_BY,CREATED_TIME, UPDATED_BY, UPDATED_TIME, IsDeleted, Remarks) 
VALUES (?,?,?,?,?,?,?,?,?,?)`
	//信息列表语句(分页)
	listCamerasettingSQL = `select id,Camera_Name,Camera_OnOff,Camera_IP,ChannelNo,ChannelName,CREATED_BY, CREATED_TIME, UPDATED_BY,
	   UPDATED_TIME,isDeleted,Remarks from camerasetting   LIMIT ? OFFSET ?`
	listPeripheralsettingsSQL = `select id,PeripheralType,Peripheral_OnOff,CREATED_BY,CREATED_TIME, UPDATED_BY,
	   UPDATED_TIME,isDeleted,Remarks from peripheralsettings   LIMIT ? OFFSET ?`
	//根据编号查询单个信息
	SingleCarsettingSQL = `select id,LocomotiveID,HostSide,CREATED_BY, CREATED_TIME, UPDATED_BY,
       UPDATED_TIME,isDeleted,Remarks from carsetting where Id=?`
	SingleSoundsettingSQL = `select id,SoundType,Volume,CustomSound,CREATED_BY, CREATED_TIME, UPDATED_BY,
       UPDATED_TIME,isDeleted,Remarks from soundsetting where Id=?`
)

//camerasetting,carsetting,peripheralsettings,soundsetting

//获取车辆配置信息
func GetCarsettingInfo(id int64) (*viewmodel.CarsettingViewModel, error) {
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	r, err := conn.Query(SingleCarsettingSQL, id)
	if err != nil {
		return nil, err
	}
	carsettingViewModel := &viewmodel.CarsettingViewModel{}
	if r.Next() {
		var id int64
		var locomotiveID, CREATED_BY, CREATED_TIME, UPDATED_BY, UPDATED_TIME, Remarks string
		var isDeleted, hostSide int
		err := r.Scan(&id, &locomotiveID, &hostSide, &CREATED_BY, &CREATED_TIME,
			&UPDATED_BY, &UPDATED_TIME, &isDeleted, &Remarks)
		if err != nil {
			return nil, err
		}
		carsettingViewModel = &viewmodel.CarsettingViewModel{
			Id:           strconv.FormatInt(id, 10), //编号(主键)
			LocomotiveID: locomotiveID,              //机车号
			HostSide:     hostSide,                  //主机所在端,1:一端,2:二端
		}
		return carsettingViewModel, nil
	}
	defer r.Close()
	err = r.Err()
	if err != nil {
		return nil, err
	}
	return carsettingViewModel, nil
}

//获取TTS语音配置信息表
func GetSoundsettingInfo(id int64) (*viewmodel.SoundsettingViewModel, error) {
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	r, err := conn.Query(SingleSoundsettingSQL, id)
	if err != nil {
		return nil, err
	}
	soundsettingViewModel := &viewmodel.SoundsettingViewModel{}
	if r.Next() {
		var id int64
		var customSound, CREATED_BY, CREATED_TIME, UPDATED_BY, UPDATED_TIME, Remarks string
		var isDeleted, soundType, volume int
		err := r.Scan(&id, &soundType, &volume, &customSound, &CREATED_BY, &CREATED_TIME,
			&UPDATED_BY, &UPDATED_TIME, &isDeleted, &Remarks)
		if err != nil {
			return nil, err
		}
		soundsettingViewModel = &viewmodel.SoundsettingViewModel{
			Id:          strconv.FormatInt(id, 10), //编号(主键)
			SoundType:   soundType,                 //机车号
			Volume:      volume,                    //主机所在端,1:一端,2:二端
			CustomSound: customSound,
		}
		return soundsettingViewModel, nil
	}
	defer r.Close()
	err = r.Err()
	if err != nil {
		return nil, err
	}
	return soundsettingViewModel, nil
}

// 获取信息列表(分页)-摄像头配置信息
func GetCamerasettingPageList(pageNumber int, pageSize int) ([]*viewmodel.CamerasettingViewModel, error) {
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	var result []*viewmodel.CamerasettingViewModel
	if err != nil {
		return result, err
	}
	offset := (pageNumber - 1) * pageSize
	r, err := conn.Query(listCamerasettingSQL, pageSize, offset)
	if err != nil {
		return result, err
	}
	for r.Next() {
		var id int64
		var camera_Name, camera_IP, channelName, CREATED_BY, CREATED_TIME, UPDATED_BY, UPDATED_TIME, Remarks string
		var isDeleted, camera_OnOff, channelNo int
		err := r.Scan(&id, &camera_Name, &camera_OnOff, &camera_IP, &channelNo, &channelName, &CREATED_BY, &CREATED_TIME,
			&UPDATED_BY, &UPDATED_TIME, &isDeleted, &Remarks)
		if err != nil {
			return nil, err
		}
		camerasettingViewModel := &viewmodel.CamerasettingViewModel{
			Id:           strconv.FormatInt(id, 10), //编号(主键)
			Camera_Name:  camera_Name,
			Camera_OnOff: camera_OnOff,
			Camera_IP:    camera_IP,
			ChannelNo:    channelNo,
			ChannelName:  channelName,
		}
		result = append(result, camerasettingViewModel)
	}
	defer r.Close()
	err = r.Err()
	if err != nil {
		return result, err
	}
	return result, nil
}

// 获取信息列表(分页)-外设设置信息表
func GetPeripheralsettingsPageList(pageNumber int, pageSize int) ([]*viewmodel.PeripheralsettingsViewModel, error) {
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	var result []*viewmodel.PeripheralsettingsViewModel
	if err != nil {
		return result, err
	}
	offset := (pageNumber - 1) * pageSize
	r, err := conn.Query(listPeripheralsettingsSQL, pageSize, offset)
	if err != nil {
		return result, err
	}
	for r.Next() {
		var id int64
		var CREATED_BY, CREATED_TIME, UPDATED_BY, UPDATED_TIME, Remarks string
		var isDeleted, peripheralType, peripheral_OnOff int
		err := r.Scan(&id, &peripheralType, &peripheral_OnOff, &CREATED_BY, &CREATED_TIME,
			&UPDATED_BY, &UPDATED_TIME, &isDeleted, &Remarks)
		if err != nil {
			return nil, err
		}
		peripheralsettingsViewModel := &viewmodel.PeripheralsettingsViewModel{
			Id:               strconv.FormatInt(id, 10), //编号(主键)
			PeripheralType:   peripheralType,
			Peripheral_OnOff: peripheral_OnOff,
		}
		result = append(result, peripheralsettingsViewModel)
	}
	defer r.Close()
	err = r.Err()
	if err != nil {
		return result, err
	}
	return result, nil
}

// insert camerasetting,carsetting,peripheralsettings,soundsetting 事务添加四个信息表
func InsertSettings(carsetting *domain.Carsetting, camerasetting []*domain.Camerasetting,
	peripheralsettings []*domain.Peripheralsettings, soundsetting *domain.Soundsetting) error {
	// 使用事务，四个表中有一个表的插入有错，则将回滚报错
	return util.ExecTransaction(func(tx *sql.Tx) error {
		_, err := tx.Exec(insertCarsettingSQL,
			carsetting.Id,
			carsetting.LocomotiveID,
			carsetting.HostSide,
			carsetting.CommonFileld.CREATED_BY,
			carsetting.CommonFileld.CREATED_TIME,
			carsetting.CommonFileld.UPDATED_BY,
			carsetting.CommonFileld.UPDATED_TIME,
			carsetting.CommonFileld.IsDeleted,
			carsetting.CommonFileld.Remarks)
		if err != nil {
			tx.Rollback()
			return err
		}
		err = BatchCamerasetting(tx, camerasetting)
		if err != nil {
			tx.Rollback()
			return err
		}
		err = BatchPeripheralsettings(tx, peripheralsettings)
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = tx.Exec(insertSoundsettingSQL,
			soundsetting.Id,
			soundsetting.SoundType,
			soundsetting.Volume,
			soundsetting.CustomSound,
			soundsetting.CommonFileld.CREATED_BY,
			soundsetting.CommonFileld.CREATED_TIME,
			soundsetting.CommonFileld.UPDATED_BY,
			soundsetting.CommonFileld.UPDATED_TIME,
			soundsetting.CommonFileld.IsDeleted,
			soundsetting.CommonFileld.Remarks)
		if err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
}

//批量添加外设配置
func BatchPeripheralsettings(tx *sql.Tx, peripheralsettings []*domain.Peripheralsettings) error {
	count := len(peripheralsettings)
	// 存放 (?, ?) 的slice
	vStrings := make([]string, 0, count)
	// 存放values的slice
	vArgs := make([]interface{}, 0, count*2)
	//遍历
	for _, a := range peripheralsettings {
		// 此处占位符要与插入值的个数对应
		vStrings = append(vStrings, "(?,?,?,?,?,?,?,?,?)")
		vArgs = append(vArgs, a.Id)
		vArgs = append(vArgs, a.PeripheralType)
		vArgs = append(vArgs, a.Peripheral_OnOff)
		vArgs = append(vArgs, a.CommonFileld.CREATED_BY)
		vArgs = append(vArgs, a.CommonFileld.CREATED_TIME)
		vArgs = append(vArgs, a.CommonFileld.UPDATED_BY)
		vArgs = append(vArgs, a.CommonFileld.UPDATED_TIME)
		vArgs = append(vArgs, a.CommonFileld.IsDeleted)
		vArgs = append(vArgs, a.CommonFileld.Remarks)
	}
	sqlHead := `INSERT INTO peripheralsettings(Id,PeripheralType,Peripheral_OnOff,CREATED_BY,CREATED_TIME,UPDATED_BY,UPDATED_TIME,IsDeleted,Remarks) VALUES %s`
	stmt := fmt.Sprintf(sqlHead, strings.Join(vStrings, ","))
	_, err := tx.Exec(stmt, vArgs...)
	if err != nil {
		return err
	}
	return nil
}

//批量添加摄像头配置
func BatchCamerasetting(tx *sql.Tx, camerasettings []*domain.Camerasetting) error {
	count := len(camerasettings)
	// 存放 (?, ?) 的slice
	vStrings := make([]string, 0, count)
	// 存放values的slice
	vArgs := make([]interface{}, 0, count*2)
	//遍历
	for _, a := range camerasettings {
		// 此处占位符要与插入值的个数对应
		vStrings = append(vStrings, "(?,?,?,?,?,?,?,?,?,?,?,?)")
		vArgs = append(vArgs, a.Id)
		vArgs = append(vArgs, a.Camera_Name)
		vArgs = append(vArgs, a.Camera_OnOff)
		vArgs = append(vArgs, a.Camera_IP)
		vArgs = append(vArgs, a.ChannelNo)
		vArgs = append(vArgs, a.ChannelName)
		vArgs = append(vArgs, a.CommonFileld.CREATED_BY)
		vArgs = append(vArgs, a.CommonFileld.CREATED_TIME)
		vArgs = append(vArgs, a.CommonFileld.UPDATED_BY)
		vArgs = append(vArgs, a.CommonFileld.UPDATED_TIME)
		vArgs = append(vArgs, a.CommonFileld.IsDeleted)
		vArgs = append(vArgs, a.CommonFileld.Remarks)
	}
	sqlHead := `INSERT INTO camerasetting(Id,Camera_Name,Camera_OnOff,Camera_IP,ChannelNo,ChannelName,CREATED_BY,CREATED_TIME,UPDATED_BY,UPDATED_TIME,IsDeleted,Remarks) VALUES %s`
	stmt := fmt.Sprintf(sqlHead, strings.Join(vStrings, ","))
	_, err := tx.Exec(stmt, vArgs...)
	if err != nil {
		return err
	}
	return nil
}
