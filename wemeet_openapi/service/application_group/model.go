// Package wemeetopenapi is auto generate
/*
   腾讯会议OpenAPI

   SAAS版RESTFUL风格API

   API version: v1.0.7
*/
package wemeetopenapi

// V1AppToolkitPostRequest struct for V1AppToolkitPostRequest
type V1AppToolkitPostRequest struct {
	// 自动打开应用的id
	AutoOpenSdkid *string `json:"auto_open_sdkid,omitempty"`
	// 用户的终端设备类型： 1 - PC 2 - Mac 3 - Android 4 - iOS 5 - Web 6 - iPad 7 - Android Pad 8 - 小程序
	Instanceid int64 `json:"instanceid"`
	// 会议id
	MeetingId string `json:"meeting_id"`
	// 工具箱应用列表
	ToolList []V1AppToolkitPostRequestToolListInner `json:"tool_list"`
	// 外显在会中工具栏的应用id（需要保证在tool_list列表中，且列表中的可见范围对此设置也生效）
	ToolbarSdkid *string `json:"toolbar_sdkid,omitempty"`
	// 调用方用于标示用户的唯一 ID（企业内部请使用企业唯一用户标识；OAuth2.0鉴权用户请使用openId） 企业唯一用户标识说明： 1、 企业对接SSO时使用的员工唯一标识ID 2、 企业调用创建用户接口时传递的userid参数
	Userid string `json:"userid"`
}

// V1AppToolkitPostRequestToolListInner 应用列表对象
type V1AppToolkitPostRequestToolListInner struct {
	// 应用是否可以拉取机器人，0:否，1:是，默认为0
	EnableAddRobot *int64 `json:"enable_add_robot,omitempty"`
	// 应用是否可以查询customerData，0:否，1:是，默认为0
	EnableCustomerData *int64 `json:"enable_customer_data,omitempty"`
	// 可见用户是否屏蔽会议创建者，visible_type=2时该字段才有效。true：屏蔽会议创建者，除非在可见用户列表中显式设置会议创建者；false：默认配置，会议创建者可见
	IsShieldCreator *bool `json:"is_shield_creator,omitempty"`
	// 工具箱应用所属企业appid
	ToolAppid string `json:"tool_appid"`
	// 工具箱应用id
	ToolSdkid string `json:"tool_sdkid"`
	// 调用方业务相关字段，最大128个字符
	UniqueCode *string `json:"unique_code,omitempty"`
	// 可见用户列表（默认会议创建者可见），visible_type=2时该字段才有效
	VisibleList []V1AppToolkitPostRequestToolListInnerVisibleListInner `json:"visible_list,omitempty"`
	// 工具箱应用可见类型。 0:所有人可见, 1:本企业可见, 2:指定用户可见，默认为0
	VisibleType *int64 `json:"visible_type,omitempty"`
}

// V1AppToolkitPostRequestToolListInnerVisibleListInner 可见列表对象
type V1AppToolkitPostRequestToolListInnerVisibleListInner struct {
	// 对哪个企业的用户可见
	VisibleAppid *string `json:"visible_appid,omitempty"`
	// 可见用户openid，OAuth2.0鉴权用户请用此字段（visible_userid和visible_openid二者选一，同时存在时以visible_openid为准）
	VisibleOpenid *string `json:"visible_openid,omitempty"`
	// 可见用户userid，若不填则对该企业下所有用户可见
	VisibleUserid *string `json:"visible_userid,omitempty"`
}
