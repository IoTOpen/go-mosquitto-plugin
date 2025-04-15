package mosquitto

/*
#include <mosquitto.h>
*/
import "C"
import (
	"unsafe"
)

type Event int

const (
	MosqEvtReload          = Event(C.MOSQ_EVT_RELOAD)
	MosqEvtACLCheck        = Event(C.MOSQ_EVT_ACL_CHECK)
	MosqEvtBasicAuth       = Event(C.MOSQ_EVT_BASIC_AUTH)
	MosqEvtEXTAuthStart    = Event(C.MOSQ_EVT_EXT_AUTH_START)
	MosqEvtEXTAuthContinue = Event(C.MOSQ_EVT_EXT_AUTH_CONTINUE)
	MosqEvtControl         = Event(C.MOSQ_EVT_CONTROL)
	MosqEvtMessageIn       = Event(C.MOSQ_EVT_MESSAGE_IN)
	MosqEvtPSKKey          = Event(C.MOSQ_EVT_PSK_KEY)
	MosqEvtTick            = Event(C.MOSQ_EVT_TICK)
	MosqEvtDisconnect      = Event(C.MOSQ_EVT_DISCONNECT)
	MosqEvtConnect         = Event(C.MOSQ_EVT_CONNECT)
	MosqEvtSubscribe       = Event(C.MOSQ_EVT_SUBSCRIBE)
	MosqEvtUnsubscribe     = Event(C.MOSQ_EVT_UNSUBSCRIBE)
	MosqEvtMessageOut      = Event(C.MOSQ_EVT_MESSAGE_OUT)
	MosqEvtClientOffline   = Event(C.MOSQ_EVT_CLIENT_OFFLINE)
)

func (e Event) String() string {
	return eventMap[e]
}

var eventMap = map[Event]string{
	MosqEvtControl:         "Control",
	MosqEvtACLCheck:        "ACLCheck",
	MosqEvtEXTAuthContinue: "AuthContinue",
	MosqEvtEXTAuthStart:    "AuthStart",
	MosqEvtPSKKey:          "PSKKey",
	MosqEvtReload:          "Reload",
	MosqEvtTick:            "Tick",
	MosqEvtMessageIn:       "MessageIn",
	MosqEvtDisconnect:      "Disconnect",
	MosqEvtBasicAuth:       "BasicAuth",
	MosqEvtConnect:         "Connect",
	MosqEvtSubscribe:       "Subscribe",
	MosqEvtUnsubscribe:     "Unsubscribe",
	MosqEvtMessageOut:      "MessageOut",
	MosqEvtClientOffline:   "ClientOffline",
}

type (
	ptrStruct[T any] struct {
		ptr unsafe.Pointer
	}
)

func (p ptrStruct[T]) getStruct() *T {
	return (*T)(p.ptr)
}
