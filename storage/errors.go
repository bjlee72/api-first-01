package storage

import "api-first-01/errors"

// Defines the list of well-known errors that can be raised from the storage.
var (
	ErrRecordNotFound = errors.NewErrorKind("storage", "record not found")
	ErrConflict       = errors.NewErrorKind("storage", "conflict - the record might be there already")
	ErrTransaction    = errors.NewErrorKind("storage", "transaction failure")
	ErrBadRequest     = errors.NewErrorKind("storage", "bad request")
)
