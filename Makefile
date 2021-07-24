plugin := $(wildcard *.go) interface.c
plugins/test/auth.so: $(plugin) plugins/test/plugin.go
	@go build -buildmode=c-shared -o test/auth.so github.com/iotopen/go-mosquitto-plugin/plugins/test 

.PHONY: test clean plugin-%.so

clean:
	@rm *.so || true

test: plugins/test/auth.so
	@mosquitto -c ./test/mqtt.cfg

plugin-%.so:
	@go build -buildmode=c-shared -o plugin-$*.so github.com/iotopen/go-mosquitto-plugin/plugins/$*
