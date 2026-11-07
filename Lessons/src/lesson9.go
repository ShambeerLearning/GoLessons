package src

import "fmt"

func Lesson9(){
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++{
			fmt.Println(i, " * ", j, " = ", i*j)
		}
		fmt.Println("")
	}
	fmt.Scanln()
}