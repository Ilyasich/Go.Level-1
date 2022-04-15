package main

import (
	"fmt"
	"math"
)

func main() {

	var s float64

	fmt.Println("Введите площадь круга: ")
	fmt.Scanln(&s)

	d := math.Sqrt(s/math.Pi) * 2

	fmt.Printf("Диаметр круга равен: %.2f\n", d)

	l := math.Pi * d

	fmt.Printf("Длина круга равна: %.2f\n", l)

}
