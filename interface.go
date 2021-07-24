package mosquitto

//#cgo LDFLAGS: -Wl,-unresolved-symbols=ignore-all
/*
#include <malloc.h>
#include <mosquitto.h>
#include <mosquitto_broker.h>
#include <mosquitto_plugin.h>

int go_mosquitto_generic_callback(int event, void* p1, void* p2);
bool go_mosquitto_topic_matches_sub(char* topic, char* subscription);
int mosquitto_callback_register2(void* id, int event, void* cb, void* eventData, void* userdata);
int mosquitto_callback_unregister2(void* id, int event, void* cb, void* eventData);
*/
import "C"
import (
	"reflect"
	"unsafe"
)

var (
	registerX = uintptr(0)
	register  = map[uintptr]interface{}{}
	pluginRef Plugin
)

type (
	PluginID uintptr
	Options  map[string]string

	Plugin interface {
		Version(versions []int) int
		Init(id PluginID, options Options) error
		Cleanup(options Options) error
	}
)

// RegisterPlugin must be called in an init function in the plugin.
func RegisterPlugin(p Plugin) {
	pluginRef = p
}

func optMap(options *C.struct_mosquitto_opt, optCount C.int) Options {
	var optionArray []C.struct_mosquitto_opt
	setSliceHeader((*reflect.SliceHeader)(unsafe.Pointer(&optionArray)), uintptr(unsafe.Pointer(options)), optCount)

	optMap := make(map[string]string)
	for _, v := range optionArray {
		optMap[C.GoString(v.key)] = C.GoString(v.value)
	}
	return optMap
}

func setSliceHeader(header *reflect.SliceHeader, data uintptr, size C.int) {
	header.Cap = int(size)
	header.Len = int(size)
	header.Data = data
}

const (
	MosqPluginVersion = C.MOSQ_PLUGIN_VERSION
)

//export goMosquittoPluginVersion
func goMosquittoPluginVersion(supportedVersionCount C.int, supportedVersions *C.int) C.int {
	var versions []C.int
	setSliceHeader((*reflect.SliceHeader)(unsafe.Pointer(&versions)), uintptr(unsafe.Pointer(supportedVersions)), supportedVersionCount)
	arg := make([]int, 0, len(versions))
	for _, v := range versions {
		arg = append(arg, int(v))
	}
	return C.int(pluginRef.Version(arg))
}

//export goMosquittoPluginInit
func goMosquittoPluginInit(identifier uintptr, options *C.struct_mosquitto_opt, optCount C.int) C.int {
	x := pluginRef.Init(PluginID(identifier), optMap(options, optCount))
	if x == nil {
		return C.int(MosqErrSuccess)
	}
	if e, ok := x.(Error); ok {
		return C.int(e)
	}
	return C.int(MosqErrUnknown)
}

//export goMosquittoPluginCleanup
func goMosquittoPluginCleanup(options *C.struct_mosquitto_opt, optCount C.int) C.int {
	x := pluginRef.Cleanup(optMap(options, optCount))
	if x == nil {
		return C.int(MosqErrSuccess)
	}
	if e, ok := x.(Error); ok {
		return C.int(e)
	}
	return C.int(MosqErrUnknown)
}

//export goGenericCallback
func goGenericCallback(event int, p1 unsafe.Pointer, p2 uintptr) C.int {
	if data, ok := register[p2]; ok {
		fn := reflect.ValueOf(data)
		args := make([]reflect.Value, 0, 1)
		switch Event(event) {
		case MosqEvtReload:
			args = append(args, reflect.ValueOf(EvtReload(p1)))
		case MosqEvtACLCheck:
			args = append(args, reflect.ValueOf(EvtAclCheck(p1)))
		case MosqEvtBasicAuth:
			args = append(args, reflect.ValueOf(EvtBasicAuth(p1)))
		case MosqEvtPSKKey:
			args = append(args, reflect.ValueOf(EvtPskKey(p1)))
		case MosqEvtEXTAuthStart, MosqEvtEXTAuthContinue:
			args = append(args, reflect.ValueOf(EvtExtendedAuth(p1)))
		case MosqEvtControl:
			args = append(args, reflect.ValueOf(EvtControl(p1)))
		case MosqEvtMessage:
			args = append(args, reflect.ValueOf(EvtMessage(p1)))
		case MosqEvtTick:
			args = append(args, reflect.ValueOf(EvtTick(p1)))
		case MosqEvtDisconnect:
			args = append(args, reflect.ValueOf(EvtDisconnect(p1)))
		case MosqEvtConnect:
			args = append(args, reflect.ValueOf(EvtConnect(p1)))
		default:
			return C.int(MosqErrUnknown)
		}
		res := fn.Call(args)
		var x = 0
		if len(res) > 0 {
			if res[0].IsNil() {
				x = int(MosqErrSuccess)
			} else if e, ok := res[0].Interface().(Error); ok {
				x = int(e)
			} else {
				x = int(MosqErrUnknown)
			}
		}
		return C.int(x)
	}
	return C.int(MosqErrUnknown)
}

/* TopicMatchesSub
 * Check whether a topic matches a subscription.
 *
 * For example:
 *
 * foo/bar would match the subscription foo/# or +/bar
 * non/matching would not match the subscription non/+/+
 *
 * Parameters:
 *	topic - topic to check.
 *	subscription - subscription string to check topic against.
 *
 * Returns:
 *	true - if topic matches subscription
 *  false - if topic doesn't match subscription, or if invalid parameters
 */
func TopicMatchesSub(topic, subscription string) bool {
	top := C.CString(topic)
	sub := C.CString(subscription)
	res := C.go_mosquitto_topic_matches_sub(top, sub)
	C.free(unsafe.Pointer(top))
	C.free(unsafe.Pointer(sub))
	return bool(res)
}

/* CallbackRegister
 * Register a callback for an event.
 *
 * Parameters:
 *  pluginID - the plugin identifier, as provided by <mosquitto_plugin_init>.
 *  event - the event to register a callback for. Can be one of:
 *          * MosqEvtReload
 *          * MosqEvtACLCheck
 *          * MosqEvtBasicAuth
 *          * MosqEvtEXTAuthStart
 *          * MosqEvtEXTAuthContinue
 *          * MosqEvtControl
 *          * MosqEvtMessage
 *          * MosqEvtPSKKey
 *          * MosqEvtTick
 *          * MosqEvtDisconnect
 *  cb - the callback function
 *  eventData - event specific data
 *
 * Returns:
 *	nil - on success
 *	MosqErrInval - if cb_func is NULL
 *	MosqErrAlreadyExists - if cb_func has already been registered for this event
 *	MosqErrNotSupported - if the event is not supported
 */
func CallbackRegister(pluginID PluginID, event Event, cb interface{}, eventData interface{}) error {
	tmp := registerX
	registerX++
	register[tmp] = cb
	ptr := unsafe.Pointer(uintptr(0))
	doFree := false
	switch tmp := eventData.(type) {
	case string:
		ptr = unsafe.Pointer(C.CString(tmp))
		doFree = true
	case uintptr:
		ptr = unsafe.Pointer(tmp)
	case unsafe.Pointer:
		ptr = tmp
	}
	x := C.mosquitto_callback_register2(unsafe.Pointer(pluginID), C.int(event), C.go_mosquitto_generic_callback, ptr, unsafe.Pointer(tmp))
	if doFree {
		C.free(ptr)
	}
	if Error(x) != MosqErrSuccess {
		return Error(x)
	}
	return nil
}

/* CallbackUnregister
 * Unregister a previously registered callback function.
 *
 * Parameters:
 *  pluginID - the plugin identifier, as provided by <mosquitto_plugin_init>.
 *  event - the event to register a callback for. Can be one of:
 *          * MosqEvtReload
 *          * MosqEvtACLCheck
 *          * MosqEvtBasicAuth
 *          * MosqEvtEXTAuthStart
 *          * MosqEvtEXTAuthContinue
 *          * MosqEvtControl
 *          * MosqEvtMessage
 *          * MosqEvtPSKKey
 *          * MosqEvtTick
 *          * MosqEvtDisconnect
 *  cb - the callback function
 *  eventData - event specific data
 *
 * Returns:
 *	nil - on success
 *	MosqErrInval - if cb_func is NULL
 *	MosqErrNotFound - if cb_func was not registered for this event
 *	MosqErrNotSupported - if the event is not supported
 */
func CallbackUnregister(pluginID PluginID, event Event, cb interface{}, eventData interface{}) error {
	ptr := unsafe.Pointer(uintptr(0))
	doFree := false
	switch tmp := eventData.(type) {
	case string:
		ptr = unsafe.Pointer(C.CString(tmp))
		doFree = true
	case uintptr:
		ptr = unsafe.Pointer(tmp)
	case unsafe.Pointer:
		ptr = tmp
	}
	x := C.mosquitto_callback_unregister2(unsafe.Pointer(pluginID), C.int(event), C.go_mosquitto_generic_callback, ptr)
	if doFree {
		C.free(ptr)
	}
	for key, v := range register {
		if reflect.ValueOf(v).Pointer() == reflect.ValueOf(cb).Pointer() {
			delete(register, key)
			break
		}
	}
	if Error(x) != MosqErrSuccess {
		return Error(x)
	}
	return nil
}

/* KickClientByClientID
 * Forcefully disconnect a client from the broker.
 *
 * If clientid != "", then the client with the matching client id is
 *   disconnected from the broker.
 * If clientid == "", then all clients are disconnected from the broker.
 *
 * If with_will == true, then if the client has a Last Will and Testament
 *   defined then this will be sent. If false, the LWT will not be sent.
 */
func KickClientByClientID(clientID string, withWill bool) int {
	var res C.int
	if clientID == "" {
		res = C.mosquitto_kick_client_by_clientid(nil, withWill == true)
	} else {
		str := C.CString(clientID)
		res = C.mosquitto_kick_client_by_clientid(str, withWill == true)
		C.free(unsafe.Pointer(str))
	}
	return int(res)
}

/* KickClientByUsername
 * Forcefully disconnect a client from the broker.
 *
 * If username != "", then all clients with a matching username are kicked
 *   from the broker.
 * If username == "", then all clients that do not have a username are
 *   kicked.
 *
 * If with_will == true, then if the client has a Last Will and Testament
 *   defined then this will be sent. If false, the LWT will not be sent.
 */
func KickClientByUsername(username string, withWill bool) int {
	var res C.int
	if username == "" {
		res = C.mosquitto_kick_client_by_username(nil, withWill == true)
	} else {
		str := C.CString(username)
		res = C.mosquitto_kick_client_by_username(str, withWill == true)
		C.free(unsafe.Pointer(str))
	}
	return int(res)
}

/* Publish
 * Publish a message from within a plugin.
 *
 * This function allows a plugin to publish a message. Messages published in
 * this way are treated as coming from the broker and so will not be passed to
 * `mosquitto_auth_acl_check(, MOSQ_ACL_WRITE, , )` for checking. Read access
 * will be enforced as normal for individual clients when they are due to
 * receive the message.
 *
 * It can be used to send messages to all clients that have a matching
 * subscription, or to a single client whether or not it has a matching
 * subscription.
 *
 * Parameters:
 *  clientID -   optional string. If set to "", the message is delivered to all
 *               clients. If non-empty, the message is delivered only to the
 *               client with the corresponding client id. If the client id
 *               specified is not connected, the message will be dropped.
 *  topic -      message topic.
 *  payload -    payload bytes.
 *  qos -        message QoS to use.
 *  retain -     should retain be set on the message. This does not apply if
 *               clientid is non-NULL.
 *  properties - MQTT v5 properties to attach to the message. If the function
 *               returns success, then properties is owned by the broker and
 *               will be freed at a later point.
 *
 * Returns:
 *   nil - on success
 *   MosqErrInval - if topic is NULL, if payloadlen < 0, if payloadlen > 0
 *                    and payload is NULL, if qos is not 0, 1, or 2.
 */
func Publish(clientID, topic string, payload []byte, qos int, retain bool) error {
	var payloadPtr unsafe.Pointer = nil
	if len(payload) > 0 {
		payloadPtr = C.CBytes(payload)
	}

	var cid *C.char
	if clientID != "" {
		cid = C.CString(clientID)
	}
	top := C.CString(topic)

	x := C.mosquitto_broker_publish(cid, top, C.int(len(payload)), payloadPtr, C.int(qos), retain == true, nil)

	C.free(unsafe.Pointer(cid))
	C.free(unsafe.Pointer(top))
	if Error(x) == MosqErrSuccess {
		return nil
	}
	if len(payload) > 0 {
		C.free(payloadPtr)
	}
	return Error(x)
}
