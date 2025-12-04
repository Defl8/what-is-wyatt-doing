package github

import (
	"strings"
	"time"
)

type EventType string

const (
	PushEvent   EventType = "PushEvent"
	CreateEvent EventType = "CreateEvent"
)

type Event struct {
	ID        string     `json:"id"`
	Type      EventType  `json:"type"`
	Actor     Actor      `json:"actor"`
	Repo      Repository `json:"repo"`
	Public    bool       `json:"public"`
	CreatedAt time.Time  `json:"created_at"`
}

type DisplayEvent struct {
	ID        string
	Type      string
	RepoName  string
	RepoURL   string
	Timestamp string
}

func (e Event) Display() DisplayEvent {
	eventType := strings.Split(string(e.Type), "Event")[0]
	return DisplayEvent{
		ID:        e.ID,
		Type:      eventType,
		RepoName:  e.Repo.Name,
		RepoURL:   "https://github.com/" + e.Repo.Name,
		Timestamp: e.CreatedAt.Local().Format("2006-01-02 15:04:05"),
	}
}
