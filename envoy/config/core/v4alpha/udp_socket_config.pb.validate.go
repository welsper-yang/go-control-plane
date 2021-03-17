// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/core/v4alpha/udp_socket_config.proto

package envoy_config_core_v4alpha

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on UdpSocketConfig with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *UdpSocketConfig) Validate() error {
	if m == nil {
		return nil
	}

	if wrapper := m.GetMaxRxDatagramSize(); wrapper != nil {

		if val := wrapper.GetValue(); val <= 0 || val >= 65536 {
			return UdpSocketConfigValidationError{
				field:  "MaxRxDatagramSize",
				reason: "value must be inside range (0, 65536)",
			}
		}

	}

	return nil
}

// UdpSocketConfigValidationError is the validation error returned by
// UdpSocketConfig.Validate if the designated constraints aren't met.
type UdpSocketConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UdpSocketConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UdpSocketConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UdpSocketConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UdpSocketConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UdpSocketConfigValidationError) ErrorName() string { return "UdpSocketConfigValidationError" }

// Error satisfies the builtin error interface
func (e UdpSocketConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUdpSocketConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UdpSocketConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UdpSocketConfigValidationError{}
