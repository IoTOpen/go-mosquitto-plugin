package mosquitto

/*
#include <mosquitto.h>
#include <mosquitto_broker.h>
*/
import "C"

type Protocol int
type ProtocolVersion int

func (p Protocol) String() string {
	return protocolMap[p]
}

func (p ProtocolVersion) String() string {
	return protocolVersionMap[p]
}

const (
	ProtocolMQTT           = Protocol(C.mp_mqtt)
	ProtocolMQTTSN         = Protocol(C.mp_mqttsn)
	ProtocolMQTTWebsockets = Protocol(C.mp_websockets)
)

const (
	MQTTProtocolV31  = ProtocolVersion(C.MQTT_PROTOCOL_V31)
	MQTTProtocolV311 = ProtocolVersion(C.MQTT_PROTOCOL_V311)
	MQTTProtocolV5   = ProtocolVersion(C.MQTT_PROTOCOL_V5)
)

var protocolMap = map[Protocol]string{
	ProtocolMQTT:           "MQTT",
	ProtocolMQTTSN:         "MQTT-SN",
	ProtocolMQTTWebsockets: "MQTT-Websocket",
}

var protocolVersionMap = map[ProtocolVersion]string{
	MQTTProtocolV5:   "v5",
	MQTTProtocolV31:  "v3.1",
	MQTTProtocolV311: "v3.1.1",
}
