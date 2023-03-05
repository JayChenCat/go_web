package service

import (
	"go_web/bll"
	"go_web/domain"
	"go_web/viewmodel"
	"strconv"
	"time"
)

//提交系统设置
func SaveSystemSettings(carsetting *viewmodel.CarsettingViewModel, camerasetting []*viewmodel.CamerasettingViewModel,
	peripheralsettings []*viewmodel.PeripheralsettingsViewModel, soundsetting *viewmodel.SoundsettingViewModel) (bool, error) {
	_time := time.Now().Format("2006-01-02 15:04:05")
	commonFileld := domain.BaseFileld{
		CREATED_BY:   "admin", //创建人
		CREATED_TIME: _time,   //创建时间
		UPDATED_BY:   "admin", //更新人
		UPDATED_TIME: _time,   //更新时间
		IsDeleted:    0,       //是否删除(逻辑删除)，1是，0否
		Remarks:      "",      //备注
	}

	num, _ := strconv.ParseInt(carsetting.Id, 0, 0)
	_carsetting := &domain.Carsetting{
		Id:           num, //编号(主键)
		LocomotiveID: carsetting.LocomotiveID,
		HostSide:     carsetting.HostSide,
		CommonFileld: commonFileld,
	}

	var resultCamerasetting []*domain.Camerasetting
	for _, item := range camerasetting {
		num, _ := strconv.ParseInt(item.Id, 0, 0)
		_camerasetting := &domain.Camerasetting{
			Id:           num, //编号(主键)
			Camera_Name:  item.Camera_Name,
			Camera_OnOff: item.Camera_OnOff,
			Camera_IP:    item.Camera_IP,
			ChannelNo:    item.ChannelNo,
			ChannelName:  item.ChannelName,
			CommonFileld: commonFileld,
		}
		resultCamerasetting = append(resultCamerasetting, _camerasetting)
	}
	var resultPeripheralsettings []*domain.Peripheralsettings
	for _, item := range peripheralsettings {
		num, _ := strconv.ParseInt(item.Id, 0, 0)
		_peripheralsettings := &domain.Peripheralsettings{
			Id:               num, //编号(主键)
			PeripheralType:   item.PeripheralType,
			Peripheral_OnOff: item.Peripheral_OnOff,
			CommonFileld:     commonFileld,
		}
		resultPeripheralsettings = append(resultPeripheralsettings, _peripheralsettings)
	}

	snum, _ := strconv.ParseInt(soundsetting.Id, 0, 0)
	_soundsetting := &domain.Soundsetting{
		Id:           snum, //编号(主键)
		SoundType:    soundsetting.SoundType,
		Volume:       soundsetting.Volume,
		CustomSound:  soundsetting.CustomSound,
		CommonFileld: commonFileld,
	}
	err := bll.InsertSettings(_carsetting, resultCamerasetting, resultPeripheralsettings, _soundsetting)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

//提交报警设置
func SaveAlarmSettings(fatiguedrivingalarms []*viewmodel.FatiguedrivingalarmViewModel,
	cdrivingrecords []*viewmodel.DrivingrecordViewModel) (bool, error) {
	_time := time.Now().Format("2006-01-02 15:04:05")
	commonFileld := domain.BaseFileld{
		CREATED_BY:   "admin", //创建人
		CREATED_TIME: _time,   //创建时间
		UPDATED_BY:   "admin", //更新人
		UPDATED_TIME: _time,   //更新时间
		IsDeleted:    0,       //是否删除(逻辑删除)，1是，0否
		Remarks:      "",      //备注
	}
	var resultFatiguedrivingalarm []*domain.Fatiguedrivingalarm
	for _, item := range fatiguedrivingalarms {
		num, _ := strconv.ParseInt(item.Id, 0, 0)
		_fatiguedrivingalarms := &domain.Fatiguedrivingalarm{
			Id:               num, //编号(主键)
			AlarmTypeName:    item.AlarmTypeName,
			Sort:             item.Sort,
			Speed:            item.Speed,
			Time:             item.Time,
			TTS:              item.TTS,
			CD:               item.CD,
			AlarmSw:          item.AlarmSw,
			AlarmSoundSw:     item.AlarmSoundSw,
			UploadingPicSw:   item.UploadingPicSw,
			UploadingVideoSw: item.UploadingVideoSw,
			CommonFileld:     commonFileld,
		}
		resultFatiguedrivingalarm = append(resultFatiguedrivingalarm, _fatiguedrivingalarms)
	}
	var resultDrivingrecord []*domain.Drivingrecord
	for _, item := range cdrivingrecords {
		num, _ := strconv.ParseInt(item.Id, 0, 0)
		_drivingrecord := &domain.Drivingrecord{
			Id:           num, //编号(主键)
			Name:         item.Name,
			IsOpen:       item.IsOpen,
			CommonFileld: commonFileld,
		}
		resultDrivingrecord = append(resultDrivingrecord, _drivingrecord)
	}
	err := bll.InsertAlarmSettings(resultFatiguedrivingalarm, resultDrivingrecord)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

//获取配置信息
func GetConfiguration() (*viewmodel.SystemConfigurationViewModel, error) {
	drivingrecords, err1 := bll.GetDrivingrecordPageList(1, 10)
	if err1 != nil {
		return nil, err1
	}
	fatiguedrivingalarms, err2 := bll.GetFatiguedrivingalarmPageList(1, 100)
	if err2 != nil {
		return nil, err2
	}
	accountmanagers, err3 := bll.GetAccountmanagerList(1, 10)
	if err3 != nil {
		return nil, err3
	}
	communicationsettings, err4 := bll.GetCommunicationsettingsList(1, 10)
	if err4 != nil {
		return nil, err4
	}
	carsetting, err5 := bll.GetCarsettingInfo(301641553292558336)
	if err5 != nil {
		return nil, err5
	}
	soundsetting, err6 := bll.GetSoundsettingInfo(301641553288364032)
	if err6 != nil {
		return nil, err6
	}
	camerasettings, err7 := bll.GetCamerasettingPageList(1, 10)
	if err7 != nil {
		return nil, err7
	}
	peripheralsettings, err8 := bll.GetPeripheralsettingsPageList(1, 10)
	if err8 != nil {
		return nil, err8
	}
	setting := &viewmodel.SystemConfigurationViewModel{}
	if err1 == nil && err2 == nil && err3 == nil && err4 == nil && err5 == nil && err6 == nil && err7 == nil && err8 == nil {
		alarmSetting := &viewmodel.AlarmSetting{
			FatiguedAlarmSetting:  fatiguedrivingalarms[0],
			CallingAlarmSetting:   fatiguedrivingalarms[1],
			LookaoundAlarmSetting: fatiguedrivingalarms[2],
			SmokingAlarmSetting:   fatiguedrivingalarms[3],
			YawnAlarmSetting:      fatiguedrivingalarms[4],
			NoFaceAlarmSetting:    fatiguedrivingalarms[5],
			MouthAlarmSetting:     fatiguedrivingalarms[6],
			OcclusionAlarmSetting: fatiguedrivingalarms[7],
			HandAlarmSetting:      fatiguedrivingalarms[8],
			ArmAlarmSetting:       fatiguedrivingalarms[9],
			BlockingAlarmSetting:  fatiguedrivingalarms[10],
			Drivingrecord:         drivingrecords,
		}
		systemSetting := &viewmodel.SystemSetting{
			CameraSetting:      camerasettings,
			Carsetting:         carsetting,
			Peripheralsettings: peripheralsettings,
			Soundsetting:       soundsetting,
		}
		setting = &viewmodel.SystemConfigurationViewModel{
			AlarmSetting:         alarmSetting,
			AccountSetting:       accountmanagers,
			CommunicationSetting: communicationsettings,
			SystemSetting:        systemSetting,
		}
		return setting, nil
	}
	return setting, err1
}
