package json

import (
	"encoding/json"
	"net/http"

	"github.com/wascript3r/gostr"
)

type Error interface {
	error
	Name() string
	Original() error
	Data() any
}

type jsonError struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func newJsonError(err Error) *jsonError {
	return &jsonError{
		Name:    err.Name(),
		Message: gostr.UpperFirst(err.Error()),
		Data:    err.Data(),
	}
}

type jsonRes struct {
	Error *jsonError  `json:"error"`
	Data  interface{} `json:"data"`
}

func SetContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func NotFound(w http.ResponseWriter, data interface{}) {
	encodeJson(w, http.StatusNotFound, NotFoundError, data)
}

func NotFoundCustom(w http.ResponseWriter, err Error, data interface{}) {
	encodeJson(w, http.StatusNotFound, err, data)
}

func Forbidden(w http.ResponseWriter, data interface{}) {
	encodeJson(w, http.StatusForbidden, ForbiddenError, data)
}

func ForbiddenCustom(w http.ResponseWriter, err Error, data interface{}) {
	encodeJson(w, http.StatusForbidden, err, data)
}

func Unauthorized(w http.ResponseWriter, data interface{}) {
	encodeJson(w, http.StatusUnauthorized, UnauthorizedError, data)
}

func UnauthorizedCustom(w http.ResponseWriter, err Error, data interface{}) {
	encodeJson(w, http.StatusUnauthorized, err, data)
}

func BadRequest(w http.ResponseWriter, data interface{}) {
	encodeJson(w, http.StatusBadRequest, BadRequestError, data)
}

func BadRequestCustom(w http.ResponseWriter, err Error, data interface{}) {
	encodeJson(w, http.StatusBadRequest, err, data)
}

func InternalError(w http.ResponseWriter, data interface{}) {
	encodeJson(w, http.StatusInternalServerError, InternalServerError, data)
}

func InternalErrorCustom(w http.ResponseWriter, err Error, data interface{}) {
	encodeJson(w, http.StatusInternalServerError, err, data)
}

func ServeErr(w http.ResponseWriter, err Error, data interface{}) {
	encodeJson(w, http.StatusUnprocessableEntity, err, data)
}

func ServeJSON(w http.ResponseWriter, data interface{}) {
	encodeJson(w, http.StatusOK, nil, data)
}

func Status(w http.ResponseWriter, status int) {
	SetContentType(w)
	w.WriteHeader(status)
}

func encodeJson(w http.ResponseWriter, status int, err Error, data interface{}) {
	SetContentType(w)
	w.WriteHeader(status)

	var jsonErr *jsonError
	if err != nil {
		jsonErr = newJsonError(err)
	}

	json.NewEncoder(w).Encode(
		jsonRes{
			Error: jsonErr,
			Data:  data,
		},
	)
}
