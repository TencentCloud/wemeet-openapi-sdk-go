package wemeethttp

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Client 封装 REST 标准客户端接口
type Client interface {
	Get(ctx context.Context, req *ApiRequest, opts ...RequestOptionFunc) (resp *ApiResponse, err error)
	Post(ctx context.Context, req *ApiRequest, opts ...RequestOptionFunc) (resp *ApiResponse, err error)
	Put(ctx context.Context, req *ApiRequest, opts ...RequestOptionFunc) (resp *ApiResponse, err error)
	Delete(ctx context.Context, req *ApiRequest, opts ...RequestOptionFunc) (resp *ApiResponse, err error)
}

// Serializable 序列化能力
type Serializable interface {
	Name() string
	ContentType() string
	Serialize(v interface{}) ([]byte, error)
	Deserialize(data []byte, v interface{}) error
}

// Authentication 鉴权能力
type Authentication interface {
	AuthHeader(httpReq *http.Request) error
}

type client struct {
	clt      *http.Client
	host     string
	protocol string

	serializer Serializable
}

type ClientOptionFunc func(c *client)

func WithSerializer(serializer Serializable) ClientOptionFunc {
	return func(c *client) {
		c.serializer = serializer
	}
}

func WithProtocol(protocol string) ClientOptionFunc {
	return func(c *client) {
		c.protocol = protocol
	}
}

func WithClient(httpClt *http.Client) ClientOptionFunc {
	return func(c *client) {
		c.clt = httpClt
	}
}

func NewClient(host string, opts ...ClientOptionFunc) (Client, error) {
	clt := &client{
		clt:      http.DefaultClient,
		host:     host,
		protocol: "http",
	}

	for _, opt := range opts {
		opt(clt)
	}

	if clt.host == "" {
		return nil, errors.New("http client's host can't be empty")
	}

	if _, err := url.Parse(fmt.Sprintf("%s://%s", clt.protocol, clt.host)); err != nil {
		return nil, fmt.Errorf("http client's protocol or host is illegal")
	}

	return clt, nil
}

func (c *client) Get(ctx context.Context, req *ApiRequest,
	opts ...RequestOptionFunc) (*ApiResponse, error) {
	// 处理空值
	if req == nil {
		return nil, fmt.Errorf("http client do request error, 'req' is nil")
	}

	return c.doRequest(ctx, req, http.MethodGet, opts...)
}

func (c *client) Put(ctx context.Context, req *ApiRequest,
	opts ...RequestOptionFunc) (*ApiResponse, error) {
	// 处理空值
	if req == nil {
		return nil, fmt.Errorf("http client do request error, 'req' is nil")
	}

	return c.doRequest(ctx, req, http.MethodPut, opts...)
}

func (c *client) Delete(ctx context.Context, req *ApiRequest,
	opts ...RequestOptionFunc) (*ApiResponse, error) {
	// 处理空值
	if req == nil {
		return nil, fmt.Errorf("http client do request error, 'req' is nil")
	}

	return c.doRequest(ctx, req, http.MethodDelete, opts...)
}

func (c *client) Post(ctx context.Context, req *ApiRequest,
	opts ...RequestOptionFunc) (*ApiResponse, error) {
	// 处理空值
	if req == nil {
		return nil, fmt.Errorf("http client do request error, 'req' is nil")
	}

	return c.doRequest(ctx, req, http.MethodPost, opts...)
}

func (c *client) doRequest(ctx context.Context, req *ApiRequest, method string,
	opts ...RequestOptionFunc) (*ApiResponse, error) {

	for _, opt := range opts {
		opt(req)
	}

	// 获取序列化器，以当前请求配置的为准
	serializer := c.serializer
	if req.serializer != nil {
		serializer = req.serializer
	}

	// 序列化请求体
	var bodyReader io.Reader
	if req.Body != nil {
		if serializer != nil {
			data, err := serializer.Serialize(req.Body)
			if err != nil {
				return nil, err
			}
			bodyReader = bytes.NewBuffer(data)
		} else if reader, ok := req.Body.(io.Reader); ok {
			bodyReader = reader
		}
	}

	// 生成 url
	u, err := req.GenerateURL(fmt.Sprintf("%s://%s", c.protocol, c.host))
	if err != nil {
		return nil, err
	}
	// 构造原生 http request 请求
	httpReq, err := http.NewRequestWithContext(ctx, method, u.String(), bodyReader)
	if err != nil {
		return nil, err
	}

	// 设置请求头
	for k, v := range req.header {
		httpReq.Header[k] = v
	}
	// 增加鉴权头
	if req.authenticators != nil {
		for _, authenticator := range req.authenticators {
			if err = authenticator.AuthHeader(httpReq); err != nil {
				return nil, err
			}
		}
	}

	// 获取原生 http client
	if c.clt == nil {
		c.clt = http.DefaultClient
	}

	// 执行发送请求
	httpRsp, err := c.clt.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer func() {
		if httpRsp.Body != nil {
			_ = httpRsp.Body.Close()
		}
	}()

	// 读取响应
	respBody, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	// 封装响应返回
	return &ApiResponse{
		StatusCode: httpRsp.StatusCode,
		Header:     httpRsp.Header,
		RawBody:    respBody,
		serializer: serializer,
	}, nil
}
