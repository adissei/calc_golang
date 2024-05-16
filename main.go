package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	)

func isArabNum(str string) bool {
	normalNum := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	for _, numeral := range normalNum {
		if numeral == str {
			return true
		}
	}
	return false
	}	
func isRomanNum(str string) bool {
	romanNum := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for _, numeral := range romanNum {
		if numeral == str {
			return true
		}
	}
	return false
	}
func romanToArabic(numeral string) int {
	switch numeral {
		case "I":
		 return 1
		case "II":
		 return 2
		case "III":
		 return 3
		case "IV":
		 return 4
		case "V":
		 return 5
		case "VI":
		 return 6
		case "VII":
		 return 7
		case "VIII":
		 return 8
		case "IX":
		 return 9
		case "X":
		 return 10
		default:
		 return 0
		}
	}

	func calc(x, operator, y string) {
		var result int
		var number1, number2 int
	
		if isRomanNum(x) && isRomanNum(y) {
			number1 = romanToArabic(x)
			number2 = romanToArabic(y)
		} else if isArabNum(x) && isArabNum(y) {
			number1, _ = strconv.Atoi(x)
			number2, _ = strconv.Atoi(y)
		} else {
			fmt.Println("Введены неправильные данные")
			return
		}
	
		switch operator {
		case "+":
			result = number1 + number2
		case "-":
			result = number1 - number2
		case "*":
			result = number1 * number2
		case "/":
			result = number1 / number2
		}
	
		fmt.Printf("Результат: %d\n", result)
	}

func main() {

	var line string
	fmt.Println("Введите выражние вида (2 + 1) или (V-III)")
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	line = reader.Text()
	match := strings.Split(line, " ")

	if len(match)==3{
			number1, operator, number2 := match[0], match[1], match[2]
			calc(number1, operator, number2)
		} else {
			fmt.Println("Ошибка выражение должно состоять из 2 элементов и знака")
			return
	}


	}
	

