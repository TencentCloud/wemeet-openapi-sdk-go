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
	   V1AppToolkitPost 绑定扩展应用[/v1/app/toolkit - Post]

	*/
	V1AppToolkitPost(ctx context.Context, request *ApiV1AppToolkitPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1AppToolkitPostResponse, err error)
}

type applicationGroupAPIService struct {
	config *core.Config
}

func NewService(config *core.Config) Service {
	return &applicationGroupAPIService{
		config: config,
	}
}

type ApiV1AppToolkitPostRequest struct {
	Body *V1AppToolkitPostRequest `json:"body,omitempty"`
}

type ApiV1AppToolkitPostResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1AppToolkitPost 绑定扩展应用[/v1/app/toolkit - Post]
*/
func (s *applicationGroupAPIService) V1AppToolkitPost(ctx context.Context, request *ApiV1AppToolkitPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1AppToolkitPostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/app/toolkit",
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

	response = &ApiV1AppToolkitPostResponse{
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
