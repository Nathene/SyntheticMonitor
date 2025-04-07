package httpclient

import (
	"io"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
)

const (
	defaultTimeout = 15 * time.Second

	defaultIdleTimeout = 30 * time.Second

	defaultRetryMax = 3

	defaultRetryWait    = 1 * time.Second
	defaultRetryMaxWait = 15 * time.Second

	defaultMaxIdleConns = 5

	defaultTLSHandshake = 5 * time.Second

	defaultMaxResponseSize = 1 * 1024 * 1024 // 1MB

)

type Client struct {
	*http.Client
	retryMax        int
	retryWaitMin    time.Duration
	retryWaitMax    time.Duration
	maxResponseSize int64
	maxIdleConns    int
	tlsHandshake    time.Duration
	backOffStrategy map[int]BackoffStrategy
}

type BackoffStrategy int

const (
	ConstantBackoff BackoffStrategy = iota
	ExponentialBackoff
)

func New(opts ...Option) *Client {
	transport := &http.Transport{
		Proxy:                  http.ProxyFromEnvironment,
		MaxIdleConns:           defaultMaxIdleConns,
		IdleConnTimeout:        defaultIdleTimeout,
		TLSHandshakeTimeout:    defaultTLSHandshake,
		MaxResponseHeaderBytes: defaultMaxResponseSize,
	}
	c := &Client{
		Client: &http.Client{
			Transport: otelhttp.NewTransport(transport),
			Timeout:   defaultTimeout,
		},
		retryMax:        defaultRetryMax,
		retryWaitMin:    defaultRetryWait,
		retryWaitMax:    defaultRetryMaxWait,
		maxResponseSize: defaultMaxResponseSize,
		maxIdleConns:    defaultMaxIdleConns,
		tlsHandshake:    defaultTLSHandshake,
		backOffStrategy: map[int]BackoffStrategy{
			http.StatusTooManyRequests:     ExponentialBackoff,
			http.StatusServiceUnavailable:  ConstantBackoff,
			http.StatusGatewayTimeout:      ConstantBackoff,
			http.StatusBadGateway:          ConstantBackoff,
			http.StatusInternalServerError: ConstantBackoff,
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	// Implement retry logic and backoff strategy here
	// ...
	tracer := otel.Tracer("httpclient")
	ctx, span := tracer.Start(req.Context(), "httpclient.Do")
	defer span.End()

	req = req.WithContext(ctx)

	wait := c.retryWaitMin

	for i := range c.retryMax {
		resp, err := c.Client.Do(req)

		_, ok := c.backOffStrategy[resp.StatusCode]
		if err == nil && !ok {
			break
		}
		if err != nil {
			if i == c.retryMax {
				return nil, err
			}

			select {
			case <-req.Context().Done():
				return nil, req.Context().Err()
			case <-time.After(wait):
			}

			wait = c.getNextWait(resp.StatusCode, wait)
			continue
		}

		if ok {
			switch c.backOffStrategy[resp.StatusCode] {
			case ConstantBackoff:
				wait = c.retryWaitMin
			case ExponentialBackoff:
				wait *= 2
				if wait > c.retryWaitMax {
					wait = c.retryWaitMax
				}
			}
		}
	}
	return c.Client.Do(req)
}

func (c *Client) getNextWait(code int, wait time.Duration) time.Duration {
	strategy, ok := c.backOffStrategy[code]
	if !ok {
		return c.retryWaitMin
	}
	switch strategy {
	case ConstantBackoff:
		return c.retryWaitMin
	case ExponentialBackoff:
		wait *= 2
	}

	if wait > c.retryWaitMax {
		return c.retryWaitMax
	}
	return wait
}

type Option func(*Client)

func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.Timeout = timeout
	}
}
func WithRetryMax(max int) Option {
	return func(c *Client) {
		c.retryMax = max
	}
}

func WithRetryWaitMin(wait time.Duration) Option {
	return func(c *Client) {
		c.retryWaitMin = wait
	}
}

func WithRetryWaitMax(wait time.Duration) Option {
	return func(c *Client) {
		c.retryWaitMax = wait
	}
}

func WithMaxIdleConns(max int) Option {
	return func(c *Client) {
		c.maxIdleConns = max
	}
}

func WithTLSHandshake(timeout time.Duration) Option {
	return func(c *Client) {
		c.tlsHandshake = timeout
	}
}

func WithMaxResponseSize(size int64) Option {
	return func(c *Client) {
		c.maxResponseSize = size
	}
}

func WithBackOffStrategy(code int, strategy BackoffStrategy) Option {
	return func(c *Client) {
		c.backOffStrategy[code] = strategy
	}
}

func (c *Client) LimitedReader(body io.ReadCloser) io.ReadCloser {
	return &limitedReader{
		Reader: io.LimitReader(body, c.maxResponseSize),
		Closer: body,
	}
}

type limitedReader struct {
	io.Reader
	io.Closer
}

func (lr *limitedReader) Read(p []byte) (n int, err error) {
	return lr.Reader.Read(p)
}

func (lr *limitedReader) Close() error {
	if closer, ok := lr.Reader.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

func (c *Client) limitedReader(r io.Reader) io.Reader {
	if c.maxResponseSize > 0 {
		return &limitedReader{
			Reader: io.LimitReader(r, c.maxResponseSize),
			Closer: r.(io.Closer),
		}
	}
	return r
}
