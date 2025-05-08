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
	   V1DevicesGet 查询设备列表[/v1/devices - Get]

	   查询企业下的可用设备列表。<span class="colour" style="color:rgb(51, 51, 51)">目前暂不支持 OAuth2.0 鉴权访问。</span>

	*/
	V1DevicesGet(ctx context.Context, request *ApiV1DevicesGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1DevicesGetResponse, err error)

	/*
	   V1MeetingRoomsCancelRoomCallPut 取消呼叫会议室[/v1/meeting-rooms/cancel-room-call - Put]

	   **描述**：会议可以通过会议室 ID 进行取消呼叫操作。

	   * \*\*权限：\*\*同会议室呼叫权限。
	   * **调用方式**：PUT

	*/
	V1MeetingRoomsCancelRoomCallPut(ctx context.Context, request *ApiV1MeetingRoomsCancelRoomCallPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingRoomsCancelRoomCallPutResponse, err error)

	/*
	   V1MeetingRoomsGet 查询会议室（Rooms）列表[/v1/meeting-rooms - Get]

	   查询企业下的会议室列表。<span class="colour" style="color:rgb(51, 51, 51)">目前暂不支持 OAuth2.0 鉴权访问。</span>

	*/
	V1MeetingRoomsGet(ctx context.Context, request *ApiV1MeetingRoomsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingRoomsGetResponse, err error)

	/*
	   V1MeetingRoomsMeetingRoomIdActiveCodePost 生成设备激活码[/v1/meeting-rooms/{meeting_room_id}/active-code - Post]

	   生成会议室中设备激活码。<span class="colour" style="color:rgb(51, 51, 51)">目前暂不支持 OAuth2.0 鉴权访问。</span>

	*/
	V1MeetingRoomsMeetingRoomIdActiveCodePost(ctx context.Context, request *ApiV1MeetingRoomsMeetingRoomIdActiveCodePostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingRoomsMeetingRoomIdActiveCodePostResponse, err error)

	/*
	   V1MeetingRoomsMeetingRoomIdBackgroundGet 查询会议室背景[/v1/meeting-rooms/{meeting_room_id}/background - Get]

	   查询会议室的会议室背景

	*/
	V1MeetingRoomsMeetingRoomIdBackgroundGet(ctx context.Context, request *ApiV1MeetingRoomsMeetingRoomIdBackgroundGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingRoomsMeetingRoomIdBackgroundGetResponse, err error)

	/*
	   V1MeetingRoomsMeetingRoomIdBackgroundPost 设置会议室背景[/v1/meeting-rooms/{meeting_room_id}/background - Post]

	   为会议室设置会议室背景，支持为批量会议室设置。异步上传背景，需要订阅素材上传结果通知。

	*/
	V1MeetingRoomsMeetingRoomIdBackgroundPost(ctx context.Context, request *ApiV1MeetingRoomsMeetingRoomIdBackgroundPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingRoomsMeetingRoomIdBackgroundPostResponse, err error)

	/*
	   V1MeetingRoomsMeetingRoomIdGet 查询会议室（Rooms）详情[/v1/meeting-rooms/{meeting_room_id} - Get]

	   根据会议室 ID 查询该会议室详细信息。<span class="colour" style="color:rgb(51, 51, 51)">目前暂不支持 OAuth2.0 鉴权访问。</span>

	*/
	V1MeetingRoomsMeetingRoomIdGet(ctx context.Context, request *ApiV1MeetingRoomsMeetingRoomIdGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingRoomsMeetingRoomIdGetResponse, err error)

	/*
	   V1MeetingRoomsModifyPut 修改会议室信息[/v1/meeting-rooms/modify - Put]

	   **描述**：对会议室信息进行设置。

	   * \*\*权限：\*\*支持 JWT 鉴权，拥有会议室管理编辑权限。暂不支持 OAuth 2\.0鉴权。
	   * **调用方式**：PUT

	*/
	V1MeetingRoomsModifyPut(ctx context.Context, request *ApiV1MeetingRoomsModifyPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingRoomsModifyPutResponse, err error)

	/*
	   V1MeetingRoomsModifyRoomConfigInfoPost 修改会议室配置项[/v1/meeting-rooms/modify-room-config-info - Post]

	   **描述**：修改会议室各种配置项。

	   * \*\*权限：\*\*JWT 鉴权，拥有会议室管理编辑权限，暂不支持 OAuth 2\.0鉴权。
	   * **调用方式**：POST

	*/
	V1MeetingRoomsModifyRoomConfigInfoPost(ctx context.Context, request *ApiV1MeetingRoomsModifyRoomConfigInfoPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingRoomsModifyRoomConfigInfoPostResponse, err error)

	/*
	   V1MeetingRoomsMonitorDeviceControllerInfoGet 查询控制器列表[/v1/meeting-rooms-monitor/device-controller-info - Get]

	   \*\*描述：\*\*查询企业下的控制器列表，目前暂不支持 OAuth2\.0 鉴权访问。

	*/
	V1MeetingRoomsMonitorDeviceControllerInfoGet(ctx context.Context, request *ApiV1MeetingRoomsMonitorDeviceControllerInfoGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingRoomsMonitorDeviceControllerInfoGetResponse, err error)

	/*
	   V1MeetingRoomsOperatorIdMeetingsGet 查询会议室（Rooms）下的会议列表[/v1/meeting-rooms/{operator_id}/meetings - Get]

	   <span class="colour" style="color:rgb(51, 51, 51)">查询指定会议室（Rooms）下的会议列表，目前暂不支持 OAuth2.0 鉴权访问。</span>

	*/
	V1MeetingRoomsOperatorIdMeetingsGet(ctx context.Context, request *ApiV1MeetingRoomsOperatorIdMeetingsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingRoomsOperatorIdMeetingsGetResponse, err error)

	/*
	   V1MeetingRoomsRoomCallInfoPost 查询会议室应答状态[/v1/meeting-rooms/room-call-info - Post]

	   **描述**：一个会议可以查询它所呼叫的会议室对其的应答状态。

	   * \*\*权限：\*\*同会议室呼叫权限。
	   * **调用方式**：POST

	*/
	V1MeetingRoomsRoomCallInfoPost(ctx context.Context, request *ApiV1MeetingRoomsRoomCallInfoPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingRoomsRoomCallInfoPostResponse, err error)

	/*
	   V1MeetingRoomsRoomCallPut 呼叫会议室[/v1/meeting-rooms/room-call - Put]

	   **描述**：会议可以通过会议室 ID 呼叫会议室邀请其入会。

	   * \*\*权限：\*\*支持 JWT 鉴权，会议创建者所在企业的管理员和会议参会者可呼叫与自己同企业下的会议室入会，若使用会议室呼叫地址，需主持人或联席主持人身份，暂不支持 OAuth 2\.0鉴权。
	   * **调用方式**：PUT

	*/
	V1MeetingRoomsRoomCallPut(ctx context.Context, request *ApiV1MeetingRoomsRoomCallPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingRoomsRoomCallPutResponse, err error)

	/*
	   V1MeetingRoomsRoomConfigInfoPost 查询会议室配置项[/v1/meeting-rooms/room-config-info - Post]

	   **描述**：查询会议室的配置项。

	   * \*\*权限：\*\*JWT 鉴权，拥有会议室查看权限，暂不支持 OAuth 2\.0鉴权。
	   * **调用方式**：POST

	*/
	V1MeetingRoomsRoomConfigInfoPost(ctx context.Context, request *ApiV1MeetingRoomsRoomConfigInfoPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingRoomsRoomConfigInfoPostResponse, err error)

	/*
	   V1MeetingRoomsScreencastCodeRoomsInfoGet 通过投屏码查询会议室信息[/v1/meeting-rooms/{screencast_code}/rooms-info - Get]

	*/
	V1MeetingRoomsScreencastCodeRoomsInfoGet(ctx context.Context, request *ApiV1MeetingRoomsScreencastCodeRoomsInfoGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingRoomsScreencastCodeRoomsInfoGetResponse, err error)

	/*
	   V1MeetingsMeetingIdBookRoomsPost 预定会议室（Rooms）[/v1/meetings/{meeting_id}/book-rooms - Post]

	   对成功预定的会议添加会议室，支持多个会议室预定。说明：会议室预定对会议时长有硬性要求，会议时长不得小于15分钟且不得大于24小时；不支持周期性会议。）

	   * 通过会议 ID 预定会议室。
	   * 目前暂不支持 OAuth2.0 鉴权访问。

	*/
	V1MeetingsMeetingIdBookRoomsPost(ctx context.Context, request *ApiV1MeetingsMeetingIdBookRoomsPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdBookRoomsPostResponse, err error)

	/*
	   V1MeetingsMeetingIdReleaseRoomsPost 释放会议室（Rooms）[/v1/meetings/{meeting_id}/release-rooms - Post]

	   通过会议 ID 释放会议室，支持多个会议室释放。<span class="colour" style="color:rgb(51, 51, 51)">目前暂不支持 OAuth2.0 鉴权访问。</span>

	*/
	V1MeetingsMeetingIdReleaseRoomsPost(ctx context.Context, request *ApiV1MeetingsMeetingIdReleaseRoomsPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdReleaseRoomsPostResponse, err error)

	/*
	   V1RoomsInventoryAccountStatisticsGet 查询账号类型资源使用统计[/v1/rooms-inventory/account-statistics - Get]

	   查询 Rooms 账号资源使用情况，暂不支持 OAuth 2.0鉴权访问。

	*/
	V1RoomsInventoryAccountStatisticsGet(ctx context.Context, request *ApiV1RoomsInventoryAccountStatisticsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1RoomsInventoryAccountStatisticsGetResponse, err error)

	/*
	   V1RoomsInventoryGet 查询账户下 Rooms 资源[/v1/rooms-inventory - Get]

	   查询企业购买的 Rooms 资源。<span class="colour" style="color: rgb(51, 51, 51);">目前暂不支持 OAuth2.0 鉴权访问。</span>

	*/
	V1RoomsInventoryGet(ctx context.Context, request *ApiV1RoomsInventoryGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1RoomsInventoryGetResponse, err error)
}

type meetingRoomAPIService struct {
	config *core.Config
}

func NewService(config *core.Config) Service {
	return &meetingRoomAPIService{
		config: config,
	}
}

type ApiV1DevicesGetRequest struct {
	// 页码，从1开始，默认1。
	Page *string `json:"-"`
	// 分页大小，从1开始，最大50，默认20。
	PageSize *string `json:"-"`
	// 会议室名称
	MeetingRoomName *string                 `json:"-"`
	Body            *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1DevicesGetResponse struct {
	*xhttp.ApiResponse
	Data *V1DevicesGet200Response `json:"data,omitempty"`
}

/*
V1DevicesGet 查询设备列表[/v1/devices - Get]

查询企业下的可用设备列表。<span class="colour" style="color:rgb(51, 51, 51)">目前暂不支持 OAuth2.0 鉴权访问。</span>
*/
func (s *meetingRoomAPIService) V1DevicesGet(ctx context.Context, request *ApiV1DevicesGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1DevicesGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/devices",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	// path 参数
	// query 参数
	if request.Page != nil {
		apiReq.QueryParams.Set("page", core.QueryValue(request.Page))
	}
	if request.PageSize != nil {
		apiReq.QueryParams.Set("page_size", core.QueryValue(request.PageSize))
	}
	if request.MeetingRoomName != nil {
		apiReq.QueryParams.Set("meeting_room_name", core.QueryValue(request.MeetingRoomName))
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

	response = &ApiV1DevicesGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1DevicesGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingRoomsCancelRoomCallPutRequest struct {
	Body *V1MeetingRoomsCancelRoomCallPutRequest `json:"body,omitempty"`
}

type ApiV1MeetingRoomsCancelRoomCallPutResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1MeetingRoomsCancelRoomCallPut 取消呼叫会议室[/v1/meeting-rooms/cancel-room-call - Put]

**描述**：会议可以通过会议室 ID 进行取消呼叫操作。

* \*\*权限：\*\*同会议室呼叫权限。
* **调用方式**：PUT
*/
func (s *meetingRoomAPIService) V1MeetingRoomsCancelRoomCallPut(ctx context.Context, request *ApiV1MeetingRoomsCancelRoomCallPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingRoomsCancelRoomCallPutResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meeting-rooms/cancel-room-call",
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

	response = &ApiV1MeetingRoomsCancelRoomCallPutResponse{
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

type ApiV1MeetingRoomsGetRequest struct {
	// 页码
	Page *string `json:"-"`
	// 分页大小
	PageSize *string `json:"-"`
	// 会议室名称
	MeetingRoomName *string                 `json:"-"`
	Body            *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1MeetingRoomsGetResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingRoomsGet200Response `json:"data,omitempty"`
}

/*
V1MeetingRoomsGet 查询会议室（Rooms）列表[/v1/meeting-rooms - Get]

查询企业下的会议室列表。<span class="colour" style="color:rgb(51, 51, 51)">目前暂不支持 OAuth2.0 鉴权访问。</span>
*/
func (s *meetingRoomAPIService) V1MeetingRoomsGet(ctx context.Context, request *ApiV1MeetingRoomsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingRoomsGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meeting-rooms",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	// path 参数
	// query 参数
	if request.Page != nil {
		apiReq.QueryParams.Set("page", core.QueryValue(request.Page))
	}
	if request.PageSize != nil {
		apiReq.QueryParams.Set("page_size", core.QueryValue(request.PageSize))
	}
	if request.MeetingRoomName != nil {
		apiReq.QueryParams.Set("meeting_room_name", core.QueryValue(request.MeetingRoomName))
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

	response = &ApiV1MeetingRoomsGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingRoomsGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingRoomsMeetingRoomIdActiveCodePostRequest struct {
	// 会议室id
	MeetingRoomId string                  `json:"-"`
	Body          *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1MeetingRoomsMeetingRoomIdActiveCodePostResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingRoomsMeetingRoomIdActiveCodePost200Response `json:"data,omitempty"`
}

/*
V1MeetingRoomsMeetingRoomIdActiveCodePost 生成设备激活码[/v1/meeting-rooms/{meeting_room_id}/active-code - Post]

生成会议室中设备激活码。<span class="colour" style="color:rgb(51, 51, 51)">目前暂不支持 OAuth2.0 鉴权访问。</span>
*/
func (s *meetingRoomAPIService) V1MeetingRoomsMeetingRoomIdActiveCodePost(ctx context.Context, request *ApiV1MeetingRoomsMeetingRoomIdActiveCodePostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingRoomsMeetingRoomIdActiveCodePostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meeting-rooms/{meeting_room_id}/active-code",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	// path 参数
	apiReq.PathParams.Set("meeting_room_id", core.PathValue(request.MeetingRoomId))
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

	response = &ApiV1MeetingRoomsMeetingRoomIdActiveCodePostResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingRoomsMeetingRoomIdActiveCodePost200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingRoomsMeetingRoomIdBackgroundGetRequest struct {
	MeetingRoomId string  `json:"-"`
	OperatorId    *string `json:"-"`
	// 1:userid
	OperatorIdType *string `json:"-"`
}

type ApiV1MeetingRoomsMeetingRoomIdBackgroundGetResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingRoomsMeetingRoomIdBackgroundGet200Response `json:"data,omitempty"`
}

/*
V1MeetingRoomsMeetingRoomIdBackgroundGet 查询会议室背景[/v1/meeting-rooms/{meeting_room_id}/background - Get]

查询会议室的会议室背景
*/
func (s *meetingRoomAPIService) V1MeetingRoomsMeetingRoomIdBackgroundGet(ctx context.Context, request *ApiV1MeetingRoomsMeetingRoomIdBackgroundGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingRoomsMeetingRoomIdBackgroundGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meeting-rooms/{meeting_room_id}/background",
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	if request.OperatorId == nil {
		return nil, fmt.Errorf("operator_id is required and must be specified")
	}

	if request.OperatorIdType == nil {
		return nil, fmt.Errorf("operator_id_type is required and must be specified")
	}

	// path 参数
	apiReq.PathParams.Set("meeting_room_id", core.PathValue(request.MeetingRoomId))
	// query 参数
	if request.OperatorId != nil {
		apiReq.QueryParams.Set("operator_id", core.QueryValue(request.OperatorId))
	}
	if request.OperatorIdType != nil {
		apiReq.QueryParams.Set("operator_id_type", core.QueryValue(request.OperatorIdType))
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

	response = &ApiV1MeetingRoomsMeetingRoomIdBackgroundGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingRoomsMeetingRoomIdBackgroundGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingRoomsMeetingRoomIdBackgroundPostRequest struct {
	MeetingRoomId string                                            `json:"-"`
	Body          *V1MeetingRoomsMeetingRoomIdBackgroundPostRequest `json:"body,omitempty"`
}

type ApiV1MeetingRoomsMeetingRoomIdBackgroundPostResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingRoomsMeetingRoomIdBackgroundPost200Response `json:"data,omitempty"`
}

/*
V1MeetingRoomsMeetingRoomIdBackgroundPost 设置会议室背景[/v1/meeting-rooms/{meeting_room_id}/background - Post]

为会议室设置会议室背景，支持为批量会议室设置。异步上传背景，需要订阅素材上传结果通知。
*/
func (s *meetingRoomAPIService) V1MeetingRoomsMeetingRoomIdBackgroundPost(ctx context.Context, request *ApiV1MeetingRoomsMeetingRoomIdBackgroundPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingRoomsMeetingRoomIdBackgroundPostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meeting-rooms/{meeting_room_id}/background",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	// path 参数
	apiReq.PathParams.Set("meeting_room_id", core.PathValue(request.MeetingRoomId))
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

	response = &ApiV1MeetingRoomsMeetingRoomIdBackgroundPostResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingRoomsMeetingRoomIdBackgroundPost200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingRoomsMeetingRoomIdGetRequest struct {
	// 会议室ID
	MeetingRoomId string                  `json:"-"`
	Body          *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1MeetingRoomsMeetingRoomIdGetResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingRoomsMeetingRoomIdGet200Response `json:"data,omitempty"`
}

/*
V1MeetingRoomsMeetingRoomIdGet 查询会议室（Rooms）详情[/v1/meeting-rooms/{meeting_room_id} - Get]

根据会议室 ID 查询该会议室详细信息。<span class="colour" style="color:rgb(51, 51, 51)">目前暂不支持 OAuth2.0 鉴权访问。</span>
*/
func (s *meetingRoomAPIService) V1MeetingRoomsMeetingRoomIdGet(ctx context.Context, request *ApiV1MeetingRoomsMeetingRoomIdGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingRoomsMeetingRoomIdGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meeting-rooms/{meeting_room_id}",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	// path 参数
	apiReq.PathParams.Set("meeting_room_id", core.PathValue(request.MeetingRoomId))
	// query 参数
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

	response = &ApiV1MeetingRoomsMeetingRoomIdGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingRoomsMeetingRoomIdGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingRoomsModifyPutRequest struct {
	Body *V1MeetingRoomsModifyPutRequest `json:"body,omitempty"`
}

type ApiV1MeetingRoomsModifyPutResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1MeetingRoomsModifyPut 修改会议室信息[/v1/meeting-rooms/modify - Put]

**描述**：对会议室信息进行设置。

* \*\*权限：\*\*支持 JWT 鉴权，拥有会议室管理编辑权限。暂不支持 OAuth 2\.0鉴权。
* **调用方式**：PUT
*/
func (s *meetingRoomAPIService) V1MeetingRoomsModifyPut(ctx context.Context, request *ApiV1MeetingRoomsModifyPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingRoomsModifyPutResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meeting-rooms/modify",
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

	response = &ApiV1MeetingRoomsModifyPutResponse{
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

type ApiV1MeetingRoomsModifyRoomConfigInfoPostRequest struct {
	Body *V1MeetingRoomsModifyRoomConfigInfoPostRequest `json:"body,omitempty"`
}

type ApiV1MeetingRoomsModifyRoomConfigInfoPostResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1MeetingRoomsModifyRoomConfigInfoPost 修改会议室配置项[/v1/meeting-rooms/modify-room-config-info - Post]

**描述**：修改会议室各种配置项。

* \*\*权限：\*\*JWT 鉴权，拥有会议室管理编辑权限，暂不支持 OAuth 2\.0鉴权。
* **调用方式**：POST
*/
func (s *meetingRoomAPIService) V1MeetingRoomsModifyRoomConfigInfoPost(ctx context.Context, request *ApiV1MeetingRoomsModifyRoomConfigInfoPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingRoomsModifyRoomConfigInfoPostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meeting-rooms/modify-room-config-info",
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

	response = &ApiV1MeetingRoomsModifyRoomConfigInfoPostResponse{
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

type ApiV1MeetingRoomsMonitorDeviceControllerInfoGetRequest struct {
	// 需要查询的设备名称（支持模糊匹配查找），如需获取全量列表，则不需要传入。
	ControllerName *string `json:"-"`
	// 页码，从1开始，默认1。
	Page *string `json:"-"`
	// 分页大小，从1开始，最大50，默认20。
	PageSize *string `json:"-"`
}

type ApiV1MeetingRoomsMonitorDeviceControllerInfoGetResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingRoomsMonitorDeviceControllerInfoGet200Response `json:"data,omitempty"`
}

/*
V1MeetingRoomsMonitorDeviceControllerInfoGet 查询控制器列表[/v1/meeting-rooms-monitor/device-controller-info - Get]

\*\*描述：\*\*查询企业下的控制器列表，目前暂不支持 OAuth2\.0 鉴权访问。
*/
func (s *meetingRoomAPIService) V1MeetingRoomsMonitorDeviceControllerInfoGet(ctx context.Context, request *ApiV1MeetingRoomsMonitorDeviceControllerInfoGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingRoomsMonitorDeviceControllerInfoGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meeting-rooms-monitor/device-controller-info",
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	// path 参数
	// query 参数
	if request.ControllerName != nil {
		apiReq.QueryParams.Set("controller_name", core.QueryValue(request.ControllerName))
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

	response = &ApiV1MeetingRoomsMonitorDeviceControllerInfoGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingRoomsMonitorDeviceControllerInfoGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingRoomsOperatorIdMeetingsGetRequest struct {
	OperatorId string `json:"-"`
	// 操作者 ID 的类型： 3. rooms 设备 rooms_id 5. 会议室ID meeting_room_id
	OperatorIdType *string `json:"-"`
	// 用户的终端设备类型： 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch Mac 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch Iphone
	Instanceid *string `json:"-"`
	// 目标查询 roomsid。
	TargetRoomsId *string `json:"-"`
	// 目标查询 roomsid 的类型： 3：rooms 设备 rooms_id 5：会议室 ID meeting_room_id
	TargetRoomsIdType *string `json:"-"`
	// Unix 时间戳。查询起始时间，时间区间不超过90天。
	StartTime *string `json:"-"`
	// Unix 时间戳。查询结束时间，时间区间不超过90天。
	EndTime *string `json:"-"`
	// 当前页，页码起始值为1，默认为1。
	Page *string `json:"-"`
	// 分页大小，默认20条，最大20条。
	PageSize *string                 `json:"-"`
	Body     *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1MeetingRoomsOperatorIdMeetingsGetResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingRoomsOperatorIdMeetingsGet200Response `json:"data,omitempty"`
}

/*
V1MeetingRoomsOperatorIdMeetingsGet 查询会议室（Rooms）下的会议列表[/v1/meeting-rooms/{operator_id}/meetings - Get]

<span class="colour" style="color:rgb(51, 51, 51)">查询指定会议室（Rooms）下的会议列表，目前暂不支持 OAuth2.0 鉴权访问。</span>
*/
func (s *meetingRoomAPIService) V1MeetingRoomsOperatorIdMeetingsGet(ctx context.Context, request *ApiV1MeetingRoomsOperatorIdMeetingsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingRoomsOperatorIdMeetingsGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meeting-rooms/{operator_id}/meetings",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	if request.OperatorIdType == nil {
		return nil, fmt.Errorf("operator_id_type is required and must be specified")
	}

	if request.Instanceid == nil {
		return nil, fmt.Errorf("instanceid is required and must be specified")
	}

	if request.TargetRoomsId == nil {
		return nil, fmt.Errorf("target_rooms_id is required and must be specified")
	}

	if request.TargetRoomsIdType == nil {
		return nil, fmt.Errorf("target_rooms_id_type is required and must be specified")
	}

	// path 参数
	apiReq.PathParams.Set("operator_id", core.PathValue(request.OperatorId))
	// query 参数
	if request.OperatorIdType != nil {
		apiReq.QueryParams.Set("operator_id_type", core.QueryValue(request.OperatorIdType))
	}
	if request.Instanceid != nil {
		apiReq.QueryParams.Set("instanceid", core.QueryValue(request.Instanceid))
	}
	if request.StartTime != nil {
		apiReq.QueryParams.Set("start_time", core.QueryValue(request.StartTime))
	}
	if request.EndTime != nil {
		apiReq.QueryParams.Set("end_time", core.QueryValue(request.EndTime))
	}
	if request.Page != nil {
		apiReq.QueryParams.Set("page", core.QueryValue(request.Page))
	}
	if request.PageSize != nil {
		apiReq.QueryParams.Set("page_size", core.QueryValue(request.PageSize))
	}
	if request.TargetRoomsId != nil {
		apiReq.QueryParams.Set("target_rooms_id", core.QueryValue(request.TargetRoomsId))
	}
	if request.TargetRoomsIdType != nil {
		apiReq.QueryParams.Set("target_rooms_id_type", core.QueryValue(request.TargetRoomsIdType))
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

	response = &ApiV1MeetingRoomsOperatorIdMeetingsGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingRoomsOperatorIdMeetingsGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingRoomsRoomCallInfoPostRequest struct {
	Body *V1MeetingRoomsRoomCallInfoPostRequest `json:"body,omitempty"`
}

type ApiV1MeetingRoomsRoomCallInfoPostResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingRoomsRoomCallInfoPost200Response `json:"data,omitempty"`
}

/*
V1MeetingRoomsRoomCallInfoPost 查询会议室应答状态[/v1/meeting-rooms/room-call-info - Post]

**描述**：一个会议可以查询它所呼叫的会议室对其的应答状态。

* \*\*权限：\*\*同会议室呼叫权限。
* **调用方式**：POST
*/
func (s *meetingRoomAPIService) V1MeetingRoomsRoomCallInfoPost(ctx context.Context, request *ApiV1MeetingRoomsRoomCallInfoPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingRoomsRoomCallInfoPostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meeting-rooms/room-call-info",
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

	response = &ApiV1MeetingRoomsRoomCallInfoPostResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingRoomsRoomCallInfoPost200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingRoomsRoomCallPutRequest struct {
	Body *V1MeetingRoomsRoomCallPutRequest `json:"body,omitempty"`
}

type ApiV1MeetingRoomsRoomCallPutResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingRoomsRoomCallPut200Response `json:"data,omitempty"`
}

/*
V1MeetingRoomsRoomCallPut 呼叫会议室[/v1/meeting-rooms/room-call - Put]

**描述**：会议可以通过会议室 ID 呼叫会议室邀请其入会。

* \*\*权限：\*\*支持 JWT 鉴权，会议创建者所在企业的管理员和会议参会者可呼叫与自己同企业下的会议室入会，若使用会议室呼叫地址，需主持人或联席主持人身份，暂不支持 OAuth 2\.0鉴权。
* **调用方式**：PUT
*/
func (s *meetingRoomAPIService) V1MeetingRoomsRoomCallPut(ctx context.Context, request *ApiV1MeetingRoomsRoomCallPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingRoomsRoomCallPutResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meeting-rooms/room-call",
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

	response = &ApiV1MeetingRoomsRoomCallPutResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingRoomsRoomCallPut200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingRoomsRoomConfigInfoPostRequest struct {
	Body *V1MeetingRoomsRoomConfigInfoPostRequest `json:"body,omitempty"`
}

type ApiV1MeetingRoomsRoomConfigInfoPostResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingRoomsRoomConfigInfoPost200Response `json:"data,omitempty"`
}

/*
V1MeetingRoomsRoomConfigInfoPost 查询会议室配置项[/v1/meeting-rooms/room-config-info - Post]

**描述**：查询会议室的配置项。

* \*\*权限：\*\*JWT 鉴权，拥有会议室查看权限，暂不支持 OAuth 2\.0鉴权。
* **调用方式**：POST
*/
func (s *meetingRoomAPIService) V1MeetingRoomsRoomConfigInfoPost(ctx context.Context, request *ApiV1MeetingRoomsRoomConfigInfoPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingRoomsRoomConfigInfoPostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meeting-rooms/room-config-info",
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

	response = &ApiV1MeetingRoomsRoomConfigInfoPostResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingRoomsRoomConfigInfoPost200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingRoomsScreencastCodeRoomsInfoGetRequest struct {
	// 投屏码
	ScreencastCode string                  `json:"-"`
	Body           *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1MeetingRoomsScreencastCodeRoomsInfoGetResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingRoomsScreencastCodeRoomsInfoGet200Response `json:"data,omitempty"`
}

/*
V1MeetingRoomsScreencastCodeRoomsInfoGet 通过投屏码查询会议室信息[/v1/meeting-rooms/{screencast_code}/rooms-info - Get]
*/
func (s *meetingRoomAPIService) V1MeetingRoomsScreencastCodeRoomsInfoGet(ctx context.Context, request *ApiV1MeetingRoomsScreencastCodeRoomsInfoGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingRoomsScreencastCodeRoomsInfoGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meeting-rooms/{screencast_code}/rooms-info",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	// path 参数
	apiReq.PathParams.Set("screencast_code", core.PathValue(request.ScreencastCode))
	// query 参数
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

	response = &ApiV1MeetingRoomsScreencastCodeRoomsInfoGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingRoomsScreencastCodeRoomsInfoGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingsMeetingIdBookRoomsPostRequest struct {
	// 会议唯一id
	MeetingId string                                   `json:"-"`
	Body      *V1MeetingsMeetingIdBookRoomsPostRequest `json:"body,omitempty"`
}

type ApiV1MeetingsMeetingIdBookRoomsPostResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingsMeetingIdBookRoomsPost200Response `json:"data,omitempty"`
}

/*
V1MeetingsMeetingIdBookRoomsPost 预定会议室（Rooms）[/v1/meetings/{meeting_id}/book-rooms - Post]

对成功预定的会议添加会议室，支持多个会议室预定。说明：会议室预定对会议时长有硬性要求，会议时长不得小于15分钟且不得大于24小时；不支持周期性会议。）

* 通过会议 ID 预定会议室。
* 目前暂不支持 OAuth2.0 鉴权访问。
*/
func (s *meetingRoomAPIService) V1MeetingsMeetingIdBookRoomsPost(ctx context.Context, request *ApiV1MeetingsMeetingIdBookRoomsPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdBookRoomsPostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings/{meeting_id}/book-rooms",
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

	response = &ApiV1MeetingsMeetingIdBookRoomsPostResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingsMeetingIdBookRoomsPost200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingsMeetingIdReleaseRoomsPostRequest struct {
	// 会议唯一id
	MeetingId string                  `json:"-"`
	Body      *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1MeetingsMeetingIdReleaseRoomsPostResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1MeetingsMeetingIdReleaseRoomsPost 释放会议室（Rooms）[/v1/meetings/{meeting_id}/release-rooms - Post]

通过会议 ID 释放会议室，支持多个会议室释放。<span class="colour" style="color:rgb(51, 51, 51)">目前暂不支持 OAuth2.0 鉴权访问。</span>
*/
func (s *meetingRoomAPIService) V1MeetingsMeetingIdReleaseRoomsPost(ctx context.Context, request *ApiV1MeetingsMeetingIdReleaseRoomsPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdReleaseRoomsPostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings/{meeting_id}/release-rooms",
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

	response = &ApiV1MeetingsMeetingIdReleaseRoomsPostResponse{
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

type ApiV1RoomsInventoryAccountStatisticsGetRequest struct {
	Body *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1RoomsInventoryAccountStatisticsGetResponse struct {
	*xhttp.ApiResponse
	Data *V1RoomsInventoryAccountStatisticsGet200Response `json:"data,omitempty"`
}

/*
V1RoomsInventoryAccountStatisticsGet 查询账号类型资源使用统计[/v1/rooms-inventory/account-statistics - Get]

查询 Rooms 账号资源使用情况，暂不支持 OAuth 2.0鉴权访问。
*/
func (s *meetingRoomAPIService) V1RoomsInventoryAccountStatisticsGet(ctx context.Context, request *ApiV1RoomsInventoryAccountStatisticsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1RoomsInventoryAccountStatisticsGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/rooms-inventory/account-statistics",
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

	response = &ApiV1RoomsInventoryAccountStatisticsGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1RoomsInventoryAccountStatisticsGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1RoomsInventoryGetRequest struct {
}

type ApiV1RoomsInventoryGetResponse struct {
	*xhttp.ApiResponse
	Data *V1RoomsInventoryGet200Response `json:"data,omitempty"`
}

/*
V1RoomsInventoryGet 查询账户下 Rooms 资源[/v1/rooms-inventory - Get]

查询企业购买的 Rooms 资源。<span class="colour" style="color: rgb(51, 51, 51);">目前暂不支持 OAuth2.0 鉴权访问。</span>
*/
func (s *meetingRoomAPIService) V1RoomsInventoryGet(ctx context.Context, request *ApiV1RoomsInventoryGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1RoomsInventoryGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/rooms-inventory",
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

	response = &ApiV1RoomsInventoryGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1RoomsInventoryGet200Response),
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
