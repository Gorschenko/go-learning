package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"unicode/utf8"
)

func main() {
	testLocalVariables()
	testEmptyVariable()
	testSumOfDifferentTypes()
	testString()
	testArray()
	testDefineTypes()
	task()
}

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

func testEmptyVariable() {
	var path = "/path/to/file"
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

func testSumOfDifferentTypes() {
	const (
		untypedNum     = 15
		typedNum   int = 10
	)

	var num int32 = 30

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
// iota - выглядит как автоинкремент, начинается с 0
const (
	_       = 10 * iota
	speed10 = 10 * iota
	speed20 = 10 * iota
	speed30 = 10 * iota
	speed40 = 10 * iota
)

// ТИПЫ ДАННЫХ

// ЧИСЛА
var num int = 100

// int8, int16, int32, int64
var num8 int8 = 1<<7 - 1

var unsignedNum uint16 = 1<<16 - 1

// Числа с плавающей точкой
var fl float64 = 10.64

// БУЛЕВЫЙ

var boolean bool
var tr = true

// СТРОКИ
// Для работы со строками имеется два типа - byte и rune,
// которые представлены типа unit8 и uint32
var newStr string = "New string"
var hello = "Hello" + "some string"

func testString() {
	var greeting = "Привет!"
	// Функция len() возвращает длину строки в байтах
	fmt.Println(len(greeting))                    // 13 байт
	fmt.Println(utf8.RuneCountInString(greeting)) // 7 символов (рун)

	// Оператор : (slice - срез) позволяет получить подстроку

	fmt.Println(greeting[3])    // 128
	fmt.Println(greeting[4:10]) // и

	// Можно опустить одно или оба значения, тогда используются значения 0 и len(greeting)
	fmt.Println(greeting[:6]) // При
	fmt.Println(greeting[:])  // Привет!

	// Изменить строку нельзя
	// ERROR
	// greeting[0] = "Л"

	// Нужно ее конвертировать в слайс рун, потом изменить и конвертировать в строку:
	var convGreeting = []rune(greeting)
	fmt.Println((convGreeting))
	convGreeting[4] = 'Е'
	fmt.Println(string(convGreeting)) // ПривЕт!
}

// МАССИВЫ

var defArr [3]string = [3]string{"one", "two", "three"}
var array [3]int // [0, 0, 0]

func testArray() {
	// Массив с заранее неизвестным количеством элементов
	dynamic := [...]bool{4: true} // [false, false, false, false, true]
	// Будет новое значение
	arr := dynamic
	fmt.Println(arr, len(arr))

	arr[2] = true
	fmt.Println(arr)     // [false, false, true, false, true]
	fmt.Println(dynamic) // [false, false, false, false, true]

	// Используем ссылку на массив
	arr2 := &dynamic
	arr2[2] = true // [false, false, true, false, true]
	fmt.Println((arr2))

	// Slice ссылается на элементы исходного массива
	sl := dynamic[:2] // [false, true]
	fmt.Println(sl)
	sl[1] = true
	fmt.Println(dynamic) // [false, true, true, false, true]
	fmt.Println(sl)      // [false, true]

	// Если ссылаться на элемент, которого нет, то будет ошибка компиляции
	// ERROR
	// fmt.Println(dynamic[7])
}

// ОПРЕДЕЛЕНИЕ ТИПОВ

type SettlementId string

func testDefineTypes() {
	var cityId SettlementId = "123-456-789"
	var strCityId string = "123-456-789"
	fmt.Println(cityId == SettlementId(strCityId))
}

// ЗАДАЧА

func task() {
	str := "104"
	number := 35

	n, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Ошибка преобразования строки в число")
	}
	fmt.Println("Преобразование строки в число", n)

	s := strconv.Itoa(number)
	fmt.Println(s)

	type AmericanVelocity float64
	type EuropeanVelocity float64

	var europeanVelocity EuropeanVelocity = 120.4 / 1000 * 3600
	var americanVelocity AmericanVelocity = AmericanVelocity(math.Round(130/1.609/1000*3600*100) / 100)
	fmt.Println(europeanVelocity, americanVelocity)
}
