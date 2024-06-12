package alert

import (
	"encoding/gob"
	"errors"
	"fmt"
	"os"
	"time"

	"observa/internal/enums"
	log "observa/internal/pkg/log"

	mymail "observa/internal/pkg/mail"
	config "observa/projects/endpoint-monitor/config"
	data "observa/projects/endpoint-monitor/internal/data"
)

type AlertStamp struct {
	Count   int
	NextRun time.Time
}

var alertControl map[string]AlertStamp = make(map[string]AlertStamp)

const (
	baseWaitTime    = 2 * time.Minute
	maxWaitTime     = 32 * time.Minute
	alertThresholds = 5
)

func TriggerAlert(endpoint data.Endpoint, err error) {
	// Load saved state
	if err := loadState(); err != nil {
		fmt.Println("Error loading state:", err)
	}

	var key string = endpoint.Url
	if endpoint.Protocol == enums.TCP {
		key = fmt.Sprintf("%s:%d", endpoint.Ip, endpoint.Port)
	}
	if shouldAlert(key) {
		fmt.Printf("Alert triggered %d time(s) for %s at %s\n", alertControl[key].Count, key, time.Now())
		sendAlertEmail(endpoint, err)
	} else {
		fmt.Printf("No alert triggered for %s at %s (wait until %s)\n", key, time.Now(), alertControl[key].NextRun)
	}

	if err := saveState(); err != nil {
		fmt.Println("Error saving state:", err)
	}
}

func shouldAlert(key string) bool {
	stamp, ok := alertControl[key]
	if !ok {
		stamp = AlertStamp{}
	}

	if time.Now().After(stamp.NextRun) {
		stamp.Count = 1
		stamp.NextRun = time.Now().Add(baseWaitTime)
	} else {
		stamp.Count++
		stamp.NextRun = stamp.NextRun.Add(calcWaitTime(stamp.Count))
	}

	alertControl[key] = stamp

	return stamp.Count <= alertThresholds
}

func calcWaitTime(count int) time.Duration {
	waitTime := baseWaitTime

	for i := 0; i < count-alertThresholds; i++ {
		waitTime *= 2
		if waitTime > maxWaitTime {
			waitTime = maxWaitTime
			break
		}
	}

	return waitTime
}

func sendAlertEmail(endpoint data.Endpoint, err error) {

	body := fmt.Sprintf("<b>URL</b>: %s<br><b>Status Code</b>: %d<br><b>Endpoint Description</b>:%s<br><b>Downtime Implications</b>:%s <br><b>Error Details</b>:%s", endpoint.Url, endpoint.StatusCode, endpoint.Description, endpoint.Implication, err.Error())

	sender := config.Env.Alert.Sender

	if _, err := mymail.SendToMultiple(sender, config.Env.Alert.Recipients, fmt.Sprintf("[%s]Endpoint Downtime Alert: %s (%s)", endpoint.Severity.String(), endpoint.Name, endpoint.FamiliarName), body); err != nil {
		log.Error("Error sending alert mail(s):", err)
	}
}

const stateFile = "alert_state.gob"

func saveState() error {
	file, err := os.Create(stateFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(alertControl); err != nil {
		return err
	}

	return nil
}

func loadState() error {

	file, err := os.Open(stateFile)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(&alertControl); err != nil {
		return err
	}

	return nil
}
