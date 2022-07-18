package main

import (
	"fmt"
	"sort"
)

func main() {

	var n int
	var list []int
	alice := 0
	bob := 0

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		tmpNum := 0
		fmt.Scan(&tmpNum)
		list = append(list, tmpNum)
	}
	// 逆順にソート
	sort.Sort(sort.Reverse(sort.IntSlice(list)))

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			alice += list[i]
		} else {
			bob += list[i]
		}
	}

	fmt.Print(alice - bob)

}
