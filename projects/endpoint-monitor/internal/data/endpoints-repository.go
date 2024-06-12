package data

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"time"

	"observa/internal/enums"
)

type Endpoint struct {
	Url          string
	Name         string
	Description  string
	Implication  string
	Severity     enums.SeverityLevel
	FamiliarName string
	StatusCode   int
	Protocol     enums.Protocols
	Ip           string
	Port         int
}

type endpointDTO struct {
	Url          string `json:"Url"`
	Name         string `json:"Name"`
	Description  string `json:"Description"`
	Implication  string `json:"Implication"`
	Severity     string `json:"Severity"`
	FamiliarName string `json:"FamiliarName"`
	StatusCode   int    `json:"StatusCode"`
	Protocol     string `json:"Protocol"`
	Ip           string `json:"Ip"`
	Port         int    `json:"Port"`
}

var Endpoints []Endpoint

//go:embed endpoints.json
var endpointsData []byte

func LoadEndpoints() {

	endpoints, err := loadEndpointsFromJSON()
	if err != nil {
		fmt.Println("Error loading endpoints:", err)
		return
	}
	Endpoints = endpoints

	// Periodic cache refresh
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if err := refreshCache(); err != nil {
					fmt.Println("Error refreshing cache:", err)
				}
			}
		}
	}()

	// Keep the main goroutine alive
	select {}
}

func loadEndpointsFromJSON() ([]Endpoint, error) {
	var endpoints []Endpoint = make([]Endpoint, 0)
	var endpointDtos []endpointDTO
	if err := json.Unmarshal(endpointsData, &endpointDtos); err != nil {
		return nil, err
	}

	for _, ep := range endpointDtos {
		endpoints = append(endpoints, Endpoint{
			Name:         ep.Name,
			FamiliarName: ep.FamiliarName,
			Url:          ep.Url,
			Description:  ep.Description,
			Implication:  ep.Implication,
			Severity:     enums.LOW.ToEnum(ep.Severity),
			Protocol:     enums.HTTP.ToEnum(ep.Protocol),
			Ip:           ep.Ip,
			Port:         ep.Port,
		})
	}
	return endpoints, nil
}

func refreshCache() error {

	endpoints, err := loadEndpointsFromJSON()
	if err != nil {
		return err
	}
	Endpoints = endpoints

	return nil
}
