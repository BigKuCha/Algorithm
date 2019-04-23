package others

import (
	"strings"
)

func YangHuiTriangle() {
	var height = 11
	arr := map[int][]int{}
	for i := 1; i <= height; i++ {
		leftSpace := ((2*height - 1) - (2*i - 1)) / 2
		for j := 1; j <= i; j++ {
			// 第一个元素 前面输出n个空格
			if j == 1 {
				print(strings.Repeat(" ", leftSpace))
			}
			// 最后一个元素输出1后回车
			if j == i {
				print(1, "\n")
				arr[i] = append(arr[i], 1)
				continue
			}
			// 第一个元素输出1
			if j == 1 {
				print(1, " ")
				arr[i] = append(arr[i], 1)
				continue
			}

			lastArr := arr[i-1]
			left := lastArr[j-2]
			right := lastArr[j-1]
			out := left + right
			print(out, " ")
			arr[i] = append(arr[i], out)
		}
	}
}
