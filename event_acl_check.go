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

// EvtAclCheck ACL Check event
type (
	EvtAclCheck ptrStruct[C.struct_mosquitto_evt_acl_check]
)

func (e EvtAclCheck) String() string {
	return MosqEvtACLCheck.String()
}

func (e EvtAclCheck) asStruct() *C.struct_mosquitto_evt_acl_check {
	return ptrStruct[C.struct_mosquitto_evt_acl_check](e).getStruct()
}

func (e EvtAclCheck) Client() Client {
	x := e.asStruct()
	return Client{unsafe.Pointer(x.client)}
}

func (e EvtAclCheck) Topic() string {
	x := e.asStruct()
	return C.GoString(x.topic)
}

func (e EvtAclCheck) Payload() []byte {
	x := e.asStruct()
	return unsafe.Slice((*byte)(x.payload), x.payloadlen)
}

func (e EvtAclCheck) QoS() int {
	x := e.asStruct()
	return int(x.qos) & 0xFF
}

func (e EvtAclCheck) Retained() bool {
	x := e.asStruct()
	return bool(x.retain)
}

func (e EvtAclCheck) Access() Access {
	x := e.asStruct()
	return Access(x.access)
}
