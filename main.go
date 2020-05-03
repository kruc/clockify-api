package clockifyapi

import (
	"github.com/kruc/clockify-api/gcclient"
	"github.com/kruc/clockify-api/gchttp"
	"github.com/kruc/clockify-api/gcproject"
	"github.com/kruc/clockify-api/gctag"
	"github.com/kruc/clockify-api/gctimeentry"
	"github.com/kruc/clockify-api/gcuser"
	"github.com/kruc/clockify-api/gcworkspace"
)

// ClockifyClient is an Clockify REST client. Created by calling NewClient.
type ClockifyClient struct {
	ClockifyHTTPClient *gchttp.ClockifyHTTPClient
	TimeEntryClient    *gctimeentry.TimeEntryClient
	UserClient         *gcuser.UserClient
	TagClient          *gctag.TagClient
	WorkspaceClient    *gcworkspace.WorkspaceClient
	ClientClient       *gcclient.ClientClient
	ProjectClient      *gcproject.ProjectClient
}

// NewClient return a new ClockifyHttpClient . An error is also returned when some configuration option is invalid
//    tc,err := clockifyapi.NewClient("token")
func NewClient(key string, options ...gchttp.ClientOptionFunc) (*ClockifyClient, error) {
	// Set up the client
	c, err := gchttp.NewClient(key, options...)
	if err != nil {
		return nil, err
	}
	th := &ClockifyClient{
		ClockifyHTTPClient: c,
		TimeEntryClient:    gctimeentry.NewClient(c),
		UserClient:         gcuser.NewClient(c),
		TagClient:          gctag.NewClient(c),
		WorkspaceClient:    gcworkspace.NewClient(c),
	}
	// Run the options on it
	for _, option := range options {
		if err := option(c); err != nil {
			return nil, err
		}
	}

	return th, nil
}
