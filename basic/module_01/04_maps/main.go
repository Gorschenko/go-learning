package main

import (
	"fmt"
	"sort"
)

func main() {
	testMaps()
}

func testMaps() {
	// Объявление и инициализация
	var book map[string]int

	var author map[string]string = map[string]string{
		"name":     "Stephen",
		"lastName": "King",
	}

	reader := map[string]string{
		"name":     "John",
		"lastName": "Smith",
		"city":     "NY",
	}

	months := map[string]int{
		"Jan": 31,
		"Feb": 28,
		"Mar": 31,
	}

	fmt.Println(book, author, reader, len(months)) // 3

	items := map[int]map[string]int{
		2020: {
			"books":      10,
			"periodical": 8,
		},
		2019: {
			"books":      12,
			"periodical": 10,
		},
	}

	// Работа с элементами

	fmt.Println(items[2020])
	items[2018] = map[string]int{
		"books": 5,
	}
	fmt.Println(items)

	book2 := map[string]int{}
	// Вернет 0 - значение по умолчанию для указанного типа
	fmt.Println(book2["Book1"])
	// Для проверки наличия ключа используется второе значение
	a, ok := book2["Book1"]
	fmt.Println(a, ok) // 0 false

	var book3 map[string]int
	fmt.Println(book3 == nil)

	// var book4 map[string]int
	// ERROR
	// Если map не был инциализирован,
	// то память под новые элементы не выделяется
	// book4["Book1"] = 2

	var book5 map[string]int = map[string]int{}
	book5["Book1"] = 2
	fmt.Println(book5)

	var book6 map[string]int
	reader2 := map[string]string{}
	fmt.Println(book6 == nil, len(book6))     // true, 0
	fmt.Println(reader2 == nil, len(reader2)) // false, 0

	book7 := make(map[string]int, 100)
	book7["Book1"] = 10
	book7["Book2"] = 6

	// Перебор элементов
	book8 := map[string]int{
		"Book2": 2,
		"Book4": 4,
		"Book1": 1,
		"Book3": 3,
	}

	keys := make([]string, 0, len(book8))

	for k := range book8 {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, v := range keys {
		fmt.Println(book8[v]) // 1 2 3 4
	}

	fmt.Println(len(keys), cap(keys))
}
