package main

import "fmt"


func main(){

	var a int
	fmt.Println("Введите число: ")
	fmt.Scanln(&a)

	fmt.Println("Результат:", fib(a))
}
	

func fib(n int) int {

	fn := make(map[int]int)
	for i := 0; i <= n; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}

	return fn[n]
}
