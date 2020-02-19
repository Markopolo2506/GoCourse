package main

import (
	"fmt"
	"math/big"
)

func main() {
	chis := make([]*big.Int, 100, 100)
	for i := 0; i < 100; i++ {
		switch i {
		case 0:
			chis[i] = big.NewInt(0)
		case 1:
			chis[i] = big.NewInt(1)
		default:
			chis[i] = big.NewInt(0).Add(chis[i-1], chis[i-2])
		}
		fmt.Printf("%v -номер. Число Фибоначчи- %v\n", i+1, chis[i])
	}
}
