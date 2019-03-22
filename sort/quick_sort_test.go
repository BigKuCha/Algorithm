package sort

import "testing"

func TestQuickSort(t *testing.T) {
	type testData struct {
		in   []int
		want []int
	}
	data := []testData{
		{[]int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{30, 20, 99, 76, 54, 55, 55, 77, 88}, []int{20, 30, 54, 55, 55, 76, 77, 88, 99}},
	}
	for _, d := range data {
		QuickSort(d.in, 0, len(d.in)-1)
		for k, v := range d.in {
			if v != d.want[k] {
				t.Error("快速排序错误")
			}
		}
	}
}
