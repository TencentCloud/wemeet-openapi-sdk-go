// Package wemeetopenapi is auto generate
/*
   腾讯会议OpenAPI

   SAAS版RESTFUL风格API

   API version: v1.0.3
*/
package wemeetopenapi

// V1DevicesGet200Response struct for V1DevicesGet200Response
type V1DevicesGet200Response struct {
	// 当前页
	CurrentPage *int64 `json:"current_page,omitempty"`
	// 单页条数
	CurrentSize *int64 `json:"current_size,omitempty"`
	// 设备对象列表
	DeviceInfoList []V1DevicesGet200ResponseDeviceInfoListInner `json:"device_info_list,omitempty"`
	// 数据总条数
	TotalCount *int64 `json:"total_count,omitempty"`
	// 数据总页数
	TotalPage *int64 `json:"total_page,omitempty"`
}

// V1DevicesGet200ResponseDeviceInfoListInner struct for V1DevicesGet200ResponseDeviceInfoListInner
type V1DevicesGet200ResponseDeviceInfoListInner struct {
	// 应用程序版本
	AppVersion *string `json:"app_version,omitempty"`
	// 设备型号
	DeviceModel       *string                                                      `json:"device_model,omitempty"`
	DeviceMonitorInfo *V1DevicesGet200ResponseDeviceInfoListInnerDeviceMonitorInfo `json:"device_monitor_info,omitempty"`
	// 会议室ID
	MeetingRoomId *string `json:"meeting_room_id,omitempty"`
	// 会议室地址
	MeetingRoomLocation *string `json:"meeting_room_location,omitempty"`
	// 会议室名称
	MeetingRoomName *string `json:"meeting_room_name,omitempty"`
	// 会议室状态：0：未激活1：未绑定2：空闲3：使用中4：离线，5:未登录
	MeetingRoomStatus *int64  `json:"meeting_room_status,omitempty"`
	RoomsId           *string `json:"rooms_id,omitempty"`
}

// V1DevicesGet200ResponseDeviceInfoListInnerDeviceMonitorInfo struct for V1DevicesGet200ResponseDeviceInfoListInnerDeviceMonitorInfo
type V1DevicesGet200ResponseDeviceInfoListInnerDeviceMonitorInfo struct {
	// 摄像头健康状态
	CameraStatus *bool `json:"camera_status,omitempty"`
	// 麦克风健康状态
	MicrophoneStatus *bool `json:"microphone_status,omitempty"`
	// 扬声器健康状态
	SpeakerStatus *bool `json:"speaker_status,omitempty"`
}

// V1MeetingRoomsCancelRoomCallPutRequest struct for V1MeetingRoomsCancelRoomCallPutRequest
type V1MeetingRoomsCancelRoomCallPutRequest struct {
	InstanceId int64 `json:"instance_id"`
	// 呼叫ID
	InviteId string `json:"invite_id"`
	// 会议ID
	MeetingId string `json:"meeting_id"`
	// 会议室 ID，与 mra_address 二选一。
	MeetingRoomId *string                                           `json:"meeting_room_id,omitempty"`
	MraAddress    *V1MeetingRoomsCancelRoomCallPutRequestMraAddress `json:"mra_address,omitempty"`
	OperatorId    string                                            `json:"operator_id"`
	// 操作者 ID 的类型： 1：userid
	OperatorIdType int64 `json:"operator_id_type"`
}

// V1MeetingRoomsCancelRoomCallPutRequestMraAddress MRA 对象
type V1MeetingRoomsCancelRoomCallPutRequestMraAddress struct {
	// 信令地址。 如果是 H.323 类型，输入 IP 地址或 E.164 号码。 如果是 SIP 类型，输入 IP 地址或 URI。
	DialString string `json:"dial_string"`
	// 信令协议。 1：SIP 2：H.323
	Protocol int64 `json:"protocol"`
}

// V1MeetingRoomsGet200Response struct for V1MeetingRoomsGet200Response
type V1MeetingRoomsGet200Response struct {
	// 当前页
	CurrentPage *int64 `json:"current_page,omitempty"`
	// 单页条数
	CurrentSize *int64 `json:"current_size,omitempty"`
	// 会议室对象列表
	MeetingRoomList []V1MeetingRoomsGet200ResponseMeetingRoomListInner `json:"meeting_room_list,omitempty"`
	// 数据总条数
	TotalCount *int64 `json:"total_count,omitempty"`
	// 数据总页数
	TotalPage *int64 `json:"total_page,omitempty"`
}

// V1MeetingRoomsGet200ResponseMeetingRoomListInner struct for V1MeetingRoomsGet200ResponseMeetingRoomListInner
type V1MeetingRoomsGet200ResponseMeetingRoomListInner struct {
	// 0:基础版账号 1:专业版账号
	AccountNewType *int64 `json:"account_new_type,omitempty"`
	// 账号类型0:普通，1:专款，2:试用
	AccountType *int64 `json:"account_type,omitempty"`
	// 激活码
	ActiveCode *string `json:"active_code,omitempty"`
	// 是否允许被呼叫
	IsAllowCall *bool `json:"is_allow_call,omitempty"`
	// 会议室ID
	MeetingRoomId *string `json:"meeting_room_id,omitempty"`
	// 会议室地址
	MeetingRoomLocation *string `json:"meeting_room_location,omitempty"`
	// 会议室名称
	MeetingRoomName *string `json:"meeting_room_name,omitempty"`
	// 会议室状态0:未激活，1:未绑定，2:空闲，3:试用中，4:离线，5:未登录
	MeetingRoomStatus *int64 `json:"meeting_room_status,omitempty"`
	// 容纳人数
	ParticipantNumber *int64 `json:"participant_number,omitempty"`
	// 预定状态
	ScheduledStatus *int64 `json:"scheduled_status,omitempty"`
}

// V1MeetingRoomsMeetingRoomIdActiveCodePost200Response struct for V1MeetingRoomsMeetingRoomIdActiveCodePost200Response
type V1MeetingRoomsMeetingRoomIdActiveCodePost200Response struct {
	// 激活码
	ActiveCode *string `json:"active_code,omitempty"`
}

// V1MeetingRoomsMeetingRoomIdBackgroundGet200Response struct for V1MeetingRoomsMeetingRoomIdBackgroundGet200Response
type V1MeetingRoomsMeetingRoomIdBackgroundGet200Response struct {
	// 背景图片地址
	BackgroundImage *string `json:"background_image,omitempty"`
}

// V1MeetingRoomsMeetingRoomIdBackgroundPost200Response struct for V1MeetingRoomsMeetingRoomIdBackgroundPost200Response
type V1MeetingRoomsMeetingRoomIdBackgroundPost200Response struct {
	JobId *string `json:"job_id,omitempty"`
}

// V1MeetingRoomsMeetingRoomIdBackgroundPostRequest struct for V1MeetingRoomsMeetingRoomIdBackgroundPostRequest
type V1MeetingRoomsMeetingRoomIdBackgroundPostRequest struct {
	// 不传或者传空则设置为默认背景，目前只能设置一张 背景图片地址，1920*1080,大小10M以内，png/jpg/jpeg格式
	BackgroundImage *string `json:"background_image,omitempty"`
	OperatorId      string  `json:"operator_id"`
	// 1:userid
	OperatorIdType int64 `json:"operator_id_type"`
}

// V1MeetingRoomsMeetingRoomIdGet200Response struct for V1MeetingRoomsMeetingRoomIdGet200Response
type V1MeetingRoomsMeetingRoomIdGet200Response struct {
	AccountInfo  *V1MeetingRoomsMeetingRoomIdGet200ResponseAccountInfo  `json:"account_info,omitempty"`
	BasicInfo    *V1MeetingRoomsMeetingRoomIdGet200ResponseBasicInfo    `json:"basic_info,omitempty"`
	HardwareInfo *V1MeetingRoomsMeetingRoomIdGet200ResponseHardwareInfo `json:"hardware_info,omitempty"`
	// 是否允许被呼叫，true：是 false：否
	IsAllowCall *bool `json:"is_allow_call,omitempty"`
	// 告警通知状态，0：未开启 1：已开启
	MonitorStatus *int64                                            `json:"monitor_status,omitempty"`
	PmiInfo       *V1MeetingRoomsMeetingRoomIdGet200ResponsePmiInfo `json:"pmi_info,omitempty"`
	// 预定状态： 0：未开放预定 1：开放预定
	ScheduledStatus *int64 `json:"scheduled_status,omitempty"`
}

// V1MeetingRoomsMeetingRoomIdGet200ResponseAccountInfo 会议室账号信息
type V1MeetingRoomsMeetingRoomIdGet200ResponseAccountInfo struct {
	AccountNewType *int64 `json:"account_new_type,omitempty"`
	// 账号类型，0：普通 1：专款 2：试用
	AccountType *int64 `json:"account_type,omitempty"`
	// 1-预装 2-体验 3-付费
	ProAccountType *int64 `json:"pro_account_type,omitempty"`
	// 有效期限
	ValidPeriod *string `json:"valid_period,omitempty"`
}

// V1MeetingRoomsMeetingRoomIdGet200ResponseBasicInfo 会议室基本信息
type V1MeetingRoomsMeetingRoomIdGet200ResponseBasicInfo struct {
	// 建筑
	Building *string `json:"building,omitempty"`
	// 城市
	City *string `json:"city,omitempty"`
	// 描述（base64）
	Desc *string `json:"desc,omitempty"`
	// 会议室设备
	Device *string `json:"device,omitempty"`
	// 楼层
	Floor *string `json:"floor,omitempty"`
	// 会议室名称
	MeetingRoomName *string `json:"meeting_room_name,omitempty"`
	// 容纳人数
	ParticipantNumber *int64 `json:"participant_number,omitempty"`
	// 管理员密码（base64）
	Password *string `json:"password,omitempty"`
}

// V1MeetingRoomsMeetingRoomIdGet200ResponseHardwareInfo 会议室硬件信息
type V1MeetingRoomsMeetingRoomIdGet200ResponseHardwareInfo struct {
	// 激活时间
	ActiveTime *string `json:"active_time,omitempty"`
	// 摄像头型号
	CameraModel *string `json:"camera_model,omitempty"`
	// CPU信息
	CpuInfo *string `json:"cpu_info,omitempty"`
	// CPU最大占用率
	CpuUsage *string `json:"cpu_usage,omitempty"`
	// 设备型号
	DeviceModel *string `json:"device_model,omitempty"`
	// 是否开启视频镜像
	EnableVideoMirror *bool `json:"enable_video_mirror,omitempty"`
	// 厂家
	Factory *string `json:"factory,omitempty"`
	// 固件版本
	FirmwareVersion *string `json:"firmware_version,omitempty"`
	// GPU信息
	GpuInfo *string `json:"gpu_info,omitempty"`
	// 健康状况
	HealthStatus *string `json:"health_status,omitempty"`
	// ip地址
	Ip *string `json:"ip,omitempty"`
	// MAC地址
	Mac *string `json:"mac,omitempty"`
	// 会议室状态
	MeetingRoomStatus *int64 `json:"meeting_room_status,omitempty"`
	// 内存信息
	MemoryInfo *string `json:"memory_info,omitempty"`
	// 麦克风信息
	MicrophoneInfo *string `json:"microphone_info,omitempty"`
	// 显示器刷新率
	MonitorFrequency *int64 `json:"monitor_frequency,omitempty"`
	// 网络类型
	NetType *string `json:"net_type,omitempty"`
	// Rooms版本
	RoomsVersion *string `json:"rooms_version,omitempty"`
	// 序列号
	Sn *string `json:"sn,omitempty"`
	// 扬声器信息
	SpeakerInfo *string `json:"speaker_info,omitempty"`
	// 设备系统
	SystemType *string `json:"system_type,omitempty"`
}

// V1MeetingRoomsMeetingRoomIdGet200ResponsePmiInfo 会议室 PMI 信息
type V1MeetingRoomsMeetingRoomIdGet200ResponsePmiInfo struct {
	// 设备会议号
	PmiCode *string `json:"pmi_code,omitempty"`
	// 设备入会密码
	PmiPwd *string `json:"pmi_pwd,omitempty"`
}

// V1MeetingRoomsModifyPutRequest struct for V1MeetingRoomsModifyPutRequest
type V1MeetingRoomsModifyPutRequest struct {
	InstanceId int64 `json:"instance_id"`
	// 是否允许被呼叫
	IsAllowCall *bool `json:"is_allow_call,omitempty"`
	// 要设置的会议室 ID。
	MeetingRoomId   string                                         `json:"meeting_room_id"`
	MeetingRoomInfo *V1MeetingRoomsModifyPutRequestMeetingRoomInfo `json:"meeting_room_info,omitempty"`
	// 操作者 ID。operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。
	OperatorId string `json:"operator_id"`
	// 操作者 ID 的类型： 1：userid
	OperatorIdType int64 `json:"operator_id_type"`
	// 是否开放预定初始值false
	ScheduledStatus *bool `json:"scheduled_status,omitempty"`
}

// V1MeetingRoomsModifyPutRequestMeetingRoomInfo 编辑会议室基本信息
type V1MeetingRoomsModifyPutRequestMeetingRoomInfo struct {
	// 建筑。若非输入城市下现有建筑则自动创建该建筑与楼层。长度不超过36个字符或18个汉字。
	Building *string `json:"building,omitempty"`
	// 城市。若非已添加城市则自动创建城市及建筑与楼层。长度不超过36个字符或18个汉字。city、building、floor 需同时传入或都不传入。
	City *string `json:"city,omitempty"`
	// 描述（base64）。长度不超过1000个字符。
	Desc *string `json:"desc,omitempty"`
	// 会议室设备，输入非现有类型内容则无效。设备类型有：TV，投影，会议电话机，MIC，视频电视，PC，无线投屏。
	Device []string `json:"device,omitempty"`
	// 楼层。若非输入建筑下现有楼层则自动创建楼层。输入应为数字或字母，长度不超过36个字符。
	Floor *string `json:"floor,omitempty"`
	// 会议室名称。长度不超过36个字符。
	MeetingRoomName string `json:"meeting_room_name"`
	// 会议室类型。 0：rooms 会议室 1：无类型会议室 2：SIP 会议室 4：H.323 会议室
	MeetingRoomType int64 `json:"meeting_room_type"`
	// 会议室信令地址。会议室类型为2或4时必填写，与mra_register_account 二选一。
	MraAddress *string `json:"mra_address,omitempty"`
	// SIP/ H.323注册账号。会议室类型为2或4时填写。
	MraRegisterAccount *string `json:"mra_register_account,omitempty"`
	// 容纳人数。不超过9位数。
	ParticipantNumber *int64 `json:"participant_number,omitempty"`
	// 使用管理员密码时必须填写管理员密码（base64）。若不使用密码，该字段无效。输入应为1-16位的数字、字母或字符。
	Password *string `json:"password,omitempty"`
	// 会议室类型为1时选择是否使用管理员密码，默认为 false。 true：使用 false：不使用
	UsePassword *bool `json:"use_password,omitempty"`
}

// V1MeetingRoomsModifyRoomConfigInfoPostRequest struct for V1MeetingRoomsModifyRoomConfigInfoPostRequest
type V1MeetingRoomsModifyRoomConfigInfoPostRequest struct {
	InstanceId int64 `json:"instance_id"`
	// 要配置的会议室 ID。
	MeetingRoomId   string                                                        `json:"meeting_room_id"`
	MeetingSettings *V1MeetingRoomsModifyRoomConfigInfoPostRequestMeetingSettings `json:"meeting_settings,omitempty"`
	OperatorId      string                                                        `json:"operator_id"`
	OperatorIdType  int64                                                         `json:"operator_id_type"`
	RecordSettings  *V1MeetingRoomsModifyRoomConfigInfoPostRequestRecordSettings  `json:"record_settings,omitempty"`
}

// V1MeetingRoomsModifyRoomConfigInfoPostRequestMeetingSettings struct for V1MeetingRoomsModifyRoomConfigInfoPostRequestMeetingSettings
type V1MeetingRoomsModifyRoomConfigInfoPostRequestMeetingSettings struct {
	// 自动接听，初始值为 false。 1：开启 2：关闭
	AutoResponse *int64 `json:"auto_response,omitempty"`
	// 字幕，初始值为 true。 true：开启 false：关闭
	Caption *bool `json:"caption,omitempty"`
	// Rooms 屏幕是否展示消息通知，初始值为 false。 true：开启 false：不开启
	RoomNotification *bool `json:"room_notification,omitempty"`
	// 专属 ID，初始值为 true。 true：开启 false：关闭
	RoomPmi         *bool                                                                        `json:"room_pmi,omitempty"`
	RoomPmiSettings *V1MeetingRoomsModifyRoomConfigInfoPostRequestMeetingSettingsRoomPmiSettings `json:"room_pmi_settings,omitempty"`
	// 水印，初始值为2。 0：关闭 1：单排水印 2：多排水印
	WaterMark *int64 `json:"water_mark,omitempty"`
}

// V1MeetingRoomsModifyRoomConfigInfoPostRequestMeetingSettingsRoomPmiSettings struct for V1MeetingRoomsModifyRoomConfigInfoPostRequestMeetingSettingsRoomPmiSettings
type V1MeetingRoomsModifyRoomConfigInfoPostRequestMeetingSettingsRoomPmiSettings struct {
	AllowInBeforeHost         *bool   `json:"allow_in_before_host,omitempty"`
	MuteEnableTypeJoin        *int64  `json:"mute_enable_type_join,omitempty"`
	OnlyEnterpriseUserAllowed *bool   `json:"only_enterprise_user_allowed,omitempty"`
	RoomPmiPsw                *string `json:"room_pmi_psw,omitempty"`
	WaitingRoom               *bool   `json:"waiting_room,omitempty"`
}

// V1MeetingRoomsModifyRoomConfigInfoPostRequestRecordSettings struct for V1MeetingRoomsModifyRoomConfigInfoPostRequestRecordSettings
type V1MeetingRoomsModifyRoomConfigInfoPostRequestRecordSettings struct {
	// 是否允许下载云录制，初始值为 false。 true：开启 false：关闭
	DownloadRecord *bool `json:"download_record,omitempty"`
	// 分享云录制，初始值为1。 0：关闭分享 1：全部成员可查看 2：仅登录成员可查看 3：仅同企业成员可查看 4：仅参会成员可查看
	ShareRecord *int64 `json:"share_record,omitempty"`
}

// V1MeetingRoomsMonitorDeviceControllerInfoGet200Response struct for V1MeetingRoomsMonitorDeviceControllerInfoGet200Response
type V1MeetingRoomsMonitorDeviceControllerInfoGet200Response struct {
	// 控制器信息对象
	ControllerInfoList []V1MeetingRoomsMonitorDeviceControllerInfoGet200ResponseControllerInfoListInner `json:"controller_info_list,omitempty"`
	// 分页查询返回当前页码。
	CurrentPage *int64 `json:"current_page,omitempty"`
	// 分页查询返回单页数据条数。
	CurrentSize *int64 `json:"current_size,omitempty"`
	// 分页查询返回数据总数。
	TotalCount *int64 `json:"total_count,omitempty"`
	// 分页查询返回分页总数。
	TotalPage *int64 `json:"total_page,omitempty"`
}

// V1MeetingRoomsMonitorDeviceControllerInfoGet200ResponseControllerInfoListInner struct for V1MeetingRoomsMonitorDeviceControllerInfoGet200ResponseControllerInfoListInner
type V1MeetingRoomsMonitorDeviceControllerInfoGet200ResponseControllerInfoListInner struct {
	// 应用程序版本。
	AppVersion *string `json:"app_version,omitempty"`
	// 控制器型号
	ControllerModel *string `json:"controller_model,omitempty"`
	// 控制器名称
	ControllerName *string `json:"controller_name,omitempty"`
	// CPU 类型
	CpuType *string `json:"cpu_type,omitempty"`
	// CPU 当前占有率
	CpuUsage *string `json:"cpu_usage,omitempty"`
	// 固件版本
	FrameworkVersion *string `json:"framework_version,omitempty"`
	// IP 地址
	IpAddress *string `json:"ip_address,omitempty"`
	// MAC 地址
	MacAddress *string `json:"mac_address,omitempty"`
	// 厂商
	ManufactureName *string `json:"manufacture_name,omitempty"`
	// 会议室地址。
	MeetingRoomLocation *string `json:"meeting_room_location,omitempty"`
	// 会议室名称。
	MeetingRoomName *string `json:"meeting_room_name,omitempty"`
	// 内存使用大小
	MemUsage *string `json:"mem_usage,omitempty"`
	// 网络类型
	NetworkType *string `json:"network_type,omitempty"`
	// 会议室 ID。
	RoomsId *string `json:"rooms_id,omitempty"`
	// 设备状态；0:离线 1:在线
	Status *string `json:"status,omitempty"`
}

// V1MeetingRoomsOperatorIdMeetingsGet200Response struct for V1MeetingRoomsOperatorIdMeetingsGet200Response
type V1MeetingRoomsOperatorIdMeetingsGet200Response struct {
	// 当前页。
	CurrentPage *int64 `json:"current_page,omitempty"`
	// 当前实际页大小。
	CurrentSize     *int64                                                               `json:"current_size,omitempty"`
	MeetingInfoList []V1MeetingRoomsOperatorIdMeetingsGet200ResponseMeetingInfoListInner `json:"meeting_info_list,omitempty"`
	// 数据总条数。
	TotalCount *int64 `json:"total_count,omitempty"`
	// 数据总页数。
	TotalPage *int64 `json:"total_page,omitempty"`
}

// V1MeetingRoomsOperatorIdMeetingsGet200ResponseMeetingInfoListInner 会议对象列表。
type V1MeetingRoomsOperatorIdMeetingsGet200ResponseMeetingInfoListInner struct {
	// 会议预订结束时间（Unix 秒）。
	EndTime *int64 `json:"end_time,omitempty"`
	// 有效会议 Code。PMI会议需返回PMI号，不是真实的meeting_code
	MeetingCode *string `json:"meeting_code,omitempty"`
	// 会议 ID。
	MeetingId *string `json:"meeting_id,omitempty"`
	// 会议类型(0-一次性会议，1-周期性会议，2-微信专属会议，4-rooms投屏会议，5-个人会议号会议，6-网络研讨会（Webinar）)
	MeetingType *int64 `json:"meeting_type,omitempty"`
	// 会议预订开始时间（Unix 秒）。
	StartTime *int64 `json:"start_time,omitempty"`
	// 会议状态。
	Status *string `json:"status,omitempty"`
	// 会议主题。
	Subject *string `json:"subject,omitempty"`
}

// V1MeetingRoomsRoomCallInfoPost200Response struct for V1MeetingRoomsRoomCallInfoPost200Response
type V1MeetingRoomsRoomCallInfoPost200Response struct {
	// 最近一次应答时间。
	ResponseTime *string `json:"response_time,omitempty"`
	// 应答状态： 0：无应答，60秒无回应 1：未呼叫 2：入会中 3：被拒绝 4：呼叫中 5：取消呼叫（仅 Rooms 会议室有该状态） 6：已离会
	Status *int64 `json:"status,omitempty"`
}

// V1MeetingRoomsRoomCallInfoPostRequest struct for V1MeetingRoomsRoomCallInfoPostRequest
type V1MeetingRoomsRoomCallInfoPostRequest struct {
	InstanceId int64 `json:"instance_id"`
	// 会议ID
	MeetingId string `json:"meeting_id"`
	// 会议室 ID，与 mra_address 二选一。
	MeetingRoomId *string                                           `json:"meeting_room_id,omitempty"`
	MraAddress    *V1MeetingRoomsCancelRoomCallPutRequestMraAddress `json:"mra_address,omitempty"`
	OperatorId    string                                            `json:"operator_id"`
	// 操作者 ID 的类型： 1：userid
	OperatorIdType int64 `json:"operator_id_type"`
}

// V1MeetingRoomsRoomCallPut200Response struct for V1MeetingRoomsRoomCallPut200Response
type V1MeetingRoomsRoomCallPut200Response struct {
	// 呼叫ID
	InviteId *string `json:"invite_id,omitempty"`
}

// V1MeetingRoomsRoomCallPutRequest struct for V1MeetingRoomsRoomCallPutRequest
type V1MeetingRoomsRoomCallPutRequest struct {
	// 会议ID
	MeetingId int64 `json:"meeting_id"`
	// 会议室 ID，与 mra_address 二选一。
	MeetingRoomId *string                                           `json:"meeting_room_id,omitempty"`
	MraAddress    *V1MeetingRoomsCancelRoomCallPutRequestMraAddress `json:"mra_address,omitempty"`
	OperatorId    string                                            `json:"operator_id"`
	// 操作者 ID 的类型： 1：userid
	OperatorIdType int64 `json:"operator_id_type"`
}

// V1MeetingRoomsRoomConfigInfoPost200Response struct for V1MeetingRoomsRoomConfigInfoPost200Response
type V1MeetingRoomsRoomConfigInfoPost200Response struct {
	MeetingSettings *V1MeetingRoomsRoomConfigInfoPost200ResponseMeetingSettings `json:"meeting_settings,omitempty"`
	RecordSettings  *V1MeetingRoomsRoomConfigInfoPost200ResponseRecordSettings  `json:"record_settings,omitempty"`
}

// V1MeetingRoomsRoomConfigInfoPost200ResponseMeetingSettings 会议室会议配置对象
type V1MeetingRoomsRoomConfigInfoPost200ResponseMeetingSettings struct {
	// 自动接听。 1：开启 2：关闭
	AutoResponse *int64 `json:"auto_response,omitempty"`
	// 字幕。 true：开启 false：关闭
	Caption *bool `json:"caption,omitempty"`
	// Rooms 屏幕是否展示消息通知。 true：开启 false：不开启
	RoomNotification *bool `json:"room_notification,omitempty"`
	// 专属 ID。 true：开启 false：关闭
	RoomPmi         *bool                                                                      `json:"room_pmi,omitempty"`
	RoomPmiSettings *V1MeetingRoomsRoomConfigInfoPost200ResponseMeetingSettingsRoomPmiSettings `json:"room_pmi_settings,omitempty"`
	// 水印。 0：关闭 1：单排水印 2：多排水印
	WaterMark *int64 `json:"water_mark,omitempty"`
}

// V1MeetingRoomsRoomConfigInfoPost200ResponseMeetingSettingsRoomPmiSettings 会议室专属会议号配置，仅专属 ID 开启时有效。
type V1MeetingRoomsRoomConfigInfoPost200ResponseMeetingSettingsRoomPmiSettings struct {
	// 是否允许成员在主持人进会前入会。
	AllowInBeforeHost *bool `json:"allow_in_before_host,omitempty"`
	// 会议指定主持人 ID。
	Hosts []string `json:"hosts,omitempty"`
	// 成员入会静音设置。 0：关闭 1：开启 2：超过6人自动开启
	MuteEnableTypeJoin *int64 `json:"mute_enable_type_join,omitempty"`
	// 入会成员设置。 true：仅企业内部用户可入会 false：所有人可入会
	OnlyEnterpriseUserAllowed *bool `json:"only_enterprise_user_allowed,omitempty"`
	// 专属会议室密码，4-6位数字。
	RoomPmiPsw *string `json:"room_pmi_psw,omitempty"`
	// 是否开启等候室。
	WaitingRoom *bool `json:"waiting_room,omitempty"`
}

// V1MeetingRoomsRoomConfigInfoPost200ResponseRecordSettings 会议室录制配置对象
type V1MeetingRoomsRoomConfigInfoPost200ResponseRecordSettings struct {
	// 是否允许下载云录制。 true：开启 false：关闭
	DownloadRecord *bool `json:"download_record,omitempty"`
	// 分享云录制。 0：关闭分享 1：全部成员可查看 2：仅登录成员可查看 3：仅同企业成员可查看 4：仅参会成员可查看
	ShareRecord *int64 `json:"share_record,omitempty"`
}

// V1MeetingRoomsRoomConfigInfoPostRequest struct for V1MeetingRoomsRoomConfigInfoPostRequest
type V1MeetingRoomsRoomConfigInfoPostRequest struct {
	// 设备类型 ID
	InstanceId int64 `json:"instance_id"`
	// 会议室 ID。
	MeetingRoomId string `json:"meeting_room_id"`
	// 操作者 ID。 operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。
	OperatorId string `json:"operator_id"`
	// 操作者 ID 的类型： 1：userid
	OperatorIdType int64 `json:"operator_id_type"`
}

// V1MeetingRoomsScreencastCodeRoomsInfoGet200Response struct for V1MeetingRoomsScreencastCodeRoomsInfoGet200Response
type V1MeetingRoomsScreencastCodeRoomsInfoGet200Response struct {
	// 中控API密码
	ApiPassword *string `json:"api_password,omitempty"`
	// 中控API开关
	CsApiEnable *bool `json:"cs_api_enable,omitempty"`
	// 会议室ID
	MeetingRoomId *string `json:"meeting_room_id,omitempty"`
	// Rooms ID
	RoomsId *string `json:"rooms_id,omitempty"`
	// rooms的IP列表
	RoomsIpList []string `json:"rooms_ip_list,omitempty"`
}

// V1MeetingsMeetingIdBookRoomsPost200Response struct for V1MeetingsMeetingIdBookRoomsPost200Response
type V1MeetingsMeetingIdBookRoomsPost200Response struct {
	// 会议室对象列表
	MeetingRoomList []V1MeetingsMeetingIdBookRoomsPost200ResponseMeetingRoomListInner `json:"meeting_room_list,omitempty"`
	// 会议室数量
	MeetingRoomNumber *int64 `json:"meeting_room_number,omitempty"`
}

// V1MeetingsMeetingIdBookRoomsPost200ResponseMeetingRoomListInner struct for V1MeetingsMeetingIdBookRoomsPost200ResponseMeetingRoomListInner
type V1MeetingsMeetingIdBookRoomsPost200ResponseMeetingRoomListInner struct {
	// 会议室ID
	MeetingRoomId *string `json:"meeting_room_id,omitempty"`
	// 会议室地址
	MeetingRoomLocation *string `json:"meeting_room_location,omitempty"`
	// 会议室名称
	MeetingRoomName *string `json:"meeting_room_name,omitempty"`
}

// V1MeetingsMeetingIdBookRoomsPostRequest struct for V1MeetingsMeetingIdBookRoomsPostRequest
type V1MeetingsMeetingIdBookRoomsPostRequest struct {
	// true：在会议开始前的一小时内，在 Room 上显示会议主题。默认值为 true。 false：在会议开始前的一小时内，在 Room 上不显示会议主题。 说明：该参数并不影响预定时间晚过当前时间一个小时以上的会议。超过一小时的会议默认不显示会议主题。
	SubjectVisible *bool `json:"subject_visible,omitempty"`
}

// V1RoomsInventoryAccountStatisticsGet200Response struct for V1RoomsInventoryAccountStatisticsGet200Response
type V1RoomsInventoryAccountStatisticsGet200Response struct {
	// 基础版账号使用数
	CustomUsedCount *int64 `json:"custom_used_count,omitempty"`
	// 专业版账号数
	ProCount *int64 `json:"pro_count,omitempty"`
	// 专业版账号使用数
	ProUsedCount *int64 `json:"pro_used_count,omitempty"`
}

// V1RoomsInventoryGet200Response struct for V1RoomsInventoryGet200Response
type V1RoomsInventoryGet200Response struct {
	// 普通设备数
	NormalCount *int64 `json:"normal_count,omitempty"`
	// 普通设备过期数
	NormalExpiredCount *int64 `json:"normal_expired_count,omitempty"`
	// 普通设备使用数
	NormalUsedCount *int64 `json:"normal_used_count,omitempty"`
	// 专款设备数
	SpecialCount *int64 `json:"special_count,omitempty"`
	// 专款设备过期数
	SpecialExpiredCount *int64 `json:"special_expired_count,omitempty"`
	// 专款设备使用数
	SpecialUsedCount *int64 `json:"special_used_count,omitempty"`
}
