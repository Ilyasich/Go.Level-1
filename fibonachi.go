package main

import (
	"fmt"
)

func main() {

	var a uint
	fmt.Println("Введите число: ")
	fmt.Scanln(&a)

	fmt.Println("Результат:", fib(a))
}

func fib(n uint) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}







