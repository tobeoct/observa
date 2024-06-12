package connection_status_checker

import (
	"fmt"
	"net/http"
	"os/exec"
	"strconv"

	log "observa/internal/pkg/log"
)

func CheckHTTPState(url string) (bool, int, error) {
	log.Infof("Checking if url:(%s) is up", url)
	statusCode, err := getStatus(url, HTTP)
	if statusCode == 0 {
		statusCode, err := getStatus(url, CLI)
		log.Debug("Tried cli method to check uptime: status-code:", statusCode, err)
	}
	if err != nil {
		log.Warn("Error checking URL:", err)
		return false, 0, err
	}
	log.Infof("Done checking result for url: (%s), status code:%d", url, statusCode)

	if statusCode < 200 || statusCode >= 400 {
		return false, statusCode, fmt.Errorf("endpoint:%s is currently not up, status code:%d", url, statusCode)
	}
	return true, statusCode, nil
}
func getStatus(url string, checker StatusChecker) (int, error) {
	if checker == CLI {
		return getStatusUsingCLI(url)
	}
	return getStatusUsingHttp(url)
}
func getStatusUsingCLI(url string) (int, error) {
	cmd := exec.Command("curl", "-s", "-o", "NUL", "-w", "%{http_code}", url)

	log.Debug("Command:", cmd.String())
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Warn("Error running command:", err.Error())
		return 0, err
	}
	statusCode, err := strconv.Atoi(string(output))

	if err != nil {
		log.Warn("Error converting status code to integer:", err.Error())
		return 0, err
	}

	return statusCode, nil
}
func getStatusUsingHttp(url string) (int, error) {

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
