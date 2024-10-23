package mosquitto

/*
#include <mosquitto.h>
#include <mosquitto_broker.h>
#include <mosquitto_plugin.h>
*/
import "C"
import "unsafe"

// EvtMessage Message event
type (
	EvtMessage ptrStruct[C.struct_mosquitto_evt_message]
)

func (e EvtMessage) String() string {
	return MosqEvtMessage.String()
}

func (e EvtMessage) asStruct() *C.struct_mosquitto_evt_message {
	return ptrStruct[C.struct_mosquitto_evt_message](e).getStruct()
}

func (e EvtMessage) Client() Client {
	x := e.asStruct()
	return Client{unsafe.Pointer(x.client)}
}

func (e EvtMessage) Payload() []byte {
	x := e.asStruct()
	return unsafe.Slice((*byte)(x.payload), x.payloadlen)
}

func (e EvtMessage) SetPayload(data []byte) {
	x := e.asStruct()
	x.payloadlen = 0
	dataPtr := C.CBytes(data)
	x.payload = dataPtr
	x.payloadlen = C.uint(len(data))
}

func (e EvtMessage) Retained() bool {
	x := e.asStruct()
	return bool(x.retain)
}

func (e EvtMessage) SetRetained(retain bool) {
	x := e.asStruct()
	x.retain = C._Bool(retain)
}

func (e EvtMessage) QoS() int {
	x := e.asStruct()
	return int(x.qos) & 0xFF
}

func (e EvtMessage) SetQoS(qos int) {
	x := e.asStruct()
	x.qos = C.uint8_t(qos & 0xFF)
}

func (e EvtMessage) Topic() string {
	x := e.asStruct()
	return C.GoString(x.topic)
}

func (e EvtMessage) SetTopic(topic string) {
	x := e.asStruct()
	x.topic = C.CString(topic)
}
