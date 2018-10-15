// Code generated by GoVPP binapi-generator. DO NOT EDIT.
// source: api/stn.api.json

/*
Package stn is a generated VPP binary API of the 'stn' VPP module.

It is generated from this file:
	stn.api.json

It contains these VPP binary API objects:
	4 messages
	2 services
*/
package stn

import "git.fd.io/govpp.git/api"
import "github.com/lunixbochs/struc"
import "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

/* Messages */

// StnAddDelRule represents the VPP binary API message 'stn_add_del_rule'.
// Generated from 'stn.api.json', line 4:
//
//            "stn_add_del_rule",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u8",
//                "is_ip4"
//            ],
//            [
//                "u8",
//                "ip_address",
//                16
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            [
//                "u8",
//                "is_add"
//            ],
//            {
//                "crc": "0x9f0bbe21"
//            }
//
type StnAddDelRule struct {
	IsIP4     uint8
	IPAddress []byte `struc:"[16]byte"`
	SwIfIndex uint32
	IsAdd     uint8
}

func (*StnAddDelRule) GetMessageName() string {
	return "stn_add_del_rule"
}
func (*StnAddDelRule) GetCrcString() string {
	return "9f0bbe21"
}
func (*StnAddDelRule) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func NewStnAddDelRule() api.Message {
	return &StnAddDelRule{}
}

// StnAddDelRuleReply represents the VPP binary API message 'stn_add_del_rule_reply'.
// Generated from 'stn.api.json', line 39:
//
//            "stn_add_del_rule_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type StnAddDelRuleReply struct {
	Retval int32
}

func (*StnAddDelRuleReply) GetMessageName() string {
	return "stn_add_del_rule_reply"
}
func (*StnAddDelRuleReply) GetCrcString() string {
	return "e8d4e804"
}
func (*StnAddDelRuleReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func NewStnAddDelRuleReply() api.Message {
	return &StnAddDelRuleReply{}
}

// StnRulesDump represents the VPP binary API message 'stn_rules_dump'.
// Generated from 'stn.api.json', line 57:
//
//            "stn_rules_dump",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            {
//                "crc": "0x51077d14"
//            }
//
type StnRulesDump struct{}

func (*StnRulesDump) GetMessageName() string {
	return "stn_rules_dump"
}
func (*StnRulesDump) GetCrcString() string {
	return "51077d14"
}
func (*StnRulesDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func NewStnRulesDump() api.Message {
	return &StnRulesDump{}
}

// StnRulesDetails represents the VPP binary API message 'stn_rules_details'.
// Generated from 'stn.api.json', line 75:
//
//            "stn_rules_details",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u8",
//                "is_ip4"
//            ],
//            [
//                "u8",
//                "ip_address",
//                16
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            {
//                "crc": "0x5eafa31e"
//            }
//
type StnRulesDetails struct {
	IsIP4     uint8
	IPAddress []byte `struc:"[16]byte"`
	SwIfIndex uint32
}

func (*StnRulesDetails) GetMessageName() string {
	return "stn_rules_details"
}
func (*StnRulesDetails) GetCrcString() string {
	return "5eafa31e"
}
func (*StnRulesDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func NewStnRulesDetails() api.Message {
	return &StnRulesDetails{}
}

/* Services */

type Services interface {
	DumpStnRules(*StnRulesDump) (*StnRulesDetails, error)
	StnAddDelRule(*StnAddDelRule) (*StnAddDelRuleReply, error)
}

func init() {
	api.RegisterMessage((*StnAddDelRule)(nil), "stn.StnAddDelRule")
	api.RegisterMessage((*StnAddDelRuleReply)(nil), "stn.StnAddDelRuleReply")
	api.RegisterMessage((*StnRulesDump)(nil), "stn.StnRulesDump")
	api.RegisterMessage((*StnRulesDetails)(nil), "stn.StnRulesDetails")
}