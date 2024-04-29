package wemeetcore

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// VersionAuthenticator SDK 版本鉴权器
type VersionAuthenticator struct{}

func (v VersionAuthenticator) AuthHeader(httpReq *http.Request) error {
	httpReq.Header.Set(HTTPHeaderUserAgent, "openapi-sdk-go/"+Version)
	return nil
}

// JWTAuthenticator jwt 鉴权器
// 提供标准的 jwt 鉴权头的设置和签名生成逻辑。
type JWTAuthenticator struct {
	config *Config

	Nonce uint64 // 此参数参与签名计算。随机正整数。SDK 默认自动随机生成。

	// 此参数参与签名计算。当前 UNIX 时间戳，可记录发起 API 请求的时间。SDK 默认当前时间戳。
	//
	// 例如1529223702，单位为秒。注意：如果与服务器时间相差超过5分钟，会引起签名过期错误。
	Timestamp string
	Action    string // 操作的接口名称。注意：某些接口不需要传递该参数，接口文档中会对此特别说明，此时即使传递该参数也不会生效。
	Region    string // 地域参数，用来标识希望操作哪个地域的数据。注意：某些接口不需要传递该参数，接口文档中会对此特别说明，此时即使传递该参数也不会生效。
	Signature string // 放置由官网的签名方法产生的签名。SDK 默认自动生成，也可替换为自己生成符合标准的签名。

	// 启用账户通讯录，传入值必须为1，创建的会议可出现在用户的会议列表中。
	//
	// 启用账户通讯录说明：
	//
	// 1. 通过 SSO 接入腾讯会议账号体系。
	//
	// 2. 通过调用接口创建企业用户。
	//
	// 3. 通过企业管理后台添加或批量导入企业用户。
	Registered string
}

func (jwt JWTAuthenticator) AuthHeader(req *http.Request) error {

	if jwt.config == nil {
		return errors.New("JWT authenticator is not available")
	}

	if req.Header.Get(HTTPHeaderContentType) == "" {
		req.Header.Set(HTTPHeaderContentType, DefaultContentType)
	}

	if jwt.config.SecretID == "" {
		return errors.New("JWT authenticator SecretId can't be empty")
	}
	req.Header.Set(XTCHeaderKey, jwt.config.SecretID)
	if jwt.config.AppId == "" {
		return errors.New("JWT authenticator AppId can't be empty")
	}
	req.Header.Set(HeaderAppId, jwt.config.AppId)
	req.Header.Set(HeaderSdkId, jwt.config.SdkId)

	if jwt.Nonce == 0 {
		ts := time.Now().UnixMilli()
		rand.Seed(time.Now().UnixNano())
		rs := fmt.Sprintf("%06d", rand.Intn(1000000))
		ri, _ := strconv.Atoi(rs)
		jwt.Nonce = uint64(ts)*1000000 + uint64(ri)
	}
	req.Header.Set(XTCHeaderNonce, fmt.Sprint(jwt.Nonce))

	if jwt.Registered == "" {
		jwt.Registered = "1"
	}
	req.Header.Set(XTCHeaderRegistered, jwt.Registered)

	if jwt.config.Version != "" {
		req.Header.Set(XTCHeaderVersion, jwt.config.Version)
	}

	if jwt.Action != "" {
		req.Header.Set(XTCHeaderAction, jwt.Action)
	}

	if jwt.Region != "" {
		req.Header.Set(XTCHeaderRegion, jwt.Region)
	}

	if jwt.Timestamp == "" {
		jwt.Timestamp = fmt.Sprint(time.Now().Unix())
	}
	req.Header.Set(XTCHeaderTimestamp, jwt.Timestamp)

	if jwt.Signature == "" {
		jwt.Signature = jwt.signature(req)
	}
	req.Header.Set(XTCHeaderSignature, jwt.Signature)
	return nil
}

// signature 生成签名
func (jwt JWTAuthenticator) signature(req *http.Request) string {

	var reqBody = ""
	if req.Body != nil {
		reqBody = string(jwt.readRequestBody(req))
	}

	signStr := fmt.Sprintf("%s\nX-TC-Key=%s&X-TC-Nonce=%d&X-TC-Timestamp=%s\n%s\n%s",
		req.Method, jwt.config.SecretID, jwt.Nonce, jwt.Timestamp, req.URL.RequestURI(), reqBody)
	signer := hmac.New(sha256.New, []byte(jwt.config.SecretKey))
	signer.Write([]byte(signStr))
	sign := hex.EncodeToString(signer.Sum(nil))
	return base64.StdEncoding.EncodeToString([]byte(sign))
}

func (jwt JWTAuthenticator) readRequestBody(req *http.Request) []byte {
	// 储存 body 的 buf
	var buf bytes.Buffer
	tee := io.TeeReader(req.Body, &buf)
	data, _ := ioutil.ReadAll(tee)
	_ = req.Body.Close()
	// 恢复 body
	req.Body = ioutil.NopCloser(&buf)
	req.ContentLength = int64(buf.Len())
	return data
}

// OAuth2Authenticator oauth2 鉴权器
type OAuth2Authenticator struct {
	config *Config

	Nonce uint64 // 此参数参与签名计算。随机正整数。SDK 默认自动随机生成。

	// 此参数参与签名计算。当前 UNIX 时间戳，可记录发起 API 请求的时间。SDK 默认当前时间戳。
	//
	// 例如1529223702，单位为秒。注意：如果与服务器时间相差超过5分钟，会引起签名过期错误。
	Timestamp string
	Action    string // 操作的接口名称。注意：某些接口不需要传递该参数，接口文档中会对此特别说明，此时即使传递该参数也不会生效。
	Region    string // 地域参数，用来标识希望操作哪个地域的数据。注意：某些接口不需要传递该参数，接口文档中会对此特别说明，此时即使传递该参数也不会生效。

	AccessToken string // OAuth2.0 鉴权成功后返回的 token 信息。
	OpenId      string // OAuth2.0 鉴权成功后的用户信息。
}

func (oauth2 OAuth2Authenticator) AuthHeader(httpReq *http.Request) error {

	if oauth2.config == nil {
		return errors.New("OAuth2 authenticator is not available")
	}

	if httpReq.Header.Get(HTTPHeaderContentType) == "" {
		httpReq.Header.Set(HTTPHeaderContentType, DefaultContentType)
	}

	if oauth2.AccessToken == "" {
		return errors.New("OAuth2 authenticator AccessToken can't be empty")
	}
	httpReq.Header.Set(HeaderAccessToken, oauth2.AccessToken)
	if oauth2.OpenId == "" {
		return errors.New("OAuth2 authenticator OpenId can't be empty")
	}
	httpReq.Header.Set(HeaderOpenId, oauth2.OpenId)
	if oauth2.Nonce == 0 {
		ts := time.Now().UnixMilli()
		rand.Seed(time.Now().UnixNano())
		rs := fmt.Sprintf("%06d", rand.Intn(1000000))
		ri, _ := strconv.Atoi(rs)
		oauth2.Nonce = uint64(ts)*1000000 + uint64(ri)
	}
	httpReq.Header.Set(XTCHeaderNonce, fmt.Sprint(oauth2.Nonce))

	if oauth2.config.Version != "" {
		httpReq.Header.Set(XTCHeaderVersion, oauth2.config.Version)
	}

	if oauth2.Action != "" {
		httpReq.Header.Set(XTCHeaderAction, oauth2.Action)
	}

	if oauth2.Region != "" {
		httpReq.Header.Set(XTCHeaderRegion, oauth2.Region)
	}

	if oauth2.Timestamp == "" {
		oauth2.Timestamp = fmt.Sprint(time.Now().Unix())
	}
	httpReq.Header.Set(XTCHeaderTimestamp, oauth2.Timestamp)
	return nil
}
