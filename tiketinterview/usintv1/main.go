package main

import "fmt"

// Input :
// []{}()
// {[(){}[]]}
// [][[]]()(){}{}
// [{]}
// }{

// Output
// []{}() = true
// {[(){}[]]} = true
// [][[]]()(){}{} = true
// [{]} = false
// }{ = false

func main() {
	input := "[]{}()"
	fmt.Println(validate(input))
	input = "{[(){}[]]}"
	fmt.Println(validate(input))
	input = "[][[]]()(){}{}"
	fmt.Println(validate(input))
	input = "[{]}"
	fmt.Println(validate(input))
	input = "}{"
	fmt.Println(validate(input))

}

func isIdentic(a, b string) bool {
	if a == "{" && b == "}" {
		return true
	}
	if a == "[" && b == "]" {
		return true
	}
	if a == "(" && b == ")" {
		return true
	}
	return false
}

func validate(input string) bool {
	charcter := map[string]string{
		"[": "]",
		"]": "]",
		"{": "}",
		"}": "{",
		"(": ")",
		")": "(",
	}
	characterClose := map[string]bool{
		"]": true,
		"}": true,
		")": true,
	}
	characterOpen := map[string]bool{
		"{": true,
		"[": true,
		"(": true,
	}

	if _, ok := characterOpen[string(input[len(input)-1])]; ok {
		return false
	}

	// var resolve bool
	firstItem := string(input[0])
	if _, ok := charcter[firstItem]; !ok {
		return false
	}
	if _, ok := characterClose[firstItem]; ok {
		return false
	}

	resolveCount := map[string]int{}
	var (
		key1 string = "keyFor}"
		key2 string = "keyFor]"
		key3 string = "keyFor)"
	)

	for i := 0; i < len(input); i++ {
		if _, ok := characterOpen[string(input[i])]; ok {
			if charSejenis, ok := charcter[string(input[i])]; ok {
				inp := string(input[i])
				if charSejenis == key1 || inp == key1 {
					resolveCount[key1]++
				} else if charSejenis == key2 || inp == key2 {
					resolveCount[key2]++
				} else if charSejenis == key3 || inp == key3 {
					resolveCount[key3]++
				}
			}
		}

		if _, ok := characterClose[string(input[i])]; ok {
			if charSejenis, ok := charcter[string(input[i])]; ok {
				inp := string(input[i])
				if charSejenis == key1 || inp == key1 {
					resolveCount[key1]--
				} else if charSejenis == key2 || inp == key2 {
					resolveCount[key2]--
				} else if charSejenis == key3 || inp == key3 {
					resolveCount[key3]--
				}
			}
		}

		// if i == len(input)-1 {
		// 	for , count := range resolveCount {
		// 		if count == 1 && string(input[i]) != {
		// 		}
		// 	}
		// 	if string(input[i]) != firstItem {
		// 		return false
		// 	}
		// }

		if i == len(input)-1 {
			a := string(input[len(input)-2])
			b := string(input[len(input)-1])

			for key, count := range resolveCount {
				if count == 1 {
					switch key {
					case key1:
						if !isIdentic("}", b) && !isIdentic("{", b) {
							return false
						}
					case key2:
						if !isIdentic("]", b) && !isIdentic("[", b) {
							return false
						}
					case key3:
						if !isIdentic("(", b) && !isIdentic(")", b) {
							return false
						}
					}
					// if !isIdentic(a, b) {
					// 	return false
					// }
				}
			}

			// charA := charcter[a]
			// charB := charcter[b]

			// if  != ) {
			// 	return false
			// }
		}
	}

	allResolved := true
	for _, val := range resolveCount {
		if val != 0 {
			allResolved = false
		}
	}

	lastItem := string(input[len(input)-1])
	if !allResolved && lastItem == firstItem {
		return false
	}
	// if allResolved && lastItem != firstItem {
	// 	return false
	// }
	// if notResolved {
	// 	lastItem := string(input[len(input)-1])
	// 	if lastItem != firstItem {
	// 		return false
	// 	}
	// }
	// for resolve {
	// 	for i := 1; i < len(input); i++ {
	// 		// if not valid char
	// 		if _, ok := charcter[string(input[i])]; !ok {
	// 			return false
	// 		}
	// 		if _, ok := characterOpen[temp]; ok {
	// 			if _, ok := characterOpen[string(input[i])]; ok {
	// 				resolve = true
	// 				temp = string(input[i])
	// 			}
	// 		}

	// 		// // if begin with char close.
	// 		// if _, ok := characterClose[string(input[i])]; ok && i == 0 {
	// 		// 	return false
	// 		// }

	// 		// if finished with open char.
	// 		// if _, ok := characterOpen[string(input[len(input)-1])]; ok {
	// 		// 	return false
	// 		// }
	// 	}
	// }

	return true
}
