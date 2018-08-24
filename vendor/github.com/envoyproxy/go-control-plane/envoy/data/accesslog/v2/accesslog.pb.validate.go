// Code generated by protoc-gen-validate
// source: envoy/data/accesslog/v2/accesslog.proto
// DO NOT EDIT!!!

package envoy_data_accesslog_v2

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

// Validate checks the field values on TCPAccessLogEntry with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *TCPAccessLogEntry) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetCommonProperties()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return TCPAccessLogEntryValidationError{
				Field:  "CommonProperties",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// TCPAccessLogEntryValidationError is the validation error returned by
// TCPAccessLogEntry.Validate if the designated constraints aren't met.
type TCPAccessLogEntryValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e TCPAccessLogEntryValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTCPAccessLogEntry.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = TCPAccessLogEntryValidationError{}

// Validate checks the field values on HTTPAccessLogEntry with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HTTPAccessLogEntry) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetCommonProperties()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return HTTPAccessLogEntryValidationError{
				Field:  "CommonProperties",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	// no validation rules for ProtocolVersion

	if v, ok := interface{}(m.GetRequest()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return HTTPAccessLogEntryValidationError{
				Field:  "Request",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetResponse()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return HTTPAccessLogEntryValidationError{
				Field:  "Response",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// HTTPAccessLogEntryValidationError is the validation error returned by
// HTTPAccessLogEntry.Validate if the designated constraints aren't met.
type HTTPAccessLogEntryValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e HTTPAccessLogEntryValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHTTPAccessLogEntry.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = HTTPAccessLogEntryValidationError{}

// Validate checks the field values on AccessLogCommon with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *AccessLogCommon) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetSampleRate() > 1 {
		return AccessLogCommonValidationError{
			Field:  "SampleRate",
			Reason: "value must be less than or equal to 1",
		}
	}

	if v, ok := interface{}(m.GetDownstreamRemoteAddress()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				Field:  "DownstreamRemoteAddress",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDownstreamLocalAddress()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				Field:  "DownstreamLocalAddress",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTlsProperties()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				Field:  "TlsProperties",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetStartTime()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				Field:  "StartTime",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTimeToLastRxByte()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				Field:  "TimeToLastRxByte",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTimeToFirstUpstreamTxByte()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				Field:  "TimeToFirstUpstreamTxByte",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTimeToLastUpstreamTxByte()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				Field:  "TimeToLastUpstreamTxByte",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTimeToFirstUpstreamRxByte()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				Field:  "TimeToFirstUpstreamRxByte",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTimeToLastUpstreamRxByte()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				Field:  "TimeToLastUpstreamRxByte",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTimeToFirstDownstreamTxByte()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				Field:  "TimeToFirstDownstreamTxByte",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTimeToLastDownstreamTxByte()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				Field:  "TimeToLastDownstreamTxByte",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpstreamRemoteAddress()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				Field:  "UpstreamRemoteAddress",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpstreamLocalAddress()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				Field:  "UpstreamLocalAddress",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	// no validation rules for UpstreamCluster

	if v, ok := interface{}(m.GetResponseFlags()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				Field:  "ResponseFlags",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMetadata()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				Field:  "Metadata",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// AccessLogCommonValidationError is the validation error returned by
// AccessLogCommon.Validate if the designated constraints aren't met.
type AccessLogCommonValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e AccessLogCommonValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccessLogCommon.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = AccessLogCommonValidationError{}

// Validate checks the field values on ResponseFlags with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ResponseFlags) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for FailedLocalHealthcheck

	// no validation rules for NoHealthyUpstream

	// no validation rules for UpstreamRequestTimeout

	// no validation rules for LocalReset

	// no validation rules for UpstreamRemoteReset

	// no validation rules for UpstreamConnectionFailure

	// no validation rules for UpstreamConnectionTermination

	// no validation rules for UpstreamOverflow

	// no validation rules for NoRouteFound

	// no validation rules for DelayInjected

	// no validation rules for FaultInjected

	// no validation rules for RateLimited

	if v, ok := interface{}(m.GetUnauthorizedDetails()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return ResponseFlagsValidationError{
				Field:  "UnauthorizedDetails",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	// no validation rules for RateLimitServiceError

	return nil
}

// ResponseFlagsValidationError is the validation error returned by
// ResponseFlags.Validate if the designated constraints aren't met.
type ResponseFlagsValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e ResponseFlagsValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResponseFlags.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = ResponseFlagsValidationError{}

// Validate checks the field values on TLSProperties with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *TLSProperties) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for TlsVersion

	if v, ok := interface{}(m.GetTlsCipherSuite()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return TLSPropertiesValidationError{
				Field:  "TlsCipherSuite",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	// no validation rules for TlsSniHostname

	return nil
}

// TLSPropertiesValidationError is the validation error returned by
// TLSProperties.Validate if the designated constraints aren't met.
type TLSPropertiesValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e TLSPropertiesValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTLSProperties.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = TLSPropertiesValidationError{}

// Validate checks the field values on HTTPRequestProperties with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HTTPRequestProperties) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RequestMethod

	// no validation rules for Scheme

	// no validation rules for Authority

	if v, ok := interface{}(m.GetPort()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return HTTPRequestPropertiesValidationError{
				Field:  "Port",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	// no validation rules for Path

	// no validation rules for UserAgent

	// no validation rules for Referer

	// no validation rules for ForwardedFor

	// no validation rules for RequestId

	// no validation rules for OriginalPath

	// no validation rules for RequestHeadersBytes

	// no validation rules for RequestBodyBytes

	// no validation rules for RequestHeaders

	return nil
}

// HTTPRequestPropertiesValidationError is the validation error returned by
// HTTPRequestProperties.Validate if the designated constraints aren't met.
type HTTPRequestPropertiesValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e HTTPRequestPropertiesValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHTTPRequestProperties.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = HTTPRequestPropertiesValidationError{}

// Validate checks the field values on HTTPResponseProperties with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HTTPResponseProperties) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetResponseCode()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return HTTPResponsePropertiesValidationError{
				Field:  "ResponseCode",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	// no validation rules for ResponseHeadersBytes

	// no validation rules for ResponseBodyBytes

	// no validation rules for ResponseHeaders

	// no validation rules for ResponseTrailers

	return nil
}

// HTTPResponsePropertiesValidationError is the validation error returned by
// HTTPResponseProperties.Validate if the designated constraints aren't met.
type HTTPResponsePropertiesValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e HTTPResponsePropertiesValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHTTPResponseProperties.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = HTTPResponsePropertiesValidationError{}

// Validate checks the field values on ResponseFlags_Unauthorized with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ResponseFlags_Unauthorized) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Reason

	return nil
}

// ResponseFlags_UnauthorizedValidationError is the validation error returned
// by ResponseFlags_Unauthorized.Validate if the designated constraints aren't met.
type ResponseFlags_UnauthorizedValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e ResponseFlags_UnauthorizedValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResponseFlags_Unauthorized.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = ResponseFlags_UnauthorizedValidationError{}
