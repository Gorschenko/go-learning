package main

import "fmt"

type SomeStruct struct{}

func foo() interface{} {
	var result *SomeStruct // = nil, здесь может быть ссылка на любой тип
	return result
}
func main() {
	res := foo()
	if res != nil {
		fmt.Println("res != nil, res =", res)
	}
}
