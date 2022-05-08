package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	if len(input) == 0 {
		return []string{}
	}
	keys := make([]int, len(input))
	for k := range input {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, key := range keys {
		result = append(result, input[key])
	}

	return result
}
