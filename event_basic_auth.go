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

// EvtBasicAuth Basic auth event
type (
	EvtBasicAuth ptrStruct[C.struct_mosquitto_evt_basic_auth]
)

func (e EvtBasicAuth) String() string {
	return MosqEvtBasicAuth.String()
}

func (e EvtBasicAuth) asStruct() *C.struct_mosquitto_evt_basic_auth {
	return ptrStruct[C.struct_mosquitto_evt_basic_auth](e).getStruct()
}

func (e EvtBasicAuth) Username() string {
	x := e.asStruct()
	return C.GoString(x.username)
}

func (e EvtBasicAuth) Password() string {
	x := e.asStruct()
	return C.GoString(x.password)
}

func (e EvtBasicAuth) Client() Client {
	x := e.asStruct()
	return Client{unsafe.Pointer(x.client)}
}
