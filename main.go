package main

import (
	"fmt"
)

func main() {
	var slices []string
	integers := [1]int{1}

	slices = []string{"first", "2º", "other"}

	length := len(slices)

	second := slices[1:2]

	slices = append(slices, "3")
	fmt.Printf("%+v slices\n", slices)
	slices = slices[:len(slices)-1]

	fmt.Println("Slices")
	for index, value := range slices {
		fmt.Printf("index[%+v], -> value %+v\n", index, value)
	}
	fmt.Println("Final Slices")

	fmt.Printf("%+v length\n", length)
	fmt.Printf("%+v integers\n", integers)
	fmt.Printf("%+v second\n", second)

}
