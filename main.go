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
		result, err := strconv.Atoi(proc(romanToNum(args), operator))
		if err != nil {
			log.Panicln(err)
		}
		if result < 1 {
			log.Panicln("Нет римских чисел меньше 1")
		}
		return numToRoman(proc(romanToNum(args), operator))
	} else {
		return proc(args, operator)
	}

}

func proc(args []string, operator string) string {
	var result int
	num1, err := strconv.Atoi(args[0])
	if err != nil || num1 > 10 {
		log.Panicln("не подходящее число")
	}
	num2, err := strconv.Atoi(args[1])
	if err != nil || num2 > 10 {
		log.Panicln("не подходящее число")
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
			log.Panicln("Деление на ноль")
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
				log.Panicln("Разные системы счисления")
			}
		} else {
			hasNum = true
			if hasRoman {
				log.Panicln("Разные системы счисления")
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
					log.Panicln("Превышено количество операторов")

				} else {
					found = true
					//log.Print("найден оператор: " + operator)
					result = operator
				}

			}
		}
	}
	if !found {
		log.Panicln("Не найдено ни одного оператора")
	}
	return result
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Panicln("Ошибка чтения строки консоли")
		}
		text = strings.TrimSpace(text)
		text = strings.ReplaceAll(text, " ", "")
		text = strings.ToUpper(text)
		fmt.Println(calculate(text))
	}

}
