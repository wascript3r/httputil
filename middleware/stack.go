package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wascript3r/httputil"
)

type Middleware func(http.ResponseWriter, *http.Request, httprouter.Params) (next bool)

type Stack struct {
	middlewares []Middleware
}

func New() *Stack {
	return &Stack{nil}
}

func (s *Stack) Use(m Middleware) {
	s.middlewares = append(s.middlewares, m)
}

func (s *Stack) Wrap(fn httputil.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		for _, m := range s.middlewares {
			if !m(w, r, p) {
				return
			}
			fn(w, r, p)
		}
	}
}
