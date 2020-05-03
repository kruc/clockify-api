package gcproject

import (
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
