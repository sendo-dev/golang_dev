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

	/*
		昇順ソート：sort.Ints(list) ※Intsは sort.Sort(sort.IntSlice(numbers)) を呼んでいるだけ
		降順ソート：sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	*/
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
