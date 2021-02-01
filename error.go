package mosquitto

/*
#include <mosquitto.h>
*/
import "C"

type Error int

func (e Error) String() string {
	return errorMap[e]
}

func (e Error) Error() string {
	return errorMap[e]
}

// Errors from mosquitto.h
const (
	MosqErrAuthContinue         = Error(C.MOSQ_ERR_AUTH_CONTINUE)
	MosqErrNoSubscribers        = Error(C.MOSQ_ERR_NO_SUBSCRIBERS)
	MosqErrSubExists            = Error(C.MOSQ_ERR_SUB_EXISTS)
	MosqErrConnPending          = Error(C.MOSQ_ERR_CONN_PENDING)
	MosqErrSuccess              = Error(C.MOSQ_ERR_SUCCESS)
	MosqErrNoMem                = Error(C.MOSQ_ERR_NOMEM)
	MosqErrProtocol             = Error(C.MOSQ_ERR_PROTOCOL)
	MosqErrInval                = Error(C.MOSQ_ERR_INVAL)
	MosqErrNoConn               = Error(C.MOSQ_ERR_NO_CONN)
	MosqErrConnRefused          = Error(C.MOSQ_ERR_CONN_REFUSED)
	MosqErrNotFound             = Error(C.MOSQ_ERR_NOT_FOUND)
	MosqErrConnLost             = Error(C.MOSQ_ERR_CONN_LOST)
	MosqErrTLS                  = Error(C.MOSQ_ERR_TLS)
	MosqErrPayloadSize          = Error(C.MOSQ_ERR_PAYLOAD_SIZE)
	MosqErrNotSupported         = Error(C.MOSQ_ERR_NOT_SUPPORTED)
	MosqErrAuth                 = Error(C.MOSQ_ERR_AUTH)
	MosqErrACLDenied            = Error(C.MOSQ_ERR_ACL_DENIED)
	MosqErrUnknown              = Error(C.MOSQ_ERR_UNKNOWN)
	MosqErrErrno                = Error(C.MOSQ_ERR_ERRNO)
	MosqErrEAI                  = Error(C.MOSQ_ERR_EAI)
	MosqErrProxy                = Error(C.MOSQ_ERR_PROXY)
	MosqErrPluginDefer          = Error(C.MOSQ_ERR_PLUGIN_DEFER)
	MosqErrMalformedUTF8        = Error(C.MOSQ_ERR_MALFORMED_UTF8)
	MosqErrKeepAlive            = Error(C.MOSQ_ERR_KEEPALIVE)
	MosqErrLookup               = Error(C.MOSQ_ERR_LOOKUP)
	MosqErrMalformedPacket      = Error(C.MOSQ_ERR_MALFORMED_PACKET)
	MosqErrDuplicateProperty    = Error(C.MOSQ_ERR_DUPLICATE_PROPERTY)
	MosqErrTLSHandshake         = Error(C.MOSQ_ERR_TLS_HANDSHAKE)
	MosqErrQOSNotSupported      = Error(C.MOSQ_ERR_QOS_NOT_SUPPORTED)
	MosqErrOversizePacket       = Error(C.MOSQ_ERR_OVERSIZE_PACKET)
	MosqErrOCSP                 = Error(C.MOSQ_ERR_OCSP)
	MosqErrTimeout              = Error(C.MOSQ_ERR_TIMEOUT)
	MosqErrRetainNotSupported   = Error(C.MOSQ_ERR_RETAIN_NOT_SUPPORTED)
	MosqErrTopicAliasInvalid    = Error(C.MOSQ_ERR_TOPIC_ALIAS_INVALID)
	MosqErrAdministrativeAction = Error(C.MOSQ_ERR_ADMINISTRATIVE_ACTION)
	MosqErrAlreadyExists        = Error(C.MOSQ_ERR_ALREADY_EXISTS)
)

var errorMap = map[Error]string{
	MosqErrAuthContinue:         "Authentication continue",
	MosqErrNoSubscribers:        "No subscribers",
	MosqErrSubExists:            "Subscription exists",
	MosqErrConnPending:          "Connection pending",
	MosqErrSuccess:              "Success",
	MosqErrNoMem:                "No memory",
	MosqErrProtocol:             "Protocol error",
	MosqErrInval:                "Invalid",
	MosqErrNoConn:               "No connection",
	MosqErrConnRefused:          "Connection refused",
	MosqErrNotFound:             "Not found",
	MosqErrConnLost:             "Connection lost",
	MosqErrTLS:                  "TLS",
	MosqErrPayloadSize:          "Payload size",
	MosqErrNotSupported:         "Not supported",
	MosqErrAuth:                 "Auth",
	MosqErrACLDenied:            "ACL denied",
	MosqErrUnknown:              "Unknown",
	MosqErrErrno:                "Errno",
	MosqErrEAI:                  "EAI",
	MosqErrProxy:                "Proxy",
	MosqErrPluginDefer:          "Plugin defer",
	MosqErrMalformedUTF8:        "Malformed UTF-8",
	MosqErrKeepAlive:            "Keep alive",
	MosqErrLookup:               "Lookup",
	MosqErrMalformedPacket:      "Malformed packet",
	MosqErrDuplicateProperty:    "Duplicate property",
	MosqErrTLSHandshake:         "TLS handshake",
	MosqErrQOSNotSupported:      "QoS not supported",
	MosqErrOversizePacket:       "Oversize packet",
	MosqErrOCSP:                 "OCSP",
	MosqErrTimeout:              "timeout",
	MosqErrRetainNotSupported:   "Retain not supported",
	MosqErrTopicAliasInvalid:    "Topic alias invalid",
	MosqErrAdministrativeAction: "Administrative action",
	MosqErrAlreadyExists:        "Already exists",
}
