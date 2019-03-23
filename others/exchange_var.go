package others

import (
	"container/list"
)

func Arithmetic(a, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b // 直接return b,a ??
}

func Pointer() {
}

func Xor(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

func List(a, b int) (int, int) {
	lk := list.New()
	lk.PushBack(a)
	lk.PushBack(b)
	a = lk.Back().Value.(int)
	b = lk.Front().Value.(int)
	return a, b
}
