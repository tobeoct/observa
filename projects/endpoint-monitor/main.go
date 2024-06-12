package main

import (
	"fmt"

	log "observa/internal/pkg/log"

	publisher "observa/internal/pkg/publisher"
	utils "observa/internal/pkg/utils"
	config "observa/projects/endpoint-monitor/config"
	data "observa/projects/endpoint-monitor/internal/data"
	alert "observa/projects/endpoint-monitor/internal/pkg/alert"
	connection_status_checker "observa/projects/endpoint-monitor/internal/pkg/connection_status_checker"
)

const confDir = ""

func main() {
	config.LoadConfig(confDir, "env")
	log.Init()
	// Schedule the check every minute
	utils.ScheduleAndStart("@every 5m", checkURLs)
	data.LoadEndpoints()
	defer log.Close()
	select {}
}

func checkURLs() {
	for _, endpoint := range data.Endpoints {
		go monitor(endpoint)
	}
}

func monitor(endpoint data.Endpoint) {

	_, statusCode, err := connection_status_checker.IsUp(endpoint)
	endpoint.StatusCode = statusCode
	publisher.Push("status_updates", fmt.Sprintf("%s,%d", endpoint.Url, statusCode))
	if err != nil {
		alert.TriggerAlert(endpoint, err)
		return
	}

}
