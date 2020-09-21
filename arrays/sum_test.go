package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got []int, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := AllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := AllTails([]int{}, []int{}, []int{3, 4, 5})
		want := []int{0, 0, 9}
		checkSums(t, got, want)
	})
}

func BenchmarkSumAllTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AllTails([]int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3})
	}
}
