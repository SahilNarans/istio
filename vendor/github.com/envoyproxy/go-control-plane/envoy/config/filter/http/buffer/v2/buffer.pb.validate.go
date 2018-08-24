// Code generated by protoc-gen-validate
// source: envoy/config/filter/http/buffer/v2/buffer.proto
// DO NOT EDIT!!!

package v2

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on Buffer with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Buffer) Validate() error {
	if m == nil {
		return nil
	}

	if wrapper := m.GetMaxRequestBytes(); wrapper != nil {

		if wrapper.GetValue() <= 0 {
			return BufferValidationError{
				Field:  "MaxRequestBytes",
				Reason: "value must be greater than 0",
			}
		}

	}

	if m.GetMaxRequestTime() == nil {
		return BufferValidationError{
			Field:  "MaxRequestTime",
			Reason: "value is required",
		}
	}

	if d := m.GetMaxRequestTime(); d != nil {
		dur := *d

		gt := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur <= gt {
			return BufferValidationError{
				Field:  "MaxRequestTime",
				Reason: "value must be greater than 0s",
			}
		}

	}

	return nil
}

// BufferValidationError is the validation error returned by Buffer.Validate if
// the designated constraints aren't met.
type BufferValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e BufferValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBuffer.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = BufferValidationError{}

// Validate checks the field values on BufferPerRoute with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *BufferPerRoute) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Override.(type) {

	case *BufferPerRoute_Disabled:

		if m.GetDisabled() != true {
			return BufferPerRouteValidationError{
				Field:  "Disabled",
				Reason: "value must equal true",
			}
		}

	case *BufferPerRoute_Buffer:

		if m.GetBuffer() == nil {
			return BufferPerRouteValidationError{
				Field:  "Buffer",
				Reason: "value is required",
			}
		}

		if v, ok := interface{}(m.GetBuffer()).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return BufferPerRouteValidationError{
					Field:  "Buffer",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	default:
		return BufferPerRouteValidationError{
			Field:  "Override",
			Reason: "value is required",
		}

	}

	return nil
}

// BufferPerRouteValidationError is the validation error returned by
// BufferPerRoute.Validate if the designated constraints aren't met.
type BufferPerRouteValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e BufferPerRouteValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBufferPerRoute.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = BufferPerRouteValidationError{}
