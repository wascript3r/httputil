package httputil

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Handle func(http.ResponseWriter, *http.Request, httprouter.Params)
type HandleCtx func(context.Context, http.ResponseWriter, *http.Request, httprouter.Params)
