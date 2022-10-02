package json

import (
	"net/http"
)

type CodeMapper struct {
	errorMap map[string]int
}

func NewCodeMapper() *CodeMapper {
	return &CodeMapper{
		errorMap: make(map[string]int),
	}
}

func (c *CodeMapper) Register(code int, errs ...Error) {
	for _, err := range errs {
		name := err.Name()
		c.errorMap[name] = code
	}
}

func (c *CodeMapper) ServeErr(w http.ResponseWriter, err Error, data interface{}) {
	code, ok := c.errorMap[err.Name()]
	if !ok {
		code = http.StatusUnprocessableEntity
	}

	encodeJson(w, code, err, data)
}
