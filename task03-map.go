package homework

func sortMapValues(input map[int]string) (result []string) {
	if len(input) == 0 {
		return nil
	}
	keys := make([]int, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}
	for i := 0; i < len(keys)-1; i++ {
		for j := i + 1; j < len(keys); j++ {
			if keys[i] > keys[j] {
				keys[i], keys[j] = keys[j], keys[i]
			}
		}
	}
	result = make([]string, 0, len(input))
	for i := 0; i < len(input); i++ {
		result[i] = input[keys[i]]
	}
	return result
}
