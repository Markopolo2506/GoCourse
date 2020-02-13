package main

import "fmt"

func main() {
	fmt.Println("Сумма вклада?")
	var rub float64 = 5000
	fmt.Println("Процент?")
	var proc float64 = 5
	c := (((rub * proc * 1825) / (365 * 100)) + rub)
	fmt.Printf("Сумма вклада через 5 лет %.2f", c)
}
