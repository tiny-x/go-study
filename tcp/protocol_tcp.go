package protocol

import (
	"encoding/binary"
	"fmt"
)

const TCP_HEAD_LEN = 20

func NewTcpProtocol(parent Protocol) *tcp {
	raw := parent.Content()[parent.HeadLength():]
	headLength := (raw[12] >> 4) * 4
	head := raw[:headLength]

	return &tcp{
		protocolType: TCP,
		headLength:   headLength,
		raw:          raw,
		head:         head,
		content:      raw[headLength:],
		srcPort:      binary.BigEndian.Uint16(head[0:2]),
		dstPort:      binary.BigEndian.Uint16(head[2:4]),
	}
}

type tcp struct {
	protocolType Type
	raw          []byte
	head         []byte
	content      []byte
	headLength   uint8
	srcPort      uint16
	dstPort      uint16
	seq          uint8
	ack          uint8
	reserved     uint8
	isNs         bool
	isCwr        bool
	isEce        bool
	isUrg        bool
	isAck        bool
	isPsh        bool
	isRst        bool
	isSyn        bool
	isFin        bool
	window       uint8
	checksum     uint8
	urgptr       uint8
	options      uint8
}

func (tcp *tcp) PolicyType() Type {
	return ICMPV4
}

func (tcp *tcp) HeadLength() byte {
	return 20
}

func (tcp *tcp) Raw() []byte {
	return tcp.raw
}

func (tcp *tcp) Head() []byte {
	return tcp.head
}

func (tcp *tcp) Content() []byte {
	return tcp.content
}

func (tcp *tcp) String() string {
	if tcp == nil {
		return "<nil>"
	}

	srcPort := tcp.SrcPort()
	dstPort := tcp.DstPort()

	return fmt.Sprintf("{\n"+
		"\t\tProtocol=TCP\n"+
		"\t\tSrcPort=%d\n"+
		"\t\tDstPort=%d\n"+
		"\t\tSeqNum=%#x\n"+
		"\t\tAckNum=%d\n"+
		"\t\tHeaderLen=%d\n"+
		"\t\tReserved=%d\n"+
		"\t\tFlags={NS=%t CWR=%t ECE=%t URG=%t ACK=%t PSH=%t RST=%t SYN=%t FIN=%t}\n"+
		"\t\tWindow=%d\n"+
		"\t\tCheckSum=%#x\n"+
		"\t\tUrgPtr=%d\n"+
		"\t\tOptions=%v\n"+
		"\t}\n",
		srcPort, dstPort, tcp.Seq(), tcp.Ack(), tcp.HeaderLen(), tcp.Reserved(), tcp.NS(), tcp.CWR(), tcp.ECE(), tcp.URG(), tcp.ACK(), tcp.PSH(), tcp.RST(), tcp.SYN(), tcp.FIN(), tcp.Window(), tcp.Checksum(), tcp.UrgPtr(), tcp.Options())
}

func (tcp *tcp) SrcPort() uint16 {
	return tcp.srcPort
}

func (tcp *tcp) DstPort() uint16 {
	return tcp.dstPort
}

func (tcp *tcp) Seq() uint32 {
	return binary.BigEndian.Uint32(tcp.content[4:8])
}

func (tcp *tcp) Ack() uint32 {
	return binary.BigEndian.Uint32(tcp.content[8:12])
}

func (tcp *tcp) HeaderLen() int {
	return int(tcp.DataOffset()) * 4
}

func (tcp *tcp) DataOffset() uint8 {
	return tcp.content[12] >> 4
}

func (tcp *tcp) Reserved() uint8 {
	return (tcp.content[12] >> 1) & 0x7
}

func (tcp *tcp) NS() bool {
	return tcp.content[12]&0x1 == 1
}

func (tcp *tcp) CWR() bool {
	return tcp.content[13]>>7 == 1
}

func (tcp *tcp) ECE() bool {
	return (tcp.content[13]>>6)&0x1 == 1
}

func (tcp *tcp) URG() bool {
	return (tcp.content[13]>>5)&0x1 == 1
}

func (tcp *tcp) ACK() bool {
	return (tcp.content[13]>>4)&0x1 == 1
}

func (tcp *tcp) PSH() bool {
	return (tcp.content[13]>>3)&0x1 == 1
}

func (tcp *tcp) RST() bool {
	return (tcp.content[13]>>2)&0x1 == 1

}

func (tcp *tcp) SYN() bool {
	return (tcp.content[13]>>1)&0x1 == 1
}

func (tcp *tcp) FIN() bool {
	return tcp.content[13]&0x1 == 1
}

func (tcp *tcp) Window() uint16 {
	return binary.BigEndian.Uint16(tcp.content[14:16])
}

func (tcp *tcp) Checksum() uint16 {
	return binary.BigEndian.Uint16(tcp.content[16:18])
}

func (tcp *tcp) UrgPtr() uint16 {
	return binary.BigEndian.Uint16(tcp.content[18:20])
}

func (tcp *tcp) Options() []byte {
	hdrLen := tcp.HeaderLen()
	if hdrLen <= 20 {
		return nil
	}
	return tcp.content[20:hdrLen]
}
