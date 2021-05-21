package protocol

import (
	"encoding/binary"
	"net"
)

func NewIpV4Protocol(raw []byte) *ipv4 {
	headLength := (raw[0] & 0xF) << 2
	head := raw[:headLength]
	return &ipv4{
		protocolType: IPV4,
		raw:          raw,
		headLength:   headLength,
		head:         head,
		content:      raw[headLength:],
		id:           binary.BigEndian.Uint16(head[4:6]),
		checksum:     binary.BigEndian.Uint16(head[10:12]),
		srcIp:        net.IPv4(head[12], head[13], head[14], head[15]),
		dstIP:        net.IPv4(head[16], head[17], head[18], head[18]),
	}
}

type ipv4 struct {
	protocolType Type
	raw          []byte
	head         []byte
	content      []byte
	headLength   uint8
	id           uint16
	checksum     uint16
	srcIp        net.IP
	dstIP        net.IP
}

func (ipv4 *ipv4) PolicyType() Type {
	return ipv4.protocolType
}

func (ipv4 *ipv4) HeadLength() uint8 {
	return ipv4.headLength
}

func (ipv4 *ipv4) Raw() []byte {
	return ipv4.raw
}

func (ipv4 *ipv4) Head() []byte {
	return ipv4.head
}

func (ipv4 *ipv4) Content() []byte {
	return ipv4.content
}

func (ipv4 *ipv4) Id() uint16 {
	return ipv4.id
}

func (ipv4 *ipv4) Checksum() uint16 {
	return ipv4.checksum
}

func (ipv4 *ipv4) SrcIp() net.IP {
	return ipv4.srcIp
}

func (ipv4 *ipv4) DstIp() net.IP {
	return ipv4.dstIP
}
