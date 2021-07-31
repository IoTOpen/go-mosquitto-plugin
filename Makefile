plugin := $(wildcard *.go) interface.c
test/auth.so: $(plugin) plugins/test/plugin.go
	@go build -buildmode=c-shared -o test/auth.so github.com/iotopen/go-mosquitto-plugin/plugins/test 

.PHONY: test clean

test: test/auth.so
	@mosquitto -c ./test/mqtt.cfg

clean:
	@rm ./test/auth.so
