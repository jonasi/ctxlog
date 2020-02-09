package ctxlog

import (
	"fmt"
	"os"
	"sync"
)

type basicLogger struct {
	mu sync.Mutex
	kv map[string]interface{}
}

func (b *basicLogger) log(level string, args ...interface{}) {
	b.logf(level, "%v", args...)
}

func (b *basicLogger) logf(level, format string, args ...interface{}) {
	b.mu.Lock()
	defer b.mu.Unlock()

	args = append([]interface{}{level}, args...)
	if len(b.kv) != 0 {
		format += " %v"
		args = append(args, b.kv)
	}

	_, _ = fmt.Fprintf(os.Stderr, "[%s] "+format+"\n", args...)
}

func (b *basicLogger) Debug(args ...interface{}) {
	b.log("debug", args...)
}
func (b *basicLogger) Info(args ...interface{}) {
	b.log("info", args...)
}
func (b *basicLogger) Warn(args ...interface{}) {
	b.log("warn", args...)
}
func (b *basicLogger) Error(args ...interface{}) {
	b.log("error", args...)
}
func (b *basicLogger) Fatal(args ...interface{}) {
	b.log("fatal", args...)
	os.Exit(1)
}
func (b *basicLogger) Debugf(format string, args ...interface{}) {
	b.logf("debug", format, args...)
}
func (b *basicLogger) Infof(format string, args ...interface{}) {
	b.logf("info", format, args...)
}
func (b *basicLogger) Warnf(format string, args ...interface{}) {
	b.logf("warn", format, args...)
}
func (b *basicLogger) Errorf(format string, args ...interface{}) {
	b.logf("error", format, args...)
}
func (b *basicLogger) Fatalf(format string, args ...interface{}) {
	b.logf("fatal", format, args...)
	os.Exit(1)
}

func (b *basicLogger) KV(k string, v interface{}) Logger {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.kv == nil {
		b.kv = map[string]interface{}{}
	}

	b.kv[k] = v
	return b
}
