package slack

import (
	"os"
)

var (
	VAPrefundAlerts = os.Getenv("SLACK_VA_ALERT_URL")
)

// FormatMonitoring ...
func FormatMonitoring(transfertype, transaction_id, status, product string, amount string, err error) {
	s := Attachment{}
	s.Color = "#2cc421"
	s.Pretext = "Hey, here's a summary of the transaction ..."
	s.Title = ":file_folder: Please check the details below!"

	f := Field{}
	f.Title = "transaction id"
	f.Value = transaction_id
	f.Short = false
	s.Fields = append(s.Fields, f)

	f = Field{}
	f.Title = "Transfer Type"
	f.Value = transfertype
	f.Short = true
	s.Fields = append(s.Fields, f)

	f = Field{}
	f.Title = "Product code"
	f.Value = product
	f.Short = true
	s.Fields = append(s.Fields, f)

	f = Field{}
	f.Title = "status"
	f.Value = status
	f.Short = true
	s.Fields = append(s.Fields, f)

	f = Field{}
	f.Title = "amount"
	f.Value = amount
	f.Short = true
	s.Fields = append(s.Fields, f)

	if err != nil {
		s.Color = "#a31205"

		f = Field{}
		f.Title = "Message"
		f.Value = "```" + err.Error() + "```"
		f.Short = false
		s.Fields = append(s.Fields, f)
	}

	a := Attachments{}
	a.Attachment = append(a.Attachment, s)

	SendNotification(VAPrefundAlerts, a)
}
