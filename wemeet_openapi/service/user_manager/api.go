// Package wemeetopenapi is auto generate
/*
   腾讯会议OpenAPI

   SAAS版RESTFUL风格API

   API version: v1.0.2
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
	   V1AuthUsersCancelAuthPut 取消用户授权[/v1/auth-users/cancel-auth - Put]

	   第三方应用可以调用该接口来取消用户的授权，针对商业版和企业版用户仅支持在授权用户所属企业开启允许企业成员自主授权应用模式时取消，且由企业管理员开通的应用无法通过接口进行取消。如果企业开启了仅管理员可授权应用，用户只能在 腾讯会议应用管理页取消授权，无法在第三方平台取消。仅支持 OAuth2.0 鉴权方式调用。

	*/
	V1AuthUsersCancelAuthPut(ctx context.Context, request *ApiV1AuthUsersCancelAuthPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1AuthUsersCancelAuthPutResponse, err error)

	/*
	   V1MeetingsMeetingIdMsOpenIdGet 查询 ms_open_id[/v1/meetings/{meeting_id}/ms-open-id - Get]

	   **查询指定会议的用户的 ms\_open\_id，支持在会议开始前查询。**
	   支持企业自建应用（JWT 鉴权），仅支持查询本企业创建的会议。
	   <span class="colour" style="color:rgb(44, 51, 60)">支持OAuth2.0鉴权，仅支持查询该应用所创建的会议。</span>

	*/
	V1MeetingsMeetingIdMsOpenIdGet(ctx context.Context, request *ApiV1MeetingsMeetingIdMsOpenIdGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdMsOpenIdGetResponse, err error)

	/*
	   V1PmiMeetingsPmiConfigGet 查询个人会议号配置信息[/v1/pmi-meetings/pmi-config - Get]

	   获取用户个人会议号配置信息。仅企业下 secret 鉴权用户可获取该用户的 pmi 配置。目前暂不支持 OAuth 2.0鉴权访问。

	*/
	V1PmiMeetingsPmiConfigGet(ctx context.Context, request *ApiV1PmiMeetingsPmiConfigGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1PmiMeetingsPmiConfigGetResponse, err error)

	/*
	   V1PmiMeetingsPmiConfigPut 修改个人会议号配置信息[/v1/pmi-meetings/pmi-config - Put]

	   修改个人会议号的基本配置信息

	*/
	V1PmiMeetingsPmiConfigPut(ctx context.Context, request *ApiV1PmiMeetingsPmiConfigPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1PmiMeetingsPmiConfigPutResponse, err error)

	/*
	   V1UsersAccountAiAccountDelete 移除AI账号能力[/v1/users/account/ai-account - Delete]

	   移除企业账号的AI账号能力
	   权限点：企业用户管理，待自建应用支持权限点需求上线后生效

	*/
	V1UsersAccountAiAccountDelete(ctx context.Context, request *ApiV1UsersAccountAiAccountDeleteRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersAccountAiAccountDeleteResponse, err error)

	/*
	   V1UsersAccountAiAccountPost 添加AI账号能力[/v1/users/account/ai-account - Post]

	   设置企业账号AI账号能力
	   权限点：企业用户管理。

	*/
	V1UsersAccountAiAccountPost(ctx context.Context, request *ApiV1UsersAccountAiAccountPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersAccountAiAccountPostResponse, err error)

	/*
	   V1UsersAccountStatisticsGet 获取账号资源统计[/v1/users/account/statistics - Get]

	   查询企业下账号资源使用情况。
	   自建应用权限点：企业用户查看

	*/
	V1UsersAccountStatisticsGet(ctx context.Context, request *ApiV1UsersAccountStatisticsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersAccountStatisticsGetResponse, err error)

	/*
	   V1UsersDelete 删除用户（通过 uuid 删除用户）[/v1/users - Delete]

	*/
	V1UsersDelete(ctx context.Context, request *ApiV1UsersDeleteRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersDeleteResponse, err error)

	/*
	   V1UsersDeleteTransferPost 用户资产转移[/v1/users/delete-transfer - Post]

	*/
	V1UsersDeleteTransferPost(ctx context.Context, request *ApiV1UsersDeleteTransferPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersDeleteTransferPostResponse, err error)

	/*
	   V1UsersGet 获取用户详情（通过 uuid 获取用户详情）[/v1/users - Get]

	   使用 uuid 获取企业用户详情。企业 secret 鉴权用户可获取该用户所属企业下的用户详情，暂不支持 OAuth2.0 鉴权访问。

	*/
	V1UsersGet(ctx context.Context, request *ApiV1UsersGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersGetResponse, err error)

	/*
	   V1UsersInfoBasicGet 获取用户基本信息[/v1/users/info/basic - Get]

	*/
	V1UsersInfoBasicGet(ctx context.Context, request *ApiV1UsersInfoBasicGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersInfoBasicGetResponse, err error)

	/*
	   V1UsersInviteActivatePost 获取账号激活链接[/v1/users/invite-activate - Post]

	   未激活的账号，可以获取激活链接，激活链接有效期是48h。
	   每次获取链接为一个新链接，账号信息不变，旧链接仍然48h有效。

	*/
	V1UsersInviteActivatePost(ctx context.Context, request *ApiV1UsersInviteActivatePostRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersInviteActivatePostResponse, err error)

	/*
	   V1UsersInviteAuthPost 获取安全验证链接[/v1/users/invite-auth - Post]

	   ●未验证的账号，可以获取验证链接，验证链接有效期是 48h，每次获取链接为一个新链接，账号信息不变，旧链接仍然48h有效。
	   ●如果没有绑定手机号，不支持调用。
	   ●每个 userid每天可获取10次验证链接。


	*/
	V1UsersInviteAuthPost(ctx context.Context, request *ApiV1UsersInviteAuthPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersInviteAuthPostResponse, err error)

	/*
	   V1UsersListGet 获取用户列表[/v1/users/list - Get]

	   获取企业用户列表，目前暂不支持 OAuth2.0 鉴权访问。

	*/
	V1UsersListGet(ctx context.Context, request *ApiV1UsersListGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersListGetResponse, err error)

	/*
	   V1UsersOpenIdToUseridPost 自建应用与三方应用 ID 转换接口[/v1/users/open-id-to-userid - Post]

	   **接口描述：**
	   <span class="colour" style="color:rgb(24, 43, 80)">将三方应用获取到open_id转换为本企业用户的userid。</span>
	   **鉴权方式：**
	   JWT鉴权~~~~

	*/
	V1UsersOpenIdToUseridPost(ctx context.Context, request *ApiV1UsersOpenIdToUseridPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersOpenIdToUseridPostResponse, err error)

	/*
	   V1UsersPost 创建用户[/v1/users - Post]

	*/
	V1UsersPost(ctx context.Context, request *ApiV1UsersPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersPostResponse, err error)

	/*
	   V1UsersPut 更新用户（通过 uuid 更新用户）[/v1/users - Put]

	*/
	V1UsersPut(ctx context.Context, request *ApiV1UsersPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersPutResponse, err error)

	/*
	   V1UsersUseridDelete 删除用户（通过 userid 删除用户）[/v1/users/{userid} - Delete]

	*/
	V1UsersUseridDelete(ctx context.Context, request *ApiV1UsersUseridDeleteRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersUseridDeleteResponse, err error)

	/*
	   V1UsersUseridEnablePut 启用与禁用用户[/v1/users/{userid}/enable - Put]

	   **接口描述：**
	   使用userid启用/禁用本企业下的用户。~~~~
	   **鉴权方式：**
	   jwt鉴权
	   **输出参数：**
	   <span class="colour" style="color:rgb(51, 51, 51)">无输出参数，成功返回空消息体，失败返回 [错误码](https://cloud.tencent.com/document/product/1095/43704) 和错误信息。</span>

	*/
	V1UsersUseridEnablePut(ctx context.Context, request *ApiV1UsersUseridEnablePutRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersUseridEnablePutResponse, err error)

	/*
	   V1UsersUseridGet 获取用户详情（通过 userid 获取用户详情）[/v1/users/{userid} - Get]

	*/
	V1UsersUseridGet(ctx context.Context, request *ApiV1UsersUseridGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersUseridGetResponse, err error)

	/*
	   V1UsersUseridInviteActivatePut 发送用户激活邀请[/v1/users/{userid}/invite-activate - Put]

	   通过 userid 发送认证短信或邮件，邀请用户认证账号，用户确认后账号变为激活态。若使用手机号创建发送短信，使用邮箱创建发送邮件。
	   仅未激活的用户能够成功发送激活邀请。
	   每个手机号或邮箱一天只能发送一次邀请


	*/
	V1UsersUseridInviteActivatePut(ctx context.Context, request *ApiV1UsersUseridInviteActivatePutRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersUseridInviteActivatePutResponse, err error)

	/*
	   V1UsersUseridInviteAuthPut 用户安全验证[/v1/users/{userid}/invite-auth - Put]

	   ●通过 userid 发送验证短信，邀请成员验证账号，成员确认后账号变为已认证状态。
	   ●仅已激活的用户能够成功发送验证短信。
	   ●每个手机号一天只能发送一次邀请验证。


	*/
	V1UsersUseridInviteAuthPut(ctx context.Context, request *ApiV1UsersUseridInviteAuthPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersUseridInviteAuthPutResponse, err error)

	/*
	   V1UsersUseridPut 更新用户（通过 userid 更新用户）[/v1/users/{userid} - Put]

	*/
	V1UsersUseridPut(ctx context.Context, request *ApiV1UsersUseridPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersUseridPutResponse, err error)
}

type userManagerAPIService struct {
	config *core.Config
}

func NewService(config *core.Config) Service {
	return &userManagerAPIService{
		config: config,
	}
}

type ApiV1AuthUsersCancelAuthPutRequest struct {
	Body *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1AuthUsersCancelAuthPutResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1AuthUsersCancelAuthPut 取消用户授权[/v1/auth-users/cancel-auth - Put]

第三方应用可以调用该接口来取消用户的授权，针对商业版和企业版用户仅支持在授权用户所属企业开启允许企业成员自主授权应用模式时取消，且由企业管理员开通的应用无法通过接口进行取消。如果企业开启了仅管理员可授权应用，用户只能在 腾讯会议应用管理页取消授权，无法在第三方平台取消。仅支持 OAuth2.0 鉴权方式调用。
*/
func (s *userManagerAPIService) V1AuthUsersCancelAuthPut(ctx context.Context, request *ApiV1AuthUsersCancelAuthPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1AuthUsersCancelAuthPutResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/auth-users/cancel-auth",
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

	response = &ApiV1AuthUsersCancelAuthPutResponse{
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

type ApiV1MeetingsMeetingIdMsOpenIdGetRequest struct {
	MeetingId string `json:"-"`
	// 操作者 ID。operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。
	OperatorId *string `json:"-"`
	// 操作者 ID 的类型： 1: userid 2: open_id 3. rooms_id
	OperatorIdType *string                 `json:"-"`
	Body           *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1MeetingsMeetingIdMsOpenIdGetResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingsMeetingIdMsOpenIdGet200Response `json:"data,omitempty"`
}

/*
V1MeetingsMeetingIdMsOpenIdGet 查询 ms_open_id[/v1/meetings/{meeting_id}/ms-open-id - Get]

**查询指定会议的用户的 ms\_open\_id，支持在会议开始前查询。**
支持企业自建应用（JWT 鉴权），仅支持查询本企业创建的会议。
<span class="colour" style="color:rgb(44, 51, 60)">支持OAuth2.0鉴权，仅支持查询该应用所创建的会议。</span>
*/
func (s *userManagerAPIService) V1MeetingsMeetingIdMsOpenIdGet(ctx context.Context, request *ApiV1MeetingsMeetingIdMsOpenIdGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdMsOpenIdGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings/{meeting_id}/ms-open-id",
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

	// path 参数
	apiReq.PathParams.Set("meeting_id", core.PathValue(request.MeetingId))
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

	response = &ApiV1MeetingsMeetingIdMsOpenIdGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingsMeetingIdMsOpenIdGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1PmiMeetingsPmiConfigGetRequest struct {
	Userid *string `json:"-"`
	// 用户的终端设备类型： 0：PSTN 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：Linux 20：Rooms for Touch Windows 21：Rooms for Touch MacOS 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch iOS
	Instanceid *string                 `json:"-"`
	Body       *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1PmiMeetingsPmiConfigGetResponse struct {
	*xhttp.ApiResponse
	Data *V1PmiMeetingsPmiConfigGet200Response `json:"data,omitempty"`
}

/*
V1PmiMeetingsPmiConfigGet 查询个人会议号配置信息[/v1/pmi-meetings/pmi-config - Get]

获取用户个人会议号配置信息。仅企业下 secret 鉴权用户可获取该用户的 pmi 配置。目前暂不支持 OAuth 2.0鉴权访问。
*/
func (s *userManagerAPIService) V1PmiMeetingsPmiConfigGet(ctx context.Context, request *ApiV1PmiMeetingsPmiConfigGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1PmiMeetingsPmiConfigGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/pmi-meetings/pmi-config",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	if request.Userid == nil {
		return nil, fmt.Errorf("userid is required and must be specified")
	}

	// path 参数
	// query 参数
	if request.Instanceid != nil {
		apiReq.QueryParams.Set("instanceid", core.QueryValue(request.Instanceid))
	}
	if request.Userid != nil {
		apiReq.QueryParams.Set("userid", core.QueryValue(request.Userid))
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

	response = &ApiV1PmiMeetingsPmiConfigGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1PmiMeetingsPmiConfigGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1PmiMeetingsPmiConfigPutRequest struct {
	Body *V1PmiMeetingsPmiConfigPutRequest `json:"body,omitempty"`
}

type ApiV1PmiMeetingsPmiConfigPutResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1PmiMeetingsPmiConfigPut 修改个人会议号配置信息[/v1/pmi-meetings/pmi-config - Put]

修改个人会议号的基本配置信息
*/
func (s *userManagerAPIService) V1PmiMeetingsPmiConfigPut(ctx context.Context, request *ApiV1PmiMeetingsPmiConfigPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1PmiMeetingsPmiConfigPutResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/pmi-meetings/pmi-config",
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

	response = &ApiV1PmiMeetingsPmiConfigPutResponse{
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

type ApiV1UsersAccountAiAccountDeleteRequest struct {
	Body *V1UsersAccountAiAccountDeleteRequest `json:"body,omitempty"`
}

type ApiV1UsersAccountAiAccountDeleteResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1UsersAccountAiAccountDelete 移除AI账号能力[/v1/users/account/ai-account - Delete]

移除企业账号的AI账号能力
权限点：企业用户管理，待自建应用支持权限点需求上线后生效
*/
func (s *userManagerAPIService) V1UsersAccountAiAccountDelete(ctx context.Context, request *ApiV1UsersAccountAiAccountDeleteRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersAccountAiAccountDeleteResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/users/account/ai-account",
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

	apiRsp, err := s.config.Clt.Delete(ctx, apiReq, httpOptions...)
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

	response = &ApiV1UsersAccountAiAccountDeleteResponse{
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

type ApiV1UsersAccountAiAccountPostRequest struct {
	Body *V1UsersAccountAiAccountPostRequest `json:"body,omitempty"`
}

type ApiV1UsersAccountAiAccountPostResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1UsersAccountAiAccountPost 添加AI账号能力[/v1/users/account/ai-account - Post]

设置企业账号AI账号能力
权限点：企业用户管理。
*/
func (s *userManagerAPIService) V1UsersAccountAiAccountPost(ctx context.Context, request *ApiV1UsersAccountAiAccountPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersAccountAiAccountPostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/users/account/ai-account",
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

	response = &ApiV1UsersAccountAiAccountPostResponse{
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

type ApiV1UsersAccountStatisticsGetRequest struct {
	// 操作人ID，用户拥有企管用户查看权限
	OperatorId *string `json:"-"`
	// 操作人ID类型 1:userid
	OperatorIdType *string                 `json:"-"`
	Body           *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1UsersAccountStatisticsGetResponse struct {
	*xhttp.ApiResponse
	Data *V1UsersAccountStatisticsGet200Response `json:"data,omitempty"`
}

/*
V1UsersAccountStatisticsGet 获取账号资源统计[/v1/users/account/statistics - Get]

查询企业下账号资源使用情况。
自建应用权限点：企业用户查看
*/
func (s *userManagerAPIService) V1UsersAccountStatisticsGet(ctx context.Context, request *ApiV1UsersAccountStatisticsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersAccountStatisticsGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/users/account/statistics",
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

	// path 参数
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

	response = &ApiV1UsersAccountStatisticsGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1UsersAccountStatisticsGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1UsersDeleteRequest struct {
	Uuid *string                 `json:"-"`
	Body *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1UsersDeleteResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1UsersDelete 删除用户（通过 uuid 删除用户）[/v1/users - Delete]
*/
func (s *userManagerAPIService) V1UsersDelete(ctx context.Context, request *ApiV1UsersDeleteRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersDeleteResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/users",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	if request.Uuid == nil {
		return nil, fmt.Errorf("uuid is required and must be specified")
	}

	// path 参数
	// query 参数
	if request.Uuid != nil {
		apiReq.QueryParams.Set("uuid", core.QueryValue(request.Uuid))
	}
	// 转换 options
	var httpOptions []xhttp.RequestOptionFunc
	for _, opt := range opts {
		httpOptions = append(httpOptions, opt(*s.config))
	}
	// 增加 SDK Version 标识
	httpOptions = append(httpOptions, xhttp.WithRequestAuthenticator(core.DefaultAuthenticator))

	apiRsp, err := s.config.Clt.Delete(ctx, apiReq, httpOptions...)
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

	response = &ApiV1UsersDeleteResponse{
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

type ApiV1UsersDeleteTransferPostRequest struct {
	Body *V1UsersDeleteTransferPostRequest `json:"body,omitempty"`
}

type ApiV1UsersDeleteTransferPostResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1UsersDeleteTransferPost 用户资产转移[/v1/users/delete-transfer - Post]
*/
func (s *userManagerAPIService) V1UsersDeleteTransferPost(ctx context.Context, request *ApiV1UsersDeleteTransferPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersDeleteTransferPostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/users/delete-transfer",
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

	response = &ApiV1UsersDeleteTransferPostResponse{
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

type ApiV1UsersGetRequest struct {
	Uuid *string                 `json:"-"`
	Body *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1UsersGetResponse struct {
	*xhttp.ApiResponse
	Data *V1UsersGet200Response `json:"data,omitempty"`
}

/*
V1UsersGet 获取用户详情（通过 uuid 获取用户详情）[/v1/users - Get]

使用 uuid 获取企业用户详情。企业 secret 鉴权用户可获取该用户所属企业下的用户详情，暂不支持 OAuth2.0 鉴权访问。
*/
func (s *userManagerAPIService) V1UsersGet(ctx context.Context, request *ApiV1UsersGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/users",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	if request.Uuid == nil {
		return nil, fmt.Errorf("uuid is required and must be specified")
	}

	// path 参数
	// query 参数
	if request.Uuid != nil {
		apiReq.QueryParams.Set("uuid", core.QueryValue(request.Uuid))
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

	response = &ApiV1UsersGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1UsersGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1UsersInfoBasicGetRequest struct {
	// 操作者 ID，该接口不支持获取 MRA、Rooms、小程序的账号。 operator_id 必须与operator_id_type 配合使用。 根据 operator_id_type 的值，operator_id 代表不同类型。
	OperatorId *string `json:"-"`
	// 操作者 ID 的类型，2：openid。
	OperatorIdType *string                 `json:"-"`
	Body           *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1UsersInfoBasicGetResponse struct {
	*xhttp.ApiResponse
	Data *V1UsersInfoBasicGet200Response `json:"data,omitempty"`
}

/*
V1UsersInfoBasicGet 获取用户基本信息[/v1/users/info/basic - Get]
*/
func (s *userManagerAPIService) V1UsersInfoBasicGet(ctx context.Context, request *ApiV1UsersInfoBasicGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersInfoBasicGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/users/info/basic",
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

	// path 参数
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

	response = &ApiV1UsersInfoBasicGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1UsersInfoBasicGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1UsersInviteActivatePostRequest struct {
	Body *V1UsersInviteActivatePostRequest `json:"body,omitempty"`
}

type ApiV1UsersInviteActivatePostResponse struct {
	*xhttp.ApiResponse
	Data *V1UsersInviteActivatePost200Response `json:"data,omitempty"`
}

/*
V1UsersInviteActivatePost 获取账号激活链接[/v1/users/invite-activate - Post]

未激活的账号，可以获取激活链接，激活链接有效期是48h。
每次获取链接为一个新链接，账号信息不变，旧链接仍然48h有效。
*/
func (s *userManagerAPIService) V1UsersInviteActivatePost(ctx context.Context, request *ApiV1UsersInviteActivatePostRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersInviteActivatePostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/users/invite-activate",
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

	response = &ApiV1UsersInviteActivatePostResponse{
		ApiResponse: apiRsp,
		Data:        new(V1UsersInviteActivatePost200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1UsersInviteAuthPostRequest struct {
	Body *V1UsersInviteAuthPostRequest `json:"body,omitempty"`
}

type ApiV1UsersInviteAuthPostResponse struct {
	*xhttp.ApiResponse
	Data *V1UsersInviteAuthPost200Response `json:"data,omitempty"`
}

/*
V1UsersInviteAuthPost 获取安全验证链接[/v1/users/invite-auth - Post]

●未验证的账号，可以获取验证链接，验证链接有效期是 48h，每次获取链接为一个新链接，账号信息不变，旧链接仍然48h有效。
●如果没有绑定手机号，不支持调用。
●每个 userid每天可获取10次验证链接。
*/
func (s *userManagerAPIService) V1UsersInviteAuthPost(ctx context.Context, request *ApiV1UsersInviteAuthPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersInviteAuthPostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/users/invite-auth",
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

	response = &ApiV1UsersInviteAuthPostResponse{
		ApiResponse: apiRsp,
		Data:        new(V1UsersInviteAuthPost200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1UsersListGetRequest struct {
	// 当前页，大于等于1。
	Page *string `json:"-"`
	// 分页大小，最大为20。
	PageSize *string                 `json:"-"`
	Body     *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1UsersListGetResponse struct {
	*xhttp.ApiResponse
	Data *V1UsersListGet200Response `json:"data,omitempty"`
}

/*
V1UsersListGet 获取用户列表[/v1/users/list - Get]

获取企业用户列表，目前暂不支持 OAuth2.0 鉴权访问。
*/
func (s *userManagerAPIService) V1UsersListGet(ctx context.Context, request *ApiV1UsersListGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersListGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/users/list",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	if request.Page == nil {
		return nil, fmt.Errorf("page is required and must be specified")
	}

	if request.PageSize == nil {
		return nil, fmt.Errorf("page_size is required and must be specified")
	}

	// path 参数
	// query 参数
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

	response = &ApiV1UsersListGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1UsersListGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1UsersOpenIdToUseridPostRequest struct {
	Body *V1UsersOpenIdToUseridPostRequest `json:"body,omitempty"`
}

type ApiV1UsersOpenIdToUseridPostResponse struct {
	*xhttp.ApiResponse
	Data *V1UsersOpenIdToUseridPost200Response `json:"data,omitempty"`
}

/*
V1UsersOpenIdToUseridPost 自建应用与三方应用 ID 转换接口[/v1/users/open-id-to-userid - Post]

**接口描述：**
<span class="colour" style="color:rgb(24, 43, 80)">将三方应用获取到open_id转换为本企业用户的userid。</span>
**鉴权方式：**
JWT鉴权~~~~
*/
func (s *userManagerAPIService) V1UsersOpenIdToUseridPost(ctx context.Context, request *ApiV1UsersOpenIdToUseridPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersOpenIdToUseridPostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/users/open-id-to-userid",
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

	response = &ApiV1UsersOpenIdToUseridPostResponse{
		ApiResponse: apiRsp,
		Data:        new(V1UsersOpenIdToUseridPost200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1UsersPostRequest struct {
	Body *V1UsersPostRequest `json:"body,omitempty"`
}

type ApiV1UsersPostResponse struct {
	*xhttp.ApiResponse
	Data *V1UsersPost200Response `json:"data,omitempty"`
}

/*
V1UsersPost 创建用户[/v1/users - Post]
*/
func (s *userManagerAPIService) V1UsersPost(ctx context.Context, request *ApiV1UsersPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersPostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/users",
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

	response = &ApiV1UsersPostResponse{
		ApiResponse: apiRsp,
		Data:        new(V1UsersPost200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1UsersPutRequest struct {
	Uuid *string                 `json:"-"`
	Body *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1UsersPutResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1UsersPut 更新用户（通过 uuid 更新用户）[/v1/users - Put]
*/
func (s *userManagerAPIService) V1UsersPut(ctx context.Context, request *ApiV1UsersPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersPutResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/users",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	if request.Uuid == nil {
		return nil, fmt.Errorf("uuid is required and must be specified")
	}

	// path 参数
	// query 参数
	if request.Uuid != nil {
		apiReq.QueryParams.Set("uuid", core.QueryValue(request.Uuid))
	}
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

	response = &ApiV1UsersPutResponse{
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

type ApiV1UsersUseridDeleteRequest struct {
	// 被删除用户的userid
	Userid string                  `json:"-"`
	Body   *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1UsersUseridDeleteResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1UsersUseridDelete 删除用户（通过 userid 删除用户）[/v1/users/{userid} - Delete]
*/
func (s *userManagerAPIService) V1UsersUseridDelete(ctx context.Context, request *ApiV1UsersUseridDeleteRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersUseridDeleteResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/users/{userid}",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	// path 参数
	apiReq.PathParams.Set("userid", core.PathValue(request.Userid))
	// query 参数
	// 转换 options
	var httpOptions []xhttp.RequestOptionFunc
	for _, opt := range opts {
		httpOptions = append(httpOptions, opt(*s.config))
	}
	// 增加 SDK Version 标识
	httpOptions = append(httpOptions, xhttp.WithRequestAuthenticator(core.DefaultAuthenticator))

	apiRsp, err := s.config.Clt.Delete(ctx, apiReq, httpOptions...)
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

	response = &ApiV1UsersUseridDeleteResponse{
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

type ApiV1UsersUseridEnablePutRequest struct {
	// 调用方用于标示用户的唯一 ID（例如：企业用户可以为企业账户英文名、个人用户可以为手机号等，暂不支持中文）。
	Userid string                         `json:"-"`
	Body   *V1UsersUseridEnablePutRequest `json:"body,omitempty"`
}

type ApiV1UsersUseridEnablePutResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1UsersUseridEnablePut 启用与禁用用户[/v1/users/{userid}/enable - Put]

**接口描述：**
使用userid启用/禁用本企业下的用户。~~~~
**鉴权方式：**
jwt鉴权
**输出参数：**
<span class="colour" style="color:rgb(51, 51, 51)">无输出参数，成功返回空消息体，失败返回 [错误码](https://cloud.tencent.com/document/product/1095/43704) 和错误信息。</span>
*/
func (s *userManagerAPIService) V1UsersUseridEnablePut(ctx context.Context, request *ApiV1UsersUseridEnablePutRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersUseridEnablePutResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/users/{userid}/enable",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	// path 参数
	apiReq.PathParams.Set("userid", core.PathValue(request.Userid))
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

	response = &ApiV1UsersUseridEnablePutResponse{
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

type ApiV1UsersUseridGetRequest struct {
	// 调用方用于标示用户的唯一 ID（企业内部请使用企业唯一用户标识；OAuth2.0 鉴权用户请使用 openId）。 企业唯一用户标识说明： 1. 企业对接 SSO 时使用的员工唯一标识 ID； 2. 企业调用创建用户接口时传递的 userid 参数。
	Userid string                  `json:"-"`
	Body   *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1UsersUseridGetResponse struct {
	*xhttp.ApiResponse
	Data *V1UsersUseridGet200Response `json:"data,omitempty"`
}

/*
V1UsersUseridGet 获取用户详情（通过 userid 获取用户详情）[/v1/users/{userid} - Get]
*/
func (s *userManagerAPIService) V1UsersUseridGet(ctx context.Context, request *ApiV1UsersUseridGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersUseridGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/users/{userid}",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	// path 参数
	apiReq.PathParams.Set("userid", core.PathValue(request.Userid))
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

	response = &ApiV1UsersUseridGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1UsersUseridGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1UsersUseridInviteActivatePutRequest struct {
	// 调用方用于标示用户的唯一 ID（例如：企业用户可以为企业账户英文名、个人用户可以为手机号等，暂不支持中文）。
	Userid string                  `json:"-"`
	Body   *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1UsersUseridInviteActivatePutResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1UsersUseridInviteActivatePut 发送用户激活邀请[/v1/users/{userid}/invite-activate - Put]

通过 userid 发送认证短信或邮件，邀请用户认证账号，用户确认后账号变为激活态。若使用手机号创建发送短信，使用邮箱创建发送邮件。
仅未激活的用户能够成功发送激活邀请。
每个手机号或邮箱一天只能发送一次邀请
*/
func (s *userManagerAPIService) V1UsersUseridInviteActivatePut(ctx context.Context, request *ApiV1UsersUseridInviteActivatePutRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersUseridInviteActivatePutResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/users/{userid}/invite-activate",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	// path 参数
	apiReq.PathParams.Set("userid", core.PathValue(request.Userid))
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

	response = &ApiV1UsersUseridInviteActivatePutResponse{
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

type ApiV1UsersUseridInviteAuthPutRequest struct {
	Userid string                             `json:"-"`
	Body   *V1UsersUseridInviteAuthPutRequest `json:"body,omitempty"`
}

type ApiV1UsersUseridInviteAuthPutResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1UsersUseridInviteAuthPut 用户安全验证[/v1/users/{userid}/invite-auth - Put]

●通过 userid 发送验证短信，邀请成员验证账号，成员确认后账号变为已认证状态。
●仅已激活的用户能够成功发送验证短信。
●每个手机号一天只能发送一次邀请验证。
*/
func (s *userManagerAPIService) V1UsersUseridInviteAuthPut(ctx context.Context, request *ApiV1UsersUseridInviteAuthPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersUseridInviteAuthPutResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/users/{userid}/invite-auth",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	// path 参数
	apiReq.PathParams.Set("userid", core.PathValue(request.Userid))
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

	response = &ApiV1UsersUseridInviteAuthPutResponse{
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

type ApiV1UsersUseridPutRequest struct {
	Userid string                   `json:"-"`
	Body   *V1UsersUseridPutRequest `json:"body,omitempty"`
}

type ApiV1UsersUseridPutResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1UsersUseridPut 更新用户（通过 userid 更新用户）[/v1/users/{userid} - Put]
*/
func (s *userManagerAPIService) V1UsersUseridPut(ctx context.Context, request *ApiV1UsersUseridPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1UsersUseridPutResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/users/{userid}",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	// path 参数
	apiReq.PathParams.Set("userid", core.PathValue(request.Userid))
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

	response = &ApiV1UsersUseridPutResponse{
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
