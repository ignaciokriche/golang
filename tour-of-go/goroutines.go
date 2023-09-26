package main

import (
	"fmt"
	"time"
)

func GoroutinesDemo() {

	fmt.Println("go routines go go")
	fmt.Println()
	go printMessage("\thello", 5, 100)
	printMessage("\tworld", 5, 200)
	fmt.Println()

	fmt.Println("channels")
	sumChannel := make(chan int)
	numbers := []int{7, 2, 8, -9, 4, 0}
	length := len(numbers)
	go sum(numbers[:length/2], sumChannel)
	go sum(numbers[length/2:], sumChannel)

	x := <-sumChannel
	y := <-sumChannel
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("sum:", x+y)

	fmt.Println("----------------------------------------------------")
	fmt.Println()

}

func sum(numbers []int, sumCh chan int) {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	sumCh <- sum
}

func printMessage(msg string, times int, delayMillisecond int) {
	for ; times > 0; times-- {
		fmt.Println(msg, times)
		time.Sleep(time.Duration(delayMillisecond) * time.Millisecond)
	}
}
