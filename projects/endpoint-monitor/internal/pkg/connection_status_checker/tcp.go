package connection_status_checker

import (
	"fmt"
	"net"
	"time"

	log "observa/internal/pkg/log"
)

func CheckTCPState(ip string, port int) (bool, int, error) {
	timeout := 5
	connectionState := checkPort(ip, port, timeout)
	switch connectionState {
	case ConnectionStateListening:
		log.Infof("Port %s:%d is actively listening and has an established connection\n", ip, port)
		return true, 200, nil
	case ConnectionStateNotListening:
		log.Errorf("Port %s:%d is not actively listening\n", ip, port)
		return false, 503, fmt.Errorf("port (%s:%d) is not actively listening", ip, port)
	case ConnectionStateError:
		log.Warnf("Error checking port %s:%d\n", ip, port)
		return false, 401, fmt.Errorf("error checking connection %s:%d", ip, port)
	}
	return false, 0, fmt.Errorf("could not determine state of the connection")
}

// ConnectionState represents the state of a port connection
type ConnectionState int

const (
	ConnectionStateListening ConnectionState = iota
	ConnectionStateNotListening
	ConnectionStateError
)

func checkPort(ip string, port int, timeout int) ConnectionState {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), time.Duration(timeout)*time.Second)
	if err != nil {
		fmt.Println("1.Error dialing remote host:", err)
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			fmt.Println("2.Error dialing remote host:", netErr, err)
			return ConnectionStateNotListening
		}
		fmt.Println("3.Error dialing remote host:", err)
		return ConnectionStateError
	}
	defer conn.Close()
	return ConnectionStateListening
}
