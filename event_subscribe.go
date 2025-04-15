package mosquitto

/*
#include <mosquitto.h>
*/
import "C"
import "unsafe"

// EvtSubscribe Subscription event.
type (
	EvtSubscribe ptrStruct[C.struct_mosquitto_evt_subscribe]
)

func (e EvtSubscribe) String() string {
	return MosqEvtConnect.String()
}

func (e EvtSubscribe) asStruct() *C.struct_mosquitto_evt_subscribe {
	return ptrStruct[C.struct_mosquitto_evt_subscribe](e).getStruct()
}

func (e EvtSubscribe) Client() Client {
	x := e.asStruct()
	return Client{unsafe.Pointer(x.client)}
}

func (e EvtSubscribe) ClientID() string {
	x := e.asStruct()
	return C.GoString(x.data.clientid)
}

func (e EvtSubscribe) TopicFilter() string {
	x := e.asStruct()
	return C.GoString(x.data.topic_filter)
}

func (e EvtSubscribe) Identifier() uint32 {
	x := e.asStruct()
	return uint32(x.data.identifier)
}

func (e EvtSubscribe) Options() uint8 {
	x := e.asStruct()
	return uint8(x.data.options)
}
