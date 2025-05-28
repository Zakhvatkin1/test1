package main

import "fmt"

func main() {
	var score int
	fmt.Print("Введите количество баллов (0-100): ")
	fmt.Scan(&score)

	switch {
	case score > 100:
		fmt.Println("Ошибка: балл не может быть больше 100!")
	case score >= 90:
		fmt.Println("Отлично")
	case score >= 70:
		fmt.Println("Хорошо")
	case score >= 50:
		fmt.Println("Удовлетворительно")
	default:
		fmt.Println("Неудовлетворительно")
	}
}
