package main

import "fmt"

func main() {
	const rubToDollar float64 = 63.60
	fmt.Println("Сколько у Вас рублей?")
	var rub float64 = 1000
	fmt.Printf("У Вас %.2f рублей\n", rub)
	fmt.Printf("Вы получите: %.2f долларов\n", rub/rubToDollar)
}
