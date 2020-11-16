package errors

import "fmt"

// ErrorKind is the type of error.
type ErrorKind struct {
	component string
	title     string
}

// NewErrorKind creates a new error kind.
func NewErrorKind(comp, title string) *ErrorKind {
	return &ErrorKind{
		component: comp,
		title:     title,
	}
}

// Error returns the string tha the kind represents.
func (kind ErrorKind) Error() string {
	return kind.title
}

// Matches decides whether the given error is the same kind.
func (kind ErrorKind) Matches(err error) bool {
	if e, ok := err.(*ErrorKind); ok {
		return e.component == kind.component && e.title == kind.title
	}
	if e, ok := err.(*Error); ok {
		return e.kind.component == kind.component && e.kind.title == kind.title
	}

	return false
}

// New returns a new Error instance.
func (kind ErrorKind) New(msg string) *Error {
	return &Error{
		kind:    kind,
		message: msg,
	}
}

// Newf returns a new Error instance.
func (kind ErrorKind) Newf(format string, args ...interface{}) *Error {
	return kind.New(fmt.Sprintf(format, args...))
}

// Wrap returns a new Error instance.
func (kind ErrorKind) Wrap(err error) *Error {
	return kind.New(err.Error())
}

// Error represents an error.
type Error struct {
	kind    ErrorKind
	message string
}

// Error returns the reason why the error happened in string.
func (e *Error) Error() string {
	return e.message
}
