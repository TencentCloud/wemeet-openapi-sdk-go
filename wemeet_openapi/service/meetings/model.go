// Package wemeetopenapi is auto generate
/*
   腾讯会议OpenAPI

   SAAS版RESTFUL风格API

   API version: v1.0.4
*/
package wemeetopenapi

// V1AsrConfigPut200Response struct for V1AsrConfigPut200Response
type V1AsrConfigPut200Response struct {
	// 热词设置结果
	CustomizeWords []string `json:"customize_words,omitempty"`
	// 自定义热词标签
	Tag *string `json:"tag,omitempty"`
}

// V1AsrConfigPutRequest struct for V1AsrConfigPutRequest
type V1AsrConfigPutRequest struct {
	//  自定义热词，不得包含数字、特殊字符、中英混合，中文十个字以内，英文 20 个字母以内。同场会议或同一个人每次设置会覆盖上次设置内容。会议维度最多支持设置 500 个，创建者维度最多支持设置 100 个。
	CustomizeWords []string `json:"customize_words"`
	// 会议ID，有该字段则对该场会议生效。不传该字段则对操作人创建的会议生效
	MeetingId *string `json:"meeting_id,omitempty"`
	// 操作者ID
	OperatorId string `json:"operator_id"`
	// 操作者ID类型 1:userid，2:openid
	OperatorIdType int64 `json:"operator_id_type"`
	// 自定义热词标签，便于热词分类，最多支持输入 32 个字符（中英文）
	Tag string `json:"tag"`
}

// V1AsrDetailsGet200Response struct for V1AsrDetailsGet200Response
type V1AsrDetailsGet200Response struct {
	// 分页查询返回分页总数
	CurrPage *int64 `json:"curr_page,omitempty"`
	// 分页查询返回当前页码
	CurrSize *int64 `json:"curr_size,omitempty"`
	// 文件下载链接列表，有效期2个小时
	DownloadUrl []string `json:"download_url,omitempty"`
	// 分页查询返回数据总数
	TotalCount *int64 `json:"total_count,omitempty"`
	// 分页查询返回单页数据条数
	TotalPage *int64 `json:"total_page,omitempty"`
}

// V1AsrPushStatusPostRequest struct for V1AsrPushStatusPostRequest
type V1AsrPushStatusPostRequest struct {
	// 开启/取消转写内容推送 true：开启推送 false：取消推送
	IsOpen bool `json:"is_open"`
	// 会议 ID。
	MeetingId string `json:"meeting_id"`
	// 操作者ID。operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。
	OperatorId string `json:"operator_id"`
	// 操作者ID类型： 1：userid 2:openid
	OperatorIdType int64 `json:"operator_id_type"`
}

// V1HistoryMeetingsUseridGet200Response struct for V1HistoryMeetingsUseridGet200Response
type V1HistoryMeetingsUseridGet200Response struct {
	// 当前页
	CurrentPage *int64 `json:"current_page,omitempty"`
	// 当前实际页大小
	CurrentSize     *int64                                                      `json:"current_size,omitempty"`
	MeetingInfoList []V1HistoryMeetingsUseridGet200ResponseMeetingInfoListInner `json:"meeting_info_list,omitempty"`
	// 数据总条数
	TotalCount *int64 `json:"total_count,omitempty"`
	// 数据总页数
	TotalPage *int64 `json:"total_page,omitempty"`
}

// V1HistoryMeetingsUseridGet200ResponseMeetingInfoListInner struct for V1HistoryMeetingsUseridGet200ResponseMeetingInfoListInner
type V1HistoryMeetingsUseridGet200ResponseMeetingInfoListInner struct {
	// 会议结束时间戳，UNIX 时间戳（单位秒）。
	EndTime *int64 `json:"end_time,omitempty"`
	// 会议 App 的呼入号码。
	MeetingCode *string `json:"meeting_code,omitempty"`
	// 会议ID
	MeetingId *string `json:"meeting_id,omitempty"`
	// 会议类型 0：一次性会议 1：周期性会议 2：微信专属会议 3：Rooms 投屏会议 5：个人会议号会议 6：网络研讨会
	MeetingType *int64 `json:"meeting_type,omitempty"`
	// 会议开始时间戳，UNIX 时间戳（单位秒）。
	StartTime    *int64  `json:"start_time,omitempty"`
	SubMeetingId *string `json:"sub_meeting_id,omitempty"`
	// 会议主题
	Subject *string `json:"subject,omitempty"`
}

// V1MeetingJobIdGet200Response struct for V1MeetingJobIdGet200Response
type V1MeetingJobIdGet200Response struct {
	// 任务错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`
	// 任务下载链接，有效期2小时。
	JobUrl *string `json:"job_url,omitempty"`
	// 任务状态。 0：未完成 1：已完成
	Status *int64 `json:"status,omitempty"`
}

// V1MeetingMeetingIdWaitingRoomGet200Response struct for V1MeetingMeetingIdWaitingRoomGet200Response
type V1MeetingMeetingIdWaitingRoomGet200Response struct {
	// 当前页码
	CurrentPage *int64 `json:"current_page,omitempty"`
	// 当前页数据条数
	CurrentSize *int64 `json:"current_size,omitempty"`
	// 会议CODE
	MeetingCode *string `json:"meeting_code,omitempty"`
	// 会议的唯一ID
	MeetingId *string `json:"meeting_id,omitempty"`
	// 人员对象数组
	Participants []V1MeetingMeetingIdWaitingRoomGet200ResponseParticipantsInner `json:"participants,omitempty"`
	// 预定结束时间（单位：毫秒）
	ScheduleEndTime *int64 `json:"schedule_end_time,omitempty"`
	// 预定开始时间（单位：毫秒）
	ScheduleStartTime *int64 `json:"schedule_start_time,omitempty"`
	// 会议主题
	Subject *string `json:"subject,omitempty"`
	// 总数据条数
	TotalCount *int64 `json:"total_count,omitempty"`
	// 总页数
	TotalPage *int64 `json:"total_page,omitempty"`
}

// V1MeetingMeetingIdWaitingRoomGet200ResponseParticipantsInner struct for V1MeetingMeetingIdWaitingRoomGet200ResponseParticipantsInner
type V1MeetingMeetingIdWaitingRoomGet200ResponseParticipantsInner struct {
	// 应用版本
	AppVersion *string `json:"app_version,omitempty"`
	// 专属链接进入等候室的昵称
	CustomerData *string `json:"customer_data,omitempty"`
	// 用户的终端设备类型： 0：PSTN 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch MacOS 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch iOS
	Instanceid *string `json:"instanceid,omitempty"`
	// 加入时间（单位：毫秒）
	JoinTime *int64 `json:"join_time,omitempty"`
	// 离开时间（单位：毫秒）
	LeftTime *int64 `json:"left_time,omitempty"`
	// 当场会议的用户临时 ID，可用于会控操作，适用于所有用户。
	MsOpenId *string `json:"ms_open_id,omitempty"`
	// OAuth2.0 鉴权用户的ID。
	OpenId *string `json:"open_id,omitempty"`
	// 入会用户名（base64）
	UserName *string `json:"user_name,omitempty"`
	// 成员用户 ID
	Userid *string `json:"userid,omitempty"`
}

// V1MeetingSetWaitingRoomWelcomeMessagePost200Response struct for V1MeetingSetWaitingRoomWelcomeMessagePost200Response
type V1MeetingSetWaitingRoomWelcomeMessagePost200Response struct {
	// 是否开启等候室欢迎语能力。
	EnableWelcome *bool `json:"enable_welcome,omitempty"`
	// 欢迎语，文本类型，最大长度1000字符。欢迎语中如果传入占位符%NICKNAME%（大小写敏感），则该占位符会被替换为被私聊用户的会中昵称。一条消息中支持多个占位符。
	WelcomeText *string `json:"welcome_text,omitempty"`
}

// V1MeetingSetWaitingRoomWelcomeMessagePostRequest struct for V1MeetingSetWaitingRoomWelcomeMessagePostRequest
type V1MeetingSetWaitingRoomWelcomeMessagePostRequest struct {
	// 是否开启等候室欢迎语能力。
	EnableWelcome bool `json:"enable_welcome"`
	// 会议ID
	MeetingId string `json:"meeting_id"`
	// 操作者 ID，设置等候室欢迎语的用户。会议的创建者、主持人、联席主持人，企业管理平台有会控权限的用户可以设置。  operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。
	OperatorId string `json:"operator_id"`
	// 操作者 ID 的类型： 1: 企业内用户 userid。 2: open_id
	OperatorIdType int64 `json:"operator_id_type"`
	// 欢迎语，文本类型，最大长度1000字符。欢迎语中如果传入占位符%NICKNAME%（大小写敏感），则该占位符会被替换为被私聊用户的会中昵称。一条消息中支持多个占位符。
	WelcomeText string `json:"welcome_text"`
}

// V1MeetingsCustomerShortUrlPost200Response struct for V1MeetingsCustomerShortUrlPost200Response
type V1MeetingsCustomerShortUrlPost200Response struct {
	MeetingShortUrlCustomerData *V1MeetingsCustomerShortUrlPost200ResponseMeetingShortUrlCustomerData `json:"meeting_short_url_customer_data,omitempty"`
}

// V1MeetingsCustomerShortUrlPost200ResponseMeetingShortUrlCustomerData 用户专属参会链接对象。
type V1MeetingsCustomerShortUrlPost200ResponseMeetingShortUrlCustomerData struct {
	// 用户专属字段
	CustomerData *string `json:"customer_data,omitempty"`
	// 用户专属参会链接
	MeetingShortUrl *string `json:"meeting_short_url,omitempty"`
}

// V1MeetingsCustomerShortUrlPostRequest struct for V1MeetingsCustomerShortUrlPostRequest
type V1MeetingsCustomerShortUrlPostRequest struct {
	// 自定义参数，长度不超过256 字节。
	CustomerData string `json:"customer_data"`
	// 会议ID
	MeetingId string `json:"meeting_id"`
}

// V1MeetingsGet200Response struct for V1MeetingsGet200Response
type V1MeetingsGet200Response struct {
	MeetingInfoList []V1MeetingsGet200ResponseMeetingInfoListInner `json:"meeting_info_list,omitempty"`
	// 会议数量。
	MeetingNumber *int64 `json:"meeting_number,omitempty"`
	// 分页获取用户会议列表，查询的会议的最后一次修改时间值，UNIX 毫秒级时间戳，分页游标。 因目前一次查询返回会议数量最多为20，当用户会议较多时，如果会议总数量超过20，则需要再次查询。此参数为非必选参数，默认值为0，表示第一次查询利用会议开始时间北京时间当日零点进行查询。 查询返回输出参数“remaining”不为0时，表示还有会议需要继续查询。返回参数“next_cursory”的值即为下一次查询的 cursory 的值。 多次调用该查询接口直到输出参数“remaining”值为0。 当只使用 pos 作为分页条件时,可能会出现查询不到第二页,数据排序出现重复数据等情况与 pos 配合使用。
	NextCursory *int64 `json:"next_cursory,omitempty"`
	// 下次查询时请求里需要携带的 pos 参数。
	NextPos *int64 `json:"next_pos,omitempty"`
	// 是否还剩下会议；因目前一次查询返回会议数量最多为20，如果会议总数量超过20则此字段被置为非0，表示需要再次查询，且下次查询的“pos”参数需从本次响应的“next_pos”字段取值
	Remaining *int64 `json:"remaining,omitempty"`
}

// V1MeetingsGet200ResponseMeetingInfoListInner struct for V1MeetingsGet200ResponseMeetingInfoListInner
type V1MeetingsGet200ResponseMeetingInfoListInner struct {
	CurrentCoHosts []V1MeetingsGet200ResponseMeetingInfoListInnerCurrentCoHostsInner `json:"current_co_hosts,omitempty"`
	CurrentHosts   []V1MeetingsGet200ResponseMeetingInfoListInnerCurrentCoHostsInner `json:"current_hosts,omitempty"`
	// 当前子会议 ID（进行中 / 即将开始）。
	CurrentSubMeetingId       *string                                                   `json:"current_sub_meeting_id,omitempty"`
	EnableDocUploadPermission *bool                                                     `json:"enable_doc_upload_permission,omitempty"`
	EnableEnroll              *bool                                                     `json:"enable_enroll,omitempty"`
	EnableHostKey             *bool                                                     `json:"enable_host_key,omitempty"`
	EnableLive                *bool                                                     `json:"enable_live,omitempty"`
	EndTime                   *string                                                   `json:"end_time,omitempty"`
	Guests                    []V1MeetingsGet200ResponseMeetingInfoListInnerGuestsInner `json:"guests,omitempty"`
	// 0：无更多。  1：有更多子会议特例。
	HasMoreSubMeeting *int64                                                            `json:"has_more_sub_meeting,omitempty"`
	HasVote           *bool                                                             `json:"has_vote,omitempty"`
	HostKey           *string                                                           `json:"host_key,omitempty"`
	Hosts             []V1MeetingsGet200ResponseMeetingInfoListInnerCurrentCoHostsInner `json:"hosts,omitempty"`
	// 查询者在会议中的角色： creator：创建者 hoster：主持人 invitee：被邀请者
	JoinMeetingRole *string                                                 `json:"join_meeting_role,omitempty"`
	JoinUrl         *string                                                 `json:"join_url,omitempty"`
	LiveConfig      *V1MeetingsGet200ResponseMeetingInfoListInnerLiveConfig `json:"live_config,omitempty"`
	Location        *string                                                 `json:"location,omitempty"`
	// 该参数仅提供给支持混合云的企业可见，默认值为0。 0：公网会议 1：专网会议
	MediaSetType  *int64                                                            `json:"media_set_type,omitempty"`
	MeetingCode   *string                                                           `json:"meeting_code,omitempty"`
	MeetingId     *string                                                           `json:"meeting_id,omitempty"`
	MeetingType   *int64                                                            `json:"meeting_type,omitempty"`
	NeedPassword  *bool                                                             `json:"need_password,omitempty"`
	Participants  []V1MeetingsGet200ResponseMeetingInfoListInnerCurrentCoHostsInner `json:"participants,omitempty"`
	Password      *string                                                           `json:"password,omitempty"`
	RecurringRule *V1MeetingsGet200ResponseMeetingInfoListInnerRecurringRule        `json:"recurring_rule,omitempty"`
	// 剩余子会议场数。
	RemainSubMeetings *int64                                                `json:"remain_sub_meetings,omitempty"`
	Settings          *V1MeetingsGet200ResponseMeetingInfoListInnerSettings `json:"settings,omitempty"`
	StartTime         *string                                               `json:"start_time,omitempty"`
	// 当前会议状态： MEETING_STATE_INVALID：非法或未知的会议状态，错误状态。 MEETING_STATE_INIT：待开始，会议预定到预定结束时间前，会议中无人。 MEETING_STATE_CANCELLED：已取消，主持人主动取消会议，待开始的会议才能取消，取消的会议无法再进入。 MEETING_STATE_STARTED：会议中，只要会议中有人即表示进行中。 MEETING_STATE_ENDED：已删除，结束时间后且会议中无人时，被主持人删除，已删除的会议无法再进入。 MEETING_STATE_NULL：无状态，过了预定结束时间，会议中无人。 MEETING_STATE_RECYCLED：已回收，过了预定开始时间30天，会议号被后台回收，无法再进入。
	Status *string `json:"status,omitempty"`
	// 周期性子会议列表。
	SubMeetings  []V1MeetingsGet200ResponseMeetingInfoListInnerSubMeetingsInner `json:"sub_meetings,omitempty"`
	Subject      *string                                                        `json:"subject,omitempty"`
	SyncToWework *bool                                                          `json:"sync_to_wework,omitempty"`
	TimeZone     *string                                                        `json:"time_zone,omitempty"`
	// 会议类型： 0：预约会议类型 1：快速会议类型
	Type *int64 `json:"type,omitempty"`
}

// V1MeetingsGet200ResponseMeetingInfoListInnerCurrentCoHostsInner struct for V1MeetingsGet200ResponseMeetingInfoListInnerCurrentCoHostsInner
type V1MeetingsGet200ResponseMeetingInfoListInnerCurrentCoHostsInner struct {
	Userid *string `json:"userid,omitempty"`
}

// V1MeetingsGet200ResponseMeetingInfoListInnerGuestsInner struct for V1MeetingsGet200ResponseMeetingInfoListInnerGuestsInner
type V1MeetingsGet200ResponseMeetingInfoListInnerGuestsInner struct {
	Area        *string `json:"area,omitempty"`
	GuestName   *string `json:"guest_name,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
}

// V1MeetingsGet200ResponseMeetingInfoListInnerLiveConfig struct for V1MeetingsGet200ResponseMeetingInfoListInnerLiveConfig
type V1MeetingsGet200ResponseMeetingInfoListInnerLiveConfig struct {
	EnableLiveIm       *bool                                                                `json:"enable_live_im,omitempty"`
	EnableLivePassword *bool                                                                `json:"enable_live_password,omitempty"`
	EnableLiveReplay   *bool                                                                `json:"enable_live_replay,omitempty"`
	LiveAddr           *string                                                              `json:"live_addr,omitempty"`
	LivePassword       *string                                                              `json:"live_password,omitempty"`
	LiveSubject        *string                                                              `json:"live_subject,omitempty"`
	LiveSummary        *string                                                              `json:"live_summary,omitempty"`
	LiveWatermark      *V1MeetingsGet200ResponseMeetingInfoListInnerLiveConfigLiveWatermark `json:"live_watermark,omitempty"`
}

// V1MeetingsGet200ResponseMeetingInfoListInnerLiveConfigLiveWatermark struct for V1MeetingsGet200ResponseMeetingInfoListInnerLiveConfigLiveWatermark
type V1MeetingsGet200ResponseMeetingInfoListInnerLiveConfigLiveWatermark struct {
	WatermarkOpt *int64 `json:"watermark_opt,omitempty"`
}

// V1MeetingsGet200ResponseMeetingInfoListInnerRecurringRule 周期性会议设置。
type V1MeetingsGet200ResponseMeetingInfoListInnerRecurringRule struct {
	CustomizedRecurringDays *int64 `json:"customized_recurring_days,omitempty"`
	CustomizedRecurringStep *int64 `json:"customized_recurring_step,omitempty"`
	CustomizedRecurringType *int64 `json:"customized_recurring_type,omitempty"`
	// 周期性会议频率，默认值为0。 0：每天 1：每周一至周五 2：每周 3：每两周 4：每月
	RecurringType *int64 `json:"recurring_type,omitempty"`
	// integer  限定会议次数（1-50次）默认值为7次。
	UntilCount *int64 `json:"until_count,omitempty"`
	// 结束日期时间戳，默认值为当前日期 + 7天。
	UntilDate *int64 `json:"until_date,omitempty"`
	// 结束重复类型，默认值为0。 0：按日期结束重复 1：按次数结束重复
	UntilType *int64 `json:"until_type,omitempty"`
}

// V1MeetingsGet200ResponseMeetingInfoListInnerSettings struct for V1MeetingsGet200ResponseMeetingInfoListInnerSettings
type V1MeetingsGet200ResponseMeetingInfoListInnerSettings struct {
	AllowInBeforeHost           *bool   `json:"allow_in_before_host,omitempty"`
	AllowScreenSharedWatermark  *bool   `json:"allow_screen_shared_watermark,omitempty"`
	AllowUnmuteSelf             *bool   `json:"allow_unmute_self,omitempty"`
	AutoInWaitingRoom           *bool   `json:"auto_in_waiting_room,omitempty"`
	AutoRecordType              *string `json:"auto_record_type,omitempty"`
	EnableHostPauseAutoRecord   *bool   `json:"enable_host_pause_auto_record,omitempty"`
	MuteEnableJoin              *bool   `json:"mute_enable_join,omitempty"`
	MuteEnableTypeJoin          *int64  `json:"mute_enable_type_join,omitempty"`
	OnlyAllowEnterpriseUserJoin *bool   `json:"only_allow_enterprise_user_join,omitempty"`
	ParticipantJoinAutoRecord   *bool   `json:"participant_join_auto_record,omitempty"`
	WaterMarkType               *int64  `json:"water_mark_type,omitempty"`
}

// V1MeetingsGet200ResponseMeetingInfoListInnerSubMeetingsInner struct for V1MeetingsGet200ResponseMeetingInfoListInnerSubMeetingsInner
type V1MeetingsGet200ResponseMeetingInfoListInnerSubMeetingsInner struct {
	EndTime   *string `json:"end_time,omitempty"`
	StartTime *string `json:"start_time,omitempty"`
	// 子会议状态： 0：默认（存在）  1：已删除
	Status       *int64  `json:"status,omitempty"`
	SubMeetingId *string `json:"sub_meeting_id,omitempty"`
}

// V1MeetingsMeetingIdCancelPostRequest struct for V1MeetingsMeetingIdCancelPostRequest
type V1MeetingsMeetingIdCancelPostRequest struct {
	// 用户的终端设备类型： 0：PSTN 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch MacOS 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch iOS
	Instanceid int64 `json:"instanceid"`
	// 会议类型，默认值为0。 0：普通会议 1：周期性会议
	MeetingType *int64 `json:"meeting_type,omitempty"`
	// 原因代码，可为用户自定义
	ReasonCode int64 `json:"reason_code"`
	// 取消原因描述
	ReasonDetail *string `json:"reason_detail,omitempty"`
	// 周期性子会议 ID，如果不传，则会取消该系列的周期性会议
	SubMeetingId *string `json:"sub_meeting_id,omitempty"`
	// 调用 API 的用户 ID（企业内部请使用企业唯一用户标识；OAuth2.0 鉴权用户请使用 openId）。 企业唯一用户标识说明： 1：企业对接 SSO 时使用的员工唯一标识 ID。 2：企业调用创建用户接口时传递的 userid 参数。
	Userid string `json:"userid"`
}

// V1MeetingsMeetingIdCustomerShortUrlGet200Response struct for V1MeetingsMeetingIdCustomerShortUrlGet200Response
type V1MeetingsMeetingIdCustomerShortUrlGet200Response struct {
	// 用户专属参会链接对象。
	MeetingShortUrlCustomerData []V1MeetingsMeetingIdCustomerShortUrlGet200ResponseMeetingShortUrlCustomerDataInner `json:"meeting_short_url_customer_data,omitempty"`
}

// V1MeetingsMeetingIdCustomerShortUrlGet200ResponseMeetingShortUrlCustomerDataInner struct for V1MeetingsMeetingIdCustomerShortUrlGet200ResponseMeetingShortUrlCustomerDataInner
type V1MeetingsMeetingIdCustomerShortUrlGet200ResponseMeetingShortUrlCustomerDataInner struct {
	// 用户专属信息字段，用户自定义参数customerData
	CustomerData *string `json:"customer_data,omitempty"`
	// 根据customerData生成的专属参会链接
	MeetingShortUrl *string `json:"meeting_short_url,omitempty"`
}

// V1MeetingsMeetingIdEnrollApprovalsGet200Response struct for V1MeetingsMeetingIdEnrollApprovalsGet200Response
type V1MeetingsMeetingIdEnrollApprovalsGet200Response struct {
	// 当前页
	CurrentPage *int64 `json:"current_page,omitempty"`
	// 当前页实际大小
	CurrentSize *int64 `json:"current_size,omitempty"`
	// 当前页的报名列表，current_size为0时无该字段
	EnrollList []V1MeetingsMeetingIdEnrollApprovalsGet200ResponseEnrollListInner `json:"enroll_list,omitempty"`
	// 根据条件筛选后的总报名人数
	TotalCount *int64 `json:"total_count,omitempty"`
	// 根据条件筛选后的总分页数
	TotalPage *int64 `json:"total_page,omitempty"`
}

// V1MeetingsMeetingIdEnrollApprovalsGet200ResponseEnrollListInner struct for V1MeetingsMeetingIdEnrollApprovalsGet200ResponseEnrollListInner
type V1MeetingsMeetingIdEnrollApprovalsGet200ResponseEnrollListInner struct {
	// 答题列表
	AnswerList []V1MeetingsMeetingIdEnrollApprovalsGet200ResponseEnrollListInnerAnswerListInner `json:"answer_list,omitempty"`
	// pstn入会凭证
	EnrollCode *string `json:"enroll_code,omitempty"`
	// 报名id
	EnrollId *int64 `json:"enroll_id,omitempty"`
	// 报名来源： 1：用户手动报名 2：批量导入报名
	EnrollSourceType *int64 `json:"enroll_source_type,omitempty"`
	// 报名时间（utc+8，非时间戳）
	EnrollTime *string `json:"enroll_time,omitempty"`
	// 当场会议的用户临时id，所有用户都有
	MsOpenId *string `json:"ms_open_id,omitempty"`
	// 昵称
	NickName *string `json:"nick_name,omitempty"`
	// 报名者已授权过会议创建的应用时返回openid，否则为空
	OpenId *string `json:"open_id,omitempty"`
	// 报名状态：1 待审批，2 已拒绝，3 已批准
	Status *int64 `json:"status,omitempty"`
	// 报名者与会议创建者为同企业时，返回userid，否则为空,导入报名入参为手机号的情况不返回userid。
	Userid *string `json:"userid,omitempty"`
}

// V1MeetingsMeetingIdEnrollApprovalsGet200ResponseEnrollListInnerAnswerListInner struct for V1MeetingsMeetingIdEnrollApprovalsGet200ResponseEnrollListInnerAnswerListInner
type V1MeetingsMeetingIdEnrollApprovalsGet200ResponseEnrollListInnerAnswerListInner struct {
	// 回答内容：单选/简答只有一个元素，多选会有多个
	AnswerContent []string `json:"answer_content,omitempty"`
	// 是否必填：1 否，2 是
	IsRequired *int64 `json:"is_required,omitempty"`
	// 问题编号，1,2,3...等形式
	QuestionNum *int64 `json:"question_num,omitempty"`
	// 问题标题
	QuestionTitle *string `json:"question_title,omitempty"`
	// 问题类型：1 单选，2 多选，3 简答
	QuestionType *int64 `json:"question_type,omitempty"`
	// 特殊问题类型：1 无，2 手机号，3 邮箱，4 姓名，5 公司名称（目前4种特殊问题均为简答题）
	SpecialType *int64 `json:"special_type,omitempty"`
}

// V1MeetingsMeetingIdEnrollApprovalsPut200Response struct for V1MeetingsMeetingIdEnrollApprovalsPut200Response
type V1MeetingsMeetingIdEnrollApprovalsPut200Response struct {
	// 成功处理的数量
	HandledCount *int64 `json:"handled_count,omitempty"`
	// 在线大会唯一标识
	MeetingId *string `json:"meeting_id,omitempty"`
}

// V1MeetingsMeetingIdEnrollApprovalsPutRequest struct for V1MeetingsMeetingIdEnrollApprovalsPutRequest
type V1MeetingsMeetingIdEnrollApprovalsPutRequest struct {
	// 审批动作：1 取消批准，2 拒绝，3.批准，取消批准后状态将变成待审批
	Action int64 `json:"action"`
	// 报名id列表效
	EnrollIdList []int64 `json:"enroll_id_list"`
	// 设备类型
	Instanceid *int64 `json:"instanceid,omitempty"`
	// 操作者 ID。会议创建者可以导入报名信息。 operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。  operator_id_type=2，operator_id必须和公共参数的openid一致。  operator_id和userid至少填写一个，两个参数如果都传了以operator_id为准。  使用OAuth公参鉴权后不能使用userid为入参。
	OperatorId *string `json:"operator_id,omitempty"`
	// 操作者 ID 的类型：  1: userid 2: open_id  如果operator_id和userid具有值，则以operator_id为准；
	OperatorIdType int64 `json:"operator_id_type"`
	// 用户id
	Userid *string `json:"userid,omitempty"`
}

// V1MeetingsMeetingIdEnrollConfigGet200Response struct for V1MeetingsMeetingIdEnrollConfigGet200Response
type V1MeetingsMeetingIdEnrollConfigGet200Response struct {
	// 审批类型：1 自动审批，2 手动审批，默认自动审批
	ApproveType *int64 `json:"approve_type,omitempty"`
	// 报名页封面图URL（base64编码）
	CoverImage []string `json:"cover_image,omitempty"`
	// 显示已报名/预约人数。0：不展示 1：展示，默认开启
	DisplayNumberOfParticipants *int64 `json:"display_number_of_participants,omitempty"`
	// 报名截止时间（秒级时间戳）
	EnrollDeadline *string `json:"enroll_deadline,omitempty"`
	// 报名页简介
	EnrollDescription *string `json:"enroll_description,omitempty"`
	// 报名人数上限
	EnrollNumber *int64 `json:"enroll_number,omitempty"`
	// 报名审批自动通知方式，1-短信通知；2-邮件中文；3-邮件英文；4-邮件中英文；5-公众号
	EnrollPushType []int64 `json:"enroll_push_type,omitempty"`
	// 是否收集问题：1 不收集，2 收集，默认不收集问题
	IsCollectQuestion *int64 `json:"is_collect_question,omitempty"`
	// 会议id
	MeetingId *string `json:"meeting_id,omitempty"`
	// 本企业用户无需报名。 true: 本企业用户无需报名。 false：默认配置，所有用户需要报名。
	NoRegistrationNeededForStaff *bool `json:"no_registration_needed_for_staff,omitempty"`
	// 报名问题列表，自定义问题按传入的顺序排序，预设问题会优先放在最前面，仅开启收集问题时有效
	QuestionList []V1MeetingsMeetingIdEnrollConfigGet200ResponseQuestionListInner `json:"question_list,omitempty"`
}

// V1MeetingsMeetingIdEnrollConfigGet200ResponseQuestionListInner struct for V1MeetingsMeetingIdEnrollConfigGet200ResponseQuestionListInner
type V1MeetingsMeetingIdEnrollConfigGet200ResponseQuestionListInner struct {
	// 是否必填：1 否，2 是
	IsRequired *int64 `json:"is_required,omitempty"`
	// 问题选项列表，按传入的顺序排序，仅单选/多选时有效，最多8个选项，每个选项限40个汉字
	OptionList []V1MeetingsMeetingIdEnrollConfigGet200ResponseQuestionListInnerOptionListInner `json:"option_list,omitempty"`
	// 问题标题，限制40个汉字
	QuestionTitle *string `json:"question_title,omitempty"`
	// 问题类型：1 单选，2 多选，3 简答
	QuestionType *int64 `json:"question_type,omitempty"`
	// 特殊问题类型：1 无，2 手机号，3 邮箱，4 姓名，5 组织名称，6 组织规模（目前除组织规模外均为简答题，组织规模为单选题）
	SpecialType *int64 `json:"special_type,omitempty"`
}

// V1MeetingsMeetingIdEnrollConfigGet200ResponseQuestionListInnerOptionListInner struct for V1MeetingsMeetingIdEnrollConfigGet200ResponseQuestionListInnerOptionListInner
type V1MeetingsMeetingIdEnrollConfigGet200ResponseQuestionListInnerOptionListInner struct {
	Content *string `json:"content,omitempty"`
}

// V1MeetingsMeetingIdEnrollConfigPut200Response struct for V1MeetingsMeetingIdEnrollConfigPut200Response
type V1MeetingsMeetingIdEnrollConfigPut200Response struct {
	// 会议的唯一标识
	MeetingId *string `json:"meeting_id,omitempty"`
	// 报名问题数量，不收集问题时该字段返回0
	QuestionCount *int64 `json:"question_count,omitempty"`
}

// V1MeetingsMeetingIdEnrollConfigPutRequest struct for V1MeetingsMeetingIdEnrollConfigPutRequest
type V1MeetingsMeetingIdEnrollConfigPutRequest struct {
	// 审批类型：1 自动审批，2 手动审批，默认自动审批
	ApproveType *int64 `json:"approve_type,omitempty"`
	// 报名封面图的URL，上传封面为异步形式，通过异步任务结果webhook获取上传结果，列表内容为空字符串时为默认图片，不传或传空列表则不做修改，最多支持5张，支持格式为jpg，jpeg，png。每张不超过5M，按照传入顺序展示。
	CoverImage []string `json:"cover_image,omitempty"`
	// 显示已报名/预约人数。0：不展示 1：展示，默认开启
	DisplayNumberOfParticipants *int64 `json:"display_number_of_participants,omitempty"`
	// 报名截止时间（秒级时间戳）
	EnrollDeadline *string `json:"enroll_deadline,omitempty"`
	// 报名页详情介绍，最多5000字符
	EnrollDescription *string `json:"enroll_description,omitempty"`
	// 报名人数上限
	EnrollNumber *int64 `json:"enroll_number,omitempty"`
	// 报名审批自动通知方式，1-短信通知；2-邮件中文；3-邮件英文；4-邮件中英文；5-公众号
	EnrollPushType []int64 `json:"enroll_push_type,omitempty"`
	// 设备类型
	Instanceid int64 `json:"instanceid"`
	// 是否收集问题：1 不收集，2 收集，默认不收集问题
	IsCollectQuestion *int64 `json:"is_collect_question,omitempty"`
	// 本企业用户无需报名。 true: 本企业用户无需报名。 false：默认配置，所有用户需要报名。
	NoRegistrationNeededForStaff *bool `json:"no_registration_needed_for_staff,omitempty"`
	// 操作者 ID。会议创建者可以导入报名信息。 operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。  operator_id_type=2，operator_id必须和公共参数的openid一致。  operator_id和userid至少填写一个，两个参数如果都传了以operator_id为准。  使用OAuth公参鉴权后不能使用userid为入参。
	OperatorId *string `json:"operator_id,omitempty"`
	// 操作者 ID 的类型：  1: userid 2: open_id  如果operator_id和userid具有值，则以operator_id为准；
	OperatorIdType *int64 `json:"operator_id_type,omitempty"`
	// 报名问题列表，非特殊问题按传入的顺序排序，特殊问题会优先放在最前面，仅开启收集问题时有效
	QuestionList []V1MeetingsMeetingIdEnrollConfigPutRequestQuestionListInner `json:"question_list,omitempty"`
	// 用户id
	Userid *string `json:"userid,omitempty"`
}

// V1MeetingsMeetingIdEnrollConfigPutRequestQuestionListInner struct for V1MeetingsMeetingIdEnrollConfigPutRequestQuestionListInner
type V1MeetingsMeetingIdEnrollConfigPutRequestQuestionListInner struct {
	// 是否必填：1 否，2 是
	IsRequired int64 `json:"is_required"`
	// 问题选项列表，按传入的顺序排序，仅单选/多选时有效，最多8个选项，每个选项限40个汉字
	OptionList []V1MeetingsMeetingIdEnrollConfigPutRequestQuestionListInnerOptionListInner `json:"option_list,omitempty"`
	// 问题标题，限制40个字符（special_type为特殊问题时，该字段无效）
	QuestionTitle *string `json:"question_title,omitempty"`
	// 问题类型：1 单选，2 多选，3 简答（special_type为特殊问题时，该字段无效）
	QuestionType *int64 `json:"question_type,omitempty"`
	// 特殊问题类型：1 无，2 手机号，3 邮箱，4 姓名，5 组织名称，6 组织规模（目前除组织规模外均为简答题，组织规模为单选题）
	SpecialType *int64 `json:"special_type,omitempty"`
}

// V1MeetingsMeetingIdEnrollConfigPutRequestQuestionListInnerOptionListInner struct for V1MeetingsMeetingIdEnrollConfigPutRequestQuestionListInnerOptionListInner
type V1MeetingsMeetingIdEnrollConfigPutRequestQuestionListInnerOptionListInner struct {
	Content *string `json:"content,omitempty"`
}

// V1MeetingsMeetingIdEnrollIdsPost200Response struct for V1MeetingsMeetingIdEnrollIdsPost200Response
type V1MeetingsMeetingIdEnrollIdsPost200Response struct {
	// 成员报名 ID 数组，仅返回已报名成员的报名 ID，若传入的用户无人报名，则无该字段。
	EnrollIdList []V1MeetingsMeetingIdEnrollIdsPost200ResponseEnrollIdListInner `json:"enroll_id_list,omitempty"`
	// 会议ID
	MeetingId *string `json:"meeting_id,omitempty"`
}

// V1MeetingsMeetingIdEnrollIdsPost200ResponseEnrollIdListInner struct for V1MeetingsMeetingIdEnrollIdsPost200ResponseEnrollIdListInner
type V1MeetingsMeetingIdEnrollIdsPost200ResponseEnrollIdListInner struct {
	// 报名ID
	EnrollId *int64 `json:"enroll_id,omitempty"`
	// 当场会议的用户临时 ID，适用于所有用户。
	MsOpenId *string `json:"ms_open_id,omitempty"`
}

// V1MeetingsMeetingIdEnrollIdsPostRequest struct for V1MeetingsMeetingIdEnrollIdsPostRequest
type V1MeetingsMeetingIdEnrollIdsPostRequest struct {
	// 用户的终端设备类型： 0：PSTN 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch MacOS 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch iOS
	Instanceid int64 `json:"instanceid"`
	// 当场会议的用户临时 ID（适用于所有用户）数组，单次最多支持500条。
	MsOpenIdList []string `json:"ms_open_id_list"`
	// 操作者 ID。会议创建者可以导入报名信息。 operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。 operator_id_type=2，operator_id 必须和公共参数的 openid 一致。 operator_id 和 userid 至少填写一个，两个参数如果都传了以 operator_id 为准。 使用 OAuth 公参鉴权后不能使用 userid 为入参。
	OperatorId *string `json:"operator_id,omitempty"`
	// 操作者 ID 的类型： 1：userid 2：open_id 如果 operator_id 和 userid 具有值，则以 operator_id 为准。
	OperatorIdType *int64 `json:"operator_id_type,omitempty"`
	// 查询报名 ID 的排序规则。当该账号存在多条报名记录（手机号导入、手动报名等）时，该接口返回的顺序。 1：优先查询手机号导入报名，再查询用户手动报名，默认值。 2：优先查询用户手动报名，再查手机号导入。
	SortingRules *int64 `json:"sorting_rules,omitempty"`
	// 会议创建者的用户 ID。为了防止现网应用报错，此参数实则仍然兼容 openid，如无 oauth 应用使用报名接口则也可做成不兼容变更。
	Userid *string `json:"userid,omitempty"`
}

// V1MeetingsMeetingIdEnrollImportPost200Response struct for V1MeetingsMeetingIdEnrollImportPost200Response
type V1MeetingsMeetingIdEnrollImportPost200Response struct {
	// 报名对象列表
	EnrollList []V1MeetingsMeetingIdEnrollImportPost200ResponseEnrollListInner `json:"enroll_list,omitempty"`
	// 成功导入的报名信息条数
	TotalCount *int64 `json:"total_count,omitempty"`
	// 未注册用户列表
	UserNonRegistered []string `json:"user_non_registered,omitempty"`
}

// V1MeetingsMeetingIdEnrollImportPost200ResponseEnrollListInner struct for V1MeetingsMeetingIdEnrollImportPost200ResponseEnrollListInner
type V1MeetingsMeetingIdEnrollImportPost200ResponseEnrollListInner struct {
	// 国家/地区代码，若使用手机号，必填（例如：中国传86，不是+86）
	Area       *string `json:"area,omitempty"`
	EnrollCode *string `json:"enroll_code,omitempty"`
	// 报名ID
	EnrollId *int64 `json:"enroll_id,omitempty"`
	// 报名的昵称，与会中昵称可能不一致
	NickName *string `json:"nick_name,omitempty"`
	// OAuth授权用户ID。  导入报名对象支持本企业（或与OAuth应用同企业）内 userid、授权用户的openid、phone_number 三种形式，三者必填其一；
	OpenId *string `json:"open_id,omitempty"`
	// 手机号
	PhoneNumber *string `json:"phone_number,omitempty"`
	// 用户的唯一 ID（企业内部请使用企业唯一用户标识）。 userid 和 phone_number 两者必填一个
	Userid *string `json:"userid,omitempty"`
}

// V1MeetingsMeetingIdEnrollImportPostRequest struct for V1MeetingsMeetingIdEnrollImportPostRequest
type V1MeetingsMeetingIdEnrollImportPostRequest struct {
	// 导入的报名对象列表，单次导入最大1000条。累计导入最大4000
	EnrollList []V1MeetingsMeetingIdEnrollImportPostRequestEnrollListInner `json:"enroll_list"`
	// 操作者的终端设备类型
	Instanceid int64 `json:"instanceid"`
	// 操作者 ID。operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。operator_id_type=2，operator_id必须和公共参数的openid一致。  使用OAuth公参鉴权后不能使用userid为入参。
	OperatorId string `json:"operator_id"`
	// 操作者 ID 的类型： 1. 企业用户 userid 2 open_id
	OperatorIdType int64 `json:"operator_id_type"`
}

// V1MeetingsMeetingIdEnrollImportPostRequestEnrollListInner struct for V1MeetingsMeetingIdEnrollImportPostRequestEnrollListInner
type V1MeetingsMeetingIdEnrollImportPostRequestEnrollListInner struct {
	// 国家/地区代码，若使用手机号，必填（例如：中国传86，不是+86）
	Area *string `json:"area,omitempty"`
	// 报名的昵称，与会中昵称可能不一致
	NickName *string `json:"nick_name,omitempty"`
	// OAuth授权用户ID。  导入报名对象支持本企业（或与OAuth应用同企业）内 userid、授权用户的openid、phone_number 三种形式，三者必填其一；  如果都传了以openid为准；（优先级为：openid -> userid -> phone_number）  JWT鉴权方式无法使用open_id导入报名。
	OpenId *string `json:"open_id,omitempty"`
	// 手机号,导入报名支持本企业（或与OAuth应用同企业）内 userid、授权用户的openid、phone_number 三种形式，三者必填其一。
	PhoneNumber *string `json:"phone_number,omitempty"`
	// 用户的唯一 ID（企业内部请使用企业唯一用户标识）。 导入报名支持本企业（或与OAuth应用同企业）内 userid、授权用户的openid、phone_number 三种形式，三者必填其一。
	Userid *string `json:"userid,omitempty"`
}

// V1MeetingsMeetingIdEnrollUnregistrationDelete200Response struct for V1MeetingsMeetingIdEnrollUnregistrationDelete200Response
type V1MeetingsMeetingIdEnrollUnregistrationDelete200Response struct {
	// 成功删除的报名信息数量
	TotalCount *int64 `json:"total_count,omitempty"`
}

// V1MeetingsMeetingIdEnrollUnregistrationDeleteRequest struct for V1MeetingsMeetingIdEnrollUnregistrationDeleteRequest
type V1MeetingsMeetingIdEnrollUnregistrationDeleteRequest struct {
	// 报名对象列表
	EnrollIdList []V1MeetingsMeetingIdEnrollUnregistrationDeleteRequestEnrollIdListInner `json:"enroll_id_list"`
	// 操作者的终端设备类型
	Instanceid int64 `json:"instanceid"`
	// 操作者 ID。operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。
	OperatorId string `json:"operator_id"`
	// 操作者 ID 的类型： 1. 企业用户 userid 2: open_id
	OperatorIdType int64 `json:"operator_id_type"`
}

// V1MeetingsMeetingIdEnrollUnregistrationDeleteRequestEnrollIdListInner struct for V1MeetingsMeetingIdEnrollUnregistrationDeleteRequestEnrollIdListInner
type V1MeetingsMeetingIdEnrollUnregistrationDeleteRequestEnrollIdListInner struct {
	// 报名ID
	EnrollId int64 `json:"enroll_id"`
}

// V1MeetingsMeetingIdGet200Response struct for V1MeetingsMeetingIdGet200Response
type V1MeetingsMeetingIdGet200Response struct {
	MeetingInfoList []V1MeetingsMeetingIdGet200ResponseMeetingInfoListInner `json:"meeting_info_list,omitempty"`
	MeetingNumber   *int64                                                  `json:"meeting_number,omitempty"`
}

// V1MeetingsMeetingIdGet200ResponseMeetingInfoListInner struct for V1MeetingsMeetingIdGet200ResponseMeetingInfoListInner
type V1MeetingsMeetingIdGet200ResponseMeetingInfoListInner struct {
	CurrentCoHosts            []V1MeetingsGet200ResponseMeetingInfoListInnerCurrentCoHostsInner       `json:"current_co_hosts,omitempty"`
	CurrentHosts              []V1MeetingsGet200ResponseMeetingInfoListInnerCurrentCoHostsInner       `json:"current_hosts,omitempty"`
	CurrentSubMeetingId       *string                                                                 `json:"current_sub_meeting_id,omitempty"`
	EnableDocUploadPermission *bool                                                                   `json:"enable_doc_upload_permission,omitempty"`
	EnableEnroll              *bool                                                                   `json:"enable_enroll,omitempty"`
	EnableHostKey             *bool                                                                   `json:"enable_host_key,omitempty"`
	EnableLive                *bool                                                                   `json:"enable_live,omitempty"`
	EndTime                   *string                                                                 `json:"end_time,omitempty"`
	Guests                    []V1MeetingsGet200ResponseMeetingInfoListInnerGuestsInner               `json:"guests,omitempty"`
	HasMoreSubMeeting         *int64                                                                  `json:"has_more_sub_meeting,omitempty"`
	HasVote                   *bool                                                                   `json:"has_vote,omitempty"`
	HostKey                   *string                                                                 `json:"host_key,omitempty"`
	Hosts                     []V1MeetingsGet200ResponseMeetingInfoListInnerCurrentCoHostsInner       `json:"hosts,omitempty"`
	JoinUrl                   *string                                                                 `json:"join_url,omitempty"`
	LiveConfig                *V1MeetingsMeetingIdGet200ResponseMeetingInfoListInnerLiveConfig        `json:"live_config,omitempty"`
	Location                  *string                                                                 `json:"location,omitempty"`
	MeetingCode               *string                                                                 `json:"meeting_code,omitempty"`
	MeetingId                 *string                                                                 `json:"meeting_id,omitempty"`
	MeetingType               *int64                                                                  `json:"meeting_type,omitempty"`
	NeedPassword              *bool                                                                   `json:"need_password,omitempty"`
	Participants              []V1MeetingsGet200ResponseMeetingInfoListInnerCurrentCoHostsInner       `json:"participants,omitempty"`
	Password                  *string                                                                 `json:"password,omitempty"`
	RecurringRule             *V1MeetingsMeetingIdGet200ResponseMeetingInfoListInnerRecurringRule     `json:"recurring_rule,omitempty"`
	RemainSubMeetings         *int64                                                                  `json:"remain_sub_meetings,omitempty"`
	Settings                  *V1MeetingsMeetingIdGet200ResponseMeetingInfoListInnerSettings          `json:"settings,omitempty"`
	StartTime                 *string                                                                 `json:"start_time,omitempty"`
	Status                    *string                                                                 `json:"status,omitempty"`
	SubMeetings               []V1MeetingsMeetingIdGet200ResponseMeetingInfoListInnerSubMeetingsInner `json:"sub_meetings,omitempty"`
	Subject                   *string                                                                 `json:"subject,omitempty"`
	SyncToWework              *bool                                                                   `json:"sync_to_wework,omitempty"`
	TimeZone                  *string                                                                 `json:"time_zone,omitempty"`
	Type                      *int64                                                                  `json:"type,omitempty"`
}

// V1MeetingsMeetingIdGet200ResponseMeetingInfoListInnerLiveConfig struct for V1MeetingsMeetingIdGet200ResponseMeetingInfoListInnerLiveConfig
type V1MeetingsMeetingIdGet200ResponseMeetingInfoListInnerLiveConfig struct {
	EnableLiveIm     *bool                                                                `json:"enable_live_im,omitempty"`
	EnableLiveReplay *bool                                                                `json:"enable_live_replay,omitempty"`
	LiveAddr         *string                                                              `json:"live_addr,omitempty"`
	LivePassword     *string                                                              `json:"live_password,omitempty"`
	LiveSubject      *string                                                              `json:"live_subject,omitempty"`
	LiveSummary      *string                                                              `json:"live_summary,omitempty"`
	LiveWatermark    *V1MeetingsGet200ResponseMeetingInfoListInnerLiveConfigLiveWatermark `json:"live_watermark,omitempty"`
}

// V1MeetingsMeetingIdGet200ResponseMeetingInfoListInnerRecurringRule struct for V1MeetingsMeetingIdGet200ResponseMeetingInfoListInnerRecurringRule
type V1MeetingsMeetingIdGet200ResponseMeetingInfoListInnerRecurringRule struct {
	CustomizedRecurringDays *int64 `json:"customized_recurring_days,omitempty"`
	CustomizedRecurringStep *int64 `json:"customized_recurring_step,omitempty"`
	CustomizedRecurringType *int64 `json:"customized_recurring_type,omitempty"`
	RecurringType           *int64 `json:"recurring_type,omitempty"`
	UntilCount              *int64 `json:"until_count,omitempty"`
	UntilDate               *int64 `json:"until_date,omitempty"`
	UntilType               *int64 `json:"until_type,omitempty"`
}

// V1MeetingsMeetingIdGet200ResponseMeetingInfoListInnerSettings struct for V1MeetingsMeetingIdGet200ResponseMeetingInfoListInnerSettings
type V1MeetingsMeetingIdGet200ResponseMeetingInfoListInnerSettings struct {
	AllowInBeforeHost          *bool   `json:"allow_in_before_host,omitempty"`
	AllowScreenSharedWatermark *bool   `json:"allow_screen_shared_watermark,omitempty"`
	AllowUnmuteSelf            *bool   `json:"allow_unmute_self,omitempty"`
	AutoInWaitingRoom          *bool   `json:"auto_in_waiting_room,omitempty"`
	AutoRecordType             *string `json:"auto_record_type,omitempty"`
	// 是否允许用户自己改名 1:允许用户自己改名，2:不允许用户自己改名，默认为1
	ChangeNickname              *int64 `json:"change_nickname,omitempty"`
	EnableHostPauseAutoRecord   *bool  `json:"enable_host_pause_auto_record,omitempty"`
	MuteEnableJoin              *bool  `json:"mute_enable_join,omitempty"`
	MuteEnableTypeJoin          *int64 `json:"mute_enable_type_join,omitempty"`
	OnlyAllowEnterpriseUserJoin *bool  `json:"only_allow_enterprise_user_join,omitempty"`
	// 是否仅受邀成员可入会，默认值为false，true：仅受邀成员可入会，false：所有成员可入会
	OnlyInviteesAllowed       *bool  `json:"only_invitees_allowed,omitempty"`
	ParticipantJoinAutoRecord *bool  `json:"participant_join_auto_record,omitempty"`
	WaterMarkType             *int64 `json:"water_mark_type,omitempty"`
}

// V1MeetingsMeetingIdGet200ResponseMeetingInfoListInnerSubMeetingsInner struct for V1MeetingsMeetingIdGet200ResponseMeetingInfoListInnerSubMeetingsInner
type V1MeetingsMeetingIdGet200ResponseMeetingInfoListInnerSubMeetingsInner struct {
	EndTime      *string `json:"end_time,omitempty"`
	StartTime    *string `json:"start_time,omitempty"`
	Status       *int64  `json:"status,omitempty"`
	SubMeetingId *string `json:"sub_meeting_id,omitempty"`
}

// V1MeetingsMeetingIdInviteesGet200Response struct for V1MeetingsMeetingIdInviteesGet200Response
type V1MeetingsMeetingIdInviteesGet200Response struct {
	// 是否还存在受邀成员需要继续查询。
	HasRemaining *bool `json:"has_remaining,omitempty"`
	// 会议邀请的参会者
	Invitees []V1MeetingsMeetingIdInviteesGet200ResponseInviteesInner `json:"invitees,omitempty"`
	// 当has_remaining为true时，下次查询时输入参数pos的值
	NextPos *int64 `json:"next_pos,omitempty"`
}

// V1MeetingsMeetingIdInviteesGet200ResponseInviteesInner struct for V1MeetingsMeetingIdInviteesGet200ResponseInviteesInner
type V1MeetingsMeetingIdInviteesGet200ResponseInviteesInner struct {
	// 用户昵称
	NickName *string `json:"nick_name,omitempty"`
	// 用户的唯一 ID（企业内部请使用企业唯一用户标识；OAuth2.0 鉴权用户请使用 openId）
	Userid *string `json:"userid,omitempty"`
}

// V1MeetingsMeetingIdInviteesPutRequest struct for V1MeetingsMeetingIdInviteesPutRequest
type V1MeetingsMeetingIdInviteesPutRequest struct {
	// 用户的终端设备类型：1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch Mac 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch Iphone
	Instanceid int64 `json:"instanceid"`
	// 会议邀请的参会者（传空数组或不传会清空受邀成员列表，最大支持2000人）
	Invitees []V1MeetingsMeetingIdInviteesPutRequestInviteesInner `json:"invitees,omitempty"`
	// 会议创建者ID。调用方用于标示用户的唯一 ID（企业内部请使用企业唯一用户标识；OAuth2.0 鉴权用户请使用 openId）。 企业唯一用户标识说明： 1. 企业对接 SSO 时使用的员工唯一标识 ID。 2. 企业调用创建用户接口时传递的 userid 参数。
	Userid string `json:"userid"`
}

// V1MeetingsMeetingIdInviteesPutRequestInviteesInner struct for V1MeetingsMeetingIdInviteesPutRequestInviteesInner
type V1MeetingsMeetingIdInviteesPutRequestInviteesInner struct {
	// 调用方用于标示用户的唯一 ID，仅支持邀请与会议创建者同企业的成员（企业内部请使用企业唯一用户标识；OAuth2.0 鉴权用户请使用 openId）。 企业唯一用户标识说明： 企业对接 SSO 时使用的员工唯一标识 ID。 企业调用创建用户接口时传递的 userid 参数。
	Userid string `json:"userid"`
}

// V1MeetingsMeetingIdParticipantsGet200Response struct for V1MeetingsMeetingIdParticipantsGet200Response
type V1MeetingsMeetingIdParticipantsGet200Response struct {
	// 是否还有未拉取的数据，该接口可多次拉取到的数据总量上限为5w条。
	HasRemaining *bool `json:"has_remaining,omitempty"`
	// 会议号码。
	MeetingCode *string `json:"meeting_code,omitempty"`
	// 会议的唯一 ID。
	MeetingId *string `json:"meeting_id,omitempty"`
	// 和“has_remaining”一起，数据量比较大的情况下支持参会成员列表的多次获取。
	NextPos *int64 `json:"next_pos,omitempty"`
	// 参会人员对象数组。
	Participants []V1MeetingsMeetingIdParticipantsGet200ResponseParticipantsInner `json:"participants,omitempty"`
	// 预定会议开始时间戳（单位秒）。
	ScheduleEndTime *string `json:"schedule_end_time,omitempty"`
	// 预定会议结束时间戳（单位秒）。
	ScheduleStartTime *string `json:"schedule_start_time,omitempty"`
	// 会议主题。
	Subject *string `json:"subject,omitempty"`
	// 当前参会总人次。
	TotalCount *int64 `json:"total_count,omitempty"`
}

// V1MeetingsMeetingIdParticipantsGet200ResponseParticipantsInner struct for V1MeetingsMeetingIdParticipantsGet200ResponseParticipantsInner
type V1MeetingsMeetingIdParticipantsGet200ResponseParticipantsInner struct {
	// 用户的客户端版本。当用户在会中时才能返回。
	AppVersion *string `json:"app_version,omitempty"`
	// 麦克风状态： true：开启 false：关闭
	AudioState *bool `json:"audio_state,omitempty"`
	// 专属字段
	CustomerData *string `json:"customer_data,omitempty"`
	// 用户的终端设备类型： 0：PSTN 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备（即MRA设备） 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch MacOS 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch iOS/iPadOS
	Instanceid *int64 `json:"instanceid,omitempty"`
	// 用户的 IP 地址。当用户在会中时才能返回。
	Ip *string `json:"ip,omitempty"`
	// 参会者加入会议时间戳（单位秒）。
	JoinTime *string `json:"join_time,omitempty"`
	// 入会方式： 0：PSTN 普通用户，标准的手机或固话类型 1：普通 VOIP 用户 2：附属投屏 VOIP 3：linux sdk for VOIP 4：附属语音 PSTN 5：附属视频 PSTN 6：linux sdk for PSTN
	JoinType *int64 `json:"join_type,omitempty"`
	// 参会者离开会议时间戳（单位秒）。
	LeftTime *string `json:"left_time,omitempty"`
	// 用户的连接方式：UDP、TCP、未知，当用户在会中时才能返回。
	LinkType *string `json:"link_type,omitempty"`
	// 用户的地理位置信息。当用户在会中时才能返回。
	Location *string `json:"location,omitempty"`
	// 当场会议的用户临时 ID，可用于会控操作，适用于所有用户。
	MsOpenId *string `json:"ms_open_id,omitempty"`
	// 网络类型：有线、WIFI、2G、3G、4G、未知。当用户在会中时才能返回。
	Net    *string `json:"net,omitempty"`
	OpenId *string `json:"open_id,omitempty"`
	// 参会者手机号 hash 值 SHA256（手机号 + \"/\" + secretid）。
	Phone *string `json:"phone,omitempty"`
	// 屏幕共享状态： true：开启 false：关闭
	ScreenSharedState *bool `json:"screen_shared_state,omitempty"`
	// 入会用户名（base64）。
	UserName *string `json:"user_name,omitempty"`
	// 用户角色： 0：普通成员角色 1：创建者角色 2：主持人 3：创建者+主持人 4：游客 5：游客+主持人 6：联席主持人 7：创建者+联席主持人
	UserRole *int64 `json:"user_role,omitempty"`
	// 参会者用户 ID。 使用企业自建应用鉴权方式（JWT）时，该值为企业唯一用户标识。
	Userid *string `json:"userid,omitempty"`
	// 用户的身份 ID，仅适用于单场会议。
	Uuid *string `json:"uuid,omitempty"`
	// 摄像头状态： true：开启 false：关闭
	VideoState *bool `json:"video_state,omitempty"`
	// 网络研讨会成员角色： 0：普通参会角色 1：内部嘉宾 2：外部嘉宾 3：邀请链接入会嘉宾 4：观众 5：有音视频权限的研讨会观众
	WebinarMemberRole *int64 `json:"webinar_member_role,omitempty"`
}

// V1MeetingsMeetingIdPut200Response struct for V1MeetingsMeetingIdPut200Response
type V1MeetingsMeetingIdPut200Response struct {
	// 会议列表
	MeetingInfoList []V1MeetingsMeetingIdPut200ResponseMeetingInfoListInner `json:"meeting_info_list,omitempty"`
	// 会议数量
	MeetingNumber *int64 `json:"meeting_number,omitempty"`
}

// V1MeetingsMeetingIdPut200ResponseMeetingInfoListInner struct for V1MeetingsMeetingIdPut200ResponseMeetingInfoListInner
type V1MeetingsMeetingIdPut200ResponseMeetingInfoListInner struct {
	// 是否开启直播
	EnableLive *bool `json:"enable_live,omitempty"`
	// 主持人密钥，仅支持6位数字。 如开启主持人密钥后没有填写此项，将自动分配一个6位数字的密钥。
	HostKey    *string                                                          `json:"host_key,omitempty"`
	LiveConfig *V1MeetingsMeetingIdPut200ResponseMeetingInfoListInnerLiveConfig `json:"live_config,omitempty"`
	// 会议号码
	MeetingCode *string `json:"meeting_code,omitempty"`
	// 会议的唯一 ID
	MeetingId *string                                                        `json:"meeting_id,omitempty"`
	Settings  *V1MeetingsMeetingIdPut200ResponseMeetingInfoListInnerSettings `json:"settings,omitempty"`
}

// V1MeetingsMeetingIdPut200ResponseMeetingInfoListInnerLiveConfig 直播配置对象，内部只返回 live_addr（直播观看地址）。
type V1MeetingsMeetingIdPut200ResponseMeetingInfoListInnerLiveConfig struct {
	// 直播观看地址
	LiveAddr *string `json:"live_addr,omitempty"`
}

// V1MeetingsMeetingIdPut200ResponseMeetingInfoListInnerSettings struct for V1MeetingsMeetingIdPut200ResponseMeetingInfoListInnerSettings
type V1MeetingsMeetingIdPut200ResponseMeetingInfoListInnerSettings struct {
	// 是否允许用户自己改名 1:允许用户自己改名，2:不允许用户自己改名，默认为1
	ChangeNickname *int64 `json:"change_nickname,omitempty"`
	// 是否仅受邀成员可入会，默认值为false，true：仅受邀成员可入会，false：所有成员可入会
	OnlyInviteesAllowed *bool `json:"only_invitees_allowed,omitempty"`
}

// V1MeetingsMeetingIdPutRequest struct for V1MeetingsMeetingIdPutRequest
type V1MeetingsMeetingIdPutRequest struct {
	// 如果没有权限创建专网会议，该字段忽略
	AllowEnterpriseIntranetOnly *bool `json:"allow_enterprise_intranet_only,omitempty"`
	// 是否允许成员上传文档，默认为允许
	EnableDocUploadPermission *bool `json:"enable_doc_upload_permission,omitempty"`
	// 是否开启报名开关，默认不开启
	EnableEnroll *bool `json:"enable_enroll,omitempty"`
	// 是否开启主持人密钥，默认为false。 true：开启 false：关闭 修改周期性会议的主持人密钥适用于接下来的所有子会议。
	EnableHostKey *bool `json:"enable_host_key,omitempty"`
	// 是否开启直播,默认不开启
	EnableLive *bool `json:"enable_live,omitempty"`
	// 会议结束时间戳（单位秒）
	EndTime *string `json:"end_time,omitempty"`
	// 会议嘉宾列表，会议嘉宾不受会议密码和等候室的限制。不传该字段代表不修改。注意：传空数组会清空嘉宾列表。
	Guests []V1MeetingsPostRequestGuestsInner `json:"guests,omitempty"`
	// 主持人密钥，仅支持6位数字。 如开启主持人密钥后没有填写此项，将自动分配一个6位数字的密钥。
	HostKey *string `json:"host_key,omitempty"`
	// 主持人列表。 不传入该字段则不修改，传入空列表即覆盖为空。
	Hosts []V1MeetingsMeetingIdPutRequestHostsInner `json:"hosts,omitempty"`
	// 用户的终端设备类型： 0：PSTN 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch MacOS 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch iOS
	Instanceid int64 `json:"instanceid"`
	// 邀请人列表，调用方用于标示用户的唯一 ID，仅支持邀请与会议创建者同企业的成员（企业内部请使用企业唯一用户标识；OAuth2.0 鉴权用户请使用 openId）。 企业唯一用户标识说明： 企业对接 SSO 时使用的员工唯一标识 ID。 企业调用创建用户接口时传递的 userid 参数。 注意：仅腾讯会议商业版和企业版可邀请参会者，邀请者列表仅支持300人；需要邀请超过300人的场景请调用 设置会议邀请成员 接口。
	Invitees   []V1MeetingsPostRequestInviteesInner     `json:"invitees,omitempty"`
	LiveConfig *V1MeetingsMeetingIdPutRequestLiveConfig `json:"live_config,omitempty"`
	// 会议地点。最长支持18个汉字或36个英文字母。
	Location *string `json:"location,omitempty"`
	// 该参数仅提供给支持混合云的企业可见，默认值为0。 0：外部会议 1：内部会议
	MediaSetType *int64 `json:"media_set_type,omitempty"`
	// 默认值为0。  0：普通会议  1：周期性会议（周期性会议时 type 不能为快速会议，同一账号同时最多可预定50场周期性会议）
	MeetingType *int64 `json:"meeting_type,omitempty"`
	// 会议密码（4~6位数字），可不填。如果将字段值改为空字符串\"\"，则表示取消会议密码
	Password      *string                                     `json:"password,omitempty"`
	RecurringRule *V1MeetingsMeetingIdPutRequestRecurringRule `json:"recurring_rule,omitempty"`
	Settings      *V1MeetingsMeetingIdPutRequestSettings      `json:"settings,omitempty"`
	// 会议开始时间戳（单位秒）
	StartTime *string `json:"start_time,omitempty"`
	// 会议主题
	Subject *string `json:"subject,omitempty"`
	// 时区，可参见 Oracle-TimeZone 标准。
	TimeZone *string `json:"time_zone,omitempty"`
	// 会议创建者 ID（企业内部请使用企业唯一用户标识；OAuth2.0 鉴权用户请使用 openId）。 企业唯一用户标识说明： 1. 企业对接 SSO 时使用的员工唯一标识 ID。 2. 企业调用创建用户接口时传递的 userid 参数。
	Userid string `json:"userid"`
}

// V1MeetingsMeetingIdPutRequestHostsInner struct for V1MeetingsMeetingIdPutRequestHostsInner
type V1MeetingsMeetingIdPutRequestHostsInner struct {
	// 用户是否匿名入会，缺省为 false，不匿名。 true：匿名 false：不匿名
	IsAnonymous *bool `json:"is_anonymous,omitempty"`
	// 用户匿名字符串。如果字段“is_anonymous”设置为“true”，但是无指定匿名字符串, 会议将分配缺省名称，例如 “会议用户xxxx”，其中“xxxx”为随机数字。
	NickName *string `json:"nick_name,omitempty"`
	// 用户 ID（企业内部请使用企业唯一用户标识；OAuth2.0 鉴权用户请使用 openId）。  企业唯一用户标识说明：  企业对接 SSO 时使用的员工唯一标识 ID，企业调用创建用户接口时传递的 userid 参数。
	Userid string `json:"userid"`
}

// V1MeetingsMeetingIdPutRequestLiveConfig 直播配置对象，enable_live 为 true 时才生效。
type V1MeetingsMeetingIdPutRequestLiveConfig struct {
	// 允许观众讨论。 true：开启 false：不开启
	EnableLiveIm *bool `json:"enable_live_im,omitempty"`
	// 是否开启直播密码，默认值false.  true：开启, false：不开启
	EnableLivePassword *bool `json:"enable_live_password,omitempty"`
	// 开启直播回看。 true：开启 false：不开启
	EnableLiveReplay *bool `json:"enable_live_replay,omitempty"`
	// 直播密码。当设置开启直播密码时，该参数必填。
	LivePassword *string `json:"live_password,omitempty"`
	// 直播主题
	LiveSubject *string `json:"live_subject,omitempty"`
	// 直播简介
	LiveSummary   *string                                               `json:"live_summary,omitempty"`
	LiveWatermark *V1MeetingsMeetingIdPutRequestLiveConfigLiveWatermark `json:"live_watermark,omitempty"`
}

// V1MeetingsMeetingIdPutRequestLiveConfigLiveWatermark 直播水印对象信息
type V1MeetingsMeetingIdPutRequestLiveConfigLiveWatermark struct {
	// 水印选项，默认为0。 0：默认水印  1：无水印
	WatermarkOpt *int64 `json:"watermark_opt,omitempty"`
}

// V1MeetingsMeetingIdPutRequestRecurringRule 周期性会议配置，meeting_type 为1时必填。
type V1MeetingsMeetingIdPutRequestRecurringRule struct {
	// 哪些天重复。根据 customized_recurring_type 和 customized_recurring_step 的不同，该字段可取值与表达含义不同。如需选择多个日期，加和即可。 customized_recurring_type = 0 时，传入该字段将被忽略。 详细请参见 自定义周期规则 API 调用示例
	CustomizedRecurringDays *int64 `json:"customized_recurring_days,omitempty"`
	// 每[n]（天、周、月）重复，使用自定义周期性会议时传入。 例如：customized_recurring_type=0 && customized_recurring_step=5 表示每5天重复一次。 customized_recurring_type=2 && customized_recurring_step=3 表示每3个月重复一次，重复的时间依赖于 customized_recurring_days 字段。
	CustomizedRecurringStep *int64 `json:"customized_recurring_step,omitempty"`
	// 自定义周期性会议的循环类型。 0：按天。 1：按周。 2：按月，以周为粒度重复。例如：每3个月的第二周的周四。 3：按月，以日期为粒度重复。例如：每3个月的16日。 按周；按月、以周为粒度； 按月、以日期为粒度时，需要包含会议开始时间所在的日期。
	CustomizedRecurringType *int64 `json:"customized_recurring_type,omitempty"`
	// 重复类型，默认值为0。 0：每天 1：每周一至周五 2：每周 3：每两周 4：每月 5：自定义，示例请参见 自定义周期规则 API 调用示例
	RecurringType *int64 `json:"recurring_type,omitempty"`
	// 子会议 ID，表示修改该子会议时间，不可与周期性会议规则同时修改。 如不填写，默认修改整个周期性会议时间。
	SubMeetingId *string `json:"sub_meeting_id,omitempty"`
	// 限定会议次数（1-50次）。
	UntilCount *int64 `json:"until_count,omitempty"`
	// 结束日期时间戳，最大支持预定50场子会议。
	UntilDate *int64 `json:"until_date,omitempty"`
	// 结束重复类型。 0：按日期结束重复 1：按次数结束重复
	UntilType *int64 `json:"until_type,omitempty"`
}

// V1MeetingsMeetingIdPutRequestSettings 会议的配置，可为缺省配置。
type V1MeetingsMeetingIdPutRequestSettings struct {
	// 允许成员在主持人进会前加入会议，默认值为 true。 true：允许 false：不允许
	AllowInBeforeHost *bool `json:"allow_in_before_host,omitempty"`
	// 是否允许多端入会
	AllowMultiDevice *bool `json:"allow_multi_device,omitempty"`
	// 是否开启屏幕共享水印，默认值为 false。 true：开启 false：不开启
	AllowScreenSharedWatermark *bool `json:"allow_screen_shared_watermark,omitempty"`
	// 允许参会者取消静音，默认值为 true。 true：开启 false：关闭
	AllowUnmuteSelf *bool `json:"allow_unmute_self,omitempty"`
	// 是否开启等候室，默认值为 false。 true：开启 false：不开启
	AutoInWaitingRoom *bool `json:"auto_in_waiting_room,omitempty"`
	// 自动录制类型： none：禁用，代表不开启自动会议录制。 local：本地录制，代表主持人入会后自动开启本地录制。 cloud：云录制，代表主持人入会后自动开启云录制。 说明： 该参数依赖企业账户设置，当企业强制锁定后，该参数必须与企业配置保持一致。 仅客户端2.7及以上版本可生效。
	AutoRecordType *string `json:"auto_record_type,omitempty"`
	// 是否允许用户自己改名 1:允许用户自己改名，2:不允许用户自己改名，默认为1
	ChangeNickname *int64 `json:"change_nickname,omitempty"`
	// 允许主持人暂停或者停止云录制，默认值为 true 开启，开启时，主持人允许暂停和停止云录制；当设置为关闭时，则主持人不允许暂停和关闭云录制。 说明： 该参数必须将 auto_record_type 设置为“cloud”时才生效，该参数依赖企业账户设置，当企业强制锁定后，该参数必须与企业配置保持一致。 仅客户端2.7及以上版本生效。
	EnableHostPauseAutoRecord *bool `json:"enable_host_pause_auto_record,omitempty"`
	// 入会时静音
	MuteEnableJoin *bool `json:"mute_enable_join,omitempty"`
	// 成员入会时静音选项，默认值为2。 当同时传入“mute_enable_join”和“mute_enable_type_join”时，将以“mute_enable_type_join”的选项为准。 0：关闭 1：开启 2：超过6人后自动开启
	MuteEnableTypeJoin *int64 `json:"mute_enable_type_join,omitempty"`
	// 是否仅企业内部成员可入会，默认值为 false。 true：仅企业内部用户可入会 false：所有人可入会
	OnlyEnterpriseUserAllowed *bool `json:"only_enterprise_user_allowed,omitempty"`
	// 是否仅受邀成员可入会，默认值为false，true：仅受邀成员可入会，false：所有成员可入会
	OnlyInviteesAllowed *bool `json:"only_invitees_allowed,omitempty"`
	// 当有参会成员入会时立即开启云录制，默认值为 false 关闭，关闭时，主持人入会自动开启云录制；当设置为开启时，则有参会成员入会自动开启云录制。 说明： 该参数必须 auto_record_type 设置为“cloud”时才生效，该参数依赖企业账户设置，当企业强制锁定后，该参数必须与企业配置保持一致。 仅客户端2.7及以上版本生效。
	ParticipantJoinAutoRecord *bool `json:"participant_join_auto_record,omitempty"`
	// 有新的与会者加入时播放提示音
	PlayIvrOnJoin *bool `json:"play_ivr_on_join,omitempty"`
	// 参会者离开时播放提示音。
	PlayIvrOnLeave *bool `json:"play_ivr_on_leave,omitempty"`
	// 水印样式，默认为单排。当屏幕共享水印参数为开启时，此参数才生效。 0：单排 1：多排
	WaterMarkType *int64 `json:"water_mark_type,omitempty"`
}

// V1MeetingsMeetingIdQosGet200Response struct for V1MeetingsMeetingIdQosGet200Response
type V1MeetingsMeetingIdQosGet200Response struct {
	// 当前页码
	CurrentPage *int64 `json:"current_page,omitempty"`
	// 当前页大小
	CurrentSize *int64 `json:"current_size,omitempty"`
	// 成员列表
	Participants []V1MeetingsMeetingIdQosGet200ResponseParticipantsInner `json:"participants,omitempty"`
	// 数据总条数
	TotalCount *int64 `json:"total_count,omitempty"`
	// 总页数
	TotalPage *int64 `json:"total_page,omitempty"`
}

// V1MeetingsMeetingIdQosGet200ResponseParticipantsInner struct for V1MeetingsMeetingIdQosGet200ResponseParticipantsInner
type V1MeetingsMeetingIdQosGet200ResponseParticipantsInner struct {
	// 会中ID
	MsOpenId *string `json:"ms_open_id,omitempty"`
	// 会中用户名
	NickName   *string                                                                `json:"nick_name,omitempty"`
	QosDetails []V1MeetingsMeetingIdQosGet200ResponseParticipantsInnerQosDetailsInner `json:"qos_details,omitempty"`
}

// V1MeetingsMeetingIdQosGet200ResponseParticipantsInnerQosDetailsInner struct for V1MeetingsMeetingIdQosGet200ResponseParticipantsInnerQosDetailsInner
type V1MeetingsMeetingIdQosGet200ResponseParticipantsInnerQosDetailsInner struct {
	Audio *V1MeetingsMeetingIdQosGet200ResponseParticipantsInnerQosDetailsInnerAudio `json:"audio,omitempty"`
	// 设备类型
	Instanceid  *int64                                                                           `json:"instanceid,omitempty"`
	Network     *V1MeetingsMeetingIdQosGet200ResponseParticipantsInnerQosDetailsInnerNetwork     `json:"network,omitempty"`
	ShareScreen *V1MeetingsMeetingIdQosGet200ResponseParticipantsInnerQosDetailsInnerShareScreen `json:"share_screen,omitempty"`
	Video       *V1MeetingsMeetingIdQosGet200ResponseParticipantsInnerQosDetailsInnerVideo       `json:"video,omitempty"`
}

// V1MeetingsMeetingIdQosGet200ResponseParticipantsInnerQosDetailsInnerAudio 音频
type V1MeetingsMeetingIdQosGet200ResponseParticipantsInnerQosDetailsInnerAudio struct {
	// 下行码率（kbps）
	DownstreamBitrate *string `json:"downstream_bitrate,omitempty"`
	// 扬声器播放音量（db）
	LoudspeakerVolume *string `json:"loudspeaker_volume,omitempty"`
	// 麦克风采集音量（db）
	MicVolume *string `json:"mic_volume,omitempty"`
	// 上行码率（kbps）
	UpstreamBitrate *string `json:"upstream_bitrate,omitempty"`
}

// V1MeetingsMeetingIdQosGet200ResponseParticipantsInnerQosDetailsInnerNetwork 网络
type V1MeetingsMeetingIdQosGet200ResponseParticipantsInnerQosDetailsInnerNetwork struct {
	// 下行带宽（kbps）
	DownstreamBindwidth *string `json:"downstream_bindwidth,omitempty"`
	// 下行丢包（%）
	DownstreamPacketLoss *string `json:"downstream_packet_loss,omitempty"`
	// 网络延迟 (ms)
	NetworkDelay *string `json:"network_delay,omitempty"`
	// 上行带宽（kbps）
	UpstreamBindwidth *string `json:"upstream_bindwidth,omitempty"`
	// 上行丢包（%）
	UpstreamPacketLoss *string `json:"upstream_packet_loss,omitempty"`
}

// V1MeetingsMeetingIdQosGet200ResponseParticipantsInnerQosDetailsInnerShareScreen 共享屏幕
type V1MeetingsMeetingIdQosGet200ResponseParticipantsInnerQosDetailsInnerShareScreen struct {
	// 下行码率（kbps）
	DownstreamBitrate *string `json:"downstream_bitrate,omitempty"`
	// 下行帧率（fps）
	DownstreamFramerate *string `json:"downstream_framerate,omitempty"`
	// 下行分辨率
	DownstreamResolution *string `json:"downstream_resolution,omitempty"`
	// 上行码率（kbps）
	UpstreamBitrate *string `json:"upstream_bitrate,omitempty"`
	// 上行帧率（fps）
	UpstreamFramerate *string `json:"upstream_framerate,omitempty"`
	// 上行分辨率
	UpstreamResolution *string `json:"upstream_resolution,omitempty"`
}

// V1MeetingsMeetingIdQosGet200ResponseParticipantsInnerQosDetailsInnerVideo 视频
type V1MeetingsMeetingIdQosGet200ResponseParticipantsInnerQosDetailsInnerVideo struct {
	// 下行码率（kbps）
	DownstreamBitrate *string `json:"downstream_bitrate,omitempty"`
	// 下行帧率（fps）
	DownstreamFramerate *string `json:"downstream_framerate,omitempty"`
	// 下行分辨率
	DownstreamResolution *string `json:"downstream_resolution,omitempty"`
	// 上行码率（kbps）
	UpstreamBitrate *string `json:"upstream_bitrate,omitempty"`
	// 上行帧率（fps）
	UpstreamFramerate *string `json:"upstream_framerate,omitempty"`
	// 上行分辨率
	UpstreamResolution *string `json:"upstream_resolution,omitempty"`
}

// V1MeetingsMeetingIdQualityGet200Response struct for V1MeetingsMeetingIdQualityGet200Response
type V1MeetingsMeetingIdQualityGet200Response struct {
	// 音频质量：0-无数据，1-好、2-较好、3-中、4-差
	AudioQuality *int64 `json:"audio_quality,omitempty"`
	// 分页查询返回当前页码
	CurrentPage *int64 `json:"current_page,omitempty"`
	// 分页查询返回单页数据条数
	CurrentSize *int64 `json:"current_size,omitempty"`
	// 会议的唯一 ID
	MeetingId *string `json:"meeting_id,omitempty"`
	// 网络质量：0-无数据，1-好、2-较好、3-中、4-差
	NetworkQuality *int64 `json:"network_quality,omitempty"`
	// 参会人员健康度对象数组（按成员入会时间正序排列，入会越早的在越上面；成员使用不同端入会时平铺返回数据，instanceid不同）
	Participants []V1MeetingsMeetingIdQualityGet200ResponseParticipantsInner `json:"participants,omitempty"`
	// 告警的具体问题
	Problems []string `json:"problems,omitempty"`
	// 健康度：0-无数据，1-健康、2-告警
	Quality *int64 `json:"quality,omitempty"`
	// 共享屏幕质量：0-无数据，1-好、2-较好、3-中、4-差
	ScreenShareQuality *int64 `json:"screen_share_quality,omitempty"`
	// 分页查询返回数据总数
	TotalCount *int64 `json:"total_count,omitempty"`
	// 分页查询返回分页总数
	TotalPage *int64 `json:"total_page,omitempty"`
	// 视频质量：0-无数据，1-好、2-较好、3-中、4-差
	VideoQuality *int64 `json:"video_quality,omitempty"`
}

// V1MeetingsMeetingIdQualityGet200ResponseParticipantsInner struct for V1MeetingsMeetingIdQualityGet200ResponseParticipantsInner
type V1MeetingsMeetingIdQualityGet200ResponseParticipantsInner struct {
	// 音频质量：0-无数据，1-好、2-较好、3-中、4-差
	AudioQuality *int64 `json:"audio_quality,omitempty"`
	// 用户的终端设备类型： 0：PSTN 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch MacOS 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch iOS
	Instanceid *string `json:"instanceid,omitempty"`
	// 当场会议的用户临时 ID，可用于会控操作，适用于所有用户。
	MsOpenId *string `json:"ms_open_id,omitempty"`
	// 网络质量：0-无数据，1-好、2-较好、3-中、4-差
	NetworkQuality *int64 `json:"network_quality,omitempty"`
	// OAuth2.0 鉴权用户请使用 open_id。
	OpenId *string `json:"open_id,omitempty"`
	// 告警的具体问题
	Problems []string `json:"problems,omitempty"`
	// 用户健康度：0-无数据，1-健康、2-告警
	Quality *int64 `json:"quality,omitempty"`
	// 共享屏幕质量：0-无数据，1-好、2-较好、3-中、4-差
	ScreenShareQuality *int64 `json:"screen_share_quality,omitempty"`
	// 同企业内部请使用企业唯一用户标识； 其他企业，个人，小程序没有 。
	Userid *string `json:"userid,omitempty"`
	// 视频质量：0-无数据，1-好、2-较好、3-中、4-差
	VideoQuality *int64 `json:"video_quality,omitempty"`
}

// V1MeetingsMeetingIdRealTimeParticipantsGet200Response struct for V1MeetingsMeetingIdRealTimeParticipantsGet200Response
type V1MeetingsMeetingIdRealTimeParticipantsGet200Response struct {
	// 当前页。
	CurrentPage *int64 `json:"current_page,omitempty"`
	// 当前页实际大小。
	CurrentSize *int64 `json:"current_size,omitempty"`
	// 会议号码。
	MeetingCode *string `json:"meeting_code,omitempty"`
	// 会议的唯一 ID。
	MeetingId *string `json:"meeting_id,omitempty"`
	// 参会人员对象数组。
	Participants []V1MeetingsMeetingIdRealTimeParticipantsGet200ResponseParticipantsInner `json:"participants,omitempty"`
	// 预定会议结束时间戳（单位秒）。
	ScheduleEndTime *string `json:"schedule_end_time,omitempty"`
	// 预定会议开始时间戳（单位秒）。
	ScheduleStartTime *string `json:"schedule_start_time,omitempty"`
	// 当前会议状态： 1. MEETING_STATE_INVALID： 非法或未知的会议状态，错误状态。 2. MEETING_STATE_INIT： 会议待开始。会议预定到预定结束时间前，会议尚无人进会。 3. MEETING_STATE_CANCELLED： 会议已取消。主持人主动取消会议，待开始的会议才能取消，且取消的会议无法再进入。 4. MEETING_STATE_STARTED： 会议已开始。会议中有人则表示会议进行中。 5. MEETING_STATE_ENDED： 会议已删除。会议已过预定结束时间且尚无人进会时，主持人删除会议，已删除的会议无法再进入。 6. MEETING_STATE_NULL： 会议无状态。会议已过预定结束时间，会议尚无人进会。 7. MEETING_STATE_RECYCLED： 会议已回收。会议已过预定开始时间30天，则会议号将被后台回收，无法再进入。
	Status *string `json:"status,omitempty"`
	// 会议主题（base64）。
	Subject *string `json:"subject,omitempty"`
	// 根据条件筛选后的总人数。
	TotalCount *int64 `json:"total_count,omitempty"`
	// 根据条件筛选后的总分页数
	TotalPage *int64 `json:"total_page,omitempty"`
}

// V1MeetingsMeetingIdRealTimeParticipantsGet200ResponseParticipantsInner struct for V1MeetingsMeetingIdRealTimeParticipantsGet200ResponseParticipantsInner
type V1MeetingsMeetingIdRealTimeParticipantsGet200ResponseParticipantsInner struct {
	// 用户的客户端版本。当用户在会中时才能返回。
	AppVersion *string `json:"app_version,omitempty"`
	// 麦克风状态： true：开启 false：关闭
	AudioState   *bool   `json:"audio_state,omitempty"`
	CustomerData *string `json:"customer_data,omitempty"`
	// 用户的终端设备类型： 0:pstn或mra 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序
	Instanceid *int64 `json:"instanceid,omitempty"`
	// 参会者加入会议时间戳（单位秒）。
	JoinTime *string `json:"join_time,omitempty"`
	// 入会方式： 0：PSTN 普通用户，标准的手机或固话类型 1：普通 VOIP 用户 2：附属投屏 VOIP 3：linux sdk for VOIP 4：附属语音 PSTN 5：附属视频 PSTN 6：linux sdk for PSTN
	JoinType *int64 `json:"join_type,omitempty"`
	// 当场会议的用户临时 ID，可用于会控操作，适用于所有用户。
	MsOpenId *string `json:"ms_open_id,omitempty"`
	// OAuth2.0 鉴权用户请使用 open_id
	OpenId *string `json:"open_id,omitempty"`
	// 屏幕共享状态： true：开启 false：关闭
	ScreenSharedState *bool `json:"screen_shared_state,omitempty"`
	// 入会用户名（base64）。
	UserName *string `json:"user_name,omitempty"`
	// 用户角色： 0：普通成员角色 1：创建者角色 2：主持人 3：创建者+主持人 4：游客 5：游客+主持人 6：联席主持人 7：创建者+联席主持人
	UserRole *int64 `json:"user_role,omitempty"`
	// 同企业内部请使用企业唯一用户标识； 其他企业，个人，小程序没有
	Userid *string `json:"userid,omitempty"`
	// 摄像头状态： true：开启 false：关闭
	VideoState *bool `json:"video_state,omitempty"`
}

// V1MeetingsMeetingIdVirtualBackgroundPost200Response struct for V1MeetingsMeetingIdVirtualBackgroundPost200Response
type V1MeetingsMeetingIdVirtualBackgroundPost200Response struct {
	// 任务ID
	JobId *string `json:"job_id,omitempty"`
}

// V1MeetingsMeetingIdVirtualBackgroundPostRequest struct for V1MeetingsMeetingIdVirtualBackgroundPostRequest
type V1MeetingsMeetingIdVirtualBackgroundPostRequest struct {
	// 背景图片 url，尺寸大小为1920*1080，小于10M，jpg/png/jpeg 格式如果不传入则使用会议室默认虚拟背景。
	Image *string `json:"image,omitempty"`
	// 用户的终端设备类型： 0：PSTN 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：Linux 20：Rooms for Touch Windows 21：Rooms for Touch MacOS 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch iOS
	Instanceid int64 `json:"instanceid"`
	// 操作者ID
	OperatorId string `json:"operator_id"`
	// 操作者ID类型。1-userid，2-opened
	OperatorIdType int64 `json:"operator_id_type"`
	// 背景生效类型。0-全部用户，1-部分用户
	Type int64 `json:"type"`
	// userid数组
	Users []string `json:"users,omitempty"`
}

// V1MeetingsMeetingIdWaitingRoomParticipantsGet200Response struct for V1MeetingsMeetingIdWaitingRoomParticipantsGet200Response
type V1MeetingsMeetingIdWaitingRoomParticipantsGet200Response struct {
	// 分页查询返回当前页码
	CurrentPage *int64 `json:"current_page,omitempty"`
	// 分页查询返回单页数据条数
	CurrentSize *int64 `json:"current_size,omitempty"`
	// 会议CODE
	MeetingCode *string `json:"meeting_code,omitempty"`
	// 会议唯一 ID
	MeetingId *string `json:"meeting_id,omitempty"`
	// 等候室人员对象数组
	Participants []V1MeetingsMeetingIdWaitingRoomParticipantsGet200ResponseParticipantsInner `json:"participants,omitempty"`
	// 预定的会议结束时间戳（单位秒）
	ScheduleEndTime *int64 `json:"schedule_end_time,omitempty"`
	// 预定的会议开始时间戳（单位秒）
	ScheduleStartTime *int64 `json:"schedule_start_time,omitempty"`
	// 会议主题 (base64 编码)
	Subject *string `json:"subject,omitempty"`
	// 分页查询返回数据总数
	TotalCount *int64 `json:"total_count,omitempty"`
	// 分页查询返回分页总数
	TotalPage *int64 `json:"total_page,omitempty"`
}

// V1MeetingsMeetingIdWaitingRoomParticipantsGet200ResponseParticipantsInner struct for V1MeetingsMeetingIdWaitingRoomParticipantsGet200ResponseParticipantsInner
type V1MeetingsMeetingIdWaitingRoomParticipantsGet200ResponseParticipantsInner struct {
	// 客户端版本
	AppVersion *string `json:"app_version,omitempty"`
	// 用户自定义参数
	CustomerData *string `json:"customer_data,omitempty"`
	// 用户的终端设备类型： 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch Mac 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch Iphone
	Instanceid *int64 `json:"instanceid,omitempty"`
	// 成员会中唯一 ID。
	MsOpenId *string `json:"ms_open_id,omitempty"`
	// 等候室 OAuth2.0 鉴权用户的 ID。
	OpenId *string `json:"open_id,omitempty"`
	// 入会用户名（base64）
	UserName *string `json:"user_name,omitempty"`
	// 等候室成员用户 ID
	Userid *string `json:"userid,omitempty"`
	// 用户的唯一标识uuid
	Uuid *string `json:"uuid,omitempty"`
}

// V1MeetingsPost200Response struct for V1MeetingsPost200Response
type V1MeetingsPost200Response struct {
	// 会议列表
	MeetingInfoList []V1MeetingsPost200ResponseMeetingInfoListInner `json:"meeting_info_list,omitempty"`
	// 会议数量
	MeetingNumber *int64 `json:"meeting_number,omitempty"`
}

// V1MeetingsPost200ResponseMeetingInfoListInner 会议列表
type V1MeetingsPost200ResponseMeetingInfoListInner struct {
	// 联席主持人
	CurrentCoHosts []V1MeetingsPost200ResponseMeetingInfoListInnerCurrentCoHostsInner `json:"current_co_hosts,omitempty"`
	// 是否开启直播（会议创建人才有权限查询）
	EnableLive *bool `json:"enable_live,omitempty"`
	// 会议结束时间戳（秒）
	EndTime *string `json:"end_time,omitempty"`
	// 主持人密钥，仅支持6位数字。
	HostKey *string `json:"host_key,omitempty"`
	// 指定主持人列表，仅商业版和企业版可指定主持人
	Hosts []V1MeetingsPost200ResponseMeetingInfoListInnerHostsInner `json:"hosts,omitempty"`
	// 加入会议 URL
	JoinUrl    *string                                                  `json:"join_url,omitempty"`
	LiveConfig *V1MeetingsPost200ResponseMeetingInfoListInnerLiveConfig `json:"live_config,omitempty"`
	// 会议 App 的呼入号码
	MeetingCode *string `json:"meeting_code,omitempty"`
	// 会议唯一id
	MeetingId *string `json:"meeting_id,omitempty"`
	// 邀请的参会人
	Participants []V1MeetingsPost200ResponseMeetingInfoListInnerParticipantsInner `json:"participants,omitempty"`
	// 会议密码
	Password *string                                                `json:"password,omitempty"`
	Settings *V1MeetingsPost200ResponseMeetingInfoListInnerSettings `json:"settings,omitempty"`
	// 会议开始时间戳（秒）
	StartTime *string `json:"start_time,omitempty"`
	// 会议主题
	Subject *string `json:"subject,omitempty"`
	// 邀请的参会者中未注册用户。注意：仅腾讯会议商业版和企业版可获取该参数。
	UserNonRegistered []string `json:"user_non_registered,omitempty"`
}

// V1MeetingsPost200ResponseMeetingInfoListInnerCurrentCoHostsInner 联席主持人
type V1MeetingsPost200ResponseMeetingInfoListInnerCurrentCoHostsInner struct {
	// 用户是否匿名入会，缺省为 false，不匿名。 true：匿名 false：不匿名
	IsAnonymous *bool `json:"is_anonymous,omitempty"`
	// 用户匿名字符串。如果字段“is_anonymous”设置为“true”，但是无指定匿名字符串, 会议将分配缺省名称，例如 “会议用户xxxx”，其中“xxxx”为随机数字
	NickName *string `json:"nick_name,omitempty"`
	// 操作者ID，根据operator_id_type的值，使用不同的类型
	OperatorId *string `json:"operator_id,omitempty"`
	// 操作者ID的类型：1:userid  2:openid 3:rooms_id  4: ms_open_id
	OperatorIdType *int64 `json:"operator_id_type,omitempty"`
	// 头像地址
	ProfilePhoto *string `json:"profile_photo,omitempty"`
	Userid       *string `json:"userid,omitempty"`
}

// V1MeetingsPost200ResponseMeetingInfoListInnerHostsInner 指定主持人列表，仅商业版和企业版可指定主持人
type V1MeetingsPost200ResponseMeetingInfoListInnerHostsInner struct {
	// 用户是否匿名入会，缺省为 false，不匿名。 true：匿名 false：不匿名
	IsAnonymous *bool `json:"is_anonymous,omitempty"`
	// 用户匿名字符串。如果字段“is_anonymous”设置为“true”，但是无指定匿名字符串, 会议将分配缺省名称，例如 “会议用户xxxx”，其中“xxxx”为随机数字
	NickName *string `json:"nick_name,omitempty"`
	// 操作者ID，根据operator_id_type的值，使用不同的类型
	OperatorId *string `json:"operator_id,omitempty"`
	// 操作者ID的类型：1:userid  2:openid 3:rooms_id  4: ms_open_id
	OperatorIdType *int64 `json:"operator_id_type,omitempty"`
	// 头像地址
	ProfilePhoto *string `json:"profile_photo,omitempty"`
	Userid       *string `json:"userid,omitempty"`
}

// V1MeetingsPost200ResponseMeetingInfoListInnerLiveConfig 会议的直播配置（会议创建人才有权限查询）
type V1MeetingsPost200ResponseMeetingInfoListInnerLiveConfig struct {
	// 是否开启直播互动
	EnableLiveIm *bool `json:"enable_live_im,omitempty"`
	// 是否开启直播密码，默认值false. true：开启, false：不开启
	EnableLivePassword *bool `json:"enable_live_password,omitempty"`
	// 是否开启直播回放
	EnableLiveReplay *bool `json:"enable_live_replay,omitempty"`
	// 直播密码
	LivePassword *string `json:"live_password,omitempty"`
	// 直播主题
	LiveSubject *string `json:"live_subject,omitempty"`
	// 直播简介
	LiveSummary   *string                                                               `json:"live_summary,omitempty"`
	LiveWatermark *V1MeetingsPost200ResponseMeetingInfoListInnerLiveConfigLiveWatermark `json:"live_watermark,omitempty"`
}

// V1MeetingsPost200ResponseMeetingInfoListInnerLiveConfigLiveWatermark 直播水印对象信息
type V1MeetingsPost200ResponseMeetingInfoListInnerLiveConfigLiveWatermark struct {
	// 水印选项，默认为0。 0：默认水印 1：无水印
	WatermarkOpt *int64 `json:"watermark_opt,omitempty"`
}

// V1MeetingsPost200ResponseMeetingInfoListInnerParticipantsInner 邀请的参会人
type V1MeetingsPost200ResponseMeetingInfoListInnerParticipantsInner struct {
	// 用户是否匿名入会，缺省为 false，不匿名。 true：匿名 false：不匿名
	IsAnonymous *bool `json:"is_anonymous,omitempty"`
	// 用户匿名字符串。如果字段“is_anonymous”设置为“true”，但是无指定匿名字符串, 会议将分配缺省名称，例如 “会议用户xxxx”，其中“xxxx”为随机数字
	NickName *string `json:"nick_name,omitempty"`
	// 操作者ID，根据operator_id_type的值，使用不同的类型
	OperatorId *string `json:"operator_id,omitempty"`
	// 操作者ID的类型：1:userid  2:openid 3:rooms_id  4: ms_open_id
	OperatorIdType *int64 `json:"operator_id_type,omitempty"`
	// 头像地址
	ProfilePhoto *string `json:"profile_photo,omitempty"`
	Userid       *string `json:"userid,omitempty"`
}

// V1MeetingsPost200ResponseMeetingInfoListInnerSettings 会议媒体设置参数
type V1MeetingsPost200ResponseMeetingInfoListInnerSettings struct {
	// 允许成员在主持人进会前加入会议
	AllowInBeforeHost *bool `json:"allow_in_before_host,omitempty"`
	// 是否允许多端入会
	AllowMultiDevice *bool `json:"allow_multi_device,omitempty"`
	// 开启屏幕共享水印
	AllowScreenSharedWatermark *bool `json:"allow_screen_shared_watermark,omitempty"`
	// 静音自解除允许
	AllowUnmuteSelf *bool `json:"allow_unmute_self,omitempty"`
	// 开启等候室
	AutoInWaitingRoom *bool `json:"auto_in_waiting_room,omitempty"`
	// 自动录制类型，仅客户端2.7及以上版本生效 none：禁用 local：本地录制 cloud：云录制
	AutoRecordType *string `json:"auto_record_type,omitempty"`
	// 是否允许用户自己改名 1:允许用户自己改名，2:不允许用户自己改名，默认为1
	ChangeNickname *int64 `json:"change_nickname,omitempty"`
	// 允许主持人暂停或者停止云录制，默认值为 true 开启，开启时，主持人允许暂停和停止云录制；当设置为关闭时，则主持人不允许暂停和关闭云录制
	EnableHostPauseAutoRecord *bool `json:"enable_host_pause_auto_record,omitempty"`
	// 加入静音状态
	MuteEnableJoin *bool `json:"mute_enable_join,omitempty"`
	// 加入静音状态
	MuteEnableTypeJoin *int64 `json:"mute_enable_type_join,omitempty"`
	// 是否仅企业内部成员可入会，默认值为 false。true：仅企业内部用户可入会 false：所有人可入会
	OnlyEnterpriseUserAllowed *bool `json:"only_enterprise_user_allowed,omitempty"`
	// 成员入会限制，1：所有成员可入会，2：仅受邀成员可入会，3：仅企业内部成员可入会 ；当only_user_join_type和only_allow_enterprise_user_join同时传的时候，以only_user_join_type为准
	OnlyUserJoinType *int64 `json:"only_user_join_type,omitempty"`
	// 当有参会成员入会时立即开启云录制，默认值为 false 关闭，关闭时，主持人入会自动开启云录制；当设置为开启时，则有参会成员入会自动开启云录制。
	ParticipantJoinAutoRecord *bool `json:"participant_join_auto_record,omitempty"`
	// 有新的与会者加入时播放提示音，暂不支持，可在客户端设置
	PlayIvrOnJoin *bool `json:"play_ivr_on_join,omitempty"`
	// 参会者离开时播放提示音，暂时不支持，可在客户端设置。
	PlayIvrOnLeave *bool `json:"play_ivr_on_leave,omitempty"`
	// 水印样式，默认为单排：0：单排 1：多排
	WaterMarkType *int64 `json:"water_mark_type,omitempty"`
}

// V1MeetingsPostRequest struct for V1MeetingsPostRequest
type V1MeetingsPostRequest struct {
	// 默认是false, 如果操作者无创建转网会议的权限，则该字段忽略
	AllowEnterpriseIntranetOnly *bool `json:"allow_enterprise_intranet_only,omitempty"`
	// 是否允许成员上传文档，默认为允许 true：允许 false：不允许
	EnableDocUploadPermission *bool `json:"enable_doc_upload_permission,omitempty"`
	// 是否开启报名开关，默认不开启 true：开启 false：不开启
	EnableEnroll *bool `json:"enable_enroll,omitempty"`
	// 是否开启主持人密钥，默认为false。 true：开启 false：关闭
	EnableHostKey *bool `json:"enable_host_key,omitempty"`
	// 是否开启同声传译，默认不开启 false：不开启 true：开启同声传译
	EnableInterpreter *bool `json:"enable_interpreter,omitempty"`
	// 是否开启直播
	EnableLive *bool `json:"enable_live,omitempty"`
	// 会议结束时间戳（单位秒）
	EndTime string `json:"end_time"`
	// 会议嘉宾列表，会议嘉宾不受会议密码和等候室的限制
	Guests []V1MeetingsPostRequestGuestsInner `json:"guests,omitempty"`
	// 主持人密钥，仅支持6位数字。 如开启主持人密钥后没有填写此项，将自动分配一个6位数字的密钥。
	HostKey *string `json:"host_key,omitempty"`
	// 主持人列表，会议指定主持人的用户 ID，如果无指定，主持人将被设定为参数 userid 的用户，即 API 调用者。 注意：仅腾讯会议商业版和企业版可指定主持人。
	Hosts []V1MeetingsPostRequestHostsInner `json:"hosts,omitempty"`
	// 用户的终端设备类型： 0：PSTN 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch MacOS 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch iOS 创建会议时 userid 对应的设备类型，不影响入会时使用的设备类型，缺省可填1。
	Instanceid int64 `json:"instanceid"`
	// 邀请人列表 仅支持邀请与会议创建者同企业的成员（企业内部请使用企业唯一用户标识；OAuth2.0 鉴权用户请使用 openId），该会议将添加至邀请成员的会议列表中。 企业唯一用户标识说明： 企业对接 SSO 时使用的员工唯一标识 ID。 企业调用创建用户接口时传递的 userid 参数。 注意：仅腾讯会议商业版和企业版可邀请参会者，邀请者列表仅支持300人；需要邀请超过300人的场景请调用 设置会议邀请成员 接口。
	Invitees   []V1MeetingsPostRequestInviteesInner `json:"invitees,omitempty"`
	LiveConfig *V1MeetingsPostRequestLiveConfig     `json:"live_config,omitempty"`
	// 会议地点。最长支持18个汉字或36个英文字母
	Location *string `json:"location,omitempty"`
	// 该参数仅提供给支持混合云的企业可见，默认值为0 0：外部会议 1：内部会议
	MediaSetType *int64 `json:"media_set_type,omitempty"`
	// 默认值为0。 0：普通会议 1：周期性会议（周期性会议时 type 不能为快速会议，同一账号同时最多可预定50场周期性会议）
	MeetingType *int64 `json:"meeting_type,omitempty"`
	// 会议密码（4~6位数字），可不填
	Password      *string                             `json:"password,omitempty"`
	RecurringRule *V1MeetingsPostRequestRecurringRule `json:"recurring_rule,omitempty"`
	Settings      *V1MeetingsPostRequestSettings      `json:"settings,omitempty"`
	// 会议开始时间戳（单位秒）
	StartTime string `json:"start_time"`
	// 会议主题
	Subject string `json:"subject"`
	// 会议是否同步至企业微信，该字段仅支持创建会议时设置，创建后无法修改。 true: 同步，默认同步  false: 不同步
	SyncToWework *bool `json:"sync_to_wework,omitempty"`
	// 时区，可参见 Oracle-TimeZone 标准
	TimeZone *string `json:"time_zone,omitempty"`
	// 会议类型 0：预约会议 1：快速会议
	Type int64 `json:"type"`
	// 调用方用于标示用户的唯一 ID（企业内部请使用企业唯一用户标识；OAuth2.0 鉴权用户请使用 openId）。 企业唯一用户标识说明： 1. 企业对接 SSO 时使用的员工唯一标识 ID； 2. 企业调用创建用户接口时传递的 userid 参数。
	Userid string `json:"userid"`
}

// V1MeetingsPostRequestGuestsInner struct for V1MeetingsPostRequestGuestsInner
type V1MeetingsPostRequestGuestsInner struct {
	// 国家/地区代码（例如：中国传86，不是+86，也不是0086）
	Area string `json:"area"`
	// 嘉宾名称
	GuestName *string `json:"guest_name,omitempty"`
	// 手机号
	PhoneNumber string `json:"phone_number"`
}

// V1MeetingsPostRequestHostsInner struct for V1MeetingsPostRequestHostsInner
type V1MeetingsPostRequestHostsInner struct {
	// 用户是否匿名入会，缺省为 false，不匿名 true：匿名 false：不匿名
	IsAnonymous *bool `json:"is_anonymous,omitempty"`
	// 用户匿名字符串。如果字段“is_anonymous”设置为“true”，但是无指定匿名字符串, 会议将分配缺省名称，例如 “会议用户xxxx”，其中“xxxx”为随机数字。
	NickName *string `json:"nick_name,omitempty"`
	// 用户 ID（企业内部请使用企业唯一用户标识；OAuth2.0 鉴权用户请使用 openId） 企业唯一用户标识说明： 企业对接 SSO 时使用的员工唯一标识 ID，企业调用创建用户接口时传递的 userid 参数。
	Userid string `json:"userid"`
}

// V1MeetingsPostRequestInviteesInner struct for V1MeetingsPostRequestInviteesInner
type V1MeetingsPostRequestInviteesInner struct {
	// 用户是否匿名入会，缺省为 false，不匿名。 true：匿名 false：不匿名
	IsAnonymous *bool `json:"is_anonymous,omitempty"`
	// 用户匿名字符串。如果字段“is_anonymous”设置为“true”，但是无指定匿名字符串, 会议将分配缺省名称，例如 “会议用户xxxx”，其中“xxxx”为随机数字。
	NickName *string `json:"nick_name,omitempty"`
	// 用户 ID（企业内部请使用企业唯一用户标识；OAuth2.0 鉴权用户请使用 openId）。 企业唯一用户标识说明： 企业对接 SSO 时使用的员工唯一标识 ID，企业调用创建用户接口时传递的 userid 参数。
	Userid string `json:"userid"`
}

// V1MeetingsPostRequestLiveConfig 直播配置
type V1MeetingsPostRequestLiveConfig struct {
	// 允许观众讨论，默认值为 false。 true：开启 false：不开启
	EnableLiveIm *bool `json:"enable_live_im,omitempty"`
	// 是否开启直播密码，默认值false. true：开启, false：不开启
	EnableLivePassword *bool `json:"enable_live_password,omitempty"`
	// 开启直播回看，默认值为 false true：开启 false：不开启
	EnableLiveReplay *bool `json:"enable_live_replay,omitempty"`
	// 直播密码。当设置开启直播密码时，该参数必填。
	LivePassword *string `json:"live_password,omitempty"`
	// 直播主题
	LiveSubject *string `json:"live_subject,omitempty"`
	// 直播简介
	LiveSummary   *string                                       `json:"live_summary,omitempty"`
	LiveWatermark *V1MeetingsPostRequestLiveConfigLiveWatermark `json:"live_watermark,omitempty"`
}

// V1MeetingsPostRequestLiveConfigLiveWatermark 直播水印对象
type V1MeetingsPostRequestLiveConfigLiveWatermark struct {
	// 水印选项，默认为0。 0：默认水印 1：无水印
	WatermarkOpt *int64 `json:"watermark_opt,omitempty"`
}

// V1MeetingsPostRequestRecurringRule 周期性会议配置
type V1MeetingsPostRequestRecurringRule struct {
	// 哪些天重复。根据 customized_recurring_type 和 customized_recurring_step 的不同，该字段可取值与表达含义不同。如需选择多个日期，加和即可。 customized_recurring_type = 0 时，传入该字段将被忽略。 详细请参见 自定义周期规则 API 调用示例
	CustomizedRecurringDays *int64 `json:"customized_recurring_days,omitempty"`
	// 每[n]（天、周、月）重复，使用自定义周期性会议时传入。 例如：customized_recurring_type=0 &amp;&amp; customized_recurring_step=5 表示每5天重复一次。 customized_recurring_type=2 &amp;&amp; customized_recurring_step=3 表示每3个月重复一次，重复的时间依赖于 customized_recurring_days 字段。
	CustomizedRecurringStep *int64 `json:"customized_recurring_step,omitempty"`
	// 自定义周期性会议的循环类型。 0：按天。 1：按周。 2：按月，以周为粒度重复。例如：每3个月的第二周的周四。 3：按月，以日期为粒度重复。例如：每3个月的16日。 按周；按月、以周为粒度； 按月、以日期为粒度时，需要包含会议开始时间所在的日期。
	CustomizedRecurringType *int64 `json:"customized_recurring_type,omitempty"`
	// 重复类型，默认值为0。 0：每天 1：每周一至周五 2：每周 3：每两周 4：每月 5：自定义，示例请参见 自定义周期规则 API 调用示例
	RecurringType *int64 `json:"recurring_type,omitempty"`
	// 限定会议次数。 说明：每天、每个工作日、每周最大支持200场子会议；每两周、每月最大支持50场子会议。如未填写，则默认为7次。
	UntilCount *int64 `json:"until_count,omitempty"`
	// 结束日期时间戳。 说明：结束日期与第一场会议的开始时间换算成的场次数不能超过以下限制：每天、每个工作日、每周最大支持200场子会议；每两周、每月最大支持50场子会议，例如：对于每天的重复类型，第一场会议开始时间为1609430400，则结束日期时间戳不能超过1609430400 + 200 × 24 × 60 × 60 - 1。如未填写，默认为当前日期往后推7天。
	UntilDate *int64 `json:"until_date,omitempty"`
	// 结束重复类型，默认值为0。  0：按日期结束重复  1：按次数结束重复
	UntilType *int64 `json:"until_type,omitempty"`
}

// V1MeetingsPostRequestSettings 会议媒体参数配置
type V1MeetingsPostRequestSettings struct {
	// 是否允许成员在主持人进会前加入会议，默认值为 true。 true：允许 false：不允许
	AllowInBeforeHost *bool `json:"allow_in_before_host,omitempty"`
	// 是否允许多端入会
	AllowMultiDevice *bool `json:"allow_multi_device,omitempty"`
	// 是否开启屏幕共享水印，默认值为 false。 true： 开启 false：不开启
	AllowScreenSharedWatermark *bool `json:"allow_screen_shared_watermark,omitempty"`
	// 允许参会者取消静音，默认值为 true。 true：开启 false：关闭
	AllowUnmuteSelf *bool `json:"allow_unmute_self,omitempty"`
	// 是否开启等候室，默认值为 false。 true：开启 false：不开启
	AutoInWaitingRoom *bool `json:"auto_in_waiting_room,omitempty"`
	// 自动会议录制类型。 none：禁用，表示不开启自动会议录制。 local：本地录制，表示主持人入会后自动开启本地录制。 cloud：云录制，表示主持人入会后自动开启云录制。 说明： 该参数依赖企业账户设置，当企业强制锁定后，该参数必须与企业配置保持一致。 仅客户端2.7及以上版本可生效。
	AutoRecordType *string `json:"auto_record_type,omitempty"`
	// 是否允许用户自己改名 1:允许用户自己改名，2:不允许用户自己改名，默认为1
	ChangeNickname *int64 `json:"change_nickname,omitempty"`
	// 允许主持人暂停或者停止云录制，默认值为 true 开启，开启时，主持人允许暂停和停止云录制；当设置为关闭时，则主持人不允许暂停和关闭云录制。 说明： 该参数必须 auto_record_type 设置为“cloud”时才生效，该参数依赖企业账户设置，当企业强制锁定后，该参数必须与企业配置保持一致。 仅客户端2.7及以上版本生效。
	EnableHostPauseAutoRecord *bool `json:"enable_host_pause_auto_record,omitempty"`
	// 入会时静音，默认值为 true true：开启 false：关闭
	MuteEnableJoin *bool `json:"mute_enable_join,omitempty"`
	// 成员入会时静音选项，默认值为2。 当同时传入“mute_enable_join”和“mute_enable_type_join”时，将以“mute_enable_type_join”的选项为准。 0：关闭 1：开启 2：超过6人后自动开启
	MuteEnableTypeJoin *int64 `json:"mute_enable_type_join,omitempty"`
	// 是否仅企业内部成员可入会，默认值为 false。 true：仅企业内部用户可入会 false：所有人可入会
	OnlyEnterpriseUserAllowed *bool `json:"only_enterprise_user_allowed,omitempty"`
	// 成员入会限制，1：所有成员可入会，2：仅受邀成员可入会，3：仅企业内部成员可入会 ；当only_user_join_type和only_allow_enterprise_user_join同时传的时候，以only_user_join_type为准
	OnlyUserJoinType *int64 `json:"only_user_join_type,omitempty"`
	// 当有参会成员入会时立即开启云录制，默认值为 false 关闭，关闭时，主持人入会自动开启云录制；当设置为开启时，则有参会成员入会自动开启云录制。 说明： 该参数必须 auto_record_type 设置为“cloud”时才生效，该参数依赖企业账户设置，当企业强制锁定后，该参数必须与企业配置保持一致。 仅客户端2.7及以上版本生效。
	ParticipantJoinAutoRecord *bool `json:"participant_join_auto_record,omitempty"`
	// 有新的与会者加入时播放提示音，暂不支持，可在客户端设置
	PlayIvrOnJoin *bool `json:"play_ivr_on_join,omitempty"`
	// 参会者离开时播放提示音，暂时不支持，可在客户端设置。
	PlayIvrOnLeave *bool `json:"play_ivr_on_leave,omitempty"`
	// 水印样式，默认为单排。当屏幕共享水印参数为开启时，此参数才生效。 0：单排 1：多排
	WaterMarkType *int64 `json:"water_mark_type,omitempty"`
}

// V1MeetingsQueryMeetingidForDevicePost200Response struct for V1MeetingsQueryMeetingidForDevicePost200Response
type V1MeetingsQueryMeetingidForDevicePost200Response struct {
	// 用户设备已加入的会议的列表
	MeetingIdMap []V1MeetingsQueryMeetingidForDevicePost200ResponseMeetingIdMapInner `json:"meeting_id_map,omitempty"`
}

// V1MeetingsQueryMeetingidForDevicePost200ResponseMeetingIdMapInner struct for V1MeetingsQueryMeetingidForDevicePost200ResponseMeetingIdMapInner
type V1MeetingsQueryMeetingidForDevicePost200ResponseMeetingIdMapInner struct {
	// 终端设备类型。
	Instanceid *int64 `json:"instanceid,omitempty"`
	// 会议ID。
	MeetingId *string `json:"meeting_id,omitempty"`
}

// V1MeetingsQueryMeetingidForDevicePostRequest struct for V1MeetingsQueryMeetingidForDevicePostRequest
type V1MeetingsQueryMeetingidForDevicePostRequest struct {
	// 操作者 ID，即查询者的信息。 operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。
	OperatorId string `json:"operator_id"`
	// 操作者 ID 的类型：  1：企业内用户 userid。JWT鉴权仅支持userid
	OperatorIdType int64 `json:"operator_id_type"`
}

// V1PmiMeetingsGet200Response struct for V1PmiMeetingsGet200Response
type V1PmiMeetingsGet200Response struct {
	// 当前页。
	CurrentPage *int64 `json:"current_page,omitempty"`
	// 当前实际页大小。
	CurrentSize *int64 `json:"current_size,omitempty"`
	// 会议列表。
	MeetingInfoList []V1PmiMeetingsGet200ResponseMeetingInfoListInner `json:"meeting_info_list,omitempty"`
	// 数据总条数。
	TotalCount *int64 `json:"total_count,omitempty"`
	// 数据总页数。
	TotalPage *int64 `json:"total_page,omitempty"`
}

// V1PmiMeetingsGet200ResponseMeetingInfoListInner struct for V1PmiMeetingsGet200ResponseMeetingInfoListInner
type V1PmiMeetingsGet200ResponseMeetingInfoListInner struct {
	// 会议预订结束时间（UTC 秒）（UTC 秒）
	EndTime *int64 `json:"end_time,omitempty"`
	// 有效会议Code
	MeetingCode *string `json:"meeting_code,omitempty"`
	// 会议 ID
	MeetingId *string `json:"meeting_id,omitempty"`
	// 会议预订开始时间（UTC 秒）（UTC 秒）
	StartTime *int64 `json:"start_time,omitempty"`
	// 会议状态
	Status *int64 `json:"status,omitempty"`
	// 会议主题
	Subject *string `json:"subject,omitempty"`
}
