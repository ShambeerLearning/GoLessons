package src

import "fmt"

func Lesson3(){
	const euro float64 = 92.01
	var inpp, result float64
	fmt.Println("Конвертер валют (Рубли -> Евро)")
	fmt.Print("Введите рубли - ")
	fmt.Scan(&inpp)
	result = inpp * 1 / euro
	fmt.Println("Итог в евро -", result)
	fmt.Scanln()
}