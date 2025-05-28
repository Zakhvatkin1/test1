package main

import (
	"fmt"
)

func main() {
	num := 10
	result := factorial(num)
	fmt.Printf("Факториал числа %d = %d\n", num, result)
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// В функции main инициализируем переменную, для которой высчитываем факториал
// Далее выводит результат факториала в числовом значении num
// В функции factorial прописываем все условия для вычисления факториала.
