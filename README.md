# Go-Mosquitto-Plugin

This module implements the mosquitto v2 plugin interface as a Go module.

The new interface offers some interesting new features like message manipulation, and a callback whether a client
disconnects. The interface is far from complete, but does implement some key concepts like registering for events and 
forcefully disconnect clients.

* KickClientByUsername
* KickClientByClientID
* Publish
* CallbackRegister
* CallbackUnregister

This module also integrates with mosquittos internal logging mechanisms by implementing a logger that simply calls
`mosquitto_log_printf` under the hood.

There is a test plugin in plugins/test which is only supposed to be used for exploring or testing different aspects of
the interface and will not make any sense from a security perspective.

The skeleton plugin provides a minimal reference on how to create a plugin.

Plugins have to be built with `go build -buildmode=c-shared`
