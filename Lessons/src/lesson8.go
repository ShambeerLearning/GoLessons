package src

import "fmt"

func Lesson8(){
	var inpp, summ int
	summ = 0
	fmt.Print("Введите число - ")
	fmt.Scanln(&inpp)
	for i := 1; i < inpp; i++ {
		summ = summ + i
	}
	fmt.Println("Сумма всех чисел от 1 до N:", summ)
	fmt.Scanln()
}