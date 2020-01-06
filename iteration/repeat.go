package iteration

const round = 5

func repeat(char string) (repeated string) {
	for i := 0; i < round; i++ {
		repeated += char
	}
	return repeated
}
