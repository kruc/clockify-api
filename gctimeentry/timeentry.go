package gctimeentry

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/kruc/clockify-api/gchttp"
)

// TimeFormat of start and end dates
const TimeFormat = "2006-01-02T15:04:05Z"

// QueryParameters of getRange method
type QueryParameters struct {
	Start    time.Time
	End      time.Time
	Hydrated bool
	PageSize int
	Project  string
}

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
func (tc *TimeEntryClient) GetRange(qp QueryParameters, workspaceID string, userID string) (TimeEntries, error) {
	v := parseQueryParameters(qp)
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

func parseQueryParameters(qp QueryParameters) url.Values {

	v := url.Values{}
	v.Set("start", qp.Start.Format(TimeFormat))
	v.Set("end", qp.End.Format(TimeFormat))
	v.Set("hydrated", strconv.FormatBool(qp.Hydrated))
	v.Set("page-size", strconv.Itoa(qp.PageSize))
	if qp.Project != "" {
		v.Set("project", qp.Project)
	}

	return v
}
