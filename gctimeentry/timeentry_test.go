package gctimeentry

import (
	"testing"

	"github.com/kruc/clockify-api/gctag"
)

func TestIsTagged(t *testing.T) {
	timeEntry := TimeEntry{
		ID:          "5ea9d03cdcc3755f82fab5ea",
		Description: "Krucafuks",
		Tags: []gctag.Tag{
			{
				ID:          "id1",
				Name:        "logged",
				WorkspaceID: "workspaceId",
				Archived:    false,
			},
			{
				ID:          "id2",
				Name:        "failed",
				WorkspaceID: "workspaceId",
				Archived:    false,
			},
		},
	}

	if timeEntry.IsTagged("logged") != true {
		t.Fatalf("Expected true got %v", timeEntry.IsTagged("logged"))
	}
	if timeEntry.IsTagged("random-tag") != false {
		t.Fatalf("Expected false got %v", timeEntry.IsTagged("random-tag"))
	}
}

func TestAddTag(t *testing.T) {
	timeEntry := TimeEntry{
		ID:          "5ea9d03cdcc3755f82fab5ea",
		Description: "Krucafuks",
		TagIds:      []string{"id1"},
	}

	timeEntry.AddTag("id1")

	if len(timeEntry.TagIds) != 1 {
		t.Fatalf("Expected 1 got %v", len(timeEntry.TagIds))
	}

	if timeEntry.TagIds[0] != "id1" {
		t.Fatalf("Expected id1 got %v", timeEntry.TagIds[0])
	}

	timeEntry.AddTag("id2")

	if len(timeEntry.TagIds) != 2 {
		t.Fatalf("Expected 2 got %v", len(timeEntry.TagIds))
	}
	if timeEntry.TagIds[0] != "id1" {
		t.Fatalf("Expected id1 got %v", timeEntry.TagIds[0])
	}
	if timeEntry.TagIds[1] != "id2" {
		t.Fatalf("Expected id2 got %v", timeEntry.TagIds[1])
	}
}

func TestRemoveTag(t *testing.T) {
	timeEntry := TimeEntry{
		ID:          "5ea9d03cdcc3755f82fab5ea",
		Description: "Krucafuks",
		TagIds:      []string{"id1", "id2"},
	}

	timeEntry.RemoveTag("id1")

	if len(timeEntry.TagIds) != 1 {
		t.Fatalf("Expected 1 got %v", len(timeEntry.TagIds))
	}

	if timeEntry.TagIds[0] != "id2" {
		t.Fatalf("Expected id2 got %v", timeEntry.TagIds[0])
	}

	timeEntry.RemoveTag("id2")

	if len(timeEntry.TagIds) != 0 {
		t.Fatalf("Expected 0 got %v", len(timeEntry.TagIds))
	}
}
