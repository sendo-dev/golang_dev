package main

import "fmt"

func main() {

	result := "Yes"
	var slice [][]int
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var tmpx, tmpy, tmpz int
		fmt.Scan(&tmpx, &tmpy, &tmpz)
		arr := []int{tmpx, tmpy, tmpz}
		slice = append(slice, arr)
	}

	for i := 0; i < N; i++ {
		if i == 0 {
			t := slice[i][0]
			x := slice[i][1]
			if t >= x {
				dif := t - x
				y := slice[i][2]
				if dif < y || (dif-y)%2 != 0 {
					result = "No"
					break
				}
			} else {
				result = "No"
				break
			}
		} else {
			dift := slice[i][0] - slice[i-1][0]
			difx := slice[i][1] - slice[i-1][1]
			if dift >= difx {
				dif := dift - difx
				dify := slice[i][2] - slice[i-1][2]
				if dif < dify || (dif-dify)%2 != 0 {
					result = "No"
					break
				}
			} else {
				result = "No"
				break
			}
		}
	}

	fmt.Print(result)

}
