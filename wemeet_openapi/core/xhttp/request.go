package wemeethttp

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type ApiRequest struct {
	ApiURI      string
	Body        interface{}
	PathParams  PathParams
	QueryParams QueryParams

	header         http.Header
	serializer     Serializable
	authenticators []Authentication
}

func (req *ApiRequest) GenerateURL(baseURL string) (*url.URL, error) {
	rawURL := req.ApiURI
	if strings.Index(rawURL, "http") != 0 {
		rawURL = fmt.Sprintf("%s%s", baseURL, rawURL)
	}

	// path
	for name, val := range req.PathParams {
		rawURL = strings.ReplaceAll(rawURL, "{"+name+"}", val)
	}

	// query
	queryPath := req.QueryParams.Encode()
	if queryPath != "" {
		rawURL = fmt.Sprintf("%s?%s", rawURL, queryPath)
	}
	return url.Parse(rawURL)
}

// RequestOptionFunc ...
type RequestOptionFunc func(request *ApiRequest)

func WithRequestHeader(header http.Header) RequestOptionFunc {
	return func(request *ApiRequest) {
		if request.header == nil {
			header = make(http.Header, len(header))
		}
		for k, v := range header {
			request.header[k] = v
		}
	}
}

func WithRequestSerializer(serializer Serializable) RequestOptionFunc {
	return func(request *ApiRequest) {
		request.serializer = serializer
	}
}

func WithRequestAuthenticator(authenticator Authentication) RequestOptionFunc {
	return func(request *ApiRequest) {
		request.authenticators = append(request.authenticators, authenticator)
	}
}

type PathParams map[string]string

func (u PathParams) Get(key string) string {
	vs := u[key]
	if len(vs) == 0 {
		return ""
	}
	return vs
}
func (u PathParams) Set(key, value string) {
	u[key] = value
}

type QueryParams map[string][]string

func (u QueryParams) Get(key string) string {
	vs := u[key]
	if len(vs) == 0 {
		return ""
	}
	return vs[0]
}

func (u QueryParams) GetAll(key string) []string {
	vs := u[key]
	return vs
}

func (u QueryParams) Set(key, value string) {
	u[key] = []string{value}
}

func (u QueryParams) Encode() string {
	return url.Values(u).Encode()
}

func (u QueryParams) Add(key, value string) {
	u[key] = append(u[key], value)
}
