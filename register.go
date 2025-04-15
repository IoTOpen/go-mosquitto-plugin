package mosquitto

type ReloadCallback func(ev EvtReload) error

// RegisterReload subscribes to reload events.
// Called when the broker is sent a signal indicating it should
// reload its configuration.
func RegisterReload(cb ReloadCallback) error {
	return CallbackRegister(MosqEvtReload, cb, nil)
}
func UnregisterReload(cb ReloadCallback) error {
	return CallbackUnregister(MosqEvtReload, cb, nil)
}

type ACLCheckCallback func(ev EvtAclCheck) error

// RegisterACLCheck subscribes to ACL check events.
// Called when a publish/subscribe/unsubscribe command is received
// and the broker wants to check when the client is allowed to carry
// out this command.
func RegisterACLCheck(cb ACLCheckCallback) error {
	return CallbackRegister(MosqEvtACLCheck, cb, nil)
}

func UnregisterACLCheck(cb ACLCheckCallback) error {
	return CallbackUnregister(MosqEvtACLCheck, cb, nil)
}

type BasicAuthCallback func(ev EvtBasicAuth) error

// RegisterBasicAuth subscribes to basic authentication events.
// Called when a client connects to the broker, to allow the
// username/password/clientid to be authenticated.
func RegisterBasicAuth(cb BasicAuthCallback) error {
	return CallbackRegister(MosqEvtBasicAuth, cb, nil)
}
func UnregisterBasicAuth(cb BasicAuthCallback) error {
	return CallbackUnregister(MosqEvtBasicAuth, cb, nil)
}

type EXTAuthCallback func(ev EvtExtendedAuth) error

// RegisterExtendedAuthStart subscribes to extended authentication events.
// Called when an MQTT v5 client connects, if it is using extended
// authentication.
func RegisterExtendedAuthStart(cb EXTAuthCallback) error {
	return CallbackRegister(MosqEvtEXTAuthStart, cb, nil)
}

// RegisterExtendedAuthContinue subscribes to extended authentication events.
// Called when an MQTT v5 client connects, if it is using extended
// authentication.
func RegisterExtendedAuthContinue(cb EXTAuthCallback) error {
	return CallbackRegister(MosqEvtEXTAuthContinue, cb, nil)
}
func UnregisterExtendedAuthStart(cb EXTAuthCallback) error {
	return CallbackUnregister(MosqEvtEXTAuthStart, cb, nil)
}
func UnregisterExtendedAuthContinue(cb EXTAuthCallback) error {
	return CallbackUnregister(MosqEvtEXTAuthContinue, cb, nil)
}

type ControlCallback func(ev EvtControl) error

// RegisterControl subscribes to control events.
// Called on receipt of a $CONTROL message that the plugin has
// registered for.
func RegisterControl(cb ControlCallback, topic string) error {
	return CallbackRegister(MosqEvtControl, cb, topic)
}

func UnregisterControl(cb ControlCallback, topic string) error {
	return CallbackUnregister(MosqEvtControl, cb, topic)
}

type MessageCallback func(ev EvtMessage) error

// RegisterMessageIn subscribes to incoming message events.
// Called for each incoming PUBLISH message after it has been received
// and authorised. The contents of the message can be modified.
func RegisterMessageIn(cb MessageCallback) error {
	return CallbackRegister(MosqEvtMessageIn, cb, nil)
}
func UnregisterMessageIn(cb MessageCallback) error {
	return CallbackUnregister(MosqEvtMessageIn, cb, nil)
}

// RegisterMessageOut subscribes to outgoing message events.
// Called for each outgoing PUBLISH message after it has been authorised,
// but before it is sent to each subscribing client. The contents of the
// message can be modified.
func RegisterMessageOut(cb MessageCallback) error {
	return CallbackRegister(MosqEvtMessageOut, cb, nil)
}
func UnregisterMessageOut(cb MessageCallback) error {
	return CallbackUnregister(MosqEvtMessageOut, cb, nil)
}

type PSKKeyCallback func(ev EvtPskKey) error

// RegisterPSKKey subscribes to PSK key events.
// Called when a client connects with TLS-PSK and the broker needs
// the PSK information.
func RegisterPSKKey(cb PSKKeyCallback) error {
	return CallbackRegister(MosqEvtPSKKey, cb, nil)
}

func UnregisterPSKKey(cb PSKKeyCallback) error {
	return CallbackUnregister(MosqEvtPSKKey, cb, nil)
}

type TickCallback func(ev EvtTick) error

// RegisterTick subscribes to tick events.
// Called periodically in the event loop. At the moment this
// occurs at a regular frequency, but this should not be relied
// upon.
func RegisterTick(cb TickCallback) error {
	return CallbackRegister(MosqEvtTick, cb, nil)
}

func UnregisterTick(cb TickCallback) error {
	return CallbackUnregister(MosqEvtTick, cb, nil)
}

type DisconnectCallback func(ev EvtDisconnect) error

// RegisterDisconnect subscribes to disconnect events.
// Called when a client disconnects from the broker.
func RegisterDisconnect(cb DisconnectCallback) error {
	return CallbackRegister(MosqEvtDisconnect, cb, nil)
}

func UnregisterDisconnect(cb DisconnectCallback) error {
	return CallbackUnregister(MosqEvtDisconnect, cb, nil)
}

type ConnectCallback func(ev EvtConnect) error

// RegisterConnect subscribes to connect events.
// Called when a client has successfully connected to the broker,
// i.e. has been authenticated.
func RegisterConnect(cb ConnectCallback) error {
	return CallbackRegister(MosqEvtConnect, cb, nil)
}

func UnregisterConnect(cb ConnectCallback) error {
	return CallbackUnregister(MosqEvtConnect, cb, nil)
}

type SubscribeCallback func(ev EvtSubscribe) error

// RegisterSubscribe subscribes to subscribe events.
// Called when a client has made a successful subscription
// request, but before the subscription is applied. The
// subscription request can be modified, although this is not
// recommended.
func RegisterSubscribe(cb SubscribeCallback) error {
	return CallbackRegister(MosqEvtSubscribe, cb, nil)
}
func UnregisterSubscribe(cb SubscribeCallback) error {
	return CallbackUnregister(MosqEvtSubscribe, cb, nil)
}

type UnsubscribeCallback func(ev EvtUnsubscribe) error

// RegisterUnsubscribe subscribes to unsubscribe events.
// Called when a client has made a successful unsubscription
// request, but before the unsubscription is applied. The
// unsubscription request can be modified, although this is not
// recommended.
func RegisterUnsubscribe(cb UnsubscribeCallback) error {
	return CallbackRegister(MosqEvtUnsubscribe, cb, nil)
}
func UnregisterUnsubscribe(cb UnsubscribeCallback) error {
	return CallbackUnregister(MosqEvtUnsubscribe, cb, nil)
}

type ClientOfflineCallback func(ev EvtClientOffline) error

// RegisterClientOffline subscribes to client offline events.
func RegisterClientOffline(cb ClientOfflineCallback) error {
	return CallbackRegister(MosqEvtClientOffline, cb, nil)
}
func UnregisterClientOffline(cb ClientOfflineCallback) error {
	return CallbackUnregister(MosqEvtClientOffline, cb, nil)
}
