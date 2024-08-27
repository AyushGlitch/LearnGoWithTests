package arrays

import "testing"

func TestSlice (t *testing.T) {
	nums := []int {1, 2, 3, 4, 5}

	got := sliceSum(nums)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, nums)
	}
}