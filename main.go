package array_arranger

func ReverseString(in []string) []string {
	output := make([]string, len(in))
	for i := 0; i < len(in); i++ {
		output[i] = in[len(in)-1-i]
	}
	return output
}

func ReverseInteger(in []int) []int {
	output := make([]int, len(in))
	for i := 0; i < len(in); i++ {
		output[i] = in[len(in)-1-i]
	}
	return output
}

func ReverseInt32(in []int32) []int32 {
	output := make([]int32, len(in))
	for i := 0; i < len(in); i++ {
		output[i] = in[len(in)-1-i]
	}
	return output
}

func ReverseFloat64(in []float64) []float64 {
	output := make([]float64, len(in))
	for i := 0; i < len(in); i++ {
		output[i] = in[len(in)-1-i]
	}
	return output
}

func ReverseFloat32(in []float32) []float32 {
	output := make([]float32, len(in))
	for i := 0; i < len(in); i++ {
		output[i] = in[len(in)-1-i]
	}
	return output
}
