package arrayslice

func Sum(numbers []int) int {
	var result int
	for _, number := range numbers {
		result += number
	}
	return result
}

func SumAll(numbersToSum ...[]int) (result []int) {
	for _, numbers := range numbersToSum {
		result = append(result, Sum(numbers))
	}
	return result 
}

func SumAllTails(sliceToSum ...[]int) (result []int) {
	for _, numbers := range sliceToSum {
		if len(numbers) == 0 {
			result = append(result, 0)
		} else {
		tail := numbers[1:]
		result = append(result, Sum(tail))
		}
	}
	return result	
}
