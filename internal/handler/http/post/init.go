package post

import (
	"context"

	"github.com/dianhadi/search/internal/entity"
)

type usecasePost interface {
	Search(ctx context.Context, keyword string) ([]entity.Post, error)
}

type Handler struct {
	usecasePost usecasePost
}

func New(post usecasePost) (*Handler, error) {
	return &Handler{
		usecasePost: post,
	}, nil
}
