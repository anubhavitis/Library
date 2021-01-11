package main

import (
	"fmt"

	"github.com/anubhavitis/Library/pkg/email")


func main() {
	ok:=email.SendWelcomeEmail("vashishtiv@gmail.com","Yashi Gupta")
	fmt.Print(ok)
}