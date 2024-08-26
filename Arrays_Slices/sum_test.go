package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := sum(numbers)
		expect := 15

		if got != expect {
			t.Errorf("got %d want %d given, %v", got, expect, numbers)
		}
	})

	t.Run("Collection any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := sum(numbers)
		expect := 6

		if got != expect {
			t.Errorf("got %d want %d given, %v", got, expect, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := sumAll([]int{1, 2}, []int{0, 9}, []int{1, 2, 3}, []int{3, 2, 1}, []int{1, 2}, []int{0, 9}, []int{1, 2, 3}, []int{3, 2, 1}, []int{1, 2}, []int{0, 9}, []int{1, 2, 3}, []int{3, 2, 1}, []int{1, 2}, []int{0, 9}, []int{1, 2, 3}, []int{3, 2, 1})
	want := []int{3, 9, 6, 6, 3, 9, 6, 6, 3, 9, 6, 6, 3, 9, 6, 6}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v want %v", got, want)
	}
}

func BenchmarkRepeatSumAllBenchmarking(b *testing.B) {
	sumAll([]int{1, 2}, []int{0, 9}, []int{1, 2, 3}, []int{3, 2, 1}, []int{1, 2}, []int{0, 9}, []int{1, 2, 3}, []int{3, 2, 1}, []int{1, 2}, []int{0, 9}, []int{1, 2, 3}, []int{3, 2, 1}, []int{1, 2}, []int{0, 9}, []int{1, 2, 3}, []int{3, 2, 1})
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v want %v", got, want)
	}
}
