package post

import (
	"context"

	"github.com/dianhadi/golib/tracer"
	"github.com/dianhadi/search/internal/entity"
	"github.com/dianhadi/search/pkg/utils"
)

func (u Usecase) Put(ctx context.Context, post entity.Post) error {
	span, ctx := tracer.StartSpanUsecase(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	return u.repoPost.Put(ctx, post)
}
