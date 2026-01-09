package main

import (
	"fmt"
	"strings"
)

func encryptWord(word string) string {
	if len(word) == 0 {
		return word
	}

	runes := []rune(word)

	if len(runes) <= 1 {
		return word
	}

	firstRune := runes[0]

	restRunes := runes[1:]
	reversedRest := reverseRunes(restRunes)

	return string(firstRune) + string(reversedRest)
}

func reverseRunes(runes []rune) []rune {
	reversed := make([]rune, len(runes))
	for i, j := 0, len(runes)-1; i <= j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = runes[j], runes[i]
	}
	return reversed
}

func encryptPhrase(phrase string) string {
	words := strings.Fields(phrase)

	if len(words) == 0 {
		return ""
	}

	encryptedWords := make([]string, len(words))
	for i, word := range words {
		encryptedWords[i] = encryptWord(word)
	}

	return strings.Join(encryptedWords, " ")
}

func main() {

	testPhrases := []string{
		"Pepe Schnele is a legend",
		"Привет мир как дела",
		"Шифрование работает отлично",
		"Съешь ещё этих мягких французских булок да выпей чаю",
		"Go is awesome",
		"А",
		"Яблоко груша слива",
		"",
	}

	for _, phrase := range testPhrases {
		encrypted := encryptPhrase(phrase)
		fmt.Printf("Исходная:  \"%s\"\n", phrase)
		fmt.Printf("Шифр:      \"%s\"\n", encrypted)
		fmt.Println()
		fmt.Println()
	}

}
