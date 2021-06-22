package main

import (
	"fmt"
	"os"
)
//map для хранения номера итерации, и значения предыдущих вычислений
var GLOBAL_prevResultsContainer = map[int32]int32{}

func main() {
	var iteration int32
	var continueStr string = "no"
	for {
		fmt.Print("Введите N ряда фибоначчи: ")
		var scN, inputErr = fmt.Scanln(&iteration)
		//Парсинг выполнен корректно? Введено число?
		if scN == 0 {
			fmt.Print(inputErr)
			os.Exit(1)
		}
		var fibValue int32
		//Проверка, вычислялось ли ранее число с данной N
		if GLOBAL_prevResultsContainer[iteration] != 0 {
			fibValue = GLOBAL_prevResultsContainer[iteration]
		}else {
			fibValue = FibonachchiCacl(iteration)
			GLOBAL_prevResultsContainer[iteration] = fibValue
		}
		fmt.Printf("Результат: %d \n", fibValue )
		fmt.Print("Exit ?: yes/no \n")
		fmt.Scanln(&continueStr)
		if continueStr == "yes" {
			break
		}
	}
}

func FibonachchiCacl(iteration int32 ) int32 {
	//условие выхода из рекурсии
	if (iteration == 1 || iteration == 2) {
		return 1
	}else{
		var curIterationValue,prevIterationValue int32
		//Ищем первый ключ
		if GLOBAL_prevResultsContainer[iteration-1] != 0 {
			curIterationValue = GLOBAL_prevResultsContainer[iteration-1]
		}else{
			curIterationValue = FibonachchiCacl(iteration - 1)
		}
		//Ищем второй ключ
		if GLOBAL_prevResultsContainer[iteration-2] != 0 {
			prevIterationValue = GLOBAL_prevResultsContainer[iteration-2]
		}else{
			prevIterationValue = FibonachchiCacl(iteration - 2)
		}
		return curIterationValue + prevIterationValue
	}
}
