// Code generated by go-enum DO NOT EDIT.
// Version: 0.3.9
// Revision: 5a95b1bbcaf1f8f32542725929d84acdf48e0259
// Build Date: 2021-11-05T17:00:39Z
// Built By: goreleaser

package schema

import (
	"fmt"
)

const (
	// MessageTypeNone is a MessageType of type None.
	MessageTypeNone MessageType = iota
	// MessageTypeChatMember is a MessageType of type Chat_member.
	MessageTypeChatMember
	// MessageTypeMax is a MessageType of type Max.
	MessageTypeMax
)

const _MessageTypeName = "nonechat_membermax"

var _MessageTypeMap = map[MessageType]string{
	MessageTypeNone:       _MessageTypeName[0:4],
	MessageTypeChatMember: _MessageTypeName[4:15],
	MessageTypeMax:        _MessageTypeName[15:18],
}

// String implements the Stringer interface.
func (x MessageType) String() string {
	if str, ok := _MessageTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("MessageType(%d)", x)
}

var _MessageTypeValue = map[string]MessageType{
	_MessageTypeName[0:4]:   MessageTypeNone,
	_MessageTypeName[4:15]:  MessageTypeChatMember,
	_MessageTypeName[15:18]: MessageTypeMax,
}

// ParseMessageType attempts to convert a string to a MessageType
func ParseMessageType(name string) (MessageType, error) {
	if x, ok := _MessageTypeValue[name]; ok {
		return x, nil
	}
	return MessageType(0), fmt.Errorf("%s is not a valid MessageType", name)
}