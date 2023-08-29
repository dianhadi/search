package post

import (
	"net/http"

	"github.com/dianhadi/golib/tracer"
	"github.com/dianhadi/search/internal/handler/helper"
	"github.com/dianhadi/search/pkg/errors"
	"github.com/dianhadi/search/pkg/utils"
)

func (h Handler) Search(w http.ResponseWriter, r *http.Request) {
	span, ctx := tracer.StartSpanHandler(r.Context(), utils.GetCurrentFunctionName())
	defer span.End()

	q := r.URL.Query().Get("q")
	if q == "" {
		err := errors.NewWithMessage(errors.StatusRequestBodyInvalid, "Query is Empty")
		helper.Write(w, ctx, err, nil)
		return
	}

	posts, err := h.usecasePost.Search(ctx, q)
	if err != nil {
		helper.Write(w, ctx, err, nil)
		return
	}

	helper.Write(w, ctx, nil, posts)
	return
}
