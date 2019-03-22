package sort

func QuickSort(arr []int, left, right int) {
	if left > right {
		return
	}
	mid := arr[left]
	low := left
	high := right
	for low < high {
		for low < high && arr[high] > mid {
			high--
		}
		arr[low] = arr[high]

		for low < high && arr[low] <= mid {
			low++
		}
		arr[high] = arr[low]
	}
	arr[low] = mid
	QuickSort(arr, left, low-1)
	QuickSort(arr, low+1, right)
}
