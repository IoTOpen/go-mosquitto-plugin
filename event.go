package mosquitto

/*
#include <mosquitto.h>
#include <mosquitto_broker.h>
#include <mosquitto_plugin.h>
*/
import "C"
import (
	"reflect"
	"time"
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
	MosqEvtConnect         = Event(C.MOSQ_EVT_CONNECT)
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
	MosqEvtConnect:         "Connect",
	MosqEvtBasicAuth:       "BasicAuth",
}

type (
	EvtReload       uintptr
	EvtAclCheck     uintptr
	EvtBasicAuth    uintptr
	EvtPskKey       uintptr
	EvtExtendedAuth uintptr
	EvtControl      uintptr
	EvtMessage      uintptr
	EvtTick         uintptr
	EvtDisconnect   uintptr
	EvtConnect      uintptr
)

// Reload event
func (e EvtReload) Options() Options {
	x := (*C.struct_mosquitto_evt_reload)(unsafe.Pointer(e))
	return optMap(x.options, x.option_count)
}

// Tick event
func (e EvtTick) Now() time.Time {
	x := (*C.struct_mosquitto_evt_tick)(unsafe.Pointer(e))
	return time.Unix(int64(x.now_s), int64(x.now_ns))
}

func (e EvtTick) Next() time.Time {
	x := (*C.struct_mosquitto_evt_tick)(unsafe.Pointer(e))
	return time.Unix(int64(x.next_s), int64(x.next_ns))
}

// ACL Check event
func (e EvtAclCheck) Client() Client {
	x := (*C.struct_mosquitto_evt_acl_check)(unsafe.Pointer(e))
	return Client(unsafe.Pointer(x.client))
}

func (e EvtAclCheck) Topic() string {
	x := (*C.struct_mosquitto_evt_acl_check)(unsafe.Pointer(e))
	return C.GoString(x.topic)
}

func (e EvtAclCheck) Payload() []byte {
	x := (*C.struct_mosquitto_evt_acl_check)(unsafe.Pointer(e))
	var res []byte
	setSliceHeader((*reflect.SliceHeader)(unsafe.Pointer(&res)), uintptr(unsafe.Pointer(x.payload)), C.int(x.payloadlen))
	return res
}

func (e EvtAclCheck) QoS() int {
	x := (*C.struct_mosquitto_evt_acl_check)(unsafe.Pointer(e))
	return int(x.qos) & 0xFF
}

func (e EvtAclCheck) Retained() bool {
	x := (*C.struct_mosquitto_evt_acl_check)(unsafe.Pointer(e))
	return x.retain == true
}

func (e EvtAclCheck) Access() Access {
	x := (*C.struct_mosquitto_evt_acl_check)(unsafe.Pointer(e))
	return Access(x.access)
}

// Basic auth event
func (e EvtBasicAuth) Username() string {
	x := (*C.struct_mosquitto_evt_basic_auth)(unsafe.Pointer(e))
	return C.GoString(x.username)
}

func (e EvtBasicAuth) Password() string {
	x := (*C.struct_mosquitto_evt_basic_auth)(unsafe.Pointer(e))
	return C.GoString(x.password)
}

func (e EvtBasicAuth) Client() Client {
	x := (*C.struct_mosquitto_evt_basic_auth)(unsafe.Pointer(e))
	return Client(unsafe.Pointer(x.client))
}

// Disconnect event
func (e EvtDisconnect) Client() Client {
	x := (*C.struct_mosquitto_evt_disconnect)(unsafe.Pointer(e))
	return Client(unsafe.Pointer(x.client))
}

func (e EvtDisconnect) Reason() int {
	x := (*C.struct_mosquitto_evt_disconnect)(unsafe.Pointer(e))
	return int(x.reason)
}

// Connect event
func (e EvtConnect) Client() Client {
	x := (*C.struct_mosquitto_evt_connect)(unsafe.Pointer(e))
	return Client(unsafe.Pointer(x.client))
}

// Message event
func (e EvtMessage) Client() Client {
	x := (*C.struct_mosquitto_evt_message)(unsafe.Pointer(e))
	return Client(unsafe.Pointer(x.client))
}

func (e EvtMessage) Payload() []byte {
	x := (*C.struct_mosquitto_evt_message)(unsafe.Pointer(e))
	var res []byte
	setSliceHeader((*reflect.SliceHeader)(unsafe.Pointer(&res)), uintptr(unsafe.Pointer(x.payload)), C.int(x.payloadlen))
	return res
}

func (e EvtMessage) SetPayload(data []byte) {
	x := (*C.struct_mosquitto_evt_message)(unsafe.Pointer(e))
	x.payloadlen = 0
	dataPtr := C.CBytes(data)
	x.payload = dataPtr
	x.payloadlen = C.uint(len(data))
}

func (e EvtMessage) Retained() bool {
	x := (*C.struct_mosquitto_evt_message)(unsafe.Pointer(e))
	return bool(x.retain)
}

func (e EvtMessage) SetRetained(retain bool) {
	x := (*C.struct_mosquitto_evt_message)(unsafe.Pointer(e))
	x.retain = retain == true
}

func (e EvtMessage) QoS() int {
	x := (*C.struct_mosquitto_evt_message)(unsafe.Pointer(e))
	return int(x.qos) & 0xFF
}

func (e EvtMessage) SetQoS(qos int) {
	x := (*C.struct_mosquitto_evt_message)(unsafe.Pointer(e))
	x.qos = C.uint8_t(qos & 0xFF)
}

func (e EvtMessage) Topic() string {
	x := (*C.struct_mosquitto_evt_message)(unsafe.Pointer(e))
	return C.GoString(x.topic)
}

func (e EvtMessage) SetTopic(topic string) {
	x := (*C.struct_mosquitto_evt_message)(unsafe.Pointer(e))
	x.topic = C.CString(topic)
}

// Control event
func (e EvtControl) Client() Client {
	x := (*C.struct_mosquitto_evt_control)(unsafe.Pointer(e))
	return Client(unsafe.Pointer(x.client))
}

func (e EvtControl) Topic() string {
	x := (*C.struct_mosquitto_evt_control)(unsafe.Pointer(e))
	return C.GoString(x.topic)
}

func (e EvtControl) Payload() []byte {
	x := (*C.struct_mosquitto_evt_control)(unsafe.Pointer(e))
	var res []byte
	setSliceHeader((*reflect.SliceHeader)(unsafe.Pointer(&res)), uintptr(unsafe.Pointer(x.payload)), C.int(x.payloadlen))
	return res
}

func (e EvtControl) QoS() int {
	x := (*C.struct_mosquitto_evt_control)(unsafe.Pointer(e))
	return int(x.qos) & 0xFF
}

func (e EvtControl) Retained() bool {
	x := (*C.struct_mosquitto_evt_control)(unsafe.Pointer(e))
	return bool(x.retain)
}
