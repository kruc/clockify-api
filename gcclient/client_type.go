package gcclient

import (
	"strings"

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

// ToMap converts []Client to map[string]Project
func (c *Clients) ToMap() map[string][]Client {
	projectMap := make(map[string][]Client, len(*c))

	for _, client := range *c {
		normalizedClientName := Normalize(client.Name)
		projectMap[normalizedClientName] = append(projectMap[normalizedClientName], client)
	}

	return projectMap
}

// Normalize modify map key to lower case
func Normalize(word string) string {
	return strings.ToLower(word)
}
