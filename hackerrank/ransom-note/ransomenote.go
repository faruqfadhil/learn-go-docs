package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the checkMagazine function below.
func checkMagazine(magazine []string, note []string) {
	magazineMap := map[string]int{}
	magazineWordsCount := map[string]int{}
	isNo := false
	for _, n := range magazine {
		if _, ok := magazineMap[n]; ok {
			magazineWordsCount[n]++
		} else {
			magazineMap[n] = 0
			magazineWordsCount[n] = 1
		}
	}

	for _, n := range note {
		if _, ok := magazineMap[n]; ok {
			magazineMap[n]++
			if magazineMap[n] > magazineWordsCount[n] {
				fmt.Println("No")
				isNo = true
				break
			}
		} else {
			fmt.Println("No")
			isNo = true
			break
		}
	}
	if !isNo {
		fmt.Println("Yes")
	}

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	mn := strings.Split(readLine(reader), " ")

	mTemp, err := strconv.ParseInt(mn[0], 10, 64)
	checkError(err)
	m := int32(mTemp)

	nTemp, err := strconv.ParseInt(mn[1], 10, 64)
	checkError(err)
	n := int32(nTemp)

	magazineTemp := strings.Split(readLine(reader), " ")

	var magazine []string

	for i := 0; i < int(m); i++ {
		magazineItem := magazineTemp[i]
		magazine = append(magazine, magazineItem)
	}

	noteTemp := strings.Split(readLine(reader), " ")

	var note []string

	for i := 0; i < int(n); i++ {
		noteItem := noteTemp[i]
		note = append(note, noteItem)
	}

	checkMagazine(magazine, note)
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
