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
	   V1MeetingsMeetingIdDismissPost 结束会议[/v1/meetings/{meeting_id}/dismiss - Post]

	   **描述**：结束一个进行中的会议。

	   * 只有会议创建者、主持人、联席主持人可以结束会议，且该会议是一个有效的进行中会议。
	   * 结束周期性会议需要传入主会议 ID。
	   * 企业 secret 鉴权用户可结束任何该企业该用户创建的进行中的会议，OAuth 2.0鉴权用户只能结束通过 OAuth 2.0鉴权创建的进行中的会议。
	   * 当您想实时监测会议结束状况时，您可以通过订阅 [会议结束](https://cloud.tencent.com/document/product/1095/51619) 的事件，接收事件通知。
	   * 此接口暂不支持 MRA 设备作为被操作者的情况。
	   * 结束会议不会释放 Rooms，如需释放请调用 [释放会议室（Rooms）](https://cloud.tencent.com/document/product/1095/55314) 接口

	*/
	V1MeetingsMeetingIdDismissPost(ctx context.Context, request *ApiV1MeetingsMeetingIdDismissPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdDismissPostResponse, err error)

	/*
	   V1RealControlMeetingsMeetingIdAsrPut 开启或关闭实时转写[/v1/real-control/meetings/{meeting_id}/asr - Put]

	   以创建者的身份开启/关闭会中实时转写，调用时需要会议处于进行中的状态；

	*/
	V1RealControlMeetingsMeetingIdAsrPut(ctx context.Context, request *ApiV1RealControlMeetingsMeetingIdAsrPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1RealControlMeetingsMeetingIdAsrPutResponse, err error)

	/*
	   V1RealControlMeetingsMeetingIdCohostsPut 设置联席主持人[/v1/real-control/meetings/{meeting_id}/cohosts - Put]

	   **描述**：设置或撤销会议中的参会者联席主持人身份，目前暂不支持 MRA 设备作为被操作者的情况。企业 secret 鉴权用户可管理任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能管理通过 OAuth2.0 鉴权创建的有效会议。

	   *

	   > 说明
	   >
	   >
	   > * 操作者必须是会议的主持人才可以设置。
	   > * 调用该接口需要权限项 ： MANAGE_MEETING 查看和管理您的会议。

	*/
	V1RealControlMeetingsMeetingIdCohostsPut(ctx context.Context, request *ApiV1RealControlMeetingsMeetingIdCohostsPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1RealControlMeetingsMeetingIdCohostsPutResponse, err error)

	/*
	   V1RealControlMeetingsMeetingIdDocPut 文档上传权限设置[/v1/real-control/meetings/{meeting_id}/doc - Put]

	   **描述**：设置会议中文档上传权限，目前暂不支持 MRA 设备作为被操作者的情况，企业 secret 鉴权用户可管理任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能管理通过 OAuth2.0 鉴权创建的有效会议。

	   *

	   > 说明
	   > 1：操作者必须是会议的主持人或者联席主持人才可以设置。
	   > 2：调用该接口需要权限项：MANAGE_MEETING 查看和管理您的会议

	*/
	V1RealControlMeetingsMeetingIdDocPut(ctx context.Context, request *ApiV1RealControlMeetingsMeetingIdDocPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1RealControlMeetingsMeetingIdDocPutResponse, err error)

	/*
	   V1RealControlMeetingsMeetingIdKickoutPut 移出用户[/v1/real-control/meetings/{meeting_id}/kickout - Put]

	   **描述**：

	   * 将会议中用户移出会议，操作者必须是会议的主持人或者联席主持人才可以设置。
	   * 支持对云会议已入会成员和 Webinar 观众移出。
	   * 支持 MRA 设备和 PSTN 设备为被操作者时的移出用户操作。
	   * 企业 secret 鉴权用户可管理任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能管理通过 OAuth2.0 鉴权创建的有效会议。

	   说明

	   *

	   > 1：操作者必须是会议的主持人或者联席主持人才可以设置。
	   > 2：调用该接口需要权限项：MANAGE_MEETING 查看和管理您的会议。

	*/
	V1RealControlMeetingsMeetingIdKickoutPut(ctx context.Context, request *ApiV1RealControlMeetingsMeetingIdKickoutPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1RealControlMeetingsMeetingIdKickoutPutResponse, err error)

	/*
	   V1RealControlMeetingsMeetingIdMutesPut 静音用户[/v1/real-control/meetings/{meeting_id}/mutes - Put]

	   **描述**：

	   * 会议中用户静音操作，操作者必须是会议的主持人或者联席主持人才可以设置。
	   * 支持对云会议已入会成员和 Webinar 观众静音。
	   * 支持 MRA 设备和 PSTN 设备作为被操作者时的静音操作。
	   * 企业 secret 鉴权用户可管理任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能管理通过 OAuth2.0 鉴权创建的有效会议。

	   说明
	   1：操作者必须是会议的主持人或者联席主持人才可以设置。
	   2：调用该接口需要权限项：MANAGE_MEETING 查看和管理您的会议。

	*/
	V1RealControlMeetingsMeetingIdMutesPut(ctx context.Context, request *ApiV1RealControlMeetingsMeetingIdMutesPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1RealControlMeetingsMeetingIdMutesPutResponse, err error)

	/*
	   V1RealControlMeetingsMeetingIdNamesPut 更改会中成员昵称[/v1/real-control/meetings/{meeting_id}/names - Put]

	   **描述：**

	   * 会中修改参会者昵称，支持主持人和联席主持人对会中成员进行改名。
	   * 此接口支持对云会议已入会成员和 Webinar 观众进行改名。
	   * 操作者必须为已在会中的主持人与联席主持人。
	   * 支持 MRA 设备和 PSTN 设备作为被操作者时的改名操作。
	   * 企业 secret 鉴权用户可管理任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能管理通过 OAuth2.0 鉴权创建的有效会议。

	   说明
	   1：操作者必须是会议的主持人或者联席主持人才可以设置。
	   2：调用该接口需要权限项：MANAGE_MEETING 查看和管理您的会议。

	*/
	V1RealControlMeetingsMeetingIdNamesPut(ctx context.Context, request *ApiV1RealControlMeetingsMeetingIdNamesPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1RealControlMeetingsMeetingIdNamesPutResponse, err error)

	/*
	   V1RealControlMeetingsMeetingIdScreenSharedPut 关闭用户屏幕共享[/v1/real-control/meetings/{meeting_id}/screen-shared - Put]

	   **描述**：关闭会议中用户屏幕共享权限，目前暂不支持 MRA 设备作为被操作者的情况。企业 secret 鉴权用户可管理任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能管理通过 OAuth2.0 鉴权创建的有效会议。

	   *

	   > 说明
	   > 1：操作者必须是会议的主持人或者联席主持人才可以设置。
	   > 2：调用该接口需要权限项：MANAGE_MEETING 查看和管理您的会议

	*/
	V1RealControlMeetingsMeetingIdScreenSharedPut(ctx context.Context, request *ApiV1RealControlMeetingsMeetingIdScreenSharedPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1RealControlMeetingsMeetingIdScreenSharedPutResponse, err error)

	/*
	   V1RealControlMeetingsMeetingIdStatusPut 会中状态设置[/v1/real-control/meetings/{meeting_id}/status - Put]

	   **描述**：设置会议中的会议属性，例如：全体静音、是否允许参会者聊天设置、锁定会议、隐藏会议号和密码、是否开启等候室等，具体设置内容可以查询接口协议，目前暂不支持 MRA 设备作为被操作者的情况。企业 secret 鉴权用户可管理任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能管理通过 OAuth2.0 鉴权创建的有效会议。

	   *

	   > 说明
	   > 1：操作者必须是会议的主持人或者联席主持人才可以设置。
	   > 2：调用该接口需要权限项 ：MANAGE_MEETING 查看和管理您的会议。

	*/
	V1RealControlMeetingsMeetingIdStatusPut(ctx context.Context, request *ApiV1RealControlMeetingsMeetingIdStatusPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1RealControlMeetingsMeetingIdStatusPutResponse, err error)

	/*
	   V1RealControlMeetingsMeetingIdVideoPut 关闭或开启用户视频[/v1/real-control/meetings/{meeting_id}/video - Put]

	   关闭指定用户视频，支持关闭或开启 MRA 设备的视频。
	   企业 secret 鉴权用户可管理任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能管理通过 OAuth2.0 鉴权创建的有效会议。
	   > 说明
	   > 1：操作者必须是会议的主持人或者联席主持人才可以设置。
	   > 2：使用 OAuth 2.0 鉴权方式时，调用该接口需要权限项：MANAGE_MEETING 查看和管理您的会议。

	*/
	V1RealControlMeetingsMeetingIdVideoPut(ctx context.Context, request *ApiV1RealControlMeetingsMeetingIdVideoPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1RealControlMeetingsMeetingIdVideoPutResponse, err error)

	/*
	   V1RealControlMeetingsMeetingIdWaitingRoomPut 用户等候室设置[/v1/real-control/meetings/{meeting_id}/waiting-room - Put]

	   会议等候室设置，允许主持人将等候室成员移入会议、主持人将会议成员移入等候室、主持人将等候室成员移出等候室等操作，目前暂不支持 MRA 设备作为被操作者的情况。企业 secret 鉴权用户可管理任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能管理通过 OAuth2.0 鉴权创建的有效会议。
	   > 说明
	   > 1：操作者必须是会议的主持人或者联席主持人才可以设置。
	   > 2：调用该接口需要权限项：MANAGE_MEETING 查看和管理您的会议

	*/
	V1RealControlMeetingsMeetingIdWaitingRoomPut(ctx context.Context, request *ApiV1RealControlMeetingsMeetingIdWaitingRoomPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1RealControlMeetingsMeetingIdWaitingRoomPutResponse, err error)
}

type meetingControlAPIService struct {
	config *core.Config
}

func NewService(config *core.Config) Service {
	return &meetingControlAPIService{
		config: config,
	}
}

type ApiV1MeetingsMeetingIdDismissPostRequest struct {
	MeetingId string                                 `json:"-"`
	Body      *V1MeetingsMeetingIdDismissPostRequest `json:"body,omitempty"`
}

type ApiV1MeetingsMeetingIdDismissPostResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1MeetingsMeetingIdDismissPost 结束会议[/v1/meetings/{meeting_id}/dismiss - Post]

**描述**：结束一个进行中的会议。

* 只有会议创建者、主持人、联席主持人可以结束会议，且该会议是一个有效的进行中会议。
* 结束周期性会议需要传入主会议 ID。
* 企业 secret 鉴权用户可结束任何该企业该用户创建的进行中的会议，OAuth 2.0鉴权用户只能结束通过 OAuth 2.0鉴权创建的进行中的会议。
* 当您想实时监测会议结束状况时，您可以通过订阅 [会议结束](https://cloud.tencent.com/document/product/1095/51619) 的事件，接收事件通知。
* 此接口暂不支持 MRA 设备作为被操作者的情况。
* 结束会议不会释放 Rooms，如需释放请调用 [释放会议室（Rooms）](https://cloud.tencent.com/document/product/1095/55314) 接口
*/
func (s *meetingControlAPIService) V1MeetingsMeetingIdDismissPost(ctx context.Context, request *ApiV1MeetingsMeetingIdDismissPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdDismissPostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings/{meeting_id}/dismiss",
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

	response = &ApiV1MeetingsMeetingIdDismissPostResponse{
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

type ApiV1RealControlMeetingsMeetingIdAsrPutRequest struct {
	MeetingId string                                       `json:"-"`
	Body      *V1RealControlMeetingsMeetingIdAsrPutRequest `json:"body,omitempty"`
}

type ApiV1RealControlMeetingsMeetingIdAsrPutResponse struct {
	*xhttp.ApiResponse
	Data *V1RealControlMeetingsMeetingIdAsrPut200Response `json:"data,omitempty"`
}

/*
V1RealControlMeetingsMeetingIdAsrPut 开启或关闭实时转写[/v1/real-control/meetings/{meeting_id}/asr - Put]

以创建者的身份开启/关闭会中实时转写，调用时需要会议处于进行中的状态；
*/
func (s *meetingControlAPIService) V1RealControlMeetingsMeetingIdAsrPut(ctx context.Context, request *ApiV1RealControlMeetingsMeetingIdAsrPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1RealControlMeetingsMeetingIdAsrPutResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/real-control/meetings/{meeting_id}/asr",
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

	response = &ApiV1RealControlMeetingsMeetingIdAsrPutResponse{
		ApiResponse: apiRsp,
		Data:        new(V1RealControlMeetingsMeetingIdAsrPut200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1RealControlMeetingsMeetingIdCohostsPutRequest struct {
	MeetingId string                                           `json:"-"`
	Body      *V1RealControlMeetingsMeetingIdCohostsPutRequest `json:"body,omitempty"`
}

type ApiV1RealControlMeetingsMeetingIdCohostsPutResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1RealControlMeetingsMeetingIdCohostsPut 设置联席主持人[/v1/real-control/meetings/{meeting_id}/cohosts - Put]

**描述**：设置或撤销会议中的参会者联席主持人身份，目前暂不支持 MRA 设备作为被操作者的情况。企业 secret 鉴权用户可管理任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能管理通过 OAuth2.0 鉴权创建的有效会议。

*

> 说明
>
>
> * 操作者必须是会议的主持人才可以设置。
> * 调用该接口需要权限项 ： MANAGE_MEETING 查看和管理您的会议。
*/
func (s *meetingControlAPIService) V1RealControlMeetingsMeetingIdCohostsPut(ctx context.Context, request *ApiV1RealControlMeetingsMeetingIdCohostsPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1RealControlMeetingsMeetingIdCohostsPutResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/real-control/meetings/{meeting_id}/cohosts",
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

	response = &ApiV1RealControlMeetingsMeetingIdCohostsPutResponse{
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

type ApiV1RealControlMeetingsMeetingIdDocPutRequest struct {
	MeetingId string                                       `json:"-"`
	Body      *V1RealControlMeetingsMeetingIdDocPutRequest `json:"body,omitempty"`
}

type ApiV1RealControlMeetingsMeetingIdDocPutResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1RealControlMeetingsMeetingIdDocPut 文档上传权限设置[/v1/real-control/meetings/{meeting_id}/doc - Put]

**描述**：设置会议中文档上传权限，目前暂不支持 MRA 设备作为被操作者的情况，企业 secret 鉴权用户可管理任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能管理通过 OAuth2.0 鉴权创建的有效会议。

*

> 说明
> 1：操作者必须是会议的主持人或者联席主持人才可以设置。
> 2：调用该接口需要权限项：MANAGE_MEETING 查看和管理您的会议
*/
func (s *meetingControlAPIService) V1RealControlMeetingsMeetingIdDocPut(ctx context.Context, request *ApiV1RealControlMeetingsMeetingIdDocPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1RealControlMeetingsMeetingIdDocPutResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/real-control/meetings/{meeting_id}/doc",
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

	response = &ApiV1RealControlMeetingsMeetingIdDocPutResponse{
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

type ApiV1RealControlMeetingsMeetingIdKickoutPutRequest struct {
	MeetingId string                                           `json:"-"`
	Body      *V1RealControlMeetingsMeetingIdKickoutPutRequest `json:"body,omitempty"`
}

type ApiV1RealControlMeetingsMeetingIdKickoutPutResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1RealControlMeetingsMeetingIdKickoutPut 移出用户[/v1/real-control/meetings/{meeting_id}/kickout - Put]

**描述**：

* 将会议中用户移出会议，操作者必须是会议的主持人或者联席主持人才可以设置。
* 支持对云会议已入会成员和 Webinar 观众移出。
* 支持 MRA 设备和 PSTN 设备为被操作者时的移出用户操作。
* 企业 secret 鉴权用户可管理任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能管理通过 OAuth2.0 鉴权创建的有效会议。

说明

*

> 1：操作者必须是会议的主持人或者联席主持人才可以设置。
> 2：调用该接口需要权限项：MANAGE_MEETING 查看和管理您的会议。
*/
func (s *meetingControlAPIService) V1RealControlMeetingsMeetingIdKickoutPut(ctx context.Context, request *ApiV1RealControlMeetingsMeetingIdKickoutPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1RealControlMeetingsMeetingIdKickoutPutResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/real-control/meetings/{meeting_id}/kickout",
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

	response = &ApiV1RealControlMeetingsMeetingIdKickoutPutResponse{
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

type ApiV1RealControlMeetingsMeetingIdMutesPutRequest struct {
	MeetingId string                                         `json:"-"`
	Body      *V1RealControlMeetingsMeetingIdMutesPutRequest `json:"body,omitempty"`
}

type ApiV1RealControlMeetingsMeetingIdMutesPutResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1RealControlMeetingsMeetingIdMutesPut 静音用户[/v1/real-control/meetings/{meeting_id}/mutes - Put]

**描述**：

* 会议中用户静音操作，操作者必须是会议的主持人或者联席主持人才可以设置。
* 支持对云会议已入会成员和 Webinar 观众静音。
* 支持 MRA 设备和 PSTN 设备作为被操作者时的静音操作。
* 企业 secret 鉴权用户可管理任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能管理通过 OAuth2.0 鉴权创建的有效会议。

说明
1：操作者必须是会议的主持人或者联席主持人才可以设置。
2：调用该接口需要权限项：MANAGE_MEETING 查看和管理您的会议。
*/
func (s *meetingControlAPIService) V1RealControlMeetingsMeetingIdMutesPut(ctx context.Context, request *ApiV1RealControlMeetingsMeetingIdMutesPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1RealControlMeetingsMeetingIdMutesPutResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/real-control/meetings/{meeting_id}/mutes",
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

	response = &ApiV1RealControlMeetingsMeetingIdMutesPutResponse{
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

type ApiV1RealControlMeetingsMeetingIdNamesPutRequest struct {
	MeetingId string                                         `json:"-"`
	Body      *V1RealControlMeetingsMeetingIdNamesPutRequest `json:"body,omitempty"`
}

type ApiV1RealControlMeetingsMeetingIdNamesPutResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1RealControlMeetingsMeetingIdNamesPut 更改会中成员昵称[/v1/real-control/meetings/{meeting_id}/names - Put]

**描述：**

* 会中修改参会者昵称，支持主持人和联席主持人对会中成员进行改名。
* 此接口支持对云会议已入会成员和 Webinar 观众进行改名。
* 操作者必须为已在会中的主持人与联席主持人。
* 支持 MRA 设备和 PSTN 设备作为被操作者时的改名操作。
* 企业 secret 鉴权用户可管理任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能管理通过 OAuth2.0 鉴权创建的有效会议。

说明
1：操作者必须是会议的主持人或者联席主持人才可以设置。
2：调用该接口需要权限项：MANAGE_MEETING 查看和管理您的会议。
*/
func (s *meetingControlAPIService) V1RealControlMeetingsMeetingIdNamesPut(ctx context.Context, request *ApiV1RealControlMeetingsMeetingIdNamesPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1RealControlMeetingsMeetingIdNamesPutResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/real-control/meetings/{meeting_id}/names",
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

	response = &ApiV1RealControlMeetingsMeetingIdNamesPutResponse{
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

type ApiV1RealControlMeetingsMeetingIdScreenSharedPutRequest struct {
	MeetingId string                                                `json:"-"`
	Body      *V1RealControlMeetingsMeetingIdScreenSharedPutRequest `json:"body,omitempty"`
}

type ApiV1RealControlMeetingsMeetingIdScreenSharedPutResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1RealControlMeetingsMeetingIdScreenSharedPut 关闭用户屏幕共享[/v1/real-control/meetings/{meeting_id}/screen-shared - Put]

**描述**：关闭会议中用户屏幕共享权限，目前暂不支持 MRA 设备作为被操作者的情况。企业 secret 鉴权用户可管理任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能管理通过 OAuth2.0 鉴权创建的有效会议。

*

> 说明
> 1：操作者必须是会议的主持人或者联席主持人才可以设置。
> 2：调用该接口需要权限项：MANAGE_MEETING 查看和管理您的会议
*/
func (s *meetingControlAPIService) V1RealControlMeetingsMeetingIdScreenSharedPut(ctx context.Context, request *ApiV1RealControlMeetingsMeetingIdScreenSharedPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1RealControlMeetingsMeetingIdScreenSharedPutResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/real-control/meetings/{meeting_id}/screen-shared",
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

	response = &ApiV1RealControlMeetingsMeetingIdScreenSharedPutResponse{
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

type ApiV1RealControlMeetingsMeetingIdStatusPutRequest struct {
	MeetingId string                                          `json:"-"`
	Body      *V1RealControlMeetingsMeetingIdStatusPutRequest `json:"body,omitempty"`
}

type ApiV1RealControlMeetingsMeetingIdStatusPutResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1RealControlMeetingsMeetingIdStatusPut 会中状态设置[/v1/real-control/meetings/{meeting_id}/status - Put]

**描述**：设置会议中的会议属性，例如：全体静音、是否允许参会者聊天设置、锁定会议、隐藏会议号和密码、是否开启等候室等，具体设置内容可以查询接口协议，目前暂不支持 MRA 设备作为被操作者的情况。企业 secret 鉴权用户可管理任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能管理通过 OAuth2.0 鉴权创建的有效会议。

*

> 说明
> 1：操作者必须是会议的主持人或者联席主持人才可以设置。
> 2：调用该接口需要权限项 ：MANAGE_MEETING 查看和管理您的会议。
*/
func (s *meetingControlAPIService) V1RealControlMeetingsMeetingIdStatusPut(ctx context.Context, request *ApiV1RealControlMeetingsMeetingIdStatusPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1RealControlMeetingsMeetingIdStatusPutResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/real-control/meetings/{meeting_id}/status",
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

	response = &ApiV1RealControlMeetingsMeetingIdStatusPutResponse{
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

type ApiV1RealControlMeetingsMeetingIdVideoPutRequest struct {
	MeetingId string                                         `json:"-"`
	Body      *V1RealControlMeetingsMeetingIdVideoPutRequest `json:"body,omitempty"`
}

type ApiV1RealControlMeetingsMeetingIdVideoPutResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1RealControlMeetingsMeetingIdVideoPut 关闭或开启用户视频[/v1/real-control/meetings/{meeting_id}/video - Put]

关闭指定用户视频，支持关闭或开启 MRA 设备的视频。
企业 secret 鉴权用户可管理任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能管理通过 OAuth2.0 鉴权创建的有效会议。
> 说明
> 1：操作者必须是会议的主持人或者联席主持人才可以设置。
> 2：使用 OAuth 2.0 鉴权方式时，调用该接口需要权限项：MANAGE_MEETING 查看和管理您的会议。
*/
func (s *meetingControlAPIService) V1RealControlMeetingsMeetingIdVideoPut(ctx context.Context, request *ApiV1RealControlMeetingsMeetingIdVideoPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1RealControlMeetingsMeetingIdVideoPutResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/real-control/meetings/{meeting_id}/video",
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

	response = &ApiV1RealControlMeetingsMeetingIdVideoPutResponse{
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

type ApiV1RealControlMeetingsMeetingIdWaitingRoomPutRequest struct {
	MeetingId string                                               `json:"-"`
	Body      *V1RealControlMeetingsMeetingIdWaitingRoomPutRequest `json:"body,omitempty"`
}

type ApiV1RealControlMeetingsMeetingIdWaitingRoomPutResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1RealControlMeetingsMeetingIdWaitingRoomPut 用户等候室设置[/v1/real-control/meetings/{meeting_id}/waiting-room - Put]

会议等候室设置，允许主持人将等候室成员移入会议、主持人将会议成员移入等候室、主持人将等候室成员移出等候室等操作，目前暂不支持 MRA 设备作为被操作者的情况。企业 secret 鉴权用户可管理任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能管理通过 OAuth2.0 鉴权创建的有效会议。
> 说明
> 1：操作者必须是会议的主持人或者联席主持人才可以设置。
> 2：调用该接口需要权限项：MANAGE_MEETING 查看和管理您的会议
*/
func (s *meetingControlAPIService) V1RealControlMeetingsMeetingIdWaitingRoomPut(ctx context.Context, request *ApiV1RealControlMeetingsMeetingIdWaitingRoomPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1RealControlMeetingsMeetingIdWaitingRoomPutResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/real-control/meetings/{meeting_id}/waiting-room",
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

	response = &ApiV1RealControlMeetingsMeetingIdWaitingRoomPutResponse{
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
