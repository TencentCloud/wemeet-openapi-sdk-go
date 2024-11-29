// Package wemeetopenapi is auto generate
/*
   腾讯会议OpenAPI

   SAAS版RESTFUL风格API

   API version: v1.0.4
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
	   V1GuestsMeetingIdGet 查询会议嘉宾列表（通过会议 ID 查询）[/v1/guests/{meeting_id} - Get]

	   通过会议 ID 查询会议嘉宾列表，只有会议创建人才有权限查询，支持 OAuth2.0 鉴权访问。

	   > 注意
	   > 只有商业版、企业版或教育版用户可以使用会议嘉宾功能，个人版尚无此功能。

	*/
	V1GuestsMeetingIdGet(ctx context.Context, request *ApiV1GuestsMeetingIdGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1GuestsMeetingIdGetResponse, err error)

	/*
	   V1GuestsMeetingIdPut 修改会议嘉宾列表（通过会议 ID 修改）[/v1/guests/{meeting_id} - Put]

	   通过会议 ID 修改会议嘉宾列表，只有会议创建人才有权限修改，支持 OAuth2.0 鉴权访问。

	   *

	   > 注意
	   > 只有商业版、企业版或教育版用户可以使用会议嘉宾功能，个人版尚无此功能。

	*/
	V1GuestsMeetingIdPut(ctx context.Context, request *ApiV1GuestsMeetingIdPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1GuestsMeetingIdPutResponse, err error)
}

type meetingGuestAPIService struct {
	config *core.Config
}

func NewService(config *core.Config) Service {
	return &meetingGuestAPIService{
		config: config,
	}
}

type ApiV1GuestsMeetingIdGetRequest struct {
	MeetingId string `json:"-"`
	// 用户的 ID（企业内部请使用企业唯一用户标识，OAuth2.0 鉴权用户请使用 openId）。
	Userid *string `json:"-"`
	// 用户的终端设备类型： 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序。9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch MacOS 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch iOS
	Instanceid *string                 `json:"-"`
	Body       *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1GuestsMeetingIdGetResponse struct {
	*xhttp.ApiResponse
	Data *V1GuestsMeetingIdGet200Response `json:"data,omitempty"`
}

/*
V1GuestsMeetingIdGet 查询会议嘉宾列表（通过会议 ID 查询）[/v1/guests/{meeting_id} - Get]

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
	MeetingId string                       `json:"-"`
	Body      *V1GuestsMeetingIdPutRequest `json:"body,omitempty"`
}

type ApiV1GuestsMeetingIdPutResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1GuestsMeetingIdPut 修改会议嘉宾列表（通过会议 ID 修改）[/v1/guests/{meeting_id} - Put]

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
