package main

import "fmt"

func main() {
	s := make([]int, 0, 5) // l=0, c=5
	s = append(s, 1, 2, 3) // l=3, c=5, s=[1,2,3],0,0
	//
	subSlice := s[1:3] // l=2, c=4, based on s, subSlice=[2,3],0,0
	//
	subSlice[0] = 99               // subSlice=[99,3],0,0, s=[1,99,3],0,0
	subSlice = append(subSlice, 4) // l=3, c=4, based on s, subSlice=[99,3,4],0, s=[1,99,3],4,0
	//
	s = append(s, 5, 6, 7) // l=6, c=10, ,new array, s=[1,99,3,5,6,7],0,0,0,0
	//
	fmt.Println("s = ", s)
	fmt.Println("subSlice = ", subSlice)
}
