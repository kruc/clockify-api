package gctag

import (
	"errors"

	"github.com/kruc/clockify-api/gchttp"
)

// TagClient type
type TagClient struct {
	cc       *gchttp.ClockifyHTTPClient
	endpoint string
}

// Tags type
type Tags struct {
	Tags []Tag `json:"tags"`
}

// Tag type
type Tag struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	WorkspaceID string `json:"workspaceId"`
	Archived    bool   `json:"archived"`
}

// Add add tag to Tags collection
func (t *Tags) Add(tag Tag) *Tags {

	tagAlreadyExists, _ := t.Get(tag.Name)

	if tagAlreadyExists == nil {
		t.Tags = append(t.Tags, tag)
	}

	return t
}

// Remove tag from Tags collection
func (t *Tags) Remove(tagName string) *Tags {

	n := 0
	for _, tag := range t.Tags {
		if t.keep(tag, tagName) {
			t.Tags[n] = tag
			n++
		}
	}
	t.Tags = t.Tags[:n]

	return t
}

func (t *Tags) keep(tag Tag, remove string) bool {
	return tag.Name != remove
}

// Get find tag by name
func (t *Tags) Get(name string) (*Tag, error) {
	for _, tag := range t.Tags {
		if tag.Name == name {
			return &tag, nil
		}
	}

	err := errors.New("Tag not found")

	return nil, err
}

// ToMap converts []Tag to map[string]Tag
func (t *Tags) ToMap() map[string]Tag {
	tagMap := make(map[string]Tag, len(t.Tags))

	for _, tag := range t.Tags {
		tagMap[tag.Name] = tag
	}

	return tagMap
}
