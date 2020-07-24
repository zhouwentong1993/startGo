package main

import "fmt"

func main() {
	r := rect{Width: 10, Height: 20}
	fmt.Println(r.area())
	fmt.Println(r.perim())

	p := Person{
		age:    10,
		height: 110,
	}
	p.growUpPointer()
	fmt.Println(p)
}

type rect struct {
	Width  int
	Height int
}

type Person struct {
	age    int
	height int
}

func (p Person) growUp() {
	p.age = p.age + 1
	p.height = p.height + 10
}

func (p *Person) growUpPointer() {
	(*p).age += 1
	p.height += 10
}

func (r *rect) area() int {
	return r.Width * r.Height
}

func (r rect) perim() int {
	return 2*r.Width + 2*r.Height
}
