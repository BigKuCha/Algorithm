package main

import (
	"fmt"
	"github.com/bigkucha/Algorithm/common"
	"os"
	"time"
)

func main() {
	timeOut := time.After(time.Second)
	for {
		select {
		case <-timeOut:
			os.Exit(0)
		default:
			id := common.Snowflake(10)
			fmt.Println(id)
		}
	}

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
