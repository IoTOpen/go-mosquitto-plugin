package main

import (
	"fmt"
	"github.com/iotopen/go-mosquitto-plugin"
	"log"
	"log/slog"
	"time"
)

type Plugin struct{}

func (p *Plugin) Version(versions []int) int {
	return mosquitto.MosqPluginVersion
}

func (p *Plugin) Init(options mosquitto.Options) error {
	log.Println("hello world", options)
	slog.Info("Hello world", "From", time.Now())
	if err := mosquitto.CallbackRegister(mosquitto.MosqEvtBasicAuth, auth, nil); err != nil {
		log.Println("Error:", err)
		return mosquitto.MosqErrUnknown
	}
	if err := mosquitto.CallbackRegister(mosquitto.MosqEvtMessage, onMessage, nil); err != nil {
		log.Println("Error:", err)
		return mosquitto.MosqErrUnknown
	}
	if err := mosquitto.CallbackRegister(mosquitto.MosqEvtDisconnect, onDisconnect, nil); err != nil {
		log.Println("Error:", err)
		return mosquitto.MosqErrUnknown
	}
	if err := mosquitto.CallbackRegister(mosquitto.MosqEvtACLCheck, aclCheck, nil); err != nil {
		log.Println("Error:", err)
		return mosquitto.MosqErrUnknown
	}
	if err := mosquitto.CallbackRegister(mosquitto.MosqEvtControl, control, "$CONTROL/test1"); err != nil {
		log.Println("Error:", err)
		return mosquitto.MosqErrUnknown
	}
	if err := mosquitto.CallbackRegister(mosquitto.MosqEvtControl, control, "$CONTROL/test2"); err != nil {
		log.Println("Error:", err)
		return mosquitto.MosqErrUnknown
	}
	if err := mosquitto.CallbackRegister(mosquitto.MosqEvtTick, ticker, nil); err != nil {
		log.Println("Error:", err)
		return mosquitto.MosqErrUnknown
	}
	return mosquitto.MosqErrSuccess
}

func (p *Plugin) Cleanup(options mosquitto.Options) error {
	log.Println("Cleanup!")
	if err := mosquitto.CallbackUnregister(mosquitto.MosqEvtBasicAuth, auth, nil); err != nil {
		log.Println("Error:", err)
		return mosquitto.MosqErrUnknown
	}
	if err := mosquitto.CallbackUnregister(mosquitto.MosqEvtControl, control, "$CONTROL/test1"); err != nil {
		log.Println("Error:", err)
		return mosquitto.MosqErrUnknown
	}
	if err := mosquitto.CallbackUnregister(mosquitto.MosqEvtControl, control, "$CONTROL/test2"); err != nil {
		log.Println("Error:", err)
		return mosquitto.MosqErrUnknown
	}
	return mosquitto.MosqErrSuccess
}

func ticker(evt mosquitto.EvtTick) error {
	return nil
}

func control(evt mosquitto.EvtControl) error {
	log.Printf("Control: %s - %s", evt.Topic(), evt.Payload())
	return nil
}

func onDisconnect(evt mosquitto.EvtDisconnect) {
	log.Println("Plugin - Client", evt.Client().ClientID(), "Disconnected")
}

func aclCheck(msg mosquitto.EvtAclCheck) error {
	log.Println("Plugin - AclCheck", msg.Client().ClientID(), msg.Client().Username(), msg.Access(), msg.Topic())
	return mosquitto.MosqErrSuccess
}

func onMessage(msg mosquitto.EvtMessage) error {
	log.Println("I: onMessage -", msg.Topic())
	msg.SetRetained(false)
	if len(msg.Payload()) != 0 {
		msg.SetPayload([]byte(fmt.Sprintf("%s says: %s", msg.Client().Username(), string(msg.Payload()))))
	}
	return nil
}

func auth(data mosquitto.EvtBasicAuth) error {
	log.Println("Auth attempt", data.Username(), "@", data.Client())
	data.Client().SetClientID("myclient")
	data.Client().SetUsername(data.Password())
	return nil
}

func init() {
	mosquitto.RegisterPlugin(&Plugin{})
}

func main() {}
