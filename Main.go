package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, res float32
	var op string
	var continueStr string = "no"
	for {
		fmt.Print("Введите первое число: ")
		var scN, inputErr = fmt.Scanln(&a)
		//Парсинг выполнен корректно? Введено число?
		if scN == 0 {
			fmt.Print(inputErr)
			os.Exit(1)
		}
		fmt.Print("Введите второе число: ")
		fmt.Scanln(&b)
		if scN == 0 {
			fmt.Print(inputErr)
			os.Exit(1)
		}

		fmt.Print("Введите арифметическую операцию (+, -, *, /, pow): ")
		fmt.Scanln(&op)
		switch op {
		case "+":
			res = a + b
		case "-":
			res = a - b
		case "*":
			res = a * b
		case "/":
			if b > 0 {
				res = a / b
			} else {
				fmt.Println("Деление на ноль")
				os.Exit(1)
			}
		case "pow":
			if !(a < 0 && b <= 1) {
				res = float32(math.Pow(float64(a), float64(b)))
			} else {
				fmt.Println("Некорректные аргументы")
				os.Exit(1)
			}
		default:
			fmt.Println("Операция выбрана неверно")
			os.Exit(1)
		}

		fmt.Printf("Результат выполнения операции: %f\n", res)
		fmt.Print("Exit ?: yes/no \n")
		fmt.Scanln(&continueStr)
		if continueStr == "yes" {
			break
		}
	}
}

