package mosquitto

/*
#include <mosquitto.h>
*/
import "C"
import "unsafe"

// EvtClientOffline Client offline event.
type (
	EvtClientOffline ptrStruct[C.struct_mosquitto_evt_client_offline]
)

func (e EvtClientOffline) String() string {
	return MosqEvtConnect.String()
}

func (e EvtClientOffline) asStruct() *C.struct_mosquitto_evt_client_offline {
	return ptrStruct[C.struct_mosquitto_evt_client_offline](e).getStruct()
}

func (e EvtClientOffline) Client() Client {
	x := e.asStruct()
	return Client{unsafe.Pointer(x.client)}
}

func (e EvtClientOffline) Reason() ReasonCode {
	x := e.asStruct()
	return ReasonCode(x.reason)
}
