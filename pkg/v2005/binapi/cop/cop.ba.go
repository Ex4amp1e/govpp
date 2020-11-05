// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              20.05-release
// source: /usr/share/vpp/api/core/cop.api.json

// Package cop contains generated bindings for API file cop.api.
//
// Contents:
//   4 messages
//
package cop

import (
	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	interface_types "github.com/edwarnicke/govpp/pkg/v2005/binapi/interface_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "cop"
	APIVersion = "1.0.1"
	VersionCrc = 0xfd7e767d
)

// CopInterfaceEnableDisable defines message 'cop_interface_enable_disable'.
type CopInterfaceEnableDisable struct {
	SwIfIndex     interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	EnableDisable bool                           `binapi:"bool,name=enable_disable" json:"enable_disable,omitempty"`
}

func (m *CopInterfaceEnableDisable) Reset()               { *m = CopInterfaceEnableDisable{} }
func (*CopInterfaceEnableDisable) GetMessageName() string { return "cop_interface_enable_disable" }
func (*CopInterfaceEnableDisable) GetCrcString() string   { return "5501adee" }
func (*CopInterfaceEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *CopInterfaceEnableDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 1 // m.EnableDisable
	return size
}
func (m *CopInterfaceEnableDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBool(m.EnableDisable)
	return buf.Bytes(), nil
}
func (m *CopInterfaceEnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.EnableDisable = buf.DecodeBool()
	return nil
}

// CopInterfaceEnableDisableReply defines message 'cop_interface_enable_disable_reply'.
type CopInterfaceEnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *CopInterfaceEnableDisableReply) Reset() { *m = CopInterfaceEnableDisableReply{} }
func (*CopInterfaceEnableDisableReply) GetMessageName() string {
	return "cop_interface_enable_disable_reply"
}
func (*CopInterfaceEnableDisableReply) GetCrcString() string { return "e8d4e804" }
func (*CopInterfaceEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *CopInterfaceEnableDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *CopInterfaceEnableDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *CopInterfaceEnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// CopWhitelistEnableDisable defines message 'cop_whitelist_enable_disable'.
type CopWhitelistEnableDisable struct {
	SwIfIndex  interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	FibID      uint32                         `binapi:"u32,name=fib_id" json:"fib_id,omitempty"`
	IP4        bool                           `binapi:"bool,name=ip4" json:"ip4,omitempty"`
	IP6        bool                           `binapi:"bool,name=ip6" json:"ip6,omitempty"`
	DefaultCop bool                           `binapi:"bool,name=default_cop" json:"default_cop,omitempty"`
}

func (m *CopWhitelistEnableDisable) Reset()               { *m = CopWhitelistEnableDisable{} }
func (*CopWhitelistEnableDisable) GetMessageName() string { return "cop_whitelist_enable_disable" }
func (*CopWhitelistEnableDisable) GetCrcString() string   { return "debe13ea" }
func (*CopWhitelistEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *CopWhitelistEnableDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 4 // m.FibID
	size += 1 // m.IP4
	size += 1 // m.IP6
	size += 1 // m.DefaultCop
	return size
}
func (m *CopWhitelistEnableDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint32(m.FibID)
	buf.EncodeBool(m.IP4)
	buf.EncodeBool(m.IP6)
	buf.EncodeBool(m.DefaultCop)
	return buf.Bytes(), nil
}
func (m *CopWhitelistEnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.FibID = buf.DecodeUint32()
	m.IP4 = buf.DecodeBool()
	m.IP6 = buf.DecodeBool()
	m.DefaultCop = buf.DecodeBool()
	return nil
}

// CopWhitelistEnableDisableReply defines message 'cop_whitelist_enable_disable_reply'.
type CopWhitelistEnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *CopWhitelistEnableDisableReply) Reset() { *m = CopWhitelistEnableDisableReply{} }
func (*CopWhitelistEnableDisableReply) GetMessageName() string {
	return "cop_whitelist_enable_disable_reply"
}
func (*CopWhitelistEnableDisableReply) GetCrcString() string { return "e8d4e804" }
func (*CopWhitelistEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *CopWhitelistEnableDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *CopWhitelistEnableDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *CopWhitelistEnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_cop_binapi_init() }
func file_cop_binapi_init() {
	api.RegisterMessage((*CopInterfaceEnableDisable)(nil), "cop_interface_enable_disable_5501adee")
	api.RegisterMessage((*CopInterfaceEnableDisableReply)(nil), "cop_interface_enable_disable_reply_e8d4e804")
	api.RegisterMessage((*CopWhitelistEnableDisable)(nil), "cop_whitelist_enable_disable_debe13ea")
	api.RegisterMessage((*CopWhitelistEnableDisableReply)(nil), "cop_whitelist_enable_disable_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*CopInterfaceEnableDisable)(nil),
		(*CopInterfaceEnableDisableReply)(nil),
		(*CopWhitelistEnableDisable)(nil),
		(*CopWhitelistEnableDisableReply)(nil),
	}
}