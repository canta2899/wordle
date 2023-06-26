package main

import (
	"bufio"
	"fmt"
	"os"
  "strings"
	"github.com/gookit/color"
)

func main() {
    reader := bufio.NewScanner(os.Stdin)
    fmt.Printf("Welcome to Wordle, you can start guessing!\n\n")

    word := strings.ToUpper(GetWord())
    attempts := []string{}
    attemptsCount := 0

    for attemptsCount < 5 {
        chosenWord := getChosenWord(reader)
        fmt.Printf("\n")
        res := compare(word, chosenWord, &attempts)
        printAttempts(attempts)
        fmt.Printf("\n\n")
        if res { break }
        attemptsCount++
    }

    fmt.Printf("\n\nThe word was %v\n\n", word)
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
            word += color.BgGreen.Sprintf(" %v ", getCharAt(chosenWord, i))
        }else if in(string(getCharAt(chosenWord, i)), trueWord) && getCharAt(chosenWord, i) != getCharAt(trueWord, i) {
            word += color.BgYellow.Sprintf(" %v ", getCharAt(chosenWord, i))
        }else {
            word += color.BgBlack.Sprintf(" %v ", getCharAt(chosenWord, i))
        }

        word += color.FgGray.Sprintf(" ")
    }

    *attempts = append(*attempts, word)

    if chosenWord == trueWord {
        return true
    }

    return false
}

func getCharAt(word string, index int) string {
    if index < 0 || index > len(word) {
        panic("Error getCharAt") 
    }

    return string(word[index])
}

func in(letter string, word string) bool{
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
        } else if !IsWord(chosenWord) {
            fmt.Println("Not a word")
        }else {
            break
        }
    }

    // removing trailing new line
    return chosenWord
}

