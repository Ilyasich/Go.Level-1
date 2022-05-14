package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func fibWithCache(n int, cache map[int]int) int {
	if val, ok := cache[n]; ok {
		return val
	}

	cache[n] = fibWithCache(n-1, cache) + fibWithCache(n-2, cache)
	return cache[n]
}

func main() {
	var a int
	fmt.Println("Введите число: ")
	fmt.Scanln(&a)
	fmt.Println("Результат:", fib(a))

	fmt.Println(fibWithCache(a, map[int]int{0: 0, 1: 1}))
}