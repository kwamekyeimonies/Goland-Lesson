package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var random_inte [6]int
	var lotto_numbers []int
	lotto_numbers = make([]int, 10)

	// var customers[string] int
	customers := make(map[string]int)
	users := make(map[int]string)

	customers["cust_1"] = rand.Intn(300)
	for dex := 0; dex < 30; dex++ {
		users[dex] = "Tenkorang Daniel"

	}
	fmt.Println(users)
	fmt.Println(customers)

	for i := 0; i <= 5; i++ {
		random_inte[i] = i
	}

	for j := 0; j < 10; j++ {
		lotto_numbers[j] = rand.Intn(100)
	}

	fmt.Println(random_inte)
	fmt.Println(lotto_numbers)
}
