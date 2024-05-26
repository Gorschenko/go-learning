package main

import "fmt"

func main() {
	testSlice()
	testTask()
}

func testSlice() {
	var s []int
	s = []int{4, 5, 6, 7}
	// len() - покажет длину слайса,
	// cap() - вместимость - максимально возможное количество элементов,
	// которое можно добавить в массив
	fmt.Println(len(s), cap(s))

	// len = 7, cap = 7
	s2 := []string{"e", "l", "e", "m", "e", "n", "t"}
	s2[0] = "E"
	fmt.Println(s2)

	// ERROR
	// s2[10] = "E"
	// fmt.Println(s2)

	// Объявление слайса и функция make()
	s3 := make([]int, 4, 6)
	fmt.Println(s3, len(s3), cap(s3))

	s4 := s2[:4]
	s5 := s2[1:6]
	fmt.Println(s4, s5) // [E l e m][l e me n]
	s4[1] = "L"
	fmt.Println(s2, s4) // [E L e m e n t][L e m en]

	// Добавление элементов в слайс

	s6 := []int{1, 2, 3}
	s7 := []int{4, 5, 6}

	// append() добавляет элементы в конец массива,
	// возращает новый слайс
	s6 = append(s6, s7...)
	fmt.Println(s6)

	s4 = append(s4, "one", "two", "three", "four")
	// Так как вместимость слайса превышена, то Go выделил новый блок в памяти,
	// перенес туда новый слайс, после чего в исходный слайс добавил новый указатель
	fmt.Println(s2, s4, s5) // [E L e m e n t] [E L e m one two three four] [L e m e n]
	fmt.Println(cap(s4))    // cap = 14.
	// Вместимость увеличена в 2 раза,
	// чтобы не производить эту операцию в последующем

	// Копирование слайсов

	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 3, 10)
	// Функция возвращает кол-во скопированных элементов
	qnt := copy(dst, src)
	fmt.Println(qnt, dst) // 3 [1 2 3]
}

func testTask() {
	days := []int{1, 2, 3, 4, 5, 6, 7}
	workDays := make([]int, 0, 5)
	// count := copy(workDays, days)
	// fmt.Println(count, workDays)
	workDays = append(workDays, days[0:5]...)
	fmt.Println(workDays)
	days = append(days[:0], days[5:]...)
	fmt.Println((days))

	newDays := append(days, workDays...)
	fmt.Println((newDays))
}
