package post

import (
	"context"

	"github.com/dianhadi/search/internal/entity"
)

type repoPost interface {
	Put(ctx context.Context, post entity.Post) error
	Search(ctx context.Context, query map[string]interface{}) ([]entity.Post, error)
}

type Usecase struct {
	repoPost repoPost
}

func New(post repoPost) (*Usecase, error) {
	return &Usecase{
		repoPost: post,
	}, nil
}
