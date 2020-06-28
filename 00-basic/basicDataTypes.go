package main

import "fmt"

/**
介绍基本的数据类型
*/
func main() {
	// 基本数据类型与 Java 的基本相同，不同之处在于 Go 将 string 类型也纳入了基础类型
	var a bool = true
	var b int = 10
	var c uint8 = 20
	var d float32 = 123.1
	var f float64 = 12311.123
	var g string = "hello world"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(f)
	fmt.Println(g)
	// 派生类型就不介绍了，比如数组类型，map 类型，Struct 等
}
