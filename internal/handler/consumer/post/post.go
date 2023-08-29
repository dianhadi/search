package post

import (
	"context"
	"encoding/json"

	"github.com/dianhadi/golib/tracer"
	"github.com/dianhadi/search/internal/entity"
	"github.com/dianhadi/search/pkg/mq"
	"github.com/dianhadi/search/pkg/utils"
)

func (h Handler) PostCreatedConsumer() mq.HandlerConsumer {
	return func(ctx context.Context, body []byte) error {
		span, ctx := tracer.StartSpanHandler(ctx, utils.GetCurrentFunctionName())
		defer span.End()

		var post entity.Post
		err := json.Unmarshal(body, &post)
		if err != nil {
			return err
		}

		return h.usecasePost.Put(ctx, post)
	}
}

// func PostEditedConsumer() (consumer.HandlerConsumer, error) {

// }
