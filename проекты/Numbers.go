package main

import (
	"fmt"
)

func main() {
	var number int

	fmt.Print("Введите целое число: ")
	fmt.Scan(&number)

	switch {
	case number > 0:
		fmt.Println("Число положительное")
	case number < 0:
		fmt.Println("Число отрицательное")
	default:
		fmt.Println("Число равно нулю")
	}
}

// 1. Запрашиваем число у пользователя
// 2. Используем switch для проверки знака числа
// 3. Выводим соответствующее сообщение
