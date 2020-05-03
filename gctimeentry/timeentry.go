package gctimeentry

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/kruc/clockify-api/gchttp"
)

// TimeFormat of start and end dates
const TimeFormat = "2006-01-02T15:04:05Z"

// NewClient return a TimeEntry Cilent. An error is also returned when some configuration option is invalid
//    clockify, err := clockifyapi.NewClient("token")
//    timeEntryClient, err := clockify.TimeEntryClient
func NewClient(cc *gchttp.ClockifyHTTPClient) *TimeEntryClient {
	tc := &TimeEntryClient{
		cc: cc,
	}
	tc.endpoint = cc.URL
	return tc
}

// GetRange get timeentries from specific time range
func (tc *TimeEntryClient) GetRange(start time.Time, end time.Time, workspaceID string, userID string) (TimeEntries, error) {
	v := url.Values{}
	v.Set("start", start.Format(TimeFormat))
	v.Set("end", end.Format(TimeFormat))
	v.Set("hydrated", "true")
	body, err := tc.cc.GetRequest(fmt.Sprintf("%s/workspaces/%s/user/%s/time-entries?%s", tc.endpoint, workspaceID, userID, v.Encode()))
	var te TimeEntries
	if err != nil {
		return te, err
	}
	if body == nil {
		return nil, nil
	}
	err = json.Unmarshal(*body, &te)

	return te, err
}

//FindTimeEntriesForUserOnWorkspace https://api.clockify.me/api/v1/workspaces/{workspaceId}/user/{userId}/time-entries
func (tc *TimeEntryClient) FindTimeEntriesForUserOnWorkspace(workspaceID, userID string) (TimeEntries, error) {
	body, err := tc.cc.GetRequest(fmt.Sprintf("%s/workspaces/%s/user/%s/time-entries", tc.endpoint, workspaceID, userID))

	var timeEntries TimeEntries
	if err != nil {
		return timeEntries, err
	}
	if body == nil {
		return nil, nil
	}
	err = json.Unmarshal(*body, &timeEntries)
	return timeEntries, err
}

//Update https://api.clockify.me/api/v1/workspaces/{workspaceId}/time-entries/{id}
func (tc *TimeEntryClient) Update(workspaceID, timeEntryID string, timeEntry *TimeEntry) (*TimeEntry, error) {
	body, err := tc.cc.PutRequest(fmt.Sprintf("%s/workspaces/%s/time-entries/%s", tc.endpoint, workspaceID, timeEntryID), timeEntry)

	var te TimeEntry
	if err != nil {
		return &te, err
	}
	if body == nil {
		return nil, nil
	}
	err = json.Unmarshal(*body, &te)
	return &te, err
}
