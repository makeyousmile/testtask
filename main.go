package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var operators = []string{"+", "-", "*", "/"}

func calculate(text string) string {
	operator := getOperator(text)
	args, roman := getArgs(text)
	if roman {
		return numToRoman(proc(romanToNum(args), operator))
	} else {
		return proc(args, operator)
	}

}

func proc(args []string, operator string) string {
	var result int
	num1, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(err)
	}
	num2, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatalln(err)
	}
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			log.Fatalln("Деление на ноль")
		}
		result = num1 / num2

	}

	return strconv.Itoa(result)
}

func romanToNum(args []string) []string {

	num1 := romans[args[0]]
	num2 := romans[args[1]]

	return []string{num1, num2}
}
func numToRoman(num string) string {
	for roman, i := range romans {

		if i == num {
			return roman
		}
	}
	return ""
}

// Возвращает аргументы. True если римские
func getArgs(text string) ([]string, bool) {
	hasRoman := false
	hasNum := false

	args := strings.Split(text, getOperator(text))

	for _, elem := range args {
		_, err := strconv.Atoi(elem)
		if err != nil {
			hasRoman = true
			if hasNum {
				log.Fatalln("Разные системы счисления")
			}
		} else {
			hasNum = true
			if hasRoman {
				log.Fatalln("Разные системы счисления")
			}
		}
	}
	return args, hasRoman
}
func getOperator(text string) string {
	result := ""
	found := false
	for _, operator := range operators {
		for _, val := range text {
			if operator == string(val) {
				if found {
					log.Fatalln("Превышено количество операторов")

				} else {
					found = true
					//log.Print("найден оператор: " + operator)
					result = operator
				}

			}
		}
	}
	if !found {
		log.Fatalln("Не найдено не одного оператора")
	}
	return result
}
func main() {
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Ошибка чтения строки консоли")
	}
	text = strings.TrimSpace(text)
	text = strings.ReplaceAll(text, " ", "")
	fmt.Println(calculate(text))
}
