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
	var result []int
	for _, x := range numbersToSum {
		result = append(result, Sum(x))
	}
	return result
}

func SumAllTrails(numbersToSum ...[]int) []int {
	var result []int
	var sum int
	for _, x := range numbersToSum {
		if len(x) < 1 {
			sum = 0
		} else {
			sum = Sum(x[1:])
		}
		result = append(result, sum)
	}
	return result
}
