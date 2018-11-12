package sort

import "testing"

func TestBubble(t *testing.T) {
	type testData struct {
		in   []int
		want []int
	}
	data := []testData{
		{[]int{4, 5, 3, 7, 8, 1, 9, 0, 2, 6}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{11, 23, 78, 9, 82, 76, 90, 12}, []int{9, 11, 12, 23, 76, 78, 82, 90}},
	}
	for _, d := range data {
		rzt := Bubble(d.in)
		for i, v := range rzt {
			if v != d.want[i] {
				t.Error("冒泡排序，排序有误")
			}

			if v < (len(rzt) - 1) {
				vAfter := rzt[i+1]
				if v > vAfter {
					t.Error("冒泡排序，排序有误")
				}
			}
		}
	}
}
