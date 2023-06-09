package iteration

func Repeat(character string) string {

	var repeated = ""

	for i := 0; i < 5; i++ {
		repeated += character
	}
	return repeated
}
