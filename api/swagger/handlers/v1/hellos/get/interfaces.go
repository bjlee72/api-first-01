package get

import (
	"context"
	"net/http"

	"api-first-01/storage"
)

// ContextCreator creates a morse request context.
type ContextCreator interface {
	Create(request *http.Request) context.Context
}

// HelloStorage represents a storage that supports hello creation operation.
//go:generate mockery -name HelloStorage -case underscore -inpkg -testonly
type HelloStorage interface {
	ReadHello(context.Context, *storage.ReadHelloRequest) (*storage.ReadHelloResponse, error)
}
