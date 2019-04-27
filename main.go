package main

import (
	"fmt"
	"github.com/bigkucha/Algorithm/common"
)

func main() {
	//bm := common.NewBitmap()
	bm := common.NewBitMap(1 << 32)
	var v uint64 = 1
	bm.SetBit(v, 1)
	fmt.Println(bm.GetBit(v))
	return

	//bm := common.NewBitMap(0x01 << 44)
	//_ = bm
	//bm.SetBit(55, 1)
	//fmt.Println(bm.GetBit(55))
	//s := common.NewSnowFlake(3)
	//to := time.After(time.Second)
	//wg1 := sync.WaitGroup{}
	//wg1.Add(1)
	//go func() {
	//	for {
	//		select {
	//		case <-to:
	//			wg1.Done()
	//			break
	//		default:
	//			id := s.GetID()
	//			fmt.Println(id)
	//		}
	//	}
	//}()
	//
	//wg1.Wait()
	//return

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
