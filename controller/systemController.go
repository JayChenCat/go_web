package controller

import (
	"go_web/config"
	"go_web/domain"
	"go_web/service"
	"go_web/util"
	"go_web/viewmodel"
	"log"
	"net/http"
	"strconv"
)

// 跳转 系统设置主页
func ViewSystem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 跳转到 系统设置主页
		skiplinks(w, r, make(map[string]interface{}), config.SystemPath)
	} else {
		getConfiguration(w, r)
	}
}

//获取配置信息
func getConfiguration(w http.ResponseWriter, r *http.Request) {
	msg := new(domain.Result)
	setting, err := service.GetConfiguration()
	if err != nil {
		log.Printf("获取配置信息 error: %v", err.Error())
	} else {
		if setting != nil {
			msg.Code = 0
			msg.Msg = "获取配置信息成功！"
			//_data, _ := json.Marshal(setting)
			msg.Data = setting
		} else {
			msg.Code = 1
			msg.Msg = "获取配置信息失败！"
		}
	}
	util.ResponseHtml(w, msg, err)
}

//提交配置
func SubmitSetting(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("提交配置信息 error: %v", err.Error())
	}
	tag := r.FormValue("tag") //提交类型的标识
	fromtoken := r.FormValue("token")
	token := GetToken(config.AddTokenKey, r)
	if fromtoken == token {
		switch tag {
		case "1":
			submitAlarmSetting(w, r)
			break
		case "2":
			submitAccountmanager(w, r)
			break
		case "3":
			submitCommunicationsettings(w, r)
			break
		case "4":
			submitSystemSettings(w, r)
			break
		}
		//添加成功后移除token
		errr := RemoveSession(w, r, config.AddTokenKey)
		if errr != nil {
			log.Printf("移除Session error: %v", errr.Error())
		}
	} else {
		msg := new(domain.Result)
		msg.Code = 3
		msg.Msg = "该配置信息已提交，请勿重复提交!"
		util.ResponseHtml(w, msg, err)
	}
}

//1.提交报警设置信息
func submitAlarmSetting(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("提交报警设置信息 error: %v", err.Error())
	}
	//疲劳驾驶报警
	fatiguedSpeed_Low := r.FormValue("fatiguedSpeed_Low")
	//fatiguedSpeed_Medium := r.FormValue("fatiguedSpeed_Medium")
	//fatiguedSpeed_Serious := r.FormValue("fatiguedSpeed_Serious")
	fatiguedTime_Low := r.FormValue("fatiguedTime_Low")
	//fatiguedTime_Medium := r.FormValue("fatiguedTime_Medium")
	//fatiguedTime_Serious := r.FormValue("fatiguedTime_Serious")
	fatiguedCD_Low := r.FormValue("fatiguedCD_Low")
	//fatiguedCD_Medium := r.FormValue("fatiguedCD_Medium")
	//fatiguedCD_Serious := r.FormValue("fatiguedCD_Serious")
	fatiguedAlarmSoundSw := r.FormValue("fatiguedAlarmSoundSw")
	fatiguedUploadingPicSw := r.FormValue("fatiguedUploadingPicSw")
	fatiguedUploadingVideoSw := r.FormValue("fatiguedUploadingVideoSw")
	fatiguedTTS := r.FormValue("fatiguedTTS")
	fatiguedAlarmSw := r.FormValue("fatiguedAlarmSw")
	//打电话报警
	callingSpeed := r.FormValue("callingSpeed")
	callingTime := r.FormValue("callingTime")
	callingCD := r.FormValue("callingCD")
	callingAlarmSoundSw := r.FormValue("callingAlarmSoundSw")
	callingUploadingPicSw := r.FormValue("callingUploadingPicSw")
	callingUploadingVideoSw := r.FormValue("callingUploadingVideoSw")
	callingTTS := r.FormValue("callingTTS")
	callingAlarmSw := r.FormValue("callingAlarmSw")
	//左顾右盼
	lookaoundSpeed := r.FormValue("lookaoundSpeed")
	lookaoundTime := r.FormValue("lookaoundTime")
	lookaoundCD := r.FormValue("lookaoundCD")
	lookaoundAlarmSoundSw := r.FormValue("lookaoundAlarmSoundSw")
	lookaoundUploadingPicSw := r.FormValue("lookaoundUploadingPicSw")
	lookaoundUploadingVideoSw := r.FormValue("lookaoundUploadingVideoSw")
	lookaoundTTS := r.FormValue("lookaoundTTS")
	lookaoundAlarmSw := r.FormValue("lookaoundAlarmSw")
	//抽烟
	smokingSpeed := r.FormValue("smokingSpeed")
	smokingTime := r.FormValue("smokingTime")
	smokingCD := r.FormValue("smokingCD")
	smokingAlarmSoundSw := r.FormValue("smokingAlarmSoundSw")
	smokingUploadingPicSw := r.FormValue("smokingUploadingPicSw")
	smokingUploadingVideoSw := r.FormValue("smokingUploadingVideoSw")
	smokingTTS := r.FormValue("smokingTTS")
	smokingAlarmSw := r.FormValue("smokingAlarmSw")
	//打哈欠
	yawnSpeed := r.FormValue("yawnSpeed")
	yawnCD := r.FormValue("yawnCD")
	yawnAlarmSoundSw := r.FormValue("yawnAlarmSoundSw")
	yawnUploadingPicSw := r.FormValue("yawnUploadingPicSw")
	yawnUploadingVideoSw := r.FormValue("yawnUploadingVideoSw")
	yawnTTS := r.FormValue("yawnTTS")
	yawnAlarmSw := r.FormValue("yawnAlarmSw")
	//离岗
	noFaceSpeed := r.FormValue("noFaceSpeed")
	noFaceTime := r.FormValue("noFaceTime")
	noFaceCD := r.FormValue("noFaceCD")
	noFaceAlarmSoundSw := r.FormValue("noFaceAlarmSoundSw")
	noFaceUploadingPicSw := r.FormValue("noFaceUploadingPicSw")
	noFaceUploadingVideoSw := r.FormValue("noFaceUploadingVideoSw")
	noFaceTTS := r.FormValue("noFaceTTS")
	noFaceAlarmSw := r.FormValue("noFaceAlarmSw")
	//吃东西/闲聊
	mouthSpeed := r.FormValue("mouthSpeed")
	mouthTime := r.FormValue("mouthTime")
	mouthCD := r.FormValue("mouthCD")
	mouthAlarmSoundSw := r.FormValue("mouthAlarmSoundSw")
	mouthUploadingPicSw := r.FormValue("mouthUploadingPicSw")
	mouthUploadingVideoSw := r.FormValue("mouthUploadingVideoSw")
	mouthTTS := r.FormValue("mouthTTS")
	mouthAlarmSw := r.FormValue("mouthAlarmSw")
	//摄像头遮挡
	occlusionSpeed := r.FormValue("occlusionSpeed")
	occlusionTime := r.FormValue("occlusionTime")
	occlusionCD := r.FormValue("occlusionCD")
	occlusionAlarmSoundSw := r.FormValue("occlusionAlarmSoundSw")
	occlusionUploadingPicSw := r.FormValue("occlusionUploadingPicSw")
	occlusionUploadingVideoSw := r.FormValue("occlusionUploadingVideoSw")
	occlusionTTS := r.FormValue("occlusionTTS")
	occlusionAlarmSw := r.FormValue("occlusionAlarmSw")
	//手比前方
	handSpeed := r.FormValue("handSpeed")
	handTime := r.FormValue("handTime")
	handCD := r.FormValue("handCD")
	handAlarmSoundSw := r.FormValue("handAlarmSoundSw")
	handUploadingPicSw := r.FormValue("handUploadingPicSw")
	handUploadingVideoSw := r.FormValue("handUploadingVideoSw")
	handTTS := r.FormValue("handTTS")
	handAlarmSw := r.FormValue("handAlarmSw")
	//摇臂
	armSpeed := r.FormValue("armSpeed")
	armTime := r.FormValue("armTime")
	armCD := r.FormValue("armCD")
	armAlarmSoundSw := r.FormValue("armAlarmSoundSw")
	armUploadingPicSw := r.FormValue("armUploadingPicSw")
	armUploadingVideoSw := r.FormValue("armUploadingVideoSw")
	armTTS := r.FormValue("armTTS")
	armAlarmSw := r.FormValue("armAlarmSw")
	//红外阻断
	blockingSpeed := r.FormValue("blockingSpeed")
	blockingTime := r.FormValue("blockingTime")
	blockingCD := r.FormValue("blockingCD")
	blockingTTS := r.FormValue("blockingTTS")
	blockingAlarmSw := r.FormValue("blockingAlarmSw")
	var fatiguedrivingalarmList []*viewmodel.FatiguedrivingalarmViewModel
	//疲劳驾驶报警
	fatiguedrivingalarm1 := &viewmodel.FatiguedrivingalarmViewModel{
		Id:               strconv.FormatInt(GenerateSnowFlakeId(16), 10),
		AlarmTypeName:    "疲劳驾驶报警",
		Sort:             1,
		Speed:            util.Tofloat64(fatiguedSpeed_Low),
		Time:             util.Tofloat64(fatiguedTime_Low),
		TTS:              fatiguedTTS,
		CD:               util.Tofloat64(fatiguedCD_Low),
		AlarmSw:          util.ToIntValue(fatiguedAlarmSw),
		AlarmSoundSw:     util.ToIntValue(fatiguedAlarmSoundSw),
		UploadingPicSw:   util.ToIntValue(fatiguedUploadingPicSw),
		UploadingVideoSw: util.ToIntValue(fatiguedUploadingVideoSw),
	}
	fatiguedrivingalarmList = append(fatiguedrivingalarmList, fatiguedrivingalarm1)
	//打电话报警
	fatiguedrivingalarm2 := &viewmodel.FatiguedrivingalarmViewModel{
		Id:               strconv.FormatInt(GenerateSnowFlakeId(17), 10),
		AlarmTypeName:    "打电话报警",
		Sort:             2,
		Speed:            util.Tofloat64(callingSpeed),
		Time:             util.Tofloat64(callingTime),
		TTS:              callingTTS,
		CD:               util.Tofloat64(callingCD),
		AlarmSw:          util.ToIntValue(callingAlarmSw),
		AlarmSoundSw:     util.ToIntValue(callingAlarmSoundSw),
		UploadingPicSw:   util.ToIntValue(callingUploadingPicSw),
		UploadingVideoSw: util.ToIntValue(callingUploadingVideoSw),
	}
	fatiguedrivingalarmList = append(fatiguedrivingalarmList, fatiguedrivingalarm2)
	//左顾右盼
	fatiguedrivingalarm3 := &viewmodel.FatiguedrivingalarmViewModel{
		Id:               strconv.FormatInt(GenerateSnowFlakeId(18), 10),
		AlarmTypeName:    "左顾右盼",
		Sort:             3,
		Speed:            util.Tofloat64(lookaoundSpeed),
		Time:             util.Tofloat64(lookaoundTime),
		TTS:              lookaoundTTS,
		CD:               util.Tofloat64(lookaoundCD),
		AlarmSw:          util.ToIntValue(lookaoundAlarmSw),
		AlarmSoundSw:     util.ToIntValue(lookaoundAlarmSoundSw),
		UploadingPicSw:   util.ToIntValue(lookaoundUploadingPicSw),
		UploadingVideoSw: util.ToIntValue(lookaoundUploadingVideoSw),
	}
	fatiguedrivingalarmList = append(fatiguedrivingalarmList, fatiguedrivingalarm3)
	//抽烟
	fatiguedrivingalarm4 := &viewmodel.FatiguedrivingalarmViewModel{
		Id:               strconv.FormatInt(GenerateSnowFlakeId(19), 10),
		AlarmTypeName:    "抽烟",
		Sort:             4,
		Speed:            util.Tofloat64(smokingSpeed),
		Time:             util.Tofloat64(smokingTime),
		TTS:              smokingTTS,
		CD:               util.Tofloat64(smokingCD),
		AlarmSw:          util.ToIntValue(smokingAlarmSw),
		AlarmSoundSw:     util.ToIntValue(smokingAlarmSoundSw),
		UploadingPicSw:   util.ToIntValue(smokingUploadingPicSw),
		UploadingVideoSw: util.ToIntValue(smokingUploadingVideoSw),
	}
	fatiguedrivingalarmList = append(fatiguedrivingalarmList, fatiguedrivingalarm4)
	//打哈欠
	fatiguedrivingalarm5 := &viewmodel.FatiguedrivingalarmViewModel{
		Id:               strconv.FormatInt(GenerateSnowFlakeId(20), 10),
		AlarmTypeName:    "打哈欠",
		Sort:             5,
		Speed:            util.Tofloat64(yawnSpeed),
		Time:             0,
		TTS:              yawnTTS,
		CD:               util.Tofloat64(yawnCD),
		AlarmSw:          util.ToIntValue(yawnAlarmSw),
		AlarmSoundSw:     util.ToIntValue(yawnAlarmSoundSw),
		UploadingPicSw:   util.ToIntValue(yawnUploadingPicSw),
		UploadingVideoSw: util.ToIntValue(yawnUploadingVideoSw),
	}
	fatiguedrivingalarmList = append(fatiguedrivingalarmList, fatiguedrivingalarm5)
	//离岗
	fatiguedrivingalarm6 := &viewmodel.FatiguedrivingalarmViewModel{
		Id:               strconv.FormatInt(GenerateSnowFlakeId(21), 10),
		AlarmTypeName:    "离岗",
		Sort:             6,
		Speed:            util.Tofloat64(noFaceSpeed),
		Time:             util.Tofloat64(noFaceTime),
		TTS:              noFaceTTS,
		CD:               util.Tofloat64(noFaceCD),
		AlarmSw:          util.ToIntValue(noFaceAlarmSw),
		AlarmSoundSw:     util.ToIntValue(noFaceAlarmSoundSw),
		UploadingPicSw:   util.ToIntValue(noFaceUploadingPicSw),
		UploadingVideoSw: util.ToIntValue(noFaceUploadingVideoSw),
	}
	fatiguedrivingalarmList = append(fatiguedrivingalarmList, fatiguedrivingalarm6)
	//吃东西/闲聊
	fatiguedrivingalarm7 := &viewmodel.FatiguedrivingalarmViewModel{
		Id:               strconv.FormatInt(GenerateSnowFlakeId(22), 10),
		AlarmTypeName:    "吃东西/闲聊",
		Sort:             7,
		Speed:            util.Tofloat64(mouthSpeed),
		Time:             util.Tofloat64(mouthTime),
		TTS:              mouthTTS,
		CD:               util.Tofloat64(mouthCD),
		AlarmSw:          util.ToIntValue(mouthAlarmSw),
		AlarmSoundSw:     util.ToIntValue(mouthAlarmSoundSw),
		UploadingPicSw:   util.ToIntValue(mouthUploadingPicSw),
		UploadingVideoSw: util.ToIntValue(mouthUploadingVideoSw),
	}
	fatiguedrivingalarmList = append(fatiguedrivingalarmList, fatiguedrivingalarm7)
	//摄像头遮挡
	fatiguedrivingalarm8 := &viewmodel.FatiguedrivingalarmViewModel{
		Id:               strconv.FormatInt(GenerateSnowFlakeId(23), 10),
		AlarmTypeName:    "摄像头遮挡",
		Sort:             8,
		Speed:            util.Tofloat64(occlusionSpeed),
		Time:             util.Tofloat64(occlusionTime),
		TTS:              occlusionTTS,
		CD:               util.Tofloat64(occlusionCD),
		AlarmSw:          util.ToIntValue(occlusionAlarmSw),
		AlarmSoundSw:     util.ToIntValue(occlusionAlarmSoundSw),
		UploadingPicSw:   util.ToIntValue(occlusionUploadingPicSw),
		UploadingVideoSw: util.ToIntValue(occlusionUploadingVideoSw),
	}
	fatiguedrivingalarmList = append(fatiguedrivingalarmList, fatiguedrivingalarm8)
	//手比前方
	fatiguedrivingalarm9 := &viewmodel.FatiguedrivingalarmViewModel{
		Id:               strconv.FormatInt(GenerateSnowFlakeId(24), 10),
		AlarmTypeName:    "手比前方",
		Sort:             9,
		Speed:            util.Tofloat64(handSpeed),
		Time:             util.Tofloat64(handTime),
		TTS:              handTTS,
		CD:               util.Tofloat64(handCD),
		AlarmSw:          util.ToIntValue(handAlarmSw),
		AlarmSoundSw:     util.ToIntValue(handAlarmSoundSw),
		UploadingPicSw:   util.ToIntValue(handUploadingPicSw),
		UploadingVideoSw: util.ToIntValue(handUploadingVideoSw),
	}
	fatiguedrivingalarmList = append(fatiguedrivingalarmList, fatiguedrivingalarm9)
	//摇臂
	fatiguedrivingalarm10 := &viewmodel.FatiguedrivingalarmViewModel{
		Id:               strconv.FormatInt(GenerateSnowFlakeId(25), 10),
		AlarmTypeName:    "摇臂",
		Sort:             10,
		Speed:            util.Tofloat64(armSpeed),
		Time:             util.Tofloat64(armTime),
		TTS:              armTTS,
		CD:               util.Tofloat64(armCD),
		AlarmSw:          util.ToIntValue(armAlarmSw),
		AlarmSoundSw:     util.ToIntValue(armAlarmSoundSw),
		UploadingPicSw:   util.ToIntValue(armUploadingPicSw),
		UploadingVideoSw: util.ToIntValue(armUploadingVideoSw),
	}
	fatiguedrivingalarmList = append(fatiguedrivingalarmList, fatiguedrivingalarm10)
	//红外阻断
	fatiguedrivingalarm11 := &viewmodel.FatiguedrivingalarmViewModel{
		Id:               strconv.FormatInt(GenerateSnowFlakeId(26), 10),
		AlarmTypeName:    "红外阻断",
		Sort:             11,
		Speed:            util.Tofloat64(blockingSpeed),
		Time:             util.Tofloat64(blockingTime),
		TTS:              blockingTTS,
		CD:               util.Tofloat64(blockingCD),
		AlarmSw:          util.ToIntValue(blockingAlarmSw),
		AlarmSoundSw:     0,
		UploadingPicSw:   0,
		UploadingVideoSw: 0,
	}
	fatiguedrivingalarmList = append(fatiguedrivingalarmList, fatiguedrivingalarm11)
	//行驶记录
	ch_Name1 := r.FormValue("ch_Name1")
	ch_Name2 := r.FormValue("ch_Name2")
	isAlarmDvr1 := r.FormValue("isAlarmDvr1")
	isAlarmDvr2 := r.FormValue("isAlarmDvr2")
	var drivingrecordList []*viewmodel.DrivingrecordViewModel
	drivingrecord1 := &viewmodel.DrivingrecordViewModel{
		Id:     strconv.FormatInt(GenerateSnowFlakeId(27), 10),
		Name:   ch_Name1,
		IsOpen: util.ToIntValue(isAlarmDvr1),
	}
	drivingrecordList = append(drivingrecordList, drivingrecord1)
	drivingrecord2 := &viewmodel.DrivingrecordViewModel{
		Id:     strconv.FormatInt(GenerateSnowFlakeId(28), 10),
		Name:   ch_Name2,
		IsOpen: util.ToIntValue(isAlarmDvr2),
	}
	drivingrecordList = append(drivingrecordList, drivingrecord2)
	msg := new(domain.Result)
	msg.Init()
	isSucc, err := service.SaveAlarmSettings(fatiguedrivingalarmList, drivingrecordList)
	if err != nil {
		log.Printf("提交报警设置信息 error: %v", err.Error())
	} else {
		if isSucc {
			msg.Code = 0
			msg.Msg = "提交配置信息成功！"
		} else {
			msg.Code = 1
			msg.Msg = "提交配置信息失败！"
			log.Printf("提交配置信息 error: %v", err.Error())
		}
	}
	util.ResponseHtml(w, msg, err)
}

//2.提交账户管理信息
func submitAccountmanager(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("提交账户管理信息 error: %v", err.Error())
	}
	ftpUser := r.FormValue("ftpUser")
	ftpPwd := r.FormValue("ftpPwd")
	ocsFtpUser := r.FormValue("ocsFtpUser")
	ocsFtpPwd := r.FormValue("ocsFtpPwd")
	rtspUser := r.FormValue("rtspUser")
	rtspPwd := r.FormValue("rtspPwd")

	//date := r.FormValue("date")
	//_time, _ := time.Parse("2006-01-02 15:04:05", date)
	//_time := time.Now().Format("2006-01-02 15:04:05")
	var accountmanagerList []*viewmodel.AccountmanagerViewModel

	accountmanager1 := &viewmodel.AccountmanagerViewModel{
		Id:          strconv.FormatInt(GenerateSnowFlakeId(10), 10),
		AccountType: 1,
		User:        ftpUser,
		Pwd:         ftpPwd,
	}
	accountmanagerList = append(accountmanagerList, accountmanager1)
	accountmanager2 := &viewmodel.AccountmanagerViewModel{
		Id:          strconv.FormatInt(GenerateSnowFlakeId(11), 10),
		AccountType: 2,
		User:        ocsFtpUser,
		Pwd:         ocsFtpPwd,
	}
	accountmanagerList = append(accountmanagerList, accountmanager2)
	accountmanager3 := &viewmodel.AccountmanagerViewModel{
		Id:          strconv.FormatInt(GenerateSnowFlakeId(12), 10),
		AccountType: 3,
		User:        rtspUser,
		Pwd:         rtspPwd,
	}
	accountmanagerList = append(accountmanagerList, accountmanager3)
	msg := new(domain.Result)
	isSucc, err := service.BatchInsertAccountmanager(0, accountmanagerList)
	if err != nil {
		log.Printf("提交账户管理信息 error: %v", err.Error())
	} else {
		if isSucc {
			msg.Code = 0
			msg.Msg = "添加成功！"
			//添加成功后移除token
			errr := RemoveSession(w, r, config.AddTokenKey)
			if errr != nil {
				log.Printf("移除Session error: %v", errr.Error())
			}
		} else {
			msg.Code = 1
			msg.Msg = "添加失败！"
			log.Printf("添加账户管理设置信息 error: %v", err.Error())
		}
	}
	util.ResponseHtml(w, msg, err)
}

//3.提交通讯地址设置信息
func submitCommunicationsettings(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("提交通讯地址设置信息 error: %v", err.Error())
	}
	localIP := r.FormValue("localIP")
	localPort := r.FormValue("localPort")
	ocsIP := r.FormValue("ocsIP")
	ocsPort := r.FormValue("ocsPort")

	var communicationsettingsList []*viewmodel.CommunicationsettingsViewModel

	communicationsettings1 := &viewmodel.CommunicationsettingsViewModel{
		Id:                strconv.FormatInt(GenerateSnowFlakeId(13), 10),
		CommunicationType: 1,
		IP:                localIP,
		Port:              localPort,
	}
	communicationsettingsList = append(communicationsettingsList, communicationsettings1)
	communicationsettings2 := &viewmodel.CommunicationsettingsViewModel{
		Id:                strconv.FormatInt(GenerateSnowFlakeId(15), 10),
		CommunicationType: 2,
		IP:                ocsIP,
		Port:              ocsPort,
	}
	communicationsettingsList = append(communicationsettingsList, communicationsettings2)

	msg := new(domain.Result)
	isSucc, err := service.BatchInsertCommunicationsettings(0, communicationsettingsList)
	if err != nil {
		log.Printf("提交通讯地址设置信息 error: %v", err.Error())
	} else {
		if isSucc {
			msg.Code = 0
			msg.Msg = "提交成功！"
		} else {
			msg.Code = 1
			msg.Msg = "提交失败！"
			log.Printf("添加用户信息 error: %v", err.Error())
		}
	}
	util.ResponseHtml(w, msg, err)
}

//4.提交系统设置信息
func submitSystemSettings(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("提交系统设置信息 error: %v", err.Error())
	}
	locomotiveID := r.FormValue("locomotiveID")
	hostSide := r.FormValue("hostSide")
	cam1Open := r.FormValue("cam1Open")
	channelID0 := r.FormValue("channelID0")
	channelName0 := r.FormValue("channelName0")
	cam2Open := r.FormValue("cam2Open")
	channelID1 := r.FormValue("channelID1")
	channelName1 := r.FormValue("channelName1")
	cam3Open := r.FormValue("cam3Open")
	channelID2 := r.FormValue("channelID2")
	channelName2 := r.FormValue("channelName2")
	gps_isOpen := r.FormValue("gps_isOpen")
	wifi_isOpen := r.FormValue("wifi_isOpen")
	G_isOpen := r.FormValue("4G_isOpen")
	soundType := r.FormValue("soundType")
	volume := r.FormValue("volume")

	_hostSide, _ := strconv.Atoi(hostSide)
	carsettingViewModel := &viewmodel.CarsettingViewModel{
		Id:           strconv.FormatInt(GenerateSnowFlakeId(6), 10),
		LocomotiveID: locomotiveID,
		HostSide:     _hostSide,
	}

	var camerasettingList []*viewmodel.CamerasettingViewModel
	_ChannelNo1, _ := strconv.Atoi(channelID0)
	camerasetting1 := &viewmodel.CamerasettingViewModel{
		Id:           strconv.FormatInt(GenerateSnowFlakeId(7), 10),
		Camera_Name:  "摄像头1",
		Camera_OnOff: util.ToIntValue(cam1Open),
		Camera_IP:    "",
		ChannelNo:    _ChannelNo1,
		ChannelName:  channelName0,
	}
	camerasettingList = append(camerasettingList, camerasetting1)
	_ChannelNo2, _ := strconv.Atoi(channelID1)
	camerasetting2 := &viewmodel.CamerasettingViewModel{
		Id:           strconv.FormatInt(GenerateSnowFlakeId(8), 10),
		Camera_Name:  "摄像头2",
		Camera_OnOff: util.ToIntValue(cam2Open),
		Camera_IP:    "",
		ChannelNo:    _ChannelNo2,
		ChannelName:  channelName1,
	}
	camerasettingList = append(camerasettingList, camerasetting2)
	_ChannelNo3, _ := strconv.Atoi(channelID2)
	camerasetting3 := &viewmodel.CamerasettingViewModel{
		Id:           strconv.FormatInt(GenerateSnowFlakeId(9), 10),
		Camera_Name:  "摄像头3",
		Camera_OnOff: util.ToIntValue(cam3Open),
		Camera_IP:    "",
		ChannelNo:    _ChannelNo3,
		ChannelName:  channelName2,
	}
	camerasettingList = append(camerasettingList, camerasetting3)

	var peripheralsettingsList []*viewmodel.PeripheralsettingsViewModel
	peripheralsettings1 := &viewmodel.PeripheralsettingsViewModel{
		Id:               strconv.FormatInt(GenerateSnowFlakeId(2), 10),
		PeripheralType:   1,
		Peripheral_OnOff: util.ToIntValue(gps_isOpen),
	}
	peripheralsettingsList = append(peripheralsettingsList, peripheralsettings1)
	peripheralsettings2 := &viewmodel.PeripheralsettingsViewModel{
		Id:               strconv.FormatInt(GenerateSnowFlakeId(3), 10),
		PeripheralType:   2,
		Peripheral_OnOff: util.ToIntValue(wifi_isOpen),
	}
	peripheralsettingsList = append(peripheralsettingsList, peripheralsettings2)
	peripheralsettings3 := &viewmodel.PeripheralsettingsViewModel{
		Id:               strconv.FormatInt(GenerateSnowFlakeId(4), 10),
		PeripheralType:   4,
		Peripheral_OnOff: util.ToIntValue(G_isOpen),
	}
	peripheralsettingsList = append(peripheralsettingsList, peripheralsettings3)
	_soundType, _ := strconv.Atoi(soundType)
	_volume, _ := strconv.Atoi(volume)
	soundsettingViewModel := &viewmodel.SoundsettingViewModel{
		Id:          strconv.FormatInt(GenerateSnowFlakeId(5), 10),
		SoundType:   _soundType,
		Volume:      _volume,
		CustomSound: "0",
	}
	msg := new(domain.Result)
	isSucc, err := service.SaveSystemSettings(
		carsettingViewModel,
		camerasettingList,
		peripheralsettingsList,
		soundsettingViewModel)
	if err != nil {
		log.Printf("提交系统设置信息 error: %v", err.Error())
	} else {
		if isSucc {
			msg.Code = 0
			msg.Msg = "提交配置信息成功！"
		} else {
			msg.Code = 1
			msg.Msg = "提交配置信息失败！"
			log.Printf("提交配置信息 error: %v", err.Error())
		}
	}
	util.ResponseHtml(w, msg, err)
}
