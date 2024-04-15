package wemeet

import (
	core "github.com/TencentCloud/wemeet-openapi-sdk-go/wemeet_openapi/core"
	xhttp "github.com/TencentCloud/wemeet-openapi-sdk-go/wemeet_openapi/core/xhttp"
	meetings "github.com/TencentCloud/wemeet-openapi-sdk-go/wemeet_openapi/service/meetings"
	records "github.com/TencentCloud/wemeet-openapi-sdk-go/wemeet_openapi/service/records"
	meeting_control "github.com/TencentCloud/wemeet-openapi-sdk-go/wemeet_openapi/service/meeting_control"
	meeting_guest "github.com/TencentCloud/wemeet-openapi-sdk-go/wemeet_openapi/service/meeting_guest"
)

type Client struct {
	config *core.Config
	// service interfaces
	MeetingsApi meetings.Service
	RecordsApi records.Service
	MeetingControlApi meeting_control.Service
	MeetingGuestApi meeting_guest.Service
}

func (clt Client) GetConfig() core.Config {
	return *clt.config
}

type ClientOptionFunc func(config *core.Config)

// WithAppId 腾讯会议分配给企业的企业 ID。
//
// 企业管理员可以登录 腾讯会议官网，单击右上角用户中心，在左侧菜单栏中的企业管理 > 账户管理 > 账户信息中进行查看。
//
// 开发者可以单击右上角用户中心，在左侧菜单栏中的企业管理 > 高级 > REST API 应用信息中查看。
func WithAppId(appId string) ClientOptionFunc {
	return func(config *core.Config) {
		config.AppId = appId
	}
}

// WithSdkId 设置用户子账号或开发的应用 ID。
//
// 企业管理员可以登录 腾讯会议官网，单击右上角用户中心，
// 在左侧菜单栏中的企业管理 > 高级 > REST API 中进行查看（如存在 SdkId 则必须填写，早期申请 API 且未分配 SdkId 的客户可不填写）。
func WithSdkId(sdkId string) ClientOptionFunc {
	return func(config *core.Config) {
		config.SdkId = sdkId
	}
}

// WithSecret 设置应用生成的 Secret ID 和 Secret Key。JWT 鉴权用。
func WithSecret(secretID string, secretKey string) ClientOptionFunc {
	return func(config *core.Config) {
		config.SecretID = secretID
		config.SecretKey = secretKey
	}
}

// WithAppVersion 设置应用 App 的版本号，建议设置，以便灰度和查找问题。
func WithAppVersion(version string) ClientOptionFunc {
	return func(config *core.Config) {
		config.Version = version
	}
}

// WithHTTPClient 设置自定义 wemeet http client
func WithHTTPClient(clt xhttp.Client) ClientOptionFunc {
	return func(config *core.Config) {
		config.Clt = clt
	}
}

func NewClient(options ...ClientOptionFunc) *Client {
	// 构建配置
	config := &core.Config{}

	// 自定义配置
	for _, opt := range options {
		opt(config)
	}

	// 构建 wemeet http client
	if config.Clt == nil {
		config.Clt, _ = xhttp.NewClient(core.OpenAPIDomain,
			xhttp.WithProtocol(core.DefaultProtocol),
			xhttp.WithSerializer(core.DefaultSerializer))
	}

	client := &Client{config: config}
	// 注册服务
	initService(client)
	return client
}

func initService(client *Client) {

	// 注册服务
	client.MeetingsApi = meetings.NewService(client.config)
	client.RecordsApi = records.NewService(client.config)
	client.MeetingControlApi = meeting_control.NewService(client.config)
	client.MeetingGuestApi = meeting_guest.NewService(client.config)
}
