package src

import "fmt"

func Lesson7(){
	var inpp uint16
	fmt.Print("Введите год, чтобы проверить его високосность - ")
	fmt.Scanln(&inpp)
	if inpp % 4 == 0 && inpp % 100 != 0 || inpp % 400 == 0 {
		fmt.Println("Год високосный")
	} else {
		fmt.Println("Год не високосный")
	}
	fmt.Scanln()
}