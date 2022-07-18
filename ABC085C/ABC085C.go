package main

import "fmt"

func main() {

	x, y, z := -1, -1, -1
	var N int
	var Y int

	fmt.Scan(&N)
	fmt.Scan(&Y)

	for i := Y / 10000; i >= 0; i-- {
		for j := (Y - 10000*i) / 5000; j >= 0; j-- {
			k := N - i - j
			if 10000*i+5000*j+1000*k == Y {
				x, y, z = i, j, k
				break
			}
		}
		if x > -1 {
			break
		}
	}

	fmt.Print(x, y, z)

}
