package gctag

import (
	"encoding/json"
	"fmt"

	"github.com/kruc/clockify-api/gchttp"
)

// NewClient return a Tag Cilent. An error is also returned when some configuration option is invalid
//    clockify, err := clockifyapi.NewClient("token")
//    tagClient := clockify.TagClient
func NewClient(cc *gchttp.ClockifyHTTPClient) *TagClient {
	tc := &TagClient{
		cc: cc,
	}
	tc.endpoint = cc.URL
	return tc
}

//GetTags https://api.clockify.me/api/v1/workspaces/{workspace-id}/tags
func (tc *TagClient) GetTags(workspaceID string) (*Tags, error) {

	body, err := tc.cc.GetRequest(fmt.Sprintf("%s/workspaces/%s/tags", tc.endpoint, workspaceID))
	var tags Tags
	if err != nil {
		return &tags, err
	}
	if body == nil {
		return nil, nil
	}
	err = json.Unmarshal(*body, &tags.Tags)
	return &tags, err
}
