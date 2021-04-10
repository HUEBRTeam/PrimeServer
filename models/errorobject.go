package models

import (
	"fmt"
	"runtime/debug"
)

var stackEnabled = true

func EnableStackTrace() {
	stackEnabled = true
}

func DisableStackTrace() {
	stackEnabled = false
}

type ErrorObject struct {
	ErrorCode  string      `json:"errorCode"`
	ErrorField string      `json:"errorField"`
	Message    string      `json:"message"`
	ErrorData  interface{} `json:"errorData"`
	StackTrace string      `json:"stackTrace"`
}

func NewErrorObject(errorCode, errorField, message string, errorData interface{}) *ErrorObject {
	return &ErrorObject{
		ErrorCode:  errorCode,
		ErrorField: errorField,
		ErrorData:  errorData,
		Message:    message,
		StackTrace: string(debug.Stack()),
	}
}

func (e *ErrorObject) Error() string {
	return e.Message
}

func (e *ErrorObject) String() string {
	o := fmt.Sprintf("Error: %s\n", e.Message)
	o += fmt.Sprintf("  Error Code: %s\n", e.ErrorCode)
	o += fmt.Sprintf("  Error Field: %s\n", e.ErrorField)
	o += fmt.Sprintf("  Error Data: %v\n", e.ErrorData)
	o += fmt.Sprintf("  Stack Trace %s\n", e.StackTrace)
	return o
}
