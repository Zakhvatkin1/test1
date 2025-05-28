package main

import "fmt"

func main() {
	a := 2
	b := 8

	fmt.Println("Кубы чисел от", a, "до", b, ":")
	for i := a; i <= b; i++ {
		cube := i * i * i
		fmt.Printf("%d в кубе = %d\n", i, cube)
	}
}

//for i := a; i <= b; i++ - цикл от наших переменных a до b включительно с шагом 1
//	// cube := i * i * i - вычисляем куб текущего числа i
//	// fmt.Printf("%d в кубе = %d\n", i, cube) - форматированный вывод числа и его куба
