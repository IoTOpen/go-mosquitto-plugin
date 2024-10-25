package mosquitto

/*
#include <mosquitto.h>
#include <mosquitto_broker.h>
#include <mqtt_protocol.h>
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

var reasonCodeMap = map[ReasonCode]string{
	ReasonSuccess: "Success",
	//	ReasonNormalDisconnection:         "Normal Disconnection",
	//	ReasonGrantedQoS0:                 "Granted QoS 0",
	ReasonGrantedQoS1:                 "Granted QoS 1",
	ReasonGrantedQoS2:                 "Granted QoS 2",
	ReasonDisconnectWithWillMsg:       "Disconnect with Will Message",
	ReasonNoMatchingSubscribers:       "No Matching Subscribers",
	ReasonNoSubscriptionExisted:       "No Subscription Existed",
	ReasonContinueAuthentication:      "Continue Authentication",
	ReasonReauthenticate:              "Reauthenticate",
	ReasonUnspecified:                 "Unspecified",
	ReasonMalformedPacket:             "Malformed Packet",
	ReasonProtocolError:               "Protocol Error",
	ReasonImplementationSpecific:      "Implementation Specific",
	ReasonUnsupportedProtocolVersion:  "Unsupported Protocol Version",
	ReasonClientIDNotValid:            "Client ID Not Valid",
	ReasonBadUsernameOrPassword:       "Bad Username or Password",
	ReasonNotAuthorized:               "Not Authorized",
	ReasonServerUnavailable:           "Server Unavailable",
	ReasonServerBusy:                  "Server Busy",
	ReasonBanned:                      "Banned",
	ReasonServerShuttingDown:          "Server Shutting Down",
	ReasonBadAuthenticationMethod:     "Bad Authentication Method",
	ReasonKeepAliveTimeout:            "Keep Alive Timeout",
	ReasonSessionTakenOver:            "Session Taken Over",
	ReasonTopicFilterInvalid:          "Topic Filter Invalid",
	ReasonTopicNameInvalid:            "Topic Name Invalid",
	ReasonPacketIDInUse:               "Packet ID In Use",
	ReasonPacketIDNotFound:            "Packet ID Not Found",
	ReasonReceiveMaximumExceeded:      "Receive Maximum Exceeded",
	ReasonTopicAliasInvalid:           "Topic Alias Invalid",
	ReasonPacketTooLarge:              "Packet Too Large",
	ReasonMessageRateTooHigh:          "Message Rate Too High",
	ReasonQuotaExceeded:               "Quota Exceeded",
	ReasonAdministrativeAction:        "Administrative Action",
	ReasonPayloadFormatInvalid:        "Payload Format Invalid",
	ReasonRetainNotSupported:          "Retain Not Supported",
	ReasonQoSNotSupported:             "QoS Not Supported",
	ReasonUseAnotherServer:            "Use Another Server",
	ReasonServerMoved:                 "Server Moved",
	ReasonSharedSubsNotSupported:      "Shared Subscriptions Not Supported",
	ReasonConnectionRateExceeded:      "Connection Rate Exceeded",
	ReasonMaximumConnectTime:          "Maximum Connect Time",
	ReasonSubscriptionIDsNotSupported: "Subscription IDs Not Supported",
	ReasonWildcardSubsNotSupported:    "Wildcard Subscriptions Not Supported",
}

type ConnACKCode311 uint8

const (
	ConnAckAccepted                   = ConnACKCode311(C.CONNACK_ACCEPTED)
	ConnAckRefusedProtocolVersion     = ConnACKCode311(C.CONNACK_REFUSED_PROTOCOL_VERSION)
	ConnAckRefusedIdentifierRejected  = ConnACKCode311(C.CONNACK_REFUSED_IDENTIFIER_REJECTED)
	ConnAckRefusedServerUnavailable   = ConnACKCode311(C.CONNACK_REFUSED_SERVER_UNAVAILABLE)
	ConnAckRefusedBadUsernamePassword = ConnACKCode311(C.CONNACK_REFUSED_BAD_USERNAME_PASSWORD)
	ConnAckRefusedNotAuthorized       = ConnACKCode311(C.CONNACK_REFUSED_NOT_AUTHORIZED)
)

type ReasonCode uint8

func (rc ReasonCode) String() string {
	return reasonCodeMap[rc]
}

const (
	ReasonSuccess                     = ReasonCode(C.MQTT_RC_SUCCESS)                        // 0
	ReasonNormalDisconnection         = ReasonCode(C.MQTT_RC_NORMAL_DISCONNECTION)           // 0
	ReasonGrantedQoS0                 = ReasonCode(C.MQTT_RC_GRANTED_QOS0)                   // 0
	ReasonGrantedQoS1                 = ReasonCode(C.MQTT_RC_GRANTED_QOS1)                   // 1
	ReasonGrantedQoS2                 = ReasonCode(C.MQTT_RC_GRANTED_QOS2)                   // 2
	ReasonDisconnectWithWillMsg       = ReasonCode(C.MQTT_RC_DISCONNECT_WITH_WILL_MSG)       // 4
	ReasonNoMatchingSubscribers       = ReasonCode(C.MQTT_RC_NO_MATCHING_SUBSCRIBERS)        // 16
	ReasonNoSubscriptionExisted       = ReasonCode(C.MQTT_RC_NO_SUBSCRIPTION_EXISTED)        // 17
	ReasonContinueAuthentication      = ReasonCode(C.MQTT_RC_CONTINUE_AUTHENTICATION)        // 24
	ReasonReauthenticate              = ReasonCode(C.MQTT_RC_REAUTHENTICATE)                 // 25
	ReasonUnspecified                 = ReasonCode(C.MQTT_RC_UNSPECIFIED)                    // 128
	ReasonMalformedPacket             = ReasonCode(C.MQTT_RC_MALFORMED_PACKET)               // 129
	ReasonProtocolError               = ReasonCode(C.MQTT_RC_PROTOCOL_ERROR)                 // 130
	ReasonImplementationSpecific      = ReasonCode(C.MQTT_RC_IMPLEMENTATION_SPECIFIC)        // 131
	ReasonUnsupportedProtocolVersion  = ReasonCode(C.MQTT_RC_UNSUPPORTED_PROTOCOL_VERSION)   // 132
	ReasonClientIDNotValid            = ReasonCode(C.MQTT_RC_CLIENTID_NOT_VALID)             // 133
	ReasonBadUsernameOrPassword       = ReasonCode(C.MQTT_RC_BAD_USERNAME_OR_PASSWORD)       // 134
	ReasonNotAuthorized               = ReasonCode(C.MQTT_RC_NOT_AUTHORIZED)                 // 135
	ReasonServerUnavailable           = ReasonCode(C.MQTT_RC_SERVER_UNAVAILABLE)             // 136
	ReasonServerBusy                  = ReasonCode(C.MQTT_RC_SERVER_BUSY)                    // 137
	ReasonBanned                      = ReasonCode(C.MQTT_RC_BANNED)                         // 138
	ReasonServerShuttingDown          = ReasonCode(C.MQTT_RC_SERVER_SHUTTING_DOWN)           // 139
	ReasonBadAuthenticationMethod     = ReasonCode(C.MQTT_RC_BAD_AUTHENTICATION_METHOD)      // 140
	ReasonKeepAliveTimeout            = ReasonCode(C.MQTT_RC_KEEP_ALIVE_TIMEOUT)             // 141
	ReasonSessionTakenOver            = ReasonCode(C.MQTT_RC_SESSION_TAKEN_OVER)             // 142
	ReasonTopicFilterInvalid          = ReasonCode(C.MQTT_RC_TOPIC_FILTER_INVALID)           // 143
	ReasonTopicNameInvalid            = ReasonCode(C.MQTT_RC_TOPIC_NAME_INVALID)             // 144
	ReasonPacketIDInUse               = ReasonCode(C.MQTT_RC_PACKET_ID_IN_USE)               // 145
	ReasonPacketIDNotFound            = ReasonCode(C.MQTT_RC_PACKET_ID_NOT_FOUND)            // 146
	ReasonReceiveMaximumExceeded      = ReasonCode(C.MQTT_RC_RECEIVE_MAXIMUM_EXCEEDED)       // 147
	ReasonTopicAliasInvalid           = ReasonCode(C.MQTT_RC_TOPIC_ALIAS_INVALID)            // 148
	ReasonPacketTooLarge              = ReasonCode(C.MQTT_RC_PACKET_TOO_LARGE)               // 149
	ReasonMessageRateTooHigh          = ReasonCode(C.MQTT_RC_MESSAGE_RATE_TOO_HIGH)          // 150
	ReasonQuotaExceeded               = ReasonCode(C.MQTT_RC_QUOTA_EXCEEDED)                 // 151
	ReasonAdministrativeAction        = ReasonCode(C.MQTT_RC_ADMINISTRATIVE_ACTION)          // 152
	ReasonPayloadFormatInvalid        = ReasonCode(C.MQTT_RC_PAYLOAD_FORMAT_INVALID)         // 153
	ReasonRetainNotSupported          = ReasonCode(C.MQTT_RC_RETAIN_NOT_SUPPORTED)           // 154
	ReasonQoSNotSupported             = ReasonCode(C.MQTT_RC_QOS_NOT_SUPPORTED)              // 155
	ReasonUseAnotherServer            = ReasonCode(C.MQTT_RC_USE_ANOTHER_SERVER)             // 156
	ReasonServerMoved                 = ReasonCode(C.MQTT_RC_SERVER_MOVED)                   // 157
	ReasonSharedSubsNotSupported      = ReasonCode(C.MQTT_RC_SHARED_SUBS_NOT_SUPPORTED)      // 158
	ReasonConnectionRateExceeded      = ReasonCode(C.MQTT_RC_CONNECTION_RATE_EXCEEDED)       // 159
	ReasonMaximumConnectTime          = ReasonCode(C.MQTT_RC_MAXIMUM_CONNECT_TIME)           // 160
	ReasonSubscriptionIDsNotSupported = ReasonCode(C.MQTT_RC_SUBSCRIPTION_IDS_NOT_SUPPORTED) // 161
	ReasonWildcardSubsNotSupported    = ReasonCode(C.MQTT_RC_WILDCARD_SUBS_NOT_SUPPORTED)    // 162
)
