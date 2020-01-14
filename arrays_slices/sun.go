package arrayslice

func Sum(numbers [5]int) int{
	var result int
	for _, number := range numbers {
		result += number
	}
	return result	
}
