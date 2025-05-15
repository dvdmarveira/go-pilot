package main

import "fmt"

func main() {
	var pessoas = map[string]int{}
	pessoas["dvd"] = 23
	pessoas["isis"] = 19

	if idade, ok := pessoas["dvd"]; ok {
			fmt.Println("Pessoa existe no map", idade, ok)
	} else {
		fmt.Println("Pessoa não existe no map")
	}

}
