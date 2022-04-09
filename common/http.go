package common

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

type IHttpClient interface {
	Do(ctx context.Context, method string, url string, options *HttpOptions) (*http.Response, error)
}

type Header struct {
	Key   string
	Value string
}
type HttpOptions struct {
	Headers []Header
}

type httpClient struct {
	logger ILogger
	client *http.Client
}

func NewHttpClient(logger ILogger) IHttpClient {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	return &httpClient{
		logger,
		client,
	}
}

func (c *httpClient) Do(ctx context.Context, method string, url string, options *HttpOptions) (*http.Response, error) {
	return FunctionRetrier(ctx, func() (*http.Response, error) {
		return c.doInternal(ctx, method, url, options)
	})
}

func (c *httpClient) doInternal(ctx context.Context, method string, url string, options *HttpOptions) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, nil)

	if err != nil {
		c.logger.Errf(ctx, err, "http err building request for url %v: ", url)
		return nil, err
	}

	if options != nil {
		for _, header := range options.Headers {
			req.Header.Add(header.Key, header.Value)
		}

	}

	c.logger.Debugf(ctx, "http request %v %v", method, url)

	resp, err := c.client.Do(req)

	if err != nil {
		c.logger.Errf(ctx, err, "http err response for %v: ", url)
		return nil, fmt.Errorf("http response err %w", err)
	}

	if resp.StatusCode == 429 {
		c.logger.Errorf(ctx, "http rate-limit status code 429 for %v", url)
		return nil, fmt.Errorf("http rate limit %w", ErrRetryable)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		c.logger.Errorf(ctx, "http invalid status code %v for %v", resp.StatusCode, url)
		return nil, errors.New(fmt.Sprintf("http invalid status code %v", resp.StatusCode))
	}

	return resp, nil
}
