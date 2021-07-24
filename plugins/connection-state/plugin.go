package main

import (
	"fmt"
	"github.com/iotopen/go-mosquitto-plugin"
)

// Plugin structure
type Plugin struct {
	id mosquitto.PluginID
}

// Version callback
func (p *Plugin) Version(versions []int) int {
	return mosquitto.MosqPluginVersion
}

// Init callback
func (p *Plugin) Init(id mosquitto.PluginID, options mosquitto.Options) error {
	p.id = id

	if err := mosquitto.CallbackRegister(id, mosquitto.MosqEvtConnect, p.ConnectCallback, nil); err != nil {
		return err
	}
	if err := mosquitto.CallbackRegister(id, mosquitto.MosqEvtDisconnect, p.DisconnectCallback, nil); err != nil {
		return err
	}

	return mosquitto.MosqErrSuccess
}

// Cleanup callback
func (p *Plugin) Cleanup(options mosquitto.Options) error {
	_ = mosquitto.CallbackUnregister(p.id, mosquitto.MosqEvtConnect, p.ConnectCallback, nil)
	return mosquitto.CallbackUnregister(p.id, mosquitto.MosqEvtDisconnect, p.DisconnectCallback, nil)
}


func (p *Plugin) ConnectCallback(connect mosquitto.EvtConnect) error {
	topic := fmt.Sprintf("$SYS/broker/connection/client/%s/state", connect.Client().ClientID())
	_ = mosquitto.Publish("", topic, []byte("1"), 0, true)
	return mosquitto.MosqErrSuccess
}

func (p *Plugin) DisconnectCallback(connect mosquitto.EvtDisconnect) error {
	topic := fmt.Sprintf("$SYS/broker/connection/client/%s/state", connect.Client().ClientID())
	_ = mosquitto.Publish("", topic, []byte("0"), 0, true)
	return mosquitto.MosqErrSuccess
}


// Register the plugin in an init function
func init() {
	mosquitto.RegisterPlugin(&Plugin{})
}

// Modules also required a main function.
func main() {}