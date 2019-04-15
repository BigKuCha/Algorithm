package string

func KmpIndex(strSrc, str string) int {
	next := Next(str)
	i := 0      // 主串从0位置开始匹配
	j := 0      // 模式串从0位置开始匹配
	index := -1 // 默认不匹配，返回-1
	for i < len(strSrc) {
		// 主串和模式串匹配成功，同时前进一位，进行下一位对比
		if strSrc[i] == str[j] {
			i++
			j++
		} else {
			// 某一位匹配失败，模式串回溯，主串不动
			j = next[j]
		}
		// 当模式串回溯到第一位时，主串需要进一位
		if j == -1 {
			i++
			j = 0
		}
		// 模式串匹配到最后一位没有回溯，说明匹配成功，返回此次主串的起始位置，也就是模式串第一次出现的位置
		if j == len(str) {
			index = i - j
			break
		}
	}
	return index
}

func Next(str string) []int {
	length := len(str)       // 模式串长度
	n := make([]int, length) // 用来存放next值
	i := 0                   // 下标
	k := -1                  // next值
	n[0] = -1
	for i < length-1 {
		if k == -1 || str[i] == str[k] {
			i++
			k++
			n[i] = k
		} else {
			k = n[k]
		}
	}
	return n
}
