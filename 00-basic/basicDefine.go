package main

import "fmt"

/**
  基础语法：涉及
	1. 变量、常量、函数、struct、interface 定义
*/
func main() {
	// 变量的定义
	var x int = 123
	var y = 123
	z := 123

	// 常量的定义
	const a = 123

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(a)
	z = 124

	// 对结构体的使用，可以指定名字
	person := Person{name: "张三", age: 18, address: "北京"}
	fmt.Println(person)

	pPerson := &person
	fmt.Println(pPerson)

	r := Rect{
		width:  10,
		height: 20,
	}

	c := Circle{radius: 15}

	s := []shape{&r, &c}

	for _, item := range s {
		fmt.Println(item.area())
		fmt.Println(item.perimeter())
	}

}

// 结构体的声明，结构体不能声明函数，函数在 Golang 中是 first class 的，通过类型确定 「this」
type Person struct {
	name    string
	age     int8
	address string
}

type Rect struct {
	width  int
	height int
}

// 求面积
//函数定义 类型（this） 方法名 参数 返回值
func (rect *Rect) area() int {
	return rect.height * rect.width
}

func (rect *Rect) perimeter() int {
	return 2 * (rect.width + rect.height)
}

// 接口定义，隐式实现。
// 如果实现了该 interface 的所有方法的函数的 this 都是同一个 struct，那么就称该 Struct 实现了该 interface
type shape interface {
	area() int
	perimeter() int
}

// 该 Circle 实现了 shape
type Circle struct {
	radius int
}

func (circle *Circle) area() int {
	return circle.radius * circle.radius
}

func (circle *Circle) perimeter() int {
	return circle.radius * circle.radius
}
