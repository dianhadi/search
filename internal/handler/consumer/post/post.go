package post

import (
	"context"
	"encoding/json"

	"github.com/dianhadi/search/internal/entity"
	"github.com/dianhadi/search/pkg/log"
	"github.com/dianhadi/search/pkg/mq"
	"github.com/dianhadi/search/pkg/tracer"
	"github.com/dianhadi/search/pkg/utils"
)

func PostCreatedConsumer() mq.HandlerConsumer {
	return func(ctx context.Context, body []byte) error {
		span, ctx := tracer.StartSpanHandler(ctx, utils.GetCurrentFunctionName())
		defer span.End()

		var post entity.Post
		err := json.Unmarshal(body, &post)
		log.Info(post, err)
		if err != nil {
			return err
		}

		return nil
	}
}

// func PostEditedConsumer() (consumer.HandlerConsumer, error) {

// }
