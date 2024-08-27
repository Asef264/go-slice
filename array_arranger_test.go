package array_arranger

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShiftElements_ZiroShift(t *testing.T) {
	//ZiroShiftTest
	sa := []string{"khaled", "hamzah", "saeed", "sara"}
	arg := 0
	want := sa
	testy := NewStringArray(sa)
	result := testy.ShiftElements(arg)
	assert.Equal(t, want, result)
}
func TestShiftElements_OneShiftToTheRight(t *testing.T) {
	//OneShiftToTheRight
	sa1 := []string{"khaled", "hamzah", "saeed", "sara"}
	arg1 := -1
	want1 := []string{"sara", "khaled", "hamzah", "saeed"}
	testy1 := NewStringArray(sa1)
	result1 := testy1.ShiftElements(arg1)
	assert.Equal(t, want1, result1)
}

func TestDeleteOneElement(t *testing.T) {
	//DeleteOneElement
	Element := "khaled"
	sa := []string{"khaled", "hamzah", "saeed", "sara"}
	want1 := []string{"hamzah", "saeed", "sara"}
	testy := NewStringArray(sa)
	result := testy.DeleteElement(Element)
	assert.Equal(t, want1, result)
}

func TestDeleteTwoElements(t *testing.T) {
	//DeleteOneElement
	Element := "khaled"
	sa := []string{"khaled", "hamzah", "saeed", "khaled", "sara"}
	want1 := []string{"hamzah", "saeed", "sara"}
	testy := NewStringArray(sa)
	result := testy.DeleteElement(Element)
	assert.Equal(t, want1, result)
}

func TestAppendElementToTheFirst(t *testing.T) {
	Element := "pari"
	sa := []string{"khaled", "hamzah", "saeed", "khaled", "sara"}
	want := []string{"pari", "khaled", "hamzah", "saeed", "khaled", "sara"}
	testy := NewStringArray(sa)
	result := testy.AppendElementToTheFirst(Element)
	assert.Equal(t, want, result)
}

func TestAppendSlices(t *testing.T) {
	first := []string{"khaled", "hamzah", "saeed", "khaled", "sara"}
	second := []string{"man", "to"}
	want1 := []string{"khaled", "hamzah", "saeed", "khaled", "sara", "man", "to"}
	result := AppendSlices(first, second)
	assert.Equal(t, want1, result)
}

func TestToDigits(t *testing.T) {
	first := []string{"0", "43", "3", "2", "1"}
	want1 := []int{0, 43, 3, 2, 1}
	testy := NewStringArray(first)
	result := testy.ToDigits()
	assert.Equal(t, want1, result)
}

func TestToDigitsError(t *testing.T) {
	first := []string{"0", "error", "3", "2", "1"}
	testy := NewStringArray(first)
	result := testy.ToDigits()
	assert.Error(t, errors.ErrUnsupported, result)
}
