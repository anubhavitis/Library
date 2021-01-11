package email

import (
	"fmt"
	"os"

	mailjet "github.com/mailjet/mailjet-apiv3-go"
)

var key, pass = os.Getenv("SMTP_KEY"), os.Getenv("SMTP_PASS")

var from = &mailjet.RecipientV31{
	Email: "ss.lfsgd@gmail.com",
	Name:  "Anubhav Singhal",
}

var client = mailjet.NewMailjetClient(key, pass)

func SendWelcomeEmail(RecipientEmail string, RecipientName string)(bool){

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
			TextPart: "Welcome!! Glad to be a part of your world of books.",
			HTMLPart: "",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := client.SendMailV31(&messages)
	if err != nil {
		fmt.Println("error while sending mail",err)
		return false
	}
   fmt.Println(res.ResultsV31)
   return true
}