package sum

func sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return
}

// AllTails returns an array with sums of elements in given arrays
func AllTails(numbersToSum ...[]int) (sums []int) {
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
