// Package wemeetopenapi is auto generate
/*
    腾讯会议OpenAPI

    SAAS版RESTFUL风格API

    API version: 1.0.0
*/
package wemeetopenapi

import (
    "context"
    "fmt"

    core "github.com/TencentCloud/wemeet-openapi-sdk-go/wemeet_openapi/core"
    xhttp "github.com/TencentCloud/wemeet-openapi-sdk-go/wemeet_openapi/core/xhttp"
)

type Service interface {
/*
V1GuestsGet 通过会议 Code 查询会议嘉宾列表[/v1/guests - Get]

通过会议 Code 查询会议嘉宾列表，只有会议创建人才有权限查询，支持 OAuth2.0 鉴权访问。

> 注意 只有商业版、企业版或教育版用户可以使用会议嘉宾功能，个人版尚无此功能。

*/
    V1GuestsGet(ctx context.Context, request *ApiV1GuestsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1GuestsGetResponse, err error)

/*
V1GuestsMeetingIdGet 通过会议 ID 查询会议嘉宾列表[/v1/guests/{meeting_id} - Get]

通过会议 ID 查询会议嘉宾列表，只有会议创建人才有权限查询，支持 OAuth2.0 鉴权访问。

> 注意
> 只有商业版、企业版或教育版用户可以使用会议嘉宾功能，个人版尚无此功能。

*/
    V1GuestsMeetingIdGet(ctx context.Context, request *ApiV1GuestsMeetingIdGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1GuestsMeetingIdGetResponse, err error)

/*
V1GuestsMeetingIdPut 通过会议 ID 修改会议嘉宾列表[/v1/guests/{meeting_id} - Put]

通过会议 ID 修改会议嘉宾列表，只有会议创建人才有权限修改，支持 OAuth2.0 鉴权访问。

*

> 注意
> 只有商业版、企业版或教育版用户可以使用会议嘉宾功能，个人版尚无此功能。

*/
    V1GuestsMeetingIdPut(ctx context.Context, request *ApiV1GuestsMeetingIdPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1GuestsMeetingIdPutResponse, err error)

/*
V1GuestsPut 通过会议 Code 修改会议嘉宾列表[/v1/guests - Put]

通过会议 Code 修改会议嘉宾列表，只有会议创建人才有权限修改，支持 OAuth2.0 鉴权访问。一天最多可设置嘉宾10次，且同一用户最多被设置为嘉宾10次。超出限制会议将不会出现在嘉宾会议列表中。

> 注意 只有商业版、企业版或教育版用户可以使用会议嘉宾功能，个人版尚无此功能。

*/
    V1GuestsPut(ctx context.Context, request *ApiV1GuestsPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1GuestsPutResponse, err error)

}

type meetingGuestAPIService struct {
    config *core.Config
}

func NewService(config *core.Config) Service {
    return &meetingGuestAPIService{
        config: config,
    }
}


type ApiV1GuestsGetRequest struct {
    // 会议 ID。
    MeetingCode *string `json:"-"`
    // 用户的 ID（企业内部请使用企业唯一用户标识，OAuth2.0 鉴权用户请使用 openId）。
    Userid *string `json:"-"`
    // 用户的终端设备类型： 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序。
    Instanceid *string `json:"-"`
    Body *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1GuestsGetResponse struct {
    *xhttp.ApiResponse
    Data *V1GuestsGet200Response `json:"data,omitempty"`
}

/*
V1GuestsGet 通过会议 Code 查询会议嘉宾列表[/v1/guests - Get]

通过会议 Code 查询会议嘉宾列表，只有会议创建人才有权限查询，支持 OAuth2.0 鉴权访问。

> 注意 只有商业版、企业版或教育版用户可以使用会议嘉宾功能，个人版尚无此功能。

*/
func (s *meetingGuestAPIService) V1GuestsGet(ctx context.Context, request *ApiV1GuestsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1GuestsGetResponse, err error) {
    apiReq := &xhttp.ApiRequest{
        ApiURI:      "/v1/guests",
        Body:        request.Body,
        PathParams:  xhttp.PathParams{},
        QueryParams: xhttp.QueryParams{},
    }

    if request.MeetingCode == nil {
        return nil, fmt.Errorf("meeting_code is required and must be specified")
    }

    if request.Userid == nil {
        return nil, fmt.Errorf("userid is required and must be specified")
    }

    if request.Instanceid == nil {
        return nil, fmt.Errorf("instanceid is required and must be specified")
    }

    // path 参数
    // query 参数
    if request.MeetingCode != nil {
        apiReq.QueryParams.Set("meeting_code", core.QueryValue(request.MeetingCode))
    }
    if request.Userid != nil {
        apiReq.QueryParams.Set("userid", core.QueryValue(request.Userid))
    }
    if request.Instanceid != nil {
        apiReq.QueryParams.Set("instanceid", core.QueryValue(request.Instanceid))
    }
    // 转换 options
    var httpOptions []xhttp.RequestOptionFunc
    for _, opt := range opts {
        httpOptions = append(httpOptions, opt(*s.config))
    }
    // 增加 SDK Version 标识
    httpOptions = append(httpOptions, xhttp.WithRequestAuthenticator(core.DefaultAuthenticator))

    apiRsp, err := s.config.Clt.Get(ctx, apiReq, httpOptions...)
    if err != nil {
        return nil, err
    }

    if apiRsp.StatusCode >= 300 {
        var svrErr = &core.ServiceError{
            ApiResponse: apiRsp,
        }
        _ = core.JsonSerializer.Deserialize(apiRsp.RawBody, svrErr)
        return nil, svrErr
    }

    response = &ApiV1GuestsGetResponse{
        ApiResponse: apiRsp,
        Data:        new(V1GuestsGet200Response),
    }
    if err = apiRsp.Translate(response.Data); err != nil {
        return nil, &core.ClientError{
            Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
                apiRsp.StatusCode, apiRsp.RawBody, err),
        }
    }
    return
}

type ApiV1GuestsMeetingIdGetRequest struct {
    MeetingId string `json:"-"`
    // 用户的 ID（企业内部请使用企业唯一用户标识，OAuth2.0 鉴权用户请使用 openId）。
    Userid *string `json:"-"`
    // 用户的终端设备类型： 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序。9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch MacOS 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch iOS
    Instanceid *string `json:"-"`
    Body *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1GuestsMeetingIdGetResponse struct {
    *xhttp.ApiResponse
    Data *V1GuestsMeetingIdGet200Response `json:"data,omitempty"`
}

/*
V1GuestsMeetingIdGet 通过会议 ID 查询会议嘉宾列表[/v1/guests/{meeting_id} - Get]

通过会议 ID 查询会议嘉宾列表，只有会议创建人才有权限查询，支持 OAuth2.0 鉴权访问。

> 注意
> 只有商业版、企业版或教育版用户可以使用会议嘉宾功能，个人版尚无此功能。

*/
func (s *meetingGuestAPIService) V1GuestsMeetingIdGet(ctx context.Context, request *ApiV1GuestsMeetingIdGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1GuestsMeetingIdGetResponse, err error) {
    apiReq := &xhttp.ApiRequest{
        ApiURI:      "/v1/guests/{meeting_id}",
        Body:        request.Body,
        PathParams:  xhttp.PathParams{},
        QueryParams: xhttp.QueryParams{},
    }

    if request.Userid == nil {
        return nil, fmt.Errorf("userid is required and must be specified")
    }

    if request.Instanceid == nil {
        return nil, fmt.Errorf("instanceid is required and must be specified")
    }

    // path 参数
    apiReq.PathParams.Set("meeting_id", core.PathValue(request.MeetingId))
    // query 参数
    if request.Userid != nil {
        apiReq.QueryParams.Set("userid", core.QueryValue(request.Userid))
    }
    if request.Instanceid != nil {
        apiReq.QueryParams.Set("instanceid", core.QueryValue(request.Instanceid))
    }
    // 转换 options
    var httpOptions []xhttp.RequestOptionFunc
    for _, opt := range opts {
        httpOptions = append(httpOptions, opt(*s.config))
    }
    // 增加 SDK Version 标识
    httpOptions = append(httpOptions, xhttp.WithRequestAuthenticator(core.DefaultAuthenticator))

    apiRsp, err := s.config.Clt.Get(ctx, apiReq, httpOptions...)
    if err != nil {
        return nil, err
    }

    if apiRsp.StatusCode >= 300 {
        var svrErr = &core.ServiceError{
            ApiResponse: apiRsp,
        }
        _ = core.JsonSerializer.Deserialize(apiRsp.RawBody, svrErr)
        return nil, svrErr
    }

    response = &ApiV1GuestsMeetingIdGetResponse{
        ApiResponse: apiRsp,
        Data:        new(V1GuestsMeetingIdGet200Response),
    }
    if err = apiRsp.Translate(response.Data); err != nil {
        return nil, &core.ClientError{
            Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
                apiRsp.StatusCode, apiRsp.RawBody, err),
        }
    }
    return
}

type ApiV1GuestsMeetingIdPutRequest struct {
    MeetingId string `json:"-"`
    Body *V1GuestsMeetingIdPutRequest `json:"body,omitempty"`
}

type ApiV1GuestsMeetingIdPutResponse struct {
    *xhttp.ApiResponse
    Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1GuestsMeetingIdPut 通过会议 ID 修改会议嘉宾列表[/v1/guests/{meeting_id} - Put]

通过会议 ID 修改会议嘉宾列表，只有会议创建人才有权限修改，支持 OAuth2.0 鉴权访问。

*

> 注意
> 只有商业版、企业版或教育版用户可以使用会议嘉宾功能，个人版尚无此功能。

*/
func (s *meetingGuestAPIService) V1GuestsMeetingIdPut(ctx context.Context, request *ApiV1GuestsMeetingIdPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1GuestsMeetingIdPutResponse, err error) {
    apiReq := &xhttp.ApiRequest{
        ApiURI:      "/v1/guests/{meeting_id}",
        Body:        request.Body,
        PathParams:  xhttp.PathParams{},
        QueryParams: xhttp.QueryParams{},
    }

    // path 参数
    apiReq.PathParams.Set("meeting_id", core.PathValue(request.MeetingId))
    // query 参数
    // 转换 options
    var httpOptions []xhttp.RequestOptionFunc
    for _, opt := range opts {
        httpOptions = append(httpOptions, opt(*s.config))
    }
    // 增加 SDK Version 标识
    httpOptions = append(httpOptions, xhttp.WithRequestAuthenticator(core.DefaultAuthenticator))

    apiRsp, err := s.config.Clt.Put(ctx, apiReq, httpOptions...)
    if err != nil {
        return nil, err
    }

    if apiRsp.StatusCode >= 300 {
        var svrErr = &core.ServiceError{
            ApiResponse: apiRsp,
        }
        _ = core.JsonSerializer.Deserialize(apiRsp.RawBody, svrErr)
        return nil, svrErr
    }

    response = &ApiV1GuestsMeetingIdPutResponse{
        ApiResponse: apiRsp,
        Data:        new(map[string]interface{}),
    }
    if err = apiRsp.Translate(response.Data); err != nil {
        return nil, &core.ClientError{
            Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
                apiRsp.StatusCode, apiRsp.RawBody, err),
        }
    }
    return
}

type ApiV1GuestsPutRequest struct {
    Body *V1GuestsPutRequest `json:"body,omitempty"`
}

type ApiV1GuestsPutResponse struct {
    *xhttp.ApiResponse
    Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1GuestsPut 通过会议 Code 修改会议嘉宾列表[/v1/guests - Put]

通过会议 Code 修改会议嘉宾列表，只有会议创建人才有权限修改，支持 OAuth2.0 鉴权访问。一天最多可设置嘉宾10次，且同一用户最多被设置为嘉宾10次。超出限制会议将不会出现在嘉宾会议列表中。

> 注意 只有商业版、企业版或教育版用户可以使用会议嘉宾功能，个人版尚无此功能。

*/
func (s *meetingGuestAPIService) V1GuestsPut(ctx context.Context, request *ApiV1GuestsPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1GuestsPutResponse, err error) {
    apiReq := &xhttp.ApiRequest{
        ApiURI:      "/v1/guests",
        Body:        request.Body,
        PathParams:  xhttp.PathParams{},
        QueryParams: xhttp.QueryParams{},
    }

    // path 参数
    // query 参数
    // 转换 options
    var httpOptions []xhttp.RequestOptionFunc
    for _, opt := range opts {
        httpOptions = append(httpOptions, opt(*s.config))
    }
    // 增加 SDK Version 标识
    httpOptions = append(httpOptions, xhttp.WithRequestAuthenticator(core.DefaultAuthenticator))

    apiRsp, err := s.config.Clt.Put(ctx, apiReq, httpOptions...)
    if err != nil {
        return nil, err
    }

    if apiRsp.StatusCode >= 300 {
        var svrErr = &core.ServiceError{
            ApiResponse: apiRsp,
        }
        _ = core.JsonSerializer.Deserialize(apiRsp.RawBody, svrErr)
        return nil, svrErr
    }

    response = &ApiV1GuestsPutResponse{
        ApiResponse: apiRsp,
        Data:        new(map[string]interface{}),
    }
    if err = apiRsp.Translate(response.Data); err != nil {
        return nil, &core.ClientError{
            Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
                apiRsp.StatusCode, apiRsp.RawBody, err),
        }
    }
    return
}

