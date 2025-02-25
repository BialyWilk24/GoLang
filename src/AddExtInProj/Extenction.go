package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {

	fmt.Println(quote.Go())

}

// func printString(str string) {
// 	fmt.Printf("%s\n", str)
// }

// func reverseString(str string) string {
// 	runes := []rune(str)
// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}
// 	return string(runes)
// }
