package main

import (
	"fmt"
	"math/rand"
	"time"
)

var channels = [3]chan int{
	nil,
	make(chan int),
	nil,
}

var numbers = []int{1, 2, 3}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i]
}

func getChan(i int) chan int {
	fmt.Printf("channels[%d]\n", i)
	return channels[i]
}
func main() {
	//intChannel := make(chan int, 3)
	//intChannel <- 1
	//intChannel <- 2
	//intChannel <- 3
	//intChannel <- 4
	//first := <- intChannel
	//fmt.Println(first)
	//fmt.Println(<- intChannel)
	//fmt.Println(<- intChannel)
	//fmt.Println(<- intChannel)

	//arr := []int{7, 2, 8, -9, 4, 0}
	//channel := make(chan int)
	//go sum(arr[:len(arr)/2], channel)
	//go sum(arr[len(arr)/2:], channel)
	//// 一直在阻塞，生产者消费者模式
	//x, y := <-channel, <-channel
	//fmt.Println(x)
	//fmt.Println(y)
	//fmt.Println(x + y)

	a := make(chan int)
	//a <- 1
	a <- 1


	// 准备好几个通道。
	intChannels := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}
	// 随机选择一个通道，并向它发送元素值。
	index := rand.Intn(3)
	fmt.Printf("The index: %d\n", index)
	intChannels[index] <- index
	x := <-intChannels[index]
	fmt.Println(x)
	// 哪一个通道中有可取的元素值，哪个对应的分支就会被执行。
	fmt.Println(len(intChannels[index]))
	intChannels[0] <- 1
	select {
	case intChannels[0] <- 1:
		fmt.Println("The first candidate case is selected.")
		fmt.Println(len(intChannels[0]))

	case <-intChannels[1]:
		fmt.Println("The second candidate case is selected.")
		fmt.Println(len(intChannels[1]))

	case elem := <-intChannels[2]:
		fmt.Printf("The third candidate case is selected, the element is %d.\n", elem)
		fmt.Println(len(intChannels[2]))
	default:
		fmt.Println("No candidate case is selected!")
	}

	//select {
	//case getChan(0) <- getNumber(0):
	//	fmt.Println("The first candidate case is selected.")
	//case getChan(1) <- getNumber(1):
	//	fmt.Println("The second candidate case is selected.")
	//case getChan(2) <- getNumber(2):
	//	fmt.Println("The third candidate case is selected")
	//default:
	//	fmt.Println("No candidate case is selected!")
	//}
}

func sum(s []int, c chan int) {
	time.Sleep(2 * time.Second)
	sum := 0
	for _, value := range s {
		sum = sum + value
	}
	c <- sum
}
