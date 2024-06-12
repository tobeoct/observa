package enums

import (
	"strings"
)

type SeverityLevel int

const (
	CRITICAL SeverityLevel = iota
	URGENT
	WARNING
	INFORMATION
	LOW
)

var severityStrings = [...]string{"CRITICAL", "URGENT", "WARNING", "INFORMATION", "LOW"}

func (d SeverityLevel) String() string {
	return severityStrings[d]
}

func (d SeverityLevel) EnumIndex() int {
	return int(d)
}

func (d SeverityLevel) ToEnum(s string) SeverityLevel {
	for i, str := range severityStrings {
		if strings.EqualFold(s, str) {
			return SeverityLevel(i)
		}
	}
	return -1
}
