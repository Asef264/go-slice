package go_slice

/*
 Package go-slice provides utility functions to work with Go slices.
 Author: Khaled Moazedi
 Email: Xhmoazedi@gmail.com
*/

/*
digit slices
*/
type IntSlice []int

func NewIntSlice(in []int) IntSlice {
	return in
}

// ReverseInteger reverses the elements in an array of integers
func ReverseInteger(in []int) []int {
	// Create an output array with the same length as the input array
	output := make([]int, len(in))
	for i := 0; i < len(in); i++ {
		// Reverse the order of elements
		output[i] = in[len(in)-1-i]
	}
	return output
}

// ReverseInt32 reverses the elements in an array of int32 numbers
func ReverseInt32(in []int32) []int32 {
	// Create an output array with the same length as the input array
	output := make([]int32, len(in))
	for i := 0; i < len(in); i++ {
		// Reverse the order of elements
		output[i] = in[len(in)-1-i]
	}
	return output
}

// ReverseFloat64 reverses the elements in an array of float64 numbers
func ReverseFloat64(in []float64) []float64 {
	// Create an output array with the same length as the input array
	output := make([]float64, len(in))
	for i := 0; i < len(in); i++ {
		// Reverse the order of elements
		output[i] = in[len(in)-1-i]
	}
	return output
}

// ReverseFloat32 reverses the elements in an array of float32 numbers
func ReverseFloat32(in []float32) []float32 {
	output := make([]float32, len(in))
	for i := 0; i < len(in); i++ {
		output[i] = in[len(in)-1-i]
	}
	return output
}

// SumInt gets a slice in type int and returns a summerize of all elements in int64
func (is IntSlice) SumInt() int64 {
	out := make([]int64, 0)
	for _, v := range is {
		out = append(out, int64(v))
	}
	var sum int64
	for _, item := range out {
		sum += item
	}
	return sum
}
