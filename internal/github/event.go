package github

import (
	"fmt"
	"strings"
	"time"
)

const GHBaseURL string = "https://github.com/"

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
	Payload   Payload    `json:"payload"`
	Public    bool       `json:"public"`
	CreatedAt time.Time  `json:"created_at"`
}

type DisplayEvent struct {
	ID        string
	Type      string
	RepoName  string
	RepoURL   string
	EventURL  string
	Timestamp string
}

func (e Event) Display() DisplayEvent {
	eventType := strings.Split(string(e.Type), "Event")[0]

	eventURL := ""
	switch e.Type {
	case PushEvent:
		eventURL = GHBaseURL + e.Repo.Name + "/commit/" + e.Payload.Head
	case CreateEvent:
		if e.Payload.RefType == "branch" {
			eventURL = GHBaseURL + e.Repo.Name + "/tree/" + e.Payload.Ref
		} else {
			eventURL = GHBaseURL + e.Repo.Name + "/releases/tag/" + e.Payload.Ref
		}
	default:
		eventURL = GHBaseURL + e.Repo.Name
	}

	return DisplayEvent{
		ID:        e.ID,
		Type:      eventType,
		RepoName:  e.Repo.Name,
		RepoURL:   GHBaseURL + e.Repo.Name,
		EventURL:  eventURL,
		Timestamp: GetRelativeTime(e.CreatedAt),
	}
}

func GetRelativeTime(unformattedTime time.Time) string {
	now := time.Now()
	diff := now.Sub(unformattedTime)
	hours := int(diff.Hours())

	if hours < 1 {
		return "just now"
	}

	if hours == 1 {
		return "one hour ago"
	}

	if hours < 24 {
		return fmt.Sprintf("%d hours ago", hours)
	}

	days := hours / 24

	if days == 1 {
		return "1 day ago"
	}

	if days < 7 {
		return fmt.Sprintf("%d days ago", days)
	}

	if unformattedTime.Year() == now.Year() {
		return unformattedTime.Local().Format("Jan 02")
	}
	return unformattedTime.Local().Format("Jan 02, 2006")
}
