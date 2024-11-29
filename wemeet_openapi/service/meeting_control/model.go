// Package wemeetopenapi is auto generate
/*
   腾讯会议OpenAPI

   SAAS版RESTFUL风格API

   API version: v1.0.4
*/
package wemeetopenapi

// V1MeetingsMeetingIdDismissPostRequest struct for V1MeetingsMeetingIdDismissPostRequest
type V1MeetingsMeetingIdDismissPostRequest struct {
	// 是否强制结束会议，默认值为1：0：不强制结束会议，会议中有参会者，则无法强制结束会议 1 ：强制结束会议，会议中有参会者，也会强制结束会议
	ForceDismissMeeting *int64 `json:"force_dismiss_meeting,omitempty"`
	// 设备类型
	Instanceid int64 `json:"instanceid"`
	// 操作者ID，根据operator_id_type的值，使用不同的类型
	OperatorId *string `json:"operator_id,omitempty"`
	// 操作者ID的类型：1:userid  2:openid（预留编号，本次不添加，未来新增接口使用）3:rooms_id  4: ms_open_id
	OperatorIdType *int64 `json:"operator_id_type,omitempty"`
	// 原因代码，可为用户自定义
	ReasonCode int64 `json:"reason_code"`
	// 取消原因
	ReasonDetail *string `json:"reason_detail,omitempty"`
	// 是否回收会议号，默认值为0： 0：不回收会议号，可以重新入会 1： 回收会议号，不可重新入会
	RetrieveCode *int64 `json:"retrieve_code,omitempty"`
	// 调用方用于标示用户的唯一 ID（企业内部请使用企业唯一用户标识；OAuth2.0 鉴权用户请使用 openId）。 企业唯一用户标识说明：企业对接 SSO 时使用的员工唯一标识 ID，企业调用创建用户接口时传递的 userid 参数。
	Userid *string `json:"userid,omitempty"`
}

// V1RealControlMeetingsMeetingIdAsrPut200Response struct for V1RealControlMeetingsMeetingIdAsrPut200Response
type V1RealControlMeetingsMeetingIdAsrPut200Response struct {
	Code    *int64  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// V1RealControlMeetingsMeetingIdAsrPutRequest struct for V1RealControlMeetingsMeetingIdAsrPutRequest
type V1RealControlMeetingsMeetingIdAsrPutRequest struct {
	// 用户的终端设备类型：  0：PSTN  1：PC  2：Mac  3：Android  4：iOS  5：Web  6：iPad  7：Android Pad  8：小程序  9：voip、sip 设备  10：Linux  20：Rooms for Touch Windows  21：Rooms for Touch MacOS  22：Rooms for Touch Android  30：Controller for Touch Windows  32：Controller for Touch Android  33：Controller for Touch iOS
	InstanceId int64 `json:"instance_id"`
	// 开启/关闭实时转写 true：开启实时转写 false：关闭实时转写
	IsOpen bool `json:"is_open"`
	// 是否自动打开转写侧边栏，仅在is_open 为 true 时生效，默认为 0， 0：打开实时转写页面 。1：不打开实时转写页面
	OpenAsrView *int64 `json:"open_asr_view,omitempty"`
	// 操作者 ID。operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。
	OperatorId *string `json:"operator_id,omitempty"`
	// 操作者 ID 的类型： 1：userid
	OperatorIdType *int64 `json:"operator_id_type,omitempty"`
}

// V1RealControlMeetingsMeetingIdCohostsPutRequest struct for V1RealControlMeetingsMeetingIdCohostsPutRequest
type V1RealControlMeetingsMeetingIdCohostsPutRequest struct {
	// 具体设置动作： true：设置联席主持人， false：撤销联席主持人
	Action bool `json:"action"`
	// 用户的终端设备类型： 0：PSTN 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch MacOS 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch iOS 说明：使用 ms_open_id 进行调用时，仅支持以上1-8的设备类型。
	Instanceid int64 `json:"instanceid"`
	// 操作者 ID。 1：operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。 2：接口输入参数如果需要传用户 ID 时，operator_id 和 uuid 不可以同时为空，两个参数如果都传了以 operator_id 为准。 3：如果 operator_id_type=2，operator_id 必须和公共参数的 openid 一致。
	OperatorId *string `json:"operator_id,omitempty"`
	// 操作者 ID 的类型： 2：openid 4：ms_open_id
	OperatorIdType *int64                                              `json:"operator_id_type,omitempty"`
	User           V1RealControlMeetingsMeetingIdCohostsPutRequestUser `json:"user"`
	// 操作者用户唯一身份 ID，仅支持主持人，且只适用于单场会议。即将废弃，推荐使用ms_open_id。
	Uuid *string `json:"uuid,omitempty"`
}

// V1RealControlMeetingsMeetingIdCohostsPutRequestUser struct for V1RealControlMeetingsMeetingIdCohostsPutRequestUser
type V1RealControlMeetingsMeetingIdCohostsPutRequestUser struct {
	// 用户的终端设备类型： 0：PSTN 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch MacOS 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch iOS 说明：请与被操作者的设备类型保持一致，否则不生效。使用 ms_open_id 进行调用时，仅支持以上1-8的设备类型。
	Instanceid int64 `json:"instanceid"`
	// 用户ID，根据to_operator_id_type的值，使用不同的类型
	ToOperatorId *string `json:"to_operator_id,omitempty"`
	// 用户ID的类型：  4: ms_open_id
	ToOperatorIdType *int64 `json:"to_operator_id_type,omitempty"`
	// 用户的唯一标识uuid
	Uuid *string `json:"uuid,omitempty"`
}

// V1RealControlMeetingsMeetingIdDocPutRequest struct for V1RealControlMeetingsMeetingIdDocPutRequest
type V1RealControlMeetingsMeetingIdDocPutRequest struct {
	// 是否允许全员上传文档  true：是 false：否
	EnableUploadDoc bool `json:"enable_upload_doc"`
	// 用户的终端设备类型： 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch Mac 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch Iphone
	Instanceid int64 `json:"instanceid"`
	// 操作者 ID。 1：operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。 2：接口输入参数如果需要传用户 ID 时，operator_id 和 uuid 不可以同时为空，两个参数如果都传了以 operator_id 为准。 3：如果 operator_id_type=2，operator_id 必须和公共参数的 openid 一致。
	OperatorId *string `json:"operator_id,omitempty"`
	// 操作者ID的类型：  2:openid  4: ms_open_id
	OperatorIdType *int64 `json:"operator_id_type,omitempty"`
	// 操作者用户唯一身份 ID，仅支持主持人和联席主持人，且只适用于单场会议。即将废弃，推荐使用ms_open_id。
	Uuid *string `json:"uuid,omitempty"`
}

// V1RealControlMeetingsMeetingIdKickoutPutRequest struct for V1RealControlMeetingsMeetingIdKickoutPutRequest
type V1RealControlMeetingsMeetingIdKickoutPutRequest struct {
	// 移出后是否允许再次入会： true：允许再次入会 false：不允许
	AllowRejoin bool `json:"allow_rejoin"`
	// 用户的终端设备类型： 0：PSTN 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch MacOS 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch iOS
	Instanceid int64 `json:"instanceid"`
	// 操作者 ID。 1：operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。 2：接口输入参数如果需要传用户 ID 时，operator_id 和 uuid 不可以同时为空，两个参数如果都传了以 operator_id 为准。 3：如果 operator_id_type=2，operator_id 必须和公共参数的 openid 一致。
	OperatorId *string `json:"operator_id,omitempty"`
	// 操作者ID的类型：2:openid 4: ms_open_id
	OperatorIdType *int64 `json:"operator_id_type,omitempty"`
	// 移出原因说明。当用户设备为 MRA 时，该参数必须填写移出原因。
	Reason *string `json:"reason,omitempty"`
	// 被操作用户对象信息列表
	Users []V1RealControlMeetingsMeetingIdKickoutPutRequestUsersInner `json:"users"`
	// 操作者用户唯一身份 ID，仅支持主持人和联席主持人，且只适用于单场会议。即将废弃，推荐使用ms_open_id。
	Uuid *string `json:"uuid,omitempty"`
}

// V1RealControlMeetingsMeetingIdKickoutPutRequestUsersInner struct for V1RealControlMeetingsMeetingIdKickoutPutRequestUsersInner
type V1RealControlMeetingsMeetingIdKickoutPutRequestUsersInner struct {
	// 用户的终端设备类型： 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch Mac 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch Iphone
	Instanceid int64 `json:"instanceid"`
	// 用户ID，根据to_operator_id_type的值，使用不同的类型
	ToOperatorId *string `json:"to_operator_id,omitempty"`
	// 用户ID的类型： 4: ms_open_id
	ToOperatorIdType *int64 `json:"to_operator_id_type,omitempty"`
	// 用户的唯一标识uuid
	Uuid *string `json:"uuid,omitempty"`
}

// V1RealControlMeetingsMeetingIdMutesPutRequest struct for V1RealControlMeetingsMeetingIdMutesPutRequest
type V1RealControlMeetingsMeetingIdMutesPutRequest struct {
	// 用户的终端设备类型： 0：PSTN 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch MacOS 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch iOS
	Instanceid int64 `json:"instanceid"`
	// 是否静音： true：静音 false：解除静音
	Mute bool `json:"mute"`
	// 操作者 ID。 1：operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。 2：接口输入参数如果需要传用户 ID 时，operator_id 和 uuid 不可以同时为空，两个参数如果都传了以 operator_id 为准。 3：如果 operator_id_type=2，operator_id 必须和公共参数的 openid 一致。
	OperatorId *string `json:"operator_id,omitempty"`
	// 操作者ID的类型：2:openid 4: ms_open_id
	OperatorIdType *int64                                            `json:"operator_id_type,omitempty"`
	User           V1RealControlMeetingsMeetingIdMutesPutRequestUser `json:"user"`
	// 操作者用户唯一身份 ID，仅支持主持人和联席主持人，且只适用于单场会议。即将废弃，推荐使用ms_open_id。
	Uuid *string `json:"uuid,omitempty"`
}

// V1RealControlMeetingsMeetingIdMutesPutRequestUser struct for V1RealControlMeetingsMeetingIdMutesPutRequestUser
type V1RealControlMeetingsMeetingIdMutesPutRequestUser struct {
	// 用户的终端设备类型： 0：PSTN 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch MacOS 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch iOS 说明：请与被操作者的设备类型保持一致，否则不生效。
	Instanceid int64 `json:"instanceid"`
	// 用户ID，根据to_operator_id_type的值，使用不同的类型
	ToOperatorId *string `json:"to_operator_id,omitempty"`
	// 用户ID的类型：  4: ms_open_id
	ToOperatorIdType *int64 `json:"to_operator_id_type,omitempty"`
	// 用户的唯一标识uuid
	Uuid *string `json:"uuid,omitempty"`
}

// V1RealControlMeetingsMeetingIdNamesPutRequest struct for V1RealControlMeetingsMeetingIdNamesPutRequest
type V1RealControlMeetingsMeetingIdNamesPutRequest struct {
	// 用户的终端设备类型： 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch Mac 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch Iphone
	Instanceid int64 `json:"instanceid"`
	// 操作者 ID。 1：operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。 2：如果 operator_id_type=2，operator_id 必须和公共参数的 openid 一致。
	OperatorId string `json:"operator_id"`
	// 操作者ID的类型：2:openid 4: ms_open_id
	OperatorIdType int64 `json:"operator_id_type"`
	// 要改名的用户
	Users []V1RealControlMeetingsMeetingIdNamesPutRequestUsersInner `json:"users"`
}

// V1RealControlMeetingsMeetingIdNamesPutRequestUsersInner struct for V1RealControlMeetingsMeetingIdNamesPutRequestUsersInner
type V1RealControlMeetingsMeetingIdNamesPutRequestUsersInner struct {
	// 用户的终端设备类型： 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch Mac 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch Iphone
	Instanceid int64 `json:"instanceid"`
	// 被操作者ms_open_id
	MsOpenId string `json:"ms_open_id"`
	// 要修改的昵称名，限制20个字符。
	NickName *string `json:"nick_name,omitempty"`
}

// V1RealControlMeetingsMeetingIdScreenSharedPutRequest struct for V1RealControlMeetingsMeetingIdScreenSharedPutRequest
type V1RealControlMeetingsMeetingIdScreenSharedPutRequest struct {
	// 用户的终端设备类型： 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch Mac 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch Iphone
	Instanceid int64 `json:"instanceid"`
	// 操作者 ID。 1：operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。 2：接口输入参数如果需要传用户 ID 时，operator_id 和 uuid 不可以同时为空，两个参数如果都传了以 operator_id 为准。 3：如果 operator_id_type=2，operator_id 必须和公共参数的 openid 一致。
	OperatorId *string `json:"operator_id,omitempty"`
	// 操作者ID的类型：  2:openid  4: ms_open_id
	OperatorIdType *int64                                                   `json:"operator_id_type,omitempty"`
	User           V1RealControlMeetingsMeetingIdScreenSharedPutRequestUser `json:"user"`
	// 操作者用户唯一身份 ID，仅支持主持人和联席主持人，且只适用于单场会议。即将废弃，推荐使用ms_open_id。
	Uuid *string `json:"uuid,omitempty"`
}

// V1RealControlMeetingsMeetingIdScreenSharedPutRequestUser struct for V1RealControlMeetingsMeetingIdScreenSharedPutRequestUser
type V1RealControlMeetingsMeetingIdScreenSharedPutRequestUser struct {
	// 用户的终端设备类型： 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch Mac 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch Iphone
	Instanceid int64 `json:"instanceid"`
	// 用户ID，根据to_operator_id_type的值，使用不同的类型
	ToOperatorId *string `json:"to_operator_id,omitempty"`
	// 用户ID的类型：4: ms_open_id
	ToOperatorIdType *int64 `json:"to_operator_id_type,omitempty"`
	// 用户的唯一标识uuid
	Uuid *string `json:"uuid,omitempty"`
}

// V1RealControlMeetingsMeetingIdStatusPutRequest struct for V1RealControlMeetingsMeetingIdStatusPutRequest
type V1RealControlMeetingsMeetingIdStatusPutRequest struct {
	// 允许参会者聊天设置  0:允许参会者自由聊天  1:仅允许参会者公开聊天  2:只允许支持人发言
	AllowChat *int64 `json:"allow_chat,omitempty"`
	// 是否允许成员自己解除静音
	AllowUnmuteBySelf *bool `json:"allow_unmute_by_self,omitempty"`
	// 是否开启等候室 true：开启 false：关闭
	AutoWaitingRoom *bool `json:"auto_waiting_room,omitempty"`
	// 是否允许参会者发送红包 true：允许 false：不允许
	EnableRedEnvelope *bool `json:"enable_red_envelope,omitempty"`
	// 隐藏会议号和密码 true：隐藏 false：不隐藏
	HideMeetingCodePassword *bool `json:"hide_meeting_code_password,omitempty"`
	// 用户的终端设备类型： 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch Mac 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch Iphone
	Instanceid int64 `json:"instanceid"`
	// 是否锁定会议 true：锁定 false：关闭锁定
	MeetingLocked *bool `json:"meeting_locked,omitempty"`
	// 是否全体静音，true：是；false关闭全体静音
	MuteAll *bool `json:"mute_all,omitempty"`
	// 是否仅企业成员可入会  true：仅企业成员可入会  false：不限制
	OnlyEnterpriseUserAllowed *bool `json:"only_enterprise_user_allowed,omitempty"`
	// 操作者ID，根据operator_id_type的值，使用不同的类型
	OperatorId *string `json:"operator_id,omitempty"`
	// 操作者ID的类型：1:userid  2:openid（预留编号，本次不添加，未来新增接口使用）3:rooms_id  4: ms_open_id
	OperatorIdType *int64 `json:"operator_id_type,omitempty"`
	// 成员入会静音 0:关闭静音 1:开启静音 2:超过6人自动开启静音
	ParticipantJoinMute *int64 `json:"participant_join_mute,omitempty"`
	// 成员入会是否播放提示音 true：成员入会播放提示音 false：不播放
	PlayIvrOnJoin *bool `json:"play_ivr_on_join,omitempty"`
	// 是否允许参会者发起屏幕共享 true：允许 false：不允许
	ShareScreen *bool `json:"share_screen,omitempty"`
	// 用户的唯一标识uuid
	Uuid *string `json:"uuid,omitempty"`
}

// V1RealControlMeetingsMeetingIdVideoPutRequest struct for V1RealControlMeetingsMeetingIdVideoPutRequest
type V1RealControlMeetingsMeetingIdVideoPutRequest struct {
	// 用户的终端设备类型： 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch Mac 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch Iphone
	Instanceid int64 `json:"instanceid"`
	// 操作者 ID。 1：operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。 2：接口输入参数如果需要传用户 ID 时，operator_id 和 uuid 不可以同时为空，两个参数如果都传了以 operator_id 为准。 3：如果 operator_id_type=2，operator_id 必须和公共参数的 openid 一致。
	OperatorId *string `json:"operator_id,omitempty"`
	// 操作者ID的类型： 2:openid 4: ms_open_id
	OperatorIdType *int64                                            `json:"operator_id_type,omitempty"`
	User           V1RealControlMeetingsMeetingIdVideoPutRequestUser `json:"user"`
	// 操作者用户唯一身份 ID，仅支持主持人和联席主持人，且只适用于单场会议。即将废弃，推荐使用ms_open_id。
	Uuid *string `json:"uuid,omitempty"`
	// 是否开启视频： false：关闭视频（默认值）。 true：开启视频， 仅支持 MRA 设备。
	Video *bool `json:"video,omitempty"`
}

// V1RealControlMeetingsMeetingIdVideoPutRequestUser struct for V1RealControlMeetingsMeetingIdVideoPutRequestUser
type V1RealControlMeetingsMeetingIdVideoPutRequestUser struct {
	// 用户的终端设备类型： 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch Mac 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch Iphone
	Instanceid int64 `json:"instanceid"`
	// 被操作者 ID，根据 operator_id_type 的值，使用不同的类型。和 uuid 不可同时为空。
	ToOperatorId *string `json:"to_operator_id,omitempty"`
	// 用户ID的类型： 4: ms_open_id
	ToOperatorIdType *int64 `json:"to_operator_id_type,omitempty"`
	// 用户的唯一标识uuid
	Uuid *string `json:"uuid,omitempty"`
}

// V1RealControlMeetingsMeetingIdWaitingRoomPutRequest struct for V1RealControlMeetingsMeetingIdWaitingRoomPutRequest
type V1RealControlMeetingsMeetingIdWaitingRoomPutRequest struct {
	// 移出后是否允许再次加入会议  true：允许 false：不允许  说明：操作类型参数 operete_type = 3 时才允许设置
	AllowRejoin *bool `json:"allow_rejoin,omitempty"`
	// 用户的终端设备类型： 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch Mac 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch Iphone
	Instanceid int64 `json:"instanceid"`
	// 操作类型： 1：主持人将等候室成员移入会议  2：主持人将会中成员移入等候室  3：主持人将等候室成员移出等候室
	OperateType int64 `json:"operate_type"`
	// 操作者 ID。 1：operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。 2：接口输入参数如果需要传用户 ID 时，operator_id 和 uuid 不可以同时为空，两个参数如果都传了以 operator_id 为准。 3：如果 operator_id_type=2，operator_id 必须和公共参数的 openid 一致。
	OperatorId *string `json:"operator_id,omitempty"`
	// 操作者ID的类型： 2:openid 4: ms_open_id
	OperatorIdType *int64 `json:"operator_id_type,omitempty"`
	// 被操作用户对象信息列表
	Users []V1RealControlMeetingsMeetingIdMutesPutRequestUser `json:"users"`
	// 操作者用户唯一身份 ID，仅支持主持人和联席主持人，且只适用于单场会议。即将废弃，推荐使用ms_open_id。
	Uuid *string `json:"uuid,omitempty"`
}
