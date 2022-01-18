package common

import (
	"errors"
	"net/http"
	"time"
)

type HttpClient struct {
	logger *Logger
}

func NewHttpClient(logger *Logger) *HttpClient {
	return &HttpClient{
		logger,
	}
}

type Header struct {
	Key   string
	Value string
}
type HttpOptions struct {
	Headers []Header
}

func (c *HttpClient) Do(method string, url string, options *HttpOptions) (*http.Response, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: time.Second * 5,
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		c.logger.Errf(err, "http err building request: ")
		return nil, err
	}

	if options != nil {
		for _, header := range options.Headers {
			req.Header.Add(header.Key, header.Value)
		}

	}

	c.logger.Debugf("http request %v %v", method, url)

	resp, err := client.Do(req)

	if err != nil {
		c.logger.Errf(err, "http response err: ")
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		c.logger.Errorf("http invalid status code %v", resp.StatusCode)
		return nil, errors.New("http invalid status code")
	}

	return resp, nil
}
