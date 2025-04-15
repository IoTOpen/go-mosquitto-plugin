package mosquitto

/*
#include <mosquitto.h>
*/
import "C"
import "unsafe"

// EvtUnsubscribe Unsubscribe event.
type (
	EvtUnsubscribe ptrStruct[C.struct_mosquitto_evt_unsubscribe]
)

func (e EvtUnsubscribe) String() string {
	return MosqEvtConnect.String()
}

func (e EvtUnsubscribe) asStruct() *C.struct_mosquitto_evt_unsubscribe {
	return ptrStruct[C.struct_mosquitto_evt_unsubscribe](e).getStruct()
}

func (e EvtUnsubscribe) Client() Client {
	x := e.asStruct()
	return Client{unsafe.Pointer(x.client)}
}

func (e EvtUnsubscribe) ClientID() string {
	x := e.asStruct()
	return C.GoString(x.data.clientid)
}

func (e EvtUnsubscribe) TopicFilter() string {
	x := e.asStruct()
	return C.GoString(x.data.topic_filter)
}

func (e EvtUnsubscribe) Identifier() uint32 {
	x := e.asStruct()
	return uint32(x.data.identifier)
}

func (e EvtUnsubscribe) Options() uint8 {
	x := e.asStruct()
	return uint8(x.data.options)
}
