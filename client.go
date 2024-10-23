package mosquitto

/*
#include <malloc.h>
#include <mosquitto_broker.h>
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

type Client struct {
	ptr unsafe.Pointer
}

func (c Client) asStruct() *C.struct_mosquitto {
	return (*C.struct_mosquitto)(c.ptr)
}

func (c Client) Address() string {
	x := c.asStruct()
	res := C.mosquitto_client_address(x)
	return C.GoString(res)
}

func (c Client) CleanSession() bool {
	x := c.asStruct()
	res := C.mosquitto_client_clean_session(x)
	return bool(res)
}

func (c Client) ClientID() string {
	x := c.asStruct()
	res := C.mosquitto_client_id(x)
	return C.GoString(res)
}

func (c Client) KeepAlive() int {
	x := c.asStruct()
	res := C.mosquitto_client_keepalive(x)
	return int(res)
}

func (c Client) Protocol() Protocol {
	x := c.asStruct()
	res := C.mosquitto_client_protocol(x)
	return Protocol(res)
}

func (c Client) ProtocolVersion() ProtocolVersion {
	x := c.asStruct()
	res := C.mosquitto_client_protocol_version(x)
	return ProtocolVersion(res)
}

func (c Client) SubscriptionCount() int {
	x := c.asStruct()
	res := C.mosquitto_client_sub_count(x)
	return int(res)
}

func (c Client) Username() string {
	x := c.asStruct()
	res := C.mosquitto_client_username(x)
	return C.GoString(res)
}

func (c Client) SetUsername(name string) error {
	x := c.asStruct()
	var res = C.int(0)
	if name == "" {
		res = C.mosquitto_set_username(x, nil)
	} else {
		tmp := C.CString(name)
		res = C.mosquitto_set_username(x, tmp)
		C.free(unsafe.Pointer(tmp))
	}
	if !errors.Is(Error(res), MosqErrSuccess) {
		return fmt.Errorf("unable to set username: %d", int(res))
	}
	return nil
}

func (c Client) SetClientID(clientID string) error {
	x := c.asStruct()
	var res = C.int(0)
	tmp := C.CString(clientID)
	res = C.mosquitto_set_clientid(x, tmp)
	C.free(unsafe.Pointer(tmp))
	if !errors.Is(Error(res), MosqErrSuccess) {
		return fmt.Errorf("unable to set username: %d", int(res))
	}
	return nil
}
