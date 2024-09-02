package go_slice

/*
 Package go-slice provides utility functions to work with Go slices.
 Author: Khaled Moazedi
 Email: Xhmoazedi@gmail.com
*/
import (
	"reflect"
)

// Reverse reverses a slice with any type
func Reverse[T any](in []T) []T {
	out := make([]T, 0)
	for i := len(in) - 1; i >= 0; i-- {
		out = append(out, in[i])
	}
	return out
}

func TypeOf[T any](in []T) string {
	valueType := reflect.TypeOf(in[0])
	return valueType.String()
}
