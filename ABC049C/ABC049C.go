package main

import "fmt"

func main() {

	result := "YES"
	var S string
	fmt.Scan(&S)

	SLength := len(S)

	index := 0
	endIndex := 0
	for index <= SLength {
		if index == SLength {
			break
		}
		endIndex = index + 5
		if endIndex > SLength {
			result = "NO"
			break
		}
		if S[index:endIndex] == "dream" {
			index += 5
			endIndex = index + 2
			if endIndex > SLength {
				continue
			}
			if S[index:endIndex] == "er" {
				endIndex = index + 5
				if endIndex > SLength {
					index += 2
					continue
				}
				if S[index:endIndex] == "erase" {
					endIndex = index + 6
					if endIndex > SLength {
						index += 5
						continue
					}
					if S[index:endIndex] == "eraser" {
						index += 6
					} else {
						index += 5
					}
				} else {
					index += 2
				}
			}
		} else if S[index:endIndex] == "erase" {
			index += 5
			endIndex = index + 1
			if endIndex > SLength {
				continue
			}
			if S[index:endIndex] == "r" {
				index += 1
			}
		} else {
			result = "NO"
			break
		}
	}

	fmt.Print(result)

}
