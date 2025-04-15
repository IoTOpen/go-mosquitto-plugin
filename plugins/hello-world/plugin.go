package main

import (
	"log"
	"time"

	"github.com/iotopen/go-mosquitto-plugin"
)

// Plugin structure
type Plugin struct {
	authChannel chan string
}

// Version callback
func (p *Plugin) Version(versions []int) int {
	return mosquitto.MosqPluginVersion
}

// Init callback
func (p *Plugin) Init(options mosquitto.Options) error {
	p.authChannel = make(chan string, 10)
	if err := mosquitto.RegisterBasicAuth(p.auth); err != nil {
		return err
	}
	if err := mosquitto.RegisterConnect(p.onConnect); err != nil {
		return err
	}
	if err := mosquitto.RegisterDisconnect(p.onDisconnect); err != nil {
		return err
	}
	if err := mosquitto.RegisterClientOffline(p.onOffline); err != nil {
		return err
	}
	if err := mosquitto.RegisterTick(p.tick); err != nil {
		return err
	}
	return mosquitto.MosqErrSuccess
}

func (p *Plugin) tick(ev mosquitto.EvtTick) error {
	select {
	case clientID := <-p.authChannel:
		log.Println("I: Completing authentication for client:", clientID)
		mosquitto.CompleteBasicAuth(clientID, mosquitto.MosqErrSuccess)
	default:
	}
	return nil
}

func (p *Plugin) onOffline(ev mosquitto.EvtClientOffline) error {
	log.Println("I: Client offline", ev.Client().ClientID())
	return nil
}

func (p *Plugin) onDisconnect(ev mosquitto.EvtDisconnect) error {
	log.Println("I: Client disconnected:", ev.Client().ClientID())
	return nil
}

func (p *Plugin) onConnect(ev mosquitto.EvtConnect) error {
	log.Println("I: Client connected:", ev.Client().ClientID())
	return nil
}

func (p *Plugin) auth(ev mosquitto.EvtBasicAuth) error {
	log.Println("I: Authentication Begin:", ev.Client().ClientID())
	go func(cid string) {
		time.Sleep(5 * time.Second)
		p.authChannel <- cid
	}(ev.Client().ClientID())
	return mosquitto.MosqErrAuthDelayed
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
