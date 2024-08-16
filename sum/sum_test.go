package sum

import "testing"

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d, want %d, given %d", got, want, numbers)
		}
	})

	t.Run("collection of any size 2", func(t *testing.T) {
		numbers := []int{5, 6, 7, 23, 45}

		got := Sum(numbers)
		want := 86

		if got != want {
			t.Errorf("got %d, want %d, given %d", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9}, []int{3, 5, 9}, []int{7, 11, 2})
	want := 49

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

}
