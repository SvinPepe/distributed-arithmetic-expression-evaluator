package main

import (
	"strconv"
	"strings"
	"time"
)

func Count(equasion string) float64 {
	postfixNotation := strings.Split(equasion, " ")
	for i, _ := range postfixNotation {
		if postfixNotation[i] == "+" {
			s, err := strconv.ParseFloat(postfixNotation[i-1], 64)
			if err != nil {
				currentStatus = "invalid"
				panic(err)
			}
			s2, err := strconv.ParseFloat(postfixNotation[i-2], 64)
			if err != nil {
				currentStatus = "invalid"
				panic(err)
			}
			postfixNotation[i] = strconv.FormatFloat(Add(s, s2), 'f', 6, 64)
			postfixNotation[i-1] = "0"
		}
		if postfixNotation[i] == "-" {
			s, err := strconv.ParseFloat(postfixNotation[i-1], 64)
			if err != nil {
				currentStatus = "invalid"
				panic(err)
			}
			s2, err := strconv.ParseFloat(postfixNotation[i-2], 64)
			if err != nil {
				currentStatus = "invalid"
				panic(err)
			}
			postfixNotation[i] = strconv.FormatFloat(Subtract(s, s2), 'f', 6, 64)
			postfixNotation[i-1] = "0"
		}
		if postfixNotation[i] == "*" {
			s, err := strconv.ParseFloat(postfixNotation[i-1], 64)
			if err != nil {
				currentStatus = "invalid"
				panic(err)
			}
			s2, err := strconv.ParseFloat(postfixNotation[i-2], 64)
			if err != nil {
				currentStatus = "invalid"
				panic(err)
			}
			postfixNotation[i] = strconv.FormatFloat(Multiply(s, s2), 'f', 6, 64)
			postfixNotation[i-1] = "0"
		}
		if postfixNotation[i] == "/" {
			s, err := strconv.ParseFloat(postfixNotation[i-1], 64)
			if err != nil {
				currentStatus = "invalid"
				panic(err)
			}
			s2, err := strconv.ParseFloat(postfixNotation[i-2], 64)
			if err != nil {
				currentStatus = "invalid"
				panic(err)
			}
			postfixNotation[i] = strconv.FormatFloat(Divide(s, s2), 'f', 6, 64)
			postfixNotation[i-1] = "0"
		}
	}
	ans, err := strconv.ParseFloat(postfixNotation[len(postfixNotation)-1], 64)
	if err != nil {
		currentStatus = "invalid"
		panic(err)
	}
	currentStatus = "done"
	currentResult = ans
	return ans
}

func Add(a, b float64) float64 {
	time.Sleep(1 * time.Second)
	return a + b
}
func Subtract(a, b float64) float64 {
	time.Sleep(1 * time.Second)
	return a - b
}
func Multiply(a, b float64) float64 {
	time.Sleep(1 * time.Second)
	return a * b
}
func Divide(a, b float64) float64 {
	time.Sleep(1 * time.Second)
	return a / b
}
