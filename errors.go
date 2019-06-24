
package avsm

import "fmt"
// ErrorCode is the type for package-specific error codes. This is used
// within the Error struct, which allows you to programatically determine
// the error cause.
type ErrorCode uint

func (e ErrorCode) String() string {
	switch e {
	case ErrorVehicleNotInitialized:
		return "VehicleNotInitialized"
	case ErrorTransitionNotPermitted:
		return "TransitionNotPermitted"
	case ErrorRolePermissionDenied:
		return "RolePermissionDenied"	
	case ErrorStateUndefined:
		return "StateUndefined"
	default:
		return "Unknown"
	}
}

const (
	// ErrorUnknown is the default value
	ErrorUnknown ErrorCode = iota

	// ErrorVehicleNotInitialized is an error returned when actions are taken on
	// a vehicle before it has been initialized. 
	ErrorVehicleNotInitialized

	// ErrorTransitionNotPermitted is the error returned when trying to
	// transition to an invalid state. 
	ErrorTransitionNotPermitted

	// ErrorRolePermissionDenied is the error returned when trying to
	// transition to a valid state but user dont have valid permission
	ErrorRolePermissionDenied

	// ErrorStateUndefined is the error returned when the requested state is
	// not defined within the machine.
	ErrorStateUndefined
)

// Error is the struct representing internal errors.
// This implements the error interface
type Error struct {
	message string
	code    ErrorCode
}

// newErrorStruct uses messge and code to create an *Error struct. The *Error
// struct implements the 'error' interface, so it should be able to be used
// wherever 'error' is expected.
func newErrorStruct(message string, code ErrorCode) *Error {
	return &Error{
		message: message,
		code:    code,
	}
}

// Message returns the error message.
func (e *Error) Message() string { return e.message }

// Code returns the error code.
func (e *Error) Code() ErrorCode { return e.code }

func (e *Error) Error() string {
	return fmt.Sprintf("%s (%d): %s", e.code, e.code, e.message)
}