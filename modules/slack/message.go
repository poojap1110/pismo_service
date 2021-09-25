package slack

import (
	logs "bitbucket.org/matchmove/logs"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"
)

// Attachments ...
type Attachments struct {
	Attachment []Attachment `json:"attachments"`
}

// Attachment ...
type Attachment struct {
	Color   string  `json:"color"`
	Pretext string  `json:"pretext"`
	Title   string  `json:"title"`
	Fields  []Field `json:"fields"`
}

// Field ...
type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

// SendNotification ...
func SendNotification(webhookURL string, a Attachments) error {
	log := logs.New()
	slackBody, _ := json.Marshal(a)
	req, err := http.NewRequest(http.MethodPost, webhookURL, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + os.Getenv("SLACK_LAMBDA_MONITORING_TOKEN")
	req.Header.Add("Authorization", bearer)

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Print("Slack Notification Error :", err.Error())
		log.Dump()
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		log.Print("Slack Notification Response :", buf.String())
		log.Dump()
		return errors.New("Non-ok response returned from Slack")
	}

	return nil
}
