package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64
	fmt.Println("Введите значекние катета а")
	fmt.Scan(&a)
	fmt.Println("Введите значение катета b")
	fmt.Scan(&b)
	c := math.Sqrt(a*a + b*b)
	fmt.Printf("Длина гипотенузы %.2f\n", c)
	per := a + b + c
	fmt.Printf("Периметр прямоугольного треугольника %.2f\n", per)
	pl := (a * b) / 2
	fmt.Printf("Площадь прямоугольного треугольника %.2f", pl)
}
