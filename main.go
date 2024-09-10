package main

import (
	"log"
	"os"

	"github.com/fatih/color"
)

func main() {
	// Инициализация логгера
	logger := log.New(os.Stdout, "LOG: ", log.Ldate|log.Ltime)

	input := []int{1, 2, 3, 4, 5}
	output := make([]int, 0)

	// Логирования каждлого действия
	logger.Println("Начало паплайна")

	for _, num := range input {
		logger.Printf(
			"Обработка числа: %s%s",
			color.GreenString("%d", num),
			color.New(color.FgWhite).Sprint(""),
		)
		result := process(num)
		logger.Printf(
			"Резильтат обработки: %s%s",
			color.BlueString("%d", result),
			color.New(color.FgWhite).Sprint(""),
		)
		output = append(output, result)
	}

	logger.Println("Завершения пайплайна")
	color.Cyan(":-)")
	color.Magenta(":-D")
}

func process(num int) int {
	return num * 2
}
