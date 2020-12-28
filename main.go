package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2/google"
)

/* Twitter Credentials
App name: bookshelf07
API key: DQMM584FmQZFzJpOw3dSM2p36
API key secret: 3rM1STdBGBd9ORHxZBCthC7AW7XjdGvIhr6q13HOBuxwR7ei0Z
Access token: 1139949042106916865-7nyRyFZFaC1ywSH8uz7vNveasDEOE5
Access token secret: eQkPSn5bCZb6eZASC71uwPhr5FnuDPp8d8rBLXSXsmAHc
Bearer token: AAAAAAAAAAAAAAAAAAAAACYXLAEAAAAAGeasPWhBLUFo03ipDjyyZZJQkS4%3Dntz5VLT5MI7up6y01lh2CJeaIQsy14BNKALI0D3DrjpbuS7lmv
*/

var (
	twitterOauthConfig = &clientcredentials.Config{
		ClientID:     "DQMM584FmQZFzJpOw3dSM2p36",
		ClientSecret: "3rM1STdBGBd9ORHxZBCthC7AW7XjdGvIhr6q13HOBuxwR7ei0Z",
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     "68309862236-0tpmf7scq3plc9ijbbfubrdh7kng9qdd.apps.googleusercontent.com",
		ClientSecret: "dLOrJJIy-xI_lczuJgAjP97G",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/drive",
		},
		Endpoint: google.Endpoint,
	}
	randomState = "random"
)

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/login2", handleLogin2)
	http.HandleFunc("/callback", handleCallback)
	http.HandleFunc("/callbackTweet", handleCallback2)
	http.ListenAndServe(":8080", nil)

}

func handleHome(w http.ResponseWriter, r *http.Request) {
	var html = `<html> <body> <div> <a href="/login"> Google Login </a></div> <a href="/login2"> Twitter Login </a> </body> </html>`
	fmt.Fprint(w, html)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(randomState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("state") != randomState {
		fmt.Println("State is not valid")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {
		fmt.Println("State is not valid")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Println("Could not get request")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Could not parse response")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Fprintf(w, "Response: %s", content)
}

func handleLogin2(w http.ResponseWriter, r *http.Request) {
	httpClient := twitterOauthConfig.Client(oauth2.NoContext)
	client := twitter.NewClient(httpClient)

}

func handleCallback2(w http.ResponseWriter, r *http.Request) {

}
