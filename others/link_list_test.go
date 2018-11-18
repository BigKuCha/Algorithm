package others

import (
	"testing"
)

var ll *LinkList

func init() {
	ll = NewLinkList()
	nodes := []*Node{
		&Node{Value: "No1"},
		&Node{Value: "No2"},
		&Node{Value: "No3"},
		&Node{Value: "No4"},
	}
	for _, node := range nodes {
		ll.Append(node)
	}
}

func TestNewLinkList(t *testing.T) {
	ll := NewLinkList()
	if !ll.IsEmpty() {
		t.Error("判断链表为空方法出错")
	}
	node := &Node{
		Value: "No1",
	}
	ll.Append(node)
	if ll.Size() != 1 {
		t.Error("链表大小方法出错")
	}
	node2 := &Node{
		Value: "No2",
	}
	ll.InsertBefore(node2, node)

	if node2.next != node {
		t.Error("InserBefore 方法出错")
	}
	ll.Remove(node2)
	if ll.Size() != 1 || node.next != nil {
		t.Error("Revemo 方法出错")
	}

	if ll.First() != ll.Head.next {
		t.Error("First 方法出错")
	}

	node3 := &Node{
		Value: "No3",
	}
	ll.Append(node3)
	node4 := &Node{
		Value: "No4",
	}
	ll.InsertAfter(node4, node)
	if ll.Size() != 3 || node.next != node4 || node4.next != node3 {
		t.Error("InsertAfter 方法出错")
	}

	if ll.Last() != node3 {
		t.Error("last 方法出错")
	}

	node5 := &Node{
		Value: "No5",
	}
	ll.PreAppend(node5)
	if ll.Head.next != node5 || node5.next != node {
		t.Error("")
	}
	ll.Print()
}
