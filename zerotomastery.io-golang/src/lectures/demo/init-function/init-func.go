package main

import (
	"fmt"
	"regexp"
)

// Cash compiled regex to variable to avoid compilation every time is used
var EmailExpr *regexp.Regexp

func init() {
	compiled, ok := regexp.Compile(`.+@.+\..+`)

	if ok != nil {
		panic("failed to compile regular expression")
	}

	EmailExpr = compiled
	fmt.Println("regular expression compiled successfully")
}

func isValidEmail(addr string) bool {
	return EmailExpr.Match([]byte(addr))
}

func main() {
	fmt.Println(isValidEmail("invalid"))
	fmt.Println(isValidEmail("valid@example.com"))
	fmt.Println(isValidEmail("invalid@example"))
}