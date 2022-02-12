package dto

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	GetItemError      = errors.New("cannot find item")
	WrongRequestError = errors.New("invalid input")
)

type DynamoDbError struct {
	Op  string
	Err error
}

func (e *DynamoDbError) Error() string {
	return fmt.Sprintf("%s: %s", e.Op, e.Err.Error())
}

type ValidationError struct {
	Field string
	Err   error
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Err.Error())
}

const (
	ValidationErrorCode = iota
	InternalServerErrorCode
)

type LambdaError struct {
	Code int
	Msg  string
}

func (e *LambdaError) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(b)
}
