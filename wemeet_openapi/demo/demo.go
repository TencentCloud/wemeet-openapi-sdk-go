package demo

import (
	"context"
	"fmt"
	wemeet "github.com/TencentCloud/wemeet-openapi-sdk-go/wemeet_openapi"
	wemeetcore "github.com/TencentCloud/wemeet-openapi-sdk-go/wemeet_openapi/core"
	meetings "github.com/TencentCloud/wemeet-openapi-sdk-go/wemeet_openapi/service/meetings"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func Demo() {
	// 1.构造 client 客户端(jwt 鉴权需要配置 appId sdkId secretID 和 secretKey)
	client := wemeet.NewClient(wemeet.WithAppId("2****46"), wemeet.WithSdkId("2****50"),
		wemeet.WithSecret("Zk*****J8h", "Y2z*****WRsVksn"))
	endTime := strconv.FormatInt(time.Now().Unix()+3600, 10)
	startTime := strconv.FormatInt(time.Now().Unix(), 10)
	subject := "测试会议"
	meetType := int64(1)
	instanceid := int64(1)
	userid := "userid"

	// 2.构造请求体
	bodyReq := &meetings.V1MeetingsPostRequest{
		EndTime:    endTime,
		StartTime:  startTime,
		Subject:    subject,
		Type:       meetType,
		Instanceid: instanceid,
		Userid:     userid,
	}

	request := &meetings.ApiV1MeetingsPostRequest{
		Body: bodyReq, // V1MeetingsPostRequest ｜
	}

	// 3.构造 JWT 鉴权器
	// 随机数
	rn := rand.New(rand.NewSource(time.Now().UnixNano()))
	nonce := uint64(100000 + rn.Intn(900000))
	// 当前时间戳
	curTs := strconv.Itoa(int(time.Now().Unix()))
	authenticator := &wemeetcore.JWTAuthenticator{
		Nonce:     nonce,
		Timestamp: curTs,
	}

	// 4.发送对应的请求
	response, err := client.MeetingsApi.V1MeetingsPost(context.Background(), request,
		wemeetcore.WithJWTAuth(authenticator))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeetingsAPI.V1MeetingsPost`: %v\n", err)
		if svrErr, ok := err.(*wemeetcore.ServiceError); ok {
			fmt.Fprintf(os.Stderr, "Full HTTP response: %s\n", svrErr.RawBody)
		}
	}

	// response from `V1MeetingsPost`: V1MeetingsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MeetingsAPI.V1MeetingsPost`: %+v\n", response)

	return
}
