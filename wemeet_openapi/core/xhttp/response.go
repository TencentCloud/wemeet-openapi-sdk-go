package wemeethttp

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
)

type ApiResponse struct {
	StatusCode int
	Header     http.Header
	RawBody    []byte

	serializer Serializable
}

func (resp ApiResponse) Write(writer http.ResponseWriter) error {
	writer.WriteHeader(resp.StatusCode)
	for k, vs := range resp.Header {
		for _, v := range vs {
			writer.Header().Add(k, v)
		}
	}
	if _, err := writer.Write(resp.RawBody); err != nil {
		return err
	}
	return nil
}

func (resp ApiResponse) Translate(dst interface{}, opts ...ResponseOptionFunc) error {

	copyResp := &ApiResponse{
		RawBody:    make([]byte, len(resp.RawBody)),
		serializer: resp.serializer, // 以 opts 中传入配置的为准，会被传入的 serializer 覆盖
	}
	copy(copyResp.RawBody, resp.RawBody)

	for _, opt := range opts {
		opt(copyResp)
	}

	v := reflect.ValueOf(dst)
	if v.Kind() != reflect.Pointer || v.IsNil() {
		return fmt.Errorf("response body translate error, dst is nil or dst's type is not the pointer")
	}
	if copyResp.serializer == nil {
		return fmt.Errorf("response body translate error, serializer is nil")
	}
	if err := copyResp.serializer.Deserialize(copyResp.RawBody, dst); err != nil {
		return errors.Join(
			fmt.Errorf("response body translate error, "+
				"body can't be translated by '%s' serializer", copyResp.serializer.Name()),
			err)
	}
	return nil
}

type ResponseOptionFunc func(resp *ApiResponse)

func WithResponseSerializer(serializer Serializable) ResponseOptionFunc {
	return func(response *ApiResponse) {
		response.serializer = serializer
	}
}
