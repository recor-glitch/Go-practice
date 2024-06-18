package doublyLinkList

import "fmt"

type Node struct {
	prev *Node
	next *Node
	data any
}

type LinkList struct {
	head *Node
	tail *Node
}

func (dl *LinkList) Insert(data any) {

	dlData := &Node{data: data}

	if dl.head == nil {
		dl.head = dlData
		dl.tail = dlData
		return
	}

	dlData.prev = dl.tail
	dl.tail.next = dlData
	dl.tail = dlData

}

func (dl *LinkList) InsertElementInIndex(index int, data any) {
	dlData := &Node{data: data}
	if dl.head == nil && index == 0 {
		dl.head = dlData
		return
	}
	if dl.head != nil && index == 0 {
		dlData.next = dl.head
		dl.head.prev = dlData
		dl.head = dlData
		return
	}
	current := dl.head

	count := 0
	for current.next != nil {
		if index == count {
			dlData.prev = current.prev
			dlData.next = current
			current.prev.next = dlData
			current.prev = dlData
			return
		}
		current = current.next
		count++
	}

	if count == index {
		dl.Insert(data)
	} else {
		println("Index out of bounds")
	}
}

func (dl *LinkList) Display() {
	current := dl.head

	for current != nil {
		fmt.Printf("%+v\n", current.data)
		current = current.next
	}
}
