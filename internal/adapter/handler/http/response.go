package http

import (
    "encoding/json"
)
// response wraps around JSON and Code.
type response interface {
    JSON() []byte
    Code() int
}

// Response wraps a message and a status code which is intended to be sent back to the client. Message takes in any json tagged struct.
type Response[T any ] struct {
    StatusCode int `json:"status_code"`
    Message T `json:"message"`
}

// newResponse creates a new response for any of the operations performed on service objects implemented by the todo-service.
func newResponse[T any ](message T, statusCode int) Response[T] {
    return Response[T]{
        StatusCode: statusCode,
        Message: message,
    }
}

// JSON returns the JSON encoding of Error.
func (e Response[_]) JSON() []byte {
    // generics wont be accessed in this method therefore we blank them
    b, _ := json.Marshal(&e)
    return b
}

// Code return the error code of Error
func (e Response[_])Code() int {
    // generics wont be accessed in this method therefore we blank them
    return e.StatusCode
}
