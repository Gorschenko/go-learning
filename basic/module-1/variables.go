package main

import (
	"fmt"
	"os"
)

// ПЕРЕМЕННЫЕ

var myVar string = "Variable string"

// Значение будет по умолчанию
var anotherVar string

// Самоопредление типа
var noType = 100

// Объявление нескольких переменных
var a, b, c int
var d, e, f = "hello", 42, true

// Блок инциализации переменных
var (
	price       int
	qty         int
	isDeletable bool
)

// КРАТКОЕ ОБЪЯВЛЕНИЕ

// Для объявление локальных переменных используется синтаксис краткого объявления
func testLocalVariables() {
	pathToFile := "somePath"
	fmt.Println(pathToFile)
	str, number, isExist := "new string", 42, false
	fmt.Println(str, number, isExist)

	// Для кратого объявление должно быть хотя бы одно новое значение
	// ERROR
	// str, number := "new string", 42
}

var path = "/path/to/file"

func testEmptyVar() {
	f, err := os.Open(path)
	fmt.Println(f, err)
	// Если функция возвращает два значения, но одно из них не будет использоваться,
	// то применяется символ _
	str, _ := os.Open(path)
	fmt.Println(str)
}

// КОНСТАНТЫ

const statusCode int = 200
const (
	orderStatusNew string = "new"
	baseDiscount          = 3.5
)

const (
	untypedNum     = 15
	typedNum   int = 10
)

var num int32 = 30

func main() {
	fmt.Println(num + untypedNum)

	// Ошибка из-за несоответствия типов
	// ERROR
	// fmt.Println(num + typedNum)
}

// Второе значение будет равно 21, четвертое - 10
const (
	num1 = 21
	num2
	num3 = 10
	num4
)

// ГЕНЕРАТОР КОНСТАНТ iota
