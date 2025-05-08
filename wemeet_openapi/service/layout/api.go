// Package wemeetopenapi is auto generate
/*
   腾讯会议OpenAPI

   SAAS版RESTFUL风格API

   API version: v1.0.7
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
	   V1MeetingsMeetingIdAdvancedLayoutsPost 添加自定义布局[/v1/meetings/{meeting_id}/advanced-layouts - Post]

	   **描述：**

	   * 对当前会议添加高级自定义布局，支持批量添加。
	   * 用户座次设置需设置参会成员。
	   * 单个会议最多允许添加20个布局。
	   * 目前暂不支持 OAuth2.0 鉴权访问。
	   * 目前仅会应用于 H.323/SIP 终端

	*/
	V1MeetingsMeetingIdAdvancedLayoutsPost(ctx context.Context, request *ApiV1MeetingsMeetingIdAdvancedLayoutsPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdAdvancedLayoutsPostResponse, err error)

	/*
	   V1MeetingsMeetingIdApplyingLayoutPut 应用布局[/v1/meetings/{meeting_id}/applying-layout - Put]

	   **描述：**

	   * 将会议中的高级自定义布局应用到指定成员或者整个会议。
	   * 恢复指定成员或整个会议的默认布局。
	   * 目前暂不支持 OAuth2.0 鉴权访问。
	   * 目前仅会应用于 H.323/SIP 终端。

	   <span class="colour" style="color:rgb(51, 51, 51)"></span>

	*/
	V1MeetingsMeetingIdApplyingLayoutPut(ctx context.Context, request *ApiV1MeetingsMeetingIdApplyingLayoutPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdApplyingLayoutPutResponse, err error)

	/*
	   V1MeetingsMeetingIdLayoutsPost 添加会议布局[/v1/meetings/{meeting_id}/layouts - Post]

	   对成功预定的会议添加会议布局，支持多个布局的添加，每个布局支持多页模板，默认选中第一页模板作为该布局的首页进行展示。

	   * 用户座次设置区分会前和会中两种方式：会前只允许设置邀请者成员，会中只允许设置参会成员。
	   * 一场会议最多添加10个布局，添加成功返回新增的会议布局信息。
	   * 目前暂不支持 OAuth2.0 鉴权访问。

	*/
	V1MeetingsMeetingIdLayoutsPost(ctx context.Context, request *ApiV1MeetingsMeetingIdLayoutsPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdLayoutsPostResponse, err error)
}

type layoutAPIService struct {
	config *core.Config
}

func NewService(config *core.Config) Service {
	return &layoutAPIService{
		config: config,
	}
}

type ApiV1MeetingsMeetingIdAdvancedLayoutsPostRequest struct {
	// 会议ID
	MeetingId string                                         `json:"-"`
	Body      *V1MeetingsMeetingIdAdvancedLayoutsPostRequest `json:"body,omitempty"`
}

type ApiV1MeetingsMeetingIdAdvancedLayoutsPostResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingsMeetingIdAdvancedLayoutsPost200Response `json:"data,omitempty"`
}

/*
V1MeetingsMeetingIdAdvancedLayoutsPost 添加自定义布局[/v1/meetings/{meeting_id}/advanced-layouts - Post]

**描述：**

* 对当前会议添加高级自定义布局，支持批量添加。
* 用户座次设置需设置参会成员。
* 单个会议最多允许添加20个布局。
* 目前暂不支持 OAuth2.0 鉴权访问。
* 目前仅会应用于 H.323/SIP 终端
*/
func (s *layoutAPIService) V1MeetingsMeetingIdAdvancedLayoutsPost(ctx context.Context, request *ApiV1MeetingsMeetingIdAdvancedLayoutsPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdAdvancedLayoutsPostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings/{meeting_id}/advanced-layouts",
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

	response = &ApiV1MeetingsMeetingIdAdvancedLayoutsPostResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingsMeetingIdAdvancedLayoutsPost200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingsMeetingIdApplyingLayoutPutRequest struct {
	// 会议ID
	MeetingId string                                       `json:"-"`
	Body      *V1MeetingsMeetingIdApplyingLayoutPutRequest `json:"body,omitempty"`
}

type ApiV1MeetingsMeetingIdApplyingLayoutPutResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1MeetingsMeetingIdApplyingLayoutPut 应用布局[/v1/meetings/{meeting_id}/applying-layout - Put]

**描述：**

* 将会议中的高级自定义布局应用到指定成员或者整个会议。
* 恢复指定成员或整个会议的默认布局。
* 目前暂不支持 OAuth2.0 鉴权访问。
* 目前仅会应用于 H.323/SIP 终端。

<span class="colour" style="color:rgb(51, 51, 51)"></span>
*/
func (s *layoutAPIService) V1MeetingsMeetingIdApplyingLayoutPut(ctx context.Context, request *ApiV1MeetingsMeetingIdApplyingLayoutPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdApplyingLayoutPutResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings/{meeting_id}/applying-layout",
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

	response = &ApiV1MeetingsMeetingIdApplyingLayoutPutResponse{
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

type ApiV1MeetingsMeetingIdLayoutsPostRequest struct {
	// 会议ID
	MeetingId string                                 `json:"-"`
	Body      *V1MeetingsMeetingIdLayoutsPostRequest `json:"body,omitempty"`
}

type ApiV1MeetingsMeetingIdLayoutsPostResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingsMeetingIdLayoutsPost200Response `json:"data,omitempty"`
}

/*
V1MeetingsMeetingIdLayoutsPost 添加会议布局[/v1/meetings/{meeting_id}/layouts - Post]

对成功预定的会议添加会议布局，支持多个布局的添加，每个布局支持多页模板，默认选中第一页模板作为该布局的首页进行展示。

* 用户座次设置区分会前和会中两种方式：会前只允许设置邀请者成员，会中只允许设置参会成员。
* 一场会议最多添加10个布局，添加成功返回新增的会议布局信息。
* 目前暂不支持 OAuth2.0 鉴权访问。
*/
func (s *layoutAPIService) V1MeetingsMeetingIdLayoutsPost(ctx context.Context, request *ApiV1MeetingsMeetingIdLayoutsPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdLayoutsPostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings/{meeting_id}/layouts",
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

	response = &ApiV1MeetingsMeetingIdLayoutsPostResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingsMeetingIdLayoutsPost200Response),
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
