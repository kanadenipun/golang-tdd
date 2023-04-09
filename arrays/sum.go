package arrays

import "example/hello/integers"

func Sum(sumArray [5]int) int {
	var sumResult int
	for _, x := range sumArray {
		sumResult = integers.Add(sumResult, x)
	}
	return sumResult
}
