package github

import "time"

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
	Timestamp time.Time
}

func (e Event) Display() DisplayEvent {
	return DisplayEvent{
		ID:        e.ID,
		Type:      string(e.Type),
		RepoName:  e.Repo.Name,
		RepoURL:   e.Repo.URL,
		Timestamp: e.CreatedAt,
	}
}
