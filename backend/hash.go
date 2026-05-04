package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	h, _ := bcrypt.GenerateFromPassword([]byte("c@ndyBAR123"), 12)
	fmt.Println(string(h))
}
