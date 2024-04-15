# 简介
欢迎使用腾讯会议开发者工具套件（SDK），为方便 GO 开发者调试和接入腾讯云会议 API，这里向您介绍适用于 GO 的腾讯会议开发工具包，并提供首次使用开发工具包的简单示例。让您快速获取腾讯会议 GO SDK 并开始调用。
# 依赖环境
1. 购买腾讯会议企业版获取 SecretID、SecretKey，接入的企业 AppId。

# 获取安装
在第一次使用云API之前，用户首先需要在腾讯云控制台上申请安全凭证，安全凭证包括 SecretID 和 SecretKey，SecretID 是用于标识 API 调用者的身份，SecretKey 是用于加密签名字符串和服务器端验证签名字符串的密钥 SecretKey 必须严格保管，避免泄露。
## 安装
```shell
require github.com/TencentCloud/wemeet-openapi-sdk-go
```

# 示例

以创建会议接口为例：
```GO
package example

import (
    wemeet "github.com/TencentCloud/wemeet-openapi-sdk-go/wemeet_openapi"
    wemeetcore "github.com/TencentCloud/wemeet-openapi-sdk-go/wemeet_openapi/core"
    wemeethttp "github.com/TencentCloud/wemeet-openapi-sdk-go/wemeet_openapi/core/xhttp"
    meetings "github.com/TencentCloud/wemeet-openapi-sdk-go/wemeet_openapi/service/meetings"
    
    "context"
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "time"
)

// CreateMeetingDemo 创建会议请求demo
func CreateMeetingDemo() {
    
    // 1.构造 client 客户端(jwt 鉴权需要配置 appId sdkId secretID 和 secretKey)
    client := wemeet.NewClient(wemeet.WithAppId("2****46"), wemeet.WithSdkId("2****50"),
        wemeet.WithSecret("Zk*****J8h", "Y2z*****WRsVksn"))
    
    // 2.构造请求体
    request := &meetings.ApiV1MeetingsPostRequest{
        Body: &meetings.V1MeetingsPostRequest{
            Instanceid:  1,
            MeetingType: wemeetcore.PtrInt32(0),
            Subject:     "测试会议",
            Type:        1,
            Userid:      "userid",
            StartTime:   "1651334400",
            EndTime:     "1651377600",
        },
    }
    
    // 3.构造 JWT 鉴权器
    // 随机数
    rn := rand.New(rand.NewSource(time.Now().UnixNano()))
    nonce := uint64(100000 + rn.Intn(900000))
    // 当前时间戳
    curTs := strconv.Itoa(int(time.Now().Unix()))
    authenticator := &wemeetcore.JWTAuthenticator{
        Nonce: nonce,
        Timestamp: curTs,
    }
    
    // 4.发送对应的请求
    response, err := client.MeetingsApi.V1MeetingsPost(context.Background(), request,
        wemeetcore.WithJWTAuth(authenticator))
    
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeetingsApi.V1MeetingsPost`: %v\n", err)
        if svrErr, ok := err.(*wemeetcore.ServiceError); ok {
            fmt.Fprintf(os.Stderr, "Full HTTP response: %s\n", svrErr.RawBody)
        }
    }
    // response from `V1MeetingsPost`: V1MeetingsPostResponse200
    fmt.Fprintf(os.Stdout, "Response from `MeetingsApi.V1MeetingsPost`: %+v\n", response)
    
    return
}
```

