package iteration

func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += character
	}
	return repeated
}

func ExampleRepeat(character string, cycle int) string {
	var repeated string
	for i := 0; i < cycle; i++ {
		repeated += character
	}
	return repeated
}
