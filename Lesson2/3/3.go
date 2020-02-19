package main

import "fmt"

func main() {
	var chislo int
	fmt.Print("Введите число\n")
	fmt.Scan(&chislo)
	if chislo%3 == 0 {
		fmt.Println(chislo, "число делится без остатка на 3.")
	} else {
		fmt.Println(chislo, "число не делится без остатка на 3.")
	}
}
