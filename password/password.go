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

	Symbols = "!@#$%^&*()_+{}[]|:;<>,./?"
)

var (
	ErrorLength                  = errors.New("password length must be greater than zero")
	ErrorNumbersMoreThanLength   = errors.New("the number of numbers must be less than the length of the password")
	ErrorSymbolsMoreThanLength   = errors.New("the number of symbols must be less than the length of the password")
	ErrorUpperCaseMoreThanLength = errors.New("the number of uppercase letters must be less than the length of the password")
)

// Generate a password with the specified length, with or without numbers, symbols and uppercase letters included.
// Returns the generated password and an error if the length is less than or equal to zero.
func Generate(length int, AllowNumbers, AllowSymbols, AllowUpperCase bool) (password string, err error) {
	if length <= 0 {
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

	rand.NewSource(time.Now().UnixNano())
	pass := make([]rune, length)
	for i := range pass {
		pass[i] = characters[rand.Intn(len(characters))]
	}

	return string(pass), nil
}

// Generate a custom password with a custom length, custom number of numbers, symbols and uppercase letters included.
// Returns the generated password and an error if the length is less than or equal to zero.
func GenerateCustom(length, numCount, symbolCount, upperCount int) (password string, err error) {
	if length <= 0 {
		return "", ErrorLength
	}

	if numCount > length {
		return "", ErrorNumbersMoreThanLength
	}

	if symbolCount > length {
		return "", ErrorSymbolsMoreThanLength
	}

	if upperCount > length {
		return "", ErrorUpperCaseMoreThanLength
	}

	var finalPass string

	randInt := generateRandomNumbers(numCount)
	randSymbols := generateRandomSymbols(symbolCount)
	randUpperLetters := generateRandomuUpperLetters(upperCount)

	numOfMissingLetters := length - numCount - symbolCount - upperCount

	randMissingLetters := generateRandomLowerLetters(numOfMissingLetters)

	finalPass = randInt + randSymbols + randUpperLetters + randMissingLetters

	finalPass = shuffleWord(finalPass)

	return finalPass, nil

}
