package go_slice

import "sort"

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

// SumInt gets a slice in type int and returns a summerize of all elements in int64
func (is IntSlice) SumInt() int64 {
	out := ToInt64(is)
	var sum int64
	for _, item := range out {
		sum += item
	}
	return sum
}

func ToInt64(in []int) []int64 {
	out := make([]int64, 0)
	for _, v := range in {
		out = append(out, int64(v))
	}
	return out
}

// Average comes on a slice in int type and calculates the average of elements
func (is IntSlice) Average() float64 {
	return float64(is.SumInt()) / float64(len(is))
}

// Sorts a slice of integers in ascending order
func SortAsc(slice []int) {
	sort.Ints(slice)
}

// Sorts a slice of integers in descending order
func SortDesc(slice []int) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})
}
