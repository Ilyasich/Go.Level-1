package main

import "fmt"

func main() {

	var a int

	fmt.Println("Введите трехзначное  число: ")
	fmt.Scan(&a)

	fmt.Println("Сотни: ", (a / 100))

	fmt.Println("Десятки: ", (a / 10))

	fmt.Println("Еденицы: ", (a / 1))

}
