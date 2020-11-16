package get

import (
	"context"
	"net/http"
)

// ContextCreator creates a morse request context.
type ContextCreator interface {
	Create(request *http.Request) context.Context
}
