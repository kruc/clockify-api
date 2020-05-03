package gcuser

import (
	"encoding/json"
	"fmt"

	"github.com/kruc/clockify-api/gchttp"
)

// NewClient return a User Cilent. An error is also returned when some configuration option is invalid
//    clockify, err := clockifyapi.NewClient("token")
//    userClient := clockify.UserClient
func NewClient(cc *gchttp.ClockifyHTTPClient) *UserClient {
	uc := &UserClient{
		cc: cc,
	}
	uc.endpoint = cc.URL
	return uc
}

//GetCurrentlyLoggedInUser https://api.clockify.me/api/v1/user
func (uc *UserClient) GetCurrentlyLoggedInUser() (*User, error) {
	body, err := uc.cc.GetRequest(fmt.Sprintf("%s/user", uc.endpoint))

	var user User
	if err != nil {
		return &user, err
	}
	if body == nil {
		return nil, nil
	}
	err = json.Unmarshal(*body, &user)
	return &user, err
}
