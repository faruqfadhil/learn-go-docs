package main

import "fmt"

func main() {
	// fmt.Println(faktorial(3))
	// toFibonacci(11)
	toPrima(15)
}

func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * faktorial(n-1)
}

func toFibonacci(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(fibonacci(i))
	}
}

func fibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func toPrima(n int) {
	for i := 2; i <= n; i++ {
		if prima(i, i-1) == 1 {
			fmt.Println(i)
		}
	}
}

func prima(a int, n int) int {
	if n == 1 {
		return 1
	}
	if a%n == 0 {
		return 0
	}
	return prima(a, n-1)
}
