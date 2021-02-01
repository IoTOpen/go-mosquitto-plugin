package main

import (
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
	return mosquitto.MosqErrSuccess
}

// Cleanup callback
func (p *Plugin) Cleanup(options mosquitto.Options) error {
	return mosquitto.MosqErrSuccess
}

// Register the plugin in an init function
func init() {
	mosquitto.RegisterPlugin(&Plugin{})
}

// Modules also required a main function.
func main() {}
