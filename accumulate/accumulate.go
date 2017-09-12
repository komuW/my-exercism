package accumulate

const testVersion = 1

// Accumulate applies the given function to the given collection
func Accumulate(inputCollection []string, applyfunc func(string) string) []string {
	var outputCollection []string

	for _, v := range inputCollection {
		transformedInput := applyfunc(v)
		outputCollection = append(outputCollection, transformedInput)
	}
	return outputCollection

}
