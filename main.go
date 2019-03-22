package main

import (
	"fmt"
	"github.com/bigkucha/Algorithm/sort"
)

func main() {
	arr := []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}
	sort.QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
	//sample.Reservoir()
	//ll := others.NewLinkList()
	//fmt.Println("链表大小", ll.Size())
	//node := &others.Node{
	//	Value: "haha",
	//}
	//ll.Append(node)
	//ll.Print()
	//node2 := &others.Node{
	//	Value: map[int]string{1: "one", 2: "two"},
	//}
	//ll.InsertBefore(node2, node)
	//ll.Print()
	//
	//ll.Remove(node2)

}
