package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/bxcodec/faker"
)

func init() {
	fmt.Println("Input a letter...")
}

func getFirstLetter(word string) string {
	return string(word[0])
}

func main() {
	var input string
	var err error

	fmt.Scanln(&input)
	for len(input) > 1 {
		err = errors.New("Error!\nEnter only a letter")
		fmt.Println(err)
		fmt.Scanln(&input)
	}

	firstLetterInput := strings.ToUpper(getFirstLetter(input))

	if firstLetterInput >= "A" && firstLetterInput <= "Z" {
		for {
			person, _ := faker.
				GetPerson().
				FirstName(reflect.Value{})

			firstLetterInName := getFirstLetter(fmt.Sprintf("%+v", person))

			if firstLetterInName == firstLetterInput {
				fmt.Println(person)
				break
			}
		}
	}
}
