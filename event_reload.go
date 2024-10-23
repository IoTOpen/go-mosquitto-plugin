package mosquitto

/*
#include <mosquitto.h>
#include <mosquitto_broker.h>
#include <mosquitto_plugin.h>
*/
import "C"

// EvtReload Reload event
type (
	EvtReload ptrStruct[C.struct_mosquitto_evt_reload]
)

func (e EvtReload) String() string {
	return MosqEvtReload.String()
}

func (e EvtReload) asStruct() *C.struct_mosquitto_evt_reload {
	return ptrStruct[C.struct_mosquitto_evt_reload](e).getStruct()
}

func (e EvtReload) Options() Options {
	x := e.asStruct()
	return optMap(x.options, x.option_count)
}
