package src

import "fmt"

func Lesson5(){
	var inpp int64
	fmt.Println("Анализ числа")
	fmt.Print("Введите число - ")
	fmt.Scanln(&inpp)
	if inpp == 0 {
		fmt.Println("Число 0")
		fmt.Scanln()
		return
	} else if inpp < 0 {
		fmt.Println("Число отрицательное")
	} else {
		fmt.Println("Число положительное")
	}
	if inpp % 2 == 0 {
		fmt.Println("Число четное")
	} else if inpp % 2 != 0 {
		fmt.Println("Число нечетное")
	}
	fmt.Scanln()
}