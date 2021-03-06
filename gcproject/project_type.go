package gcproject

import (
	"github.com/kruc/clockify-api/gchttp"
)

// ProjectClient type
type ProjectClient struct {
	cc       *gchttp.ClockifyHTTPClient
	endpoint string
}

//Project type
type Project struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	HourlyRate  HourlyRate   `json:"hourlyRate"`
	ClientID    string       `json:"clientId"`
	WorkspaceID string       `json:"workspaceId"`
	Billable    bool         `json:"billable"`
	Memberships []Membership `json:"memberships"`
	Color       string       `json:"color"`
	Estimate    Estimate     `json:"estimate"`
	Archived    bool         `json:"archived"`
	Duration    string       `json:"duration"`
	ClientName  string       `json:"clientName"`
	Note        string       `json:"note"`
	Public      bool         `json:"public"`
	Template    bool         `json:"template"`
}

// Estimate type
type Estimate struct {
	Estimate string `json:"estimate"`
	Type     string `json:"type"`
}

// HourlyRate type
type HourlyRate struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

// Membership type
type Membership struct {
	UserID           string      `json:"userId"`
	HourlyRate       interface{} `json:"hourlyRate"`
	TargetID         string      `json:"targetId"`
	MembershipType   string      `json:"membershipType"`
	MembershipStatus string      `json:"membershipStatus"`
}

// Projects list of projects
type Projects []Project
