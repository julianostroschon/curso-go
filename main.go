package main

import "fmt"

func main() {
	m := map[string]int{"a": 1}
	m["b"] = 2

	for key, value := range m {

		fmt.Printf("[%+v] -> %+v\n", key, value)
	}
}
