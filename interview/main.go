package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5} // l=5, c=5
	b := a[2:4]               // l=2, c=3, based on a, b=[3,4],5
	c := append(b, 10)        // l=3, c=3, based on b, c=[3,4,10], b=[3,4],10 , a=[1,2,3,4,10]
	c[1] = 55                 // c=[3,55,10], b=[3,55],10, a=[1,2,3,55,10]
	//
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("c =", c)
	//
	d := b[:3]
	fmt.Println("d =", d) // d=[3,55,10]
}
