package sum

func sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return
}

// SumAll returns an array with sums of elements in given arrays
func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, sum(numbers))
	}

	return
}

// SumAllTails returns an array with sums of tail elements in given arrays
func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, sum(tail))
		}
	}

	return
}
