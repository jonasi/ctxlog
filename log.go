package ctxlog

import (
	"context"
)

type ctxKey struct{}

// L returns the logger found on ctx
func L(ctx context.Context) Logger {
	if ctx == nil {
		return &basicLogger{}
	}

	l, ok := ctx.Value(ctxKey{}).(Logger)
	if !ok {
		return &basicLogger{}
	}

	return l
}

// WithLogger stores a new logger on the context
func WithLogger(ctx context.Context, l Logger) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}

	return context.WithValue(ctx, ctxKey{}, l)
}

// WithKV stores a new logger on the context with k=>v set on it's
// kv
func WithKV(ctx context.Context, k string, v interface{}) context.Context {
	return WithLogger(ctx, KV(ctx, k, v))
}

// Debug logs with level=debug
func Debug(ctx context.Context, args ...interface{}) {
	L(ctx).Debug(args...)
}

// Info logs with level=info
func Info(ctx context.Context, args ...interface{}) {
	L(ctx).Info(args...)
}

// Warn logs with level=warn
func Warn(ctx context.Context, args ...interface{}) {
	L(ctx).Warn(args...)
}

// Error logs with level=error
func Error(ctx context.Context, args ...interface{}) {
	L(ctx).Error(args...)
}

// Fatal logs with level=fatal
func Fatal(ctx context.Context, args ...interface{}) {
	L(ctx).Fatal(args...)
}

// Debugf logs with level=debug
func Debugf(ctx context.Context, str string, args ...interface{}) {
	L(ctx).Debugf(str, args...)
}

// Infof logs with level=info
func Infof(ctx context.Context, str string, args ...interface{}) {
	L(ctx).Infof(str, args...)
}

// Warnf logs with level=warn
func Warnf(ctx context.Context, str string, args ...interface{}) {
	L(ctx).Warnf(str, args...)
}

// Errorf logs with level=error
func Errorf(ctx context.Context, str string, args ...interface{}) {
	L(ctx).Errorf(str, args...)
}

// Fatalf logs with level=fatal
func Fatalf(ctx context.Context, str string, args ...interface{}) {
	L(ctx).Fatalf(str, args...)
}

// KV returns a logger with key=value added to context
func KV(ctx context.Context, key string, value interface{}) Logger {
	return L(ctx).KV(key, value)
}
