package protocol

type Type int

const (
	IPV4 Type = 4
	IPV6 Type = 6

	ICMPV4 Type = 1
	ICMPV6 Type = 58

	TCP Type = 6
	UDP Type = 17
)

type Protocol interface {

	PolicyType() Type

	HeadLength() uint8

	Raw() []byte

	Head() []byte

	Content() []byte
}

func ParseProtocol(packet []byte) Protocol {
	v := packet[0] >> 4

	var t uint8
	var parentProtocol Protocol
	if v == 4 {
		t = packet[9]
		parentProtocol = NewIpV4Protocol(packet)
	} else if v == 6 {
		t = packet[6]
		return nil
	}

	switch t {
	case ICMPV4:
		return NewIcmpV4Protocol(parentProtocol)
	case TCP:
		return NewTcpProtocol(parentProtocol)
	}
	return nil
}
