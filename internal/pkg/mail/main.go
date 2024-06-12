package mymail

import (
	"fmt"
	"strings"

	log "observa/internal/pkg/log"

	"github.com/go-gomail/gomail"
)

var mailClient *gomail.Dialer = gomail.NewDialer("smtp-relay.sendinblue.com", 587, "appzonegridtesting@gmail.com", "89BtFA3bWGU27PMx")

func Send(sender string, recipient string, subject string, body string) (bool, error) {
	recipients := make([]string, 1)
	recipients[0] = recipient
	return SendToMultiple(sender, recipients, subject, body)
}
func SendToMultiple(sender string, recipients []string, subject string, body string) (bool, error) {

	if len(strings.TrimSpace(sender)) == 0 {
		return false, fmt.Errorf("specify a sender")
	}
	if len(recipients) <= 0 {
		return false, fmt.Errorf("a recipient must be provided")
	}

	if len(strings.TrimSpace(subject)) <= 0 && len(strings.TrimSpace(body)) <= 0 {
		return false, fmt.Errorf("a subject or body must be provided")

	}
	m := gomail.NewMessage()
	m.SetHeader("From", sender)

	addresses := make([]string, len(recipients))
	for i, recipient := range recipients {
		addresses[i] = m.FormatAddress(recipient, "")
	}

	m.SetHeader("To", addresses...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", fmt.Sprintf(body))

	if err := mailClient.DialAndSend(m); err != nil {
		log.Warn("Error sending email:", err)
		return false, err
	}

	return true, nil
}
