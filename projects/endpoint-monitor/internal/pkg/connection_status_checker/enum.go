package connection_status_checker

type StatusChecker int

const (
	CLI  StatusChecker = iota + 1 // EnumIndex = 1
	HTTP                          // EnumIndex = 2
)

func (d StatusChecker) String() string {
	return [...]string{"CLI", "HTTP"}[d-1]
}

func (d StatusChecker) EnumIndex() int {
	return int(d)
}
