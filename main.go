package main

import (
	"bufio"
	"fmt"
	"github.com/gookit/color"
	"os"
	"strings"
)

var green = color.S256(233, 34)
var yellow = color.S256(233, 220)
var gray = color.S256(255, 241)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Printf("Welcome to Wordle, you can start guessing!\n\n")

	word := strings.ToUpper(GetWord())
	attempts := []string{}
	attemptsCount := 0

	for attemptsCount < 5 {
		chosenWord := getChosenWord(reader)
		fmt.Printf("\n")
		match := compare(word, chosenWord, &attempts)
		printAttempts(attempts)
		fmt.Printf("\n\n")
		if match {
			fmt.Println("Congrats!")
			break
		}
		attemptsCount++
	}

	fmt.Printf("The word was %v\n\n", word)
}

func printAttempts(attempts []string) {
	for _, word := range attempts {
		fmt.Println(word)
	}
}

func compare(trueWord string, chosenWord string, attempts *[]string) bool {

	word := " "

	for i := 0; i < 5; i++ {
		if getCharAt(chosenWord, i) == getCharAt(trueWord, i) {
			word += green.Sprintf(" %v ", getCharAt(chosenWord, i))
		} else if in(string(getCharAt(chosenWord, i)), trueWord) && getCharAt(chosenWord, i) != getCharAt(trueWord, i) {
			word += yellow.Sprintf(" %v ", getCharAt(chosenWord, i))
		} else {
			word += gray.Sprintf(" %v ", getCharAt(chosenWord, i))
		}

		word += fmt.Sprintf(" ")
	}

	*attempts = append(*attempts, word)

	return chosenWord == trueWord
}

func getCharAt(word string, index int) string {
	if index < 0 || index > len(word) {
		panic("unable to parse word")
	}

	return strings.ToUpper(string(word[index]))
}

func in(letter string, word string) bool {
	for _, l := range word {
		if letter == string(l) {
			return true
		}
	}
	return false
}

func getChosenWord(scanner *bufio.Scanner) string {
	var chosenWord string

	for {
		fmt.Printf("Guess your word: ")
		scanner.Scan()
		chosenWord = scanner.Text()

		if len(chosenWord) != 5 {
			fmt.Println("Word does not have five characters")
		} else {
			break
		}
	}

	return chosenWord
}
