package mosquitto

/*
#include <mosquitto.h>
*/
import "C"
import "unsafe"

// EvtConnect Client connect event.
type (
	EvtConnect ptrStruct[C.struct_mosquitto_evt_connect]
)

func (e EvtConnect) String() string {
	return MosqEvtConnect.String()
}

func (e EvtConnect) asStruct() *C.struct_mosquitto_evt_connect {
	return ptrStruct[C.struct_mosquitto_evt_connect](e).getStruct()
}

func (e EvtConnect) Client() Client {
	x := e.asStruct()
	return Client{unsafe.Pointer(x.client)}
}
