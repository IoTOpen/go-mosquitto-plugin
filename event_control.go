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

// EvtControl Control event
type (
	EvtControl ptrStruct[C.struct_mosquitto_evt_control]
)

func (e EvtControl) String() string {
	return MosqEvtControl.String()
}

func (e EvtControl) asStruct() *C.struct_mosquitto_evt_control {
	return ptrStruct[C.struct_mosquitto_evt_control](e).getStruct()
}

func (e EvtControl) Client() Client {
	x := e.asStruct()
	return Client{unsafe.Pointer(x.client)}
}

func (e EvtControl) Topic() string {
	x := e.asStruct()
	return C.GoString(x.topic)
}

func (e EvtControl) Payload() []byte {
	x := e.asStruct()
	return unsafe.Slice((*byte)(x.payload), x.payloadlen)
}

func (e EvtControl) QoS() int {
	x := e.asStruct()
	return int(x.qos) & 0xFF
}

func (e EvtControl) Retained() bool {
	x := e.asStruct()
	return bool(x.retain)
}
