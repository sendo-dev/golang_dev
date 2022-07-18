package main

import (
	"fmt"
	"strconv"
)

func main() {

	var n int
	var a int
	var b int
	var sum int = 0

	fmt.Scan(&n)
	fmt.Scan(&a)
	fmt.Scan(&b)

	for i := 1; i <= n; i++ {
		tmpSum := 0
		for j := 0; j < len(strconv.Itoa(i)); j++ {
			tmpInt, _ := strconv.Atoi(strconv.Itoa(i)[j : j+1])
			tmpSum += tmpInt
			if tmpSum > b {
				break
			}
		}
		if tmpSum >= a && tmpSum <= b {
			sum += i
		}
	}

	fmt.Print(sum)

}
