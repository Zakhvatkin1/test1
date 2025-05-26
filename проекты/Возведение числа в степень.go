package main

import (
	"fmt"
	"math"
)

func main() {
	base := 3.00
	exponent := 2.55

	result := math.Pow(base, exponent)

	fmt.Printf("%.2f в степени %.2f = %.2f\n", base, exponent, result)
}

// Задаём основание степени через base и саму степень через exponen
// Задаём основание степени через base и саму степень через exponent
//Вычисляем результат с помощью функции math.Pow
