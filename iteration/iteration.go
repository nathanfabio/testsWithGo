package iteration



func Repeat(char string, num int) string {
	var repetitions string

	for i := 0; i < num; i++ {
		repetitions += char
	}

	return repetitions
}