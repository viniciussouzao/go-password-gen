package password

import (
	"math/rand"
	"time"
)

func generateRandomNumbers(count int) string {
	rand.NewSource(time.Now().UnixNano())
	result := ""
	for i := 0; i < count; i++ {
		randomIndex := rand.Intn(len(Numerals))
		result += string(Numerals[randomIndex])
	}

	return result
}

func generateRandomLowerLetters(count int) string {
	var letters = LettersLower

	rand.NewSource(time.Now().UnixNano())
	result := ""
	for i := 0; i < count; i++ {
		randomIndex := rand.Intn(len(letters))
		result += string(letters[randomIndex])
	}

	return result
}

func generateRandomuUpperLetters(count int) string {
	var letters = LettersUpper

	rand.NewSource(time.Now().UnixNano())
	result := ""
	for i := 0; i < count; i++ {
		randomIndex := rand.Intn(len(letters))
		result += string(letters[randomIndex])
	}

	return result
}

func generateRandomSymbols(count int) string {
	rand.NewSource(time.Now().UnixNano())
	result := ""
	for i := 0; i < count; i++ {
		randomIndex := rand.Intn(len(Symbols))
		result += string(Symbols[randomIndex])
	}

	return result
}

func shuffleWord(word string) string {

	runes := []rune(word)

	rand.NewSource(time.Now().UnixNano())
	rand.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})

	shuffledWord := string(runes)

	return shuffledWord
}
