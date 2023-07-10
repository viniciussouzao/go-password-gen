package main

import (
	"fmt"
	"log"

	"github.com/viniciussouzao/go-password-gen/password"
)

func main() {
	//Generate a simple password with 8 characters, without numbers and symbols, but with uppercase letters
	pass, err := password.Generate(8, false, false, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Password generated without numbers and symbols, but with uppercase letters:", pass)

	//Generate a password with 16 characters, with numbers and symbols, but without uppercase letters
	pass, err = password.Generate(16, true, true, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Password generated with numbers and symbols, but without uppercase letters:", pass)

	//Generate a password with 32 characters, with numbers, symbols and uppercase letters
	pass, err = password.Generate(32, true, true, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Password generated with numbers, symbols and uppercase letters:", pass)

	//Generate a password with 48 characters, without numbers, symbols and uppercase letters
	pass, err = password.Generate(48, false, false, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Password generated without numbers, symbols and uppercase letters:", pass)

}
