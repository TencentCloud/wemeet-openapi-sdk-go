// Package wemeetopenapi is auto generate
/*
   腾讯会议OpenAPI

   SAAS版RESTFUL风格API

   API version: v1.0.5
*/
package wemeetopenapi

// V1MeetingsMeetingIdAdvancedLayoutsPost200Response struct for V1MeetingsMeetingIdAdvancedLayoutsPost200Response
type V1MeetingsMeetingIdAdvancedLayoutsPost200Response struct {
	// 布局对象列表
	LayoutList []V1MeetingsMeetingIdAdvancedLayoutsPost200ResponseLayoutListInner `json:"layout_list,omitempty"`
}

// V1MeetingsMeetingIdAdvancedLayoutsPost200ResponseLayoutListInner struct for V1MeetingsMeetingIdAdvancedLayoutsPost200ResponseLayoutListInner
type V1MeetingsMeetingIdAdvancedLayoutsPost200ResponseLayoutListInner struct {
	// 布局 ID
	LayoutId *string `json:"layout_id,omitempty"`
	// 布局名称
	LayoutName *string `json:"layout_name,omitempty"`
	// 布局单页对象列表
	PageList []V1MeetingsMeetingIdAdvancedLayoutsPost200ResponseLayoutListInnerPageListInner `json:"page_list,omitempty"`
}

// V1MeetingsMeetingIdAdvancedLayoutsPost200ResponseLayoutListInnerPageListInner struct for V1MeetingsMeetingIdAdvancedLayoutsPost200ResponseLayoutListInnerPageListInner
type V1MeetingsMeetingIdAdvancedLayoutsPost200ResponseLayoutListInnerPageListInner struct {
	// 开启/关闭轮询
	EnablePolling *bool `json:"enable_polling,omitempty"`
	// 布局模板 ID
	LayoutTemplateId *string                                                                                      `json:"layout_template_id,omitempty"`
	PollingSetting   *V1MeetingsMeetingIdAdvancedLayoutsPost200ResponseLayoutListInnerPageListInnerPollingSetting `json:"polling_setting,omitempty"`
	// 用户座次对象列表
	UserSeatList []V1MeetingsMeetingIdAdvancedLayoutsPost200ResponseLayoutListInnerPageListInnerUserSeatListInner `json:"user_seat_list,omitempty"`
}

// V1MeetingsMeetingIdAdvancedLayoutsPost200ResponseLayoutListInnerPageListInnerPollingSetting 轮询参数设置对象
type V1MeetingsMeetingIdAdvancedLayoutsPost200ResponseLayoutListInnerPageListInnerPollingSetting struct {
	// 轮询开启后设置参数， 设置是否忽略未入会成员
	IgnoreUserAbsence *bool `json:"ignore_user_absence,omitempty"`
	// 轮询开启后设置参数，设置是否忽略没开启视频成员
	IgnoreUserNovideo *bool `json:"ignore_user_novideo,omitempty"`
	// 轮询开启后设置参数 轮询间隔时长， 允许取值范围1～999999
	PollingInterval *int64 `json:"polling_interval,omitempty"`
	// 轮询开启后设置参数。 轮询间隔时间类型： 1-秒 2-分钟
	PollingIntervalUnit *int64 `json:"polling_interval_unit,omitempty"`
}

// V1MeetingsMeetingIdAdvancedLayoutsPost200ResponseLayoutListInnerPageListInnerUserSeatListInner struct for V1MeetingsMeetingIdAdvancedLayoutsPost200ResponseLayoutListInnerPageListInnerUserSeatListInner
type V1MeetingsMeetingIdAdvancedLayoutsPost200ResponseLayoutListInnerPageListInnerUserSeatListInner struct {
	// 宫格 ID
	GridId *string `json:"grid_id,omitempty"`
	// 宫格类型： 1-视频画面 2-共享画面
	GridType *int64 `json:"grid_type,omitempty"`
	// 宫格中的用户列表 ● 轮询关闭， 只有一个用户 ● 轮询开启后， 可以包含多个用户
	UserList []V1MeetingsMeetingIdAdvancedLayoutsPost200ResponseLayoutListInnerPageListInnerUserSeatListInnerUserListInner `json:"user_list,omitempty"`
	// 视频画面来源 1-演讲者 2-自动填充 3-指定人员，根据user_list的定义显示视频内容（此类型需传递 userid 或 ms_open_id、username 入参，作为视频画面展示；若会中参会成员有外部企业用户，需传递该用户的 ms_open_id；如果 userid、ms_open_id 同时传递则以 ms_open_id 为准）
	VideoType *int64 `json:"video_type,omitempty"`
}

// V1MeetingsMeetingIdAdvancedLayoutsPost200ResponseLayoutListInnerPageListInnerUserSeatListInnerUserListInner struct for V1MeetingsMeetingIdAdvancedLayoutsPost200ResponseLayoutListInnerPageListInnerUserSeatListInnerUserListInner
type V1MeetingsMeetingIdAdvancedLayoutsPost200ResponseLayoutListInnerPageListInnerUserSeatListInnerUserListInner struct {
	// 用户当前会议临时身份 ID，单场会议唯一
	MsOpenId *string `json:"ms_open_id,omitempty"`
	// 用户 ID
	Userid *string `json:"userid,omitempty"`
	// 用户昵称，base64编码
	Username *string `json:"username,omitempty"`
}

// V1MeetingsMeetingIdAdvancedLayoutsPostRequest struct for V1MeetingsMeetingIdAdvancedLayoutsPostRequest
type V1MeetingsMeetingIdAdvancedLayoutsPostRequest struct {
	// 设备类型 ID 0：PSTN 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch MacOS 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch iOS
	Instanceid int64 `json:"instanceid"`
	// 布局对象列表
	LayoutList []V1MeetingsMeetingIdAdvancedLayoutsPostRequestLayoutListInner `json:"layout_list"`
	// 操作人ID
	OperatorId string `json:"operator_id"`
	// 操作人id类型，1：userid，4 :ms_open_id
	OperatorIdType int64 `json:"operator_id_type"`
}

// V1MeetingsMeetingIdAdvancedLayoutsPostRequestLayoutListInner struct for V1MeetingsMeetingIdAdvancedLayoutsPostRequestLayoutListInner
type V1MeetingsMeetingIdAdvancedLayoutsPostRequestLayoutListInner struct {
	// 布局名称
	LayoutName *string `json:"layout_name,omitempty"`
	// 布局单页对象列表
	PageList []V1MeetingsMeetingIdAdvancedLayoutsPostRequestLayoutListInnerPageListInner `json:"page_list"`
}

// V1MeetingsMeetingIdAdvancedLayoutsPostRequestLayoutListInnerPageListInner struct for V1MeetingsMeetingIdAdvancedLayoutsPostRequestLayoutListInnerPageListInner
type V1MeetingsMeetingIdAdvancedLayoutsPostRequestLayoutListInnerPageListInner struct {
	// 开启/关闭轮询
	EnablePolling *bool `json:"enable_polling,omitempty"`
	// 布局模板 ID
	LayoutTemplateId string                                                                                   `json:"layout_template_id"`
	PollingSetting   *V1MeetingsMeetingIdAdvancedLayoutsPostRequestLayoutListInnerPageListInnerPollingSetting `json:"polling_setting,omitempty"`
	// 用户座次对象列表
	UserSeatList []V1MeetingsMeetingIdAdvancedLayoutsPostRequestLayoutListInnerPageListInnerUserSeatListInner `json:"user_seat_list"`
}

// V1MeetingsMeetingIdAdvancedLayoutsPostRequestLayoutListInnerPageListInnerPollingSetting 轮询参数设置对象
type V1MeetingsMeetingIdAdvancedLayoutsPostRequestLayoutListInnerPageListInnerPollingSetting struct {
	// 轮询开启后设置参数， 设置是否忽略未入会成员
	IgnoreUserAbsence bool `json:"ignore_user_absence"`
	// 轮询开启后设置参数，设置是否忽略没开启视频成员
	IgnoreUserNovideo bool `json:"ignore_user_novideo"`
	// 轮询开启后设置参数 轮询间隔时长， 允许取值范围1～999999
	PollingInterval int64 `json:"polling_interval"`
	// 轮询开启后设置参数。 轮询间隔时间类型： 1-秒 2-分钟
	PollingIntervalUnit int64 `json:"polling_interval_unit"`
}

// V1MeetingsMeetingIdAdvancedLayoutsPostRequestLayoutListInnerPageListInnerUserSeatListInner struct for V1MeetingsMeetingIdAdvancedLayoutsPostRequestLayoutListInnerPageListInnerUserSeatListInner
type V1MeetingsMeetingIdAdvancedLayoutsPostRequestLayoutListInnerPageListInnerUserSeatListInner struct {
	// 宫格 ID
	GridId string `json:"grid_id"`
	// 宫格类型： 1-视频画面 2-共享画面
	GridType int64 `json:"grid_type"`
	// 宫格中的用户列表 ● 轮询关闭， 只有一个用户 ● 轮询开启后， 可以包含多个用户
	UserList []V1MeetingsMeetingIdAdvancedLayoutsPostRequestLayoutListInnerPageListInnerUserSeatListInnerUserListInner `json:"user_list,omitempty"`
	// 视频画面来源 1-演讲者 2-自动填充 3-指定人员，根据user_list的定义显示视频内容（此类型需传递 userid 或 ms_open_id、username 入参，作为视频画面展示；若会中参会成员有外部企业用户，需传递该用户的 ms_open_id；如果 userid、ms_open_id 同时传递则以 ms_open_id 为准）
	VideoType *int64 `json:"video_type,omitempty"`
}

// V1MeetingsMeetingIdAdvancedLayoutsPostRequestLayoutListInnerPageListInnerUserSeatListInnerUserListInner struct for V1MeetingsMeetingIdAdvancedLayoutsPostRequestLayoutListInnerPageListInnerUserSeatListInnerUserListInner
type V1MeetingsMeetingIdAdvancedLayoutsPostRequestLayoutListInnerPageListInnerUserSeatListInnerUserListInner struct {
	// 用户当前会议临时身份 ID，单场会议唯一
	MsOpenId *string `json:"ms_open_id,omitempty"`
	// 用户 ID
	Userid *string `json:"userid,omitempty"`
	// 用户昵称，base64编码
	Username *string `json:"username,omitempty"`
}

// V1MeetingsMeetingIdApplyingLayoutPutRequest struct for V1MeetingsMeetingIdApplyingLayoutPutRequest
type V1MeetingsMeetingIdApplyingLayoutPutRequest struct {
	// 设备类型 ID 0：PSTN 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch MacOS 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch iOS
	Instanceid int64 `json:"instanceid"`
	// 选择应用的布局 ID （若送空\"\"，表示恢复成当前会议的默认布局）  备注：应用布局的优先级从高到低为： ● 个性布局 ● 自定义布局 ● 默认布局（MRA不支持同框模式， 如果会议设置为同框模式， MRA应用默认布局））
	LayoutId *string `json:"layout_id,omitempty"`
	// 操作人ID
	OperatorId string `json:"operator_id"`
	// 操作人id类型，1 :userid，4 :ms_open_id
	OperatorIdType int64 `json:"operator_id_type"`
	// 用户列表对象数组。如果该字段为空， 为会议设置高级自定义布局；如果该字段携带用户， 则只为指定用户设置个性布局。 单次最多支持20个用户。
	UserList []V1MeetingsMeetingIdApplyingLayoutPutRequestUserListInner `json:"user_list,omitempty"`
}

// V1MeetingsMeetingIdApplyingLayoutPutRequestUserListInner struct for V1MeetingsMeetingIdApplyingLayoutPutRequestUserListInner
type V1MeetingsMeetingIdApplyingLayoutPutRequestUserListInner struct {
	// 用户的终端设备类型：  9：voip、sip 设备（H.323/SIP设备） 说明：请与被操作者的设备类型保持一致，否则不生效。
	Instanceid int64 `json:"instanceid"`
	// 被操作者用户当前会议中的临时身份 ID，单场会议唯一。
	MsOpenId string `json:"ms_open_id"`
}

// V1MeetingsMeetingIdLayoutsPost200Response struct for V1MeetingsMeetingIdLayoutsPost200Response
type V1MeetingsMeetingIdLayoutsPost200Response struct {
	// 布局对象列表
	LayoutList []V1MeetingsMeetingIdLayoutsPost200ResponseLayoutListInner `json:"layout_list,omitempty"`
	// 布局数量
	LayoutNumber *int64 `json:"layout_number,omitempty"`
	// 会议应用的布局ID
	SelectedLayoutId *string `json:"selected_layout_id,omitempty"`
}

// V1MeetingsMeetingIdLayoutsPost200ResponseLayoutListInner struct for V1MeetingsMeetingIdLayoutsPost200ResponseLayoutListInner
type V1MeetingsMeetingIdLayoutsPost200ResponseLayoutListInner struct {
	LayoutId *string `json:"layout_id,omitempty"`
	// 布局单页对象列表
	PageList []V1MeetingsMeetingIdLayoutsPost200ResponseLayoutListInnerPageListInner `json:"page_list,omitempty"`
}

// V1MeetingsMeetingIdLayoutsPost200ResponseLayoutListInnerPageListInner struct for V1MeetingsMeetingIdLayoutsPost200ResponseLayoutListInnerPageListInner
type V1MeetingsMeetingIdLayoutsPost200ResponseLayoutListInnerPageListInner struct {
	// 布局模板ID
	LayoutTemplateId *string `json:"layout_template_id,omitempty"`
	// 用户座次对象列表
	UserSeatList []V1MeetingsMeetingIdLayoutsPost200ResponseLayoutListInnerPageListInnerUserSeatListInner `json:"user_seat_list,omitempty"`
}

// V1MeetingsMeetingIdLayoutsPost200ResponseLayoutListInnerPageListInnerUserSeatListInner struct for V1MeetingsMeetingIdLayoutsPost200ResponseLayoutListInnerPageListInnerUserSeatListInner
type V1MeetingsMeetingIdLayoutsPost200ResponseLayoutListInnerPageListInnerUserSeatListInner struct {
	// 宫格ID
	GridId *string `json:"grid_id,omitempty"`
	// 宫格类型
	GridType *int64 `json:"grid_type,omitempty"`
	// 当场会议的用户临时 ID。
	MsOpenId *string `json:"ms_open_id,omitempty"`
	// 工具箱id
	ToolSdkid *string `json:"tool_sdkid,omitempty"`
	// 用户ID
	Userid *string `json:"userid,omitempty"`
	// 用户昵称
	Username *string `json:"username,omitempty"`
	// 用户身份 ID（腾讯会议颁发的用于开放平台的唯一用户 ID）
	Uuid *string `json:"uuid,omitempty"`
}

// V1MeetingsMeetingIdLayoutsPostRequest struct for V1MeetingsMeetingIdLayoutsPostRequest
type V1MeetingsMeetingIdLayoutsPostRequest struct {
	// 布局列表中会议需要应用的布局序号，从1开始计数（首次添加时若该参数不送，则默认选中第一个布局作为会议应用的布局）
	DefaultLayoutOrder *int64 `json:"default_layout_order,omitempty"`
	// 用户的终端设备类型：1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch Mac 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch Iphone
	Instanceid int64 `json:"instanceid"`
	// 布局对象列表
	LayoutList []V1MeetingsMeetingIdLayoutsPostRequestLayoutListInner `json:"layout_list"`
	// 操作者ID
	OperatorId string `json:"operator_id"`
	// 操作者id的类型，1:userid
	OperatorIdType int64 `json:"operator_id_type"`
	// 会议创建者ID
	Userid string `json:"userid"`
}

// V1MeetingsMeetingIdLayoutsPostRequestLayoutListInner struct for V1MeetingsMeetingIdLayoutsPostRequestLayoutListInner
type V1MeetingsMeetingIdLayoutsPostRequestLayoutListInner struct {
	LayoutId string `json:"layout_id"`
	// 布局单页对象列表
	PageList []V1MeetingsMeetingIdLayoutsPostRequestLayoutListInnerPageListInner `json:"page_list"`
}

// V1MeetingsMeetingIdLayoutsPostRequestLayoutListInnerPageListInner struct for V1MeetingsMeetingIdLayoutsPostRequestLayoutListInnerPageListInner
type V1MeetingsMeetingIdLayoutsPostRequestLayoutListInnerPageListInner struct {
	// 布局模板ID
	LayoutTemplateId string `json:"layout_template_id"`
	// 用户座次对象列表
	UserSeatList []V1MeetingsMeetingIdLayoutsPostRequestLayoutListInnerPageListInnerUserSeatListInner `json:"user_seat_list,omitempty"`
}

// V1MeetingsMeetingIdLayoutsPostRequestLayoutListInnerPageListInnerUserSeatListInner struct for V1MeetingsMeetingIdLayoutsPostRequestLayoutListInnerPageListInnerUserSeatListInner
type V1MeetingsMeetingIdLayoutsPostRequestLayoutListInnerPageListInnerUserSeatListInner struct {
	// 宫格ID
	GridId string `json:"grid_id"`
	// 宫格类型，注意：多次传入同一宫格ID的对象，仅第一次出现的对象生效。 宫格类型： 1. 视频画面（此类型需传递 userid 或 uuid、username 入参，作为视频画面展示；若会中参会成员有外部企业用户，需传递该用户的 uuid；如果 userid、uuid 同时传递则以 uuid 为准）。 2. 共享画面。 3. 拓展应用（目前一页仅可添加一个应用）。 添加的应用需满足以下条件： 与会议绑定。 开启网页服务。 同企业下的仅企业内可见应用或外部企业可见应用。
	GridType int64 `json:"grid_type"`
	// 当场会议的用户临时 ID。
	MsOpenId *string `json:"ms_open_id,omitempty"`
	// 工具箱id
	ToolSdkid *string `json:"tool_sdkid,omitempty"`
	// 用户ID
	Userid *string `json:"userid,omitempty"`
	// 用户昵称
	Username *string `json:"username,omitempty"`
	// 用户身份 ID（腾讯会议颁发的用于开放平台的唯一用户 ID）
	Uuid *string `json:"uuid,omitempty"`
}
