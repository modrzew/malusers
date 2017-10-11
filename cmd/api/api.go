package main

import (
	"fmt"

	"github.com/modrzew/malusers"
)

func main() {
	user := &malusers.User{Username: "abc"}
	fmt.Println(user)
}
