package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Credential struct {
	Cid     string `json:"cid,omitempty"`
	Csecret string `json:"csecret,omitempty"`
}

var cred Credential

func init() {

	f, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println("could not read the file:",err)
	}
	err=json.Unmarshal(f, &cred)
	fmt.Println(err, cred)
}

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8000/google/callback",
		ClientID:     cred.Cid,
		ClientSecret: cred.Csecret,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
	randomState = "random"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	var html = `<html> <body> <div> <a href="/login"> Google Login </a></div> </body> </html>`
	fmt.Fprint(w, html)
}

func GoogleSignupHandler(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(randomState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
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

// func main() {
// 	http.HandleFunc("/", handleHome)

// 	//Google Oauths
// 	http.HandleFunc("/login", handleLogin)
// 	http.HandleFunc("/callback", handleCallback)

// 	http.ListenAndServe(":8080", nil)
// }
