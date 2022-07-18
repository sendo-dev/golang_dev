package main

import "fmt"

func main() {

	const tenThousand int = 10000
	const fiveThousand int = 5000
	const thousand int = 1000
	resultMap := map[string]int{"i": -1, "j": -1, "k": -1}
	var n int
	var y int

	fmt.Scan(&n)
	fmt.Scan(&y)

	endFlg := false
	for i := y / tenThousand; i >= 0; i-- {
		for j := 0; j <= n-i; j++ {
			k := n - i - j
			if tenThousand*i+fiveThousand*j+thousand*k == y {
				resultMap["i"] = i
				resultMap["j"] = j
				resultMap["k"] = k
				endFlg = true
				break
			}
		}
		if endFlg {
			break
		}
	}

	fmt.Print(resultMap["i"], resultMap["j"], resultMap["k"])

}
