package gcworkspace

import (
	"encoding/json"
	"fmt"

	"github.com/kruc/clockify-api/gchttp"
)

// NewClient return a Workspace Cilent. An error is also returned when some configuration option is invalid
//    clockify, err := clockifyapi.NewClient("token")
//    workspaceClient, err := clockify.WorkspaceClient
func NewClient(cc *gchttp.ClockifyHTTPClient) *WorkspaceClient {
	wc := &WorkspaceClient{
		cc: cc,
	}
	wc.endpoint = cc.URL
	return wc
}

//List https://api.clockify.me/api/v1/workspaces
func (wc *WorkspaceClient) List() (Workspaces, error) {
	body, err := wc.cc.GetRequest(fmt.Sprintf("%s/workspaces", wc.endpoint))
	var workspaces Workspaces
	if err != nil {
		return workspaces, err
	}
	if body == nil {
		return nil, nil
	}
	err = json.Unmarshal(*body, &workspaces)
	return workspaces, err
}
