package mosquitto

/*
#include <mosquitto_broker.h>
#include <mosquitto_plugin.h>
*/
import "C"

type Access int

func (a Access) String() string {
	return accessMap[a]
}

const (
	MosqAclNone        = Access(C.MOSQ_ACL_NONE)
	MosqAclRead        = Access(C.MOSQ_ACL_READ)
	MosqAclWrite       = Access(C.MOSQ_ACL_WRITE)
	MosqAclSubscribe   = Access(C.MOSQ_ACL_SUBSCRIBE)
	MosqAclUnsubscribe = Access(C.MOSQ_ACL_UNSUBSCRIBE)
)

var accessMap = map[Access]string{
	MosqAclNone:        "None",
	MosqAclRead:        "Read",
	MosqAclSubscribe:   "Subscribe",
	MosqAclWrite:       "Write",
	MosqAclUnsubscribe: "Unsubscribe",
}
