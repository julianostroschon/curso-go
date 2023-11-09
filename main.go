package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "TESJTEa"

	// some := strings.Split(str, "T")
	some := strings.TrimFunc(str, func(r rune) bool {
		return r == 'T' || r == 'E'
	})

	fmt.Printf("%+v\n", some)
}
