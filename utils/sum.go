package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func Sum(a string, b string) float64 {

	float1, err := strconv.ParseFloat(strings.TrimSpace(a),64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}

	float2, err := strconv.ParseFloat(strings.TrimSpace(b),64)
	if err != nil {
		fmt.Println(err)
		panic("Value 2 must be a number")
	}

	return float1 + float2
}