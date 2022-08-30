package apperr

import "fmt"

type PointGraphError struct {
	error
}

func NewPointGraphErr(err error) error {
	return &PointGraphError{error: fmt.Errorf("point graph error: %w", err)}
}
func (p *PointGraphError) Error() string {
	return p.error.Error()
}
