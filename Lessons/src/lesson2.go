package src

import "fmt"

func Lesson2(){
	var inpp string
	fmt.Print("Введите имя - ")
	fmt.Scanln(&inpp)
	fmt.Print("Привет, ", inpp, "!")
	fmt.Scanln()
}