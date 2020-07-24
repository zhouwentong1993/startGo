package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4))

}

// 多参数函数
func sum(nums ...int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}
