package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Введите значекние катета а")
	var a float64 = 3
	fmt.Println("Введите значение катета b")
	var b float64 = 4
	c := math.Sqrt(a*a + b*b)
	fmt.Printf("Длина гипотенузы %.2f\n", c)
	per := a + b + c
	fmt.Printf("Периметр прямоугольного треугольника %.2f\n", per)
	pl := ((a * b) / 2)
	fmt.Printf("Площадь прямоугольного треугольника %.2f", pl)
}
