package main

import (
	"bufio"
	"fmt"
	"os"
    "strings"
	"github.com/gookit/color"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    word := strings.ToUpper(GetWord())
    fmt.Println("word is", word)

    attempts := []string{}
    attemptsCount := 0

    for attemptsCount < 5 {

        chosenWord := getChosenWord(reader)

        res := compare(word, chosenWord, &attempts)
        printAttempts(attempts)

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

        word += color.FgGray.Sprintf("|")
    }

    *attempts = append(*attempts, word)

    if chosenWord == trueWord {
        return true
    }

    return false
}

func getCharAt(word string, index int) string {
    if index < 0 || index > len(word) {
        panic("Error getCharAt for word")
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

func getChosenWord(reader *bufio.Reader) string {
    var chosenWord string

    for {
        chosenWord, _ = reader.ReadString('\n')

        if len(chosenWord) != 6 {
            fmt.Println("Word does not have five characters")
        } else {
            break
        }
    }

    // removing trailing new line
    return strings.ToUpper(chosenWord[:5])
}
