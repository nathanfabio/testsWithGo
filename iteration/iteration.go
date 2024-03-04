package iteration


func Repeat(char string) string {
	var repetitions string

	for i := 0; i < 5; i++ {
		repetitions += char
	}

	return repetitions
}