#!/bin/bash
mosquitto_sub -v -t "#" \
	-p 8884 \
	-h localhost \
	--tls-use-os-certs \
	--cafile ../ca/rootCA.pem  \
	--cert ./client.crt \
	--key ./client.key

