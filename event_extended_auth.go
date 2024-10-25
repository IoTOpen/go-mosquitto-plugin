package mosquitto

/*
#include <mosquitto.h>
#include <mosquitto_broker.h>
#include <mosquitto_plugin.h>
*/
import "C"
import "unsafe"

type (
	EvtExtendedAuth ptrStruct[C.struct_mosquitto_evt_extended_auth]
)

func (e EvtExtendedAuth) String() string {
	return "ExtendedAuth"
}

func (e EvtExtendedAuth) asStruct() *C.struct_mosquitto_evt_extended_auth {
	return ptrStruct[C.struct_mosquitto_evt_extended_auth](e).getStruct()
}

func (e EvtExtendedAuth) Client() Client {
	x := e.asStruct()
	return Client{unsafe.Pointer(x.client)}
}

func (e EvtExtendedAuth) AuthMethod() string {
	x := e.asStruct()
	return C.GoString(x.auth_method)
}

func (e EvtExtendedAuth) DataIn() []byte {
	x := e.asStruct()
	dataSize := C.int(x.data_in_len)
	data := C.GoBytes(x.data_in, dataSize) // Read only
	return data
}

func (e EvtExtendedAuth) DataOut() []byte {
	x := e.asStruct()
	dataSize := C.int(x.data_out_len)
	data := unsafe.Slice((*byte)(x.data_out), dataSize) // Place to write
	return data
}
