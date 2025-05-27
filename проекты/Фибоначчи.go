package main

import (
	"fmt"
)

func main() {
	n := 10
	fmt.Printf("Ряд чисел Фибоначчи, состоящий из %d элементов:\n", n)

	// Используем итеративный подход вместо рекурсивного
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a)
		a, b = b, a+b
	}
	fmt.Println()

}
