package gcproject

import (
	"encoding/json"
	"fmt"

	"github.com/kruc/clockify-api/gchttp"
)

// NewClient return a Project Cilent. An error is also returned when some configuration option is invalid
//    clockify,err := clockifyapi.NewClient("token")
//    projectClient := clockify.ProjectClient
func NewClient(cc *gchttp.ClockifyHTTPClient) *ProjectClient {
	pc := &ProjectClient{
		cc: cc,
	}
	pc.endpoint = cc.URL
	return pc
}

// FindProjectsOnWorkspace get projects from specific workspace
func (pc *ProjectClient) FindProjectsOnWorkspace(workspaceID string) (Projects, error) {

	body, err := pc.cc.GetRequest(fmt.Sprintf("%s/workspaces/%s/projects", pc.endpoint, workspaceID))

	var projects Projects

	if err != nil {
		return projects, err
	}
	if body == nil {
		return nil, nil
	}
	err = json.Unmarshal(*body, &projects)

	return projects, err
}
