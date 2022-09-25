package errcode

type ErrName string

type Error struct {
	name     ErrName
	original error
}

func New(name ErrName, original error) *Error {
	return &Error{name, original}
}

func (e *Error) Error() string {
	return e.original.Error()
}

func (e *Error) Name() string {
	return string(e.name)
}

func (e *Error) Original() error {
	return e.original
}

func (e *Error) Data() any {
	return nil
}
