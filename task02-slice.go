package homework

func reverse(input []int64) (result []int64) {
	if len(input) == 0 {
		return input
	}

	for i := len(input); i > -1; i-- {
		result = append(result, input[i])
	}

	return result
}
