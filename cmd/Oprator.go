package cmd

import (
	"errors"
	"fmt"
	"strconv"
)

func Add(arg1, arg2 string) (error, string) {
	num1, err := strconv.ParseFloat(arg1, 64)
	if err != nil {
		return errors.New("first value is invalid"), ""
	}
	num2, err := strconv.ParseFloat(arg2, 64)
	if err != nil {
		return errors.New("secound arg is not valid"), ""
	}
	return nil, fmt.Sprintf("%f", num2+num1)

}

func Substract(arg1, arg2 string) (error, string) {
	num1, err := strconv.ParseFloat(arg1, 64)
	if err != nil {
		return errors.New("first value is invalid"), ""
	}
	num2, err := strconv.ParseFloat(arg2, 64)
	if err != nil {
		return errors.New("secound arg is not valid"), ""
	}
	return nil, fmt.Sprintf("%f", num1-num2)
}

func Multiply(arg1, arg2 string, round bool) (error, string) {
	num1, err := strconv.ParseFloat(arg1, 64)
	if err != nil {
		return errors.New("first value is invalid"), ""
	}
	num2, err := strconv.ParseFloat(arg2, 64)
	if err != nil {
		return errors.New("secound arg is not valid"), ""
	}
	if round {
		return nil, fmt.Sprintf("%.2f", num2*num1)
	}
	return nil, fmt.Sprintf("%f", num2*num1)

}
