package sort

// 选择出数组中的最小元素，将它与数组的第一个元素交换位置。
// 再从剩下的元素中选择出最小的元素，将它与数组的第二个元素交换位置。
// 不断进行这样的操作，直到将整个数组排序。
func Select(arr []int) []int {
	length := len(arr)
	for i, _ := range arr {
		min := arr[i]
		index := -1
		for j := i + 1; j < length; j++ {
			if min > arr[j] {
				min = arr[j]
				index = j
			}
		}
		if index >= 0 {
			arr[index] = arr[i]
			arr[i] = min
		}
	}
	return arr
}
