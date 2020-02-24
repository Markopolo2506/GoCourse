package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	addressBook := make(map[string][]int)

	addressBook["Alex"] = []int{8527885555}
	addressBook["Bob"] = []int{67243805212888}
	addressBook["Bob"] = append(addressBook["Bob"], 89528653752525)

	fmt.Println(addressBook)

	for name, numbers := range addressBook {
		fmt.Println("Номер:", name)
		for i, number := range numbers {
			fmt.Printf("\t %v) %v \n", i+1, number)
		}
	}

	addBookJson, err := json.Marshal(addressBook)
	if err != nil {
		fmt.Printf("Ошибка конвертации в json: %v\n", err)
	}

	err = ioutil.WriteFile("addBookJson.json", addBookJson, 0644)
	if err != nil {
		fmt.Printf("Ошибка записи в файл: %v\n", err)
	}
}
