package libs

import (
	"fmt"
	"os"
)

func ConsoleExit() {
	fmt.Println("\nНажмите Enter для выхода...")
	var input string
	fmt.Scanln(&input)
	os.Exit(0)
}

func ConsolePause() {
	fmt.Println("\nНажмите Enter для выхода...")
	var input string
	fmt.Scanln(&input)
}
