package main

import (
	"fmt"
	"math"
)

func main() {
	testLinks()
	testFirstTask()
	testSecondTask()
}

func testLinks() {
	var x int
	// Получить адрес переменной, то есть указатель
	// Значение по умолчанию для указателя является nil
	link := &x

	fmt.Println(link) // 0xc00000a098

	x = 64

	// Для получения значения по адресу нужно разыменовать указатель
	fmt.Println(*link) // 64

	// Можно присвоить новое значение, используя указатель
	*link = 31
	fmt.Println(*link) // 31
	fmt.Println(x)     // 31

	// Для удобства работы с указателями есть new()

	var value int
	link2 := &value
	fmt.Println(link2) // 0xc00000a098

	link3 := new(int)
	fmt.Println(link3) // 0xc00000a099
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

	radiusLink := &radius

	circleArea := math.Round(*radiusLink**radiusLink*math.Pi*100) / 100
	fmt.Println(circleArea)
}
