package middleware

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wascript3r/httputil"
)

type MiddlewareCtx func(context.Context, http.ResponseWriter, *http.Request, httprouter.Params) (next bool)

type StackCtx struct {
	middlewares []MiddlewareCtx
}

func NewCtx() *StackCtx {
	return &StackCtx{nil}
}

func (s *StackCtx) Use(m MiddlewareCtx) {
	s.middlewares = append(s.middlewares, m)
}

func (s *StackCtx) Wrap(ctx context.Context, fn httputil.HandleCtx) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		for _, m := range s.middlewares {
			if !m(ctx, w, r, p) {
				return
			}
			fn(ctx, w, r, p)
		}
	}
}
