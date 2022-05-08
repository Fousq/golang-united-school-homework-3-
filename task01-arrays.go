package homework

func average(input [15]float32) (result float32) {
	if len(input) == 0 {
		return 0.0
	}
	var sum float32 = 0.0
	for i := 0; i < len(input); i++ {
		sum += input[i]
	}
	return sum / float32(len(input))
}
