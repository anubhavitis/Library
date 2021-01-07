package main

import (
	"fmt"

	database "github.com/anubhavitis/Library/databases"
	"github.com/anubhavitis/Library/pkg/auth"
)

func main() {
	var user = database.Member{
		UserName: "yashiG",
		Email:    "yashi",
		Password: "yashi",
		College:  "JSS",
	}
	token, err := auth.GenerateToken(user)
	fmt.Println(token)
	fmt.Println(err)
	return
}
