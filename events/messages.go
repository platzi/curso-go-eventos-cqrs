package events

import "platzi.com/go/cqrs/models"

type Message interface {
	Type() string
}

type CreatedFeedMessage struct {
	Feed *models.Feed
}

func (m CreatedFeedMessage) Type() string {
	return "created_feed"
}
