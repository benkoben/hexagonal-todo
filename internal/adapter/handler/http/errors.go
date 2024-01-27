package http

import "encoding/json"

const (
    errInternalServer = "Internal server error."
    errMethodNotAllowed = "Method not allowed."
    errNotAuthorized = "Not authorized." 
)

type Error struct {
    StatusCode int `json:"status_code"`
    Message string `json:"message"`
}

// newError creates a new Error
func newError(message string, statusCode int) Error {
    return Error{
        StatusCode: statusCode,
        Message: message,
    }
}

// JSON returns the JSON encoding of Error.
func (e Error) JSON() []byte {
    b, _ := json.Marshal(&e)
    return b
}

// Code return the error code of Error
func (e Error)Code() int {
    return e.StatusCode
}
