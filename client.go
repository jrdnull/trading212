package trading212

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"time"
)

const (
	liveBaseURL = "https://live.trading212.com"
	demoBaseURL = "https://demo.trading212.com"
)

// Client for interacting with the undocumented Trading 212 website API.
type Client struct {
	baseURL string
	c       *http.Client
}

// Option sets an optional setting on the Client.
type Option func(*Client)

// NewClient returns a new client initialised with opts.
func NewClient(opts ...Option) *Client {
	jar, _ := cookiejar.New(nil)
	client := &Client{liveBaseURL, &http.Client{
		Timeout: 10 * time.Second,
		Jar:     jar,
	}}
	for _, opt := range opts {
		opt(client)
	}
	return client
}

// WithHTTPClient returns an Option to set the http.Client to be used.
func WithHTTPClient(httpClient *http.Client) Option {
	if httpClient.Jar == nil {
		panic("client must have Jar set")
	}
	return func(c *Client) {
		c.c = httpClient
	}
}

// WithPracticeMode returns an Option to set the Client to make requests to the
// demo mode environment for practice.
func WithPracticeMode() Option {
	return func(c *Client) {
		c.baseURL = demoBaseURL
	}
}

func (c *Client) newRequest(
	ctx context.Context, method, endpoint string, body interface{},
) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}

	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		req.Header.Add("Content-Type", "application/json; charset=utf8")
	}
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) error {
	resp, err := c.c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close() // nolint: errcheck

	if resp.StatusCode >= 299 {
		return NewError(resp)
	}

	if v != nil {
		return json.NewDecoder(resp.Body).Decode(v)
	}
	return nil
}
