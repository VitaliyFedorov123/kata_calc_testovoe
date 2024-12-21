package main

import (
	"fmt"
	"strconv"
	"strings"
)

func romanToArabic(roman string) int {
	romanNumbers := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}
	return romanNumbers[roman]
}

func arabicToRoman(arabic int) string {
	if arabic < 1 || arabic > 100 {
		return ""
	}

	values := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""
	for i := 0; i < len(values); i++ {
		for arabic >= values[i] {
			result += symbols[i]
			arabic -= values[i]
		}
	}
	return result
}

func main() {
	fmt.Println("Input:")

	for {
		fmt.Print("Введите выражение (числа от 1 до 10): ")
		var input string
		fmt.Scanln(&input)

		if input == "exit" {
			break
		}

		var operator string
		if strings.Contains(input, "+") {
			operator = "+"
		} else if strings.Contains(input, "-") {
			operator = "-"
		} else if strings.Contains(input, "*") {
			operator = "*"
		} else if strings.Contains(input, "/") {
			operator = "/"
		} else {
			fmt.Println("Выдача паники, так как строка не является математической операцией.")
			return
		}

		parts := strings.Split(input, operator)
		if len(parts) != 2 {
			fmt.Println("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
			return
		}

		num1Str := strings.TrimSpace(parts[0])
		num2Str := strings.TrimSpace(parts[1])

		var num1, num2 int
		var isRoman bool

		if strings.ContainsAny(num1Str, "IVX") && strings.ContainsAny(num2Str, "IVX") {
			num1 = romanToArabic(num1Str)
			num2 = romanToArabic(num2Str)
			isRoman = true
		} else if strings.ContainsAny(num1Str, "0123456789") && strings.ContainsAny(num2Str, "0123456789") {
			var err1, err2 error
			num1, err1 = strconv.Atoi(num1Str)
			num2, err2 = strconv.Atoi(num2Str)
			if err1 != nil || err2 != nil {
				fmt.Println("Ошибка: некорректный ввод чисел. Программа завершает работу.")
				return
			}
		} else {
			fmt.Println("Выдача паники, так как используются одновременно разные системы счисления.")
			return
		}

		if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
			fmt.Println("Выдача паники, ввод числел от 1-10 либо I-X")
			return
		}

		var result int
		switch operator {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			if num2 == 0 {
				fmt.Println("Нельзя делить на ноль.")
				return
			}
			result = num1 / num2
		}

		if result < 1 || result > 100 {
			fmt.Println("Ввод числел от 1-10 либо I-X.")
			return
		}

		if isRoman {
			fmt.Println("Output:", arabicToRoman(result))
		} else {
			fmt.Println("Output:", result)
		}
	}
}
