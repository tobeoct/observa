package enums

import "strings"

type Protocols int

const (
	HTTP Protocols = iota
	TCP
	UDP
)

var protocolsStrings = [...]string{"HTTP", "TCP", "UDP"}

func (d Protocols) String() string {
	return protocolsStrings[d]
}

func (d Protocols) EnumIndex() int {
	return int(d)
}

func (d Protocols) ToEnum(s string) Protocols {
	for i, str := range protocolsStrings {
		if strings.EqualFold(s, str) {
			return Protocols(i)
		}
	}
	return -1
}
