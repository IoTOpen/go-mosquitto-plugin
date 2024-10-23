package mosquitto

/*
#include <mosquitto.h>
#include <mosquitto_broker.h>
#include <mosquitto_plugin.h>
*/
import "C"
import "unsafe"

// EvtDisconnect Disconnect event
type (
	EvtDisconnect ptrStruct[C.struct_mosquitto_evt_disconnect]
)

func (e EvtDisconnect) String() string {
	return MosqEvtDisconnect.String()
}

func (e EvtDisconnect) asStruct() *C.struct_mosquitto_evt_disconnect {
	return ptrStruct[C.struct_mosquitto_evt_disconnect](e).getStruct()
}

func (e EvtDisconnect) Client() Client {
	x := e.asStruct()
	return Client{unsafe.Pointer(x.client)}
}
