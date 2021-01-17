package email

import (
	"fmt"

	mailjet "github.com/mailjet/mailjet-apiv3-go"
)

// var key, pass = os.Getenv("SMTP_KEY"), os.Getenv("SMTP_PASS")

var key = "4552e05a89258be962e3ae25c446a00c"
var pass = "e479991038b74f2fb43feaf8ea78fd4b"

var from = &mailjet.RecipientV31{
	Email: "ss.lfsgd@gmail.com",
	Name:  "Anubhav Singhal",
}

var client = mailjet.NewMailjetClient(key, pass)

//SendWelcomeEmail func
func SendWelcomeEmail(RecipientEmail string, RecipientName string, Body string) bool {
	messagesInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: from,
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: RecipientEmail,
					Name:  RecipientName,
				},
			},
			Subject:  "Glad to have you onboard!!",
			TextPart: "Welcome!! Glad to be a part of your world of books. " + Body,
			HTMLPart: `<br><html><body><center><a href="` + Body + `"> Click here </a></center></body></html>`,
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := client.SendMailV31(&messages)
	if err != nil {
		fmt.Println("error while sending mail", err)
		return false
	}

	fmt.Println(res.ResultsV31)
	return true
}
