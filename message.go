package mosquitto

/*
#include <mosquitto.h>
 */
import "C"
import "unsafe"

type Message ptrStruct[C.struct_mosquitto_message]

func (m Message) asStruct() *C.struct_mosquitto_message {
	return (*C.struct_mosquitto_message)(m.ptr)
}

func (m Message) Topic() string {
	x := m.asStruct()
	return C.GoString(x.topic)
}

func (m Message) Payload() []byte {
	x := m.asStruct()
	return unsafe.Slice((*byte)(x.payload), x.payloadlen)
}

func (m Message) QoS() int {
	x := m.asStruct()
	return int(x.qos) & 0xFF
}

func (m Message) Retained() bool {
	x := m.asStruct()
	return bool(x.retain)
}
