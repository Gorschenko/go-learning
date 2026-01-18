package main

type ABC interface {
	A()
	B()
	C()
}
type AB interface {
	A()
	B()
}
type BC interface {
	B()
	C()
}
type abc struct{}

//
func (a abc) A() {}
func (a abc) B() {}
func (a abc) C() {}

//
type ab struct{}

//
func (a ab) A() {}
func (a ab) B() {}

//
func main() {
	var a interface{}
	a = abc{}

	ab := a.(AB)
	ab.A()
	ab.C()
}
