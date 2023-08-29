package post

import (
	"context"

	"github.com/dianhadi/golib/tracer"
	"github.com/dianhadi/search/internal/entity"
	"github.com/dianhadi/search/pkg/utils"
)

func (u Usecase) Search(ctx context.Context, keyword string) ([]entity.Post, error) {
	span, ctx := tracer.StartSpanUsecase(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	// Build the Elasticsearch query
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"content": keyword,
			},
		},
	}

	return u.repoPost.Search(ctx, query)
}
