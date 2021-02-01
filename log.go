package mosquitto

/*
#include <malloc.h>
#include <mosquitto.h>
#include <mosquitto_broker.h>

void go_mosquitto_log_printf(int level, const char* fmt, const char* string);
*/
import "C"
import (
	"log"
	"strings"
	"unsafe"
)

type logger struct{}

func (logger) Write(p []byte) (n int, err error) {
	f := C.CString("%s")
	s := C.CString(strings.TrimRight(string(p), " \t\n"))
	var level = C.MOSQ_LOG_INFO
	if len(p) > 0 {
		switch p[0] {
		case 'I':
			level = C.MOSQ_LOG_INFO
		case 'N':
			level = C.MOSQ_LOG_NOTICE
		case 'W':
			level = C.MOSQ_LOG_WARNING
		case 'E':
			level = C.MOSQ_LOG_ERR
		case 'D':
			level = C.MOSQ_LOG_DEBUG
		}
	}
	C.go_mosquitto_log_printf(C.int(level), f, s)
	C.free(unsafe.Pointer(f))
	C.free(unsafe.Pointer(s))
	return len(p), nil
}

var mosquittoLogger = &logger{}

func init() {
	log.SetFlags(0)
	log.SetOutput(mosquittoLogger)
}
