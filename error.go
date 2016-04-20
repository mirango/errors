package errors

import (
	"fmt"
)

type Error struct {
	Status  int
	Code    string
	Type    string
	Message string
	Values  map[string]interface{}
}

func New(status int, msg string) *Error {
	return &Error{Status: status, Message: msg}
}

func (e *Error) GetCode() string {
	return e.Code
}

func (e *Error) GetType() string {
	return e.Type
}

func (e *Error) GetMessage() string {
	return e.Message
}

func (e *Error) GetStatus() int {
	return e.Status
}

func (e *Error) SetValue(key string, value interface{}) {
	if e.Values == nil {
		e.Values = map[string]interface{}{}
	}
	e.Values[key] = value
}

func (e *Error) Error() string {
	if e.Type == "" && e.Code == "" {
		return fmt.Sprintf("[Error] %s", e.Message)
	} else if e.Type == "" {
		return fmt.Sprintf("[%s Error] %s", e.Type, e.Message)
	} else if e.Code == "" {
		return fmt.Sprintf("[Error: %s] %s", e.Code, e.Message)
	}
	return fmt.Sprintf("[%s Error: %s] %s", e.Type, e.Code, e.Message)
}
