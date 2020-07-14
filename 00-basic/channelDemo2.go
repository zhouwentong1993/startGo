package main

import (
	"fmt"
)

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}
	for _, ch := range chs {
		<-ch
	}

	ch := make(chan int, 1)
	for {
		select {
		// 这里相当于对 ch 进行操作了，这合适吗？
		case ch <- 0:
		case ch <- 1:
		}
		i := <-ch
		fmt.Println("Value received:", i)
	}

	// 会发生死锁
	//channel := make(chan string)
	//func() {
	//	channel <- "hello"
	//}()
	//fmt.Println(<-channel)
	//channel := make(chan string, 5)
	//go say("goroutine", channel)
	//say("main", channel)
	//for s := range channel {
	//	fmt.Println("channel:" + s)
	//}
}

func Count(ch chan int) {
	ch <- 1
	fmt.Println("Counting")
}

func say(word string, channel chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println("hello " + word)
		channel <- "hello" + word
	}
}
