package client

/*
#cgo linux CFLAGS: -I/usr/local/include
#cgo linux LDFLAGS: -L/usr/local/lib -ltdjson_static -ltdjson_private -ltdclient -ltdcore -ltdapi -ltdactor -ltddb -ltdsqlite -ltdnet -ltdutils -lgcc -lssl -lcrypto -ldl -lz -lm
#cgo darwin CFLAGS: -I/usr/local/include
#cgo darwin LDFLAGS: -L/usr/local/lib -L/usr/local/opt/openssl/lib -ltdjson_static -ltdjson_private -ltdclient -ltdcore -ltdapi -ltdactor -ltddb -ltdsqlite -ltdnet -ltdutils -lc++ -lssl -lcrypto -ldl -lz -lm
#cgo windows CFLAGS: -I${SRCDIR}/../../td -I${SRCDIR}/../../td/build
#cgo windows LDFLAGS: -L${SRCDIR}/../../td/tdlib/lib -ltdjson
#include <stdlib.h>
#include <td/telegram/td_json_client.h>
*/
import "C"

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
	"unsafe"
)

type JsonClient struct {
	id int
}

func NewJsonClient() *JsonClient {
	return &JsonClient{
		id: int(C.td_create_client_id()),
	}
}

//Sends request to the TDLib client. May be called from any thread.
func (jsonClient *JsonClient) Send(req Request) {
	data, _ := json.Marshal(req)

	query := C.CString(string(data))
	defer C.free(unsafe.Pointer(query))

	C.td_send(C.int(jsonClient.id), query)
}

// Receives incoming updates and request responses from the TDLib client. May be called from any thread, but
// shouldn't be called simultaneously from two different threads.
// Returned pointer will be deallocated by TDLib during next call to td_json_client_receive or td_json_client_execute
// in the same thread, so it can't be used after that.
func (jsonClient *JsonClient) Receive(timeout time.Duration) (*Response, error) {
	result := C.td_receive(C.double(float64(timeout) / float64(time.Second)))
	if result == nil {
		return nil, errors.New("update receiving timeout")
	}

	data := []byte(C.GoString(result))

	var resp Response

	err := json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}

	if resp.ClientId != jsonClient.id {
		return nil, errors.New("wrong @client_id")
	}

	resp.Data = data

	return &resp, nil
}

// Synchronously executes TDLib request. May be called from any thread.
// Only a few requests can be executed synchronously.
// Returned pointer will be deallocated by TDLib during next call to td_json_client_receive or td_json_client_execute
// in the same thread, so it can't be used after that.
func (jsonClient *JsonClient) Execute(req Request) (*Response, error) {
	data, _ := json.Marshal(req)

	query := C.CString(string(data))
	defer C.free(unsafe.Pointer(query))
	result := C.td_execute(query)
	if result == nil {
		return nil, errors.New("request can't be parsed")
	}

	data = []byte(C.GoString(result))

	var resp Response

	err := json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}

	resp.Data = data

	return &resp, nil
}

type meta struct {
	Type     string `json:"@type"`
	Extra    string `json:"@extra"`
	ClientId int    `json:"@client_id"`
}

type Request struct {
	meta
	Data map[string]interface{}
}

func (req Request) MarshalJSON() ([]byte, error) {
	req.Data["@type"] = req.Type
	req.Data["@extra"] = req.Extra

	return json.Marshal(req.Data)
}

type Response struct {
	meta
	Data json.RawMessage
}

type ResponseError struct {
	Err *Error
}

func (responseError ResponseError) Error() string {
	return fmt.Sprintf("%d %s", responseError.Err.Code, responseError.Err.Message)
}

func buildResponseError(data json.RawMessage) error {
	respErr, err := UnmarshalError(data)
	if err != nil {
		return err
	}

	return ResponseError{
		Err: respErr,
	}
}

// JsonInt64 alias for int64, in order to deal with json big number problem
type JsonInt64 int64

// MarshalJSON marshals to json
func (jsonInt64 JsonInt64) MarshalJSON() ([]byte, error) {
	return []byte(`"` + strconv.FormatInt(int64(jsonInt64), 10) + `"`), nil
}

// UnmarshalJSON unmarshals from json
func (jsonInt64 *JsonInt64) UnmarshalJSON(data []byte) error {
	if len(data) > 2 && data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}

	jsonBigInt, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}

	*jsonInt64 = JsonInt64(jsonBigInt)

	return nil
}

type Type interface {
	GetType() string
	GetClass() string
}
