package mosquitto

/*
#include <mosquitto.h>
#include <mosquitto_broker.h>
#include <mosquitto_plugin.h>
*/
import "C"

import "time"

// EvtTick Tick event
type (
	EvtTick ptrStruct[C.struct_mosquitto_evt_tick]
)

func (e EvtTick) String() string {
	return MosqEvtTick.String()
}

func (e EvtTick) asStruct() *C.struct_mosquitto_evt_tick {
	return ptrStruct[C.struct_mosquitto_evt_tick](e).getStruct()
}

func (e EvtTick) Now() time.Time {
	x := e.asStruct()
	return time.Unix(int64(x.now_s), int64(x.now_ns))
}

func (e EvtTick) Next() time.Time {
	x := e.asStruct()
	return time.Unix(int64(x.next_s), int64(x.next_ns))
}
