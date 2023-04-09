package arrays

func Sum(sumArray [5]int) int {
	var sumResult int
	for _, x := range sumArray {
		sumResult += x
	}
	return sumResult
}
