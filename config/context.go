package config

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/rs/xid"
)

// NewContextFactory creates a new context factory from the given environment.
func NewContextFactory(environment *Environment) *ContextFactory {
	return &ContextFactory{
		projectID: environment.ProjectID(),
		stage:     environment.Stage(),
	}
}

// ContextFactory handles the swagger-based request context object creation.
type ContextFactory struct {
	projectID string
	stage     string
}

// Create creates a context.
func (f *ContextFactory) Create(request *http.Request) context.Context {
	// If no trace is given, we use this default value.
	trace := fmt.Sprintf("projects/%s/traces/%s", f.projectID, xid.New().String())

	traceHeader := request.Header.Get("X-Cloud-Trace-Context")
	traceParts := strings.Split(traceHeader, "/")
	if len(traceParts) > 0 && len(traceParts[0]) > 0 {
		trace = fmt.Sprintf("projects/%s/traces/%s", f.projectID, traceParts[0])
	}

	ctx := context.Background()
	ctx = WithStage(ctx, f.stage)
	ctx = WithTrace(ctx, trace)

	return ctx
}

type contextKey struct {
	Key string
}

var (
	contextKeyTrace = contextKey{Key: "trace"}
	contextKeyStage = contextKey{Key: "stage"}
)

// WithTrace adds a trace information to the given context.
func WithTrace(ctx context.Context, trace string) context.Context {
	return context.WithValue(ctx, contextKeyTrace, trace)
}

// Trace extracts trace information from the given context.
func Trace(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	v := ctx.Value(contextKeyTrace)
	if str, ok := v.(string); ok {
		return str
	}
	return ""
}

// MustTrace extracts trace information from the given context and panics if it's empty.
func MustTrace(ctx context.Context) string {
	if trace := Trace(ctx); trace != "" {
		return trace
	}
	panic("trace is empty in context")
}

// WithStage sets stage information to the given context.
func WithStage(ctx context.Context, stage string) context.Context {
	return context.WithValue(ctx, contextKeyStage, stage)
}

// Stage extracts stage information from the given context.
func Stage(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	v := ctx.Value(contextKeyStage)
	if str, ok := v.(string); ok {
		return str
	}
	return ""
}

// MustStage extracts stage information from the given context and panics if it's empty.
func MustStage(ctx context.Context) string {
	if stage := Stage(ctx); stage != "" {
		return stage
	}
	panic("stage is empty in context")
}
