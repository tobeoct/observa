package publisher

import (
	"fmt"
	"net/http"
	"strings"

	log "observa/internal/pkg/log"
	config "observa/projects/endpoint-monitor/config"
)

func Push(channel string, message string) {
	go func() {
		var url = fmt.Sprintf("%s?channelName=%s&Message=%s", config.Env.Collector.Url, channel, strings.ReplaceAll(message, " ", ","))
		resp, err := http.PostForm(url, nil)
		if err != nil {
			log.Errorf("Error pushing to Publisher API: %s %s", url, err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Warnf("Publisher API: %s returned non-200 status code: %d", url, resp.StatusCode)
		}
	}()

}
