package main

import "fmt"

func main() {
	ints := [10]int{}
	for i := range ints {
		ints[i] = i
	}
	fmt.Println(ints)

	slice := []int{1, 2, 3}
	slice = append(slice, 4)
	fmt.Println(slice)
	fmt.Println(slice)
	slice1 := make([]int, 10)
	fmt.Println(slice1[10])



}
