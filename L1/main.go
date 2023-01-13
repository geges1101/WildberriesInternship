package main

import "fmt"

func read(a chan int, b chan int, nums []int) {
	for _, x := range nums {
		a <- x
		b <- x * 2
	}
}

func main() {
	a := make(chan int)
	b := make(chan int)
	length := 0
	fmt.Println("Enter the number of inputs")
	fmt.Scanln(&length)
	fmt.Println("Enter the inputs")
	numbers := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scanln(&numbers[i])
	}

	go read(a, b, numbers)

	res := <-b
	fmt.Printf(string(rune(res)))
}
