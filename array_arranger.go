package array_arranger

import (
	"strconv"
)

/*

	string arrays

*/

type StringArray []string

// ReverseString gets an array of strings and reverses its elements
func ReverseString(in []string) []string {
	// Create an output array with the same length as the input array
	output := make([]string, len(in))
	if len(in) == 0 {
		return output
	}
	for i := 0; i < len(in); i++ {
		// Reverse the order of elements
		output[i] = in[len(in)-1-i]
	}
	return output
}

// NewStringArray converts an array of strings to a StringArray
func NewStringArray(in []string) StringArray {
	return in
}

// AttachElements concatenates the elements of a StringArray into a single string
func (sa StringArray) AttachElements() string {
	output := ""
	if len(sa) == 0 {
		return output
	}
	for _, v := range sa {
		output += v
	}
	return output
}

// AttachElementsBySpace concatenates the elements of a StringArray into a single string separated by spaces
func (sa StringArray) AttachElementsBySpace() string {
	output := ""
	if len(sa) == 0 {
		return output
	}
	for _, v := range sa {
		output = output + " " + v
	}
	return output
}

// CountElements returns the number of elements in a StringArray
func (sa StringArray) CountElements() int {
	out := 0
	if len(sa) == 0 {
		return out
	}
	for range sa {
		out++
	}
	return out
}

type ElementExist struct {
	Exist bool
	Index int
}

// ElementExist checks if an element exists in a StringArray and returns its index
func (sa StringArray) ElementExist(element string) ElementExist {
	if len(sa) == 0 {
		return ElementExist{}
	}
	for i, v := range sa {
		if v == element {
			return ElementExist{
				Exist: true,
				Index: i,
			}
		}
	}
	return ElementExist{}
}

// ElementsIndexMapper maps arrays element to its index
func (sa StringArray) ElementsIndexMapper() map[int]string {
	m := make(map[int]string)
	if len(sa) == 0 {
		return m
	}
	for i, v := range sa {
		m[i] = v
	}

	return m
}

// TODO: complete this ...
func (sa StringArray) DeleteElement(element string) []string {
	if len(sa) == 0 {
		return sa
	}
	for i, v := range sa {
		if v == element {
			sa = append(sa[:i], sa[i+1:]...)
		}
	}
	return sa
}

// ShiftElements to shift to the right use minus digits and positive digits to shift to the left
func (sa StringArray) ShiftElements(n int) []string {
	output := make([]string, len(sa))

	for i := range sa {
		j := n + i
		if j >= len(sa) {
			j -= len(sa)
		} else if j < 0 {
			j += len(sa)
		}
		output[i] = sa[j]
	}

	return output
}

// AppendElementToTheEnd appends an element to the end of an string arrey
func (sa StringArray) AppendElementToTheEnd(element string) []string {
	return append(sa, element)
}

// AppendElementToTheFirst appends an element to the very first of an string arrey
func (sa StringArray) AppendElementToTheFirst(element string) []string {
	output := make([]string, len(sa)+1)
	output[0] = element
	copy(output[1:], sa)
	return output
}

// AppendSlices appends two slices
func AppendSlices(first, second []string) []string {
	output := make([]string, 0)
	output = append(output, first...)
	output = append(output, second...)
	return output
}

// ToDigits converts the string elements to digits and return the corresponding int array
func (sa StringArray) ToDigits() []int {
	output := make([]int, 0)
	for _, v := range sa {

		val, err := strconv.Atoi(v)
		if err != nil {
			return output
		}
		output = append(output, val)
	}
	return output

}

/*

	integer arrays

*/

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
