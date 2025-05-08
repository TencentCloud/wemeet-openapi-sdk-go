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
	   V1AddressesGet 查询会议录制地址[/v1/addresses - Get]

	   **描述：**

	   * 查询会议录制地址，可获取会议云录制的播放地址和下载地址。
	   * 企业 secret 鉴权用户可获取该用户所属企业下的会议录制地址，OAuth2.0 鉴权用户只能获取该企业下 OAuth2.0 应用的会议录制地址。
	   * 当您想实时监测会议录制相关状况时，您可以通过订阅 [录制管理](https://cloud.tencent.com/document/product/1095/53226) 中的相关事件，接收事件通知。
	   * 当前同一场会议的不同录制文件共用分享链接。

	*/
	V1AddressesGet(ctx context.Context, request *ApiV1AddressesGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1AddressesGetResponse, err error)

	/*
	   V1AddressesRecordFileIdGet 查询单个录制详情（文件、转写、纪要）[/v1/addresses/{record_file_id} - Get]

	   查询单个云录制的详情信息，包括录制文件和会议纪要，并可获取播放地址和下载地址。企业 secert 鉴权用户可获取该用户所属企业下的单个录制列表，OAuth2.0 鉴权用户只能获取该企业下 OAuth2.0 应用的单个录制列表。

	*/
	V1AddressesRecordFileIdGet(ctx context.Context, request *ApiV1AddressesRecordFileIdGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1AddressesRecordFileIdGetResponse, err error)

	/*
	   V1FilesRecordsUploadAllPost 上传录制文件[/v1/files/records/upload-all - Post]

	*/
	V1FilesRecordsUploadAllPost(ctx context.Context, request *ApiV1FilesRecordsUploadAllPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1FilesRecordsUploadAllPostResponse, err error)

	/*
	   V1FilesRecordsUploadFinishPost 分块上传录制文件 - 上传完成[/v1/files/records/upload-finish - Post]

	*/
	V1FilesRecordsUploadFinishPost(ctx context.Context, request *ApiV1FilesRecordsUploadFinishPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1FilesRecordsUploadFinishPostResponse, err error)

	/*
	   V1FilesRecordsUploadPartPost 分块上传录制文件 - 上传[/v1/files/records/upload-part - Post]

	*/
	V1FilesRecordsUploadPartPost(ctx context.Context, request *ApiV1FilesRecordsUploadPartPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1FilesRecordsUploadPartPostResponse, err error)

	/*
	   V1FilesRecordsUploadPreparePost 分块上传录制文件 - 预上传[/v1/files/records/upload-prepare - Post]

	   分块上传录制文件 - 预上传

	*/
	V1FilesRecordsUploadPreparePost(ctx context.Context, request *ApiV1FilesRecordsUploadPreparePostRequest, opts ...core.RequestOptionFunc) (response *ApiV1FilesRecordsUploadPreparePostResponse, err error)

	/*
	   V1MetricsRecordsGet 查询录制文件访问数据[/v1/metrics/records - Get]

	   \*\*描述：\*\*查询会议录制 ID 对应的访问数据，按照天维度返回，支持 OAuth2\.0 鉴权调用。

	   * \*\*所需权限点为：\*\*查看会议录制（VIEW\_VIDEO） 或 管理会议录制（MANAGE\_VIDEO）。
	   * \*\*接口请求方法：\*\*GET

	*/
	V1MetricsRecordsGet(ctx context.Context, request *ApiV1MetricsRecordsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MetricsRecordsGetResponse, err error)

	/*
	   V1RecordsAccessMeetingRecordIdDelete 移除录制访问成员[/v1/records/access/{meeting_record_id} - Delete]

	   仅会议创建者、企业超级管理员或有企业录制管理权限的用户可调用。
	   权限点：管理会议录制

	*/
	V1RecordsAccessMeetingRecordIdDelete(ctx context.Context, request *ApiV1RecordsAccessMeetingRecordIdDeleteRequest, opts ...core.RequestOptionFunc) (response *ApiV1RecordsAccessMeetingRecordIdDeleteResponse, err error)

	/*
	   V1RecordsAccessMeetingRecordIdPost 添加录制访问成员[/v1/records/access/{meeting_record_id} - Post]

	   仅会议创建者、企业超级管理员或有企业录制管理权限的用户可调用，可以添加参会成员或企业内成员为访问成员。
	   权限点：管理会议录制

	*/
	V1RecordsAccessMeetingRecordIdPost(ctx context.Context, request *ApiV1RecordsAccessMeetingRecordIdPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1RecordsAccessMeetingRecordIdPostResponse, err error)

	/*
	   V1RecordsApprovalsMeetingRecordIdPut 审批云录制权限[/v1/records/approvals/{meeting_record_id} - Put]

	   会议创建者，企业超级管理员，有企业录制管理权限的用户，可以对云录制观看申请进行审批操作。OAuth权限点录制管理

	*/
	V1RecordsApprovalsMeetingRecordIdPut(ctx context.Context, request *ApiV1RecordsApprovalsMeetingRecordIdPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1RecordsApprovalsMeetingRecordIdPutResponse, err error)

	/*
	   V1RecordsDelete 删除会议录制[/v1/records - Delete]

	   删除会议的所有录制文件，该接口会删除会议录制 ID 里对应的所有云录制文件。企业 secret 鉴权用户可删除该用户所属企业下的会议录制，OAuth2.0 鉴权用户只能删除该企业下 OAuth2.0 应用的会议录制。
	   <span class="colour" style="color:rgb(51, 51, 51)">当您想实时监测会议录制删除状况时，您可以通过订阅 </span>[删除云录制](https://cloud.tencent.com/document/product/1095/53232)<span class="colour" style="color:rgb(51, 51, 51)"> 的事件，接收事件通知。</span>

	*/
	V1RecordsDelete(ctx context.Context, request *ApiV1RecordsDeleteRequest, opts ...core.RequestOptionFunc) (response *ApiV1RecordsDeleteResponse, err error)

	/*
	   V1RecordsEventsGet 获取会议录制操作（查看、下载）记录[/v1/records/events - Get]

	   \*\*描述：\*\*查询会议录制 ID 对应的操作记录，包括用户查看和下载，支持 OAuth2\.0 鉴权调用。

	   * \*\*所需权限点为：\*\*查看会议录制（VIEW\_VIDEO） 或 管理会议录制（MANAGE\_VIDEO）。
	   * \*\*接口请求方法：\*\*GET

	*/
	V1RecordsEventsGet(ctx context.Context, request *ApiV1RecordsEventsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1RecordsEventsGetResponse, err error)

	/*
	   V1RecordsGet 查询会议录制列表[/v1/records - Get]

	   获取用户云录制记录，根据用户 ID、会议 ID、会议 code 进行查询，支持根据时间区间分页获取。
	   企业 secret 鉴权用户可获取该用户所属企业下的会议录制列表，OAuth2.0 鉴权用户只能获取该企业下 OAuth2.0 应用的会议录制列表。
	   当您想实时监测会议录制相关状况时，您可以通过订阅 [录制管理](https://cloud.tencent.com/document/product/1095/53226) 中的相关事件，接收事件通知。
	   当前同一场会议的不同录制文件共用分享链接。

	*/
	V1RecordsGet(ctx context.Context, request *ApiV1RecordsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1RecordsGetResponse, err error)

	/*
	   V1RecordsRecordFileIdDelete 删除单个录制文件[/v1/records/{record_file_id} - Delete]

	   删除单个录制文件，该接口支持从会议中删除指定的某个录制文件。企业 secret 鉴权用户可删除该用户所属企业下的单个录制文件，OAuth2.0 鉴权用户只能删除该企业下 OAuth2.0 应用的单个录制文件。
	   <span class="colour" style="color:rgb(51, 51, 51)">当您想实时监测会议录制删除状况时，您可以通过订阅 </span>[删除云录制](https://cloud.tencent.com/document/product/1095/53232)<span class="colour" style="color:rgb(51, 51, 51)"> 的事件，接收事件通知。</span>

	*/
	V1RecordsRecordFileIdDelete(ctx context.Context, request *ApiV1RecordsRecordFileIdDeleteRequest, opts ...core.RequestOptionFunc) (response *ApiV1RecordsRecordFileIdDeleteResponse, err error)

	/*
	   V1RecordsSettingsMeetingRecordIdGet 查询会议录制共享设置[/v1/records/settings/{meeting_record_id} - Get]

	   根据会议录制 ID 查询共享等配置，支持修改共享权限、共享密码、共享有效期等信息，

	*/
	V1RecordsSettingsMeetingRecordIdGet(ctx context.Context, request *ApiV1RecordsSettingsMeetingRecordIdGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1RecordsSettingsMeetingRecordIdGetResponse, err error)

	/*
	   V1RecordsSettingsMeetingRecordIdPut 修改会议录制共享设置[/v1/records/settings/{meeting_record_id} - Put]

	   根据会议录制 ID 修改共享等配置，支持修改共享权限、共享密码、共享有效期等信息，支持 OAuth2.0 鉴权调用。
	   所需权限点为：管理会议录制（MANAGE\_VIDEO）。

	*/
	V1RecordsSettingsMeetingRecordIdPut(ctx context.Context, request *ApiV1RecordsSettingsMeetingRecordIdPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1RecordsSettingsMeetingRecordIdPutResponse, err error)

	/*
	   V1RecordsTranscriptsDetailsGet 查询录制转写详情[/v1/records/transcripts/details - Get]

	   获取云录制会议纪要的详情，包含时间戳、文本等内容。支持 OAuth2.0 鉴权调用，仅支持授权用户为商业版、企业版、教育版。

	   所需权限点为：查看会议录制（VIEW\_VIDEO） 或 管理会议录制（MANAGE\_VIDEO）。

	*/
	V1RecordsTranscriptsDetailsGet(ctx context.Context, request *ApiV1RecordsTranscriptsDetailsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1RecordsTranscriptsDetailsGetResponse, err error)

	/*
	   V1RecordsTranscriptsParagraphsGet 查询录制转写段落信息[/v1/records/transcripts/paragraphs - Get]

	   获取云录制会议纪要的段落信息（段落总数、段落 ID）。支持 OAuth2\.0 鉴权调用，仅支持授权用户为商业版、企业版、教育版。

	   所需权限点为：查看会议录制（VIEW\_VIDEO） 或 管理会议录制（MANAGE\_VIDEO）。

	*/
	V1RecordsTranscriptsParagraphsGet(ctx context.Context, request *ApiV1RecordsTranscriptsParagraphsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1RecordsTranscriptsParagraphsGetResponse, err error)

	/*
	   V1RecordsTranscriptsSearchGet 查询录制转写搜索结果[/v1/records/transcripts/search - Get]

	   根据指定内容搜索会议纪要。支持 OAuth2\.0 鉴权调用，仅支持授权用户为商业版、企业版、教育版。

	   所需权限点为：查看会议录制（VIEW\_VIDEO） 或 管理会议录制（MANAGE\_VIDEO）。

	*/
	V1RecordsTranscriptsSearchGet(ctx context.Context, request *ApiV1RecordsTranscriptsSearchGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1RecordsTranscriptsSearchGetResponse, err error)

	/*
	   V1RecordsTransferRecordingPut 设置专网会议转存[/v1/records/transfer-recording - Put]

	   描述：设置指定会议的录制文件是否转存
	   企业 secret 鉴权用户和OAuth2.0 鉴权用户并且有录制访问权限可指定会议录制设置。
	   设置仅支持对专网会议录制生效，如果会议为公网会议则返回失败
	   通过会议录制ID设置录制是否转存， 根据混合云集群是否开启转存：
	   如果混合云集群已开启录制转存功能
	   对指定的会议录制可通过API设置转存，和转存完成后的删除策略
	   如果录制未加入转存任务或转存失败， 则将录制加入转存任务
	   如果录制已加入转存任务， 或转存已完成， 则返回失败
	   如果混合云集群未开启专网会议录制转存
	   不支持通过API设置会议录制的转存， 返回失败

	*/
	V1RecordsTransferRecordingPut(ctx context.Context, request *ApiV1RecordsTransferRecordingPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1RecordsTransferRecordingPutResponse, err error)

	/*
	   V1RecordsViewDetailsGet 查询录制文件观看流水记录[/v1/records/view-details - Get]

	   查询会议云录制在一段时间内的观看记录，每次播放都会有一条记录。
	   支持 JWT 和 OAuth，OAuth 2.0鉴权用户只能获取该企业下 OAuth 2.0应用创建的会议记录
	   权限点：查看会议录制或管理会议录制。

	*/
	V1RecordsViewDetailsGet(ctx context.Context, request *ApiV1RecordsViewDetailsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1RecordsViewDetailsGetResponse, err error)
}

type recordsAPIService struct {
	config *core.Config
}

func NewService(config *core.Config) Service {
	return &recordsAPIService{
		config: config,
	}
}

type ApiV1AddressesGetRequest struct {
	// 会议录制 ID。
	MeetingRecordId *string `json:"-"`
	// 操作者ID 必须与operator_id_type 同时提供
	OperatorId *string `json:"-"`
	// 操作者ID的类型 3为rooms_id 必须与operator_id_type 同时提供
	OperatorIdType *string `json:"-"`
	// 用户 ID（企业内部请使用企业唯一用户标识；OAuth2.0 鉴权用户请使用 openId）。
	Userid *string `json:"-"`
	// 分页size
	PageSize *string `json:"-"`
	// 分页page
	Page        *string                 `json:"-"`
	AddressType *string                 `json:"-"`
	Body        *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1AddressesGetResponse struct {
	*xhttp.ApiResponse
	Data *V1AddressesGet200Response `json:"data,omitempty"`
}

/*
V1AddressesGet 查询会议录制地址[/v1/addresses - Get]

**描述：**

* 查询会议录制地址，可获取会议云录制的播放地址和下载地址。
* 企业 secret 鉴权用户可获取该用户所属企业下的会议录制地址，OAuth2.0 鉴权用户只能获取该企业下 OAuth2.0 应用的会议录制地址。
* 当您想实时监测会议录制相关状况时，您可以通过订阅 [录制管理](https://cloud.tencent.com/document/product/1095/53226) 中的相关事件，接收事件通知。
* 当前同一场会议的不同录制文件共用分享链接。
*/
func (s *recordsAPIService) V1AddressesGet(ctx context.Context, request *ApiV1AddressesGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1AddressesGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/addresses",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	if request.MeetingRecordId == nil {
		return nil, fmt.Errorf("meeting_record_id is required and must be specified")
	}

	// path 参数
	// query 参数
	if request.MeetingRecordId != nil {
		apiReq.QueryParams.Set("meeting_record_id", core.QueryValue(request.MeetingRecordId))
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
	if request.PageSize != nil {
		apiReq.QueryParams.Set("page_size", core.QueryValue(request.PageSize))
	}
	if request.Page != nil {
		apiReq.QueryParams.Set("page", core.QueryValue(request.Page))
	}
	if request.AddressType != nil {
		apiReq.QueryParams.Set("address_type", core.QueryValue(request.AddressType))
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

	response = &ApiV1AddressesGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1AddressesGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1AddressesRecordFileIdGetRequest struct {
	RecordFileId string `json:"-"`
	// 操作者ID，必须与operator_id_type同时出现。
	OperatorId *string `json:"-"`
	// 操作者ID的类型 rooms_Id是3，必须与operator_id同时出现。
	OperatorIdType *string `json:"-"`
	// 用户 ID（企业内部请使用企业唯一用户标识；OAuth2.0 鉴权用户请使用 openId）。
	Userid      *string                 `json:"-"`
	AddressType *string                 `json:"-"`
	Body        *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1AddressesRecordFileIdGetResponse struct {
	*xhttp.ApiResponse
	Data *V1AddressesRecordFileIdGet200Response `json:"data,omitempty"`
}

/*
V1AddressesRecordFileIdGet 查询单个录制详情（文件、转写、纪要）[/v1/addresses/{record_file_id} - Get]

查询单个云录制的详情信息，包括录制文件和会议纪要，并可获取播放地址和下载地址。企业 secert 鉴权用户可获取该用户所属企业下的单个录制列表，OAuth2.0 鉴权用户只能获取该企业下 OAuth2.0 应用的单个录制列表。
*/
func (s *recordsAPIService) V1AddressesRecordFileIdGet(ctx context.Context, request *ApiV1AddressesRecordFileIdGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1AddressesRecordFileIdGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/addresses/{record_file_id}",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	// path 参数
	apiReq.PathParams.Set("record_file_id", core.PathValue(request.RecordFileId))
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
	if request.AddressType != nil {
		apiReq.QueryParams.Set("address_type", core.QueryValue(request.AddressType))
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

	response = &ApiV1AddressesRecordFileIdGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1AddressesRecordFileIdGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1FilesRecordsUploadAllPostRequest struct {
	OperatorId     *string `json:"-"`
	OperatorIdType *int32  `json:"-"`
	FileName       *string `json:"-"`
	FileType       *string `json:"-"`
	FileSize       *int32  `json:"-"`
	FileChecksum   *string `json:"-"`
	FileContent    *[]byte `json:"-"`
	SpeakNumber    *int32  `json:"-"`
	AiRecord       *bool   `json:"-"`
}

type ApiV1FilesRecordsUploadAllPostResponse struct {
	*xhttp.ApiResponse
	Data *V1FilesRecordsUploadAllPost200Response `json:"data,omitempty"`
}

/*
V1FilesRecordsUploadAllPost 上传录制文件[/v1/files/records/upload-all - Post]
*/
func (s *recordsAPIService) V1FilesRecordsUploadAllPost(ctx context.Context, request *ApiV1FilesRecordsUploadAllPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1FilesRecordsUploadAllPostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/files/records/upload-all",
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	if request.OperatorId == nil {
		return nil, fmt.Errorf("operator_id is required and must be specified")
	}

	if request.OperatorIdType == nil {
		return nil, fmt.Errorf("operator_id_type is required and must be specified")
	}

	if request.FileName == nil {
		return nil, fmt.Errorf("file_name is required and must be specified")
	}

	if request.FileType == nil {
		return nil, fmt.Errorf("file_type is required and must be specified")
	}

	if request.FileSize == nil {
		return nil, fmt.Errorf("file_size is required and must be specified")
	}

	if request.FileChecksum == nil {
		return nil, fmt.Errorf("file_checksum is required and must be specified")
	}

	if request.FileContent == nil {
		return nil, fmt.Errorf("file_content is required and must be specified")
	}

	if request.SpeakNumber == nil {
		return nil, fmt.Errorf("speak_number is required and must be specified")
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

	// 创建一个新的FormData
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("operator_id", core.ToString(*request.OperatorId))
	writer.WriteField("operator_id_type", core.ToString(*request.OperatorIdType))
	writer.WriteField("file_name", core.ToString(*request.FileName))
	writer.WriteField("file_type", core.ToString(*request.FileType))
	writer.WriteField("file_size", core.ToString(*request.FileSize))
	writer.WriteField("file_checksum", core.ToString(*request.FileChecksum))
	// 添加文件到FormData
	fileContentPart, err := writer.CreateFormFile("file_content", "fileContent")
	if err != nil {
		fmt.Println("创建表单文件失败:", err)
		return
	}
	_, err = io.Copy(fileContentPart, bytes.NewReader(*request.FileContent))
	if err != nil {
		fmt.Println("复制文件内容失败:", err)
		return
	}
	writer.WriteField("speak_number", core.ToString(*request.SpeakNumber))
	writer.WriteField("ai_record", core.ToString(*request.AiRecord))

	// 设置请求头
	// from-data设置请求头
	_ = writer.Close()
	apiReq.Body = body
	header := make(http.Header)
	header.Set("Content-Type", writer.FormDataContentType())
	httpOptions = append(httpOptions, xhttp.WithRequestHeader(header))
	apiReq.Body = body

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

	response = &ApiV1FilesRecordsUploadAllPostResponse{
		ApiResponse: apiRsp,
		Data:        new(V1FilesRecordsUploadAllPost200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1FilesRecordsUploadFinishPostRequest struct {
	Body *V1FilesRecordsUploadFinishPostRequest `json:"body,omitempty"`
}

type ApiV1FilesRecordsUploadFinishPostResponse struct {
	*xhttp.ApiResponse
	Data *V1FilesRecordsUploadAllPost200Response `json:"data,omitempty"`
}

/*
V1FilesRecordsUploadFinishPost 分块上传录制文件 - 上传完成[/v1/files/records/upload-finish - Post]
*/
func (s *recordsAPIService) V1FilesRecordsUploadFinishPost(ctx context.Context, request *ApiV1FilesRecordsUploadFinishPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1FilesRecordsUploadFinishPostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/files/records/upload-finish",
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

	response = &ApiV1FilesRecordsUploadFinishPostResponse{
		ApiResponse: apiRsp,
		Data:        new(V1FilesRecordsUploadAllPost200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1FilesRecordsUploadPartPostRequest struct {
	// String 操作者 ID。operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。
	OperatorId *string `json:"-"`
	// 操作人 ID 类型： 1：userid
	OperatorIdType *int32 `json:"-"`
	// 上传事务 ID。
	UploadId *string `json:"-"`
	// 文件大小（以字节为单位），需按预上传返回的 block_size 填写（最后一个文件块按照实际大小填写）。
	FileSize *int32 `json:"-"`
	// 文件块号，从1开始计数。最后一个文件块允许小于 block_size 的值。
	FileSeq *int32 `json:"-"`
	// 文件校验和，文件内容 MD5 结果的十六进制表示。
	FileChecksum *string `json:"-"`
	// 文件二进制内容。
	FileContent *[]byte `json:"-"`
}

type ApiV1FilesRecordsUploadPartPostResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1FilesRecordsUploadPartPost 分块上传录制文件 - 上传[/v1/files/records/upload-part - Post]
*/
func (s *recordsAPIService) V1FilesRecordsUploadPartPost(ctx context.Context, request *ApiV1FilesRecordsUploadPartPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1FilesRecordsUploadPartPostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/files/records/upload-part",
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	if request.OperatorId == nil {
		return nil, fmt.Errorf("operator_id is required and must be specified")
	}

	if request.OperatorIdType == nil {
		return nil, fmt.Errorf("operator_id_type is required and must be specified")
	}

	if request.UploadId == nil {
		return nil, fmt.Errorf("upload_id is required and must be specified")
	}

	if request.FileSize == nil {
		return nil, fmt.Errorf("file_size is required and must be specified")
	}

	if request.FileSeq == nil {
		return nil, fmt.Errorf("file_seq is required and must be specified")
	}

	if request.FileChecksum == nil {
		return nil, fmt.Errorf("file_checksum is required and must be specified")
	}

	if request.FileContent == nil {
		return nil, fmt.Errorf("file_content is required and must be specified")
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

	// 创建一个新的FormData
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("operator_id", core.ToString(*request.OperatorId))
	writer.WriteField("operator_id_type", core.ToString(*request.OperatorIdType))
	writer.WriteField("upload_id", core.ToString(*request.UploadId))
	writer.WriteField("file_size", core.ToString(*request.FileSize))
	writer.WriteField("file_seq", core.ToString(*request.FileSeq))
	writer.WriteField("file_checksum", core.ToString(*request.FileChecksum))
	// 添加文件到FormData
	fileContentPart, err := writer.CreateFormFile("file_content", "fileContent")
	if err != nil {
		fmt.Println("创建表单文件失败:", err)
		return
	}
	_, err = io.Copy(fileContentPart, bytes.NewReader(*request.FileContent))
	if err != nil {
		fmt.Println("复制文件内容失败:", err)
		return
	}

	// 设置请求头
	// from-data设置请求头
	_ = writer.Close()
	apiReq.Body = body
	header := make(http.Header)
	header.Set("Content-Type", writer.FormDataContentType())
	httpOptions = append(httpOptions, xhttp.WithRequestHeader(header))
	apiReq.Body = body

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

	response = &ApiV1FilesRecordsUploadPartPostResponse{
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

type ApiV1FilesRecordsUploadPreparePostRequest struct {
	Body *V1FilesRecordsUploadPreparePostRequest `json:"body,omitempty"`
}

type ApiV1FilesRecordsUploadPreparePostResponse struct {
	*xhttp.ApiResponse
	Data *V1FilesRecordsUploadPreparePost200Response `json:"data,omitempty"`
}

/*
V1FilesRecordsUploadPreparePost 分块上传录制文件 - 预上传[/v1/files/records/upload-prepare - Post]

分块上传录制文件 - 预上传
*/
func (s *recordsAPIService) V1FilesRecordsUploadPreparePost(ctx context.Context, request *ApiV1FilesRecordsUploadPreparePostRequest, opts ...core.RequestOptionFunc) (response *ApiV1FilesRecordsUploadPreparePostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/files/records/upload-prepare",
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

	response = &ApiV1FilesRecordsUploadPreparePostResponse{
		ApiResponse: apiRsp,
		Data:        new(V1FilesRecordsUploadPreparePost200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1MetricsRecordsGetRequest struct {
	// 会议录制 ID。
	MeetingRecordId *string `json:"-"`
	// 查询起始时间戳，UNIX 时间戳（单位秒）。说明：时间区间不允许超过31天。
	StartTime *string `json:"-"`
	// 查询结束时间戳，UNIX 时间戳（单位秒）。说明：时间区间不允许超过31天。
	EndTime *string                 `json:"-"`
	Body    *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1MetricsRecordsGetResponse struct {
	*xhttp.ApiResponse
	Data *V1MetricsRecordsGet200Response `json:"data,omitempty"`
}

/*
V1MetricsRecordsGet 查询录制文件访问数据[/v1/metrics/records - Get]

\*\*描述：\*\*查询会议录制 ID 对应的访问数据，按照天维度返回，支持 OAuth2\.0 鉴权调用。

* \*\*所需权限点为：\*\*查看会议录制（VIEW\_VIDEO） 或 管理会议录制（MANAGE\_VIDEO）。
* \*\*接口请求方法：\*\*GET
*/
func (s *recordsAPIService) V1MetricsRecordsGet(ctx context.Context, request *ApiV1MetricsRecordsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1MetricsRecordsGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/metrics/records",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	if request.MeetingRecordId == nil {
		return nil, fmt.Errorf("meeting_record_id is required and must be specified")
	}

	// path 参数
	// query 参数
	if request.MeetingRecordId != nil {
		apiReq.QueryParams.Set("meeting_record_id", core.QueryValue(request.MeetingRecordId))
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

	response = &ApiV1MetricsRecordsGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1MetricsRecordsGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1RecordsAccessMeetingRecordIdDeleteRequest struct {
	// 会议录制ID
	MeetingRecordId string                                       `json:"-"`
	Body            *V1RecordsAccessMeetingRecordIdDeleteRequest `json:"body,omitempty"`
}

type ApiV1RecordsAccessMeetingRecordIdDeleteResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1RecordsAccessMeetingRecordIdDelete 移除录制访问成员[/v1/records/access/{meeting_record_id} - Delete]

仅会议创建者、企业超级管理员或有企业录制管理权限的用户可调用。
权限点：管理会议录制
*/
func (s *recordsAPIService) V1RecordsAccessMeetingRecordIdDelete(ctx context.Context, request *ApiV1RecordsAccessMeetingRecordIdDeleteRequest, opts ...core.RequestOptionFunc) (response *ApiV1RecordsAccessMeetingRecordIdDeleteResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/records/access/{meeting_record_id}",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	// path 参数
	apiReq.PathParams.Set("meeting_record_id", core.PathValue(request.MeetingRecordId))
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

	response = &ApiV1RecordsAccessMeetingRecordIdDeleteResponse{
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

type ApiV1RecordsAccessMeetingRecordIdPostRequest struct {
	// 会议录制ID
	MeetingRecordId string                                     `json:"-"`
	Body            *V1RecordsAccessMeetingRecordIdPostRequest `json:"body,omitempty"`
}

type ApiV1RecordsAccessMeetingRecordIdPostResponse struct {
	*xhttp.ApiResponse
	Data *V1RecordsAccessMeetingRecordIdPost200Response `json:"data,omitempty"`
}

/*
V1RecordsAccessMeetingRecordIdPost 添加录制访问成员[/v1/records/access/{meeting_record_id} - Post]

仅会议创建者、企业超级管理员或有企业录制管理权限的用户可调用，可以添加参会成员或企业内成员为访问成员。
权限点：管理会议录制
*/
func (s *recordsAPIService) V1RecordsAccessMeetingRecordIdPost(ctx context.Context, request *ApiV1RecordsAccessMeetingRecordIdPostRequest, opts ...core.RequestOptionFunc) (response *ApiV1RecordsAccessMeetingRecordIdPostResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/records/access/{meeting_record_id}",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	// path 参数
	apiReq.PathParams.Set("meeting_record_id", core.PathValue(request.MeetingRecordId))
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

	response = &ApiV1RecordsAccessMeetingRecordIdPostResponse{
		ApiResponse: apiRsp,
		Data:        new(V1RecordsAccessMeetingRecordIdPost200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1RecordsApprovalsMeetingRecordIdPutRequest struct {
	// 会议录制 ID，列表查询接口返回的 meeting_record_id。
	MeetingRecordId string                                       `json:"-"`
	Body            *V1RecordsApprovalsMeetingRecordIdPutRequest `json:"body,omitempty"`
}

type ApiV1RecordsApprovalsMeetingRecordIdPutResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1RecordsApprovalsMeetingRecordIdPut 审批云录制权限[/v1/records/approvals/{meeting_record_id} - Put]

会议创建者，企业超级管理员，有企业录制管理权限的用户，可以对云录制观看申请进行审批操作。OAuth权限点录制管理
*/
func (s *recordsAPIService) V1RecordsApprovalsMeetingRecordIdPut(ctx context.Context, request *ApiV1RecordsApprovalsMeetingRecordIdPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1RecordsApprovalsMeetingRecordIdPutResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/records/approvals/{meeting_record_id}",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	// path 参数
	apiReq.PathParams.Set("meeting_record_id", core.PathValue(request.MeetingRecordId))
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

	response = &ApiV1RecordsApprovalsMeetingRecordIdPutResponse{
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

type ApiV1RecordsDeleteRequest struct {
	// 会议录制 ID。
	MeetingRecordId *string `json:"-"`
	// 会议 ID。
	MeetingId *string `json:"-"`
	// 操作者ID，根据operator_id_type的值，使用不同的类型
	OperatorId *string `json:"-"`
	// 操作者ID的类型，必须与operator_id同时出现
	OperatorIdType *string `json:"-"`
	// 用户 ID（企业内部请使用企业唯一用户标识；OAuth2.0 鉴权用户请使用 openId）。
	Userid *string                 `json:"-"`
	Body   *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1RecordsDeleteResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1RecordsDelete 删除会议录制[/v1/records - Delete]

删除会议的所有录制文件，该接口会删除会议录制 ID 里对应的所有云录制文件。企业 secret 鉴权用户可删除该用户所属企业下的会议录制，OAuth2.0 鉴权用户只能删除该企业下 OAuth2.0 应用的会议录制。
<span class="colour" style="color:rgb(51, 51, 51)">当您想实时监测会议录制删除状况时，您可以通过订阅 </span>[删除云录制](https://cloud.tencent.com/document/product/1095/53232)<span class="colour" style="color:rgb(51, 51, 51)"> 的事件，接收事件通知。</span>
*/
func (s *recordsAPIService) V1RecordsDelete(ctx context.Context, request *ApiV1RecordsDeleteRequest, opts ...core.RequestOptionFunc) (response *ApiV1RecordsDeleteResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/records",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	if request.MeetingRecordId == nil {
		return nil, fmt.Errorf("meeting_record_id is required and must be specified")
	}

	// path 参数
	// query 参数
	if request.MeetingId != nil {
		apiReq.QueryParams.Set("meeting_id", core.QueryValue(request.MeetingId))
	}
	if request.MeetingRecordId != nil {
		apiReq.QueryParams.Set("meeting_record_id", core.QueryValue(request.MeetingRecordId))
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

	response = &ApiV1RecordsDeleteResponse{
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

type ApiV1RecordsEventsGetRequest struct {
	// 会议录制 ID，列表接口返回的是 meeting_record_id。
	MeetingRecordId *string `json:"-"`
	// 查询事件类型：1：下载，2：查看。
	EventType *string `json:"-"`
	// 分页大小，默认值为20，最大为50。
	PageSize *string `json:"-"`
	// 页码，从1开始，默认值为1。
	Page *string `json:"-"`
	// 查询起始时间戳，UNIX 时间戳（单位秒）。说明：时间区间不允许超过31天。
	StartTime *string `json:"-"`
	// 查询结束时间戳，UNIX 时间戳（单位秒）。说明：时间区间不允许超过31天。
	EndTime *string `json:"-"`
}

type ApiV1RecordsEventsGetResponse struct {
	*xhttp.ApiResponse
	Data *V1RecordsEventsGet200Response `json:"data,omitempty"`
}

/*
V1RecordsEventsGet 获取会议录制操作（查看、下载）记录[/v1/records/events - Get]

\*\*描述：\*\*查询会议录制 ID 对应的操作记录，包括用户查看和下载，支持 OAuth2\.0 鉴权调用。

* \*\*所需权限点为：\*\*查看会议录制（VIEW\_VIDEO） 或 管理会议录制（MANAGE\_VIDEO）。
* \*\*接口请求方法：\*\*GET
*/
func (s *recordsAPIService) V1RecordsEventsGet(ctx context.Context, request *ApiV1RecordsEventsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1RecordsEventsGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/records/events",
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	if request.MeetingRecordId == nil {
		return nil, fmt.Errorf("meeting_record_id is required and must be specified")
	}

	if request.EventType == nil {
		return nil, fmt.Errorf("event_type is required and must be specified")
	}

	// path 参数
	// query 参数
	if request.MeetingRecordId != nil {
		apiReq.QueryParams.Set("meeting_record_id", core.QueryValue(request.MeetingRecordId))
	}
	if request.EventType != nil {
		apiReq.QueryParams.Set("event_type", core.QueryValue(request.EventType))
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

	response = &ApiV1RecordsEventsGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1RecordsEventsGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1RecordsGetRequest struct {
	// 操作者ID，必须与operator_id_type同时出现。
	OperatorId *string `json:"-"`
	// 操作者ID的类型，必须与operator_id同时出现。
	OperatorIdType *string `json:"-"`
	// 查询起始时间戳，UNIX 时间戳（单位秒）。说明：时间区间不允许超过31天。
	StartTime *string `json:"-"`
	// 查询结束时间戳，UNIX 时间戳（单位秒）。说明：时间区间不允许超过31天。
	EndTime *string `json:"-"`
	// 会议的唯一 ID，不为空时优先根据会议 ID 查询。
	MeetingId *string `json:"-"`
	// 会议 code，当 meeting_id 为空且 meeting_code 不为空时根据会议 code 查询。
	MeetingCode *string `json:"-"`
	// 分页大小，默认值为10，最大为20。
	PageSize *string `json:"-"`
	// 页码，从1开始，默认值为1。
	Page         *string `json:"-"`
	MediaSetType *string `json:"-"`
	// 录制文件类型： 0：全部、1：云录制、2：上传录制，默认为 1
	QueryRecordType *string                 `json:"-"`
	Body            *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1RecordsGetResponse struct {
	*xhttp.ApiResponse
	Data *V1RecordsGet200Response `json:"data,omitempty"`
}

/*
V1RecordsGet 查询会议录制列表[/v1/records - Get]

获取用户云录制记录，根据用户 ID、会议 ID、会议 code 进行查询，支持根据时间区间分页获取。
企业 secret 鉴权用户可获取该用户所属企业下的会议录制列表，OAuth2.0 鉴权用户只能获取该企业下 OAuth2.0 应用的会议录制列表。
当您想实时监测会议录制相关状况时，您可以通过订阅 [录制管理](https://cloud.tencent.com/document/product/1095/53226) 中的相关事件，接收事件通知。
当前同一场会议的不同录制文件共用分享链接。
*/
func (s *recordsAPIService) V1RecordsGet(ctx context.Context, request *ApiV1RecordsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1RecordsGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/records",
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

	if request.StartTime == nil {
		return nil, fmt.Errorf("start_time is required and must be specified")
	}

	if request.EndTime == nil {
		return nil, fmt.Errorf("end_time is required and must be specified")
	}

	// path 参数
	// query 参数
	if request.OperatorId != nil {
		apiReq.QueryParams.Set("operator_id", core.QueryValue(request.OperatorId))
	}
	if request.OperatorIdType != nil {
		apiReq.QueryParams.Set("operator_id_type", core.QueryValue(request.OperatorIdType))
	}
	if request.MeetingId != nil {
		apiReq.QueryParams.Set("meeting_id", core.QueryValue(request.MeetingId))
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
	if request.PageSize != nil {
		apiReq.QueryParams.Set("page_size", core.QueryValue(request.PageSize))
	}
	if request.Page != nil {
		apiReq.QueryParams.Set("page", core.QueryValue(request.Page))
	}
	if request.MediaSetType != nil {
		apiReq.QueryParams.Set("media_set_type", core.QueryValue(request.MediaSetType))
	}
	if request.QueryRecordType != nil {
		apiReq.QueryParams.Set("query_record_type", core.QueryValue(request.QueryRecordType))
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

	response = &ApiV1RecordsGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1RecordsGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1RecordsRecordFileIdDeleteRequest struct {
	// 录制文件 ID。
	RecordFileId string `json:"-"`
	// 会议 ID。
	MeetingId *string `json:"-"`
	// 操作者ID，根据operator_id_type的值，使用不同的类型，必须与operator_id_type同时出现
	OperatorId *string `json:"-"`
	// 操作者ID的类型，必须与operator_id同时出现
	OperatorIdType *string `json:"-"`
	// 用户 ID（企业内部请使用企业唯一用户标识；OAuth2.0 鉴权用户请使用 openId）。
	Userid *string                 `json:"-"`
	Body   *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1RecordsRecordFileIdDeleteResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1RecordsRecordFileIdDelete 删除单个录制文件[/v1/records/{record_file_id} - Delete]

删除单个录制文件，该接口支持从会议中删除指定的某个录制文件。企业 secret 鉴权用户可删除该用户所属企业下的单个录制文件，OAuth2.0 鉴权用户只能删除该企业下 OAuth2.0 应用的单个录制文件。
<span class="colour" style="color:rgb(51, 51, 51)">当您想实时监测会议录制删除状况时，您可以通过订阅 </span>[删除云录制](https://cloud.tencent.com/document/product/1095/53232)<span class="colour" style="color:rgb(51, 51, 51)"> 的事件，接收事件通知。</span>
*/
func (s *recordsAPIService) V1RecordsRecordFileIdDelete(ctx context.Context, request *ApiV1RecordsRecordFileIdDeleteRequest, opts ...core.RequestOptionFunc) (response *ApiV1RecordsRecordFileIdDeleteResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/records/{record_file_id}",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	// path 参数
	apiReq.PathParams.Set("record_file_id", core.PathValue(request.RecordFileId))
	// query 参数
	if request.MeetingId != nil {
		apiReq.QueryParams.Set("meeting_id", core.QueryValue(request.MeetingId))
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

	response = &ApiV1RecordsRecordFileIdDeleteResponse{
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

type ApiV1RecordsSettingsMeetingRecordIdGetRequest struct {
	// 会议录制ID
	MeetingRecordId string `json:"-"`
	// 操作人ID,录制管理者、企业超级管理员或有企业录制管理权限的用
	OperatorId *string `json:"-"`
	// 操作人ID 类型 1-userid，2-openid,3-rooms_id
	OperatorIdType *string                 `json:"-"`
	Body           *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1RecordsSettingsMeetingRecordIdGetResponse struct {
	*xhttp.ApiResponse
	Data *V1RecordsSettingsMeetingRecordIdGet200Response `json:"data,omitempty"`
}

/*
V1RecordsSettingsMeetingRecordIdGet 查询会议录制共享设置[/v1/records/settings/{meeting_record_id} - Get]

根据会议录制 ID 查询共享等配置，支持修改共享权限、共享密码、共享有效期等信息，
*/
func (s *recordsAPIService) V1RecordsSettingsMeetingRecordIdGet(ctx context.Context, request *ApiV1RecordsSettingsMeetingRecordIdGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1RecordsSettingsMeetingRecordIdGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/records/settings/{meeting_record_id}",
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
	apiReq.PathParams.Set("meeting_record_id", core.PathValue(request.MeetingRecordId))
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

	response = &ApiV1RecordsSettingsMeetingRecordIdGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1RecordsSettingsMeetingRecordIdGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1RecordsSettingsMeetingRecordIdPutRequest struct {
	// 会议录制 ID。
	MeetingRecordId string                                      `json:"-"`
	Body            *V1RecordsSettingsMeetingRecordIdPutRequest `json:"body,omitempty"`
}

type ApiV1RecordsSettingsMeetingRecordIdPutResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1RecordsSettingsMeetingRecordIdPut 修改会议录制共享设置[/v1/records/settings/{meeting_record_id} - Put]

根据会议录制 ID 修改共享等配置，支持修改共享权限、共享密码、共享有效期等信息，支持 OAuth2.0 鉴权调用。
所需权限点为：管理会议录制（MANAGE\_VIDEO）。
*/
func (s *recordsAPIService) V1RecordsSettingsMeetingRecordIdPut(ctx context.Context, request *ApiV1RecordsSettingsMeetingRecordIdPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1RecordsSettingsMeetingRecordIdPutResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/records/settings/{meeting_record_id}",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	// path 参数
	apiReq.PathParams.Set("meeting_record_id", core.PathValue(request.MeetingRecordId))
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

	response = &ApiV1RecordsSettingsMeetingRecordIdPutResponse{
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

type ApiV1RecordsTranscriptsDetailsGetRequest struct {
	// 录制id
	RecordFileId *string `json:"-"`
	// 操作者ID。operator_id 必须与 operator_id_type 配合使用。根据operator_id_type的值，operator_id 代表不同类型。
	OperatorId *string `json:"-"`
	// 操作者ID的类型：  1. 企业用户userid 2：open_id 3. rooms设备rooms_id
	OperatorIdType *string `json:"-"`
	// 会议id
	MeetingId *string `json:"-"`
	// 查询的起始段落 ID。获取 pid 后（含）的段落，默认从0开始。
	Pid *string `json:"-"`
	// 查询的段落数，默认查询全量数据
	Limit *string `json:"-"`
	// 转写类型，默认是0。 0：原文版 1：智能优化版
	TranscriptsType *string                 `json:"-"`
	Body            *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1RecordsTranscriptsDetailsGetResponse struct {
	*xhttp.ApiResponse
	Data *V1RecordsTranscriptsDetailsGet200Response `json:"data,omitempty"`
}

/*
V1RecordsTranscriptsDetailsGet 查询录制转写详情[/v1/records/transcripts/details - Get]

获取云录制会议纪要的详情，包含时间戳、文本等内容。支持 OAuth2.0 鉴权调用，仅支持授权用户为商业版、企业版、教育版。

所需权限点为：查看会议录制（VIEW\_VIDEO） 或 管理会议录制（MANAGE\_VIDEO）。
*/
func (s *recordsAPIService) V1RecordsTranscriptsDetailsGet(ctx context.Context, request *ApiV1RecordsTranscriptsDetailsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1RecordsTranscriptsDetailsGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/records/transcripts/details",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	if request.RecordFileId == nil {
		return nil, fmt.Errorf("record_file_id is required and must be specified")
	}

	if request.OperatorId == nil {
		return nil, fmt.Errorf("operator_id is required and must be specified")
	}

	if request.OperatorIdType == nil {
		return nil, fmt.Errorf("operator_id_type is required and must be specified")
	}

	// path 参数
	// query 参数
	if request.MeetingId != nil {
		apiReq.QueryParams.Set("meeting_id", core.QueryValue(request.MeetingId))
	}
	if request.RecordFileId != nil {
		apiReq.QueryParams.Set("record_file_id", core.QueryValue(request.RecordFileId))
	}
	if request.Pid != nil {
		apiReq.QueryParams.Set("pid", core.QueryValue(request.Pid))
	}
	if request.Limit != nil {
		apiReq.QueryParams.Set("limit", core.QueryValue(request.Limit))
	}
	if request.OperatorId != nil {
		apiReq.QueryParams.Set("operator_id", core.QueryValue(request.OperatorId))
	}
	if request.OperatorIdType != nil {
		apiReq.QueryParams.Set("operator_id_type", core.QueryValue(request.OperatorIdType))
	}
	if request.TranscriptsType != nil {
		apiReq.QueryParams.Set("transcripts_type", core.QueryValue(request.TranscriptsType))
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

	response = &ApiV1RecordsTranscriptsDetailsGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1RecordsTranscriptsDetailsGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1RecordsTranscriptsParagraphsGetRequest struct {
	// 录制文件 ID。
	RecordFileId *string `json:"-"`
	// 操作者ID的类型：  1. 企业用户userid 2：open_id 3. rooms设备rooms_id
	OperatorIdType *string `json:"-"`
	// 操作者ID。operator_id 必须与 operator_id_type 配合使用。根据operator_id_type的值，operator_id 代表不同类型。
	OperatorId *string `json:"-"`
	// 会议 ID。
	MeetingId *string                 `json:"-"`
	Body      *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1RecordsTranscriptsParagraphsGetResponse struct {
	*xhttp.ApiResponse
	Data *V1RecordsTranscriptsParagraphsGet200Response `json:"data,omitempty"`
}

/*
V1RecordsTranscriptsParagraphsGet 查询录制转写段落信息[/v1/records/transcripts/paragraphs - Get]

获取云录制会议纪要的段落信息（段落总数、段落 ID）。支持 OAuth2\.0 鉴权调用，仅支持授权用户为商业版、企业版、教育版。

所需权限点为：查看会议录制（VIEW\_VIDEO） 或 管理会议录制（MANAGE\_VIDEO）。
*/
func (s *recordsAPIService) V1RecordsTranscriptsParagraphsGet(ctx context.Context, request *ApiV1RecordsTranscriptsParagraphsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1RecordsTranscriptsParagraphsGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/records/transcripts/paragraphs",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	if request.RecordFileId == nil {
		return nil, fmt.Errorf("record_file_id is required and must be specified")
	}

	if request.OperatorIdType == nil {
		return nil, fmt.Errorf("operator_id_type is required and must be specified")
	}

	if request.OperatorId == nil {
		return nil, fmt.Errorf("operator_id is required and must be specified")
	}

	// path 参数
	// query 参数
	if request.MeetingId != nil {
		apiReq.QueryParams.Set("meeting_id", core.QueryValue(request.MeetingId))
	}
	if request.RecordFileId != nil {
		apiReq.QueryParams.Set("record_file_id", core.QueryValue(request.RecordFileId))
	}
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

	response = &ApiV1RecordsTranscriptsParagraphsGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1RecordsTranscriptsParagraphsGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1RecordsTranscriptsSearchGetRequest struct {
	// 录制文件id
	RecordFileId *string `json:"-"`
	// 用户名
	OperatorId *string `json:"-"`
	// id类型: 1: 常规用户 2：open_id 3:rooms
	OperatorIdType *string `json:"-"`
	// 搜索的文本, 如果是中文, 需要urlencode一下
	Text *string `json:"-"`
	// 会议ID
	MeetingId *string `json:"-"`
	// 转写类型，默认是0。 0：原文版 1：智能优化版
	TranscriptsType *string                 `json:"-"`
	Body            *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1RecordsTranscriptsSearchGetResponse struct {
	*xhttp.ApiResponse
	Data *V1RecordsTranscriptsSearchGet200Response `json:"data,omitempty"`
}

/*
V1RecordsTranscriptsSearchGet 查询录制转写搜索结果[/v1/records/transcripts/search - Get]

根据指定内容搜索会议纪要。支持 OAuth2\.0 鉴权调用，仅支持授权用户为商业版、企业版、教育版。

所需权限点为：查看会议录制（VIEW\_VIDEO） 或 管理会议录制（MANAGE\_VIDEO）。
*/
func (s *recordsAPIService) V1RecordsTranscriptsSearchGet(ctx context.Context, request *ApiV1RecordsTranscriptsSearchGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1RecordsTranscriptsSearchGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/records/transcripts/search",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	if request.RecordFileId == nil {
		return nil, fmt.Errorf("record_file_id is required and must be specified")
	}

	if request.OperatorId == nil {
		return nil, fmt.Errorf("operator_id is required and must be specified")
	}

	if request.OperatorIdType == nil {
		return nil, fmt.Errorf("operator_id_type is required and must be specified")
	}

	if request.Text == nil {
		return nil, fmt.Errorf("text is required and must be specified")
	}

	// path 参数
	// query 参数
	if request.MeetingId != nil {
		apiReq.QueryParams.Set("meeting_id", core.QueryValue(request.MeetingId))
	}
	if request.RecordFileId != nil {
		apiReq.QueryParams.Set("record_file_id", core.QueryValue(request.RecordFileId))
	}
	if request.OperatorId != nil {
		apiReq.QueryParams.Set("operator_id", core.QueryValue(request.OperatorId))
	}
	if request.OperatorIdType != nil {
		apiReq.QueryParams.Set("operator_id_type", core.QueryValue(request.OperatorIdType))
	}
	if request.Text != nil {
		apiReq.QueryParams.Set("text", core.QueryValue(request.Text))
	}
	if request.TranscriptsType != nil {
		apiReq.QueryParams.Set("transcripts_type", core.QueryValue(request.TranscriptsType))
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

	response = &ApiV1RecordsTranscriptsSearchGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1RecordsTranscriptsSearchGet200Response),
	}
	if err = apiRsp.Translate(response.Data); err != nil {
		return nil, &core.ClientError{
			Err: fmt.Errorf("http status code: %d, response: %s, err: %v",
				apiRsp.StatusCode, apiRsp.RawBody, err),
		}
	}
	return
}

type ApiV1RecordsTransferRecordingPutRequest struct {
	// 会议ID
	MeetingId *string `json:"-"`
	// 会议录制ID
	MeetingRecordId *string                               `json:"-"`
	Body            *V1RecordsTransferRecordingPutRequest `json:"body,omitempty"`
}

type ApiV1RecordsTransferRecordingPutResponse struct {
	*xhttp.ApiResponse
	Data *map[string]interface{} `json:"data,omitempty"`
}

/*
V1RecordsTransferRecordingPut 设置专网会议转存[/v1/records/transfer-recording - Put]

描述：设置指定会议的录制文件是否转存
企业 secret 鉴权用户和OAuth2.0 鉴权用户并且有录制访问权限可指定会议录制设置。
设置仅支持对专网会议录制生效，如果会议为公网会议则返回失败
通过会议录制ID设置录制是否转存， 根据混合云集群是否开启转存：
如果混合云集群已开启录制转存功能
对指定的会议录制可通过API设置转存，和转存完成后的删除策略
如果录制未加入转存任务或转存失败， 则将录制加入转存任务
如果录制已加入转存任务， 或转存已完成， 则返回失败
如果混合云集群未开启专网会议录制转存
不支持通过API设置会议录制的转存， 返回失败
*/
func (s *recordsAPIService) V1RecordsTransferRecordingPut(ctx context.Context, request *ApiV1RecordsTransferRecordingPutRequest, opts ...core.RequestOptionFunc) (response *ApiV1RecordsTransferRecordingPutResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/records/transfer-recording",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	if request.MeetingId == nil {
		return nil, fmt.Errorf("meeting_id is required and must be specified")
	}

	if request.MeetingRecordId == nil {
		return nil, fmt.Errorf("meeting_record_id is required and must be specified")
	}

	// path 参数
	// query 参数
	if request.MeetingId != nil {
		apiReq.QueryParams.Set("meeting_id", core.QueryValue(request.MeetingId))
	}
	if request.MeetingRecordId != nil {
		apiReq.QueryParams.Set("meeting_record_id", core.QueryValue(request.MeetingRecordId))
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

	response = &ApiV1RecordsTransferRecordingPutResponse{
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

type ApiV1RecordsViewDetailsGetRequest struct {
	// 录制文件 ID。
	RecordFileId *string `json:"-"`
	// 操作者 ID 的类型： 1：userid 2：open_id
	OperatorIdType *string `json:"-"`
	// 操作者 ID。operator_id 必须与 operator_id_type 配合使用。根据 operator_id_type 的值，operator_id 代表不同类型。
	OperatorId *string `json:"-"`
	// 分页大小，默认20，最大50
	PageSize *string `json:"-"`
	// 页码，从1开始，默认1
	Page *string `json:"-"`
	// 查询起始时间戳，UNIX 时间戳（单位毫秒）。 说明：仅存储最近31天的数据，默认展示最近31天的数据。
	StartTime *string `json:"-"`
	// 查询结束时间戳，UNIX 时间戳（单位毫秒）。 说明：仅存储最近31天的数据，默认展示最近31天的数据
	EndTime *string                 `json:"-"`
	Body    *map[string]interface{} `json:"body,omitempty"`
}

type ApiV1RecordsViewDetailsGetResponse struct {
	*xhttp.ApiResponse
	Data *V1RecordsViewDetailsGet200Response `json:"data,omitempty"`
}

/*
V1RecordsViewDetailsGet 查询录制文件观看流水记录[/v1/records/view-details - Get]

查询会议云录制在一段时间内的观看记录，每次播放都会有一条记录。
支持 JWT 和 OAuth，OAuth 2.0鉴权用户只能获取该企业下 OAuth 2.0应用创建的会议记录
权限点：查看会议录制或管理会议录制。
*/
func (s *recordsAPIService) V1RecordsViewDetailsGet(ctx context.Context, request *ApiV1RecordsViewDetailsGetRequest, opts ...core.RequestOptionFunc) (response *ApiV1RecordsViewDetailsGetResponse, err error) {
	apiReq := &xhttp.ApiRequest{
		ApiURI:      "/v1/records/view-details",
		Body:        request.Body,
		PathParams:  xhttp.PathParams{},
		QueryParams: xhttp.QueryParams{},
	}

	if request.RecordFileId == nil {
		return nil, fmt.Errorf("record_file_id is required and must be specified")
	}

	if request.OperatorIdType == nil {
		return nil, fmt.Errorf("operator_id_type is required and must be specified")
	}

	if request.OperatorId == nil {
		return nil, fmt.Errorf("operator_id is required and must be specified")
	}

	// path 参数
	// query 参数
	if request.RecordFileId != nil {
		apiReq.QueryParams.Set("record_file_id", core.QueryValue(request.RecordFileId))
	}
	if request.OperatorIdType != nil {
		apiReq.QueryParams.Set("operator_id_type", core.QueryValue(request.OperatorIdType))
	}
	if request.OperatorId != nil {
		apiReq.QueryParams.Set("operator_id", core.QueryValue(request.OperatorId))
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

	response = &ApiV1RecordsViewDetailsGetResponse{
		ApiResponse: apiRsp,
		Data:        new(V1RecordsViewDetailsGet200Response),
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
