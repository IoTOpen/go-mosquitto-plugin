package mosquitto

//#cgo LDFLAGS: -Wl,-unresolved-symbols=ignore-all
//#cgo CFLAGS: -I /usr/include -fPIC
/*
#include <malloc.h>
#include <mosquitto.h>

int go_mosquitto_generic_callback(int event, void* p1, void* p2);
bool go_mosquitto_topic_matches_sub(char* topic, char* subscription);
bool go_mosquitto_sub_matches_acl(char* acl, char* sub);
int mosquitto_callback_register2(mosquitto_plugin_id_t *id, int event, void* cb, void* eventData, uintptr_t userdata);
int mosquitto_callback_unregister2(mosquitto_plugin_id_t *id, int event, void* cb, void* eventData);
*/
import "C"
import (
	"errors"
	"reflect"
	"unsafe"
)

var (
	registerX = uintptr(0)
	register  = map[uintptr]interface{}{}
	pluginRef Plugin
	pluginID  *C.mosquitto_plugin_id_t
)

type (
	PluginID uintptr
	Options  map[string]string

	Plugin interface {
		Version(versions []int) int
		Init(options Options) error
		Cleanup(options Options) error
	}
)

// RegisterPlugin must be called in an init function in the plugin.
func RegisterPlugin(p Plugin) {
	pluginRef = p
}

func optMap(options *C.struct_mosquitto_opt, optCount C.int) Options {
	var optionArray []C.struct_mosquitto_opt = unsafe.Slice(options, int(optCount))
	optMap := make(map[string]string)
	for _, v := range optionArray {
		optMap[C.GoString(v.key)] = C.GoString(v.value)
	}
	return optMap
}

const (
	MosqPluginVersion = int(C.MOSQ_PLUGIN_VERSION)
)

//export goMosquittoPluginVersion
func goMosquittoPluginVersion(supportedVersionCount C.int, supportedVersions *C.int) C.int {
	var versions []C.int = unsafe.Slice(supportedVersions, int(supportedVersionCount))
	arg := make([]int, 0, len(versions))
	for _, v := range versions {
		arg = append(arg, int(v))
	}
	return C.int(pluginRef.Version(arg))
}

//export goMosquittoPluginInit
func goMosquittoPluginInit(identifier *C.mosquitto_plugin_id_t, options *C.struct_mosquitto_opt, optCount C.int) C.int {
	pluginID = identifier
	x := pluginRef.Init(optMap(options, optCount))
	if x == nil {
		return C.int(MosqErrSuccess)
	}
	var e Error
	if errors.As(x, &e) {
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
	var e Error
	if errors.As(x, &e) {
		return C.int(e)
	}
	return C.int(MosqErrUnknown)
}

//export goGenericCallback
func goGenericCallback(event int, callbackData unsafe.Pointer, regNo uintptr) C.int {
	if data, ok := register[regNo]; ok {
		fn := reflect.ValueOf(data)
		args := make([]reflect.Value, 0, 1)
		switch Event(event) {
		case MosqEvtReload:
			args = append(args, reflect.ValueOf(EvtReload{callbackData}))
		case MosqEvtACLCheck:
			args = append(args, reflect.ValueOf(EvtAclCheck{callbackData}))
		case MosqEvtBasicAuth:
			args = append(args, reflect.ValueOf(EvtBasicAuth{callbackData}))
		case MosqEvtPSKKey:
			args = append(args, reflect.ValueOf(EvtPskKey{callbackData}))
		case MosqEvtEXTAuthStart, MosqEvtEXTAuthContinue:
			args = append(args, reflect.ValueOf(EvtExtendedAuth{callbackData}))
		case MosqEvtControl:
			args = append(args, reflect.ValueOf(EvtControl{callbackData}))
		case MosqEvtMessageIn, MosqEvtMessageOut:
			args = append(args, reflect.ValueOf(EvtMessage{callbackData}))
		case MosqEvtTick:
			args = append(args, reflect.ValueOf(EvtTick{callbackData}))
		case MosqEvtDisconnect:
			args = append(args, reflect.ValueOf(EvtDisconnect{callbackData}))
		case MosqEvtConnect:
			args = append(args, reflect.ValueOf(EvtConnect{callbackData}))
		case MosqEvtSubscribe:
			args = append(args, reflect.ValueOf(EvtSubscribe{callbackData}))
		case MosqEvtUnsubscribe:
			args = append(args, reflect.ValueOf(EvtUnsubscribe{callbackData}))
		case MosqEvtClientOffline:
			args = append(args, reflect.ValueOf(EvtClientOffline{callbackData}))
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

// TopicMatchesSub
// Check whether a topic matches a subscription.
//
// For example:
//
// foo/bar would match the subscription foo/# or +/bar
// non/matching would not match the subscription non/+/+
//
// Parameters:
//
//	topic - topic to check.
//	subscription - subscription string to check topic against.
//
// Returns:
//
//		true - if topic matches subscription
//	 false - if topic doesn't match subscription, or if invalid parameters
func TopicMatchesSub(topic, subscription string) bool {
	top := C.CString(topic)
	sub := C.CString(subscription)
	res := C.go_mosquitto_topic_matches_sub(top, sub)
	C.free(unsafe.Pointer(top))
	C.free(unsafe.Pointer(sub))
	return bool(res)
}

func SubMatchesACL(sub, acl string) bool {
	subPtr := C.CString(sub)
	aclPtr := C.CString(acl)
	res := C.go_mosquitto_sub_matches_acl(subPtr, aclPtr)
	C.free(unsafe.Pointer(subPtr))
	C.free(unsafe.Pointer(aclPtr))
	return bool(res)
}

// CallbackRegister
// Register a callback for an event.
//
// Parameters:
//
//	pluginID - the plugin identifier, as provided by <mosquitto_plugin_init>.
//	event - the event to register a callback for. Can be one of:
//	        * MosqEvtReload
//	        * MosqEvtACLCheck
//	        * MosqEvtBasicAuth
//	        * MosqEvtEXTAuthStart
//	        * MosqEvtEXTAuthContinue
//	        * MosqEvtControl
//	        * MosqEvtMessage
//	        * MosqEvtPSKKey
//	        * MosqEvtTick
//	        * MosqEvtDisconnect
//	cb - the callback function
//	eventData - event specific data
//
// Returns:
//
//	nil - on success
//	MosqErrInval - if cb_func is NULL
//	MosqErrAlreadyExists - if cb_func has already been registered for this event
//	MosqErrNotSupported - if the event is not supported
func CallbackRegister(event Event, cb, eventData any) error {
	regNo := registerX
	registerX++
	register[regNo] = cb
	var ptr unsafe.Pointer
	doFree := false
	switch tmp := eventData.(type) {
	case nil:
	case string:
		ptr = unsafe.Pointer(C.CString(tmp))
		doFree = true
	case unsafe.Pointer:
		ptr = tmp
	default:
		panic("Register: couldn't pass eventdata, bad type")
	}
	x := C.mosquitto_callback_register2(pluginID, C.int(event), C.go_mosquitto_generic_callback, ptr, C.uintptr_t(regNo))
	if doFree {
		C.free(ptr)
	}
	if !errors.Is(Error(x), MosqErrSuccess) {
		return Error(x)
	}
	return nil
}

// CallbackUnregister
// Unregister a previously registered callback function.
//
// Parameters:
//
//	pluginID - the plugin identifier, as provided by <mosquitto_plugin_init>.
//	event - the event to register a callback for. Can be one of:
//	        * MosqEvtReload
//	        * MosqEvtACLCheck
//	        * MosqEvtBasicAuth
//	        * MosqEvtEXTAuthStart
//	        * MosqEvtEXTAuthContinue
//	        * MosqEvtControl
//	        * MosqEvtMessage
//	        * MosqEvtPSKKey
//	        * MosqEvtTick
//	        * MosqEvtDisconnect
//	cb - the callback function
//	eventData - event specific data
//
// Returns:
//
//	nil - on success
//	MosqErrInval - if cb_func is NULL
//	MosqErrNotFound - if cb_func was not registered for this event
//	MosqErrNotSupported - if the event is not supported
func CallbackUnregister(event Event, cb, eventData any) error {
	var ptr unsafe.Pointer
	doFree := false
	switch tmp := eventData.(type) {
	case nil:
	case string:
		ptr = unsafe.Pointer(C.CString(tmp))
		doFree = true
	case unsafe.Pointer:
		ptr = tmp
	default:
		panic("Unregister: couldn't pass eventData, bad type")
	}
	x := C.mosquitto_callback_unregister2(pluginID, C.int(event), C.go_mosquitto_generic_callback, ptr)
	if doFree {
		C.free(ptr)
	}
	for key, v := range register {
		if reflect.ValueOf(v).Pointer() == reflect.ValueOf(cb).Pointer() {
			delete(register, key)
			break
		}
	}
	if !errors.Is(Error(x), MosqErrSuccess) {
		return Error(x)
	}
	return nil
}

// KickClientByClientID
// Forcefully disconnect a client from the broker.
//
// If clientid != "", then the client with the matching client id is
//
//	disconnected from the broker.
//
// If clientid == "", then all clients are disconnected from the broker.
//
// If with_will == true, then if the client has a Last Will and Testament
//
//	defined then this will be sent. If false, the LWT will not be sent.
func KickClientByClientID(clientID string, withWill bool) int {
	var res C.int
	if clientID == "" {
		res = C.mosquitto_kick_client_by_clientid(nil, C._Bool(withWill))
	} else {
		str := C.CString(clientID)
		res = C.mosquitto_kick_client_by_clientid(str, C._Bool(withWill))
		C.free(unsafe.Pointer(str))
	}
	return int(res)
}

// KickClientByUsername
// Forcefully disconnect a client from the broker.
//
// If username != "", then all clients with a matching username are kicked
//
//	from the broker.
//
// If username == "", then all clients that do not have a username are
//
//	kicked.
//
// If with_will == true, then if the client has a Last Will and Testament
//
//	defined then this will be sent. If false, the LWT will not be sent.
func KickClientByUsername(username string, withWill bool) int {
	var res C.int
	if username == "" {
		res = C.mosquitto_kick_client_by_username(nil, C._Bool(withWill))
	} else {
		str := C.CString(username)
		res = C.mosquitto_kick_client_by_username(str, C._Bool(withWill))
		C.free(unsafe.Pointer(str))
	}
	return int(res)
}

// Publish a message from within a plugin.
//
// This function allows a plugin to publish a message. Messages published in
// this way are treated as coming from the broker and so will not be passed to
// `mosquitto_auth_acl_check(, MOSQ_ACL_WRITE, , )` for checking. Read access
// will be enforced as normal for individual clients when they are due to
// receive the message.
//
// It can be used to send messages to all clients that have a matching
// subscription, or to a single client whether or not it has a matching
// subscription.
//
// Parameters:
//
//	clientID -   optional string. If set to "", the message is delivered to all
//	             clients. If non-empty, the message is delivered only to the
//	             client with the corresponding client id. If the client id
//	             specified is not connected, the message will be dropped.
//	topic -      message topic.
//	payload -    payload bytes.
//	qos -        message QoS to use.
//	retain -     should retain be set on the message. This does not apply if
//	             clientid is non-NULL.
//	properties - MQTT v5 properties to attach to the message. If the function
//	             returns success, then properties is owned by the broker and
//	             will be freed at a later point.
//
// Returns:
//
//	nil - on success
//	MosqErrInval - if topic is NULL, if payloadlen < 0, if payloadlen > 0
//	                 and payload is NULL, if qos is not 0, 1, or 2.
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

	x := C.mosquitto_broker_publish(cid, top, C.int(len(payload)), payloadPtr, C.int(qos), C._Bool(retain), nil)

	C.free(unsafe.Pointer(cid))
	C.free(unsafe.Pointer(top))
	if errors.Is(Error(x), MosqErrSuccess) {
		return nil
	}
	if len(payload) > 0 {
		C.free(payloadPtr)
	}
	return Error(x)
}

// GetClient returns a Client object for the given clientID.
// Retrieve the mosquitto client for a client id.
// If the client is not connected, ok will be false and the client is unusable.
func GetClient(clientID string) (client Client, ok bool) {
	var cid *C.char
	if clientID != "" {
		cid = C.CString(clientID)
	}
	x := C.mosquitto_client(cid)
	C.free(unsafe.Pointer(cid))
	if x == nil {
		return Client{}, false
	}
	return Client{unsafe.Pointer(x)}, true
}

// CompleteBasicAuth Complete a delayed authentication request.
// Useful for plugins that subscribe to the MosqEvtBasicAuth event. If your
// plugin makes authentication requests that are not "instant", in particular
// if they communicate with an external service, then instead of blocking for a
// reply and returning MosqErrSuccess or MosqErrAuth, the plugin can return
// MosqErrAuthDelayed. This means that the plugin is promising to tell the
// broker the authentication result in the future. Once the plugin has an
// answer, it should call CompleteBasicAuth passing the client
// id and the result.
//
// Result:
//
// MosqErrSuccess - the client successfully authenticated
// MosqErrAuth - the client authentication failed
//
// Other error codes can be used if more appropriate, and the client connection
// will still be rejected, e.g. MosqErrNoMem.
//
// The plugin may use extra threads to handle the authentication requests, but
// the call to CompleteBasicAuth must happen in the main
// mosquitto thread. Using the MosqEvtTick event for this is suggested.
func CompleteBasicAuth(clientID string, result error) {
	var cid *C.char
	if clientID != "" {
		cid = C.CString(clientID)
	}
	var err Error
	if result == nil {
		err = MosqErrSuccess
	}
	if !errors.As(result, &err) {
		err = MosqErrAuth
	}
	C.mosquitto_complete_basic_auth(cid, C.int(err))
	C.free(unsafe.Pointer(cid))
}
