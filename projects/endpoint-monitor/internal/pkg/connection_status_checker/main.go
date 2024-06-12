package connection_status_checker

import (
	"fmt"

	enums "observa/internal/enums"
	data "observa/projects/endpoint-monitor/internal/data"
)

func IsUp(endpoint data.Endpoint) (bool, int, error) {
	if endpoint.Protocol == enums.HTTP {
		return CheckHTTPState(endpoint.Url)
	}
	if endpoint.Protocol == enums.TCP {
		return CheckTCPState(endpoint.Ip, endpoint.Port)
	}
	return false, 0, fmt.Errorf("protocol not supported for endpoint monitoring")
}
