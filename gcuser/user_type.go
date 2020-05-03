package gcuser

import (
	"github.com/kruc/clockify-api/gchttp"
)

// UserClient type
type UserClient struct {
	cc       *gchttp.ClockifyHTTPClient
	endpoint string
}

// User type
type User struct {
	ID               string      `json:"id"`
	Email            string      `json:"email"`
	Name             string      `json:"name"`
	Memberships      Memberships `json:"memberships"`
	ProfilePicture   string      `json:"profilePicture"`
	ActiveWorkspace  string      `json:"activeWorkspace"`
	DefaultWorkspace string      `json:"defaultWorkspace"`
	Settings         Settings    `json:"settings"`
	Status           string      `json:"status"`
}

// Settings type
type Settings struct {
	WeekStart                   string                `json:"weekStart"`
	TimeZone                    string                `json:"timeZone"`
	TimeFormat                  string                `json:"timeFormat"`
	DateFormat                  string                `json:"dateFormat"`
	SendNewsletter              bool                  `json:"sendNewsletter"`
	WeeklyUpdates               bool                  `json:"weeklyUpdates"`
	LongRunning                 bool                  `json:"longRunning"`
	TimeTrackingManual          bool                  `json:"timeTrackingManual"`
	SummaryReportSettings       SummaryReportSettings `json:"summaryReportSettings"`
	IsCompactViewOn             bool                  `json:"isCompactViewOn"`
	DashboardSelection          string                `json:"dashboardSelection"`
	DashboardViewType           string                `json:"dashboardViewType"`
	DashboardPinToTop           bool                  `json:"dashboardPinToTop"`
	ProjectListCollapse         int                   `json:"projectListCollapse"`
	CollapseAllProjectLists     bool                  `json:"collapseAllProjectLists"`
	GroupSimilarEntriesDisabled bool                  `json:"groupSimilarEntriesDisabled"`
	MyStartOfDay                string                `json:"myStartOfDay"`
}

// SummaryReportSettings type
type SummaryReportSettings struct {
	Group    string `json:"group"`
	Subgroup string `json:"subgroup"`
}

// Membership type
type Membership struct {
	UserID           string      `json:"userId"`
	HourlyRate       interface{} `json:"hourlyRate"`
	TargetID         string      `json:"targetId"`
	MembershipType   string      `json:"membershipType"`
	MembershipStatus string      `json:"membershipStatus"`
}

// Memberships type
type Memberships []Membership

// Users list of User
type Users []User
