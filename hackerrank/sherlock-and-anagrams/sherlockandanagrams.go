package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the sherlockAndAnagrams function below.
func sherlockAndAnagrams(s string) int32 {
	// i := 0
	// for jarak := 0; jarak < len(s); jarak++ {
	// 	key := i + jarak
	// 	for key < len(s) {
	// 		targetKey := string(s[i:key])
	// 		j := i + jarak
	// 		jKey := j + jarak
	// 		for jKey < len(s) {
	// 			current := string(s[j:jKey])

	// 			j++
	// 			jKey = j + jarak
	// 		}
	// 		i++
	// 		key = i + jarak
	// 	}
	// }
	return 0
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := sherlockAndAnagrams(s)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
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
