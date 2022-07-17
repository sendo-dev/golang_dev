package main

import "fmt"

func main() {

	var fiveHundredsCoin int
	var oneHundredsCoin int
	var fiftyCoin int
	var sum int
	var result int = 0

	fmt.Scan(&fiveHundredsCoin)
	fmt.Scan(&oneHundredsCoin)
	fmt.Scan(&fiftyCoin)
	fmt.Scan(&sum)

	for i := 0; i <= fiveHundredsCoin; i++ {
		if sum == 500*i {
			result++
		} else if sum > 500*i {
			for j := 0; j <= oneHundredsCoin; j++ {
				if sum == 500*i+100*j {
					result++
				} else if sum > 500*i+100*j {
					for k := 0; k <= fiftyCoin; k++ {
						if sum == 500*i+100*j+50*k {
							result++
						} else if sum < 500*i+100*j+50*k {
							break
						}
					}
				} else if sum < 500*i+100*j {
					break
				}
			}
		} else if sum < 500*i {
			break
		}
	}

	fmt.Print(result)

}
