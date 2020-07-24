package main

import "fmt"

func main() {
	i := 1
	value(i)
	fmt.Println(i)
	pointer(&i)
	fmt.Println(i)
	s := Student{
		name: "Jack",
		age:  10,
	}
	studentValue(s)
	fmt.Println(s)
	studentPointer(&s)
	fmt.Println(s)

	arr := make([]int, 10)
	arrValue(arr)
	fmt.Println(arr)
	arrPointer(&arr)
	fmt.Println(arr)

}

type Student struct {
	name string
	age  int
}

// 传递过来的是指针，就是指向 i 这个元素的地址
// 该函数通过 *p 取出值并赋值为 0
func pointer(p *int) {
	*p = 0
}

// 传递过来的是 i 的副本，改变副本不会影响原有的值
func value(i int) {
	i = 0
}

func studentPointer(s *Student) {
	(*s).name = "Pointer"
	(*s).age = 12
}

func studentValue(s Student) {
	s.name = "value"
	s.age = 100
}

func arrPointer(arr *[]int) {
	*arr = append(*arr, 10)
}

func arrValue(arr []int) {
	arr = append(arr, 111)
}
