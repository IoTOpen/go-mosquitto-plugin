package mosquitto

/*
#include <malloc.h>
#include <mosquitto.h>
#include <mosquitto_broker.h>

void go_mosquitto_log_printf(int level, const char* fmt, const char* string);
*/
import "C"
import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"strings"
	"unsafe"
)

type logger struct{}

func (logger) Write(p []byte) (n int, err error) {
	f := C.CString("%s")
	logString := string(p)
	s := C.CString(strings.TrimRight(logString, " \t\n"))
	var level = C.MOSQ_LOG_INFO
	parts := strings.SplitN(logString, ":", 2)
	if len(parts) == 2 && len(parts[0]) < 8 {
		switch strings.ToLower(parts[0]) {
		case "i", "info":
			level = C.MOSQ_LOG_INFO
		case "n", "notice":
			level = C.MOSQ_LOG_NOTICE
		case "w", "warn", "warning":
			level = C.MOSQ_LOG_WARNING
		case "e", "err", "error":
			level = C.MOSQ_LOG_ERR
		case "d", "debug":
			level = C.MOSQ_LOG_DEBUG
		}
	}
	C.go_mosquitto_log_printf(C.int(level), f, s)
	C.free(unsafe.Pointer(f))
	C.free(unsafe.Pointer(s))
	return len(p), nil
}

var mosquittoLogger = &logger{}

type slogHandler struct {
	attrs []slog.Attr
	group string
}

func (s slogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return true
}

func (s slogHandler) Handle(ctx context.Context, record slog.Record) error {
	level := C.MOSQ_LOG_INFO
	switch record.Level {
	case slog.LevelDebug:
		level = C.MOSQ_LOG_DEBUG
	case slog.LevelInfo:
		level = C.MOSQ_LOG_INFO
	case slog.LevelError:
		level = C.MOSQ_LOG_ERR
	case slog.LevelWarn:
		level = C.MOSQ_LOG_WARNING
	}

	msg := record.Message
	if s.group != "" {
		msg = s.group + ": " + msg
	}

	var collectedAttrs []slog.Attr
	record.Attrs(func(attr slog.Attr) bool {
		collectedAttrs = append(collectedAttrs, attr)
		return true
	})

	collectedAttrs = append(collectedAttrs, s.attrs...)

	for _, attr := range collectedAttrs {
		msg += fmt.Sprintf(" %s=%v", attr.Key, attr.Value)
	}

	f := C.CString("%s")
	msgString := C.CString(msg)
	C.go_mosquitto_log_printf(C.int(level), f, msgString)
	C.free(unsafe.Pointer(f))
	C.free(unsafe.Pointer(msgString))
	return nil
}

func (s slogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newAttrs := append(s.attrs, attrs...)
	return slogHandler{attrs: newAttrs, group: s.group}
}

func (s slogHandler) WithGroup(name string) slog.Handler {
	return slogHandler{attrs: s.attrs, group: name}
}

func init() {
	log.SetFlags(0)
	log.SetOutput(mosquittoLogger)
	slog.SetDefault(slog.New(slogHandler{}))
}
