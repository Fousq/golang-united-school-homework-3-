package homework

func reverse(input []int64) (result []int64) {
	if input == nil || len(input) == 0 {
		return input
	}
	reversedInput := make([]int64, len(input))

	copy(reversedInput, input)

	for i := 0; i < len(reversedInput)-1; i++ {
		for j := i + 1; j < len(reversedInput); j++ {
			reversedInput[i], reversedInput[j] = reversedInput[j], reversedInput[i]
		}
	}

	return reversedInput
}
