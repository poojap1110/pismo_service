package finance

import (
	"fmt"
	"strings"

	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/config"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/mailer"
)

// SendMismatchNotification function - sends the mismatch notification to the finance team.
func SendMismatchNotification(c *container.Container, csvFiles []string, args string) error {
	if len(csvFiles) == 0 {
		return nil
	}

	financeMail := config.GetInstance(c).GetTechnicalConfigs(constant.TechnicalFinanceMail)
	recipients := strings.Split(financeMail.(map[string]string)["recipients"], ",")
	ccRecipients := financeMail.(map[string]string)["ccRecipients"]
	bccRecipients := financeMail.(map[string]string)["bccRecipients"]
	subject := financeMail.(map[string]string)["subject"]
	htmlBody := financeMail.(map[string]string)["htmlBody"]

	i := mailer.New(c)
	i = (*i).AddRecipient(recipients...)
	i = (*i).SetSubject(subject)
	i = (*i).SetHTMLBody(fmt.Sprintf(htmlBody, args))

	if ccRecipients != "" {
		ccR := strings.Split(ccRecipients, ",")

		for _, v := range ccR {
			i = (*i).AddCC(v, v)
		}
	}

	if bccRecipients != "" {
		bccR := strings.Split(bccRecipients, ",")

		for _, v := range bccR {
			i = (*i).AddBCC(v, v)
		}
	}

	for _, v := range csvFiles {
		i = (*i).AttachFile(v)
	}

	return (*i).Send()
}

// SendFundSettlementReport - sends email with the fund statement
func SendFundSettlementReport(c *container.Container, csvFiles []string, args string) error {
	if len(csvFiles) == 0 {
		return nil
	}

	financeMail := config.GetInstance(c).GetTechnicalConfigs(constant.TechnicalFinanceMail)
	recipients := strings.Split(financeMail.(map[string]string)["recipients"], ",")
	ccRecipients := financeMail.(map[string]string)["ccRecipients"]
	bccRecipients := financeMail.(map[string]string)["bccRecipients"]
	subject := "Platform SGPrefund Settlement Notification"
	htmlBody := financeMail.(map[string]string)["htmlBody"]

	i := mailer.New(c)
	i = (*i).AddRecipient(recipients...)
	i = (*i).SetSubject(subject)
	i = (*i).SetHTMLBody(fmt.Sprintf(htmlBody, args))

	if ccRecipients != "" {
		ccR := strings.Split(ccRecipients, ",")

		for _, v := range ccR {
			i = (*i).AddCC(v, v)
		}
	}

	if bccRecipients != "" {
		bccR := strings.Split(bccRecipients, ",")

		for _, v := range bccR {
			i = (*i).AddBCC(v, v)
		}
	}

	for _, v := range csvFiles {
		i = (*i).AttachFile(v)
	}

	return (*i).Send()
}
