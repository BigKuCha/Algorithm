package others

import (
	"errors"
	"fmt"
)

type Node struct {
	Value interface{}
	next  *Node
}

// 单链表
type LinkList struct {
	Head *Node
}

func NewLinkList() *LinkList {
	Head := &Node{
		Value: 0,
	}
	Head.next = nil
	return &LinkList{
		Head: Head,
	}
}

// 链表大小
func (l *LinkList) Size() int {
	return l.Head.Value.(int)
}

// 链表是否为空
func (l *LinkList) IsEmpty() bool {
	if l.Head.next == nil {
		return true
	}
	return false
}

// 第一个节点
func (l *LinkList) First() *Node {
	return l.Head.next
}

// 最后一个节点
func (l *LinkList) Last() *Node {
	current := l.Head
	for {
		if current.next == nil {
			break
		}
		current = current.next
	}
	return current
}

// 插入节点到链表尾部
func (l *LinkList) Append(n *Node) *LinkList {
	last := l.Last()
	last.next = n
	l.Head.Value = l.Size() + 1
	return l
}

// 插入节点到链表头部
func (l *LinkList) PreAppend(n *Node) *Node {
	n.next = l.Head.next
	l.Head.next = n
	l.sizeInc()
	return n
}

// 指定位置前面插入节点
func (l *LinkList) InsertBefore(n, at *Node) (*Node, error) {
	p := l.Head
	for {
		if p.next == nil {
			return nil, errors.New("没有找到目标节点")
		}
		if p.next == at {
			break
		}
	}
	n.next = at
	p.next = n
	l.sizeInc()
	return n, nil
}

// 指定位置后面插入节点
func (l *LinkList) InsertAfter(n, at *Node) (*Node, error) {
	p := l.Head
	for {
		if p.next == nil {
			return nil, errors.New("没有找到目标节点")
		}
		if p == at {
			break
		}
		p = p.next
	}
	n.next = p.next
	p.next = n
	l.sizeInc()
	return n, nil
}

// 移除节点
func (l *LinkList) Remove(n *Node) (int, error) {
	current := l.Head
	for {
		if current.next == nil {
			return l.Size(), errors.New("未找到节点")
		}
		if current.next == n {
			break
		}
	}
	current.next = current.next.next
	l.sizeDec()
	return l.Size(), nil
}

// 打印链表
func (l *LinkList) Print() {
	fmt.Println("↓-----------------↓")
	fmt.Println("链表大小:", l.Size())
	i := 1
	current := l.Head.next
	for {
		fmt.Printf("第%d个元素为%+v -> \n", i, current.Value)
		if current.next == nil {
			break
		}
		current = current.next
		i++
	}
	fmt.Println("↑-----------------↑")
}

func (l *LinkList) sizeInc() {
	c := l.Head.Value.(int)
	l.Head.Value = c + 1
}

func (l *LinkList) sizeDec() {
	c := l.Head.Value.(int)
	l.Head.Value = c - 1
}
