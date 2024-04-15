package wemeetcore

import (
	"net/http"

	xhttp "github.com/TencentCloud/wemeet-openapi-sdk-go/wemeet_openapi/core/xhttp"
)

// Config 通用配置
type Config struct {
	Clt xhttp.Client // wemeet 封装的 http client

	Version string // 应用 App 的版本号，建议设置，以便灰度和查找问题。通过设置该字段，API 会把该版本信息传递给会议后台，以控制一些和 App 版本有关的特性。

	// 腾讯会议分配给企业的企业 ID。
	//
	// 企业管理员可以登录 腾讯会议官网，单击右上角用户中心，在左侧菜单栏中的企业管理 > 账户管理 > 账户信息中进行查看。
	//
	// 开发者可以单击右上角用户中心，在左侧菜单栏中的企业管理 > 高级  > REST API 应用信息中查看。
	AppId string

	// 用户子账号或开发的应用 ID。
	//
	// 企业管理员可以登录 腾讯会议官网，单击右上角用户中心，在左侧菜单栏中的企业管理 > 高级 > REST API 中进行查看（如存在 SdkId 则必须填写，早期申请 API 且未分配 SdkId 的客户可不填写）。
	SdkId     string
	SecretID  string // 应用生成的 Secret ID。JWT 鉴权用。
	SecretKey string // 应用生成的 Secret Key。JWT 鉴权用。
}

// RequestOptionFunc 对外请求配置
type RequestOptionFunc func(config Config) xhttp.RequestOptionFunc

// WithJWTAuth 配置 JWT 鉴权
func WithJWTAuth(jwt *JWTAuthenticator) RequestOptionFunc {
	return func(config Config) xhttp.RequestOptionFunc {
		jwt.config = &config
		return xhttp.WithRequestAuthenticator(jwt)
	}
}

// WithOAuth2Auth 配置 OAuth2 鉴权
func WithOAuth2Auth(oauth2 *OAuth2Authenticator) RequestOptionFunc {
	return func(config Config) xhttp.RequestOptionFunc {
		oauth2.config = &config
		return xhttp.WithRequestAuthenticator(oauth2)
	}
}

// WithRequestHeader 配置自定义请求头
func WithRequestHeader(header http.Header) RequestOptionFunc {
	return func(_ Config) xhttp.RequestOptionFunc {
		return xhttp.WithRequestHeader(header)
	}
}
