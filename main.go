package main

import (
	"fmt"
)

func getUniqueValues(slices []string) []string {
	var mapi = map[string]int{}
	for key, value := range slices {
		mapi[value] = key
	}
	var result = []string{}

	for key := range mapi {
		result = append(result, (key))
	}
	return result
}

func main() {
	var slices []string

	slices = []string{"a", "a", "c", "c", "b", "b", "d", "d"}

	fmt.Printf("%+v\n", getUniqueValues(slices))
}
