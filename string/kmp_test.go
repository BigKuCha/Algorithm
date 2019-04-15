package string

import (
	"fmt"
	"testing"
)

func TestKmpIndex(t *testing.T) {
	type testData struct {
		src  string
		str  string
		want int
	}

	data := []testData{
		testData{src: "abcdefjflej", str: "fjf", want: 5},
		testData{src: "abaaaaaab", str: "aa", want: 2},
	}
	for _, td := range data {
		index := KmpIndex(td.src, td.str)
		if index != td.want {
			t.Error("查找错误")
		}
	}
}

func TestNext(t *testing.T) {
	type testData struct {
		in   string
		want []int
	}
	data := []testData{
		{"abcdefaa", []int{-1, 0, 0, 0, 0, 0, 0, 1}},
		{"ababaabb", []int{-1, 0, 0, 1, 2, 3, 1, 2}},
	}
	for _, d := range data {
		next := Next(d.in)
		if len(next) != len(d.want) {
			t.Error("next数组长度有误")
		}
		for k, v := range next {
			if v != d.want[k] {
				fmt.Println(k, v)
				t.Error("next方法有误")
			}
		}
	}
}
