package sum

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbers ...[]int) int {
	lengthOfNumbers := len(numbers)
	sums := make([]int, lengthOfNumbers)

	for i, number := range numbers {
		sums[i] = Sum(number)
	}

	sumAll := 0
	for _, number := range sums {
		sumAll += number
	}

	return sumAll
}
