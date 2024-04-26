package arraysandslices

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, v := range numbersToSum {
		sums = append(sums, Sum(v))
	}
	return sums
}

func SumTail(numbersToSum ...[]int) (result []int) {
	for _, value := range numbersToSum {
		var tail int
		if len(value) > 0 {
			tail = len(value) - 1
			result = append(result, value[tail])
		}
	}
	return
}
