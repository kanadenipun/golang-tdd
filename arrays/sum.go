package arrays

func Sum(sumArray [5]int) int {
	var sumResult int
	for i := 0; i < len(sumArray); i++ {
		sumResult += sumArray[i]
	}
	return sumResult
}
