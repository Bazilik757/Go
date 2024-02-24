package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		var operator string
		fmt.Print("Введите операцию (или 'exit' для выхода): ")
		fmt.Scanln(&operator)
		if operator == "exit" {
			fmt.Println("Программа завершена.")
			os.Exit(0)
		}
		result := evalExpression(operator)
		fmt.Println("Результат:", result)
		fmt.Println()
	}
}

func evalExpression(operator string) float64 {
	var a, b float64
	var op rune
	_, err := fmt.Sscanf(operator, "%f%c%f", &a, &op, &b)
	if err != nil {
		fmt.Println("Ошибка ввода")
		return 0
	}

	var result float64
	switch op {
	case '+':
		result = a + b
	case '-':
		result = a - b
	case '*':
		result = a * b
	case '/':
		result = a / b
		if b == 0 {
			fmt.Println("На ноль делить нельзя!")
			return 0
		}
	default:
		fmt.Println("Неподдерживаемая операция")
		return 0
	}

	return result
}
