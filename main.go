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

func getFirstLetter(input string) string {
	return string(input[0])
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

	y := input[0]

	if (y >= 'a' && y <= 'z') || (y >= 'A' && y <= 'Z') {
		for {
			person, _ := faker.
				GetPerson().
				FirstName(reflect.Value{})
			newPerson := fmt.Sprintf("%+v", person)
			firstLetter := getFirstLetter(newPerson)
			if firstLetter == strings.ToUpper(string(y)) {
				fmt.Println(person)
				break
			}
		}
	}
}
