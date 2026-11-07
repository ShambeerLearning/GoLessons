package src

import "fmt"

func Lesson4(){
	var str1, str2, P, S int64
	fmt.Println("Расчёт периметра и площади прямоугольника")
	fmt.Print("Введите длину - ")
	fmt.Scanln(&str1)
	fmt.Print("Введите ширину - ")
	fmt.Scanln(&str2)
	P = str1 * 2 + str2 * 2
	S = str1 * str2
	fmt.Println("\nПериметр прямоугольника -", P)
	fmt.Println("Площадь прямоугольника -", S)
	fmt.Scanln()
}