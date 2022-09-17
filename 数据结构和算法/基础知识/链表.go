package main

import (
	"fmt"
	"strconv"
)

//链表可以将数据和数据之间关联起来，从一个数据指向另外一个数据。
func main() {
	test1()
	test2()
}

type LinkNode struct {
	Data     int64
	NextNode *LinkNode
}

//最简单的链表
func test1() {
	node := new(LinkNode)
	node.Data = 2

	node1 := new(LinkNode)
	node1.Data = 32
	node.NextNode = node1

	node2 := new(LinkNode)
	node2.Data = 50
	node1.NextNode = node2

	nowNode := node
	for {
		if nowNode != nil {
			fmt.Println(nowNode.Data)
			nowNode = nowNode.NextNode
		} else {
			// 如果下一个节点为空，表示链表结束了
			break
		}
	}

}

// Ring 循环链表
type Ring struct {
	next, prev *Ring       // 前驱和后驱节点
	Value      interface{} // 数据
}

//初始化循环链表

func (r *Ring) init() *Ring {
	r.prev = r
	r.next = r
	return r
}

// Next 获取下一个节点
func (r *Ring) Next() *Ring {
	if r.next ==nil {
		return r.init()
	}
	return r.next
}

// Prev 获取上一个节点
func (r *Ring) Prev() *Ring {
	if r.next==nil {
		r.init()
	}
	return r.prev
}


// New 因为绑定前驱和后驱节点为自己，没有循环，时间复杂度为：O(1)。
//创建一个指定大小 N 的循环链表，值全为空：
func New(n int ) * Ring{
	if n<=0{
		return nil
	}
	r:=new(Ring)
	p:=r
	for i := 1; i < n; i++ {
		p.Value="This node default is: "+strconv.FormatInt(int64(i),10)
		p.next=&Ring{prev: p}
		p=p.next
	}
	p.next=r
	p.prev=p
	return r

}
//Move 因为链表是循环的，当 n 为负数，表示从前面往前遍历，否则往后面遍历：
//因为需要遍历 n 次，所以时间复杂度为：O(n)。
func (r *Ring) Move(n int) *Ring {
	if r.next==nil {
		return r.init()
	}
	switch {
	case n<0:
		for; n<0;n++{
			r=r.prev
		}
	case n>0:
		for; n>0;n--{
			r=r.Next()
		}
	}
	return r;
}

// Link 往节点A，链接一个节点，并且返回之前节点A的后驱节点
func (r *Ring) Link(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
		p := s.Prev()
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}
func test2() {
fmt.Println("Hello World!")
}
