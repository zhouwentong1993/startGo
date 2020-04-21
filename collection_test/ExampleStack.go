package collection_test

import (
	"fmt"
	"startGo/collection"
)

func main() {
	print("hello world!")
}

func ExampleStack() {
	var s collection.Stack
	s.Push("World!")
	s.Push("Hello, ")
	for s.Size() > 0 {
		fmt.Print(s.Pop())
	}
	fmt.Println()
}
