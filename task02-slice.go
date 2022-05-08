package homework

func reverse(input []int64) (result []int64) {
	if len(input) == 0 {
		return input
	}
	reversedInput := make([]int64, len(input))

	copy(reversedInput, input)

	for i := 0; i < len(reversedInput)/2; i++ {
		j := len(reversedInput) - i
		reversedInput[i], reversedInput[j] = reversedInput[j], reversedInput[i]
	}

	return reversedInput
}
