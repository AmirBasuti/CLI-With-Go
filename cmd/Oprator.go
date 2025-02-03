package cmd

import (
	"fmt"
	"strconv"
)

func Add(arg1, arg2 string) string {
	num1, err := strconv.ParseFloat(arg1, 64)
	if err != nil {
		fmt.Println("first value is invalid")
		return ""
	}
	num2, err := strconv.ParseFloat(arg2, 64)
	if err != nil {
		fmt.Println("secound arg is not valid")
		return ""
	}
	return fmt.Sprintf("%f", num2+num1)

}

func Substract(arg1, arg2 string) string {
	num1, err := strconv.ParseFloat(arg1, 64)
	if err != nil {
		fmt.Println("first value is invalid")
		return ""
	}
	num2, err := strconv.ParseFloat(arg2, 64)
	if err != nil {
		fmt.Println("secound arg is not valid")
		return ""
	}
	return fmt.Sprintf("%f", num1-num2)
}
