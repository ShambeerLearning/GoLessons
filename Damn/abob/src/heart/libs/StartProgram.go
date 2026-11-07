package libs

import "fmt"

func StartProgram() {
	var inpp, result float64
	const euro float64 = 0.94
	fmt.Printf("Введите рубли - ")
	fmt.Scanln(&inpp)
	result = inpp * euro
	fmt.Print("Итог в евро - ", result)
	ConsolePause()
}
