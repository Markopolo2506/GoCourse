package main

import "fmt"

func main() {
	var rub float64
	var proc float64
	fmt.Print("Сумма вклада?\n")
	fmt.Scan(&rub)
	fmt.Print("Процент?\n")
	fmt.Scan(&proc)
	c := (((rub * proc * 1825) / (365 * 100)) + rub)
	fmt.Printf("Сумма вклада через 5 лет %.2f", c)
}
