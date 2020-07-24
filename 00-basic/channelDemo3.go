package main

import (
	"fmt"
)

func main() {
	t := new([]int)
	fmt.Println(cap((*t)))
	*t = append(*t, 10)
	fmt.Println(len((*t)))
	t1 := []int{1, 2, 3, 4}
	t1[10] = 1

	fmt.Println(t1 == nil)
	fmt.Println((*t) == nil)
	(*t)[0] = 10
	fmt.Println(*t)

	ints := make([]int, 10)
	ints[1] = 1
	ints[10] =10
	fmt.Println(ints)

}

func player(name string, court chan int) {

}
