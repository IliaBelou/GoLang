package main

import (
	"fmt"
	"math"
)

func main() {
	var continueStr string = "no"
	var programNumberToExec int16
	for {
		fmt.Print("Choose program to execute: \n")
		fmt.Scanln(&programNumberToExec)
		switch programNumberToExec {
		case 1:
			//Задание 1
			//Напишите программу для вычисления площади прямоугольника.
			//Длины сторон прямоугольника должны вводиться пользователем с клавиатуры.
			CalcSquareArea()
		case 2:
			//Задание 2
			//Напишите программу, вычисляющую диаметр и длину окружности по
			//заданной площади круга. Площадь круга должна вводиться пользователем с клавиатуры.
			CalcCircleDandL()
		case 3:
			//Задание 3
			//С клавиатуры вводится трехзначное число. Выведите цифры, соответствующие количество сотен,
			//десятков и единиц  в этом числе.
			NumberDecomp()
		}
		fmt.Print("Exit ?: yes/no \n")
		fmt.Scanln(&continueStr)
		if continueStr == "yes" {
			break
		}
	}
}
//Задание 1
func CalcSquareArea() {
	var squareLenght,squareHeight int32
	fmt.Print("SquareAreaCalculation \n")
	fmt.Print("Input square lenght in mm: \n")
	fmt.Scanln(&squareLenght)

	fmt.Print("Input square height in mm: \n")
	fmt.Scanln(&squareHeight)

	fmt.Printf("SquareArea: %d mm \n", squareLenght*squareHeight)
}
//Задание 2
func CalcCircleDandL() {
	var circleArea,circleDiameter,circleLenght float64
	fmt.Print("Circle diameter and lenght calculation  \n")
	fmt.Print("Input circle area in mm: \n")
	fmt.Scanln(&circleArea)

	circleDiameter = math.Pow((circleArea/math.Pi),0.5)*2
	fmt.Printf("Circle diameter: %f mm \n", circleDiameter  )

	circleLenght = 2*math.Pi*(circleDiameter/2)
	fmt.Printf("Circle fullarc lenght: %f mm \n", circleLenght  )
}
//Задание 3
func NumberDecomp() {
	var numberToDecomp,hundredsInNumber,dozensInNumber,digitsInNumber int32
	fmt.Print("Number decomposition  \n")
	fmt.Print("Input number to decompose: \n")
	fmt.Scanln(&numberToDecomp)

	if numberToDecomp > 99 {
		hundredsInNumber = numberToDecomp / 100
	}
	if (numberToDecomp - hundredsInNumber*100) > 9 {
		dozensInNumber = (numberToDecomp - hundredsInNumber*100)/10
	}
	digitsInNumber = numberToDecomp - hundredsInNumber*100 - dozensInNumber*10

	fmt.Printf("Hundreds:%d Dozens:%d Digits:%d \n",hundredsInNumber,dozensInNumber,digitsInNumber )
}
