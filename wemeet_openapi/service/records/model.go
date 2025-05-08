// Package wemeetopenapi is auto generate
/*
   腾讯会议OpenAPI

   SAAS版RESTFUL风格API

   API version: v1.0.7
*/
package wemeetopenapi

// V1AddressesGet200Response struct for V1AddressesGet200Response
type V1AddressesGet200Response struct {
	// 当前页
	CurrentPage *int64 `json:"current_page,omitempty"`
	// 当前size
	CurrentSize *int64 `json:"current_size,omitempty"`
	// 会议 code。
	MeetingCode *string `json:"meeting_code,omitempty"`
	// 会议唯一 ID。
	MeetingId *string `json:"meeting_id,omitempty"`
	// 会议录制 ID。
	MeetingRecordId *string `json:"meeting_record_id,omitempty"`
	// 录制文件列表。
	RecordFiles []V1AddressesGet200ResponseRecordFilesInner `json:"record_files,omitempty"`
	// 会议主题。
	Subject *string `json:"subject,omitempty"`
	// 录制总数
	TotalCount *int64 `json:"total_count,omitempty"`
	// 总页数
	TotalPage *int64 `json:"total_page,omitempty"`
}

// V1AddressesGet200ResponseRecordFilesInner struct for V1AddressesGet200ResponseRecordFilesInner
type V1AddressesGet200ResponseRecordFilesInner struct {
	// 音频下载地址。OAuth 鉴权方式下，账号类型为个人免费版、企微创建企业时，该值返回为空。
	AudioAddress *string `json:"audio_address,omitempty"`
	// 下载音频文件格式，例如：m4a。OAuth 鉴权方式下，账号类型为个人免费版、企微创建企业时，该值返回为空。
	AudioAddressFileType *string `json:"audio_address_file_type,omitempty"`
	// 下载地址，过期时间6小时。OAuth 鉴权方式下，账号类型为个人免费版、企微创建企业时，该值返回为空。
	DownloadAddress *string `json:"download_address,omitempty"`
	// 下载视频文件格式，例如：mp4。
	DownloadAddressFileType *string `json:"download_address_file_type,omitempty"`
	// 会议纪要文件列表。。OAuth 鉴权方式下，账号类型为个人免费版、企微创建企业时，该值返回为空。
	MeetingSummary []V1AddressesGet200ResponseRecordFilesInnerMeetingSummaryInner `json:"meeting_summary,omitempty"`
	// 录制文件 ID。
	RecordFileId *string `json:"record_file_id,omitempty"`
	// 播放地址。
	ViewAddress *string `json:"view_address,omitempty"`
}

// V1AddressesGet200ResponseRecordFilesInnerMeetingSummaryInner struct for V1AddressesGet200ResponseRecordFilesInnerMeetingSummaryInner
type V1AddressesGet200ResponseRecordFilesInnerMeetingSummaryInner struct {
	// 下载文件地址。OAuth 鉴权方式下，账号类型为个人免费版、企微创建企业时，该值返回为空。
	DownloadAddress *string `json:"download_address,omitempty"`
	// 下载文件类型，例如：txt、pdf、docs。OAuth 鉴权方式下，账号类型为个人免费版、企微创建企业时，该值返回为空。
	FileType *string `json:"file_type,omitempty"`
}

// V1AddressesRecordFileIdGet200Response struct for V1AddressesRecordFileIdGet200Response
type V1AddressesRecordFileIdGet200Response struct {
	// 录制转写文件（智能优化版）列表。OAuth 鉴权方式下，账号类型为个人免费版、企微创建企业时，该值返回为空。
	AiMeetingTranscripts []V1AddressesRecordFileIdGet200ResponseAiMeetingTranscriptsInner `json:"ai_meeting_transcripts,omitempty"`
	// 智能纪要列表。OAuth 鉴权方式下，账号类型为个人免费版、企微创建企业时，该值返回为空。
	AiMinutes []V1AddressesRecordFileIdGet200ResponseAiMeetingTranscriptsInner `json:"ai_minutes,omitempty"`
	// 音频下载地址。OAuth 鉴权方式下，账号类型为个人免费版、企微创建企业时，该值返回为空。
	AudioAddress *string `json:"audio_address,omitempty"`
	// 下载音频文件格式，例如：m4a。OAuth 鉴权方式下，账号类型为个人免费版、企微创建企业时，该值返回为空。
	AudioAddressFileType *string `json:"audio_address_file_type,omitempty"`
	// 下载地址，过期时间6小时。OAuth 鉴权方式下，账号类型为个人免费版、企微创建企业时，该值返回为空。
	DownloadAddress *string `json:"download_address,omitempty"`
	// 下载视频文件格式，例如：mp4。
	DownloadAddressFileType *string `json:"download_address_file_type,omitempty"`
	// 会议 code。
	MeetingCode *string `json:"meeting_code,omitempty"`
	// 会议唯一 ID。
	MeetingId *string `json:"meeting_id,omitempty"`
	// 会议转写文件列表。OAuth 鉴权方式下，账号类型为个人免费版、企微创建企业时，该值返回为空。
	MeetingSummary []V1AddressesGet200ResponseRecordFilesInnerMeetingSummaryInner `json:"meeting_summary,omitempty"`
	// 录制文件 ID。
	RecordFileId *string `json:"record_file_id,omitempty"`
	// 播放地址。
	ViewAddress *string `json:"view_address,omitempty"`
}

// V1AddressesRecordFileIdGet200ResponseAiMeetingTranscriptsInner struct for V1AddressesRecordFileIdGet200ResponseAiMeetingTranscriptsInner
type V1AddressesRecordFileIdGet200ResponseAiMeetingTranscriptsInner struct {
	// 录制转写文件下载文件地址，链接有效期为5分钟。OAuth 鉴权方式下，账号类型为个人免费版、企微创建企业时，该值返回为空。
	DownloadAddress *string `json:"download_address,omitempty"`
	// 下载文件类型，例如：txt、pdf、docx。OAuth 鉴权方式下，账号类型为个人免费版、企微创建企业时，该值返回为空。
	FileType *string `json:"file_type,omitempty"`
}

// V1FilesRecordsUploadAllPost200Response struct for V1FilesRecordsUploadAllPost200Response
type V1FilesRecordsUploadAllPost200Response struct {
	// 任务ID
	JobId *string `json:"job_id,omitempty"`
}

// V1FilesRecordsUploadFinishPostRequest struct for V1FilesRecordsUploadFinishPostRequest
type V1FilesRecordsUploadFinishPostRequest struct {
	// 自动生成智能转写和智能纪要：true：自动生成（默认）  false：不生成
	AiRecord *bool `json:"ai_record,omitempty"`
	// 操作者ID
	OperatorId string `json:"operator_id"`
	// 操作者ID类型
	OperatorIdType int64 `json:"operator_id_type"`
	// 上传文件中的发言人数：传具体数值代表几人发言，最多支持12人，其中0代表多人发言
	SpeakNumber int64 `json:"speak_number"`
	// 上传事务ID，由upload-prepare接口返回
	UploadId string `json:"upload_id"`
}

// V1FilesRecordsUploadPreparePost200Response struct for V1FilesRecordsUploadPreparePost200Response
type V1FilesRecordsUploadPreparePost200Response struct {
	// 分块数量
	BlockNum *int64 `json:"block_num,omitempty"`
	// 分块大小策略（以字节为单位）
	BlockSize *int64 `json:"block_size,omitempty"`
	// 上传事务 ID
	UploadId *string `json:"upload_id,omitempty"`
}

// V1FilesRecordsUploadPreparePostRequest struct for V1FilesRecordsUploadPreparePostRequest
type V1FilesRecordsUploadPreparePostRequest struct {
	// 文件名base64编码
	FileName string `json:"file_name"`
	// 文件大小（以字节为单位）
	FileSize int64 `json:"file_size"`
	// 文件类型。voice：音频；video：视频
	FileType string `json:"file_type"`
	// 操作者ID
	OperatorId string `json:"operator_id"`
	// 操作者ID类型
	OperatorIdType int64 `json:"operator_id_type"`
}

// V1MetricsRecordsGet200Response struct for V1MetricsRecordsGet200Response
type V1MetricsRecordsGet200Response struct {
	// 概览数据集合。
	Summaries []V1MetricsRecordsGet200ResponseSummariesInner `json:"summaries,omitempty"`
}

// V1MetricsRecordsGet200ResponseSummariesInner struct for V1MetricsRecordsGet200ResponseSummariesInner
type V1MetricsRecordsGet200ResponseSummariesInner struct {
	// 统计时间，格式：yyyy-MM-dd。
	Date *string `json:"date,omitempty"`
	// 下载次数（当天数据），默认0。
	DownloadCounts *int64 `json:"download_counts,omitempty"`
	// 观看次数（当天数据），默认0。
	ViewCounts *int64 `json:"view_counts,omitempty"`
}

// V1RecordsAccessMeetingRecordIdDeleteRequest struct for V1RecordsAccessMeetingRecordIdDeleteRequest
type V1RecordsAccessMeetingRecordIdDeleteRequest struct {
	// 成员列表，如果传入非有权限的成员，则不生效
	AccessMembers []V1RecordsAccessMeetingRecordIdDeleteRequestAccessMembersInner `json:"access_members"`
	// 操作者 ID。operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。
	OperatorId string `json:"operator_id"`
	// 操作者 ID 的类型： 1:userid 2:open_id
	OperatorIdType int64 `json:"operator_id_type"`
}

// V1RecordsAccessMeetingRecordIdDeleteRequestAccessMembersInner struct for V1RecordsAccessMeetingRecordIdDeleteRequestAccessMembersInner
type V1RecordsAccessMeetingRecordIdDeleteRequestAccessMembersInner struct {
	// 被操作者ID，根据 to_operator_id_type 的值，使用不同的类型
	ToOperatorId string `json:"to_operator_id"`
	// 被操作者ID类型  1:userid  2:open_id  4:ms_open_id
	ToOperatorIdType int64 `json:"to_operator_id_type"`
}

// V1RecordsAccessMeetingRecordIdPost200Response struct for V1RecordsAccessMeetingRecordIdPost200Response
type V1RecordsAccessMeetingRecordIdPost200Response struct {
	// 未添加成功的成员列表
	FailAccessMembers []V1RecordsAccessMeetingRecordIdPost200ResponseFailAccessMembersInner `json:"fail_access_members,omitempty"`
}

// V1RecordsAccessMeetingRecordIdPost200ResponseFailAccessMembersInner struct for V1RecordsAccessMeetingRecordIdPost200ResponseFailAccessMembersInner
type V1RecordsAccessMeetingRecordIdPost200ResponseFailAccessMembersInner struct {
	// 成员访问权限，默认为 0 ； 0：仅查看，1：可管理
	Permission *int64 `json:"permission,omitempty"`
	// 被操作者ID，根据 to_operator_id_type 的值，使用不同的类型
	ToOperatorId *string `json:"to_operator_id,omitempty"`
	// 被操作者ID类型  1:userid  2:open_id  4:ms_open_id
	ToOperatorIdType *int64 `json:"to_operator_id_type,omitempty"`
}

// V1RecordsAccessMeetingRecordIdPostRequest struct for V1RecordsAccessMeetingRecordIdPostRequest
type V1RecordsAccessMeetingRecordIdPostRequest struct {
	// 成员列表，一次最多传入200个，可以多次调用接口，累加型接口
	AccessMembers []V1RecordsAccessMeetingRecordIdPostRequestAccessMembersInner `json:"access_members"`
	// 操作者 ID。operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。
	OperatorId string `json:"operator_id"`
	// 操作者 ID 的类型： 1:userid 2:open_id
	OperatorIdType int64 `json:"operator_id_type"`
}

// V1RecordsAccessMeetingRecordIdPostRequestAccessMembersInner struct for V1RecordsAccessMeetingRecordIdPostRequestAccessMembersInner
type V1RecordsAccessMeetingRecordIdPostRequestAccessMembersInner struct {
	// 成员访问权限，默认为 0 ； 0：仅查看，1：可管理
	Permission *int64 `json:"permission,omitempty"`
	// 被操作者ID，根据 to_operator_id_type 的值，使用不同的类型
	ToOperatorId string `json:"to_operator_id"`
	// 被操作者ID类型,只支持设置参会成员  1:userid  2:open_id  4:ms_open_id
	ToOperatorIdType int64 `json:"to_operator_id_type"`
}

// V1RecordsApprovalsMeetingRecordIdPutRequest struct for V1RecordsApprovalsMeetingRecordIdPutRequest
type V1RecordsApprovalsMeetingRecordIdPutRequest struct {
	// 审批动作。 1：同意 2：忽略
	Action int64 `json:"action"`
	// 申请 ID 列表，通过订阅云录制查看申请事件（可跳转链接），可以获取申请 ID。
	ApplyIdList []string `json:"apply_id_list"`
	// 操作者 ID。 operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。 operator_id_type=2，operator_id 必须和公共参数的 openid 一致。 operator_id 和 userid 至少填写一个，两个参数如果都传了以 operator_id 为准。 使用 OAuth 公参鉴权后不能使用 userid 为入参。
	OperatorId string `json:"operator_id"`
	// 操作者 ID 的类型： 1：userid 2：open_id 如果 operator_id 和 userid 具有值，则以 operator_id 为准。
	OperatorIdType int64 `json:"operator_id_type"`
}

// V1RecordsEventsGet200Response struct for V1RecordsEventsGet200Response
type V1RecordsEventsGet200Response struct {
	// 分页查询返回当前页码。
	CurrentPage *int64 `json:"current_page,omitempty"`
	// 分页查询返回单页数据条数。
	CurrentSize *int64 `json:"current_size,omitempty"`
	// 事件明细集合。
	Events []V1RecordsEventsGet200ResponseEventsInner `json:"events,omitempty"`
	// 分页查询返回数据总数。
	TotalCount *int64 `json:"total_count,omitempty"`
	// 分页查询返回分页总数。
	TotalPage *int64 `json:"total_page,omitempty"`
}

// V1RecordsEventsGet200ResponseEventsInner struct for V1RecordsEventsGet200ResponseEventsInner
type V1RecordsEventsGet200ResponseEventsInner struct {
	// 查询事件类型：1：下载，2：查看。
	EventType *int64 `json:"event_type,omitempty"`
	// 操作时间，UNIX 时间戳（单位毫秒）。
	OperateTime *int64 `json:"operate_time,omitempty"`
	// 录制文件名称。
	RecordName *string `json:"record_name,omitempty"`
	// 用户名称。
	UserName *string `json:"user_name,omitempty"`
	// 用户 ID。
	Userid *string `json:"userid,omitempty"`
}

// V1RecordsGet200Response struct for V1RecordsGet200Response
type V1RecordsGet200Response struct {
	// 分页查询返回当前页码。
	CurrentPage *int64 `json:"current_page,omitempty"`
	// 分页查询返回单页数据条数。
	CurrentSize *int64 `json:"current_size,omitempty"`
	// 会议录制列表。
	RecordMeetings []V1RecordsGet200ResponseRecordMeetingsInner `json:"record_meetings,omitempty"`
	// 分页查询返回数据总数。
	TotalCount *int64 `json:"total_count,omitempty"`
	// 分页查询返回分页总数。
	TotalPage *int64 `json:"total_page,omitempty"`
}

// V1RecordsGet200ResponseRecordMeetingsInner struct for V1RecordsGet200ResponseRecordMeetingsInner
type V1RecordsGet200ResponseRecordMeetingsInner struct {
	// 主持人用户 ID。(查询会议录制文件列表接口返回。)
	HostUserId *string `json:"host_user_id,omitempty"`
	// 会议类型，0-全部 1-外部会议 2-内部会议
	MediaSetType *int64 `json:"media_set_type,omitempty"`
	// 会议开始时间，UNIX 时间戳（单位毫秒）。
	MediaStartTime *int64 `json:"media_start_time,omitempty"`
	// 会议 code。
	MeetingCode *string `json:"meeting_code,omitempty"`
	// 会议唯一 ID。
	MeetingId *string `json:"meeting_id,omitempty"`
	// 会议录制 ID。
	MeetingRecordId *string `json:"meeting_record_id,omitempty"`
	// 录制文件列表。
	RecordFiles []V1RecordsGet200ResponseRecordMeetingsInnerRecordFilesInner `json:"record_files,omitempty"`
	// 返回的录制文件类型，0：云录制、2：上传录制
	RecordType *int64 `json:"record_type,omitempty"`
	// 录制状态。1：录制中，2：转码中，3：转码完成，当状态为转码完成才会返回录制文件列表。
	State *int64 `json:"state,omitempty"`
	// 会议主题。
	Subject *string `json:"subject,omitempty"`
}

// V1RecordsGet200ResponseRecordMeetingsInnerRecordFilesInner struct for V1RecordsGet200ResponseRecordMeetingsInnerRecordFilesInner
type V1RecordsGet200ResponseRecordMeetingsInnerRecordFilesInner struct {
	// 是否允许下载，开启共享时返回。
	AllowDownload *bool `json:"allow_download,omitempty"`
	// 访问密码，开启共享时返回。
	Password *string `json:"password,omitempty"`
	// 结束录制时间，UNIX 时间戳（单位毫秒）。
	RecordEndTime *int64 `json:"record_end_time,omitempty"`
	// 录制文件 ID。
	RecordFileId *string `json:"record_file_id,omitempty"`
	// 文件大小（单位字节）。
	RecordSize *int64 `json:"record_size,omitempty"`
	// 开始录制时间，UNIX 时间戳（单位毫秒）。
	RecordStartTime *int64 `json:"record_start_time,omitempty"`
	// 仅参会成员可查看，开启共享时返回。
	RequiredParticipant *bool `json:"required_participant,omitempty"`
	// 仅企业用户可查看，开启共享时返回。
	RequiredSameCorp *bool `json:"required_same_corp,omitempty"`
	// 共享链接有效期（单位毫秒），当未开启共享时，返回0表示永久有效；开启共享时返回。
	SharingExpire *int64 `json:"sharing_expire,omitempty"`
	// 共享状态，是否开启共享。0：未开启1：开启，当开启共享时返回访问权限、访问密码、共享链接有效期、是否允许下载。
	SharingState *int64 `json:"sharing_state,omitempty"`
	// 共享链接，开启共享时返回。
	SharingUrl *string `json:"sharing_url,omitempty"`
}

// V1RecordsSettingsMeetingRecordIdGet200Response struct for V1RecordsSettingsMeetingRecordIdGet200Response
type V1RecordsSettingsMeetingRecordIdGet200Response struct {
	// 会议ID
	MeetingId *string `json:"meeting_id,omitempty"`
	// 录制名称
	MeetingRecordName *string                                                      `json:"meeting_record_name,omitempty"`
	SharingConfig     *V1RecordsSettingsMeetingRecordIdGet200ResponseSharingConfig `json:"sharing_config,omitempty"`
}

// V1RecordsSettingsMeetingRecordIdGet200ResponseSharingConfig 共享配置对象
type V1RecordsSettingsMeetingRecordIdGet200ResponseSharingConfig struct {
	// 是否允许下载
	AllowDownload *bool `json:"allow_download,omitempty"`
	// 是否允许查看录制转写
	AllowViewTranscripts *bool `json:"allow_view_transcripts,omitempty"`
	// 是否需要审批
	EnableApprove *bool `json:"enable_approve,omitempty"`
	// 是否开启秘密
	EnablePassword *bool `json:"enable_password,omitempty"`
	// 共享链接开关，true-开启，false-未开启
	EnableSharing *bool `json:"enable_sharing,omitempty"`
	// 是否开启共享链接有效期
	EnableSharingExpire *bool `json:"enable_sharing_expire,omitempty"`
	// 录制上传者、会议创建者、企业超级管理员或有企业录制管理权限的用户返回，
	Password *string `json:"password,omitempty"`
	// 访问范围，0-所有人，1-同企业
	ShareScope *int64 `json:"share_scope,omitempty"`
	// 共享链接有效期，毫秒级时间戳
	SharingExpire *int64 `json:"sharing_expire,omitempty"`
}

// V1RecordsSettingsMeetingRecordIdPutRequest struct for V1RecordsSettingsMeetingRecordIdPutRequest
type V1RecordsSettingsMeetingRecordIdPutRequest struct {
	// 会议id
	MeetingId *string `json:"meeting_id,omitempty"`
	// 操作者ID
	OperatorId *string `json:"operator_id,omitempty"`
	// 操作者ID的类型。3. rooms_id 说明：当前仅支持 rooms_id。如操作者为企业内 userid 或 openId，请使用 userid 字段。
	OperatorIdType *int64                                                   `json:"operator_id_type,omitempty"`
	SharingConfig  *V1RecordsSettingsMeetingRecordIdPutRequestSharingConfig `json:"sharing_config,omitempty"`
	// 用户id。仅会议创建者、企业超级管理员或有企业录制管理权限的用户可调用。调用方用于标示用户的唯一 ID（企业内部请使用企业唯一用户标识；OAuth2.0 鉴权用户请使用 openId）。 企业唯一用户标识说明： 1. 企业对接 SSO 时使用的员工唯一标识 ID。 2. 企业调用创建用户接口时传递的 userid 参数。
	Userid *string `json:"userid,omitempty"`
}

// V1RecordsSettingsMeetingRecordIdPutRequestSharingConfig struct for V1RecordsSettingsMeetingRecordIdPutRequestSharingConfig
type V1RecordsSettingsMeetingRecordIdPutRequestSharingConfig struct {
	// 是否允许下载，默认为 false。 true：允许下载
	AllowDownload *bool `json:"allow_download,omitempty"`
	// 是否允许查看会议纪要，默认为 true。
	AllowViewTranscripts *bool `json:"allow_view_transcripts,omitempty"`
	// 是否需要审批，true：需要审批 false：不需要审批
	EnableApprove *bool `json:"enable_approve,omitempty"`
	// 是否开启密码，默认为 true。 true：开启
	EnablePassword *bool `json:"enable_password,omitempty"`
	// 共享开关，是否开启共享，默认为 true。 true：开启 false：未开启 说明： 未开启时不允许设置以下参数。 修改为 false 关闭共享后，之前设置的共享设置将不保存。
	EnableSharing *bool `json:"enable_sharing,omitempty"`
	// 是否开启共享链接有效期，默认为 false。 true：开启
	EnableSharingExpire *bool `json:"enable_sharing_expire,omitempty"`
	// 共享密码，默认随机生成。 说明：当 enable_password = true 时，必传；当 enable_password = false 时，不可传。
	Password *string `json:"password,omitempty"`
	// 访问范围，0：所有人 1：同企业
	ShareScope *int64 `json:"share_scope,omitempty"`
	// 共享权限类型。 0：仅允许登录用户查看 1：仅企业用户成员可查看 2：仅参会成员可查看
	SharingAuthType *int64 `json:"sharing_auth_type,omitempty"`
	// 共享链接有效期，unix 时间戳（单位毫秒），默认为空。 说明：当 enable_sharing_expire = true 时，必传；当 enable_sharing_expire = false 时，不可传。
	SharingExpire *int64 `json:"sharing_expire,omitempty"`
}

// V1RecordsTranscriptsDetailsGet200Response struct for V1RecordsTranscriptsDetailsGet200Response
type V1RecordsTranscriptsDetailsGet200Response struct {
	Minutes *V1RecordsTranscriptsDetailsGet200ResponseMinutes `json:"minutes,omitempty"`
	// 是否还有更多
	More *bool `json:"more,omitempty"`
}

// V1RecordsTranscriptsDetailsGet200ResponseMinutes 会议纪要对象。
type V1RecordsTranscriptsDetailsGet200ResponseMinutes struct {
	// 声纹识别状态0-未完成 1-已完成。说明：声纹识别是针对 Rooms 等设备出现一台设备多人讲话场景时，自动区分为多个发言人的能力。声纹识别与纪要生成的过程独立。无需声纹识别或声纹识别已完成时，该值为1。
	AudioDetect *int64 `json:"audio_detect,omitempty"`
	// 段落对象列表
	Paragraphs []V1RecordsTranscriptsDetailsGet200ResponseMinutesParagraphsInner `json:"paragraphs,omitempty"`
}

// V1RecordsTranscriptsDetailsGet200ResponseMinutesParagraphsInner struct for V1RecordsTranscriptsDetailsGet200ResponseMinutesParagraphsInner
type V1RecordsTranscriptsDetailsGet200ResponseMinutesParagraphsInner struct {
	// 录制文件中的段落结束时间（毫秒）。
	EndTime *int64 `json:"end_time,omitempty"`
	// 段落id
	Pid         *string                                                                         `json:"pid,omitempty"`
	Sentences   []V1RecordsTranscriptsDetailsGet200ResponseMinutesParagraphsInnerSentencesInner `json:"sentences,omitempty"`
	SpeakerInfo *V1RecordsTranscriptsDetailsGet200ResponseMinutesParagraphsInnerSpeakerInfo     `json:"speaker_info,omitempty"`
	// 录制文件中的段落开始时间（毫秒）。
	StartTime *int64 `json:"start_time,omitempty"`
}

// V1RecordsTranscriptsDetailsGet200ResponseMinutesParagraphsInnerSentencesInner struct for V1RecordsTranscriptsDetailsGet200ResponseMinutesParagraphsInnerSentencesInner
type V1RecordsTranscriptsDetailsGet200ResponseMinutesParagraphsInnerSentencesInner struct {
	// 录制文件中的句子结束时间（毫秒）。
	EndTime *int64 `json:"end_time,omitempty"`
	// 句子 ID。
	Sid *string `json:"sid,omitempty"`
	// 录制文件中的句子开始时间（毫秒）。
	StartTime *int64 `json:"start_time,omitempty"`
	// 词对象列表
	Words []V1RecordsTranscriptsDetailsGet200ResponseMinutesParagraphsInnerSentencesInnerWordsInner `json:"words,omitempty"`
}

// V1RecordsTranscriptsDetailsGet200ResponseMinutesParagraphsInnerSentencesInnerWordsInner struct for V1RecordsTranscriptsDetailsGet200ResponseMinutesParagraphsInnerSentencesInnerWordsInner
type V1RecordsTranscriptsDetailsGet200ResponseMinutesParagraphsInnerSentencesInnerWordsInner struct {
	// 录制文件中的词结束时间（毫秒）。
	EndTime *int64 `json:"end_time,omitempty"`
	// 录制文件中的词开始时间（毫秒）
	StartTime *int64 `json:"start_time,omitempty"`
	// 文本内容。
	Text *string `json:"text,omitempty"`
	// 词 ID
	Wid *string `json:"wid,omitempty"`
}

// V1RecordsTranscriptsDetailsGet200ResponseMinutesParagraphsInnerSpeakerInfo 发言人信息对象。
type V1RecordsTranscriptsDetailsGet200ResponseMinutesParagraphsInnerSpeakerInfo struct {
	// 同企业返回企业用户 userid。
	Userid *string `json:"userid,omitempty"`
	// 昵称
	Username *string `json:"username,omitempty"`
}

// V1RecordsTranscriptsParagraphsGet200Response struct for V1RecordsTranscriptsParagraphsGet200Response
type V1RecordsTranscriptsParagraphsGet200Response struct {
	// 声纹识别状态0-未完成 1-已完成。说明：声纹识别是针对 Rooms 等设备出现一台设备多人讲话场景时，自动区分为多个发言人的能力。声纹识别与纪要生成的过程独立。无需声纹识别或声纹识别已完成时，该值为1。
	AudioDetect *int64 `json:"audio_detect,omitempty"`
	// 段落对象列表
	Pids []V1RecordsTranscriptsParagraphsGet200ResponsePidsInner `json:"pids,omitempty"`
	// 段落总数。
	Total *int64 `json:"total,omitempty"`
}

// V1RecordsTranscriptsParagraphsGet200ResponsePidsInner struct for V1RecordsTranscriptsParagraphsGet200ResponsePidsInner
type V1RecordsTranscriptsParagraphsGet200ResponsePidsInner struct {
	// 录制文件中的段落结束时间（毫秒）。
	EndTime *int64 `json:"end_time,omitempty"`
	// 段落 ID。
	Pid *string `json:"pid,omitempty"`
	// 录制文件中的段落开始时间（毫秒）。
	StartTime *int64 `json:"start_time,omitempty"`
}

// V1RecordsTranscriptsSearchGet200Response struct for V1RecordsTranscriptsSearchGet200Response
type V1RecordsTranscriptsSearchGet200Response struct {
	// 搜索结果列表
	Hits []V1RecordsTranscriptsSearchGet200ResponseHitsInner `json:"hits,omitempty"`
	// 搜索结果时间戳对象列表
	Timelines []V1RecordsTranscriptsSearchGet200ResponseTimelinesInner `json:"timelines,omitempty"`
}

// V1RecordsTranscriptsSearchGet200ResponseHitsInner struct for V1RecordsTranscriptsSearchGet200ResponseHitsInner
type V1RecordsTranscriptsSearchGet200ResponseHitsInner struct {
	// 匹配长度
	Length *int64 `json:"length,omitempty"`
	// text 相对词的偏移。
	Offset *int64 `json:"offset,omitempty"`
	// 段落 ID
	Pid *string `json:"pid,omitempty"`
	// 句子 ID
	Sid *string `json:"sid,omitempty"`
}

// V1RecordsTranscriptsSearchGet200ResponseTimelinesInner struct for V1RecordsTranscriptsSearchGet200ResponseTimelinesInner
type V1RecordsTranscriptsSearchGet200ResponseTimelinesInner struct {
	// 段落id
	Pid *string `json:"pid,omitempty"`
	// 句子id
	Sid *string `json:"sid,omitempty"`
	// 录制文件中的词开始时间（毫秒）
	StartTime *int64 `json:"start_time,omitempty"`
}

// V1RecordsTransferRecordingPutRequest struct for V1RecordsTransferRecordingPutRequest
type V1RecordsTransferRecordingPutRequest struct {
	// 如果参数未带， 则按集群删除策略对指定录制删除操作。    转存完成后删除录制策略：  0 - 转存完成后立刻删除录制文件  1～30 - 转存完成后1～30天后删除录制文件  不删除 - 转存完成后不删除录制文件
	DeleteRecordingAfterTransfer *string `json:"delete_recording_after_transfer,omitempty"`
	Instanceid                   int64   `json:"instanceid"`
	// 操作者 ID，根据 operator_id_type 的值，使用不同的类型。
	OperatorId string `json:"operator_id"`
	// 操作者 ID 的类型： 1：userid
	OperatorIdType int64 `json:"operator_id_type"`
}

// V1RecordsViewDetailsGet200Response struct for V1RecordsViewDetailsGet200Response
type V1RecordsViewDetailsGet200Response struct {
	CurrentPage *int64 `json:"current_page,omitempty"`
	CurrentSize *int64 `json:"current_size,omitempty"`
	TotalCount  *int64 `json:"total_count,omitempty"`
	TotalPage   *int64 `json:"total_page,omitempty"`
	// 事件明细集合。
	ViewDetails []V1RecordsViewDetailsGet200ResponseViewDetailsInner `json:"view_details,omitempty"`
}

// V1RecordsViewDetailsGet200ResponseViewDetailsInner struct for V1RecordsViewDetailsGet200ResponseViewDetailsInner
type V1RecordsViewDetailsGet200ResponseViewDetailsInner struct {
	// 录制文件名称。
	RecordFileName *string `json:"record_file_name,omitempty"`
	// 观看完成度（百分比），该用户累计观看该视频的完成度。
	TotalViewProgress *string `json:"total_view_progress,omitempty"`
	// 用户名称。匿名用户给出匿名用户标识。
	UserName *string `json:"user_name,omitempty"`
	// 所在同一企业下的用户 ID。
	Userid *string `json:"userid,omitempty"`
	// 观看结束时间，UNIX时间戳（单位毫秒）。
	ViewEndTime *int64 `json:"view_end_time,omitempty"`
	// Integer 观看开始时间，UNIX 时间戳（单位毫秒）。
	ViewStartTime *int64 `json:"view_start_time,omitempty"`
	// 实际观看时长（单位毫秒）。
	ViewTime *int64 `json:"view_time,omitempty"`
}
