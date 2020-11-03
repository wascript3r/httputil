package httputil

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func WrapCtx(ctx context.Context, handle HandleCtx) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		handle(ctx, w, r, p)
	}
}
