package github

import "time"

type EventType string

const (
	PushEvent EventType = "PushEvent"
	CreateEvent EventType = "CreateEvent"
)

type Event struct {
	ID   string `json:"id"`
	Type EventType `json:"type"`
	Public bool `json:"public"`
	CreatedAt time.Time `json:"created_at"`
}
