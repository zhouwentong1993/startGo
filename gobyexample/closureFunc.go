package main

import "fmt"

func main() {
	oneStep := intStep()
	fmt.Println(oneStep())
	fmt.Println(oneStep())
	fmt.Println(oneStep())
	newStep := intStep()
	fmt.Println(newStep())
}

// 闭包，让函数返回函数
func intStep() func() int{
	i := 0
	return func() int {
		i++
		return i
	}
}
