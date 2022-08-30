package apperr

import "fmt"

type SaveGraphError struct {
	error
}

func NewSaveGraphErr(err error) error {
	return &SaveGraphError{error: fmt.Errorf("save graph error: %w", err)}
}

func (s *SaveGraphError) Error() string {
	return s.error.Error()
}
