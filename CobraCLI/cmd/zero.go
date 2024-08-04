package cmd

import (
	"fmt"
	"strconv"
)

func checkNums(first string, second string) (num1 float64, num2 float64, valid bool) {
	num1, err := strconv.ParseFloat(first, 64)
	if err != nil {
		fmt.Println("Invalid Num1 ", err)
		return -1, -1, false
	}
	num2, err = strconv.ParseFloat(second, 64)
	if err != nil {
		fmt.Println("Invalid Num2 ", err)
		return -1, -1, false
	}
	fmt.Println("Num1 and Num2 are valid")
	return num1, num2, true
}

func Add(first string, second string) (result string) {
	num1, num2, valid := checkNums(first, second)
	if !valid {
		fmt.Println("Invalid Num1 or Num2 ", first, second)
		return
	}
	return fmt.Sprintf("%f", num1+num2)
}
func Subtract(first string, second string) (result string) {
	num1, num2, valid := checkNums(first, second)
	if !valid {
		fmt.Println("Invalid Num1 or Num2 ", first, second)
		return
	}
	return fmt.Sprintf("%f", num1-num2)
}
func Multiple(first string, second string) (result string) {
	num1, num2, valid := checkNums(first, second)
	if !valid {
		fmt.Println("Invalid Num1 or Num2 ", first, second)
		return
	}
	return fmt.Sprintf("%f", num1*num2)
}
func Divide(first string, second string) (result string) {
	num1, num2, valid := checkNums(first, second)
	if !valid {
		fmt.Println("Invalid Num1 or Num2 ", first, second)
		return
	}
	if num2 == 0 {
		fmt.Println("Divide by Zero error")
		return
	}
	return fmt.Sprintf("%f", num1/num2)
}
