package main

import (
	"errors"

	liberr "github.com/jortel/go-utils/error"
)

var (
	Wrap = liberr.Wrap
)

type ManifestNotFound struct {
}

func (m *ManifestNotFound) Error() (s string) {
	s = "No manifest associated with the application or found in the source repository."
	return
}

func (e *ManifestNotFound) Is(err error) (matched bool) {
	var inst *ManifestNotFound
	matched = errors.As(err, &inst)
	return
}

type RepositoryNotDefined struct {
	Role string
}

func (m *RepositoryNotDefined) Error() (s string) {
	s = m.Role + " repository not defined."
	return
}

func (e *RepositoryNotDefined) Is(err error) (matched bool) {
	var inst *RepositoryNotDefined
	matched = errors.As(err, &inst)
	return
}

type PlatformNotDefined struct {
}

func (m *PlatformNotDefined) Error() (s string) {
	s = "Application not associated with any platform."
	return
}

func (e *PlatformNotDefined) Is(err error) (matched bool) {
	var inst *PlatformNotDefined
	matched = errors.As(err, &inst)
	return
}
