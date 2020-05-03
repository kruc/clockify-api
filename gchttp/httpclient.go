package gchttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/throttled/throttled"
	"github.com/throttled/throttled/store/memstore"
)

// Default const values
const (
	DefaultToken              = "api_token"
	DefaultURL                = "https://api.clockify.me/api/v1"
	DefaultMaxRetries         = 5
	DefaultVersion            = "v1"
	defaultBucket             = "clockify"
	DefaultRateLimitPerSecond = 3
)

func (e *ClockifyError) Error() string {
	return fmt.Sprintf("%s\t%s\n", e.Status, e.Msg)
}

// NewClient return a new CloudifyHTTPClient . An error is also returned when some configuration option is invalid
//    clockify ,err := clockifyapi.NewClient("token")
func NewClient(key string, options ...ClientOptionFunc) (*ClockifyHTTPClient, error) {
	c := &ClockifyHTTPClient{
		client:     http.DefaultClient,
		maxRetries: DefaultMaxRetries,
		URL:        DefaultURL,
		version:    DefaultVersion,
		token:      DefaultToken,
		errorLog:   defaultLogger,
		infoLog:    defaultLogger,
		traceLog:   defaultLogger,
	}

	err := SetRateLimit(DefaultRateLimitPerSecond)(c)
	if err != nil {
		return nil, err
	}

	// Run the options on it
	for _, option := range options {
		if err := option(c); err != nil {
			return nil, err
		}
	}
	c.infoLog.Printf("Logging in with token: %s\n", key)

	if len(key) < 1 {
		c.errorLog.Printf("%s\n", "valid token required")
		return nil, errors.New("Token required")
	}
	c.token = key

	return c, nil
}

// SetRateLimit Set custom rate limit per second
func SetRateLimit(perSec int) ClientOptionFunc {
	return func(c *ClockifyHTTPClient) error {
		store, err := memstore.New(65536)
		if err != nil {
			return err
		}
		quota := throttled.RateQuota{throttled.PerSec(perSec), 1}
		c.rateLimiter, err = throttled.NewGCRARateLimiter(store, quota)
		if err != nil {
			return err
		}
		c.perSec = perSec
		return nil
	}
}

// SetHTTPClient can be used to specify the http.Client to use when making
// HTTP requests to Toggl
func SetHTTPClient(httpClient *http.Client) ClientOptionFunc {
	return func(c *ClockifyHTTPClient) error {
		if httpClient != nil {
			c.client = httpClient
		} else {
			c.client = http.DefaultClient
		}
		return nil
	}
}

// SetURL defines the base URL. See DefaultURL
func SetURL(url string) ClientOptionFunc {
	return func(c *ClockifyHTTPClient) error {
		c.URL = url
		return nil
	}
}

//SetTraceLogger logger to print HTTP requests
func SetTraceLogger(l Logger) ClientOptionFunc {
	return func(c *ClockifyHTTPClient) error {
		c.traceLog = l
		return nil
	}
}

//SetErrorLogger logger to handle error messages
func SetErrorLogger(l Logger) ClientOptionFunc {
	return func(c *ClockifyHTTPClient) error {
		c.errorLog = l
		return nil
	}
}

//SetInfoLogger logger to handle info messages
func SetInfoLogger(l Logger) ClientOptionFunc {
	return func(c *ClockifyHTTPClient) error {
		c.infoLog = l
		return nil
	}
}

func (l *nullLogger) Printf(format string, v ...interface{}) {
}

var defaultLogger = &nullLogger{}

func requestWithLimit(c *ClockifyHTTPClient, method, endpoint string, b interface{}, attempt int) (*json.RawMessage, error) {
	c.infoLog.Printf("Request attempt %d for %s %s\n", attempt, method, endpoint)
	if attempt > DefaultMaxRetries {
		return nil, errors.New("Max Retries exceeded: " + strconv.FormatInt(DefaultMaxRetries, 10))
	}
	var body []byte
	var err error

	limited, reason, err := c.rateLimiter.RateLimit(defaultBucket, 1)
	if err != nil {
		return nil, err
	}

	if limited {
		c.traceLog.Printf("Hit rate limit. Sleeping for %f ms.\n", float64(reason.RetryAfter)/1000000)
		time.Sleep(reason.RetryAfter)
		return requestWithLimit(c, method, endpoint, b, attempt+1)
	}

	if b != nil {
		if body, err = json.Marshal(b); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Api-Key", c.token)
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	c.dumpRequest(req)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	c.dumpResponse(resp)
	defer resp.Body.Close()

	js, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 429 {
		c.errorLog.Printf("Hit (429) rate limit. Sleeping for %d ms.\n", attempt*1000)
		time.Sleep(time.Millisecond * time.Duration(attempt*1000))
		return requestWithLimit(c, method, endpoint, b, attempt+1)
	}
	if resp.StatusCode == 404 {
		return nil, nil
	}
	if resp.StatusCode >= 400 {
		return nil, &ClockifyError{Code: resp.StatusCode, Status: resp.Status, Msg: string(js)}
	}
	var raw json.RawMessage
	if json.Unmarshal(js, &raw) != nil {
		return nil, err
	}
	return &raw, err
}

func request(c *ClockifyHTTPClient, method, endpoint string, b interface{}) (*json.RawMessage, error) {
	return requestWithLimit(c, method, endpoint, b, 1)
}

// GetRequest Utility to GET requests
func (c *ClockifyHTTPClient) GetRequest(endpoint string) (*json.RawMessage, error) {
	return request(c, "GET", endpoint, nil)
}

// PutRequest Utility to PUT requests
func (c *ClockifyHTTPClient) PutRequest(endpoint string, body interface{}) (*json.RawMessage, error) {
	return request(c, "PUT", endpoint, body)
}
