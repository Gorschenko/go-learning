package main

import (
	"fmt"
	"math"
)

func main() {
	// .Println после вывода перейдет на новую строку,
	// тогда не нужно указывать \n
	fmt.Print("Калькулятор индекса массы тела \n")
	// Альтернативный вариант записи переменной
	// Позволяет указать тип после переменной
	// var userHeight: float64 = 1.83
	// Позволяет позже присвоить значение
	// var userHeight float64
	// userHeight = 1.83
	userHeight := 1.83
	var userKg float64 = 83
	userHeight = 2.0
	// Можно сразу объявить две переменных
	// var userHeight, userKg float64 = 1.8, 100
	// userHeight, userKg := 1.8, 100

	fmt.Print("Введите свой рост в метрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес в кг: ")
	fmt.Scan((&userKg))

	// Переназначить const не получится
	// У переменной untyped int тип, так как
	// не явно был задан тип. Это упрощает работу с float64
	const IMTPower = 2

	// Второй вариант работы с типом:
	// var IMT = float64(userKg) / userHeight
	var IMT = userKg / math.Pow(userHeight, IMTPower)
	fmt.Printf("Ваш индекс массы тела %v", IMT)
}

func outputResult(IMT float64) {
	result := fmt.Sprintf("\nВаш индекс массы тела: %v", IMT)
	fmt.Print(result)
}

func calculateIMT(userKg float64, userHeight float64) float64 {
	const IMTPower = 2
	IMT := userKg / math.Pow(userHeight, IMTPower)
	return IMT
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Print("Введите свой рост в метрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес в кг: ")
	fmt.Scan((&userKg))

	return userHeight, userKg
}

func getUserInputTwo() (userHeight float64, userKg float64) {
	return userHeight, 85
}
