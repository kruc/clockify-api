package gcworkspace

import "github.com/kruc/clockify-api/gchttp"

// WorkspaceClient type
type WorkspaceClient struct {
	cc       *gchttp.ClockifyHTTPClient
	endpoint string
}

// Workspace type
type Workspace struct {
	ID                      string            `json:"id"`
	Name                    string            `json:"name"`
	HourlyRate              HourlyRate        `json:"hourlyRate"`
	Memberships             Memberships       `json:"memberships"`
	WorkspaceSettings       WorkspaceSettings `json:"workspaceSettings"`
	ImageURL                string            `json:"imageUrl"`
	FeatureSubscriptionType interface{}       `json:"featureSubscriptionType"`
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

// HourlyRate type
type HourlyRate struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}

// WorkspaceSettings type
type WorkspaceSettings struct {
	TimeRoundingInReports              bool          `json:"timeRoundingInReports"`
	OnlyAdminsSeeBillableRates         bool          `json:"onlyAdminsSeeBillableRates"`
	OnlyAdminsCreateProject            bool          `json:"onlyAdminsCreateProject"`
	OnlyAdminsSeeDashboard             bool          `json:"onlyAdminsSeeDashboard"`
	DefaultBillableProjects            bool          `json:"defaultBillableProjects"`
	LockTimeEntries                    interface{}   `json:"lockTimeEntries"`
	Round                              Round         `json:"round"`
	ProjectFavorites                   bool          `json:"projectFavorites"`
	CanSeeTimeSheet                    bool          `json:"canSeeTimeSheet"`
	CanSeeTracker                      bool          `json:"canSeeTracker"`
	ProjectPickerSpecialFilter         bool          `json:"projectPickerSpecialFilter"`
	ForceProjects                      bool          `json:"forceProjects"`
	ForceTasks                         bool          `json:"forceTasks"`
	ForceTags                          bool          `json:"forceTags"`
	ForceDescription                   bool          `json:"forceDescription"`
	OnlyAdminsSeeAllTimeEntries        bool          `json:"onlyAdminsSeeAllTimeEntries"`
	OnlyAdminsSeePublicProjectsEntries bool          `json:"onlyAdminsSeePublicProjectsEntries"`
	TrackTimeDownToSecond              bool          `json:"trackTimeDownToSecond"`
	ProjectGroupingLabel               string        `json:"projectGroupingLabel"`
	AdminOnlyPages                     []interface{} `json:"adminOnlyPages"`
	AutomaticLock                      interface{}   `json:"automaticLock"`
	OnlyAdminsCreateTag                bool          `json:"onlyAdminsCreateTag"`
	OnlyAdminsCreateTask               bool          `json:"onlyAdminsCreateTask"`
	IsProjectPublicByDefault           bool          `json:"isProjectPublicByDefault"`
}

// Round type
type Round struct {
	Round   string `json:"round"`
	Minutes string `json:"minutes"`
}

// Workspaces list of Workspace
type Workspaces []Workspace
