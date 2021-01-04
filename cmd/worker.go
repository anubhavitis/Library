package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", profileHandler)
	mux.HandleFunc("/logout", logoutHandler)
	// 1. Register Twitter login and callback handlers

	mux.Handle("/twitter/login", twitter.LoginHandler(oauth1Config, nil))
	mux.Handle("/twitter/callback", twitter.CallbackHandler(oauth1Config, issueSession(), nil))
	log.Printf("Starting Server listening on %s\n", address)
	err := http.ListenAndServe(address, New(config))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
