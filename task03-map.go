package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	if len(input) == 0 {
		return []string{}
	}
	keys := make([]int, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	result = make([]string, 0, len(input))
	for i := 0; i < len(input); i++ {
		result[i] = input[keys[i]]
	}

	return result
}
