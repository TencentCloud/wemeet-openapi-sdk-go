// Package wemeetopenapi is auto generate
/*
   腾讯会议OpenAPI

   SAAS版RESTFUL风格API

   API version: v1.0.4
*/
package wemeetopenapi

// V1MeetingsMeetingIdMsOpenIdGet200Response struct for V1MeetingsMeetingIdMsOpenIdGet200Response
type V1MeetingsMeetingIdMsOpenIdGet200Response struct {
	// 会议唯一id
	MeetingId *string `json:"meeting_id,omitempty"`
	// 当场会议的用户临时 ID，可用于会控操作，适用于所有用户。
	MsOpenId *string `json:"ms_open_id,omitempty"`
}

// V1PmiMeetingsPmiConfigGet200Response struct for V1PmiMeetingsPmiConfigGet200Response
type V1PmiMeetingsPmiConfigGet200Response struct {
	// 是否允许成员在主持人前入会
	AllowInBeforeHost *bool `json:"allow_in_before_host,omitempty"`
	// 是否开启等候室
	AllowInWaitingRoom *bool `json:"allow_in_waiting_room,omitempty"`
	// 是否允许多端入会
	AllowMultiDevice *bool `json:"allow_multi_device,omitempty"`
	// 是否禁止笔记截屏
	DisableNoteCapture *bool `json:"disable_note_capture,omitempty"`
	// 指定主持人列表
	Hosts []V1PmiMeetingsPmiConfigGet200ResponseHostsInner `json:"hosts,omitempty"`
	// 成员入会静音设置，0-关闭，1-开启，2-超过6人后自动开启
	MuteEnableTypeJoin *int64 `json:"mute_enable_type_join,omitempty"`
	// 是否仅企业内部成员可入会
	OnlyEnterpriseUserAllowed *bool `json:"only_enterprise_user_allowed,omitempty"`
	// 个人会议号
	PmiCode *string `json:"pmi_code,omitempty"`
	// 个人会议室名称
	PmiName *string `json:"pmi_name,omitempty"`
	// 个人会议号密码，经过base64处理
	PmiPassword *string `json:"pmi_password,omitempty"`
	// 水印样式，0-单排，1-多排
	WaterMarkType *int64 `json:"water_mark_type,omitempty"`
	// 是否开启会议水印
	Watermark *bool `json:"watermark,omitempty"`
}

// V1PmiMeetingsPmiConfigGet200ResponseHostsInner userid
type V1PmiMeetingsPmiConfigGet200ResponseHostsInner struct {
	Userid *string `json:"userid,omitempty"`
}

// V1PmiMeetingsPmiConfigPutRequest struct for V1PmiMeetingsPmiConfigPutRequest
type V1PmiMeetingsPmiConfigPutRequest struct {
	// 是否允许成员在主持人进会前加入会议
	AllowInBeforeHost *bool `json:"allow_in_before_host,omitempty"`
	// 是否允许成员多端入会
	AllowMultiDevice *bool `json:"allow_multi_device,omitempty"`
	// 是否开启等候室
	AutoInWaitingRoom *bool `json:"auto_in_waiting_room,omitempty"`
	// 禁止笔记截屏，true-禁止，false-不禁止。当水印参数开启时生效
	DisableNoteCapture *bool `json:"disable_note_capture,omitempty"`
	// 是否需要密码
	EnablePassword *bool `json:"enable_password,omitempty"`
	// 指定主持人列表
	Hosts []V1PmiMeetingsPmiConfigPutRequestHostsInner `json:"hosts,omitempty"`
	// 设备id
	Instanceid int64 `json:"instanceid"`
	// 成员入会静音选项，0-关闭，1-开启，2-超过6人开启
	MuteEnableTypeJoin *int64 `json:"mute_enable_type_join,omitempty"`
	// 是否仅企业内部成员可入会
	OnlyEnterpriseUserAllowed *bool `json:"only_enterprise_user_allowed,omitempty"`
	// 根据type类型传相应内容
	OperatorId string `json:"operator_id"`
	// 操作者ID类型，1 - userid
	OperatorIdType int64 `json:"operator_id_type"`
	// 个人会议室名称，最大支持18个汉字或36个英文字母。
	PmiName *string `json:"pmi_name,omitempty"`
	// 入会密码
	PmiPassword *string `json:"pmi_password,omitempty"`
	// 水印样式。当水印参数为开启时，此参数才生效。 0：单排 1：多排
	WaterMarkType *int64 `json:"water_mark_type,omitempty"`
	// 是否开启会议水印
	Watermark *bool `json:"watermark,omitempty"`
}

// V1PmiMeetingsPmiConfigPutRequestHostsInner struct for V1PmiMeetingsPmiConfigPutRequestHostsInner
type V1PmiMeetingsPmiConfigPutRequestHostsInner struct {
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

// V1UsersAccountAiAccountDeleteRequest struct for V1UsersAccountAiAccountDeleteRequest
type V1UsersAccountAiAccountDeleteRequest struct {
	// 用户拥有企管用户管理权限
	OperatorId string `json:"operator_id"`
	// 1:userid
	OperatorIdType int64  `json:"operator_id_type"`
	ToOperatorId   string `json:"to_operator_id"`
	// 1:userid
	ToOperatorIdType int64 `json:"to_operator_id_type"`
}

// V1UsersAccountAiAccountPostRequest struct for V1UsersAccountAiAccountPostRequest
type V1UsersAccountAiAccountPostRequest struct {
	// 1:购买版 2:赠送版AI账号类型  1：购买版  2：赠送版  如果未传入该字段，默认分配赠送版AI账号
	AiAccountType *int64 `json:"ai_account_type,omitempty"`
	// 操作者ID，拥有用户管理权限
	OperatorId string `json:"operator_id"`
	// ID类型，1:userid
	OperatorIdType int64 `json:"operator_id_type"`
	// 被操作者ID，仅支持企业版/教育版高级账号被设置，其他类型账号会报错
	ToOperatorId string `json:"to_operator_id"`
	// ID类型  1:userid
	ToOperatorIdType int64 `json:"to_operator_id_type"`
}

// V1UsersAccountStatisticsGet200Response struct for V1UsersAccountStatisticsGet200Response
type V1UsersAccountStatisticsGet200Response struct {
	// ai账号类型使用对象（商业版不返回）
	AiAccountDetails []V1UsersAccountStatisticsGet200ResponseAiAccountDetailsInner `json:"ai_account_details,omitempty"`
	// 账号类型使用对象
	UserAccountDetails []V1UsersAccountStatisticsGet200ResponseUserAccountDetailsInner `json:"user_account_details,omitempty"`
	// 当前用户数
	UserCount *int64 `json:"user_count,omitempty"`
}

// V1UsersAccountStatisticsGet200ResponseAiAccountDetailsInner struct for V1UsersAccountStatisticsGet200ResponseAiAccountDetailsInner
type V1UsersAccountStatisticsGet200ResponseAiAccountDetailsInner struct {
	// 账号数
	AiAccountCount *int64 `json:"ai_account_count,omitempty"`
	// ai账号类型，1:购买版 2:赠送版
	AiAccountType *int64 `json:"ai_account_type,omitempty"`
	// 已分配的账号数
	AiAccountUsedCount *int64 `json:"ai_account_used_count,omitempty"`
}

// V1UsersAccountStatisticsGet200ResponseUserAccountDetailsInner struct for V1UsersAccountStatisticsGet200ResponseUserAccountDetailsInner
type V1UsersAccountStatisticsGet200ResponseUserAccountDetailsInner struct {
	// 账号数
	UserAccountCount *int64 `json:"user_account_count,omitempty"`
	// 账号类型，1：高级账号 （企业版，教育版）  2：免费账号  （企业版，教育版，商业版）  3：免费账号100方 （商业版）  4：高级账号300方（商业版）  5：高级账号500方（商业版）  6：高级账号1000方（商业版）  7：高级账号2000方（商业版）
	UserAccountType *int64 `json:"user_account_type,omitempty"`
	// 已分配账号数
	UserAccountUsedCount *int64 `json:"user_account_used_count,omitempty"`
}

// V1UsersAdvanceListGet200Response struct for V1UsersAdvanceListGet200Response
type V1UsersAdvanceListGet200Response struct {
	// 是否还有未拉取的数据
	HasRemaining *bool `json:"has_remaining,omitempty"`
	// 下一次查询pos位置
	NextPos *string                                      `json:"next_pos,omitempty"`
	Users   []V1UsersAdvanceListGet200ResponseUsersInner `json:"users,omitempty"`
}

// V1UsersAdvanceListGet200ResponseUsersInner struct for V1UsersAdvanceListGet200ResponseUsersInner
type V1UsersAdvanceListGet200ResponseUsersInner struct {
	// 账号版本。 0：其他 1：商业版 2：企业版 3：教育版
	AccountVersion *int64 `json:"account_version,omitempty"`
	// AI 账号类型。 0：无账号 1：购买版 2：赠送版
	AiAccountType *int64 `json:"ai_account_type,omitempty"`
	// 手机区号
	Area *string `json:"area,omitempty"`
	// 头像地址
	AvatarUrl *string `json:"avatar_url,omitempty"`
	// 用户部门信息
	DepartmentList []V1UsersAdvanceListGet200ResponseUsersInnerDepartmentListInner `json:"department_list,omitempty"`
	// 邮箱
	Email *string `json:"email,omitempty"`
	// 是否有 AI 账号能力。 true：有  false：无  教育版/企业版存在有 AI 账号，商业版都具有 AI 能力，其余为 false。
	EnableAiAccount *bool `json:"enable_ai_account,omitempty"`
	// 入职时间
	EntryTime *string `json:"entry_time,omitempty"`
	// 员工职位
	JobTitle *string `json:"job_title,omitempty"`
	// 手机号
	Phone *string `json:"phone,omitempty"`
	// 手机号验证状态。 0：未知 1：已验证 2：未验证
	PhoneStatus *int64 `json:"phone_status,omitempty"`
	// 角色类型
	RoleCode *string `json:"role_code,omitempty"`
	// 角色名称
	RoleName *string `json:"role_name,omitempty"`
	// 员工工号
	StaffId *string `json:"staff_id,omitempty"`
	// 账号状态。账号状态： 1：正常 2：注销 3：未激活 4：禁用 5：预注册
	Status *string `json:"status,omitempty"`
	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`
	// 账号类型。 1：高级账号（企业版/教育版） 2：免费账号（企业版/教育版） 3：免费账号100方 （商业版） 4：高级账号300方（商业版） 5：高级账号500方（商业版） 6：高级账号1000方（商业版） 7：高级账号2000方（商业版） 8：高级账号100方（商业版）
	UserAccountType *int64 `json:"user_account_type,omitempty"`
	// 用户userid
	Userid *string `json:"userid,omitempty"`
	// 用户名称
	Username *string `json:"username,omitempty"`
	// 用户uuid
	Uuid *string `json:"uuid,omitempty"`
}

// V1UsersAdvanceListGet200ResponseUsersInnerDepartmentListInner struct for V1UsersAdvanceListGet200ResponseUsersInnerDepartmentListInner
type V1UsersAdvanceListGet200ResponseUsersInnerDepartmentListInner struct {
	// 部门全路径
	DepartmentFullName *string `json:"department_full_name,omitempty"`
	// 部门ID
	DepartmentId *string `json:"department_id,omitempty"`
	// 部门名称
	DepartmentName *string `json:"department_name,omitempty"`
	// 是否主部门
	IsMain *bool `json:"is_main,omitempty"`
}

// V1UsersDeleteTransferPostRequest struct for V1UsersDeleteTransferPostRequest
type V1UsersDeleteTransferPostRequest struct {
	// 删除用户的数据处理方式： 1=彻底删除； 2=转移给指定成员；
	DataProcess *string `json:"data_process,omitempty"`
	// 操作者 ID。 operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。 operator_id_type=2，operator_id 必须和公共参数的 openid 一致。
	OperatorId string `json:"operator_id"`
	// 操作者 ID 的类型： 1：userid 2：open_id
	OperatorIdType int64 `json:"operator_id_type"`
	// 数据接收者的ID，根据receiver_id_type的值，使用不同的类型。； data_process为2时生效； 该userid不存在时，将报错；
	ReceiverId *string `json:"receiver_id,omitempty"`
	// 数据接收者 ID 的类型：  1：userid  2：open_id
	ReceiverIdType *int64 `json:"receiver_id_type,omitempty"`
	// 被操作者 ID，根据 to_operator_id_type 的值，使用不同的类型，这里指被删除的用户。
	ToOperatorId string `json:"to_operator_id"`
	// 被操作者 ID 的类型： 1：userid 2：open_id
	ToOperatorIdType int64 `json:"to_operator_id_type"`
	// 转移的具体数据； 0=全部； 1=云录制； 2=会议列表； data_process为2时生效； 不传时默认为0；
	TransferData *string `json:"transfer_data,omitempty"`
}

// V1UsersGet200Response struct for V1UsersGet200Response
type V1UsersGet200Response struct {
	// ai账号类型 1：购买版 2:赠送版
	AiAccountType *int64 `json:"ai_account_type,omitempty"`
	// 地区编码（国内默认86）。
	Area *string `json:"area,omitempty"`
	// 用户部门信息。
	DepartmentList []V1UsersGet200ResponseDepartmentListInner `json:"department_list,omitempty"`
	// 邮箱地址。
	Email *string `json:"email,omitempty"`
	// 是否有ai账号能力
	EnableAiAccount *bool `json:"enable_ai_account,omitempty"`
	// 入职时间。
	EntryTime *string `json:"entry_time,omitempty"`
	// 员工职位。
	JobTitle *string `json:"job_title,omitempty"`
	// 企业员工手机号码。
	Phone *string `json:"phone,omitempty"`
	// 手机号验证状态。 0：未知 1：已验证 2：未验证
	PhoneStatus *int64 `json:"phone_status,omitempty"`
	// 角色类型。
	RoleCode *string `json:"role_code,omitempty"`
	// 角色名称。
	RoleName *string `json:"role_name,omitempty"`
	// 员工工号。
	StaffId *string `json:"staff_id,omitempty"`
	// 用户状态： 1：正常 2：注销 3：未激活 4：禁用
	Status *string `json:"status,omitempty"`
	// 更新时间，格式：yyyy-MM-dd HH:mm:ss。
	UpdateTime *string `json:"update_time,omitempty"`
	// 调用方用于标示用户的唯一 ID（例如企业用户可以为企业账户英文名、个人用户可以为手机号等）。
	Userid *string `json:"userid,omitempty"`
	// 用户昵称。
	Username *string `json:"username,omitempty"`
	// String 用户身份 ID（腾讯会议颁发的用于开放平台的唯一用户 ID）。
	Uuid *string `json:"uuid,omitempty"`
}

// V1UsersGet200ResponseDepartmentListInner struct for V1UsersGet200ResponseDepartmentListInner
type V1UsersGet200ResponseDepartmentListInner struct {
	// 部门 ID。
	DepartmentId *string `json:"department_id,omitempty"`
	// 部门名称。
	DepartmentName *string `json:"department_name,omitempty"`
}

// V1UsersInfoBasicGet200Response struct for V1UsersInfoBasicGet200Response
type V1UsersInfoBasicGet200Response struct {
	AccountType *int64 `json:"account_type,omitempty"`
	// 商企版计费需求，账号版本
	AccountVersion *int64 `json:"account_version,omitempty"`
	// AI账号类型 1:购买版 2:赠送版
	AiAccountType *int64  `json:"ai_account_type,omitempty"`
	AvatarUrl     *string `json:"avatar_url,omitempty"`
	// 是否有AI账号能力，true：有， false：无，教育版/企业版存在ai账号，商业版都是ai账号，其余为false
	EnableAiAccount *bool   `json:"enable_ai_account,omitempty"`
	OpenCorpId      *string `json:"open_corp_id,omitempty"`
	OpenCorpName    *string `json:"open_corp_name,omitempty"`
	// 手机号验证状态。 0：未知 1：已验证 2：未验证
	PhoneStatus *int64  `json:"phone_status,omitempty"`
	Status      *string `json:"status,omitempty"`
	// 账号类型 1：高级账号  2：免费账号  3：免费账号100方 4:高级账号300方，5:高级账号500方，6：高级账号1000方，7:高级账号2000方
	UserAccountType *int64  `json:"user_account_type,omitempty"`
	Username        *string `json:"username,omitempty"`
}

// V1UsersInviteActivatePost200Response struct for V1UsersInviteActivatePost200Response
type V1UsersInviteActivatePost200Response struct {
	// 未激活用户对象列表
	InactivateUserList []V1UsersInviteActivatePost200ResponseInactivateUserListInner `json:"inactivate_user_list,omitempty"`
}

// V1UsersInviteActivatePost200ResponseInactivateUserListInner struct for V1UsersInviteActivatePost200ResponseInactivateUserListInner
type V1UsersInviteActivatePost200ResponseInactivateUserListInner struct {
	// 激活链接
	ActivateUrl *string `json:"activate_url,omitempty"`
	Userid      *string `json:"userid,omitempty"`
}

// V1UsersInviteActivatePostRequest struct for V1UsersInviteActivatePostRequest
type V1UsersInviteActivatePostRequest struct {
	// 操作者ID
	OperatorId string `json:"operator_id"`
	// 操作者ID类型，1:userid
	OperatorIdType int64 `json:"operator_id_type"`
	// 未激活的账号列表，最多支持传100个
	UseridList []string `json:"userid_list"`
}

// V1UsersInviteAuthPost200Response struct for V1UsersInviteAuthPost200Response
type V1UsersInviteAuthPost200Response struct {
	// 未验证用户对象列表
	AuthUserList  []V1UsersInviteAuthPost200ResponseAuthUserListInner  `json:"auth_user_list,omitempty"`
	ErrorUserList []V1UsersInviteAuthPost200ResponseErrorUserListInner `json:"error_user_list,omitempty"`
}

// V1UsersInviteAuthPost200ResponseAuthUserListInner struct for V1UsersInviteAuthPost200ResponseAuthUserListInner
type V1UsersInviteAuthPost200ResponseAuthUserListInner struct {
	// 验证链接
	AuthUrl *string `json:"auth_url,omitempty"`
	// 账号 ID
	Userid *string `json:"userid,omitempty"`
}

// V1UsersInviteAuthPost200ResponseErrorUserListInner struct for V1UsersInviteAuthPost200ResponseErrorUserListInner
type V1UsersInviteAuthPost200ResponseErrorUserListInner struct {
	// 错误码
	ErrorCode *int64 `json:"error_code,omitempty"`
	// 错误描述
	ErrorMsg *string `json:"error_msg,omitempty"`
	// 账号ID
	Userid *string `json:"userid,omitempty"`
}

// V1UsersInviteAuthPostRequest struct for V1UsersInviteAuthPostRequest
type V1UsersInviteAuthPostRequest struct {
	// 操作者 ID。operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。
	OperatorId string `json:"operator_id"`
	// 操作者 ID 的类型： 1：userid
	OperatorIdType int64 `json:"operator_id_type"`
	// 未验证 userid 列表，最多一次支持传100个
	UseridList []string `json:"userid_list"`
}

// V1UsersListGet200Response struct for V1UsersListGet200Response
type V1UsersListGet200Response struct {
	// 当前页数。
	CurrentPage *int64 `json:"current_page,omitempty"`
	// 当前页实际大小。
	CurrentSize *int64 `json:"current_size,omitempty"`
	// 分页大小。
	PageSize *int64 `json:"page_size,omitempty"`
	// 总数。
	TotalCount *int64 `json:"total_count,omitempty"`
	// 数组格式，item 为用户对象。
	Users []V1UsersListGet200ResponseUsersInner `json:"users,omitempty"`
}

// V1UsersListGet200ResponseUsersInner struct for V1UsersListGet200ResponseUsersInner
type V1UsersListGet200ResponseUsersInner struct {
	// 账号版本
	AccountVersion *int64 `json:"account_version,omitempty"`
	// ai账号类型 1:购买版 2:赠送版
	AiAccountType *int64 `json:"ai_account_type,omitempty"`
	// 手机区号。
	Area *string `json:"area,omitempty"`
	// 用户图像地址。
	AvatarUrl *string `json:"avatar_url,omitempty"`
	// 用户部门信息。
	DepartmentList []V1UsersListGet200ResponseUsersInnerDepartmentListInner `json:"department_list,omitempty"`
	// 邮箱。
	Email *string `json:"email,omitempty"`
	// 是否有ai账号能力  true：有  false：无  教育版/企业版存在有ai账号，商业版都具有ai能力，其余为false
	EnableAiAccount *bool `json:"enable_ai_account,omitempty"`
	// 入职时间。
	EntryTime *string `json:"entry_time,omitempty"`
	// 员工职位。
	JobTitle *string `json:"job_title,omitempty"`
	// 手机号。
	Phone *string `json:"phone,omitempty"`
	// 手机号验证状态。 0：未知 1：已验证 2：未验证
	PhoneStatus *int64 `json:"phone_status,omitempty"`
	// 角色类型。
	RoleCode *string `json:"role_code,omitempty"`
	// 角色名称。
	RoleName *string `json:"role_name,omitempty"`
	// String  员工工号。
	StaffId *string `json:"staff_id,omitempty"`
	// 账号状态： 1：正常 2：注销 3：未激活 4：禁用
	Status *string `json:"status,omitempty"`
	// String  更新时间。
	UpdateTime *string `json:"update_time,omitempty"`
	// 账号类型    1：高级账号 （企业版，教育版）  2：免费账号  （企业版，教育版，商业版）  3：免费账号100方 （商业版）  4：高级账号300方（商业版）  5：高级账号500方（商业版）  6：高级账号1000方（商业版）  7:高级账号2000方（商业版）
	UserAccountType *int64 `json:"user_account_type,omitempty"`
	// String  用户 userid。
	Userid *string `json:"userid,omitempty"`
	// 用户 name。
	Username *string `json:"username,omitempty"`
	// 用户身份 ID（腾讯会议颁发的用于开放平台的唯一用户 ID）。
	Uuid *string `json:"uuid,omitempty"`
}

// V1UsersListGet200ResponseUsersInnerDepartmentListInner DepartmentInfo对象数组
type V1UsersListGet200ResponseUsersInnerDepartmentListInner struct {
	// 部门 ID。
	DepartmentId *string `json:"department_id,omitempty"`
	// String  部门名称。
	DepartmentName *string `json:"department_name,omitempty"`
}

// V1UsersOpenIdToUseridPost200Response struct for V1UsersOpenIdToUseridPost200Response
type V1UsersOpenIdToUseridPost200Response struct {
	// 转换成功的该自建应用所在企业下的userid、open_id对应关系列表。
	UseridList []V1UsersOpenIdToUseridPost200ResponseUseridListInner `json:"userid_list,omitempty"`
}

// V1UsersOpenIdToUseridPost200ResponseUseridListInner struct for V1UsersOpenIdToUseridPost200ResponseUseridListInner
type V1UsersOpenIdToUseridPost200ResponseUseridListInner struct {
	// 需要转换的open_id
	OpenId *string `json:"open_id,omitempty"`
	// 转换成功后，该open_id所对应的本企业下用户的userid。
	Userid *string `json:"userid,omitempty"`
}

// V1UsersOpenIdToUseridPostRequest struct for V1UsersOpenIdToUseridPostRequest
type V1UsersOpenIdToUseridPostRequest struct {
	// 操作者ID
	OperatorId string `json:"operator_id"`
	// 操作者ID类型
	OperatorIdType int64 `json:"operator_id_type"`
	// 第三方应用的sdkid。需要转换的open_id应为腾讯会议为该三方应用提供的open_id。
	Sdkid string `json:"sdkid"`
}

// V1UsersPost200Response struct for V1UsersPost200Response
type V1UsersPost200Response struct {
	Email    *string `json:"email,omitempty"`
	Phone    *string `json:"phone,omitempty"`
	Userid   *string `json:"userid,omitempty"`
	Username *string `json:"username,omitempty"`
	Uuid     *string `json:"uuid,omitempty"`
}

// V1UsersPostRequest struct for V1UsersPostRequest
type V1UsersPostRequest struct {
	Area *string `json:"area,omitempty"`
	// 自动发送邀请，开启之后调用接口后自动发送激活邀请 true：开启，默认开启;false：关闭
	AutoInvite *bool   `json:"auto_invite,omitempty"`
	Email      *string `json:"email,omitempty"`
	EntryTime  *int64  `json:"entry_time,omitempty"`
	JobTitle   *string `json:"job_title,omitempty"`
	// 操作者ID
	OperatorId string `json:"operator_id"`
	// 操作者ID类型，1:userid
	OperatorIdType int64   `json:"operator_id_type"`
	Phone          string  `json:"phone"`
	StaffId        *string `json:"staff_id,omitempty"`
	// 1：高级账号  2：免费账号  3：免费账号100方 4:高级账号300方，5:高级账号500方，6：高级账号1000方，7:高级账号2000方     其中企业版/教育版：1，2 。免费组织 2。 商业版：2-7      根据传入的参数判断是否有该类型账号，没有则报错。创建成功即锁定该账号资源。默认值：商业版默认为高级账号，绑定资源为由小到大，资源消耗完账号为免费账号，企业版-高级账号
	UserAccountType *int64 `json:"user_account_type,omitempty"`
	Userid          string `json:"userid"`
	Username        string `json:"username"`
}

// V1UsersPutRequest struct for V1UsersPutRequest
type V1UsersPutRequest struct {
	AvatarUrl *string `json:"avatar_url,omitempty"`
	// 员工部门，暂只支持为用户分配1个部门。
	DepartmentList []string `json:"department_list,omitempty"`
	Email          *string  `json:"email,omitempty"`
	EntryTime      *int64   `json:"entry_time,omitempty"`
	JobTitle       *string  `json:"job_title,omitempty"`
	// 操作者ID
	OperatorId string `json:"operator_id"`
	// 操作者ID类型，1:userid
	OperatorIdType int64   `json:"operator_id_type"`
	Phone          *string `json:"phone,omitempty"`
	StaffId        *string `json:"staff_id,omitempty"`
	Username       *string `json:"username,omitempty"`
}

// V1UsersUseridEnablePutRequest struct for V1UsersUseridEnablePutRequest
type V1UsersUseridEnablePutRequest struct {
	// 是否启用用户： true：启用 false：禁用
	Enable bool `json:"enable"`
	// 操作者ID
	OperatorId string `json:"operator_id"`
	// 操作者ID类型，1:userid
	OperatorIdType int64 `json:"operator_id_type"`
}

// V1UsersUseridGet200Response struct for V1UsersUseridGet200Response
type V1UsersUseridGet200Response struct {
	AccountType *int64 `json:"account_type,omitempty"`
	// 账号版本
	AccountVersion *int64 `json:"account_version,omitempty"`
	// ai账号类型 1:购买版 2:赠送版
	AiAccountType  *int64                                           `json:"ai_account_type,omitempty"`
	Area           *string                                          `json:"area,omitempty"`
	AvatarUrl      *string                                          `json:"avatar_url,omitempty"`
	DepartmentList []V1UsersUseridGet200ResponseDepartmentListInner `json:"department_list,omitempty"`
	Email          *string                                          `json:"email,omitempty"`
	// 是否有ai账号能力，true：有，false：无
	EnableAiAccount *bool   `json:"enable_ai_account,omitempty"`
	EntryTime       *string `json:"entry_time,omitempty"`
	JobTitle        *string `json:"job_title,omitempty"`
	Phone           *string `json:"phone,omitempty"`
	// 手机号验证状态。 0：未知 1：已验证 2：未验证
	PhoneStatus *int64  `json:"phone_status,omitempty"`
	RoleCode    *string `json:"role_code,omitempty"`
	RoleName    *string `json:"role_name,omitempty"`
	StaffId     *string `json:"staff_id,omitempty"`
	Status      *string `json:"status,omitempty"`
	UpdateTime  *string `json:"update_time,omitempty"`
	//  1：高级账号  2：免费账号  3：免费账号100方 4:高级账号300方，5:高级账号500方，6：高级账号1000方，7:高级账号2000方
	UserAccountType *int64  `json:"user_account_type,omitempty"`
	Userid          *string `json:"userid,omitempty"`
	Username        *string `json:"username,omitempty"`
	Uuid            *string `json:"uuid,omitempty"`
}

// V1UsersUseridGet200ResponseDepartmentListInner struct for V1UsersUseridGet200ResponseDepartmentListInner
type V1UsersUseridGet200ResponseDepartmentListInner struct {
	DepartmentId   *string `json:"department_id,omitempty"`
	DepartmentName *string `json:"department_name,omitempty"`
}

// V1UsersUseridInviteAuthPutRequest struct for V1UsersUseridInviteAuthPutRequest
type V1UsersUseridInviteAuthPutRequest struct {
	// 操作者 ID。operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。
	OperatorId string `json:"operator_id"`
	// 操作者 ID 的类型： 1：userid
	OperatorIdType int64 `json:"operator_id_type"`
}

// V1UsersUseridPutRequest struct for V1UsersUseridPutRequest
type V1UsersUseridPutRequest struct {
	Area      *string `json:"area,omitempty"`
	AvatarUrl *string `json:"avatar_url,omitempty"`
	Email     *string `json:"email,omitempty"`
	EntryTime *int64  `json:"entry_time,omitempty"`
	JobTitle  *string `json:"job_title,omitempty"`
	// 操作者ID
	OperatorId string `json:"operator_id"`
	// 操作者ID类型，1:userid
	OperatorIdType int64   `json:"operator_id_type"`
	Phone          *string `json:"phone,omitempty"`
	StaffId        *string `json:"staff_id,omitempty"`
	// 1：高级账号 2：免费账号 3：免费账号100方 4:高级账号300方，5:高级账号500方，6：高级账号1000方，7:高级账号2000方 其中企业版/教育版：1，2 。免费组织 2。 商业版：2-7 根据传入的参数判断是否有该类型账号，没有则报错。更新后，原类型账号资源释放。
	UserAccountType *int64  `json:"user_account_type,omitempty"`
	Userid          *string `json:"userid,omitempty"`
	Username        *string `json:"username,omitempty"`
}
