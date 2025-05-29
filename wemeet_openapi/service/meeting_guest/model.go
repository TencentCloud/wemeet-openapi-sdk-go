// Package wemeetopenapi is auto generate
/*
    腾讯会议OpenAPI

    SAAS版RESTFUL风格API

    API version: v1.0.8
*/
package wemeetopenapi


// V1GuestsMeetingIdGet200Response struct for V1GuestsMeetingIdGet200Response
type V1GuestsMeetingIdGet200Response struct {
    // 会议嘉宾列表数组。
    Guests []V1GuestsMeetingIdGet200ResponseGuestsInner `json:"guests,omitempty"`
    // 会议 Code。
    MeetingCode *string `json:"meeting_code,omitempty"`
    // 会议 ID。
    MeetingId *string `json:"meeting_id,omitempty"`
    // 会议主题。
    Subject *string `json:"subject,omitempty"`
}

// V1GuestsMeetingIdGet200ResponseGuestsInner struct for V1GuestsMeetingIdGet200ResponseGuestsInner
type V1GuestsMeetingIdGet200ResponseGuestsInner struct {
    // 国家/地区代码（例如：中国传86，不是+86，也不是0086）。
    Area *string `json:"area,omitempty"`
    // 嘉宾名称。
    GuestName *string `json:"guest_name,omitempty"`
    // 手机号。
    PhoneNumber *string `json:"phone_number,omitempty"`
}

// V1GuestsMeetingIdPutRequest struct for V1GuestsMeetingIdPutRequest
type V1GuestsMeetingIdPutRequest struct {
    //  会议嘉宾列表（传空数组会清空嘉宾列表）。
    Guests []V1GuestsMeetingIdPutRequestGuestsInner `json:"guests"`
    // 用户的终端设备类型： 0：PSTN 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch MacOS 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch iOS
    Instanceid int64 `json:"instanceid"`
    // 用户的 ID（企业内部请使用企业唯一用户标识，OAuth2.0 鉴权用户请使用 openId）。
    Userid string `json:"userid"`
}

// V1GuestsMeetingIdPutRequestGuestsInner struct for V1GuestsMeetingIdPutRequestGuestsInner
type V1GuestsMeetingIdPutRequestGuestsInner struct {
    // 国家/地区代码（例如：中国传86，不是+86，也不是0086）。
    Area string `json:"area"`
    // 嘉宾名称
    GuestName *string `json:"guest_name,omitempty"`
    // 手机号
    PhoneNumber string `json:"phone_number"`
}
