package mosquitto

/*
#include <mosquitto.h>
#include <mosquitto_broker.h>
#include <mosquitto_plugin.h>
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
	MosqEvtMessage         = Event(C.MOSQ_EVT_MESSAGE)
	MosqEvtPSKKey          = Event(C.MOSQ_EVT_PSK_KEY)
	MosqEvtTick            = Event(C.MOSQ_EVT_TICK)
	MosqEvtDisconnect      = Event(C.MOSQ_EVT_DISCONNECT)
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
	MosqEvtMessage:         "Message",
	MosqEvtDisconnect:      "Disconnect",
	MosqEvtBasicAuth:       "BasicAuth",
}

type (
	ptrStruct[T any] struct {
		ptr unsafe.Pointer
	}
	EvtPskKey       ptrStruct[C.struct_mosquitto_evt_psk_key]
	EvtExtendedAuth ptrStruct[C.struct_mosquitto_evt_extended_auth]
)

func (p ptrStruct[T]) getStruct() *T {
	return (*T)(p.ptr)
}
