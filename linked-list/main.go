package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type LinkedList struct {
	head *Node
}

type Node struct {
	el int
	next *Node
}

func InitList(el int) *LinkedList {
	n := Node{el: el}
	return &LinkedList{head: &n}
}

func (l* LinkedList) IsEmpty() bool {
	return l.head == nil
}

func (l* LinkedList) Last() *Node {
	if l.IsEmpty() {
		return nil
	}

	it := l.head
	for ; it.next != nil; {
		it = it.next
	}
	return it
}

func (l* LinkedList) Append(el int) {
	lastElement := l.Last()
	newElement := Node{el: el, next: nil}
	lastElement.next = &newElement
}

func (l* LinkedList) Size() int {
	if l.IsEmpty() {
		return 0
	}
	it := l.head
	i := 1
	for ; it.next != nil; {
		i++
		it = it.next
	}
	return i
}

func (l* LinkedList) ToSlice() []int {
	var sl []int

	if l.IsEmpty() {
		return sl
	}

	it := l.head
	sl = append(sl, it.el)
	for ; it.next != nil; {
		it = it.next
		sl = append(sl, it.el)
	}
	return sl
}

func (l* LinkedList) Remove(index int) (err error) {
	if l.IsEmpty() {
		return errors.New("list is empty")
	}
	var prev *Node
	el := l.head
	for i := 0; i < index; i++ {
		prev = el
		el = el.next

		if el == nil {
			return errors.New("list index is out of range")
		}
	}

	if prev == nil {
		l.head = el.next
	} else {
		prev.next = el.next
	}
	return nil
}

func displayList(l* LinkedList) {
	it := l.head
	for ; it.next != nil; {
		fmt.Println(it.el)
		it = it.next
	}
	fmt.Println(it.el)
}

func main() {
	l := InitList(1)
	l.Append(5)
	l.Append(6)
	l.Append(10)
	fmt.Println(l.Size())
	fmt.Println(l.ToSlice())
	err := l.Remove(2)
	if err != nil {
		panic(err)
	}

	fmt.Println(l.ToSlice())
}
