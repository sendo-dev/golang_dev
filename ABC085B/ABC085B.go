package main

import "fmt"

func main() {

	var n int
	mapEx := make(map[int]int)

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		mapEx[tmp] = tmp
	}

	fmt.Print(len(mapEx))

}
