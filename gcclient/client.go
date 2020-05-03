package gcclient

import (
	"github.com/kruc/clockify-api/gchttp"
)

// NewClient return a User Cilent. An error is also returned when some configuration option is invalid
//    clockify, err := clockifyapi.NewClient("token")
//    userClient := clockify.UserClient
func NewClient(cc *gchttp.ClockifyHTTPClient) *ClientClient {
	client := &ClientClient{
		cc: cc,
	}
	client.endpoint = cc.URL
	return client
}
