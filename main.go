package main

import (
    "fmt"
	"github.com/fatih/color"
)

func main() {
    green :=  color.New(color.FgHiGreen).SprintFunc()
    word := GetWord()
    fmt.Println("Welcome to worldle, your word is", green(word))
}
