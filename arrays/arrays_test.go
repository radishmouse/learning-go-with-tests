package arrays

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
			t.Errorf("want %d got %d given %v", want, got, numbers)
		}
	})
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}

		got := Sum(numbers)
		want := 21
		if got != want {
			t.Errorf("expected %d but got %d given %v", want, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("returns a slice containing a sum for each slice passed in", func(t *testing.T) {
		nums1 := []int{1, 2, 3}
		nums2 := []int{4, 5, 6}
		nums3 := []int{7, 8, 9}

		got := SumAll(nums1, nums2, nums3)
		want := []int{6, 15, 24}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v but got %v", want, got)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v got %v", want, got)
		}
	}

	t.Run("sums of tails of slices", func(t *testing.T) {
		nums1 := []int{1, 2, 3}
		nums2 := []int{4, 5, 6}
		nums3 := []int{7, 8, 9}
		nums4 := []int{0, 8, 9}

		got := SumAllTails(nums1, nums2, nums3, nums4)
		want := []int{5, 11, 17, 17}
		checkSums(t, got, want)
	})
	t.Run("safely sums empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{}, []int{3, 4, 5})
		want := []int{0, 0, 9}
		checkSums(t, got, want)
	})
}
