// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: protobuf/api/vacation/messages.proto

package vacation

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on GetVacationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetVacationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetVacationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetVacationRequestMultiError, or nil if none found.
func (m *GetVacationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetVacationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for CountDays

	if len(errors) > 0 {
		return GetVacationRequestMultiError(errors)
	}

	return nil
}

// GetVacationRequestMultiError is an error wrapping multiple validation errors
// returned by GetVacationRequest.ValidateAll() if the designated constraints
// aren't met.
type GetVacationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetVacationRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetVacationRequestMultiError) AllErrors() []error { return m }

// GetVacationRequestValidationError is the validation error returned by
// GetVacationRequest.Validate if the designated constraints aren't met.
type GetVacationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetVacationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetVacationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetVacationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetVacationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetVacationRequestValidationError) ErrorName() string {
	return "GetVacationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetVacationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetVacationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetVacationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetVacationRequestValidationError{}

// Validate checks the field values on GetVacationResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetVacationResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetVacationResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetVacationResponseMultiError, or nil if none found.
func (m *GetVacationResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetVacationResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Approve

	if len(errors) > 0 {
		return GetVacationResponseMultiError(errors)
	}

	return nil
}

// GetVacationResponseMultiError is an error wrapping multiple validation
// errors returned by GetVacationResponse.ValidateAll() if the designated
// constraints aren't met.
type GetVacationResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetVacationResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetVacationResponseMultiError) AllErrors() []error { return m }

// GetVacationResponseValidationError is the validation error returned by
// GetVacationResponse.Validate if the designated constraints aren't met.
type GetVacationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetVacationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetVacationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetVacationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetVacationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetVacationResponseValidationError) ErrorName() string {
	return "GetVacationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetVacationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetVacationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetVacationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetVacationResponseValidationError{}
