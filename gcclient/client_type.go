package gcclient

import (
	"github.com/kruc/clockify-api/gchttp"
)

// ClientClient type
type ClientClient struct {
	cc       *gchttp.ClockifyHTTPClient
	endpoint string
}

// Client type
type Client struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	WorkspaceID string `json:"workspaceId"`
	Archived    bool   `json:"archived"`
}

// Clients list of Client
type Clients []Client
