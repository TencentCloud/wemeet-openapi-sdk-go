// Package wemeetopenapi is auto generate
/*
    腾讯会议OpenAPI

    SAAS版RESTFUL风格API

    API version: v1.0.8
*/
package wemeetopenapi

import (
    "bytes"
    "context"
    "fmt"

    core "github.com/TencentCloud/wemeet-openapi-sdk-go/wemeet_openapi/core"
    xhttp "github.com/TencentCloud/wemeet-openapi-sdk-go/wemeet_openapi/core/xhttp"
    "io"
    "mime/multipart"
    "net/http"
)

type Service interface {
/*
V1MeetingMeetingIdPhoneCalloutPost 批量外呼[/v1/meeting/{meeting_id}/phone/callout - Post]

**描述**：

* 会议创建者、主持人、联席主持人可以批量外呼电话入会。
* 在拨打后，立刻返回，无需等待，客户通过查询接口和 Webhook 获得外呼状态。
* 需要支持在会议未开始、会中外呼。
* 每次调用支持批量外呼50路。
* \*\*鉴权方式：\*\*JWT 鉴权

*/
    V1MeetingMeetingIdPhoneCalloutPost(ctx context.Context, request *ApiV1MeetingMeetingIdPhoneCalloutPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingMeetingIdPhoneCalloutPostResponse, err error)

/*
V1MeetingMeetingIdPhoneCancelcallPost 批量取消外呼[/v1/meeting/{meeting_id}/phone/cancelcall - Post]

*/
    V1MeetingMeetingIdPhoneCancelcallPost(ctx context.Context, request *ApiV1MeetingMeetingIdPhoneCancelcallPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingMeetingIdPhoneCancelcallPostResponse, err error)

}

type pstnAPIService struct {
    config *core.Config
}

func NewService(config *core.Config) Service {
    return &pstnAPIService{
        config: config,
    }
}


type ApiV1MeetingMeetingIdPhoneCalloutPostRequest struct {
    // 会议的唯一ID
    MeetingId string `json:"-"`
    Body *V1MeetingMeetingIdPhoneCalloutPostRequest `json:"body,omitempty"`
}

type ApiV1MeetingMeetingIdPhoneCalloutPostResponse struct {
    *xhttp.ApiResponse
    Data *V1MeetingMeetingIdPhoneCalloutPost200Response `json:"data,omitempty"`
}

/*
V1MeetingMeetingIdPhoneCalloutPost 批量外呼[/v1/meeting/{meeting_id}/phone/callout - Post]

**描述**：

* 会议创建者、主持人、联席主持人可以批量外呼电话入会。
* 在拨打后，立刻返回，无需等待，客户通过查询接口和 Webhook 获得外呼状态。
* 需要支持在会议未开始、会中外呼。
* 每次调用支持批量外呼50路。
* \*\*鉴权方式：\*\*JWT 鉴权

*/
func (s *pstnAPIService) V1MeetingMeetingIdPhoneCalloutPost(ctx context.Context, request *ApiV1MeetingMeetingIdPhoneCalloutPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingMeetingIdPhoneCalloutPostResponse, err error) {
    apiReq := &xhttp.ApiRequest{
        ApiURI:      "/v1/meeting/{meeting_id}/phone/callout",
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





    apiRsp, err := s.config.Clt.Post(ctx, apiReq, httpOptions...)
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

    response = &ApiV1MeetingMeetingIdPhoneCalloutPostResponse{
        ApiResponse: apiRsp,
        Data:        new(V1MeetingMeetingIdPhoneCalloutPost200Response),
    }
    if err = apiRsp.Translate(response.Data); err != nil {
        return nil, &core.ClientError{
            Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
                apiRsp.StatusCode, apiRsp.RawBody, err),
        }
    }
    return
}

type ApiV1MeetingMeetingIdPhoneCancelcallPostRequest struct {
    MeetingId string `json:"-"`
    Body *V1MeetingMeetingIdPhoneCancelcallPostRequest `json:"body,omitempty"`
}

type ApiV1MeetingMeetingIdPhoneCancelcallPostResponse struct {
    *xhttp.ApiResponse
    Data *V1MeetingMeetingIdPhoneCancelcallPost200Response `json:"data,omitempty"`
}

/*
V1MeetingMeetingIdPhoneCancelcallPost 批量取消外呼[/v1/meeting/{meeting_id}/phone/cancelcall - Post]

*/
func (s *pstnAPIService) V1MeetingMeetingIdPhoneCancelcallPost(ctx context.Context, request *ApiV1MeetingMeetingIdPhoneCancelcallPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingMeetingIdPhoneCancelcallPostResponse, err error) {
    apiReq := &xhttp.ApiRequest{
        ApiURI:      "/v1/meeting/{meeting_id}/phone/cancelcall",
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





    apiRsp, err := s.config.Clt.Post(ctx, apiReq, httpOptions...)
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

    response = &ApiV1MeetingMeetingIdPhoneCancelcallPostResponse{
        ApiResponse: apiRsp,
        Data:        new(V1MeetingMeetingIdPhoneCancelcallPost200Response),
    }
    if err = apiRsp.Translate(response.Data); err != nil {
        return nil, &core.ClientError{
            Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
                apiRsp.StatusCode, apiRsp.RawBody, err),
        }
    }
    return
}



// GetFromDataDemo fromdata对象构造demo演示
func GetFromDataDemo() {
    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)

    writer.WriteField("operator_id", "xx")
    // 添加文件到FormData
    fileContentPart, err := writer.CreateFormFile("file_content", "zy.mp4")
    if err != nil {
        fmt.Println("创建表单文件失败:", err)
        return
    }

    data := []byte("Hello, World!")
    _, err = io.Copy(fileContentPart, bytes.NewReader(data))
    if err != nil {
        fmt.Println("复制文件内容失败:", err)
    }

    // 设置请求头
    // from-data设置请求头
    _ = writer.Close()
    header := make(http.Header)
    header.Set("Content-Type", writer.FormDataContentType())
}