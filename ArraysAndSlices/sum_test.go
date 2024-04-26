package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("want = %d but got = %d, given = %v", want, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v want %v", want, got)
	}
}

func TestSumTail(t *testing.T) {
	assertion := func(want, got []int) {
		t.Helper()
		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("two not empty arrays", func(t *testing.T) {
		got := SumTail([]int{0, 2}, []int{1, 2, 3, 9})
		want := []int{2, 9}

		assertion(want, got)
	})

	t.Run("one empty array", func(t *testing.T) {
		got := SumTail([]int{0, 2}, []int{})
		want := []int{2}
		assertion(want, got)
	})
}
