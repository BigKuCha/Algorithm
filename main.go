package main

import (
	"fmt"
	"github.com/bigkucha/Algorithm/common"
	"sync"
)

func main() {
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

	snow1 := common.NewSnowFlake(1)

	wg := sync.WaitGroup{}
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			id := snow1.GetID()
			fmt.Println(id)
			wg.Done()
		}()
	}
	wg.Wait()

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
