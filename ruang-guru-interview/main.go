package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var testCases int
	fmt.Scan(&testCases)
	for i := 0; i < testCases; i++ {
		strInput := strings.Split(readLine(reader), ",")
		fmt.Println(isAnagram(strInput))
	}

}

func isAnagram(inputs []string) bool {
	mapKeys := map[string][]string{}
	for _, i := range inputs {
		sTemp := strings.Split(string(i), "")
		sort.Strings(sTemp)
		key := strings.Join(sTemp, "")

		mapKeys[key] = append(mapKeys[key], string(i))
	}

	for key, v := range mapKeys {
		if len(v) > 1 {
			fmt.Printf("Found anagram words!:%s\t in slice:%s\n", v, key)
			return true
		}

	}

	return false
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
