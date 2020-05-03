package gchttp

import (
	"net/http"

	"github.com/throttled/throttled"
)

// ClockifyHTTPClient is an Cloudify REST client. Created by calling NewClient.
type ClockifyHTTPClient struct {
	client      *http.Client // net/http Client to use for requests
	version     string       // v8
	URL         string       // set of URLs passed initially to the client
	errorLog    Logger       // error log for critical messages
	infoLog     Logger       // information log for e.g. response times
	traceLog    Logger       // trace log for debugging
	token       string       // auth token
	maxRetries  uint
	rateLimiter *throttled.GCRARateLimiter
	perSec      int
	cookie      *http.Cookie
}

// ClockifyError type
type ClockifyError struct {
	Code   int
	Status string
	Msg    string
}

// ClientOptionFunc is a function that configures a Client.
// It is used in NewClient.
type ClientOptionFunc func(*ClockifyHTTPClient) error

type nullLogger struct{}
