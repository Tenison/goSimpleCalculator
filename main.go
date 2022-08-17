package main

import (
	"fmt"
	"goCalculator/calculator"
)


func main(){
	var (
		numberOne int
		numberTwo int
		operator string
	)
	fmt.Printf("Please input first number: ")
	fmt.Scanf("%d\n", &numberOne)

	fmt.Printf("Please enter calculation operator(+,-,/,*,%%): ")
	fmt.Scanf("%s\n", &operator)

	fmt.Printf("Please input second number: ")
	fmt.Scanf("%d\n", &numberTwo)

	switch string(operator){
	case "+":
		fmt.Println("The answer is ", calculator.Sum(numberOne, numberTwo))
	case "/":
		fmt.Println("The answer is ", calculator.Divide(numberOne, numberTwo))
	case "*":
		fmt.Println("The answer is ", calculator.Multiplication(numberOne, numberTwo))
	case "-":
		fmt.Println("The answer is ", calculator.Substract(numberOne, numberTwo))
	case "%":
		fmt.Println("The answer is ", calculator.Reminder(numberOne, numberTwo))
	default:
		fmt.Println("Not a valid operations")
	}
}