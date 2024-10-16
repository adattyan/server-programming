package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Mul(a, b int) int {
	return a * b
}

func Reusu(scores []int, opertion func(int, int) int, def_val int) int {
	result := def_val

	for _, v := range scores {
		result = opertion(result, v)
	}
	return result
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	//スライスの足し算
	sum := Reusu(numbers, Add, 0)
	fmt.Println("Sum: ", sum)

	//スライスの掛け算
	mul := Reusu(numbers, Mul, 1)
	fmt.Println("Mul: ", mul)
}
