package circularLinkList

import "fmt"

type Node struct {
	next *Node
	prev *Node
	data any
}

type LinkList struct {
	head *Node
	tail *Node
}

func (cl *LinkList) Insert(data any) {

	clData := &Node{data: data}

	if cl.head == nil {
		cl.head = clData
		cl.tail = clData
		cl.head.prev = cl.tail
		cl.tail.next = cl.head

		return
	}

	clData.prev = cl.tail
	cl.tail.next = clData
	cl.tail = clData

	cl.tail.next = cl.head
	cl.head.prev = cl.tail

}

func (cl *LinkList) Display() {
	current := cl.head

	for current != nil && current.next != cl.tail {
		fmt.Printf("%+v\n", current.data)
		current = current.next
	}
}