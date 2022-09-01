package main

import "fmt"

func main() {

	s := "stando"

	sRune := s[0]
	testRune := 's'

	fmt.Printf("This is the value of the letter %v: %v\n", s, s)
	fmt.Printf("This is the value of the first letter of %v: %v\n", s, sRune)
	fmt.Printf("This is the value of s: %v\n", testRune)
}
