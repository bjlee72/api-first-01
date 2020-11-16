package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewErrorKind(t *testing.T) {
	// component part should not be the part of the error message.
	assert.Equal(t, "bad request", NewErrorKind("auth", "bad request").Error())
}

func TestErrorKind_Matches(t *testing.T) {
	errorKind := NewErrorKind("auth", "bad request")
	anotherErrorKind := NewErrorKind("handler", "bad request")

	assert.True(t, errorKind.Matches(errorKind))
	assert.True(t, errorKind.Matches(errorKind.New("property foo is missing")))
	assert.True(t, errorKind.Matches(errorKind.Newf("property %s is missing", "bar")))
	assert.True(t, errorKind.Matches(errorKind.Wrap(fmt.Errorf("error"))))
	assert.False(t, errorKind.Matches(anotherErrorKind))
	assert.False(t, errorKind.Matches(anotherErrorKind.New("property foo is missing")))
	assert.False(t, errorKind.Matches(anotherErrorKind.Newf("property %s is missing", "bar")))
	assert.False(t, errorKind.Matches(anotherErrorKind.Wrap(fmt.Errorf("error"))))
}
