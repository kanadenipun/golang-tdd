package arrays

import "example/hello/integers"

func Sum(sumArray []int) int {
	var sumResult int
	for _, x := range sumArray {
		sumResult = integers.Add(sumResult, x)
	}
	return sumResult
}

func SumAll(numbersToSum ...[]int) []int {
	var result = make([]int, len(numbersToSum))
	for i, x := range numbersToSum {
		result[i] = Sum(x)
	}
	return result
}
