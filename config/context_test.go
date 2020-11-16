package config

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithTrace(t *testing.T) {
	ctx := WithTrace(context.Background(), "test trace")
	assert.Equal(t, Trace(ctx), "test trace")
}

func TestMustTrace(t *testing.T) {
	ctx := WithTrace(context.Background(), "test trace")
	assert.Equal(t, MustTrace(ctx), "test trace")

	ctx = WithTrace(context.Background(), "")
	assert.Panics(t, func() { MustTrace(ctx) })
}

func TestWithStage(t *testing.T) {
	ctx := WithStage(context.Background(), "test")
	assert.Equal(t, Stage(ctx), "test")
}

func TestMustStage(t *testing.T) {
	ctx := WithStage(context.Background(), "test")
	assert.Equal(t, MustStage(ctx), "test")

	ctx = WithStage(context.Background(), "")
	assert.Panics(t, func() { MustStage(ctx) })
}
