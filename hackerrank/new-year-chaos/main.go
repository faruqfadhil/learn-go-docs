package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the minimumBribes function below.
func minimumBribes(q []int32) {
	n := len(q) - 1
	tempStorage := map[int32]int{}
	isSwap := true
	isChaotic := false
	res := 0
ForLoop:
	for isSwap {
		isSwap = false
		for i := 0; i <= n-1; i++ {
			if q[i] > q[i+1] {
				isSwap = true
				if _, ok := tempStorage[q[i]]; ok {
					tempStorage[q[i]]++
					if tempStorage[q[i]] > 2 {
						isChaotic = true
						fmt.Println("Too chaotic")
						break ForLoop
					}
				} else {
					tempStorage[q[i]] = 1
				}
				res++
				q[i], q[i+1] = q[i+1], q[i]
			}
		}
	}
	if !isChaotic {
		fmt.Println(res)
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(readLine(reader), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
