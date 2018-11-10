package sort

import (
	"fmt"
)

// 冒泡排序
//从左到右不断交换相邻逆序的元素，在一轮的循环之后，可以让未排序的最大元素上浮到右侧。
//在一轮循环中，如果没有发生交换，就说明数组已经是有序的，此时可以直接退出。
func Bubble() {
	arr := []int{1, 23, 12, 35, 69, 3, 4, 345, 41, 45, 325, 65}
	length := len(arr)
	sorted := false // 假如给定的数组是已经排好序的，有sorted变量，只需要循环2次即可排好
	num := 0
	for i := 0; i < length && !sorted; i++ {
		sorted = true
		for j := 1; j < length; j++ {
			num++
			if arr[j] < arr[j-1] {
				sorted = false
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	fmt.Printf("总共运行了 %d 次 \n", num)
	fmt.Println(arr)
}
