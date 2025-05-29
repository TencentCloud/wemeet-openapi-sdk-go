// Package wemeetopenapi is auto generate
/*
    腾讯会议OpenAPI

    SAAS版RESTFUL风格API

    API version: v1.0.8
*/
package wemeetopenapi


// V1MeetingMeetingIdPhoneCalloutPost200Response struct for V1MeetingMeetingIdPhoneCalloutPost200Response
type V1MeetingMeetingIdPhoneCalloutPost200Response struct {
    // 外呼的电话号码对象列表。
    InvalidPhoneNumbers []V1MeetingMeetingIdPhoneCalloutPost200ResponseInvalidPhoneNumbersInner `json:"invalid_phone_numbers,omitempty"`
    // 会议的唯一ID
    MeetingId *string `json:"meeting_id,omitempty"`
    // 外呼的电话号码对象列表。
    PhoneNumbers []V1MeetingMeetingIdPhoneCalloutPost200ResponseInvalidPhoneNumbersInner `json:"phone_numbers,omitempty"`
}

// V1MeetingMeetingIdPhoneCalloutPost200ResponseInvalidPhoneNumbersInner 电话号码对象
type V1MeetingMeetingIdPhoneCalloutPost200ResponseInvalidPhoneNumbersInner struct {
    // 电话区号
    Area *int64 `json:"area,omitempty"`
    ExtensionNumber *string `json:"extension_number,omitempty"`
    Phone *string `json:"phone,omitempty"`
}

// V1MeetingMeetingIdPhoneCalloutPostRequest struct for V1MeetingMeetingIdPhoneCalloutPostRequest
type V1MeetingMeetingIdPhoneCalloutPostRequest struct {
    // 操作者 ID。会议创建者、主持人、联席主持人可以调用该接口。 operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。
    OperatorId string `json:"operator_id"`
    // 操作者 ID 的类型： 1：企业内用户 userid。 2: openid 3. rooms_id 4: ms_open_id
    OperatorIdType int64 `json:"operator_id_type"`
    // 外呼的电话号码对象列表。
    PhoneNumbers []V1MeetingMeetingIdPhoneCalloutPostRequestPhoneNumbersInner `json:"phone_numbers"`
}

// V1MeetingMeetingIdPhoneCalloutPostRequestPhoneNumbersInner 电话号码对象
type V1MeetingMeetingIdPhoneCalloutPostRequestPhoneNumbersInner struct {
    // 电话区号
    Area int64 `json:"area"`
    // 国家/地区代码。（例如：中国是86） 当前仅支持呼叫中国大陆、中国香港、美国的号码。
    CallingPartyArea *int64 `json:"calling_party_area,omitempty"`
    // 电话号码或固定电话总机号。
    CallingPartyPhoneNumber *string `json:"calling_party_phone_number,omitempty"`
    ExtensionNumber *string `json:"extension_number,omitempty"`
    NickName *string `json:"nick_name,omitempty"`
    Phone string `json:"phone"`
}

// V1MeetingMeetingIdPhoneCancelcallPost200Response struct for V1MeetingMeetingIdPhoneCancelcallPost200Response
type V1MeetingMeetingIdPhoneCancelcallPost200Response struct {
    // 取消呼叫失败的列表
    FailedList []V1MeetingMeetingIdPhoneCancelcallPost200ResponseFailedListInner `json:"failed_list,omitempty"`
    // 取消呼叫成功的列表
    SuccessList []V1MeetingMeetingIdPhoneCancelcallPost200ResponseSuccessListInner `json:"success_list,omitempty"`
}

// V1MeetingMeetingIdPhoneCancelcallPost200ResponseFailedListInner struct for V1MeetingMeetingIdPhoneCancelcallPost200ResponseFailedListInner
type V1MeetingMeetingIdPhoneCancelcallPost200ResponseFailedListInner struct {
    // 国家/地区代码。（例如：中国是86）
    Area *int64 `json:"area,omitempty"`
    // 错误信息
    ErrorMsg *string `json:"error_msg,omitempty"`
    // 固定电话分机号。
    ExtensionNumber *string `json:"extension_number,omitempty"`
    // 电话号码或固定电话总机号。
    Phone *string `json:"phone,omitempty"`
}

// V1MeetingMeetingIdPhoneCancelcallPost200ResponseSuccessListInner struct for V1MeetingMeetingIdPhoneCancelcallPost200ResponseSuccessListInner
type V1MeetingMeetingIdPhoneCancelcallPost200ResponseSuccessListInner struct {
    // 国家/地区代码。（例如：中国是86）
    Area *int64 `json:"area,omitempty"`
    // 固定电话分机号。
    ExtensionNumber *string `json:"extension_number,omitempty"`
    // 电话号码或固定电话总机号。
    Phone *string `json:"phone,omitempty"`
}

// V1MeetingMeetingIdPhoneCancelcallPostRequest struct for V1MeetingMeetingIdPhoneCancelcallPostRequest
type V1MeetingMeetingIdPhoneCancelcallPostRequest struct {
    // 操作者 ID。会议创建者、主持人、联席主持人可以调用该接口。 operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。
    OperatorId string `json:"operator_id"`
    // 操作者 ID 的类型： 1：企业内用户 userid。 
    OperatorIdType int64 `json:"operator_id_type"`
    // 电话号码对象数组
    PhoneNumbers []V1MeetingMeetingIdPhoneCancelcallPostRequestPhoneNumbersInner `json:"phone_numbers"`
}

// V1MeetingMeetingIdPhoneCancelcallPostRequestPhoneNumbersInner struct for V1MeetingMeetingIdPhoneCancelcallPostRequestPhoneNumbersInner
type V1MeetingMeetingIdPhoneCancelcallPostRequestPhoneNumbersInner struct {
    // 国家/地区代码。（例如：中国是86）
    Area int64 `json:"area"`
    // 固定电话分机号。
    ExtensionNumber *string `json:"extension_number,omitempty"`
    // 电话号码或固定电话总机号。
    Phone string `json:"phone"`
}
