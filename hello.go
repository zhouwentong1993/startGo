package main

import (
	"fmt"
	"startGo/collection"
)

func main() {
	var s collection.Stack
	s.Push("World!")
	s.Push("Hello, ")
	for s.Size() > 0 {
		fmt.Print(s.Pop())
	}
	fmt.Println()

	// 函数是一等公民
	type binOp func(int, int) int
	var op binOp
	add := func(i, j int) int { return i + j }
	op = add
	n := op(10, 20)
	println(n)

	i, i2 := multiple(10, 20)
	println(i, i2)

	threeComponentLoop()

	var x MyStruct
	var px *MyStruct // px 是指针，现在是 nil
	x.s = "hello"
	px = new(MyStruct) // 现在 px 指向了新的 Struct 的地址
	px.n = 100
	println(&x)
	println(&px)

}

// Go 可以返回多个值，元组
func multiple(i, j int) (int, int) {
	return j, i
}

func conditional(int2 int) bool {
	// if 可以包含一个初始化语句，感觉没啥用
	if int2 = 10; int2 > 9 {
		println("eh")
	}
	// if 里面不需要加括号
	if int2 > 10 {
		return true
	} else {
		return false
	}
}

func threeComponentLoop() {
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	println(sum)
}

func forLikeWhile() {
	i := 1
	for i < 10 {
		i++
	}
	println(i)
}

func infiniteLoop() {
	i := 1
	for {
		i++
	}
	println(i)
}

func foreachLoop() {
	strings := []string{"a", "b", "c"}
	for i, s := range strings {
		println(i)
		println(s)
	}
}

func switchStatement() {
	n := 10
	switch n {
	case 0:
	case 10:
		println("hello")
	}
}

func switchStatement2() {
	n := 10
	switch {
	case n > 100:
	case n < 19:
		println("haha")
	}
}

func deferStatement() {

}

type MyStruct struct {
	s string
	n int64
}

func pointer() {
	p := new(MyStruct)
	p.n = 100
}
