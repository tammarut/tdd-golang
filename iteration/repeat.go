package iteration

func repeat(char string) (repeated string) {
	for i := 0; i<5; i++{
		repeated += char
	}
	return repeated
}
