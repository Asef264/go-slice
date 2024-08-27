package array_arranger

import (
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
