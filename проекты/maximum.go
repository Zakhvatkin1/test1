package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	fmt.Print("Введите первое число: ")
	fmt.Scan(&num1)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&num2)

	if num1 > num2 {
		fmt.Println("Большее число:", num1)
	} else if num2 > num1 {
		fmt.Println("Большее число:", num2)
	} else {
		fmt.Println("Числа равны")
	}

}

// - Объявляем переменные одной строкой
//- Сравниваем числа через if-else
// - Выводим результат в зависимости от переменных
