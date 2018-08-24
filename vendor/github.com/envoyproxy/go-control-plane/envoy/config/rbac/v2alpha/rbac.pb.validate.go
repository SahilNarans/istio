// Code generated by protoc-gen-validate
// source: envoy/config/rbac/v2alpha/rbac.proto
// DO NOT EDIT!!!

package v2alpha

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

// Validate checks the field values on RBAC with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *RBAC) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Action

	// no validation rules for Policies

	return nil
}

// RBACValidationError is the validation error returned by RBAC.Validate if the
// designated constraints aren't met.
type RBACValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e RBACValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRBAC.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = RBACValidationError{}

// Validate checks the field values on Policy with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Policy) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetPermissions()) < 1 {
		return PolicyValidationError{
			Field:  "Permissions",
			Reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetPermissions() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return PolicyValidationError{
					Field:  fmt.Sprintf("Permissions[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	if len(m.GetPrincipals()) < 1 {
		return PolicyValidationError{
			Field:  "Principals",
			Reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetPrincipals() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return PolicyValidationError{
					Field:  fmt.Sprintf("Principals[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// PolicyValidationError is the validation error returned by Policy.Validate if
// the designated constraints aren't met.
type PolicyValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e PolicyValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPolicy.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = PolicyValidationError{}

// Validate checks the field values on Permission with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Permission) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Rule.(type) {

	case *Permission_AndRules:

		if v, ok := interface{}(m.GetAndRules()).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return PermissionValidationError{
					Field:  "AndRules",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *Permission_OrRules:

		if v, ok := interface{}(m.GetOrRules()).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return PermissionValidationError{
					Field:  "OrRules",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *Permission_Any:

		if m.GetAny() != true {
			return PermissionValidationError{
				Field:  "Any",
				Reason: "value must equal true",
			}
		}

	case *Permission_Header:

		if v, ok := interface{}(m.GetHeader()).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return PermissionValidationError{
					Field:  "Header",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *Permission_DestinationIp:

		if v, ok := interface{}(m.GetDestinationIp()).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return PermissionValidationError{
					Field:  "DestinationIp",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *Permission_DestinationPort:

		if m.GetDestinationPort() > 65535 {
			return PermissionValidationError{
				Field:  "DestinationPort",
				Reason: "value must be less than or equal to 65535",
			}
		}

	case *Permission_Metadata:

		if v, ok := interface{}(m.GetMetadata()).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return PermissionValidationError{
					Field:  "Metadata",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *Permission_NotRule:

		if v, ok := interface{}(m.GetNotRule()).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return PermissionValidationError{
					Field:  "NotRule",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	default:
		return PermissionValidationError{
			Field:  "Rule",
			Reason: "value is required",
		}

	}

	return nil
}

// PermissionValidationError is the validation error returned by
// Permission.Validate if the designated constraints aren't met.
type PermissionValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e PermissionValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPermission.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = PermissionValidationError{}

// Validate checks the field values on Principal with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Principal) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Identifier.(type) {

	case *Principal_AndIds:

		if v, ok := interface{}(m.GetAndIds()).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return PrincipalValidationError{
					Field:  "AndIds",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *Principal_OrIds:

		if v, ok := interface{}(m.GetOrIds()).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return PrincipalValidationError{
					Field:  "OrIds",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *Principal_Any:

		if m.GetAny() != true {
			return PrincipalValidationError{
				Field:  "Any",
				Reason: "value must equal true",
			}
		}

	case *Principal_Authenticated_:

		if v, ok := interface{}(m.GetAuthenticated()).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return PrincipalValidationError{
					Field:  "Authenticated",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *Principal_SourceIp:

		if v, ok := interface{}(m.GetSourceIp()).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return PrincipalValidationError{
					Field:  "SourceIp",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *Principal_Header:

		if v, ok := interface{}(m.GetHeader()).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return PrincipalValidationError{
					Field:  "Header",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *Principal_Metadata:

		if v, ok := interface{}(m.GetMetadata()).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return PrincipalValidationError{
					Field:  "Metadata",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *Principal_NotId:

		if v, ok := interface{}(m.GetNotId()).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return PrincipalValidationError{
					Field:  "NotId",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	default:
		return PrincipalValidationError{
			Field:  "Identifier",
			Reason: "value is required",
		}

	}

	return nil
}

// PrincipalValidationError is the validation error returned by
// Principal.Validate if the designated constraints aren't met.
type PrincipalValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e PrincipalValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPrincipal.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = PrincipalValidationError{}

// Validate checks the field values on Permission_Set with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Permission_Set) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetRules()) < 1 {
		return Permission_SetValidationError{
			Field:  "Rules",
			Reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetRules() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return Permission_SetValidationError{
					Field:  fmt.Sprintf("Rules[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// Permission_SetValidationError is the validation error returned by
// Permission_Set.Validate if the designated constraints aren't met.
type Permission_SetValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e Permission_SetValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPermission_Set.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = Permission_SetValidationError{}

// Validate checks the field values on Principal_Set with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Principal_Set) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetIds()) < 1 {
		return Principal_SetValidationError{
			Field:  "Ids",
			Reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetIds() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return Principal_SetValidationError{
					Field:  fmt.Sprintf("Ids[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// Principal_SetValidationError is the validation error returned by
// Principal_Set.Validate if the designated constraints aren't met.
type Principal_SetValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e Principal_SetValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPrincipal_Set.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = Principal_SetValidationError{}

// Validate checks the field values on Principal_Authenticated with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Principal_Authenticated) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	return nil
}

// Principal_AuthenticatedValidationError is the validation error returned by
// Principal_Authenticated.Validate if the designated constraints aren't met.
type Principal_AuthenticatedValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e Principal_AuthenticatedValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPrincipal_Authenticated.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = Principal_AuthenticatedValidationError{}
