package main

import "fmt"

func Multiply(num1 int, num2 int) {

	var mult int=0

	mult = 0
	for count:=1; count<=num2; count++{
		mult = mult + num1
	}

	fmt.Printf("Multiplication is:%d", mult)
}