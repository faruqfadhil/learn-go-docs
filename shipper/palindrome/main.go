// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// /** To execute Go, please define "main()" as below **/
// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	var testCases int
// 	fmt.Scan(&testCases)
// 	for i := 0; i < testCases; i++ {
// 		strInput := strings.Split(readLine(reader), ",")
// 		// fmt.Println(isPalindrome(strInput[0]))
// 		i, _ := strconv.Atoi(strInput[0])
// 		fmt.Println(isPalindrome(i))
// 	}
// }

// func readLine(reader *bufio.Reader) string {
// 	str, _, err := reader.ReadLine()
// 	if err == io.EOF {
// 		return ""
// 	}

// 	return strings.TrimRight(string(str), "\r\n")
// }

// func recur(n, i int) int {
// 	if n == 0 {
// 		return i
// 	}

// 	i = (i * 10) + (n % 10)
// 	return recur(n/10, i)
// }

// func isPalindrome(n int) string {
// 	out := recur(n, 0)
// 	if out == n {
// 		return "Yes"
// 	}
// 	return "No"
// }

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/** To execute Go, please define "main()" as below **/
func main() {
	reader := bufio.NewReader(os.Stdin)
	var testCases int
	fmt.Scan(&testCases)
	for i := 0; i < testCases; i++ {
		strInput := strings.Split(readLine(reader), ",")
		i, _ := strconv.Atoi(strInput[0])
		fmt.Println(isPalindrome(i))
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func recur(n, i int) int {
	if n == 0 {
		return i
	}

	i = (i * 10) + (n % 10)
	return recur(n/10, i)
}

func isPalindrome(n int) string {
	out := recur(n, 0)
	if out == n {
		return "Yes"
	}
	return "No"
}

// Your Previous Javascript code is preserved below:
// /**
// 	Default code for reading input data
//
// process.stdin.resume();
// process.stdin.setEncoding("utf-8");
// var stdin_input = "";

// process.stdin.on("data", function (input) {
//     stdin_input += input;
// });

// process.stdin.on("end", function () {
// 	main(stdin_input);
// });

// /**
// 	Below "main" function will be called with input as string argument
//
// function main(input) {
// 	/**
// 		NOTE: Start modifying below code
// 		If necessary parse input as required in question and
// 		print your program's output using console.log
//
//  	console.log("Code output is: " + input + ".");
// }
