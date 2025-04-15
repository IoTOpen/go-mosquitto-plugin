plugin := $(wildcard *.go) interface.c ssl.c

test/auth.so: $(plugin) plugins/test/plugin.go
	@go build -buildmode=c-shared -o test/auth.so github.com/iotopen/go-mosquitto-plugin/plugins/test 
test/hello-world.so: $(plugin) plugins/hello-world/plugin.go
	@go build -buildmode=c-shared -o test/hello-world.so github.com/iotopen/go-mosquitto-plugin/plugins/hello-world

.PHONY: test clean

test: test/auth.so
	@mosquitto -c ./test/mqtt.cfg


clean:
	@rm ./test/auth.so
