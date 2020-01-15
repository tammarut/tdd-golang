package arrayslice

func Sum(numbers []int) int {
	var result int
	for _, number := range numbers {
		result += number
	}
	return result
}

func SumAll(numbersToSum ...[]int) (result []int) {
	lengthOfNumbers := len(numbersToSum)
	result = make([]int, lengthOfNumbers)	
	
	for i, numbers := range numbersToSum {
		result[i] = Sum(numbers)
	}
	return result 
}
