package mosquitto

/*
#include <mosquitto.h>
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

func (e EvtDisconnect) Reason() ReasonCode {
	x := e.asStruct()
	return ReasonCode(x.reason)
}
