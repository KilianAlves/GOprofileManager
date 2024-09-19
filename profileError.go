package profileManager

import (
	"errors"
)

var (
	// ErrProfileNotFound is returned when the profile is not found
	ErrProfileNotFound = errors.New("profile not found")
	// ErrProfileAlreadyExists is returned when the profile already exists
	ErrProfileAlreadyExists = errors.New("profile already exists")
	// ErrProfileNotSet is returned when the profile is not set
	ErrProfileNotSet = errors.New("profile not set")
	// ErrEmptyStruct is returned when the provided struct has no fields
	ErrEmptyStruct = errors.New("provided struct has no fields")
)
