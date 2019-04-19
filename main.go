package main

import (
	"fmt"
	"github.com/bigkucha/Algorithm/common"
	"os"
	"time"
)

func main() {
	snowFlake := common.NewSnowFlake(1)
	timeout := time.After(time.Second)
	for {
		select {
		case <-timeout:
			fmt.Println("os exit")
			os.Exit(0)
		default:
			id := snowFlake.GetID()
			fmt.Println(id)
		}
	}
	return

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
