package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// usingScanf()
	// usingScan()
	// usingScanLn()
	// usingBufio()
	// intoIntSlice()
	intoStringSlice()
}

func usingScanf() {
	var i int
	fmt.Scanf("%d", &i)
	fmt.Println("hasil = ", i)

	var j string
	fmt.Scanf("%s", &j)
	fmt.Println("hasil = ", j)
}

func usingScan() {
	var i int
	fmt.Scan(&i)
	fmt.Println("hasil = ", i)

	var j string
	fmt.Scan(&j)
	fmt.Println("hasil = ", j)
}

func usingScanLn() {
	var i int
	fmt.Scanln(&i)
	fmt.Println("hasil = ", i)

	var j string
	fmt.Scanln(&j)
	fmt.Println("hasil = ", j)
}

func usingBufio() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func intoIntSlice() {
	var n int
	fmt.Scan(&n)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		text := strings.Split(readLine(reader), " ")
		anuInt := []int{}
		for _, n := range text {
			i, _ := strconv.Atoi(string(n))
			anuInt = append(anuInt, i)
		}
		fmt.Println(anuInt)
	}
}

func intoStringSlice() {
	var n int
	fmt.Scan(&n)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		text := strings.Split(readLine(reader), " ")
		fmt.Println(text)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
