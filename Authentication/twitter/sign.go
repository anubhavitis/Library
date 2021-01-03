package twitter

import "os"

var key = os.Getenv("Twitter_Key")
var secret = os.Getenv("Twitter_Secret")

var Config := &oauth2.Config{
	ConsumerKey:    key,
	ConsumerSecret: secret,
	CallbackURL:    "http://localhost:8080/twitter/callback",
	Endpoint:       twitterOAuth1.AuthorizeEndpoint,
}

func LoginHandler(w http.ResponseWriter, r *http.Request){
	ctx:= r.Context();
	requestToken, _, err := RequestTokenFromContext(ctx)
	if err != nil {
		fmt.Println("err1",err)
	}
	authorizationURL, err := config.AuthorizationURL(requestToken)
	if err != nil {
		fmt.Println("err2",err)

	}
	http.Redirect(w, req, authorizationURL.String(), http.StatusFound)

}
// profileHandler shows a personal profile or a login button.
func profileHandler(w http.ResponseWriter, req *http.Request) {
	session, err := sessionStore.Get(req, sessionName)
	if err != nil {
		page, _ := ioutil.ReadFile("home.html")
		fmt.Fprintf(w, string(page))
		return
	}

	// authenticated profile
	fmt.Fprintf(w, `<p>You are logged in %s!</p><form action="/logout" method="post"><input type="submit" value="Logout"></form>`, session.Values[sessionUsername])
}

// logoutHandler destroys the session on POSTs and redirects to home.
func logoutHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		sessionStore.Destroy(w, sessionName)
	}
	http.Redirect(w, req, "/", http.StatusFound)
}