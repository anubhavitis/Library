package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dghubble/oauth1"
	twauth "github.com/dghubble/oauth1/twitter"
)

var consumerKey = os.Getenv("Twitter_Key")
var consumerSecret = os.Getenv("Twitter_Secret")
var requestSecret, requestToken string
var err error

var config = oauth1.Config{
	ConsumerKey:    consumerKey,
	ConsumerSecret: consumerSecret,
	CallbackURL:    "http://localhost:8000/twitter/callback",
	Endpoint:       twauth.AuthorizeEndpoint,
}

func TwitterLoginHandler(w http.ResponseWriter, r *http.Request) {
	requestToken, requestSecret, err = config.RequestToken()
	if err != nil {
		fmt.Println("error while generating request token", err)
	}
	authorizationURL, errr := config.AuthorizationURL(requestToken)
	fmt.Println(authorizationURL.String())
	if errr != nil {
		fmt.Println("err2", errr)
	}
	http.Redirect(w, r, authorizationURL.String(), 302)
}

func TwitterCallbackHandler(w http.ResponseWriter, r *http.Request) {
	checkToken, verifier, err := oauth1.ParseAuthorizationCallback(r)
	if err != nil {
		fmt.Println("error while verifying:", err)
		return
	}

	accessToken, accessSecret, err := config.AccessToken(checkToken, "checkSecret", verifier)
	if err != nil {
		fmt.Println("Access token not generated:", err)
		return
	}
	fmt.Println(accessSecret, accessToken)
	w.Write([]byte("okay"))
}
