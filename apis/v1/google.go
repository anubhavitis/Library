package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/anubhavitis/Library/apis/utility"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type credential struct {
	Cid      string `json:"client_id"`
	Csecret  string `json:"client_secret"`
	Redirect string `json:"redirect_uris"`
}

var cred credential

func init() {

	f, err := ioutil.ReadFile("apis/v1/googleauth.json")
	if err != nil {
		fmt.Println("could not read the file:", err)
	}
	err = json.Unmarshal(f, &cred)
	// fmt.Println(err, cred)
}

var (
	googleOauthConfig = &oauth2.Config{

		RedirectURL:  "http://localhost:8000/google_callback",
		ClientID:     "68309862236-0tpmf7scq3plc9ijbbfubrdh7kng9qdd.apps.googleusercontent.com",
		ClientSecret: "dLOrJJIy-xI_lczuJgAjP97G",
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

//GoogleSignupHandler func
func GoogleSignupHandler(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(randomState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

//GoogleCallbackHandler func
func GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	var res utility.Result

	if r.FormValue("state") != randomState {
		res.Error = "State is not valid"
		utility.SendResponse(w, http.StatusConflict, res)
		return
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {
		res.Error = fmt.Sprintf("%s", err)
		utility.SendResponse(w, http.StatusConflict, res)
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		res.Error = fmt.Sprintf("%s", err)
		utility.SendResponse(w, http.StatusConflict, res)
		return
	}

	defer resp.Body.Close()

	var user struct {
		ID      string `json:"id"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Fname   string `json:"given_name"`
		Sname   string `json:"family_name"`
		Picture string `json:"picture"`
	}

	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		res.Error = fmt.Sprintf("%s", err)
		utility.SendResponse(w, 400, res)
	}

	res.Success = true
	res.Body = map[string]interface{}{
		"content": user,
	}
	utility.SendResponse(w, 202, res)
}

func main() {

	//Google Oauths
	http.HandleFunc("/login", GoogleSignupHandler)
	http.HandleFunc("/callback", GoogleCallbackHandler)

}
