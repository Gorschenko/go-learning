package main

import (
	"fmt"
	"math"
)

func main() {
	testPointers()
	testFirstTask()
	testSecondTask()
}

func testPointers() {
	var x int
	// Получить адрес переменной, то есть указатель
	// Значение по умолчанию для указателя является nil
	pointer := &x

	fmt.Println(pointer) // 0xc00000a098

	x = 64

	// Для получения значения по адресу нужно разыменовать указатель
	fmt.Println(*pointer) // 64

	// Можно присвоить новое значение, используя указатель
	*pointer = 31
	fmt.Println(*pointer) // 31
	fmt.Println(x)        // 31

	// Для удобства работы с указателями есть new()

	var value int
	pointer2 := &value
	fmt.Println(pointer2) // 0xc00000a098

	pointer3 := new(int)
	fmt.Println(pointer3) // 0xc00000a099
}

func testFirstTask() {
	fmt.Println("First task is started")
	a := new(int)
	b := 100

	*a = b
	fmt.Println(*a)

	b = *a + 10
	fmt.Println(b)
}

func testSecondTask() {
	const circleLength int = 35

	var radius float64 = math.Round(float64(circleLength)/2/math.Pi*100) / 100
	fmt.Println(radius)

	radiusPointer := &radius

	circleArea := math.Round(*radiusPointer**radiusPointer*math.Pi*100) / 100
	fmt.Println(circleArea)
}
