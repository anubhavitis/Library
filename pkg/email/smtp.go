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
			HTMLPart: `
			<html>
				<body>
					<center>
						<img src="https://user-images.githubusercontent.com/26124625/104884138-dd1e2a80-598b-11eb-8b8e-bb2a23db022e.jpg" style=" height:auto; width: 500px; border-radius: 25px; margin:5px;"/>
						<div style="border: 1px solid grey; border-radius: 25px; box-shadow: 2px 2px 2px; background:lightgrey; margin:10px; width: 500px; padding:10px;">
							<h1> Glad to have you onboard! </h1>
							<p> Thanks for registering with <b> Library </b> , you are just one step away from being a member.
								<br>
								For security purpose, we would like you to verify your email, Click on the button below to verify!
								<br>
								<br>
								<button	type="box" 
									style="background:blue; box-shadow:2px 2px 0; border:1px solid white; font-size: 16px; border-radius: 10px; padding:5px; text-align: center;"
									> 
										<a href="` + Body + `" style="color:white; margin:5px 15px;"> VERIFY </a>
								</button>
								<br> <br>
								Thank You!, have a nice day.
							</p>
						</div>
					</center>
				</body>
			</html>`,
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
