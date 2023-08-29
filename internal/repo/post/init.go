package post

import (
	"context"
)

const (
	postIndex string = "blog-posts"
)

type elastic interface {
	Put(ctx context.Context, index string, value interface{}) error
	Search(ctx context.Context, index string, query map[string]interface{}) ([]interface{}, error)
}

type Repo struct {
	elastic elastic
}

func New(elastic elastic) (Repo, error) {
	return Repo{
		elastic: elastic,
	}, nil
}
