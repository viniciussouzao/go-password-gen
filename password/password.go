package password

import (
	"errors"
	"math/rand"
	"time"
)

const (
	LettersUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	LettersLower = "abcdefghijklmnopqrstuvwxyz"

	Numerals = "0123456789"

	Symbols = "!@#$%^&*()_+{}[]|:;<>,.?"
)

var (
	ErrorLength = errors.New("password length must be greater than zero")
)

func Generate(lenght int, AllowNumbers, AllowSymbols, AllowUpperCase bool) (password string, err error) {
	if lenght <= 0 {
		return "", ErrorLength
	}

	numbers := ""
	symbols := ""
	defaultLetters := LettersLower

	if AllowNumbers {
		numbers = Numerals
	}

	if AllowSymbols {
		symbols = Symbols
	}

	if AllowUpperCase {
		defaultLetters += LettersUpper
	}

	var characters = []rune(defaultLetters + numbers + symbols)

	//rand.Seed(time.Now().UnixNano())
	rand.NewSource(time.Now().UnixNano())
	pass := make([]rune, lenght)
	for i := range pass {
		pass[i] = characters[rand.Intn(len(characters))]
	}

	return string(pass), nil
}
