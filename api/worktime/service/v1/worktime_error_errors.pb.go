// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsUnknownError(err error) bool {
	e := errors.FromError(err)
	return e.Reason == WorkTimeServiceErrorReason_UNKNOWN_ERROR.String() && e.Code == 500
}

func ErrorUnknownError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, WorkTimeServiceErrorReason_UNKNOWN_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsRecordExisted(err error) bool {
	e := errors.FromError(err)
	return e.Reason == WorkTimeServiceErrorReason_RECORD_EXISTED.String() && e.Code == 406
}

func ErrorRecordExisted(format string, args ...interface{}) *errors.Error {
	return errors.New(406, WorkTimeServiceErrorReason_RECORD_EXISTED.String(), fmt.Sprintf(format, args...))
}