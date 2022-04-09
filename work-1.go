package main

import "fmt"

func main() {

	var a, b float32

	fmt.Println("Введите высоту: ")
	fmt.Scanln(&a)

	fmt.Println("Введите ширину: ")
	fmt.Scanln(&b)

	fmt.Println("Площадь прямоугольника равна: ", a*b)

}
