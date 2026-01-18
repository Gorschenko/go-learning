package main

import "fmt"

func main() {
	greetings := "привет как дела"
	fmt.Println(greetings)
	runes := []rune(greetings)
	runes[1] = 's'
	result := string(runes)
	fmt.Println(result)
}
