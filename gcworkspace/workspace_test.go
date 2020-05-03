package gcworkspace

import (
	"testing"

	gctest "github.com/kruc/clockify-api/test"
)

func workspaceClient(t *testing.T) *WorkspaceClient {
	tu := &gctest.TestUtil{}
	client := tu.MockClient(t)
	return NewClient(client)
}

func TestWorkspaceList(t *testing.T) {
	workspaceClient := workspaceClient(t)
	workspaces, err := workspaceClient.List()
	if err != nil {
		t.Fatal(err)
	}
	if len(workspaces) != 2 {
		t.Fatalf("Expected 2 Workspaces got %d", len(workspaces))
	}
	if workspaces[0].ID != "1" {
		t.Errorf("Expected Workspace Id 1 got %s", workspaces[0].ID)
	}
	if workspaces[0].Name != "DevOps" {
		t.Errorf("Expected Workspace name DevOps got %s", workspaces[0].Name)
	}
	if workspaces[1].ID != "2" {
		t.Errorf("Expected Workspace Id 2 got %s", workspaces[0].ID)
	}
	if workspaces[1].Name != "Kruc" {
		t.Errorf("Expected Workspace name Kruc got %s", workspaces[0].Name)
	}
}
