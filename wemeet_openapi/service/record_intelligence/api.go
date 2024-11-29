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
	   V1SmartChaptersGet 智能章节[/v1/smart/chapters - Get]

	   查询单个云录制的智能章节，支持OAuth2.0鉴权调用，仅支持授权用户为商业版、企业版、教育版。
	   当该录制文件未开启相关智能化功能或内容处于生成中状态时，结果返回错误码

	*/
	V1SmartChaptersGet(ctx context.Context, request *ApiV1SmartChaptersGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1SmartChaptersGetResponse, err error)

	/*
	   V1SmartFullsummaryGet 智能总结[/v1/smart/fullsummary - Get]

	   查询单个云录制的智能总结。支持OAuth2.0鉴权调用，仅支持授权用户为商业版、企业版、教育版。
	   当该录制文件未开启相关智能化功能或内容处于生成中状态时，结果返回错误码

	*/
	V1SmartFullsummaryGet(ctx context.Context, request *ApiV1SmartFullsummaryGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1SmartFullsummaryGetResponse, err error)

	/*
	   V1SmartSpeakersGet 智能发言人[/v1/smart/speakers - Get]

	   查询单个云录制的发言人列表。支持OAuth2.0鉴权调用，仅支持授权用户为商业版、企业版、教育版。
	   当该录制文件未开启相关智能化功能或内容处于生成中状态时，结果返回错误码

	*/
	V1SmartSpeakersGet(ctx context.Context, request *ApiV1SmartSpeakersGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1SmartSpeakersGetResponse, err error)

	/*
	   V1SmartTopicsGet 智能话题[/v1/smart/topics - Get]

	   查询单个云录制的智能话题。支持OAuth2.0鉴权调用，仅支持授权用户为商业版、企业版、教育版。
	   当该录制文件未开启相关智能化功能或内容处于生成中状态时，结果返回错误码

	*/
	V1SmartTopicsGet(ctx context.Context, request *ApiV1SmartTopicsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1SmartTopicsGetResponse, err error)
}

type recordIntelligenceAPIService struct {
	config *core.Config
}

func NewService(config *core.Config) Service {
	return &recordIntelligenceAPIService{
		config: config,
	}
}

type ApiV1SmartChaptersGetRequest struct {
	// 操作者ID。operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。
	OperatorId *string `json:"-"`
	// 操作者ID类型： 1：userid 2:openid
	OperatorIdType *string `json:"-"`
	// 录制文件ID，列表接口返回的 record_file_id。
	RecordFileId *string `json:"-"`
	// 翻译类型，默认原文展示 \"default\" ：原文（不翻译） \"zh\" ：简体中文 \"en\" ：英文 \"ja\"： 日语
	Lang *string                 `json:"-"`
	Pwd  *string                 `json:"-"`
	Body *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1SmartChaptersGetResponse struct {
	*xhttp.ApiResponse
	Data *V1SmartChaptersGet200Response `json:"data,omitempty"`
}

/*
V1SmartChaptersGet 智能章节[/v1/smart/chapters - Get]

查询单个云录制的智能章节，支持OAuth2.0鉴权调用，仅支持授权用户为商业版、企业版、教育版。
当该录制文件未开启相关智能化功能或内容处于生成中状态时，结果返回错误码
*/
func (s *recordIntelligenceAPIService) V1SmartChaptersGet(ctx context.Context, request *ApiV1SmartChaptersGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1SmartChaptersGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/smart/chapters",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	if request.OperatorId == nil {
		return nil, fmt.Errorf("operator_id is required and must be specified")
	}

	if request.OperatorIdType == nil {
		return nil, fmt.Errorf("operator_id_type is required and must be specified")
	}

	if request.RecordFileId == nil {
		return nil, fmt.Errorf("record_file_id is required and must be specified")
	}

	// path 参数
	// query 参数
	if request.OperatorId != nil {
		apiReq.QueryParams.Set("operator_id", core.QueryValue(request.OperatorId))
	}
	if request.OperatorIdType != nil {
		apiReq.QueryParams.Set("operator_id_type", core.QueryValue(request.OperatorIdType))
	}
	if request.RecordFileId != nil {
		apiReq.QueryParams.Set("record_file_id", core.QueryValue(request.RecordFileId))
	}
	if request.Lang != nil {
		apiReq.QueryParams.Set("lang", core.QueryValue(request.Lang))
	}
	if request.Pwd != nil {
		apiReq.QueryParams.Set("pwd", core.QueryValue(request.Pwd))
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

	response = &ApiV1SmartChaptersGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1SmartChaptersGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1SmartFullsummaryGetRequest struct {
	// 操作者ID。operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。
	OperatorId *string `json:"-"`
	// 操作者ID类型： 1：userid 2:openid
	OperatorIdType *string `json:"-"`
	// 录制文件ID，列表接口返回的 record_file_id。
	RecordFileId *string `json:"-"`
	// 翻译类型，默认原文展示 \"default\" 原文（不翻译） \"zh\" 简体中文 \"en\" 英文 \"ja\" 日语
	Lang *string `json:"-"`
	// 录制文件的访问密码，非必填
	Pwd  *string                 `json:"-"`
	Body *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1SmartFullsummaryGetResponse struct {
	*xhttp.ApiResponse
	Data *V1SmartFullsummaryGet200Response `json:"data,omitempty"`
}

/*
V1SmartFullsummaryGet 智能总结[/v1/smart/fullsummary - Get]

查询单个云录制的智能总结。支持OAuth2.0鉴权调用，仅支持授权用户为商业版、企业版、教育版。
当该录制文件未开启相关智能化功能或内容处于生成中状态时，结果返回错误码
*/
func (s *recordIntelligenceAPIService) V1SmartFullsummaryGet(ctx context.Context, request *ApiV1SmartFullsummaryGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1SmartFullsummaryGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/smart/fullsummary",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	if request.OperatorId == nil {
		return nil, fmt.Errorf("operator_id is required and must be specified")
	}

	if request.OperatorIdType == nil {
		return nil, fmt.Errorf("operator_id_type is required and must be specified")
	}

	if request.RecordFileId == nil {
		return nil, fmt.Errorf("record_file_id is required and must be specified")
	}

	// path 参数
	// query 参数
	if request.OperatorId != nil {
		apiReq.QueryParams.Set("operator_id", core.QueryValue(request.OperatorId))
	}
	if request.OperatorIdType != nil {
		apiReq.QueryParams.Set("operator_id_type", core.QueryValue(request.OperatorIdType))
	}
	if request.RecordFileId != nil {
		apiReq.QueryParams.Set("record_file_id", core.QueryValue(request.RecordFileId))
	}
	if request.Lang != nil {
		apiReq.QueryParams.Set("lang", core.QueryValue(request.Lang))
	}
	if request.Pwd != nil {
		apiReq.QueryParams.Set("pwd", core.QueryValue(request.Pwd))
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

	response = &ApiV1SmartFullsummaryGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1SmartFullsummaryGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1SmartSpeakersGetRequest struct {
	// 操作者ID。operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。
	OperatorId *string `json:"-"`
	// 操作者ID类型： 1：userid 2:openid
	OperatorIdType *string `json:"-"`
	// 录制文件ID，列表接口返回的 record_file_id。
	RecordFileId *string `json:"-"`
	// 页码，从1开始
	Page *string `json:"-"`
	// 返回数量，最大50
	PageSize *string `json:"-"`
	// 录制文件的访问密码，非必填
	Pwd  *string                 `json:"-"`
	Body *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1SmartSpeakersGetResponse struct {
	*xhttp.ApiResponse
	Data *V1SmartSpeakersGet200Response `json:"data,omitempty"`
}

/*
V1SmartSpeakersGet 智能发言人[/v1/smart/speakers - Get]

查询单个云录制的发言人列表。支持OAuth2.0鉴权调用，仅支持授权用户为商业版、企业版、教育版。
当该录制文件未开启相关智能化功能或内容处于生成中状态时，结果返回错误码
*/
func (s *recordIntelligenceAPIService) V1SmartSpeakersGet(ctx context.Context, request *ApiV1SmartSpeakersGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1SmartSpeakersGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/smart/speakers",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	if request.OperatorId == nil {
		return nil, fmt.Errorf("operator_id is required and must be specified")
	}

	if request.OperatorIdType == nil {
		return nil, fmt.Errorf("operator_id_type is required and must be specified")
	}

	if request.RecordFileId == nil {
		return nil, fmt.Errorf("record_file_id is required and must be specified")
	}

	if request.Page == nil {
		return nil, fmt.Errorf("page is required and must be specified")
	}

	if request.PageSize == nil {
		return nil, fmt.Errorf("page_size is required and must be specified")
	}

	// path 参数
	// query 参数
	if request.OperatorId != nil {
		apiReq.QueryParams.Set("operator_id", core.QueryValue(request.OperatorId))
	}
	if request.OperatorIdType != nil {
		apiReq.QueryParams.Set("operator_id_type", core.QueryValue(request.OperatorIdType))
	}
	if request.RecordFileId != nil {
		apiReq.QueryParams.Set("record_file_id", core.QueryValue(request.RecordFileId))
	}
	if request.Pwd != nil {
		apiReq.QueryParams.Set("pwd", core.QueryValue(request.Pwd))
	}
	if request.Page != nil {
		apiReq.QueryParams.Set("page", core.QueryValue(request.Page))
	}
	if request.PageSize != nil {
		apiReq.QueryParams.Set("page_size", core.QueryValue(request.PageSize))
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

	response = &ApiV1SmartSpeakersGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1SmartSpeakersGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1SmartTopicsGetRequest struct {
	// 操作者ID。operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。
	OperatorId *string `json:"-"`
	// 操作者ID类型： 1：userid 2:openid
	OperatorIdType *string `json:"-"`
	// 录制文件ID，列表接口返回的 record_file_id。
	RecordFileId *string `json:"-"`
	// 翻译类型，默认原文展示 \"default\" 原文（不翻译） \"zh\" 简体中文 \"en\" 英文 \"ja\" 日语
	Lang *string `json:"-"`
	// 录制文件的访问密码，非必填
	Pwd  *string                 `json:"-"`
	Body *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1SmartTopicsGetResponse struct {
	*xhttp.ApiResponse
	Data *V1SmartTopicsGet200Response `json:"data,omitempty"`
}

/*
V1SmartTopicsGet 智能话题[/v1/smart/topics - Get]

查询单个云录制的智能话题。支持OAuth2.0鉴权调用，仅支持授权用户为商业版、企业版、教育版。
当该录制文件未开启相关智能化功能或内容处于生成中状态时，结果返回错误码
*/
func (s *recordIntelligenceAPIService) V1SmartTopicsGet(ctx context.Context, request *ApiV1SmartTopicsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1SmartTopicsGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/smart/topics",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	if request.OperatorId == nil {
		return nil, fmt.Errorf("operator_id is required and must be specified")
	}

	if request.OperatorIdType == nil {
		return nil, fmt.Errorf("operator_id_type is required and must be specified")
	}

	if request.RecordFileId == nil {
		return nil, fmt.Errorf("record_file_id is required and must be specified")
	}

	// path 参数
	// query 参数
	if request.OperatorId != nil {
		apiReq.QueryParams.Set("operator_id", core.QueryValue(request.OperatorId))
	}
	if request.OperatorIdType != nil {
		apiReq.QueryParams.Set("operator_id_type", core.QueryValue(request.OperatorIdType))
	}
	if request.RecordFileId != nil {
		apiReq.QueryParams.Set("record_file_id", core.QueryValue(request.RecordFileId))
	}
	if request.Lang != nil {
		apiReq.QueryParams.Set("lang", core.QueryValue(request.Lang))
	}
	if request.Pwd != nil {
		apiReq.QueryParams.Set("pwd", core.QueryValue(request.Pwd))
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

	response = &ApiV1SmartTopicsGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1SmartTopicsGet200Response),
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
