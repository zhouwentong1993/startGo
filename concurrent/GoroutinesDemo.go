package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		go foo(i)
	}
	var input string
	fmt.Scanln(&input)
}

func foo(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n,":", i)
	}
}
