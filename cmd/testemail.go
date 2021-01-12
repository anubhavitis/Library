package main

import (
	"fmt"

	"github.com/anubhavitis/Library/pkg/email"
)

func main() {
	// ok:=email.SendWelcomeEmail("vashishtiv@gmail.com","Yashi Gupta")
	ok1 := email.SendWelcomeEmail("anubhavitis@gmail.com", "Anubhav Singhal")
	fmt.Print(ok1)
}
