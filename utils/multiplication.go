package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func Multiplication(a string, b string) float64 {
	number1, err := strconv.ParseFloat(strings.TrimSpace(a), 64)
	if err != nil {
		fmt.Println(err)
	}

	number2, err := strconv.ParseFloat(strings.TrimSpace(b), 64)
	if err != nil {
		fmt.Println(err)
	}

	return number1 * number2
}