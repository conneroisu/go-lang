package slicesandarrays

import "testing"


func TestSum (t *testing.T) {
    numbers := []int{1, 2, 3, 4, 5}
    got := Sum(numbers)
    want := 15

    if got != want {
	t.Errorf("got %d want %d given, %v", got, want, numbers)
    }
}


func TestSumAll(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if got != nil && want != nil{
			t.Errorf("got %v want %v", got, want)
		}
		
	})

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAll([]int{1, 1}, []int{0, 0})
		want := []int{2, 0}

		if got != nil && want != nil{
			t.Errorf("got %v want %v", got, want)
		}
	})

}


