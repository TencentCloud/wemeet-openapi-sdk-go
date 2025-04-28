// Package wemeetopenapi is auto generate
/*
   腾讯会议OpenAPI

   SAAS版RESTFUL风格API

   API version: v1.0.6
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
	   V1AsrConfigPut 设置语音识别热词[/v1/asr/config - Put]

	   用户可以通过接口进行您创建的会议的语音识别设置。
	   权限点：查看或管理您的会议

	*/
	V1AsrConfigPut(ctx context.Context, request *ApiV1AsrConfigPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1AsrConfigPutResponse, err error)

	/*
	   V1AsrDetailsGet 导出实时转写记录[/v1/asr/details - Get]

	   如果会议开启了会议转写，可以导出转写记录。主持人可以设置导出权限，默认主持人可以导出，支持会中和会后导出。

	*/
	V1AsrDetailsGet(ctx context.Context, request *ApiV1AsrDetailsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1AsrDetailsGetResponse, err error)

	/*
	   V1AsrPushStatusPost 开启/关闭实时转写推送[/v1/asr/push-status - Post]

	    会议创建者可开启本场会议的实时转写内容推送，待开始的会议或未打开实时转写功能的会议也支持开启推送，开启推送后如果会议打开转写，则可通过webhook 实时转写推送 实时获取到转写内容。
	   企业 secret 鉴权用户可开启该用户所属企业下的所有会议转写推送，OAuth2.0 鉴权用户只能开启通过 OAuth2.0 鉴权创建的会议转写推送。
	   企业级自建应用通过 webhook 可以接收到企业下所有开启推送的会议的转写内容，应用级自建应用通过 webhook 可以接收到本应用创建的会议的转写内容。OAuth2.0 应用通过 webhook 可以接收到 OAuth2.0 鉴权创建的会议的转写内容。


	*/
	V1AsrPushStatusPost(ctx context.Context, request *ApiV1AsrPushStatusPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1AsrPushStatusPostResponse, err error)

	/*
	   V1HistoryMeetingsUseridGet 查询用户已结束会议列表[/v1/history/meetings/{userid} - Get]

	   获取某指定用户的已结束的会议列表，可返回用户创建与参加过的已结束会议列表，支持 OAuth2.0 鉴权和企微鉴权。
	   请优先使用operator_id和operator_id_type查询，当使用operator_id和operator_id_type时，userid设置为operator_id的值即可

	*/
	V1HistoryMeetingsUseridGet(ctx context.Context, request *ApiV1HistoryMeetingsUseridGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1HistoryMeetingsUseridGetResponse, err error)

	/*
	   V1MeetingJobIdGet 获取导出 PSTN 通话详单任务结果[/v1/meeting/{job_id} - Get]

	   获取异步导出任务的结果。

	*/
	V1MeetingJobIdGet(ctx context.Context, request *ApiV1MeetingJobIdGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingJobIdGetResponse, err error)

	/*
	   V1MeetingMeetingIdWaitingRoomGet 查询等候室成员记录[/v1/meeting/{meeting_id}/waiting-room - Get]

	   会议创建者、主持人、联席主持人可以查询等候室成员列表，包括等候室内所有用户的进出记录。会前、会中、会后都可以查询。
	   “查询等候室成员列表”改为“获取实时等候室成员列表”，只有当前等候室的成员。如果是PMI会议，返回的是PMI的meeting_code。
	   鉴权方式：JWT鉴权、OAuth鉴权、SDK鉴权
	   OAuth鉴权的权限为：查询会议、查询和管理会议

	*/
	V1MeetingMeetingIdWaitingRoomGet(ctx context.Context, request *ApiV1MeetingMeetingIdWaitingRoomGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingMeetingIdWaitingRoomGetResponse, err error)

	/*
	   V1MeetingMeetingIdWaitingRoomWelcomeMessageGet 获取等候室欢迎语[/v1/meeting/{meeting_id}/waiting-room-welcome-message - Get]

	   查询会议的等候室欢迎语。当有用户进入等候室时，会收到来自会议主办方的私聊消息引导。
	   鉴权方式: JWT鉴权、OAuth鉴权

	*/
	V1MeetingMeetingIdWaitingRoomWelcomeMessageGet(ctx context.Context, request *ApiV1MeetingMeetingIdWaitingRoomWelcomeMessageGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingMeetingIdWaitingRoomWelcomeMessageGetResponse, err error)

	/*
	   V1MeetingSetWaitingRoomWelcomeMessagePost 设置等候室欢迎语[/v1/meeting/set-waiting-room-welcome-message - Post]

	   为已开启等候室的会议配置等候室欢迎语。当有用户进入等候室时，会收到来自会议主办方的私聊消息引导。

	   鉴权方式:
	   JWT鉴权、OAuth鉴权

	*/
	V1MeetingSetWaitingRoomWelcomeMessagePost(ctx context.Context, request *ApiV1MeetingSetWaitingRoomWelcomeMessagePostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingSetWaitingRoomWelcomeMessagePostResponse, err error)

	/*
	   V1MeetingsCustomerShortUrlPost 创建用户专属参会链接[/v1/meetings/customer-short-url - Post]

	   使用该接口，可用 `customer_data` 进行区分，为一场会议生成多个会议链接。通过用户入会、用户进入等候室等事件，或通过获取等候室成员列表的 API 查询到该参数。
	   该接口不支持个人会议号会议、网络研讨会（Webinar）。支持企业品牌化链接。
	   参会者腾讯会议客户端版本需大于等于 3.2.0。
	   暂不支持 OAuth 2.0 鉴权方式访问。

	*/
	V1MeetingsCustomerShortUrlPost(ctx context.Context, request *ApiV1MeetingsCustomerShortUrlPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsCustomerShortUrlPostResponse, err error)

	/*
	   V1MeetingsGet 查询用户的会议列表[/v1/meetings - Get]

	   通过会议CODE查询会议详情/查询用户的会议列表
	   ① 通过会议CODE查询会议详情
	   企业 secret 鉴权用户可查询到任何该用户创建的企业下的会议，OAuth2.0 鉴权用户只能查询到通过 OAuth2.0 鉴权创建的会议。
	   支持企业管理员查询企业下的会议。
	   本接口的邀请参会成员限制调整至300人。
	   当会议为周期性会议时，主持人密钥每场会议固定，但单场会议只能获取一次。支持查询周期性会议的主持人密钥。
	   支持查询 MRA 当前所在会议信息。
	   若会议号被回收则无法通过 Code 查询，您可以通过会议 ID 查询到该会议。
	   ② 查询用户的会议列表
	   获取某指定用户的进行中或待开始的会议列表。企业 secret 鉴权用户可查询任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能查询通过 OAuth2.0 鉴权创建的有效会议。

	*/
	V1MeetingsGet(ctx context.Context, request *ApiV1MeetingsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsGetResponse, err error)

	/*
	   V1MeetingsMeetingIdCancelPost 取消会议[/v1/meetings/{meeting_id}/cancel - Post]

	   取消用户创建的会议。用户只能取消自己创建的会议，且该会议是一个有效的会议。如果不是会议创建者或者无效会议号将会返回错误。
	   企业 secret 鉴权用户可取消任何该用户企业下创建的有效会议，OAuth2.0 鉴权用户只能取消通过 OAuth2.0 鉴权创建的有效会议。
	   当您想实时监测会议取消状况时，您可以通过订阅 [会议取消](https://cloud.tencent.com/document/product/1095/51616) 的事件，接收事件通知。

	*/
	V1MeetingsMeetingIdCancelPost(ctx context.Context, request *ApiV1MeetingsMeetingIdCancelPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdCancelPostResponse, err error)

	/*
	   V1MeetingsMeetingIdCustomerShortUrlGet 获取用户专属参会链接[/v1/meetings/{meeting_id}/customer-short-url - Get]

	   **描述**：

	   * 可以获取指定会议的所有专属参会链接及 `customer_data`。
	   * 该接口不支持个人会议号会议、网络研讨会（Webinar）。支持企业品牌化链接。
	   * 参会者腾讯会议客户端版本需大于等于 3.2.0。
	   * 暂不支持 OAuth 2.0 鉴权方式访问。

	*/
	V1MeetingsMeetingIdCustomerShortUrlGet(ctx context.Context, request *ApiV1MeetingsMeetingIdCustomerShortUrlGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdCustomerShortUrlGetResponse, err error)

	/*
	   V1MeetingsMeetingIdEnrollApprovalsGet 查询会议报名信息[/v1/meetings/{meeting_id}/enroll/approvals - Get]

	   查询已报名观众数量和报名观众答题详情，仅会议创建者可查询。
	   企业 secret 鉴权用户可修改任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能修改通过 OAuth2.0 鉴权创建的有效会议。
	   用户必须是注册用户，请求头部 X-TC-Registered 字段必须传入为1。

	*/
	V1MeetingsMeetingIdEnrollApprovalsGet(ctx context.Context, request *ApiV1MeetingsMeetingIdEnrollApprovalsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdEnrollApprovalsGetResponse, err error)

	/*
	   V1MeetingsMeetingIdEnrollApprovalsPut 审批会议报名信息[/v1/meetings/{meeting_id}/enroll/approvals - Put]

	   批量云会议的报名信息，仅会议创建者可审批。
	   企业 secret 鉴权用户可审批任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能审批通过 OAuth2.0 鉴权创建的有效会议。
	   用户必须是注册用户，请求头部 X-TC-Registered 字段必须传入为1。

	*/
	V1MeetingsMeetingIdEnrollApprovalsPut(ctx context.Context, request *ApiV1MeetingsMeetingIdEnrollApprovalsPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdEnrollApprovalsPutResponse, err error)

	/*
	   V1MeetingsMeetingIdEnrollConfigGet 查询会议报名配置[/v1/meetings/{meeting_id}/enroll/config - Get]

	   查询云会议的报名配置和报名问题，仅会议创建者可查询。会议未开启报名时会返回未开启报名错误。
	   企业 secret 鉴权用户可查询任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能查询通过 OAuth2.0 鉴权创建的有效会议。
	   用户必须是注册用户，请求头部 X-TC-Registered 字段必须传入为1。

	*/
	V1MeetingsMeetingIdEnrollConfigGet(ctx context.Context, request *ApiV1MeetingsMeetingIdEnrollConfigGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdEnrollConfigGetResponse, err error)

	/*
	   V1MeetingsMeetingIdEnrollConfigPut 修改会议报名配置[/v1/meetings/{meeting_id}/enroll/config - Put]

	   修改云会议的报名配置和报名问题，仅会议创建者可修改，且需要会议已开启报名。
	   企业 secret 鉴权用户可修改任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能修改通过 OAuth2.0 鉴权创建的有效会议。
	   用户必须是注册用户，请求头部 X-TC-Registered 字段必须传入为1。

	*/
	V1MeetingsMeetingIdEnrollConfigPut(ctx context.Context, request *ApiV1MeetingsMeetingIdEnrollConfigPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdEnrollConfigPutResponse, err error)

	/*
	   V1MeetingsMeetingIdEnrollIdsPost 查询会议成员报名 ID[/v1/meetings/{meeting_id}/enroll/ids - Post]

	   描述： 支持查询会议中已报名成员的报名 ID，仅会议创建者可查询。

	*/
	V1MeetingsMeetingIdEnrollIdsPost(ctx context.Context, request *ApiV1MeetingsMeetingIdEnrollIdsPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdEnrollIdsPostResponse, err error)

	/*
	   V1MeetingsMeetingIdEnrollImportPost 导入会议报名信息[/v1/meetings/{meeting_id}/enroll/import - Post]

	   指定会议中导入报名信息。

	   企业 secret 鉴权用户可通过同企业下用户 userid 和手机号导入报名信息，OAuth2.0 鉴权用户能通过用户 open_id，与应用同企业下的 userid 以及手机号导入报名信息。
	   用户必须是注册用户，请求头部 X-TC-Registered 字段必须传入为1。
	   商业版单场会议导入上限1000条，企业版单场会议导入上限4000条。如需提升，请联系我们。

	*/
	V1MeetingsMeetingIdEnrollImportPost(ctx context.Context, request *ApiV1MeetingsMeetingIdEnrollImportPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdEnrollImportPostResponse, err error)

	/*
	   V1MeetingsMeetingIdEnrollUnregistrationDelete 删除会议报名信息[/v1/meetings/{meeting_id}/enroll/unregistration - Delete]

	   描述：
	   删除指定会议的报名信息，支持删除用户手动报名的信息和导入的报名信息。
	   企业 secret 鉴权用户可删除该用户企业会议下的报名信息，OAuth2.0 鉴权用户只能删除通过 OAuth2.0 鉴权创建的有效会议的报名信息。
	   用户必须是注册用户，请求头部 X-TC-Registered 字段必须传入为1。

	*/
	V1MeetingsMeetingIdEnrollUnregistrationDelete(ctx context.Context, request *ApiV1MeetingsMeetingIdEnrollUnregistrationDeleteRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdEnrollUnregistrationDeleteResponse, err error)

	/*
	   V1MeetingsMeetingIdGet 查询会议[/v1/meetings/{meeting_id} - Get]

	   通过会议 ID 查询会议详情。
	   企业 secret 鉴权用户可查询到任何该用户创建的企业下的会议，OAuth2.0 鉴权用户只能查询到通过 OAuth2.0 鉴权创建的会议。
	   本接口的邀请参会成员限制调整至300人。
	   当会议为周期性会议时，主持人密钥每场会议固定，但单场会议只能获取一次。支持查询周期性会议的主持人密钥。
	   支持查询 MRA 当前所在会议信息。

	*/
	V1MeetingsMeetingIdGet(ctx context.Context, request *ApiV1MeetingsMeetingIdGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdGetResponse, err error)

	/*
	   V1MeetingsMeetingIdInviteesGet 获取会议受邀成员列表[/v1/meetings/{meeting_id}/invitees - Get]

	   根据会议ID获取受邀成员列表，支持分页获取，只有会议的创建者才有权限获取。

	*/
	V1MeetingsMeetingIdInviteesGet(ctx context.Context, request *ApiV1MeetingsMeetingIdInviteesGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdInviteesGetResponse, err error)

	/*
	   V1MeetingsMeetingIdInviteesPut 设置会议邀请成员[/v1/meetings/{meeting_id}/invitees - Put]

	   根据会议ID设置邀请成员，只有会议的创建者才有权限设置。
	   最多可以设置2000名邀请者。
	   注：本接口为覆盖式设置。

	*/
	V1MeetingsMeetingIdInviteesPut(ctx context.Context, request *ApiV1MeetingsMeetingIdInviteesPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdInviteesPutResponse, err error)

	/*
	   V1MeetingsMeetingIdParticipantsGet 获取参会成员列表[/v1/meetings/{meetingId}/participants - Get]

	   会议创建者和企业管理员可以查询参会成员的列表，其他用户的调用会被拒绝。

	   支持查询网络研讨会参会成员列表。
	   如果会议还未开始，调用此接口查询会返回空列表。
	   企业 secret 鉴权用户（会议创建者）可获取任何该企业该用户创建的有效会议中的参会成员，企业 secret 鉴权用户（企业超级管理员）可获取任何该企业下创建的有效会议中的参会成员，OAuth2.0 鉴权用户（会议创建者）只能获取用户通过 OAuth2.0 鉴权创建的有效会议中的参会成员。
	   当您需要实时监测参会成员入会状态或退会状态时，您可以通过订阅 [用户入会](https://cloud.tencent.com/document/product/1095/51620)和 [用户离会](https://cloud.tencent.com/document/product/1095/51622) 的事件，接收事件通知。

	*/
	V1MeetingsMeetingIdParticipantsGet(ctx context.Context, request *ApiV1MeetingsMeetingIdParticipantsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdParticipantsGetResponse, err error)

	/*
	   V1MeetingsMeetingIdPut 修改会议[/v1/meetings/{meeting_id} - Put]

	   修改某指定会议的会议信息。

	   企业 secret 鉴权用户可修改任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能修改通过 OAuth2.0 鉴权创建的有效会议。
	   当您想实时监测会议修改状况时，您可以通过订阅 [会议更新](https://cloud.tencent.com/document/product/1095/51615) 的事件，接收事件通知。
	   本接口的邀请参会成员限制调整至300人。
	   当会议为周期性会议时，主持人密钥每场会议固定，但单场会议只能获取一次。支持修改周期性会议的主持人密钥。

	*/
	V1MeetingsMeetingIdPut(ctx context.Context, request *ApiV1MeetingsMeetingIdPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdPutResponse, err error)

	/*
	   V1MeetingsMeetingIdQosGet 获取会议实时质量检测数据[/v1/meetings/{meeting_id}/qos - Get]

	   拥有企业“会议列表--会控”权限的成员，能够获取实时会议质量检测数据。
	   支持云会议和Webinar会议的数据。会议状态为进行中。

	*/
	V1MeetingsMeetingIdQosGet(ctx context.Context, request *ApiV1MeetingsMeetingIdQosGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdQosGetResponse, err error)

	/*
	   V1MeetingsMeetingIdQualityGet 查询会议健康度[/v1/meetings/{meeting_id}/quality - Get]

	   查询会议及参会成员的健康度，付费开通该服务的企业管理员、超管可以查询，与是否为会议创建者/主持人/联席主持人无关。
	   鉴权方式：支持 JWT 鉴权 和 Oauth 鉴权

	*/
	V1MeetingsMeetingIdQualityGet(ctx context.Context, request *ApiV1MeetingsMeetingIdQualityGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdQualityGetResponse, err error)

	/*
	   V1MeetingsMeetingIdRealTimeParticipantsGet 查询实时会中成员列表[/v1/meetings/{meeting_id}/real-time-participants - Get]

	   查询当前会中成员列表，仅包括会中的成员，如果已离会，则不展示
	   企业超级管理员、会议创建者、会议主持人、会议联席主持人可以查询该数据。

	*/
	V1MeetingsMeetingIdRealTimeParticipantsGet(ctx context.Context, request *ApiV1MeetingsMeetingIdRealTimeParticipantsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdRealTimeParticipantsGetResponse, err error)

	/*
	   V1MeetingsMeetingIdVirtualBackgroundPost 设置会议统一虚拟背景[/v1/meetings/{meeting_id}/virtual-background - Post]

	   非进行中非已结束的会议，会议创建者可以设置统一虚拟背景，并设置生效范围。如果企业未开启虚拟背景开关，则该企业下会议不可进行该设置。异步方式上传。支持云会议和Webinar会议，其中Webinar会议设置为对嘉宾生效，且不能指定成员

	*/
	V1MeetingsMeetingIdVirtualBackgroundPost(ctx context.Context, request *ApiV1MeetingsMeetingIdVirtualBackgroundPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdVirtualBackgroundPostResponse, err error)

	/*
	   V1MeetingsMeetingIdWaitingRoomParticipantsGet 获取实时等候室成员列表[/v1/meetings/{meeting_id}/waiting-room-participants - Get]

	   **描述**：

	   * 会议拥有者获取某指定会议的等候室成员列表，需开启等候室且为“会议进行中”状态。
	   * 只有会议的拥有者即创建者可以查询等候室成员列表，其他用户的调用会被拒绝。如果会议非进行中，调用此接口查询会返回空列表。
	   * 企业 secret 鉴权用户（会议创建者）可获取任何该企业该用户创建的会议中的等候室成员列表，OAuth2.0 鉴权用户（会议创建者）只能获取用户通过 OAuth2.0 鉴权创建的会议中的等候室成员列表。
	   * 此接口暂不支持 MRA 设备作为被操作者的情况。

	*/
	V1MeetingsMeetingIdWaitingRoomParticipantsGet(ctx context.Context, request *ApiV1MeetingsMeetingIdWaitingRoomParticipantsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdWaitingRoomParticipantsGetResponse, err error)

	/*
	   V1MeetingsPost 创建会议[/v1/meetings - Post]

	   快速创建或预定一个会议。

	   企业 secret 鉴权用户可创建该用户所属企业下的会议，OAuth2.0 鉴权用户只能创建该企业下 OAuth2.0 应用的会议。
	   用户必须是注册用户，请求头部 X-TC-Registered 字段必须传入为1。
	   当您想实时监测会议创建状况时，您可以通过订阅 [会议创建](https://cloud.tencent.com/document/product/1095/51614) 的事件，接收事件通知。
	   本接口的邀请参会成员限制调整至300人。
	   当会议为周期性会议时，主持人密钥每场会议固定，但单场会议只能获取一次。支持创建周期性会议的主持人密钥。

	*/
	V1MeetingsPost(ctx context.Context, request *ApiV1MeetingsPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsPostResponse, err error)

	/*
	   V1MeetingsQueryMeetingidForDevicePost 查询用户设备是否入会[/v1/meetings/query/meetingid-for-device - Post]

	   查询用户设备是否入会接口，用来查询本企业用户在当前时间是否有设备进入指定的会议中。
	   不支持OAuth2.0鉴权方式访问。

	*/
	V1MeetingsQueryMeetingidForDevicePost(ctx context.Context, request *ApiV1MeetingsQueryMeetingidForDevicePostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsQueryMeetingidForDevicePostResponse, err error)

	/*
	   V1PmiMeetingsGet 查询个人会议号会议列表[/v1/pmi-meetings - Get]

	   查询个人会议号（PMI）会议的会议列表（待开始、进行中、已结束），目前暂不支持 OAuth2.0 鉴权访问。

	*/
	V1PmiMeetingsGet(ctx context.Context, request *ApiV1PmiMeetingsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1PmiMeetingsGetResponse, err error)
}

type meetingsAPIService struct {
	config *core.Config
}

func NewService(config *core.Config) Service {
	return &meetingsAPIService{
		config: config,
	}
}

type ApiV1AsrConfigPutRequest struct {
	Body *V1AsrConfigPutRequest `json:"body,omitempty"`
}

type ApiV1AsrConfigPutResponse struct {
	*xhttp.ApiResponse
	Data *V1AsrConfigPut200Response `json:"data,omitempty"`
}

/*
V1AsrConfigPut 设置语音识别热词[/v1/asr/config - Put]

用户可以通过接口进行您创建的会议的语音识别设置。
权限点：查看或管理您的会议
*/
func (s *meetingsAPIService) V1AsrConfigPut(ctx context.Context, request *ApiV1AsrConfigPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1AsrConfigPutResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/asr/config",
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

	response = &ApiV1AsrConfigPutResponse{
		ApiResponse: apiRsp,
		Data:        new(V1AsrConfigPut200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1AsrDetailsGetRequest struct {
	// 操作者 ID 的类型： 1：userid  2：openid
	OperatorIdType *string `json:"-"`
	// 操作者 ID。operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。
	OperatorId *string `json:"-"`
	// 可以查询某次会议的数据
	MeetingId *string `json:"-"`
	// 查询起始时间戳，UNIX 时间戳（单位秒）。
	StartTime *string `json:"-"`
	// 查询结束时间戳，UNIX 时间戳（单位秒）。
	EndTime *string `json:"-"`
	// 导出文件类型,默认导出txt  0：txt  1：word  2：pdf
	FileType *string `json:"-"`
	// 页码，默认1
	Page *string `json:"-"`
	// 分页大小，默认10，最大50
	PageSize *string `json:"-"`
	// 是否展示双语，默认不展示双语
	ShowBilingual *string                 `json:"-"`
	Body          *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1AsrDetailsGetResponse struct {
	*xhttp.ApiResponse
	Data *V1AsrDetailsGet200Response `json:"data,omitempty"`
}

/*
V1AsrDetailsGet 导出实时转写记录[/v1/asr/details - Get]

如果会议开启了会议转写，可以导出转写记录。主持人可以设置导出权限，默认主持人可以导出，支持会中和会后导出。
*/
func (s *meetingsAPIService) V1AsrDetailsGet(ctx context.Context, request *ApiV1AsrDetailsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1AsrDetailsGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/asr/details",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	if request.OperatorIdType == nil {
		return nil, fmt.Errorf("operator_id_type is required and must be specified")
	}

	if request.OperatorId == nil {
		return nil, fmt.Errorf("operator_id is required and must be specified")
	}

	if request.MeetingId == nil {
		return nil, fmt.Errorf("meeting_id is required and must be specified")
	}

	// path 参数
	// query 参数
	if request.OperatorIdType != nil {
		apiReq.QueryParams.Set("operator_id_type", core.QueryValue(request.OperatorIdType))
	}
	if request.OperatorId != nil {
		apiReq.QueryParams.Set("operator_id", core.QueryValue(request.OperatorId))
	}
	if request.MeetingId != nil {
		apiReq.QueryParams.Set("meeting_id", core.QueryValue(request.MeetingId))
	}
	if request.StartTime != nil {
		apiReq.QueryParams.Set("start_time", core.QueryValue(request.StartTime))
	}
	if request.EndTime != nil {
		apiReq.QueryParams.Set("end_time", core.QueryValue(request.EndTime))
	}
	if request.FileType != nil {
		apiReq.QueryParams.Set("file_type", core.QueryValue(request.FileType))
	}
	if request.Page != nil {
		apiReq.QueryParams.Set("page", core.QueryValue(request.Page))
	}
	if request.PageSize != nil {
		apiReq.QueryParams.Set("page_size", core.QueryValue(request.PageSize))
	}
	if request.ShowBilingual != nil {
		apiReq.QueryParams.Set("show_bilingual", core.QueryValue(request.ShowBilingual))
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

	response = &ApiV1AsrDetailsGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1AsrDetailsGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1AsrPushStatusPostRequest struct {
	Body *V1AsrPushStatusPostRequest `json:"body,omitempty"`
}

type ApiV1AsrPushStatusPostResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1AsrPushStatusPost 开启/关闭实时转写推送[/v1/asr/push-status - Post]

	会议创建者可开启本场会议的实时转写内容推送，待开始的会议或未打开实时转写功能的会议也支持开启推送，开启推送后如果会议打开转写，则可通过webhook 实时转写推送 实时获取到转写内容。

企业 secret 鉴权用户可开启该用户所属企业下的所有会议转写推送，OAuth2.0 鉴权用户只能开启通过 OAuth2.0 鉴权创建的会议转写推送。
企业级自建应用通过 webhook 可以接收到企业下所有开启推送的会议的转写内容，应用级自建应用通过 webhook 可以接收到本应用创建的会议的转写内容。OAuth2.0 应用通过 webhook 可以接收到 OAuth2.0 鉴权创建的会议的转写内容。
*/
func (s *meetingsAPIService) V1AsrPushStatusPost(ctx context.Context, request *ApiV1AsrPushStatusPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1AsrPushStatusPostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/asr/push-status",
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

	response = &ApiV1AsrPushStatusPostResponse{
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

type ApiV1HistoryMeetingsUseridGetRequest struct {
	Userid string `json:"-"`
	// 当前页大小，最小值为1，最大值为20。
	PageSize *string `json:"-"`
	// 当前页，从1开始。
	Page *string `json:"-"`
	// 会议号。
	MeetingCode *string `json:"-"`
	// 搜索的开始时间（预定会议的开始时间），单位秒。
	StartTime *string `json:"-"`
	// 搜索的结束时间（预定会议的开始时间），单位秒。
	EndTime *string                 `json:"-"`
	Body    *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1HistoryMeetingsUseridGetResponse struct {
	*xhttp.ApiResponse
	Data *V1HistoryMeetingsUseridGet200Response `json:"data,omitempty"`
}

/*
V1HistoryMeetingsUseridGet 查询用户已结束会议列表[/v1/history/meetings/{userid} - Get]

获取某指定用户的已结束的会议列表，可返回用户创建与参加过的已结束会议列表，支持 OAuth2.0 鉴权和企微鉴权。
请优先使用operator_id和operator_id_type查询，当使用operator_id和operator_id_type时，userid设置为operator_id的值即可
*/
func (s *meetingsAPIService) V1HistoryMeetingsUseridGet(ctx context.Context, request *ApiV1HistoryMeetingsUseridGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1HistoryMeetingsUseridGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/history/meetings/{userid}",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	if request.PageSize == nil {
		return nil, fmt.Errorf("page_size is required and must be specified")
	}

	if request.Page == nil {
		return nil, fmt.Errorf("page is required and must be specified")
	}

	// path 参数
	apiReq.PathParams.Set("userid", core.PathValue(request.Userid))
	// query 参数
	if request.PageSize != nil {
		apiReq.QueryParams.Set("page_size", core.QueryValue(request.PageSize))
	}
	if request.Page != nil {
		apiReq.QueryParams.Set("page", core.QueryValue(request.Page))
	}
	if request.MeetingCode != nil {
		apiReq.QueryParams.Set("meeting_code", core.QueryValue(request.MeetingCode))
	}
	if request.StartTime != nil {
		apiReq.QueryParams.Set("start_time", core.QueryValue(request.StartTime))
	}
	if request.EndTime != nil {
		apiReq.QueryParams.Set("end_time", core.QueryValue(request.EndTime))
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

	response = &ApiV1HistoryMeetingsUseridGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1HistoryMeetingsUseridGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingJobIdGetRequest struct {
	// 任务ID
	JobId string `json:"-"`
	// 操作者 ID 的类型： 1：userid
	OperatorIdType *string `json:"-"`
	// 操作者 ID。operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型
	OperatorId *string                 `json:"-"`
	Body       *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1MeetingJobIdGetResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingJobIdGet200Response `json:"data,omitempty"`
}

/*
V1MeetingJobIdGet 获取导出 PSTN 通话详单任务结果[/v1/meeting/{job_id} - Get]

获取异步导出任务的结果。
*/
func (s *meetingsAPIService) V1MeetingJobIdGet(ctx context.Context, request *ApiV1MeetingJobIdGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingJobIdGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meeting/{job_id}",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	if request.OperatorIdType == nil {
		return nil, fmt.Errorf("operator_id_type is required and must be specified")
	}

	if request.OperatorId == nil {
		return nil, fmt.Errorf("operator_id is required and must be specified")
	}

	// path 参数
	apiReq.PathParams.Set("job_id", core.PathValue(request.JobId))
	// query 参数
	if request.OperatorIdType != nil {
		apiReq.QueryParams.Set("operator_id_type", core.QueryValue(request.OperatorIdType))
	}
	if request.OperatorId != nil {
		apiReq.QueryParams.Set("operator_id", core.QueryValue(request.OperatorId))
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

	response = &ApiV1MeetingJobIdGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingJobIdGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingMeetingIdWaitingRoomGetRequest struct {
	// 会议的唯一ID
	MeetingId string `json:"-"`
	// 操作者 ID。会议创建者、主持人、联席主持人可以调用该接口。 operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。
	OperatorId *string `json:"-"`
	// 操作者 ID 的类型： 1：企业内用户 userid。 2：企微或者oauth open_id
	OperatorIdType *string `json:"-"`
	// 当前页，页码起始值为1。page最大值为2000
	Page *string `json:"-"`
	// 每页数据条数。默认20，最大50
	PageSize *string                 `json:"-"`
	Body     *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1MeetingMeetingIdWaitingRoomGetResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingMeetingIdWaitingRoomGet200Response `json:"data,omitempty"`
}

/*
V1MeetingMeetingIdWaitingRoomGet 查询等候室成员记录[/v1/meeting/{meeting_id}/waiting-room - Get]

会议创建者、主持人、联席主持人可以查询等候室成员列表，包括等候室内所有用户的进出记录。会前、会中、会后都可以查询。
“查询等候室成员列表”改为“获取实时等候室成员列表”，只有当前等候室的成员。如果是PMI会议，返回的是PMI的meeting_code。
鉴权方式：JWT鉴权、OAuth鉴权、SDK鉴权
OAuth鉴权的权限为：查询会议、查询和管理会议
*/
func (s *meetingsAPIService) V1MeetingMeetingIdWaitingRoomGet(ctx context.Context, request *ApiV1MeetingMeetingIdWaitingRoomGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingMeetingIdWaitingRoomGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meeting/{meeting_id}/waiting-room",
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

	if request.Page == nil {
		return nil, fmt.Errorf("page is required and must be specified")
	}

	if request.PageSize == nil {
		return nil, fmt.Errorf("page_size is required and must be specified")
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

	response = &ApiV1MeetingMeetingIdWaitingRoomGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingMeetingIdWaitingRoomGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingMeetingIdWaitingRoomWelcomeMessageGetRequest struct {
	MeetingId      string                  `json:"-"`
	OperatorId     *string                 `json:"-"`
	OperatorIdType *string                 `json:"-"`
	Body           *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1MeetingMeetingIdWaitingRoomWelcomeMessageGetResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingSetWaitingRoomWelcomeMessagePost200Response `json:"data,omitempty"`
}

/*
V1MeetingMeetingIdWaitingRoomWelcomeMessageGet 获取等候室欢迎语[/v1/meeting/{meeting_id}/waiting-room-welcome-message - Get]

查询会议的等候室欢迎语。当有用户进入等候室时，会收到来自会议主办方的私聊消息引导。
鉴权方式: JWT鉴权、OAuth鉴权
*/
func (s *meetingsAPIService) V1MeetingMeetingIdWaitingRoomWelcomeMessageGet(ctx context.Context, request *ApiV1MeetingMeetingIdWaitingRoomWelcomeMessageGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingMeetingIdWaitingRoomWelcomeMessageGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meeting/{meeting_id}/waiting-room-welcome-message",
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

	response = &ApiV1MeetingMeetingIdWaitingRoomWelcomeMessageGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingSetWaitingRoomWelcomeMessagePost200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingSetWaitingRoomWelcomeMessagePostRequest struct {
	Body *V1MeetingSetWaitingRoomWelcomeMessagePostRequest `json:"body,omitempty"`
}

type ApiV1MeetingSetWaitingRoomWelcomeMessagePostResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingSetWaitingRoomWelcomeMessagePost200Response `json:"data,omitempty"`
}

/*
V1MeetingSetWaitingRoomWelcomeMessagePost 设置等候室欢迎语[/v1/meeting/set-waiting-room-welcome-message - Post]

为已开启等候室的会议配置等候室欢迎语。当有用户进入等候室时，会收到来自会议主办方的私聊消息引导。

鉴权方式:
JWT鉴权、OAuth鉴权
*/
func (s *meetingsAPIService) V1MeetingSetWaitingRoomWelcomeMessagePost(ctx context.Context, request *ApiV1MeetingSetWaitingRoomWelcomeMessagePostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingSetWaitingRoomWelcomeMessagePostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meeting/set-waiting-room-welcome-message",
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

	response = &ApiV1MeetingSetWaitingRoomWelcomeMessagePostResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingSetWaitingRoomWelcomeMessagePost200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingsCustomerShortUrlPostRequest struct {
	Body *V1MeetingsCustomerShortUrlPostRequest `json:"body,omitempty"`
}

type ApiV1MeetingsCustomerShortUrlPostResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingsCustomerShortUrlPost200Response `json:"data,omitempty"`
}

/*
V1MeetingsCustomerShortUrlPost 创建用户专属参会链接[/v1/meetings/customer-short-url - Post]

使用该接口，可用 `customer_data` 进行区分，为一场会议生成多个会议链接。通过用户入会、用户进入等候室等事件，或通过获取等候室成员列表的 API 查询到该参数。
该接口不支持个人会议号会议、网络研讨会（Webinar）。支持企业品牌化链接。
参会者腾讯会议客户端版本需大于等于 3.2.0。
暂不支持 OAuth 2.0 鉴权方式访问。
*/
func (s *meetingsAPIService) V1MeetingsCustomerShortUrlPost(ctx context.Context, request *ApiV1MeetingsCustomerShortUrlPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsCustomerShortUrlPostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings/customer-short-url",
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

	response = &ApiV1MeetingsCustomerShortUrlPostResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingsCustomerShortUrlPost200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingsGetRequest struct {
	// 用户的终端设备类型： 0：PSTN 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：Linux 20：Rooms for Touch Windows 21：Rooms for Touch MacOS 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch iOS
	Instanceid *string `json:"-"`
	// 操作者 ID。operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。 说明：userid 字段和 operator_id 字段二者必填一项。若两者都填，以 operator_id 字段为准。
	OperatorId *string `json:"-"`
	// 操作者 ID 的类型： 3：rooms_id 说明：当前仅支持 rooms_id。如操作者为企业内 userid 或 openId，请使用 userid 字段。
	OperatorIdType *string `json:"-"`
	// 调用方用于标示用户的唯一 ID（企业内部请使用企业唯一用户标识；OAuth2.0 鉴权用户请使用 openId）。 企业唯一用户标识说明： 1：企业对接 SSO 时使用的员工唯一标识 ID。 2：企业调用创建用户接口时传递的 userid 参数。
	Userid *string `json:"-"`
	// 有效的9位数字会议号码。（通过会议CODE查询会议详情时，必传）
	MeetingCode *string `json:"-"`
	// 分页游标
	Cursory *string `json:"-"`
	// 分页获取用户会议列表的查询起始时间值，unix 秒级时间戳
	Pos *string `json:"-"`
	// 是否显示周期性会议的所有子会议，默认值为0。 0：只显示周期性会议的第一个子会议 1：显示所有周期性会议的子会议
	IsShowAllSubMeetings *string                 `json:"-"`
	Body                 *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1MeetingsGetResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingsGet200Response `json:"data,omitempty"`
}

/*
V1MeetingsGet 查询用户的会议列表[/v1/meetings - Get]

通过会议CODE查询会议详情/查询用户的会议列表
① 通过会议CODE查询会议详情
企业 secret 鉴权用户可查询到任何该用户创建的企业下的会议，OAuth2.0 鉴权用户只能查询到通过 OAuth2.0 鉴权创建的会议。
支持企业管理员查询企业下的会议。
本接口的邀请参会成员限制调整至300人。
当会议为周期性会议时，主持人密钥每场会议固定，但单场会议只能获取一次。支持查询周期性会议的主持人密钥。
支持查询 MRA 当前所在会议信息。
若会议号被回收则无法通过 Code 查询，您可以通过会议 ID 查询到该会议。
② 查询用户的会议列表
获取某指定用户的进行中或待开始的会议列表。企业 secret 鉴权用户可查询任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能查询通过 OAuth2.0 鉴权创建的有效会议。
*/
func (s *meetingsAPIService) V1MeetingsGet(ctx context.Context, request *ApiV1MeetingsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	if request.Instanceid == nil {
		return nil, fmt.Errorf("instanceid is required and must be specified")
	}

	// path 参数
	// query 参数
	if request.OperatorId != nil {
		apiReq.QueryParams.Set("operator_id", core.QueryValue(request.OperatorId))
	}
	if request.OperatorIdType != nil {
		apiReq.QueryParams.Set("operator_id_type", core.QueryValue(request.OperatorIdType))
	}
	if request.Userid != nil {
		apiReq.QueryParams.Set("userid", core.QueryValue(request.Userid))
	}
	if request.Instanceid != nil {
		apiReq.QueryParams.Set("instanceid", core.QueryValue(request.Instanceid))
	}
	if request.MeetingCode != nil {
		apiReq.QueryParams.Set("meeting_code", core.QueryValue(request.MeetingCode))
	}
	if request.Cursory != nil {
		apiReq.QueryParams.Set("cursory", core.QueryValue(request.Cursory))
	}
	if request.Pos != nil {
		apiReq.QueryParams.Set("pos", core.QueryValue(request.Pos))
	}
	if request.IsShowAllSubMeetings != nil {
		apiReq.QueryParams.Set("is_show_all_sub_meetings", core.QueryValue(request.IsShowAllSubMeetings))
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

	response = &ApiV1MeetingsGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingsGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingsMeetingIdCancelPostRequest struct {
	MeetingId string                                `json:"-"`
	Body      *V1MeetingsMeetingIdCancelPostRequest `json:"body,omitempty"`
}

type ApiV1MeetingsMeetingIdCancelPostResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1MeetingsMeetingIdCancelPost 取消会议[/v1/meetings/{meeting_id}/cancel - Post]

取消用户创建的会议。用户只能取消自己创建的会议，且该会议是一个有效的会议。如果不是会议创建者或者无效会议号将会返回错误。
企业 secret 鉴权用户可取消任何该用户企业下创建的有效会议，OAuth2.0 鉴权用户只能取消通过 OAuth2.0 鉴权创建的有效会议。
当您想实时监测会议取消状况时，您可以通过订阅 [会议取消](https://cloud.tencent.com/document/product/1095/51616) 的事件，接收事件通知。
*/
func (s *meetingsAPIService) V1MeetingsMeetingIdCancelPost(ctx context.Context, request *ApiV1MeetingsMeetingIdCancelPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdCancelPostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings/{meeting_id}/cancel",
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

	response = &ApiV1MeetingsMeetingIdCancelPostResponse{
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

type ApiV1MeetingsMeetingIdCustomerShortUrlGetRequest struct {
	MeetingId string                  `json:"-"`
	Body      *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1MeetingsMeetingIdCustomerShortUrlGetResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingsMeetingIdCustomerShortUrlGet200Response `json:"data,omitempty"`
}

/*
V1MeetingsMeetingIdCustomerShortUrlGet 获取用户专属参会链接[/v1/meetings/{meeting_id}/customer-short-url - Get]

**描述**：

* 可以获取指定会议的所有专属参会链接及 `customer_data`。
* 该接口不支持个人会议号会议、网络研讨会（Webinar）。支持企业品牌化链接。
* 参会者腾讯会议客户端版本需大于等于 3.2.0。
* 暂不支持 OAuth 2.0 鉴权方式访问。
*/
func (s *meetingsAPIService) V1MeetingsMeetingIdCustomerShortUrlGet(ctx context.Context, request *ApiV1MeetingsMeetingIdCustomerShortUrlGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdCustomerShortUrlGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings/{meeting_id}/customer-short-url",
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

	response = &ApiV1MeetingsMeetingIdCustomerShortUrlGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingsMeetingIdCustomerShortUrlGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingsMeetingIdEnrollApprovalsGetRequest struct {
	// 会议ID
	MeetingId string `json:"-"`
	// 操作者 ID。会议创建者可以导入报名信息。 operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。  operator_id_type=2，operator_id必须和公共参数的openid一致。  operator_id和userid至少填写一个，两个参数如果都传了以operator_id为准。  使用OAuth公参鉴权后不能使用userid为入参。
	OperatorId *string `json:"-"`
	// 操作者 ID 的类型：  1: userid 2: open_id  如果operator_id和userid具有值，则以operator_id为准；
	OperatorIdType *string `json:"-"`
	// 用户的终端设备类型： 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序
	Instanceid *string `json:"-"`
	// 当前页，页码起始值为1
	Page *string `json:"-"`
	// 分页大小，最大50条
	PageSize *string `json:"-"`
	// 审批状态筛选字段，审批状态：0 全部，1 待审批，2 已拒绝，3 已批准，默认返回全部
	Status *string                 `json:"-"`
	Body   *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1MeetingsMeetingIdEnrollApprovalsGetResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingsMeetingIdEnrollApprovalsGet200Response `json:"data,omitempty"`
}

/*
V1MeetingsMeetingIdEnrollApprovalsGet 查询会议报名信息[/v1/meetings/{meeting_id}/enroll/approvals - Get]

查询已报名观众数量和报名观众答题详情，仅会议创建者可查询。
企业 secret 鉴权用户可修改任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能修改通过 OAuth2.0 鉴权创建的有效会议。
用户必须是注册用户，请求头部 X-TC-Registered 字段必须传入为1。
*/
func (s *meetingsAPIService) V1MeetingsMeetingIdEnrollApprovalsGet(ctx context.Context, request *ApiV1MeetingsMeetingIdEnrollApprovalsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdEnrollApprovalsGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings/{meeting_id}/enroll/approvals",
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

	if request.Instanceid == nil {
		return nil, fmt.Errorf("instanceid is required and must be specified")
	}

	if request.Page == nil {
		return nil, fmt.Errorf("page is required and must be specified")
	}

	if request.PageSize == nil {
		return nil, fmt.Errorf("page_size is required and must be specified")
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
	if request.Instanceid != nil {
		apiReq.QueryParams.Set("instanceid", core.QueryValue(request.Instanceid))
	}
	if request.Page != nil {
		apiReq.QueryParams.Set("page", core.QueryValue(request.Page))
	}
	if request.PageSize != nil {
		apiReq.QueryParams.Set("page_size", core.QueryValue(request.PageSize))
	}
	if request.Status != nil {
		apiReq.QueryParams.Set("status", core.QueryValue(request.Status))
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

	response = &ApiV1MeetingsMeetingIdEnrollApprovalsGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingsMeetingIdEnrollApprovalsGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingsMeetingIdEnrollApprovalsPutRequest struct {
	// 会议ID
	MeetingId string                                        `json:"-"`
	Body      *V1MeetingsMeetingIdEnrollApprovalsPutRequest `json:"body,omitempty"`
}

type ApiV1MeetingsMeetingIdEnrollApprovalsPutResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingsMeetingIdEnrollApprovalsPut200Response `json:"data,omitempty"`
}

/*
V1MeetingsMeetingIdEnrollApprovalsPut 审批会议报名信息[/v1/meetings/{meeting_id}/enroll/approvals - Put]

批量云会议的报名信息，仅会议创建者可审批。
企业 secret 鉴权用户可审批任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能审批通过 OAuth2.0 鉴权创建的有效会议。
用户必须是注册用户，请求头部 X-TC-Registered 字段必须传入为1。
*/
func (s *meetingsAPIService) V1MeetingsMeetingIdEnrollApprovalsPut(ctx context.Context, request *ApiV1MeetingsMeetingIdEnrollApprovalsPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdEnrollApprovalsPutResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings/{meeting_id}/enroll/approvals",
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

	response = &ApiV1MeetingsMeetingIdEnrollApprovalsPutResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingsMeetingIdEnrollApprovalsPut200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingsMeetingIdEnrollConfigGetRequest struct {
	// 会议ID
	MeetingId string `json:"-"`
	// 操作者 ID。会议创建者可以导入报名信息。 operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。  operator_id_type=2，operator_id必须和公共参数的openid一致。  operator_id和userid至少填写一个，两个参数如果都传了以operator_id为准。  使用OAuth公参鉴权后不能使用userid为入参。
	OperatorId *string `json:"-"`
	// 操作者 ID 的类型：  1: userid 2: open_id  如果operator_id和userid具有值，则以operator_id为准；
	OperatorIdType *string `json:"-"`
	// 用户的终端设备类型： 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序
	Instanceid *string                 `json:"-"`
	Body       *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1MeetingsMeetingIdEnrollConfigGetResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingsMeetingIdEnrollConfigGet200Response `json:"data,omitempty"`
}

/*
V1MeetingsMeetingIdEnrollConfigGet 查询会议报名配置[/v1/meetings/{meeting_id}/enroll/config - Get]

查询云会议的报名配置和报名问题，仅会议创建者可查询。会议未开启报名时会返回未开启报名错误。
企业 secret 鉴权用户可查询任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能查询通过 OAuth2.0 鉴权创建的有效会议。
用户必须是注册用户，请求头部 X-TC-Registered 字段必须传入为1。
*/
func (s *meetingsAPIService) V1MeetingsMeetingIdEnrollConfigGet(ctx context.Context, request *ApiV1MeetingsMeetingIdEnrollConfigGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdEnrollConfigGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings/{meeting_id}/enroll/config",
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

	if request.Instanceid == nil {
		return nil, fmt.Errorf("instanceid is required and must be specified")
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

	response = &ApiV1MeetingsMeetingIdEnrollConfigGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingsMeetingIdEnrollConfigGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingsMeetingIdEnrollConfigPutRequest struct {
	// 会议ID
	MeetingId string                                     `json:"-"`
	Body      *V1MeetingsMeetingIdEnrollConfigPutRequest `json:"body,omitempty"`
}

type ApiV1MeetingsMeetingIdEnrollConfigPutResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingsMeetingIdEnrollConfigPut200Response `json:"data,omitempty"`
}

/*
V1MeetingsMeetingIdEnrollConfigPut 修改会议报名配置[/v1/meetings/{meeting_id}/enroll/config - Put]

修改云会议的报名配置和报名问题，仅会议创建者可修改，且需要会议已开启报名。
企业 secret 鉴权用户可修改任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能修改通过 OAuth2.0 鉴权创建的有效会议。
用户必须是注册用户，请求头部 X-TC-Registered 字段必须传入为1。
*/
func (s *meetingsAPIService) V1MeetingsMeetingIdEnrollConfigPut(ctx context.Context, request *ApiV1MeetingsMeetingIdEnrollConfigPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdEnrollConfigPutResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings/{meeting_id}/enroll/config",
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

	response = &ApiV1MeetingsMeetingIdEnrollConfigPutResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingsMeetingIdEnrollConfigPut200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingsMeetingIdEnrollIdsPostRequest struct {
	// 会议ID
	MeetingId string                                   `json:"-"`
	Body      *V1MeetingsMeetingIdEnrollIdsPostRequest `json:"body,omitempty"`
}

type ApiV1MeetingsMeetingIdEnrollIdsPostResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingsMeetingIdEnrollIdsPost200Response `json:"data,omitempty"`
}

/*
V1MeetingsMeetingIdEnrollIdsPost 查询会议成员报名 ID[/v1/meetings/{meeting_id}/enroll/ids - Post]

描述： 支持查询会议中已报名成员的报名 ID，仅会议创建者可查询。
*/
func (s *meetingsAPIService) V1MeetingsMeetingIdEnrollIdsPost(ctx context.Context, request *ApiV1MeetingsMeetingIdEnrollIdsPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdEnrollIdsPostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings/{meeting_id}/enroll/ids",
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

	response = &ApiV1MeetingsMeetingIdEnrollIdsPostResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingsMeetingIdEnrollIdsPost200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingsMeetingIdEnrollImportPostRequest struct {
	// 会议id
	MeetingId string                                      `json:"-"`
	Body      *V1MeetingsMeetingIdEnrollImportPostRequest `json:"body,omitempty"`
}

type ApiV1MeetingsMeetingIdEnrollImportPostResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingsMeetingIdEnrollImportPost200Response `json:"data,omitempty"`
}

/*
V1MeetingsMeetingIdEnrollImportPost 导入会议报名信息[/v1/meetings/{meeting_id}/enroll/import - Post]

指定会议中导入报名信息。

企业 secret 鉴权用户可通过同企业下用户 userid 和手机号导入报名信息，OAuth2.0 鉴权用户能通过用户 open_id，与应用同企业下的 userid 以及手机号导入报名信息。
用户必须是注册用户，请求头部 X-TC-Registered 字段必须传入为1。
商业版单场会议导入上限1000条，企业版单场会议导入上限4000条。如需提升，请联系我们。
*/
func (s *meetingsAPIService) V1MeetingsMeetingIdEnrollImportPost(ctx context.Context, request *ApiV1MeetingsMeetingIdEnrollImportPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdEnrollImportPostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings/{meeting_id}/enroll/import",
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

	response = &ApiV1MeetingsMeetingIdEnrollImportPostResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingsMeetingIdEnrollImportPost200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingsMeetingIdEnrollUnregistrationDeleteRequest struct {
	MeetingId string                                                `json:"-"`
	Body      *V1MeetingsMeetingIdEnrollUnregistrationDeleteRequest `json:"body,omitempty"`
}

type ApiV1MeetingsMeetingIdEnrollUnregistrationDeleteResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingsMeetingIdEnrollUnregistrationDelete200Response `json:"data,omitempty"`
}

/*
V1MeetingsMeetingIdEnrollUnregistrationDelete 删除会议报名信息[/v1/meetings/{meeting_id}/enroll/unregistration - Delete]

描述：
删除指定会议的报名信息，支持删除用户手动报名的信息和导入的报名信息。
企业 secret 鉴权用户可删除该用户企业会议下的报名信息，OAuth2.0 鉴权用户只能删除通过 OAuth2.0 鉴权创建的有效会议的报名信息。
用户必须是注册用户，请求头部 X-TC-Registered 字段必须传入为1。
*/
func (s *meetingsAPIService) V1MeetingsMeetingIdEnrollUnregistrationDelete(ctx context.Context, request *ApiV1MeetingsMeetingIdEnrollUnregistrationDeleteRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdEnrollUnregistrationDeleteResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings/{meeting_id}/enroll/unregistration",
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

	response = &ApiV1MeetingsMeetingIdEnrollUnregistrationDeleteResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingsMeetingIdEnrollUnregistrationDelete200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingsMeetingIdGetRequest struct {
	MeetingId string `json:"-"`
	// 用户的终端设备类型： 0：PSTN 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch MacOS 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch iOS
	Instanceid *string `json:"-"`
	// 操作者ID，根据operator_id_type的值，使用不同的类型
	OperatorId *string `json:"-"`
	// 操作者ID的类型：1.userid 2.openid 3.rooms_id
	OperatorIdType *string                 `json:"-"`
	Body           *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1MeetingsMeetingIdGetResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingsMeetingIdGet200Response `json:"data,omitempty"`
}

/*
V1MeetingsMeetingIdGet 查询会议[/v1/meetings/{meeting_id} - Get]

通过会议 ID 查询会议详情。
企业 secret 鉴权用户可查询到任何该用户创建的企业下的会议，OAuth2.0 鉴权用户只能查询到通过 OAuth2.0 鉴权创建的会议。
本接口的邀请参会成员限制调整至300人。
当会议为周期性会议时，主持人密钥每场会议固定，但单场会议只能获取一次。支持查询周期性会议的主持人密钥。
支持查询 MRA 当前所在会议信息。
*/
func (s *meetingsAPIService) V1MeetingsMeetingIdGet(ctx context.Context, request *ApiV1MeetingsMeetingIdGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings/{meeting_id}",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	if request.Instanceid == nil {
		return nil, fmt.Errorf("instanceid is required and must be specified")
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

	response = &ApiV1MeetingsMeetingIdGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingsMeetingIdGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingsMeetingIdInviteesGetRequest struct {
	// 会议ID
	MeetingId string `json:"-"`
	// 会议创建者ID.调用方用于标示用户的唯一 ID（企业内部请使用企业唯一用户标识；OAuth2.0 鉴权用户请使用 openId）。 企业唯一用户标识说明： 1. 企业对接 SSO 时使用的员工唯一标识 ID。 2. 企业调用创建用户接口时传递的 userid 参数。
	Userid *string `json:"-"`
	// 用户的终端设备类型：1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch Mac 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch Iphone
	Instanceid *string `json:"-"`
	// 分页获取受邀成员列表的查询起始位置值。此参数为非必选参数，默认值为0，从头开始查询。（输出参数has_remaining为 true，表示还存在受邀成员需要继续查询；输出参数next_pos即下一次查询的“pos”值。多次调用该接口直到输出参数“has_remaining”为 false。
	Pos  *string                 `json:"-"`
	Body *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1MeetingsMeetingIdInviteesGetResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingsMeetingIdInviteesGet200Response `json:"data,omitempty"`
}

/*
V1MeetingsMeetingIdInviteesGet 获取会议受邀成员列表[/v1/meetings/{meeting_id}/invitees - Get]

根据会议ID获取受邀成员列表，支持分页获取，只有会议的创建者才有权限获取。
*/
func (s *meetingsAPIService) V1MeetingsMeetingIdInviteesGet(ctx context.Context, request *ApiV1MeetingsMeetingIdInviteesGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdInviteesGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings/{meeting_id}/invitees",
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
	if request.Pos != nil {
		apiReq.QueryParams.Set("pos", core.QueryValue(request.Pos))
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

	response = &ApiV1MeetingsMeetingIdInviteesGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingsMeetingIdInviteesGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingsMeetingIdInviteesPutRequest struct {
	// 会议ID
	MeetingId string                                 `json:"-"`
	Body      *V1MeetingsMeetingIdInviteesPutRequest `json:"body,omitempty"`
}

type ApiV1MeetingsMeetingIdInviteesPutResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1MeetingsMeetingIdInviteesPut 设置会议邀请成员[/v1/meetings/{meeting_id}/invitees - Put]

根据会议ID设置邀请成员，只有会议的创建者才有权限设置。
最多可以设置2000名邀请者。
注：本接口为覆盖式设置。
*/
func (s *meetingsAPIService) V1MeetingsMeetingIdInviteesPut(ctx context.Context, request *ApiV1MeetingsMeetingIdInviteesPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdInviteesPutResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings/{meeting_id}/invitees",
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

	response = &ApiV1MeetingsMeetingIdInviteesPutResponse{
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

type ApiV1MeetingsMeetingIdParticipantsGetRequest struct {
	MeetingId string `json:"-"`
	// 周期性会议子会议 ID。说明：可通过查询用户的会议列表、查询会议接口获取返回的子会议 ID，即 current_sub_meeting_id；如果是周期性会议，此参数必传。
	SubMeetingId *string `json:"-"`
	// 操作者ID，根据operator_id_type的值，使用不同的类型
	OperatorId *string `json:"-"`
	// 操作者ID的类型：1.userid 2.open_id 3.rooms_id
	OperatorIdType *string `json:"-"`
	// 会议创建者的用户 ID（企业内部请使用企业唯一用户标识；OAuth2.0 鉴权用户请使用 openId）。
	Userid *string `json:"-"`
	// 分页获取参会成员列表的查询起始位置值。当参会成员较多时，建议使用此参数进行分页查询，避免查询超时。此参数为非必选参数，默认值为0，从头开始查询。设置每页返回的数量，请参考参数“size”的说明。查询返回输出参数“has_remaining”为 true，表示该会议人数较多，还有一定数量的参会成员需要继续查询。返回参数“next_pos”的值即为下一次查询的 pos 的值。多次调用该查询接口直到输出参数“has_remaining”值为 false。
	Pos *string `json:"-"`
	// 拉取参会成员条数，目前每页支持最大100条。
	Size *string `json:"-"`
	// 参会时间过滤起始时间（单位秒）。说明：时间区间不允许超过31天，如果为空默认当前时间前推31天；start_time 和 end_time 都没传时最大查询时间跨度90天；对于周期性会议查询暂时不生效，请使用分页参数查询。
	StartTime *string `json:"-"`
	//  参会时间过滤终止时间（单位秒）。说明：时间区间不允许超过31天，如果为空默认取当前时间；start_time 和 end_time 都没传时最大查询时间跨度90天；对于周期性会议查询暂时不生效，请使用分页参数查询。
	EndTime *string                 `json:"-"`
	Body    *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1MeetingsMeetingIdParticipantsGetResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingsMeetingIdParticipantsGet200Response `json:"data,omitempty"`
}

/*
V1MeetingsMeetingIdParticipantsGet 获取参会成员列表[/v1/meetings/{meetingId}/participants - Get]

会议创建者和企业管理员可以查询参会成员的列表，其他用户的调用会被拒绝。

支持查询网络研讨会参会成员列表。
如果会议还未开始，调用此接口查询会返回空列表。
企业 secret 鉴权用户（会议创建者）可获取任何该企业该用户创建的有效会议中的参会成员，企业 secret 鉴权用户（企业超级管理员）可获取任何该企业下创建的有效会议中的参会成员，OAuth2.0 鉴权用户（会议创建者）只能获取用户通过 OAuth2.0 鉴权创建的有效会议中的参会成员。
当您需要实时监测参会成员入会状态或退会状态时，您可以通过订阅 [用户入会](https://cloud.tencent.com/document/product/1095/51620)和 [用户离会](https://cloud.tencent.com/document/product/1095/51622) 的事件，接收事件通知。
*/
func (s *meetingsAPIService) V1MeetingsMeetingIdParticipantsGet(ctx context.Context, request *ApiV1MeetingsMeetingIdParticipantsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdParticipantsGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings/{meetingId}/participants",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	// path 参数
	apiReq.PathParams.Set("meetingId", core.PathValue(request.MeetingId))
	// query 参数
	if request.SubMeetingId != nil {
		apiReq.QueryParams.Set("sub_meeting_id", core.QueryValue(request.SubMeetingId))
	}
	if request.OperatorId != nil {
		apiReq.QueryParams.Set("operator_id", core.QueryValue(request.OperatorId))
	}
	if request.OperatorIdType != nil {
		apiReq.QueryParams.Set("operator_id_type", core.QueryValue(request.OperatorIdType))
	}
	if request.Userid != nil {
		apiReq.QueryParams.Set("userid", core.QueryValue(request.Userid))
	}
	if request.Pos != nil {
		apiReq.QueryParams.Set("pos", core.QueryValue(request.Pos))
	}
	if request.Size != nil {
		apiReq.QueryParams.Set("size", core.QueryValue(request.Size))
	}
	if request.StartTime != nil {
		apiReq.QueryParams.Set("start_time", core.QueryValue(request.StartTime))
	}
	if request.EndTime != nil {
		apiReq.QueryParams.Set("end_time", core.QueryValue(request.EndTime))
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

	response = &ApiV1MeetingsMeetingIdParticipantsGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingsMeetingIdParticipantsGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingsMeetingIdPutRequest struct {
	MeetingId string                         `json:"-"`
	Body      *V1MeetingsMeetingIdPutRequest `json:"body,omitempty"`
}

type ApiV1MeetingsMeetingIdPutResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingsMeetingIdPut200Response `json:"data,omitempty"`
}

/*
V1MeetingsMeetingIdPut 修改会议[/v1/meetings/{meeting_id} - Put]

修改某指定会议的会议信息。

企业 secret 鉴权用户可修改任何该企业该用户创建的有效会议，OAuth2.0 鉴权用户只能修改通过 OAuth2.0 鉴权创建的有效会议。
当您想实时监测会议修改状况时，您可以通过订阅 [会议更新](https://cloud.tencent.com/document/product/1095/51615) 的事件，接收事件通知。
本接口的邀请参会成员限制调整至300人。
当会议为周期性会议时，主持人密钥每场会议固定，但单场会议只能获取一次。支持修改周期性会议的主持人密钥。
*/
func (s *meetingsAPIService) V1MeetingsMeetingIdPut(ctx context.Context, request *ApiV1MeetingsMeetingIdPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdPutResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings/{meeting_id}",
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

	response = &ApiV1MeetingsMeetingIdPutResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingsMeetingIdPut200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingsMeetingIdQosGetRequest struct {
	// 会议ID
	MeetingId string `json:"-"`
	// 操作者ID
	OperatorId *string `json:"-"`
	// 操作者ID类型
	OperatorIdType *string `json:"-"`
	// 分页大小，20-100
	PageSize *string `json:"-"`
	// 页码
	Page *string `json:"-"`
	// 筛选的用户ID
	ToOperatorId *string `json:"-"`
	// 筛选的用户ID类型 4:ms_open_id
	ToOperatorIdType *string `json:"-"`
	// 搜索key(格式：模块_指标)
	Key *string `json:"-"`
	// 搜索值,搜索大于等于该值的数据
	MinValue *string `json:"-"`
	// 最大搜索值，搜索小于等于该值的数据
	MaxValue *string                 `json:"-"`
	Body     *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1MeetingsMeetingIdQosGetResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingsMeetingIdQosGet200Response `json:"data,omitempty"`
}

/*
V1MeetingsMeetingIdQosGet 获取会议实时质量检测数据[/v1/meetings/{meeting_id}/qos - Get]

拥有企业“会议列表--会控”权限的成员，能够获取实时会议质量检测数据。
支持云会议和Webinar会议的数据。会议状态为进行中。
*/
func (s *meetingsAPIService) V1MeetingsMeetingIdQosGet(ctx context.Context, request *ApiV1MeetingsMeetingIdQosGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdQosGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings/{meeting_id}/qos",
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
	if request.PageSize != nil {
		apiReq.QueryParams.Set("page_size", core.QueryValue(request.PageSize))
	}
	if request.Page != nil {
		apiReq.QueryParams.Set("page", core.QueryValue(request.Page))
	}
	if request.ToOperatorId != nil {
		apiReq.QueryParams.Set("to_operator_id", core.QueryValue(request.ToOperatorId))
	}
	if request.ToOperatorIdType != nil {
		apiReq.QueryParams.Set("to_operator_id_type", core.QueryValue(request.ToOperatorIdType))
	}
	if request.Key != nil {
		apiReq.QueryParams.Set("key", core.QueryValue(request.Key))
	}
	if request.MinValue != nil {
		apiReq.QueryParams.Set("min_value", core.QueryValue(request.MinValue))
	}
	if request.MaxValue != nil {
		apiReq.QueryParams.Set("max_value", core.QueryValue(request.MaxValue))
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

	response = &ApiV1MeetingsMeetingIdQosGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingsMeetingIdQosGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingsMeetingIdQualityGetRequest struct {
	// 会议唯一ID
	MeetingId string `json:"-"`
	// 操作者 ID。operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。角色校验：付费开通该服务的企业管理员/超管可以查询健康度。
	OperatorId *string `json:"-"`
	// 操作者 ID 的类型： 1：企业内用户 userid。 2: open_id 4: ms_open_id
	OperatorIdType *string `json:"-"`
	// 分页大小，最小1，最大50。
	PageSize *string `json:"-"`
	// 当前页，页码起始值，最小1，最大2000。
	Page *string `json:"-"`
	// 参会时间过滤起始时间，UNIX 时间戳（单位秒），可查询的时间区间为过去7天到现在。 返回meeting_id对应会议房间下，开始时间大于等于start_time且离start_time最近的一个媒体房间数据（从第一个人入会到会中成员全部离开会议形成一个媒体房间，若同一会议号下再次有人入会则形成新的媒体房间） 如果同一会议号下有多个媒体房间，请先使用“获取账户级已结束会议列表”接口查询，获知需查询的媒体房间的start_time。
	StartTime *string `json:"-"`
	// 用户的终端设备类型： 0：PSTN 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch MacOS 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch iOS
	Instanceid *string `json:"-"`
	// 周期性会议子会议 ID。说明：可通过查询用户的会议列表、查询会议接口获取返回的子会议 ID，即 current_sub_meeting_id；如果是周期性会议，此参数必传。
	SubMeetingId *string                 `json:"-"`
	Body         *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1MeetingsMeetingIdQualityGetResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingsMeetingIdQualityGet200Response `json:"data,omitempty"`
}

/*
V1MeetingsMeetingIdQualityGet 查询会议健康度[/v1/meetings/{meeting_id}/quality - Get]

查询会议及参会成员的健康度，付费开通该服务的企业管理员、超管可以查询，与是否为会议创建者/主持人/联席主持人无关。
鉴权方式：支持 JWT 鉴权 和 Oauth 鉴权
*/
func (s *meetingsAPIService) V1MeetingsMeetingIdQualityGet(ctx context.Context, request *ApiV1MeetingsMeetingIdQualityGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdQualityGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings/{meeting_id}/quality",
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

	if request.PageSize == nil {
		return nil, fmt.Errorf("page_size is required and must be specified")
	}

	if request.Page == nil {
		return nil, fmt.Errorf("page is required and must be specified")
	}

	if request.StartTime == nil {
		return nil, fmt.Errorf("start_time is required and must be specified")
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
	if request.Instanceid != nil {
		apiReq.QueryParams.Set("instanceid", core.QueryValue(request.Instanceid))
	}
	if request.SubMeetingId != nil {
		apiReq.QueryParams.Set("sub_meeting_id", core.QueryValue(request.SubMeetingId))
	}
	if request.PageSize != nil {
		apiReq.QueryParams.Set("page_size", core.QueryValue(request.PageSize))
	}
	if request.Page != nil {
		apiReq.QueryParams.Set("page", core.QueryValue(request.Page))
	}
	if request.StartTime != nil {
		apiReq.QueryParams.Set("start_time", core.QueryValue(request.StartTime))
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

	response = &ApiV1MeetingsMeetingIdQualityGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingsMeetingIdQualityGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingsMeetingIdRealTimeParticipantsGetRequest struct {
	// 会议唯一ID
	MeetingId string `json:"-"`
	// 操作者 ID。operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。
	OperatorId *string `json:"-"`
	// Integer 操作者 ID 的类型：1：企业内用户 userid。2: open_id3. rooms_id
	OperatorIdType *string `json:"-"`
	// 当前页，页码起始值为1。
	Page *string `json:"-"`
	// 分页大小，最大50条。
	PageSize *string `json:"-"`
	// String 周期性会议子会议 ID。说明：可通过查询用户的会议列表、查询会议接口获取返回的子会议 ID，即 current_sub_meeting_id；如果是周期性会议，此参数必传。
	SubMeetingId *string                 `json:"-"`
	Body         *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1MeetingsMeetingIdRealTimeParticipantsGetResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingsMeetingIdRealTimeParticipantsGet200Response `json:"data,omitempty"`
}

/*
V1MeetingsMeetingIdRealTimeParticipantsGet 查询实时会中成员列表[/v1/meetings/{meeting_id}/real-time-participants - Get]

查询当前会中成员列表，仅包括会中的成员，如果已离会，则不展示
企业超级管理员、会议创建者、会议主持人、会议联席主持人可以查询该数据。
*/
func (s *meetingsAPIService) V1MeetingsMeetingIdRealTimeParticipantsGet(ctx context.Context, request *ApiV1MeetingsMeetingIdRealTimeParticipantsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdRealTimeParticipantsGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings/{meeting_id}/real-time-participants",
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

	if request.Page == nil {
		return nil, fmt.Errorf("page is required and must be specified")
	}

	if request.PageSize == nil {
		return nil, fmt.Errorf("page_size is required and must be specified")
	}

	// path 参数
	apiReq.PathParams.Set("meeting_id", core.PathValue(request.MeetingId))
	// query 参数
	if request.SubMeetingId != nil {
		apiReq.QueryParams.Set("sub_meeting_id", core.QueryValue(request.SubMeetingId))
	}
	if request.OperatorId != nil {
		apiReq.QueryParams.Set("operator_id", core.QueryValue(request.OperatorId))
	}
	if request.OperatorIdType != nil {
		apiReq.QueryParams.Set("operator_id_type", core.QueryValue(request.OperatorIdType))
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

	response = &ApiV1MeetingsMeetingIdRealTimeParticipantsGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingsMeetingIdRealTimeParticipantsGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingsMeetingIdVirtualBackgroundPostRequest struct {
	// 会议ID
	MeetingId string                                           `json:"-"`
	Body      *V1MeetingsMeetingIdVirtualBackgroundPostRequest `json:"body,omitempty"`
}

type ApiV1MeetingsMeetingIdVirtualBackgroundPostResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingsMeetingIdVirtualBackgroundPost200Response `json:"data,omitempty"`
}

/*
V1MeetingsMeetingIdVirtualBackgroundPost 设置会议统一虚拟背景[/v1/meetings/{meeting_id}/virtual-background - Post]

非进行中非已结束的会议，会议创建者可以设置统一虚拟背景，并设置生效范围。如果企业未开启虚拟背景开关，则该企业下会议不可进行该设置。异步方式上传。支持云会议和Webinar会议，其中Webinar会议设置为对嘉宾生效，且不能指定成员
*/
func (s *meetingsAPIService) V1MeetingsMeetingIdVirtualBackgroundPost(ctx context.Context, request *ApiV1MeetingsMeetingIdVirtualBackgroundPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdVirtualBackgroundPostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings/{meeting_id}/virtual-background",
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

	response = &ApiV1MeetingsMeetingIdVirtualBackgroundPostResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingsMeetingIdVirtualBackgroundPost200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingsMeetingIdWaitingRoomParticipantsGetRequest struct {
	MeetingId string `json:"-"`
	// 会议创建者的用户 ID（企业内部请使用企业唯一用户标识；OAuth2.0 鉴权用户请使用 openId）
	Userid *string `json:"-"`
	// 分页大小，默认10，最大50
	PageSize *string `json:"-"`
	// 页码，从1开始
	Page *string                 `json:"-"`
	Body *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1MeetingsMeetingIdWaitingRoomParticipantsGetResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingsMeetingIdWaitingRoomParticipantsGet200Response `json:"data,omitempty"`
}

/*
V1MeetingsMeetingIdWaitingRoomParticipantsGet 获取实时等候室成员列表[/v1/meetings/{meeting_id}/waiting-room-participants - Get]

**描述**：

* 会议拥有者获取某指定会议的等候室成员列表，需开启等候室且为“会议进行中”状态。
* 只有会议的拥有者即创建者可以查询等候室成员列表，其他用户的调用会被拒绝。如果会议非进行中，调用此接口查询会返回空列表。
* 企业 secret 鉴权用户（会议创建者）可获取任何该企业该用户创建的会议中的等候室成员列表，OAuth2.0 鉴权用户（会议创建者）只能获取用户通过 OAuth2.0 鉴权创建的会议中的等候室成员列表。
* 此接口暂不支持 MRA 设备作为被操作者的情况。
*/
func (s *meetingsAPIService) V1MeetingsMeetingIdWaitingRoomParticipantsGet(ctx context.Context, request *ApiV1MeetingsMeetingIdWaitingRoomParticipantsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsMeetingIdWaitingRoomParticipantsGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings/{meeting_id}/waiting-room-participants",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	if request.Userid == nil {
		return nil, fmt.Errorf("userid is required and must be specified")
	}

	// path 参数
	apiReq.PathParams.Set("meeting_id", core.PathValue(request.MeetingId))
	// query 参数
	if request.Userid != nil {
		apiReq.QueryParams.Set("userid", core.QueryValue(request.Userid))
	}
	if request.PageSize != nil {
		apiReq.QueryParams.Set("page_size", core.QueryValue(request.PageSize))
	}
	if request.Page != nil {
		apiReq.QueryParams.Set("page", core.QueryValue(request.Page))
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

	response = &ApiV1MeetingsMeetingIdWaitingRoomParticipantsGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingsMeetingIdWaitingRoomParticipantsGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingsPostRequest struct {
	Body *V1MeetingsPostRequest `json:"body,omitempty"`
}

type ApiV1MeetingsPostResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingsPost200Response `json:"data,omitempty"`
}

/*
V1MeetingsPost 创建会议[/v1/meetings - Post]

快速创建或预定一个会议。

企业 secret 鉴权用户可创建该用户所属企业下的会议，OAuth2.0 鉴权用户只能创建该企业下 OAuth2.0 应用的会议。
用户必须是注册用户，请求头部 X-TC-Registered 字段必须传入为1。
当您想实时监测会议创建状况时，您可以通过订阅 [会议创建](https://cloud.tencent.com/document/product/1095/51614) 的事件，接收事件通知。
本接口的邀请参会成员限制调整至300人。
当会议为周期性会议时，主持人密钥每场会议固定，但单场会议只能获取一次。支持创建周期性会议的主持人密钥。
*/
func (s *meetingsAPIService) V1MeetingsPost(ctx context.Context, request *ApiV1MeetingsPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsPostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings",
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

	response = &ApiV1MeetingsPostResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingsPost200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MeetingsQueryMeetingidForDevicePostRequest struct {
	Body *V1MeetingsQueryMeetingidForDevicePostRequest `json:"body,omitempty"`
}

type ApiV1MeetingsQueryMeetingidForDevicePostResponse struct {
	*xhttp.ApiResponse
	Data *V1MeetingsQueryMeetingidForDevicePost200Response `json:"data,omitempty"`
}

/*
V1MeetingsQueryMeetingidForDevicePost 查询用户设备是否入会[/v1/meetings/query/meetingid-for-device - Post]

查询用户设备是否入会接口，用来查询本企业用户在当前时间是否有设备进入指定的会议中。
不支持OAuth2.0鉴权方式访问。
*/
func (s *meetingsAPIService) V1MeetingsQueryMeetingidForDevicePost(ctx context.Context, request *ApiV1MeetingsQueryMeetingidForDevicePostRequest, opts ...core.RequestOptionFunc) (response *ApiV1MeetingsQueryMeetingidForDevicePostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/meetings/query/meetingid-for-device",
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

	response = &ApiV1MeetingsQueryMeetingidForDevicePostResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MeetingsQueryMeetingidForDevicePost200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1PmiMeetingsGetRequest struct {
	// 企业下操作者ID，根据operator_id_type的值，使用不同的类型
	OperatorId *string `json:"-"`
	// 操作者ID类型： 1.企业用户userid 3.rooms设备rooms_id
	OperatorIdType *string `json:"-"`
	// 用户的终端设备类型： 1：PC 2：Mac 3：Android 4：iOS 5：Web 6：iPad 7：Android Pad 8：小程序 9：voip、sip 设备 10：linux 20：Rooms for Touch Windows 21：Rooms for Touch Mac 22：Rooms for Touch Android 30：Controller for Touch Windows 32：Controller for Touch Android 33：Controller for Touch Iphone
	Instanceid *string `json:"-"`
	// 查询起始时间，时间区间不超过90天
	StartTime *string `json:"-"`
	// 查询结束时间，时间区间不超过90天
	EndTime *string `json:"-"`
	// 当前页，页码起始值为1，默认为1
	Page *string `json:"-"`
	// 分页大小，默认20条，最大20条
	PageSize *string                 `json:"-"`
	Body     *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1PmiMeetingsGetResponse struct {
	*xhttp.ApiResponse
	Data *V1PmiMeetingsGet200Response `json:"data,omitempty"`
}

/*
V1PmiMeetingsGet 查询个人会议号会议列表[/v1/pmi-meetings - Get]

查询个人会议号（PMI）会议的会议列表（待开始、进行中、已结束），目前暂不支持 OAuth2.0 鉴权访问。
*/
func (s *meetingsAPIService) V1PmiMeetingsGet(ctx context.Context, request *ApiV1PmiMeetingsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1PmiMeetingsGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/pmi-meetings",
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

	if request.Instanceid == nil {
		return nil, fmt.Errorf("instanceid is required and must be specified")
	}

	// path 参数
	// query 参数
	if request.OperatorId != nil {
		apiReq.QueryParams.Set("operator_id", core.QueryValue(request.OperatorId))
	}
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

	response = &ApiV1PmiMeetingsGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1PmiMeetingsGet200Response),
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
