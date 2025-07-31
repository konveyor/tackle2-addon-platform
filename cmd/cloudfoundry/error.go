package cloudfoundry

import (
	"errors"

	liberr "github.com/jortel/go-utils/error"
)

var (
	Wrap = liberr.Wrap
)

type CoordinatesError struct {
}

func (m *CoordinatesError) Error() (s string) {
	s = "Application coordinates not defined."
	return
}

func (e *CoordinatesError) Is(err error) (matched bool) {
	var inst *CoordinatesError
	matched = errors.As(err, &inst)
	return
}
