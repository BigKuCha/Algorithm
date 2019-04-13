package sort

func MergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	mid := length / 2
	left := arr[0:mid]
	right := arr[mid:]
	return merge(MergeSort(left), MergeSort(right))
}
func merge(left, right []int) []int {
	var tmpArr []int
	var i, j int
	// 在一个数组走到头之前一直比较两个数组中的元素
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			tmpArr = append(tmpArr, left[i])
			i++
		} else {
			tmpArr = append(tmpArr, right[j])
			j++
		}
	}
	// 当第一个数组到头之后，第二个数组直接并入到临时数组
	if i == len(left) {
		tmpArr = append(tmpArr, right[j:]...)
	} else {
		tmpArr = append(tmpArr, left[i:]...)
	}
	return tmpArr
}
