package wemeetcore

import xhttp "github.com/TencentCloud/wemeet-openapi-sdk-go/wemeet_openapi/core/xhttp"

// OpenAPIDomain 域名
const OpenAPIDomain = "api.meeting.qq.com"

// DefaultProtocol 默认协议
const DefaultProtocol = "https"

// SDK 使用的原生+自定义请求头
const (
	HTTPHeaderContentType = "Content-Type"
	HTTPHeaderUserAgent   = "User-Agent"

	XTCHeaderNonce      = "X-TC-Nonce"      // 随机正整数。
	XTCHeaderKey        = "X-TC-Key"        // 应用安全凭证密钥对中的 SecretId。
	XTCHeaderAction     = "X-TC-Action"     // 操作的接口名称。
	XTCHeaderRegion     = "X-TC-Region"     // 地域参数，用来标识希望操作哪个地域的数据。
	XTCHeaderTimestamp  = "X-TC-Timestamp"  // 当前 UNIX 时间戳，可记录发起 API 请求的时间。例如1529223702，单位为秒。
	XTCHeaderVersion    = "X-TC-Version"    // 应用 App 的版本号，建议设置，以便灰度和查找问题。
	XTCHeaderSignature  = "X-TC-Signature"  // 签名方法产生的签名。
	XTCHeaderRegistered = "X-TC-Registered" // 启用账户通讯录，传入值必须为1，创建的会议可出现在用户的会议列表中。
	HeaderAppId         = "AppId"           // 腾讯会议分配给企业的企业 ID。
	HeaderSdkId         = "SdkId"           // 用户子账号或开发的应用 ID。
	HeaderAccessToken   = "AccessToken"     // OAuth2.0 鉴权成功后返回的 token 信息。
	HeaderOpenId        = "OpenId"          // OAuth2.0 鉴权成功后的用户信息。
)

const (
	DefaultContentType = "application/json; charset=utf-8"
)

// JsonSerializer 全局 JSON 序列化器
var JsonSerializer = &JSONSerializer{}

// DefaultSerializer 默认序列化器
var DefaultSerializer xhttp.Serializable = JsonSerializer

// DefaultAuthenticator 默认鉴权器，用于增加 SDK 版本标识头
var DefaultAuthenticator xhttp.Authentication = &VersionAuthenticator{}
