package main

import (
	"errors"

	liberr "github.com/jortel/go-utils/error"
)

var (
	Wrap = liberr.Wrap
)

type ManifestError struct {
}

func (m *ManifestError) Error() (s string) {
	s = "No manifest associated with the application or found in the source repository."
	return
}

func (e *ManifestError) Is(err error) (matched bool) {
	var inst *ManifestError
	matched = errors.As(err, &inst)
	return
}

type SourceRepoError struct {
}

func (m *SourceRepoError) Error() (s string) {
	s = "Application source repository not defined."
	return
}

func (e *SourceRepoError) Is(err error) (matched bool) {
	var inst *AssetRepoError
	matched = errors.As(err, &inst)
	return
}

type AssetRepoError struct {
}

func (m *AssetRepoError) Error() (s string) {
	s = "Application asset repository not defined."
	return
}

func (e *AssetRepoError) Is(err error) (matched bool) {
	var inst *AssetRepoError
	matched = errors.As(err, &inst)
	return
}

type PlatformError struct {
}

func (m *PlatformError) Error() (s string) {
	s = "Application not associated with platform."
	return
}

func (e *PlatformError) Is(err error) (matched bool) {
	var inst *PlatformError
	matched = errors.As(err, &inst)
	return
}
