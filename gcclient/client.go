package gcclient

import (
	"encoding/json"
	"fmt"

	"github.com/kruc/clockify-api/gchttp"
)

// NewClient return a User Cilent. An error is also returned when some configuration option is invalid
//    clockify, err := clockifyapi.NewClient("token")
//    clientClient := clockify.UserClient
func NewClient(cc *gchttp.ClockifyHTTPClient) *ClientClient {
	client := &ClientClient{
		cc: cc,
	}
	client.endpoint = cc.URL
	return client
}

// FindClientsOnWorkspace get clients from specific workspace
func (cc *ClientClient) FindClientsOnWorkspace(workspaceID string) (Clients, error) {

	test := fmt.Sprintf("%s/workspaces/%s/clients", cc.endpoint, workspaceID)
	fmt.Println(test)
	body, err := cc.cc.GetRequest(fmt.Sprintf("%s/workspaces/%s/clients", cc.endpoint, workspaceID))

	var clients Clients

	if err != nil {
		return clients, err
	}
	if body == nil {
		return nil, nil
	}
	err = json.Unmarshal(*body, &clients)

	return clients, err
}
