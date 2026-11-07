package src

import "fmt"

func Lesson6(){
	fmt.Println("Калькулятор")
	var inpp1, inpp2, res float64
	var oper string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&inpp1)
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&inpp2)

	fmt.Println("\nВыбор действий:")
	fmt.Println("+ - Сложение")
	fmt.Println("- - Вычитание")
	fmt.Println("* - Умножение")
	fmt.Println("/ - Деление")
	fmt.Print("Введите действие: ")
	fmt.Scanln(&oper)

	if oper == "+" {
		res = inpp1 + inpp2
		fmt.Println("\nРезультат сложения:", res)
	} else if oper == "-" {
		res = inpp1 - inpp2
		fmt.Println("\nРезультат вычитания:", res)
	} else if oper == "*" {
		res = inpp1 * inpp2
		fmt.Println("\nРезультат умножения:", res)
	} else if oper == "/" && inpp2 != 0 {
		res = inpp1 / inpp2
		fmt.Println("\nРезультат деления:", res)
	}

	fmt.Scanln()
}