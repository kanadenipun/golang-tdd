package iteration

func Repeat(toBeRepeated string) string {
	var resultString string
	for i := 0; i < 5; i++ {
		resultString += toBeRepeated
	}
	return resultString
}
