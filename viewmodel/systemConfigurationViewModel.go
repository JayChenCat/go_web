package viewmodel

//获取配置信息对象
type SystemConfigurationViewModel struct {
	AlarmSetting         *AlarmSetting                     `json:"AlarmSetting"`
	AccountSetting       []*AccountmanagerViewModel        `json:"AccountSetting"`
	CommunicationSetting []*CommunicationsettingsViewModel `json:"CommunicationSetting"`
	SystemSetting        *SystemSetting                    `json:"SystemSetting"`
}

type AlarmSetting struct {
	FatiguedAlarmSetting  *FatiguedrivingalarmViewModel `json:"FatiguedAlarmSetting"`
	CallingAlarmSetting   *FatiguedrivingalarmViewModel `json:"CallingAlarmSetting"`
	LookaoundAlarmSetting *FatiguedrivingalarmViewModel `json:"LookaoundAlarmSetting"`
	SmokingAlarmSetting   *FatiguedrivingalarmViewModel `json:"SmokingAlarmSetting"`
	YawnAlarmSetting      *FatiguedrivingalarmViewModel `json:"YawnAlarmSetting"`
	NoFaceAlarmSetting    *FatiguedrivingalarmViewModel `json:"NoFaceAlarmSetting"`
	MouthAlarmSetting     *FatiguedrivingalarmViewModel `json:"MouthAlarmSetting"`
	OcclusionAlarmSetting *FatiguedrivingalarmViewModel `json:"OcclusionAlarmSetting"`
	HandAlarmSetting      *FatiguedrivingalarmViewModel `json:"HandAlarmSetting"`
	ArmAlarmSetting       *FatiguedrivingalarmViewModel `json:"ArmAlarmSetting"`
	BlockingAlarmSetting  *FatiguedrivingalarmViewModel `json:"BlockingAlarmSetting"`
	Drivingrecord         []*DrivingrecordViewModel     `json:"DrivingrecordSetting"`
}

type SystemSetting struct {
	CameraSetting      []*CamerasettingViewModel      `json:"cameraSetting"`
	Carsetting         *CarsettingViewModel           `json:"carsetting"`
	Peripheralsettings []*PeripheralsettingsViewModel `json:"peripheralsettings"`
	Soundsetting       *SoundsettingViewModel         `json:"soundsetting"`
}
