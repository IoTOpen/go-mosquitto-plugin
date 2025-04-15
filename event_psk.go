package mosquitto

/*
#include <mosquitto.h>
*/
import "C"
import "unsafe"

// EvtPskKey Pre Shared Key event.
type (
	EvtPskKey ptrStruct[C.struct_mosquitto_evt_psk_key]
)

func (e EvtPskKey) asStruct() *C.struct_mosquitto_evt_psk_key {
	return ptrStruct[C.struct_mosquitto_evt_psk_key](e).getStruct()
}

func (e EvtPskKey) String() string {
	return MosqEvtPSKKey.String()
}

func (e EvtPskKey) Client() Client {
	x := e.asStruct()
	return Client{unsafe.Pointer(x.client)}
}

func (e EvtPskKey) Hint() string {
	x := e.asStruct()
	return C.GoString(x.hint)
}

func (e EvtPskKey) Identity() string {
	x := e.asStruct()
	return C.GoString(x.identity)
}

func (e EvtPskKey) Key() []byte {
	x := e.asStruct()
	data := unsafe.Slice((*byte)(unsafe.Pointer(x.key)), x.max_key_len)
	return data
}
