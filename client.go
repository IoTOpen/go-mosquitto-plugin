package mosquitto

/*
#include <malloc.h>
#include <mosquitto_broker.h>

char* x509_to_pem(void *cert);
char* x509_to_der(void *cert, int *der_length);
char* convert_x509(void* cert, int *der_length);
*/
import "C"
import (
	"crypto/x509"
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

func (c Client) X509() *x509.Certificate {
	x := c.asStruct()
	var size C.int
	certPointer := C.mosquitto_client_certificate(x)
	derPointer := C.convert_x509(certPointer, &size)
	if derPointer == nil {
		return nil
	}
	defer C.free(unsafe.Pointer(derPointer))
	derData := unsafe.Slice((*byte)(unsafe.Pointer(derPointer)), int(size))
	cert, err := x509.ParseCertificate(derData)
	if err != nil {
		return nil
	}
	return cert
}
