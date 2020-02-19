package main

import "fmt"

func main() {
	const rubToDollar float64 = 63.60
	var rub float64
	fmt.Print("Сколько у Вас рублей?\n")
	fmt.Scan(&rub)
	fmt.Printf("У Вас %.2f рублей\n", rub)
	fmt.Printf("Вы получите: %.2f долларов\n", rub/rubToDollar)
}
