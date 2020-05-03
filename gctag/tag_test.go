package gctag

import (
	"testing"
)

func TestAddTag(t *testing.T) {
	tags := Tags{}

	if len(tags.Tags) != 0 {
		t.Fatalf("Expected 0 Tags got %d", len(tags.Tags))
	}

	tag1 := Tag{
		ID:          "tag1_id",
		Name:        "tag1_name",
		WorkspaceID: "workspace_id",
		Archived:    false,
	}
	tag2 := Tag{
		ID:          "tag2_id",
		Name:        "tag2_name",
		WorkspaceID: "workspace_id",
		Archived:    false,
	}

	tags.Add(tag1)
	tags.Add(tag2)
	tags.Add(tag1)

	if len(tags.Tags) != 2 {
		t.Fatalf("Expected 2 Tags got %d", len(tags.Tags))
	}

	if tags.Tags[0].ID != "tag1_id" {
		t.Fatalf("Expected tag_id got %s", tags.Tags[0].ID)
	}

	if tags.Tags[1].ID != "tag2_id" {
		t.Fatalf("Expected tag_id got %s", tags.Tags[1].ID)
	}
}

func TestFindTagByName(t *testing.T) {
	tag1 := Tag{
		ID:          "tag1_id",
		Name:        "tag1_name",
		WorkspaceID: "workspace_id",
		Archived:    false,
	}
	tag2 := Tag{
		ID:          "tag2_id",
		Name:        "tag2_name",
		WorkspaceID: "workspace_id",
		Archived:    false,
	}
	tags := Tags{
		[]Tag{tag1, tag2},
	}

	expectedTag, err := tags.Get("tag2_name")
	if *expectedTag != tag2 {
		t.Fatalf("Expected tag2_name tag got %v", expectedTag.Name)
	}

	if err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}

	emptyTag, err := tags.Get("non-exist")
	if emptyTag != nil {
		t.Fatalf("Non exist")
	}
}

func TestRemoveTag(t *testing.T) {
	tags := Tags{}

	tag1 := Tag{
		ID:          "tag1_id",
		Name:        "tag1_name",
		WorkspaceID: "workspace_id",
		Archived:    false,
	}
	tag2 := Tag{
		ID:          "tag2_id",
		Name:        "tag2_name",
		WorkspaceID: "workspace_id",
		Archived:    false,
	}

	tags.Add(tag1)
	tags.Add(tag2)

	tags.Remove("tag1_name")

	if len(tags.Tags) != 1 {
		t.Fatalf("Expected 1 Tags got %d", len(tags.Tags))
	}

	if tags.Tags[0].ID != "tag2_id" {
		t.Fatalf("Expected tag2_id got %s", tags.Tags[0].ID)
	}
}

func TestToMap(t *testing.T) {
	tags := Tags{}

	if len(tags.Tags) != 0 {
		t.Fatalf("Expected 0 Tags got %d", len(tags.Tags))
	}

	tag1 := Tag{
		ID:          "tag1_id",
		Name:        "tag1_name",
		WorkspaceID: "workspace_id",
		Archived:    false,
	}
	tag2 := Tag{
		ID:          "tag2_id",
		Name:        "tag2_name",
		WorkspaceID: "workspace_id",
		Archived:    false,
	}

	tags.Add(tag1)
	tags.Add(tag2)
	tags.Add(tag1)

	tagsMap := tags.ToMap()
	_, exists := tagsMap["tag2_name"]

	if exists != true {
		t.Fatalf("Expected map element tag2 got %v", tagsMap)
	}
}
