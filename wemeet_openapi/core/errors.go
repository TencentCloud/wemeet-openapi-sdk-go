package wemeetcore

import (
	"fmt"
    "strings"

	xhttp "github.com/TencentCloud/wemeet-openapi-sdk-go/wemeet_openapi/core/xhttp"
)

// ClientError 客户端异常
type ClientError struct {
	Err error
}

func (ce ClientError) Error() string {
	return fmt.Sprintf("[wemeet client error] %s", ce.Err.Error())
}

// ServiceError 接口服务异常
type ServiceError struct {
	*xhttp.ApiResponse

	// ErrorInfo 业务错误
	ErrorInfo *struct {
		ErrorCode    int64  `json:"error_code"`
		NewErrorCode int64  `json:"new_error_code"`
		Message      string `json:"message"`
	} `json:"error_info"`
}

func (se ServiceError) Error() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("http status code: %d, ", se.StatusCode))
	if se.ErrorInfo != nil {
		sb.WriteString(fmt.Sprintf("error code: %d, new error code: %d, message: %s",
			se.ErrorInfo.ErrorCode, se.ErrorInfo.NewErrorCode, se.ErrorInfo.Message))
	} else {
		sb.WriteString(fmt.Sprintf("response body: %s", se.RawBody))
	}
	return fmt.Sprintf("[wemeet service error] %s", sb.String())
}
