package sort

import "testing"

func TestMergeSort(t *testing.T) {
	type testData struct {
		in   []int
		want []int
	}
	data := []testData{
		{[]int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{30, 20, 99, 76, 54, 55, 55, 77, 88}, []int{20, 30, 54, 55, 55, 76, 77, 88, 99}},
		{[]int{1}, []int{1}},
	}
	for _, d := range data {
		var s []int
		s = MergeSort(d.in)
		for k, v := range s {
			if v != d.want[k] {
				t.Error("归并排序错误")
			}
		}
	}
}
