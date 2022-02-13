package events

import "time"

type Message interface {
	Type() string
}

type CreatedFeedMessage struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func (m CreatedFeedMessage) Type() string {
	return "created_feed"
}
