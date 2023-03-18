package arrays

func Sum(numbers []int) (sum int) {
	for _, v := range numbers {
		sum += v
	}
	return
}

func SumAll(numbersToSum ...[]int) (result []int) {
	for _, s := range numbersToSum {
		result = append(result, Sum(s))
	}
	return
}

func SumAllTails(numbersToSum ...[]int) (result []int) {

	for _, s := range numbersToSum {
		tail := []int{0}
		if len(s) != 0 {
			tail = s[1:]
		}
		result = append(result, Sum(tail))
	}

	return
}
