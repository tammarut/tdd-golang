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
		tail := numbers[1:]
		result = append(result, Sum(tail))
	}
	return result	
}
