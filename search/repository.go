package search

import (
	"context"

	"platzi.com/go/cqrs/models"
)

type SearchRepository interface {
	Close()
	IndexFeed(ctx context.Context, feed models.Feed) error
	SearchFeed(ctx context.Context, query string) ([]models.Feed, error)
}

var repo SearchRepository

func SetSearchRepository(r SearchRepository) {
	repo = r
}

func Close() {
	repo.Close()
}

func IndexFeed(ctx context.Context, feed models.Feed) error {
	return repo.IndexFeed(ctx, feed)
}

func SearchFeed(ctx context.Context, query string) ([]models.Feed, error) {
	return repo.SearchFeed(ctx, query)
}
