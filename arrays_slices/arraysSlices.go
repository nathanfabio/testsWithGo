package arraysslices


func Sum(numbers []int) int {
	sum:= 0

	for _, num := range numbers {
		sum += num
		
	}
	return sum
}

// func SumAll(numbersToSum ...[]int) [] int {
// 	var sums []int

// 	for _, numbers := range numbersToSum {
// 		sums = append(sums, Sum(numbers))		
// 	}

// 	return sums
// }

func SumAllTheRest(numbersToSum ...[]int) [] int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			final:= numbers[1:]
			sums = append(sums, Sum(final))
		}
	}

	return sums
}