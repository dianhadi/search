package post

import (
	"context"

	"github.com/dianhadi/golib/tracer"
	"github.com/dianhadi/search/internal/entity"
	"github.com/dianhadi/search/pkg/utils"
)

func (r Repo) Put(ctx context.Context, post entity.Post) error {
	span, ctx := tracer.StartSpanRepo(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	return r.elastic.Put(ctx, postIndex, post)
}

func (r Repo) Search(ctx context.Context, query map[string]interface{}) ([]entity.Post, error) {
	span, ctx := tracer.StartSpanRepo(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	hits, err := r.elastic.Search(ctx, postIndex, query)
	if err != nil {
		return []entity.Post{}, err
	}

	// Parse hits into an array of Post
	var posts []entity.Post
	for _, hit := range hits {
		source, ok := hit.(map[string]interface{})["_source"].(map[string]interface{})
		if !ok {
			return []entity.Post{}, err
		}
		var post entity.Post
		post.ID = source["id"].(string)
		post.Title = source["title"].(string)
		post.SeoTitle = source["seo-title"].(string)
		post.Content = source["content"].(string)
		post.PreviewContent = source["preview-content"].(string)
		post.PublishDate = source["publish-date"].(string)
		post.AuthorID = source["author-id"].(string)

		posts = append(posts, post)
	}

	return posts, nil
}
