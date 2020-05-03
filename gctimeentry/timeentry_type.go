package gctimeentry

import (
	"time"

	"github.com/kruc/clockify-api/gchttp"
	"github.com/kruc/clockify-api/gcproject"
	"github.com/kruc/clockify-api/gctag"
)

// TimeEntryClient type
type TimeEntryClient struct {
	cc       *gchttp.ClockifyHTTPClient
	endpoint string
}

// TimeEntry type
type TimeEntry struct {
	Billable     bool              `json:"billable"`
	Description  string            `json:"description"`
	ID           string            `json:"id"`
	IsLocked     bool              `json:"isLocked"`
	ProjectID    string            `json:"projectId"`
	Project      gcproject.Project `json:"project"`
	TagIds       []string          `json:"tagIds"`
	Tags         []gctag.Tag       `json:"tags"`
	TaskID       string            `json:"taskId"`
	End          time.Time         `json:"end,omitempty"`
	Start        time.Time         `json:"start,omitempty"`
	TimeInterval TimeInterval      `json:"timeInterval"`
	UserID       string            `json:"userId"`
	WorkspaceID  string            `json:"workspaceId"`
}

// TimeInterval type
type TimeInterval struct {
	Duration string    `json:"duration"`
	End      time.Time `json:"end"`
	Start    time.Time `json:"start"`
}

// TimeEntries list of TimeEntry
type TimeEntries []TimeEntry

// IsTagged returns true if timeEntry is tagged with specify tag name
func (te *TimeEntry) IsTagged(tagName string) bool {
	for _, tag := range te.Tags {
		if tag.Name == tagName {
			return true
		}
	}

	return false
}

// AddTag tag to timeEntry
func (te *TimeEntry) AddTag(tagID string) []string {
	for _, ele := range te.TagIds {
		if ele == tagID {
			return te.TagIds
		}
	}
	te.TagIds = append(te.TagIds, tagID)

	return te.TagIds
}

// RemoveTag tag to timeEntry
func (te *TimeEntry) RemoveTag(tagID string) []string {
	for i, ele := range te.TagIds {
		if ele == tagID {
			te.TagIds[i] = te.TagIds[len(te.TagIds)-1]
			te.TagIds[len(te.TagIds)-1] = ""
			te.TagIds = te.TagIds[:len(te.TagIds)-1]
		}
	}

	return te.TagIds
}
