// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              21.06.0-9~ga41932662
// source: /usr/share/vpp/api/plugins/lacp.api.json

// Package lacp contains generated bindings for API file lacp.api.
//
// Contents:
//   2 messages
//
package lacp

import (
	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	ethernet_types "github.com/edwarnicke/govpp/binapi/ethernet_types"
	interface_types "github.com/edwarnicke/govpp/binapi/interface_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "lacp"
	APIVersion = "2.0.0"
	VersionCrc = 0xe1609dab
)

// SwInterfaceLacpDetails defines message 'sw_interface_lacp_details'.
type SwInterfaceLacpDetails struct {
	SwIfIndex             interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	InterfaceName         string                         `binapi:"string[64],name=interface_name" json:"interface_name,omitempty"`
	RxState               uint32                         `binapi:"u32,name=rx_state" json:"rx_state,omitempty"`
	TxState               uint32                         `binapi:"u32,name=tx_state" json:"tx_state,omitempty"`
	MuxState              uint32                         `binapi:"u32,name=mux_state" json:"mux_state,omitempty"`
	PtxState              uint32                         `binapi:"u32,name=ptx_state" json:"ptx_state,omitempty"`
	BondInterfaceName     string                         `binapi:"string[64],name=bond_interface_name" json:"bond_interface_name,omitempty"`
	ActorSystemPriority   uint16                         `binapi:"u16,name=actor_system_priority" json:"actor_system_priority,omitempty"`
	ActorSystem           ethernet_types.MacAddress      `binapi:"mac_address,name=actor_system" json:"actor_system,omitempty"`
	ActorKey              uint16                         `binapi:"u16,name=actor_key" json:"actor_key,omitempty"`
	ActorPortPriority     uint16                         `binapi:"u16,name=actor_port_priority" json:"actor_port_priority,omitempty"`
	ActorPortNumber       uint16                         `binapi:"u16,name=actor_port_number" json:"actor_port_number,omitempty"`
	ActorState            uint8                          `binapi:"u8,name=actor_state" json:"actor_state,omitempty"`
	PartnerSystemPriority uint16                         `binapi:"u16,name=partner_system_priority" json:"partner_system_priority,omitempty"`
	PartnerSystem         ethernet_types.MacAddress      `binapi:"mac_address,name=partner_system" json:"partner_system,omitempty"`
	PartnerKey            uint16                         `binapi:"u16,name=partner_key" json:"partner_key,omitempty"`
	PartnerPortPriority   uint16                         `binapi:"u16,name=partner_port_priority" json:"partner_port_priority,omitempty"`
	PartnerPortNumber     uint16                         `binapi:"u16,name=partner_port_number" json:"partner_port_number,omitempty"`
	PartnerState          uint8                          `binapi:"u8,name=partner_state" json:"partner_state,omitempty"`
}

func (m *SwInterfaceLacpDetails) Reset()               { *m = SwInterfaceLacpDetails{} }
func (*SwInterfaceLacpDetails) GetMessageName() string { return "sw_interface_lacp_details" }
func (*SwInterfaceLacpDetails) GetCrcString() string   { return "d9a83d2f" }
func (*SwInterfaceLacpDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SwInterfaceLacpDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4     // m.SwIfIndex
	size += 64    // m.InterfaceName
	size += 4     // m.RxState
	size += 4     // m.TxState
	size += 4     // m.MuxState
	size += 4     // m.PtxState
	size += 64    // m.BondInterfaceName
	size += 2     // m.ActorSystemPriority
	size += 1 * 6 // m.ActorSystem
	size += 2     // m.ActorKey
	size += 2     // m.ActorPortPriority
	size += 2     // m.ActorPortNumber
	size += 1     // m.ActorState
	size += 2     // m.PartnerSystemPriority
	size += 1 * 6 // m.PartnerSystem
	size += 2     // m.PartnerKey
	size += 2     // m.PartnerPortPriority
	size += 2     // m.PartnerPortNumber
	size += 1     // m.PartnerState
	return size
}
func (m *SwInterfaceLacpDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeString(m.InterfaceName, 64)
	buf.EncodeUint32(m.RxState)
	buf.EncodeUint32(m.TxState)
	buf.EncodeUint32(m.MuxState)
	buf.EncodeUint32(m.PtxState)
	buf.EncodeString(m.BondInterfaceName, 64)
	buf.EncodeUint16(m.ActorSystemPriority)
	buf.EncodeBytes(m.ActorSystem[:], 6)
	buf.EncodeUint16(m.ActorKey)
	buf.EncodeUint16(m.ActorPortPriority)
	buf.EncodeUint16(m.ActorPortNumber)
	buf.EncodeUint8(m.ActorState)
	buf.EncodeUint16(m.PartnerSystemPriority)
	buf.EncodeBytes(m.PartnerSystem[:], 6)
	buf.EncodeUint16(m.PartnerKey)
	buf.EncodeUint16(m.PartnerPortPriority)
	buf.EncodeUint16(m.PartnerPortNumber)
	buf.EncodeUint8(m.PartnerState)
	return buf.Bytes(), nil
}
func (m *SwInterfaceLacpDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.InterfaceName = buf.DecodeString(64)
	m.RxState = buf.DecodeUint32()
	m.TxState = buf.DecodeUint32()
	m.MuxState = buf.DecodeUint32()
	m.PtxState = buf.DecodeUint32()
	m.BondInterfaceName = buf.DecodeString(64)
	m.ActorSystemPriority = buf.DecodeUint16()
	copy(m.ActorSystem[:], buf.DecodeBytes(6))
	m.ActorKey = buf.DecodeUint16()
	m.ActorPortPriority = buf.DecodeUint16()
	m.ActorPortNumber = buf.DecodeUint16()
	m.ActorState = buf.DecodeUint8()
	m.PartnerSystemPriority = buf.DecodeUint16()
	copy(m.PartnerSystem[:], buf.DecodeBytes(6))
	m.PartnerKey = buf.DecodeUint16()
	m.PartnerPortPriority = buf.DecodeUint16()
	m.PartnerPortNumber = buf.DecodeUint16()
	m.PartnerState = buf.DecodeUint8()
	return nil
}

// SwInterfaceLacpDump defines message 'sw_interface_lacp_dump'.
type SwInterfaceLacpDump struct{}

func (m *SwInterfaceLacpDump) Reset()               { *m = SwInterfaceLacpDump{} }
func (*SwInterfaceLacpDump) GetMessageName() string { return "sw_interface_lacp_dump" }
func (*SwInterfaceLacpDump) GetCrcString() string   { return "51077d14" }
func (*SwInterfaceLacpDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SwInterfaceLacpDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *SwInterfaceLacpDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *SwInterfaceLacpDump) Unmarshal(b []byte) error {
	return nil
}

func init() { file_lacp_binapi_init() }
func file_lacp_binapi_init() {
	api.RegisterMessage((*SwInterfaceLacpDetails)(nil), "sw_interface_lacp_details_d9a83d2f")
	api.RegisterMessage((*SwInterfaceLacpDump)(nil), "sw_interface_lacp_dump_51077d14")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*SwInterfaceLacpDetails)(nil),
		(*SwInterfaceLacpDump)(nil),
	}
}
