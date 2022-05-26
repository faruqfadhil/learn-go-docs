package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// // // ----------------------------------------------------------
// // // Q1:
// // // make a function to solve a simple mathematic expression
// // // that only consists of +, -, * operations and positive integers
// // // input: String
// // // output: integer

// // // example
// // // input: "0 - 5"
// // // output: -5

//
// // // input: "1 - 7 * 2 + 5" -> 1 - 14 + 5 -> -13 + 5
// // // output: -8

// // // input: "12 * 5 * 3"
// // // output: 180
// // // -------------------------

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func main() {
	fmt.Println("Input the mathematical expression\nNote: between operand and operator should be separated by whitespace")
	reader := bufio.NewReader(os.Stdin)
	inputInArray := strings.Split(readLine(reader), " ")

	newArr := inputInArray
	needMultiplied := true
	for needMultiplied {
		needMultiplied = false
		last := false
		inputInArray = newArr
		for i := 1; i < len(newArr); i += 2 {
			firstNum := toNumber(newArr[i-1])
			secondNum := toNumber(newArr[i+1])

			// Handle if in the last operator.
			if i+2 > len(newArr)-1 {
				last = true
			}

			if newArr[i] == "*" {
				newArr = inputInArray[:i-1]
				newArr = append(newArr, fmt.Sprintf("%d", firstNum*secondNum))
				if !last {
					newArr = append(newArr, inputInArray[i+2:]...)
				}
				inputInArray = newArr
				needMultiplied = true
			} else if newArr[i] == "/" {
				newArr = inputInArray[:i-1]
				newArr = append(newArr, fmt.Sprintf("%d", firstNum/secondNum))
				if !last {
					newArr = append(newArr, inputInArray[i+2:]...)
				}
				inputInArray = newArr
				needMultiplied = true
			}
		}
	}

	agg, err := strconv.Atoi(newArr[0])
	if err != nil {
		log.Fatalf(err.Error())
	}
	for i := 1; i < len(newArr); i += 2 {
		num, err := strconv.Atoi(newArr[i+1])
		if err != nil {
			log.Fatalf(err.Error())
		}
		if newArr[i] == "-" {
			agg -= num
		}
		if newArr[i] == "+" {
			agg += num
		}
	}
	fmt.Println("Result = ", agg)
}

func toNumber(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("error when parsing to num: %s", err)
	}
	return num
}
