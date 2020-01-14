package iteration

const round = 5

func repeat(char string, n int) (repeated string) {
	for i := 0; i < n; i++ {
		repeated += char
	}
	return repeated
}
