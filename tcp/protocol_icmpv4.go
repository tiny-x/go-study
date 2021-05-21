package protocol

import (
	"encoding/binary"
)

const ICMPV4_HEAD_LEN = 8

func NewIcmpV4Protocol(parent Protocol) *icmpv4 {
	raw := parent.Content()[parent.HeadLength():]
	head := raw[:ICMPV4_HEAD_LEN]

	return &icmpv4{
		protocolType: ICMPV4,
		parent:       parent,
		raw:          raw,
		head:         head,
		content:      raw[ICMPV4_HEAD_LEN:],
		type_:        head[0],
		code:         head[1],
		checksum:     binary.BigEndian.Uint16(head[2:4]),
		body:         binary.BigEndian.Uint32(head[4:8]),
	}
}

type icmpv4 struct {
	protocolType Type
	parent       Protocol
	raw          []byte
	head         []byte
	content      []byte
	headLength   uint8
	type_        uint8
	code         uint8
	checksum     uint16
	body         uint32
}

func (icmpv4 *icmpv4) PolicyType() Type {
	return icmpv4.protocolType
}

func (icmpv4 *icmpv4) HeadLength() byte {
	return ICMPV4_HEAD_LEN
}

func (icmpv4 *icmpv4) Raw() []byte {
	return icmpv4.raw
}

func (icmpv4 *icmpv4) Head() []byte {
	return icmpv4.head
}

func (icmpv4 *icmpv4) Content() []byte {
	return icmpv4.content
}

func (icmpv4 *icmpv4) Type() byte {
	return icmpv4.type_
}

func (icmpv4 *icmpv4) Code() byte {
	return icmpv4.code
}

func (icmpv4 *icmpv4) Checksum() uint16 {
	return icmpv4.checksum
}

func (icmpv4 *icmpv4) Body() uint32 {
	return icmpv4.body
}
