package post

import (
	"context"

	"github.com/dianhadi/search/internal/entity"
)

type usecasePost interface {
	Put(ctx context.Context, post entity.Post) error
}

type Handler struct {
	usecasePost usecasePost
}

func New(post usecasePost) (*Handler, error) {
	return &Handler{
		usecasePost: post,
	}, nil
}
